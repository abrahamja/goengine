// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hellofresh/goengine/driver/sql (interfaces: PersistenceStrategy)

// Package sql is a generated GoMock package.
package sql

import (
	gomock "github.com/golang/mock/gomock"
	goengine "github.com/hellofresh/goengine"
	reflect "reflect"
)

// PersistenceStrategy is a mock of PersistenceStrategy interface
type PersistenceStrategy struct {
	ctrl     *gomock.Controller
	recorder *PersistenceStrategyMockRecorder
}

// PersistenceStrategyMockRecorder is the mock recorder for PersistenceStrategy
type PersistenceStrategyMockRecorder struct {
	mock *PersistenceStrategy
}

// NewPersistenceStrategy creates a new mock instance
func NewPersistenceStrategy(ctrl *gomock.Controller) *PersistenceStrategy {
	mock := &PersistenceStrategy{ctrl: ctrl}
	mock.recorder = &PersistenceStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *PersistenceStrategy) EXPECT() *PersistenceStrategyMockRecorder {
	return m.recorder
}

// ColumnNames mocks base method
func (m *PersistenceStrategy) ColumnNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ColumnNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ColumnNames indicates an expected call of ColumnNames
func (mr *PersistenceStrategyMockRecorder) ColumnNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColumnNames", reflect.TypeOf((*PersistenceStrategy)(nil).ColumnNames))
}

// CreateSchema mocks base method
func (m *PersistenceStrategy) CreateSchema(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchema", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// CreateSchema indicates an expected call of CreateSchema
func (mr *PersistenceStrategyMockRecorder) CreateSchema(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchema", reflect.TypeOf((*PersistenceStrategy)(nil).CreateSchema), arg0)
}

// GenerateTableName mocks base method
func (m *PersistenceStrategy) GenerateTableName(arg0 goengine.StreamName) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTableName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTableName indicates an expected call of GenerateTableName
func (mr *PersistenceStrategyMockRecorder) GenerateTableName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTableName", reflect.TypeOf((*PersistenceStrategy)(nil).GenerateTableName), arg0)
}

// PrepareData mocks base method
func (m *PersistenceStrategy) PrepareData(arg0 []goengine.Message) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareData", arg0)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareData indicates an expected call of PrepareData
func (mr *PersistenceStrategyMockRecorder) PrepareData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareData", reflect.TypeOf((*PersistenceStrategy)(nil).PrepareData), arg0)
}