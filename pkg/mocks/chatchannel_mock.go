// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/chat-service/pkg/repository/chatchannel.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/dev/sro/chat-service/pkg/repository/chatchannel.go -destination=/home/wil/dev/sro/chat-service/pkg/mocks/chatchannel_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	chat "github.com/ShatteredRealms/chat-service/pkg/model/chat"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockChatChannelRepository is a mock of ChatChannelRepository interface.
type MockChatChannelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatChannelRepositoryMockRecorder
	isgomock struct{}
}

// MockChatChannelRepositoryMockRecorder is the mock recorder for MockChatChannelRepository.
type MockChatChannelRepositoryMockRecorder struct {
	mock *MockChatChannelRepository
}

// NewMockChatChannelRepository creates a new mock instance.
func NewMockChatChannelRepository(ctrl *gomock.Controller) *MockChatChannelRepository {
	mock := &MockChatChannelRepository{ctrl: ctrl}
	mock.recorder = &MockChatChannelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatChannelRepository) EXPECT() *MockChatChannelRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChatChannelRepository) Create(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, channel)
	ret0, _ := ret[0].(*chat.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockChatChannelRepositoryMockRecorder) Create(ctx, channel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChatChannelRepository)(nil).Create), ctx, channel)
}

// Delete mocks base method.
func (m *MockChatChannelRepository) Delete(ctx context.Context, channelId *uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, channelId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockChatChannelRepositoryMockRecorder) Delete(ctx, channelId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockChatChannelRepository)(nil).Delete), ctx, channelId)
}

// GetAll mocks base method.
func (m *MockChatChannelRepository) GetAll(ctx context.Context) (*chat.Channels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(*chat.Channels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockChatChannelRepositoryMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockChatChannelRepository)(nil).GetAll), ctx)
}

// GetById mocks base method.
func (m *MockChatChannelRepository) GetById(ctx context.Context, id *uuid.UUID) (*chat.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*chat.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockChatChannelRepositoryMockRecorder) GetById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockChatChannelRepository)(nil).GetById), ctx, id)
}

// Save mocks base method.
func (m *MockChatChannelRepository) Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, channel)
	ret0, _ := ret[0].(*chat.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockChatChannelRepositoryMockRecorder) Save(ctx, channel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockChatChannelRepository)(nil).Save), ctx, channel)
}