package inmemory

import (
	"context"
	"errors"
	"reflect"
	"sync"

	goengine_dev "github.com/hellofresh/goengine-dev"
	"github.com/hellofresh/goengine/log"
	"github.com/hellofresh/goengine/metadata"
)

var (
	// ErrStreamExistsAlready occurs when create is called for an already created stream
	ErrStreamExistsAlready = errors.New("stream already exists")
	// ErrStreamNotFound occurs when an unknown streamName is provided
	ErrStreamNotFound = errors.New("unknown stream")
	// ErrNilMessage occurs when a goengine_dev.Message that is being appended to a stream is nil or a reference to nil
	ErrNilMessage = errors.New("nil is not a valid message")
	// Ensure that we satisfy the eventstore.EventStore interface
	_ goengine_dev.EventStore = &EventStore{}
)

// EventStore a in memory event store implementation
type EventStore struct {
	sync.RWMutex

	logger  log.Logger
	streams map[goengine_dev.StreamName][]goengine_dev.Message
}

// NewEventStore return a new inmemory.EventStore
func NewEventStore(logger log.Logger) *EventStore {
	return &EventStore{
		logger:  logger,
		streams: map[goengine_dev.StreamName][]goengine_dev.Message{},
	}
}

// Create creates a event stream
func (i *EventStore) Create(ctx context.Context, streamName goengine_dev.StreamName) error {
	if _, found := i.streams[streamName]; found {
		return ErrStreamExistsAlready
	}

	i.streams[streamName] = []goengine_dev.Message{}

	return nil
}

// HasStream returns true if the stream exists
func (i *EventStore) HasStream(ctx context.Context, streamName goengine_dev.StreamName) bool {
	_, found := i.streams[streamName]

	return found
}

// Load returns a list of events based on the provided conditions
func (i *EventStore) Load(
	ctx context.Context,
	streamName goengine_dev.StreamName,
	fromNumber int64,
	count *uint,
	matcher metadata.Matcher,
) (goengine_dev.EventStream, error) {
	i.RLock()
	defer i.RUnlock()

	storedEvents, knownStream := i.streams[streamName]
	if !knownStream {
		return nil, ErrStreamNotFound
	}

	metadataMatcher, err := NewMetadataMatcher(matcher, i.logger)
	if err != nil {
		return nil, err
	}

	var messages []goengine_dev.Message
	var messageNumbers []int64
	var found uint

	for idx, event := range storedEvents {
		messageNumber := int64(idx + 1)
		if messageNumber >= fromNumber && metadataMatcher.Matches(event.Metadata()) {
			found++
			messages = append(messages, event)
			messageNumbers = append(messageNumbers, messageNumber)
			if count != nil && found == *count {
				break
			}
		}
	}

	return NewEventStream(messages, messageNumbers)
}

// AppendTo appends the provided messages to the stream
func (i *EventStore) AppendTo(ctx context.Context, streamName goengine_dev.StreamName, streamEvents []goengine_dev.Message) error {
	i.Lock()
	defer i.Unlock()

	storedEvents, knownStream := i.streams[streamName]
	if !knownStream {
		return ErrStreamNotFound
	}

	for _, msg := range streamEvents {
		if msg == nil || reflect.ValueOf(msg).IsNil() {
			return ErrNilMessage
		}
	}

	storedEventCount := len(storedEvents)

	eventsToStore := make([]goengine_dev.Message, storedEventCount, storedEventCount+len(streamEvents))
	copy(eventsToStore, storedEvents)
	i.streams[streamName] = append(eventsToStore, streamEvents...)

	return nil
}
