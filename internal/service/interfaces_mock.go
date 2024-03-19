// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service is a generated GoMock package.
package service

import (
	dto "Hezzl_testing/internal/service/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(projectId int, name string) (*dto.CreateGoodResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", projectId, name)
	ret0, _ := ret[0].(*dto.CreateGoodResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(projectId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), projectId, name)
}

// Delete mocks base method.
func (m *MockRepository) Delete(id, projectId int) (*dto.DeleteGoodResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, projectId)
	ret0, _ := ret[0].(*dto.DeleteGoodResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(id, projectId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id, projectId)
}

// Get mocks base method.
func (m *MockRepository) Get(limit, page int) (*dto.GetGoodsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", limit, page)
	ret0, _ := ret[0].(*dto.GetGoodsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(limit, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), limit, page)
}

// PatchPriority mocks base method.
func (m *MockRepository) PatchPriority(id, projectId, newPriority int) (*dto.ReprioritizeResponse, *dto.UpdateGoodsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchPriority", id, projectId, newPriority)
	ret0, _ := ret[0].(*dto.ReprioritizeResponse)
	ret1, _ := ret[1].(*dto.UpdateGoodsResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PatchPriority indicates an expected call of PatchPriority.
func (mr *MockRepositoryMockRecorder) PatchPriority(id, projectId, newPriority interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchPriority", reflect.TypeOf((*MockRepository)(nil).PatchPriority), id, projectId, newPriority)
}

// Update mocks base method.
func (m *MockRepository) Update(id, projectId int, name, description string) (*dto.UpdateGoodResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, projectId, name, description)
	ret0, _ := ret[0].(*dto.UpdateGoodResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(id, projectId, name, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), id, projectId, name, description)
}
