// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/chat-service/pkg/service/chatchannel.go
//
// Generated by this command:
//
//	mockgen -source=/home/wil/dev/sro/chat-service/pkg/service/chatchannel.go -destination=/home/wil/dev/sro/chat-service/pkg/service/mocks/chatchannel.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	chat "github.com/ShatteredRealms/chat-service/pkg/model/chat"
	pb "github.com/ShatteredRealms/chat-service/pkg/pb"
	gomock "go.uber.org/mock/gomock"
)

// MockChatChannelService is a mock of ChatChannelService interface.
type MockChatChannelService struct {
	ctrl     *gomock.Controller
	recorder *MockChatChannelServiceMockRecorder
	isgomock struct{}
}

// MockChatChannelServiceMockRecorder is the mock recorder for MockChatChannelService.
type MockChatChannelServiceMockRecorder struct {
	mock *MockChatChannelService
}

// NewMockChatChannelService creates a new mock instance.
func NewMockChatChannelService(ctrl *gomock.Controller) *MockChatChannelService {
	mock := &MockChatChannelService{ctrl: ctrl}
	mock.recorder = &MockChatChannelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatChannelService) EXPECT() *MockChatChannelServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChatChannelService) Create(ctx context.Context, name, dimensionId string) (*chat.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, name, dimensionId)
	ret0, _ := ret[0].(*chat.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockChatChannelServiceMockRecorder) Create(ctx, name, dimensionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChatChannelService)(nil).Create), ctx, name, dimensionId)
}

// Delete mocks base method.
func (m *MockChatChannelService) Delete(ctx context.Context, channelId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, channelId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockChatChannelServiceMockRecorder) Delete(ctx, channelId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockChatChannelService)(nil).Delete), ctx, channelId)
}

// GetAll mocks base method.
func (m *MockChatChannelService) GetAll(ctx context.Context) (*chat.Channels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(*chat.Channels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockChatChannelServiceMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockChatChannelService)(nil).GetAll), ctx)
}

// GetById mocks base method.
func (m *MockChatChannelService) GetById(ctx context.Context, id string) (*chat.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*chat.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockChatChannelServiceMockRecorder) GetById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockChatChannelService)(nil).GetById), ctx, id)
}

// Save mocks base method.
func (m *MockChatChannelService) Save(ctx context.Context, channel *chat.Channel) (*chat.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, channel)
	ret0, _ := ret[0].(*chat.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockChatChannelServiceMockRecorder) Save(ctx, channel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockChatChannelService)(nil).Save), ctx, channel)
}

// Update mocks base method.
func (m *MockChatChannelService) Update(ctx context.Context, pbRequest *pb.UpdateChatChannelRequest) (*chat.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, pbRequest)
	ret0, _ := ret[0].(*chat.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockChatChannelServiceMockRecorder) Update(ctx, pbRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockChatChannelService)(nil).Update), ctx, pbRequest)
}
