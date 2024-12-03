// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/chat-service/pkg/repository/chat_permission.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/dev/sro/chat-service/pkg/repository/chat_permission.go -destination=/home/wil/dev/sro/chat-service/pkg/mocks/chat_permission_mock.go
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

// MockChatChannelPermissionRepository is a mock of ChatChannelPermissionRepository interface.
type MockChatChannelPermissionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatChannelPermissionRepositoryMockRecorder
	isgomock struct{}
}

// MockChatChannelPermissionRepositoryMockRecorder is the mock recorder for MockChatChannelPermissionRepository.
type MockChatChannelPermissionRepositoryMockRecorder struct {
	mock *MockChatChannelPermissionRepository
}

// NewMockChatChannelPermissionRepository creates a new mock instance.
func NewMockChatChannelPermissionRepository(ctrl *gomock.Controller) *MockChatChannelPermissionRepository {
	mock := &MockChatChannelPermissionRepository{ctrl: ctrl}
	mock.recorder = &MockChatChannelPermissionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatChannelPermissionRepository) EXPECT() *MockChatChannelPermissionRepositoryMockRecorder {
	return m.recorder
}

// GetForCharacter mocks base method.
func (m *MockChatChannelPermissionRepository) GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForCharacter", ctx, characterId)
	ret0, _ := ret[0].(*chat.Channels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForCharacter indicates an expected call of GetForCharacter.
func (mr *MockChatChannelPermissionRepositoryMockRecorder) GetForCharacter(ctx, characterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForCharacter", reflect.TypeOf((*MockChatChannelPermissionRepository)(nil).GetForCharacter), ctx, characterId)
}

// HasAccess mocks base method.
func (m *MockChatChannelPermissionRepository) HasAccess(ctx context.Context, channelId *uuid.UUID, characterId string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAccess", ctx, channelId, characterId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAccess indicates an expected call of HasAccess.
func (mr *MockChatChannelPermissionRepositoryMockRecorder) HasAccess(ctx, channelId, characterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAccess", reflect.TypeOf((*MockChatChannelPermissionRepository)(nil).HasAccess), ctx, channelId, characterId)
}

// SaveForCharacter mocks base method.
func (m *MockChatChannelPermissionRepository) SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveForCharacter", ctx, characterId, channelIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveForCharacter indicates an expected call of SaveForCharacter.
func (mr *MockChatChannelPermissionRepositoryMockRecorder) SaveForCharacter(ctx, characterId, channelIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveForCharacter", reflect.TypeOf((*MockChatChannelPermissionRepository)(nil).SaveForCharacter), ctx, characterId, channelIds)
}