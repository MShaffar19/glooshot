// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/api_snapshot_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/glooshot/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockApiEmitter is a mock of ApiEmitter interface
type MockApiEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockApiEmitterMockRecorder
}

// MockApiEmitterMockRecorder is the mock recorder for MockApiEmitter
type MockApiEmitterMockRecorder struct {
	mock *MockApiEmitter
}

// NewMockApiEmitter creates a new mock instance
func NewMockApiEmitter(ctrl *gomock.Controller) *MockApiEmitter {
	mock := &MockApiEmitter{ctrl: ctrl}
	mock.recorder = &MockApiEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiEmitter) EXPECT() *MockApiEmitterMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockApiEmitter) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockApiEmitterMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockApiEmitter)(nil).Register))
}

// Experiment mocks base method
func (m *MockApiEmitter) Experiment() v1.ExperimentClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Experiment")
	ret0, _ := ret[0].(v1.ExperimentClient)
	return ret0
}

// Experiment indicates an expected call of Experiment
func (mr *MockApiEmitterMockRecorder) Experiment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Experiment", reflect.TypeOf((*MockApiEmitter)(nil).Experiment))
}

// Snapshots mocks base method
func (m *MockApiEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *v1.ApiSnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", watchNamespaces, opts)
	ret0, _ := ret[0].(<-chan *v1.ApiSnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockApiEmitterMockRecorder) Snapshots(watchNamespaces, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockApiEmitter)(nil).Snapshots), watchNamespaces, opts)
}