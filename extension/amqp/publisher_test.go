package amqp_test

import (
	"context"
	"testing"
	"time"

	"github.com/hellofresh/goengine"
	"github.com/hellofresh/goengine/driver/sql"
	goengineAmqp "github.com/hellofresh/goengine/extension/amqp"
	goengineLogger "github.com/hellofresh/goengine/extension/logrus"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotificationPublisher_Publish(t *testing.T) {

	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second)
	defer ctxCancel()

	channel := &mockChannel{}
	connection := &mockConnection{}

	t.Run("Invalid arguments", func(t *testing.T) {
		logger, _ := getLogger()

		_, err := goengineAmqp.NewNotificationPublisher("http://localhost:5672/", "my-queue", logger, connection, channel)
		assert.Equal(t, goengine.InvalidArgumentError("amqpDSN"), err)

		_, err = goengineAmqp.NewNotificationPublisher("amqp://localhost:5672/", "", logger, connection, channel)
		assert.Equal(t, goengine.InvalidArgumentError("queue"), err)

	})

	t.Run("Publish Nil Notification Message", func(t *testing.T) {
		ensure := require.New(t)
		logger, loggerHook := getLogger()

		publisher, err := goengineAmqp.NewNotificationPublisher("amqp://localhost:5672/", "my-queue", logger, connection, channel)
		ensure.NoError(err)
		err = publisher.Publish(ctx, nil)
		ensure.Nil(err)
		ensure.Len(loggerHook.Entries, 1)
		ensure.Equal("unable to handle nil notification, skipping", loggerHook.LastEntry().Message)
	})

	t.Run("Publish Message", func(t *testing.T) {
		ensure := require.New(t)
		logger, loggerHook := getLogger()

		publisher, err := goengineAmqp.NewNotificationPublisher("amqp://localhost:5672/", "my-queue", logger, connection, channel)
		ensure.NoError(err)

		err = publisher.Publish(ctx, &sql.ProjectionNotification{No: 1, AggregateID: "8150276e-34fe-49d9-aeae-a35af0040a4f"})

		ensure.NoError(err)
		ensure.Len(loggerHook.Entries, 0)
	})
}

func getLogger() (goengine.Logger, *test.Hook) {
	logger, loggerHook := test.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	return goengineLogger.Wrap(logger), loggerHook
}
