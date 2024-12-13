// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/chat-service/pkg/repository/chat_history.go
//
// Generated by this command:
//
//	mockgen -source=/home/wil/dev/sro/chat-service/pkg/repository/chat_history.go -destination=/home/wil/dev/sro/chat-service/pkg/repository/mocks/chat_history.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"
	time "time"

	chat "github.com/ShatteredRealms/chat-service/pkg/model/chat"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockChatLogRepository is a mock of ChatLogRepository interface.
type MockChatLogRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatLogRepositoryMockRecorder
	isgomock struct{}
}

// MockChatLogRepositoryMockRecorder is the mock recorder for MockChatLogRepository.
type MockChatLogRepositoryMockRecorder struct {
	mock *MockChatLogRepository
}

// NewMockChatLogRepository creates a new mock instance.
func NewMockChatLogRepository(ctrl *gomock.Controller) *MockChatLogRepository {
	mock := &MockChatLogRepository{ctrl: ctrl}
	mock.recorder = &MockChatLogRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatLogRepository) EXPECT() *MockChatLogRepositoryMockRecorder {
	return m.recorder
}

// AddMessage mocks base method.
func (m *MockChatLogRepository) AddMessage(ctx context.Context, channelId, dimensionId *uuid.UUID, message *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMessage", ctx, channelId, dimensionId, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMessage indicates an expected call of AddMessage.
func (mr *MockChatLogRepositoryMockRecorder) AddMessage(ctx, channelId, dimensionId, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessage", reflect.TypeOf((*MockChatLogRepository)(nil).AddMessage), ctx, channelId, dimensionId, message)
}

// GetMessages mocks base method.
func (m *MockChatLogRepository) GetMessages(ctx context.Context, channelId *uuid.UUID, limit, offset *uint, before, after *time.Time, sender *uuid.UUID) (*chat.MessageLogs, uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", ctx, channelId, limit, offset, before, after, sender)
	ret0, _ := ret[0].(*chat.MessageLogs)
	ret1, _ := ret[1].(uint)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMessages indicates an expected call of GetMessages.
func (mr *MockChatLogRepositoryMockRecorder) GetMessages(ctx, channelId, limit, offset, before, after, sender any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockChatLogRepository)(nil).GetMessages), ctx, channelId, limit, offset, before, after, sender)
}
