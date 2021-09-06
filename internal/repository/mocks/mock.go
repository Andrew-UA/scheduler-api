// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	model "scheduler/internal/model"

	gomock "github.com/golang/mock/gomock"
)

// MockISchedule is a mock of ISchedule interface.
type MockISchedule struct {
	ctrl     *gomock.Controller
	recorder *MockIScheduleMockRecorder
}

// MockIScheduleMockRecorder is the mock recorder for MockISchedule.
type MockIScheduleMockRecorder struct {
	mock *MockISchedule
}

// NewMockISchedule creates a new mock instance.
func NewMockISchedule(ctrl *gomock.Controller) *MockISchedule {
	mock := &MockISchedule{ctrl: ctrl}
	mock.recorder = &MockIScheduleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISchedule) EXPECT() *MockIScheduleMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockISchedule) Create(m model.ScheduleEvent) (model.ScheduleEvent, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", m)
	ret0, _ := ret[0].(model.ScheduleEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIScheduleMockRecorder) Create(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockISchedule)(nil).Create), m)
}

// Delete mocks base method.
func (m *MockISchedule) Delete(ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIScheduleMockRecorder) Delete(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockISchedule)(nil).Delete), ID)
}

// List mocks base method.
func (m *MockISchedule) List(params map[string]string) ([]model.ScheduleEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", params)
	ret0, _ := ret[0].([]model.ScheduleEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIScheduleMockRecorder) List(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockISchedule)(nil).List), params)
}

// Show mocks base method.
func (m *MockISchedule) Show(D int) (model.ScheduleEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", D)
	ret0, _ := ret[0].(model.ScheduleEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockIScheduleMockRecorder) Show(D interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockISchedule)(nil).Show), D)
}

// Update mocks base method.
func (m_2 *MockISchedule) Update(ID int, m model.ScheduleEvent) (model.ScheduleEvent, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ID, m)
	ret0, _ := ret[0].(model.ScheduleEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIScheduleMockRecorder) Update(ID, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockISchedule)(nil).Update), ID, m)
}

// MockIUser is a mock of IUser interface.
type MockIUser struct {
	ctrl     *gomock.Controller
	recorder *MockIUserMockRecorder
}

// MockIUserMockRecorder is the mock recorder for MockIUser.
type MockIUserMockRecorder struct {
	mock *MockIUser
}

// NewMockIUser creates a new mock instance.
func NewMockIUser(ctrl *gomock.Controller) *MockIUser {
	mock := &MockIUser{ctrl: ctrl}
	mock.recorder = &MockIUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUser) EXPECT() *MockIUserMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockIUser) FindByID(ID int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ID)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockIUserMockRecorder) FindByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockIUser)(nil).FindByID), ID)
}

// FindByLogin mocks base method.
func (m *MockIUser) FindByLogin(login string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLogin", login)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLogin indicates an expected call of FindByLogin.
func (mr *MockIUserMockRecorder) FindByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*MockIUser)(nil).FindByLogin), login)
}

// Update mocks base method.
func (m *MockIUser) Update(userID int, user model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userID, user)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIUserMockRecorder) Update(userID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUser)(nil).Update), userID, user)
}
