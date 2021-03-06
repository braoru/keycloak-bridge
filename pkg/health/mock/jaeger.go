// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/keycloak-bridge/pkg/health (interfaces: JaegerModule,SystemDConn)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	health "github.com/cloudtrust/keycloak-bridge/pkg/health"
	dbus "github.com/coreos/go-systemd/dbus"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// JaegerModule is a mock of JaegerModule interface
type JaegerModule struct {
	ctrl     *gomock.Controller
	recorder *JaegerModuleMockRecorder
}

// JaegerModuleMockRecorder is the mock recorder for JaegerModule
type JaegerModuleMockRecorder struct {
	mock *JaegerModule
}

// NewJaegerModule creates a new mock instance
func NewJaegerModule(ctrl *gomock.Controller) *JaegerModule {
	mock := &JaegerModule{ctrl: ctrl}
	mock.recorder = &JaegerModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *JaegerModule) EXPECT() *JaegerModuleMockRecorder {
	return m.recorder
}

// HealthChecks mocks base method
func (m *JaegerModule) HealthChecks(arg0 context.Context) []health.JaegerReport {
	ret := m.ctrl.Call(m, "HealthChecks", arg0)
	ret0, _ := ret[0].([]health.JaegerReport)
	return ret0
}

// HealthChecks indicates an expected call of HealthChecks
func (mr *JaegerModuleMockRecorder) HealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthChecks", reflect.TypeOf((*JaegerModule)(nil).HealthChecks), arg0)
}

// SystemDConn is a mock of SystemDConn interface
type SystemDConn struct {
	ctrl     *gomock.Controller
	recorder *SystemDConnMockRecorder
}

// SystemDConnMockRecorder is the mock recorder for SystemDConn
type SystemDConnMockRecorder struct {
	mock *SystemDConn
}

// NewSystemDConn creates a new mock instance
func NewSystemDConn(ctrl *gomock.Controller) *SystemDConn {
	mock := &SystemDConn{ctrl: ctrl}
	mock.recorder = &SystemDConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SystemDConn) EXPECT() *SystemDConnMockRecorder {
	return m.recorder
}

// ListUnitsByNames mocks base method
func (m *SystemDConn) ListUnitsByNames(arg0 []string) ([]dbus.UnitStatus, error) {
	ret := m.ctrl.Call(m, "ListUnitsByNames", arg0)
	ret0, _ := ret[0].([]dbus.UnitStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnitsByNames indicates an expected call of ListUnitsByNames
func (mr *SystemDConnMockRecorder) ListUnitsByNames(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnitsByNames", reflect.TypeOf((*SystemDConn)(nil).ListUnitsByNames), arg0)
}
