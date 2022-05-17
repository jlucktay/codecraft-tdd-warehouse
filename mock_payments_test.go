// Code generated by MockGen. DO NOT EDIT.
// Source: payments.go

// Package warehouse_test is a generated GoMock package.
package warehouse_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockpaymentProvider is a mock of paymentProvider interface.
type MockpaymentProvider struct {
	ctrl     *gomock.Controller
	recorder *MockpaymentProviderMockRecorder
}

// MockpaymentProviderMockRecorder is the mock recorder for MockpaymentProvider.
type MockpaymentProviderMockRecorder struct {
	mock *MockpaymentProvider
}

// NewMockpaymentProvider creates a new mock instance.
func NewMockpaymentProvider(ctrl *gomock.Controller) *MockpaymentProvider {
	mock := &MockpaymentProvider{ctrl: ctrl}
	mock.recorder = &MockpaymentProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpaymentProvider) EXPECT() *MockpaymentProviderMockRecorder {
	return m.recorder
}

// Refund mocks base method.
func (m *MockpaymentProvider) Refund(price float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refund", price)
	ret0, _ := ret[0].(error)
	return ret0
}

// Refund indicates an expected call of Refund.
func (mr *MockpaymentProviderMockRecorder) Refund(price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refund", reflect.TypeOf((*MockpaymentProvider)(nil).Refund), price)
}

// Sell mocks base method.
func (m *MockpaymentProvider) Sell(price float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sell", price)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sell indicates an expected call of Sell.
func (mr *MockpaymentProviderMockRecorder) Sell(price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sell", reflect.TypeOf((*MockpaymentProvider)(nil).Sell), price)
}