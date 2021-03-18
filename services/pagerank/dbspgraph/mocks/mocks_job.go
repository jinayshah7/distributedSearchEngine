// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jinayshah7/distributedSearchEngine/services/pagerank/dbspgraph/job (interfaces: Runner)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	bspgraph "github.com/jinayshah7/distributedSearchEngine/services/pagerank/bspgraph"
	job "github.com/jinayshah7/distributedSearchEngine/services/pagerank/dbspgraph/job"
	reflect "reflect"
)

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// AbortJob mocks base method
func (m *MockRunner) AbortJob(arg0 job.Details) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AbortJob", arg0)
}

// AbortJob indicates an expected call of AbortJob
func (mr *MockRunnerMockRecorder) AbortJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortJob", reflect.TypeOf((*MockRunner)(nil).AbortJob), arg0)
}

// CompleteJob mocks base method
func (m *MockRunner) CompleteJob(arg0 job.Details) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteJob", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteJob indicates an expected call of CompleteJob
func (mr *MockRunnerMockRecorder) CompleteJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteJob", reflect.TypeOf((*MockRunner)(nil).CompleteJob), arg0)
}

// StartJob mocks base method
func (m *MockRunner) StartJob(arg0 job.Details, arg1 bspgraph.ExecutorFactory) (*bspgraph.Executor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartJob", arg0, arg1)
	ret0, _ := ret[0].(*bspgraph.Executor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartJob indicates an expected call of StartJob
func (mr *MockRunnerMockRecorder) StartJob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartJob", reflect.TypeOf((*MockRunner)(nil).StartJob), arg0, arg1)
}
