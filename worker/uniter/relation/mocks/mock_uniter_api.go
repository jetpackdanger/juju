// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/uniter/relation (interfaces: Unit,Relation,RelationUnit)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	uniter "github.com/juju/juju/api/uniter"
	params "github.com/juju/juju/apiserver/params"
	life "github.com/juju/juju/core/life"
	relation "github.com/juju/juju/core/relation"
	watcher "github.com/juju/juju/core/watcher"
	relation0 "github.com/juju/juju/worker/uniter/relation"
	names "github.com/juju/names/v4"
	reflect "reflect"
)

// MockUnit is a mock of Unit interface
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// Application mocks base method
func (m *MockUnit) Application() (relation0.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application")
	ret0, _ := ret[0].(relation0.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application
func (mr *MockUnitMockRecorder) Application() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockUnit)(nil).Application))
}

// ApplicationTag mocks base method
func (m *MockUnit) ApplicationTag() names.ApplicationTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationTag")
	ret0, _ := ret[0].(names.ApplicationTag)
	return ret0
}

// ApplicationTag indicates an expected call of ApplicationTag
func (mr *MockUnitMockRecorder) ApplicationTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationTag", reflect.TypeOf((*MockUnit)(nil).ApplicationTag))
}

// Destroy mocks base method
func (m *MockUnit) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockUnitMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockUnit)(nil).Destroy))
}

// Life mocks base method
func (m *MockUnit) Life() life.Value {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(life.Value)
	return ret0
}

// Life indicates an expected call of Life
func (mr *MockUnitMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockUnit)(nil).Life))
}

// Name mocks base method
func (m *MockUnit) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockUnitMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockUnit)(nil).Name))
}

// Refresh mocks base method
func (m *MockUnit) Refresh() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh")
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh
func (mr *MockUnitMockRecorder) Refresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockUnit)(nil).Refresh))
}

// RelationsStatus mocks base method
func (m *MockUnit) RelationsStatus() ([]uniter.RelationStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationsStatus")
	ret0, _ := ret[0].([]uniter.RelationStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelationsStatus indicates an expected call of RelationsStatus
func (mr *MockUnitMockRecorder) RelationsStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationsStatus", reflect.TypeOf((*MockUnit)(nil).RelationsStatus))
}

// SetState mocks base method
func (m *MockUnit) SetState(arg0 params.SetUnitStateArg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetState indicates an expected call of SetState
func (mr *MockUnitMockRecorder) SetState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockUnit)(nil).SetState), arg0)
}

// State mocks base method
func (m *MockUnit) State() (params.UnitStateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(params.UnitStateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State
func (mr *MockUnitMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockUnit)(nil).State))
}

// Tag mocks base method
func (m *MockUnit) Tag() names.UnitTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.UnitTag)
	return ret0
}

// Tag indicates an expected call of Tag
func (mr *MockUnitMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockUnit)(nil).Tag))
}

// Watch mocks base method
func (m *MockUnit) Watch() (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockUnitMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockUnit)(nil).Watch))
}

// MockRelation is a mock of Relation interface
type MockRelation struct {
	ctrl     *gomock.Controller
	recorder *MockRelationMockRecorder
}

// MockRelationMockRecorder is the mock recorder for MockRelation
type MockRelationMockRecorder struct {
	mock *MockRelation
}

// NewMockRelation creates a new mock instance
func NewMockRelation(ctrl *gomock.Controller) *MockRelation {
	mock := &MockRelation{ctrl: ctrl}
	mock.recorder = &MockRelationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelation) EXPECT() *MockRelationMockRecorder {
	return m.recorder
}

// Endpoint mocks base method
func (m *MockRelation) Endpoint() (*uniter.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoint")
	ret0, _ := ret[0].(*uniter.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Endpoint indicates an expected call of Endpoint
func (mr *MockRelationMockRecorder) Endpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*MockRelation)(nil).Endpoint))
}

// Id mocks base method
func (m *MockRelation) Id() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(int)
	return ret0
}

// Id indicates an expected call of Id
func (mr *MockRelationMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockRelation)(nil).Id))
}

// Life mocks base method
func (m *MockRelation) Life() life.Value {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(life.Value)
	return ret0
}

// Life indicates an expected call of Life
func (mr *MockRelationMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockRelation)(nil).Life))
}

// OtherApplication mocks base method
func (m *MockRelation) OtherApplication() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OtherApplication")
	ret0, _ := ret[0].(string)
	return ret0
}

// OtherApplication indicates an expected call of OtherApplication
func (mr *MockRelationMockRecorder) OtherApplication() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OtherApplication", reflect.TypeOf((*MockRelation)(nil).OtherApplication))
}

// Refresh mocks base method
func (m *MockRelation) Refresh() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh")
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh
func (mr *MockRelationMockRecorder) Refresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockRelation)(nil).Refresh))
}

// SetStatus mocks base method
func (m *MockRelation) SetStatus(arg0 relation.Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockRelationMockRecorder) SetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockRelation)(nil).SetStatus), arg0)
}

// String mocks base method
func (m *MockRelation) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockRelationMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockRelation)(nil).String))
}

// Suspended mocks base method
func (m *MockRelation) Suspended() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suspended")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Suspended indicates an expected call of Suspended
func (mr *MockRelationMockRecorder) Suspended() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suspended", reflect.TypeOf((*MockRelation)(nil).Suspended))
}

// Tag mocks base method
func (m *MockRelation) Tag() names.RelationTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.RelationTag)
	return ret0
}

// Tag indicates an expected call of Tag
func (mr *MockRelationMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockRelation)(nil).Tag))
}

// Unit mocks base method
func (m *MockRelation) Unit(arg0 names.UnitTag) (relation0.RelationUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(relation0.RelationUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit
func (mr *MockRelationMockRecorder) Unit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockRelation)(nil).Unit), arg0)
}

// UpdateSuspended mocks base method
func (m *MockRelation) UpdateSuspended(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateSuspended", arg0)
}

// UpdateSuspended indicates an expected call of UpdateSuspended
func (mr *MockRelationMockRecorder) UpdateSuspended(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSuspended", reflect.TypeOf((*MockRelation)(nil).UpdateSuspended), arg0)
}

// MockRelationUnit is a mock of RelationUnit interface
type MockRelationUnit struct {
	ctrl     *gomock.Controller
	recorder *MockRelationUnitMockRecorder
}

// MockRelationUnitMockRecorder is the mock recorder for MockRelationUnit
type MockRelationUnitMockRecorder struct {
	mock *MockRelationUnit
}

// NewMockRelationUnit creates a new mock instance
func NewMockRelationUnit(ctrl *gomock.Controller) *MockRelationUnit {
	mock := &MockRelationUnit{ctrl: ctrl}
	mock.recorder = &MockRelationUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelationUnit) EXPECT() *MockRelationUnitMockRecorder {
	return m.recorder
}

// ApplicationSettings mocks base method
func (m *MockRelationUnit) ApplicationSettings() (*uniter.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationSettings")
	ret0, _ := ret[0].(*uniter.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationSettings indicates an expected call of ApplicationSettings
func (mr *MockRelationUnitMockRecorder) ApplicationSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationSettings", reflect.TypeOf((*MockRelationUnit)(nil).ApplicationSettings))
}

// Endpoint mocks base method
func (m *MockRelationUnit) Endpoint() uniter.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoint")
	ret0, _ := ret[0].(uniter.Endpoint)
	return ret0
}

// Endpoint indicates an expected call of Endpoint
func (mr *MockRelationUnitMockRecorder) Endpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*MockRelationUnit)(nil).Endpoint))
}

// EnterScope mocks base method
func (m *MockRelationUnit) EnterScope() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnterScope")
	ret0, _ := ret[0].(error)
	return ret0
}

// EnterScope indicates an expected call of EnterScope
func (mr *MockRelationUnitMockRecorder) EnterScope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnterScope", reflect.TypeOf((*MockRelationUnit)(nil).EnterScope))
}

// LeaveScope mocks base method
func (m *MockRelationUnit) LeaveScope() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveScope")
	ret0, _ := ret[0].(error)
	return ret0
}

// LeaveScope indicates an expected call of LeaveScope
func (mr *MockRelationUnitMockRecorder) LeaveScope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveScope", reflect.TypeOf((*MockRelationUnit)(nil).LeaveScope))
}

// ReadSettings mocks base method
func (m *MockRelationUnit) ReadSettings(arg0 string) (params.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSettings", arg0)
	ret0, _ := ret[0].(params.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSettings indicates an expected call of ReadSettings
func (mr *MockRelationUnitMockRecorder) ReadSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSettings", reflect.TypeOf((*MockRelationUnit)(nil).ReadSettings), arg0)
}

// Relation mocks base method
func (m *MockRelationUnit) Relation() relation0.Relation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relation")
	ret0, _ := ret[0].(relation0.Relation)
	return ret0
}

// Relation indicates an expected call of Relation
func (mr *MockRelationUnitMockRecorder) Relation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relation", reflect.TypeOf((*MockRelationUnit)(nil).Relation))
}

// Settings mocks base method
func (m *MockRelationUnit) Settings() (*uniter.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Settings")
	ret0, _ := ret[0].(*uniter.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Settings indicates an expected call of Settings
func (mr *MockRelationUnitMockRecorder) Settings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settings", reflect.TypeOf((*MockRelationUnit)(nil).Settings))
}

// UpdateRelationSettings mocks base method
func (m *MockRelationUnit) UpdateRelationSettings(arg0, arg1 params.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRelationSettings", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRelationSettings indicates an expected call of UpdateRelationSettings
func (mr *MockRelationUnitMockRecorder) UpdateRelationSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRelationSettings", reflect.TypeOf((*MockRelationUnit)(nil).UpdateRelationSettings), arg0, arg1)
}
