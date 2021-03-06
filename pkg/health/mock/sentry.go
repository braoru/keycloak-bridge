// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/keycloak-bridge/pkg/health (interfaces: SentryModule,Sentry)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	health "github.com/cloudtrust/keycloak-bridge/pkg/health"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// SentryModule is a mock of SentryModule interface
type SentryModule struct {
	ctrl     *gomock.Controller
	recorder *SentryModuleMockRecorder
}

// SentryModuleMockRecorder is the mock recorder for SentryModule
type SentryModuleMockRecorder struct {
	mock *SentryModule
}

// NewSentryModule creates a new mock instance
func NewSentryModule(ctrl *gomock.Controller) *SentryModule {
	mock := &SentryModule{ctrl: ctrl}
	mock.recorder = &SentryModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SentryModule) EXPECT() *SentryModuleMockRecorder {
	return m.recorder
}

// HealthChecks mocks base method
func (m *SentryModule) HealthChecks(arg0 context.Context) []health.SentryReport {
	ret := m.ctrl.Call(m, "HealthChecks", arg0)
	ret0, _ := ret[0].([]health.SentryReport)
	return ret0
}

// HealthChecks indicates an expected call of HealthChecks
func (mr *SentryModuleMockRecorder) HealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthChecks", reflect.TypeOf((*SentryModule)(nil).HealthChecks), arg0)
}

// Sentry is a mock of Sentry interface
type Sentry struct {
	ctrl     *gomock.Controller
	recorder *SentryMockRecorder
}

// SentryMockRecorder is the mock recorder for Sentry
type SentryMockRecorder struct {
	mock *Sentry
}

// NewSentry creates a new mock instance
func NewSentry(ctrl *gomock.Controller) *Sentry {
	mock := &Sentry{ctrl: ctrl}
	mock.recorder = &SentryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Sentry) EXPECT() *SentryMockRecorder {
	return m.recorder
}

// URL mocks base method
func (m *Sentry) URL() string {
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL
func (mr *SentryMockRecorder) URL() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*Sentry)(nil).URL))
}
