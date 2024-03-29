// Code generated by MockGen. DO NOT EDIT.
// Source: chart.go

// Package warehouse_test is a generated GoMock package.
package warehouse_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChart is a mock of Chart interface.
type MockChart struct {
	ctrl     *gomock.Controller
	recorder *MockChartMockRecorder
}

// MockChartMockRecorder is the mock recorder for MockChart.
type MockChartMockRecorder struct {
	mock *MockChart
}

// NewMockChart creates a new mock instance.
func NewMockChart(ctrl *gomock.Controller) *MockChart {
	mock := &MockChart{ctrl: ctrl}
	mock.recorder = &MockChartMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChart) EXPECT() *MockChartMockRecorder {
	return m.recorder
}

// Notify mocks base method.
func (m *MockChart) Notify(artist, title string, copies uint) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", artist, title, copies)
}

// Notify indicates an expected call of Notify.
func (mr *MockChartMockRecorder) Notify(artist, title, copies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockChart)(nil).Notify), artist, title, copies)
}
