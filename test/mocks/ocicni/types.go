// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cri-o/ocicni/pkg/ocicni (interfaces: CNIPlugin)

// Package ocicnitypesmock is a generated GoMock package.
package ocicnitypesmock

import (
	context "context"
	reflect "reflect"

	ocicni "github.com/cri-o/ocicni/pkg/ocicni"
	gomock "github.com/golang/mock/gomock"
)

// MockCNIPlugin is a mock of CNIPlugin interface.
type MockCNIPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockCNIPluginMockRecorder
}

// MockCNIPluginMockRecorder is the mock recorder for MockCNIPlugin.
type MockCNIPluginMockRecorder struct {
	mock *MockCNIPlugin
}

// NewMockCNIPlugin creates a new mock instance.
func NewMockCNIPlugin(ctrl *gomock.Controller) *MockCNIPlugin {
	mock := &MockCNIPlugin{ctrl: ctrl}
	mock.recorder = &MockCNIPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCNIPlugin) EXPECT() *MockCNIPluginMockRecorder {
	return m.recorder
}

// GetDefaultNetworkName mocks base method.
func (m *MockCNIPlugin) GetDefaultNetworkName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultNetworkName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDefaultNetworkName indicates an expected call of GetDefaultNetworkName.
func (mr *MockCNIPluginMockRecorder) GetDefaultNetworkName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultNetworkName", reflect.TypeOf((*MockCNIPlugin)(nil).GetDefaultNetworkName))
}

// GetPodNetworkStatus mocks base method.
func (m *MockCNIPlugin) GetPodNetworkStatus(arg0 ocicni.PodNetwork) ([]ocicni.NetResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodNetworkStatus", arg0)
	ret0, _ := ret[0].([]ocicni.NetResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodNetworkStatus indicates an expected call of GetPodNetworkStatus.
func (mr *MockCNIPluginMockRecorder) GetPodNetworkStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodNetworkStatus", reflect.TypeOf((*MockCNIPlugin)(nil).GetPodNetworkStatus), arg0)
}

// GetPodNetworkStatusWithContext mocks base method.
func (m *MockCNIPlugin) GetPodNetworkStatusWithContext(arg0 context.Context, arg1 ocicni.PodNetwork) ([]ocicni.NetResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodNetworkStatusWithContext", arg0, arg1)
	ret0, _ := ret[0].([]ocicni.NetResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodNetworkStatusWithContext indicates an expected call of GetPodNetworkStatusWithContext.
func (mr *MockCNIPluginMockRecorder) GetPodNetworkStatusWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodNetworkStatusWithContext", reflect.TypeOf((*MockCNIPlugin)(nil).GetPodNetworkStatusWithContext), arg0, arg1)
}

// Name mocks base method.
func (m *MockCNIPlugin) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockCNIPluginMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCNIPlugin)(nil).Name))
}

// SetUpPod mocks base method.
func (m *MockCNIPlugin) SetUpPod(arg0 ocicni.PodNetwork) ([]ocicni.NetResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUpPod", arg0)
	ret0, _ := ret[0].([]ocicni.NetResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUpPod indicates an expected call of SetUpPod.
func (mr *MockCNIPluginMockRecorder) SetUpPod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpPod", reflect.TypeOf((*MockCNIPlugin)(nil).SetUpPod), arg0)
}

// SetUpPodWithContext mocks base method.
func (m *MockCNIPlugin) SetUpPodWithContext(arg0 context.Context, arg1 ocicni.PodNetwork) ([]ocicni.NetResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUpPodWithContext", arg0, arg1)
	ret0, _ := ret[0].([]ocicni.NetResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUpPodWithContext indicates an expected call of SetUpPodWithContext.
func (mr *MockCNIPluginMockRecorder) SetUpPodWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpPodWithContext", reflect.TypeOf((*MockCNIPlugin)(nil).SetUpPodWithContext), arg0, arg1)
}

// Shutdown mocks base method.
func (m *MockCNIPlugin) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockCNIPluginMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockCNIPlugin)(nil).Shutdown))
}

// Status mocks base method.
func (m *MockCNIPlugin) Status() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(error)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockCNIPluginMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockCNIPlugin)(nil).Status))
}

// TearDownPod mocks base method.
func (m *MockCNIPlugin) TearDownPod(arg0 ocicni.PodNetwork) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TearDownPod", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TearDownPod indicates an expected call of TearDownPod.
func (mr *MockCNIPluginMockRecorder) TearDownPod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TearDownPod", reflect.TypeOf((*MockCNIPlugin)(nil).TearDownPod), arg0)
}

// TearDownPodWithContext mocks base method.
func (m *MockCNIPlugin) TearDownPodWithContext(arg0 context.Context, arg1 ocicni.PodNetwork) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TearDownPodWithContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TearDownPodWithContext indicates an expected call of TearDownPodWithContext.
func (mr *MockCNIPluginMockRecorder) TearDownPodWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TearDownPodWithContext", reflect.TypeOf((*MockCNIPlugin)(nil).TearDownPodWithContext), arg0, arg1)
}
