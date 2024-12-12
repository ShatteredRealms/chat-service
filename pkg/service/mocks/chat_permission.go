// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/chat-service/pkg/service/chat_permission.go
//
// Generated by this command:
//
//	mockgen -source=/home/wil/dev/sro/chat-service/pkg/service/chat_permission.go -destination=/home/wil/dev/sro/chat-service/pkg/service/mocks/chat_permission.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	chat "github.com/ShatteredRealms/chat-service/pkg/model/chat"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockChatChannelPermissionService is a mock of ChatChannelPermissionService interface.
type MockChatChannelPermissionService struct {
	ctrl     *gomock.Controller
	recorder *MockChatChannelPermissionServiceMockRecorder
	isgomock struct{}
}

// MockChatChannelPermissionServiceMockRecorder is the mock recorder for MockChatChannelPermissionService.
type MockChatChannelPermissionServiceMockRecorder struct {
	mock *MockChatChannelPermissionService
}

// NewMockChatChannelPermissionService creates a new mock instance.
func NewMockChatChannelPermissionService(ctrl *gomock.Controller) *MockChatChannelPermissionService {
	mock := &MockChatChannelPermissionService{ctrl: ctrl}
	mock.recorder = &MockChatChannelPermissionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatChannelPermissionService) EXPECT() *MockChatChannelPermissionServiceMockRecorder {
	return m.recorder
}

// AddForCharacter mocks base method.
func (m *MockChatChannelPermissionService) AddForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddForCharacter", ctx, characterId, channelIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddForCharacter indicates an expected call of AddForCharacter.
func (mr *MockChatChannelPermissionServiceMockRecorder) AddForCharacter(ctx, characterId, channelIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddForCharacter", reflect.TypeOf((*MockChatChannelPermissionService)(nil).AddForCharacter), ctx, characterId, channelIds)
}

// GetForCharacter mocks base method.
func (m *MockChatChannelPermissionService) GetForCharacter(ctx context.Context, characterId string) (*chat.Channels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForCharacter", ctx, characterId)
	ret0, _ := ret[0].(*chat.Channels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForCharacter indicates an expected call of GetForCharacter.
func (mr *MockChatChannelPermissionServiceMockRecorder) GetForCharacter(ctx, characterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForCharacter", reflect.TypeOf((*MockChatChannelPermissionService)(nil).GetForCharacter), ctx, characterId)
}

// HasAccess mocks base method.
func (m *MockChatChannelPermissionService) HasAccess(ctx context.Context, channelId *uuid.UUID, characterId string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAccess", ctx, channelId, characterId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAccess indicates an expected call of HasAccess.
func (mr *MockChatChannelPermissionServiceMockRecorder) HasAccess(ctx, channelId, characterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAccess", reflect.TypeOf((*MockChatChannelPermissionService)(nil).HasAccess), ctx, channelId, characterId)
}

// RemForCharacter mocks base method.
func (m *MockChatChannelPermissionService) RemForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemForCharacter", ctx, characterId, channelIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemForCharacter indicates an expected call of RemForCharacter.
func (mr *MockChatChannelPermissionServiceMockRecorder) RemForCharacter(ctx, characterId, channelIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemForCharacter", reflect.TypeOf((*MockChatChannelPermissionService)(nil).RemForCharacter), ctx, characterId, channelIds)
}

// SaveForCharacter mocks base method.
func (m *MockChatChannelPermissionService) SaveForCharacter(ctx context.Context, characterId string, channelIds []*uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveForCharacter", ctx, characterId, channelIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveForCharacter indicates an expected call of SaveForCharacter.
func (mr *MockChatChannelPermissionServiceMockRecorder) SaveForCharacter(ctx, characterId, channelIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveForCharacter", reflect.TypeOf((*MockChatChannelPermissionService)(nil).SaveForCharacter), ctx, characterId, channelIds)
}