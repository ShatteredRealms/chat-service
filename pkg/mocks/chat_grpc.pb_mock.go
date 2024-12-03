// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/chat-service/pkg/pb/chat_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/dev/sro/chat-service/pkg/pb/chat_grpc.pb.go -destination=/home/wil/dev/sro/chat-service/pkg/mocks/chat_grpc.pb_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	pb "github.com/ShatteredRealms/chat-service/pkg/pb"
	pb0 "github.com/ShatteredRealms/go-common-service/pkg/pb"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockChatServiceClient is a mock of ChatServiceClient interface.
type MockChatServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceClientMockRecorder
	isgomock struct{}
}

// MockChatServiceClientMockRecorder is the mock recorder for MockChatServiceClient.
type MockChatServiceClientMockRecorder struct {
	mock *MockChatServiceClient
}

// NewMockChatServiceClient creates a new mock instance.
func NewMockChatServiceClient(ctrl *gomock.Controller) *MockChatServiceClient {
	mock := &MockChatServiceClient{ctrl: ctrl}
	mock.recorder = &MockChatServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServiceClient) EXPECT() *MockChatServiceClientMockRecorder {
	return m.recorder
}

// ConnectChatChannel mocks base method.
func (m *MockChatServiceClient) ConnectChatChannel(ctx context.Context, in *pb.ConnectChatChannelRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.ChatMessage], error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConnectChatChannel", varargs...)
	ret0, _ := ret[0].(grpc.ServerStreamingClient[pb.ChatMessage])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectChatChannel indicates an expected call of ConnectChatChannel.
func (mr *MockChatServiceClientMockRecorder) ConnectChatChannel(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectChatChannel", reflect.TypeOf((*MockChatServiceClient)(nil).ConnectChatChannel), varargs...)
}

// ConnectDirectMessages mocks base method.
func (m *MockChatServiceClient) ConnectDirectMessages(ctx context.Context, in *pb0.TargetId, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.ChatMessage], error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConnectDirectMessages", varargs...)
	ret0, _ := ret[0].(grpc.ServerStreamingClient[pb.ChatMessage])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectDirectMessages indicates an expected call of ConnectDirectMessages.
func (mr *MockChatServiceClientMockRecorder) ConnectDirectMessages(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectDirectMessages", reflect.TypeOf((*MockChatServiceClient)(nil).ConnectDirectMessages), varargs...)
}

// CreateChatChannel mocks base method.
func (m *MockChatServiceClient) CreateChatChannel(ctx context.Context, in *pb.CreateChatChannelMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateChatChannel", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChatChannel indicates an expected call of CreateChatChannel.
func (mr *MockChatServiceClientMockRecorder) CreateChatChannel(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChatChannel", reflect.TypeOf((*MockChatServiceClient)(nil).CreateChatChannel), varargs...)
}

// DeleteChatChannel mocks base method.
func (m *MockChatServiceClient) DeleteChatChannel(ctx context.Context, in *pb0.TargetId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteChatChannel", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteChatChannel indicates an expected call of DeleteChatChannel.
func (mr *MockChatServiceClientMockRecorder) DeleteChatChannel(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChatChannel", reflect.TypeOf((*MockChatServiceClient)(nil).DeleteChatChannel), varargs...)
}

// EditChatChannel mocks base method.
func (m *MockChatServiceClient) EditChatChannel(ctx context.Context, in *pb.UpdateChatChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditChatChannel", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditChatChannel indicates an expected call of EditChatChannel.
func (mr *MockChatServiceClientMockRecorder) EditChatChannel(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditChatChannel", reflect.TypeOf((*MockChatServiceClient)(nil).EditChatChannel), varargs...)
}

// GetAuthorizedChatChannels mocks base method.
func (m *MockChatServiceClient) GetAuthorizedChatChannels(ctx context.Context, in *pb0.TargetId, opts ...grpc.CallOption) (*pb.ChatChannels, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthorizedChatChannels", varargs...)
	ret0, _ := ret[0].(*pb.ChatChannels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizedChatChannels indicates an expected call of GetAuthorizedChatChannels.
func (mr *MockChatServiceClientMockRecorder) GetAuthorizedChatChannels(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizedChatChannels", reflect.TypeOf((*MockChatServiceClient)(nil).GetAuthorizedChatChannels), varargs...)
}

// GetChatChannel mocks base method.
func (m *MockChatServiceClient) GetChatChannel(ctx context.Context, in *pb0.TargetId, opts ...grpc.CallOption) (*pb.ChatChannel, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChatChannel", varargs...)
	ret0, _ := ret[0].(*pb.ChatChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatChannel indicates an expected call of GetChatChannel.
func (mr *MockChatServiceClientMockRecorder) GetChatChannel(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatChannel", reflect.TypeOf((*MockChatServiceClient)(nil).GetChatChannel), varargs...)
}

// GetChatChannels mocks base method.
func (m *MockChatServiceClient) GetChatChannels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.ChatChannels, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChatChannels", varargs...)
	ret0, _ := ret[0].(*pb.ChatChannels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatChannels indicates an expected call of GetChatChannels.
func (mr *MockChatServiceClientMockRecorder) GetChatChannels(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatChannels", reflect.TypeOf((*MockChatServiceClient)(nil).GetChatChannels), varargs...)
}

// SendChatChannelMessage mocks base method.
func (m *MockChatServiceClient) SendChatChannelMessage(ctx context.Context, in *pb.SendChatChannelMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendChatChannelMessage", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendChatChannelMessage indicates an expected call of SendChatChannelMessage.
func (mr *MockChatServiceClientMockRecorder) SendChatChannelMessage(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChatChannelMessage", reflect.TypeOf((*MockChatServiceClient)(nil).SendChatChannelMessage), varargs...)
}

// SendDirectMessage mocks base method.
func (m *MockChatServiceClient) SendDirectMessage(ctx context.Context, in *pb.SendDirectMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendDirectMessage", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDirectMessage indicates an expected call of SendDirectMessage.
func (mr *MockChatServiceClientMockRecorder) SendDirectMessage(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDirectMessage", reflect.TypeOf((*MockChatServiceClient)(nil).SendDirectMessage), varargs...)
}

// SetCharacterChatChannelAuth mocks base method.
func (m *MockChatServiceClient) SetCharacterChatChannelAuth(ctx context.Context, in *pb.RequestSetCharacterSetChatChannelAuth, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetCharacterChatChannelAuth", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCharacterChatChannelAuth indicates an expected call of SetCharacterChatChannelAuth.
func (mr *MockChatServiceClientMockRecorder) SetCharacterChatChannelAuth(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharacterChatChannelAuth", reflect.TypeOf((*MockChatServiceClient)(nil).SetCharacterChatChannelAuth), varargs...)
}

// UpdateCharacterChatChannelAuth mocks base method.
func (m *MockChatServiceClient) UpdateCharacterChatChannelAuth(ctx context.Context, in *pb.RequestUpdateCharacterSetChatChannelAuth, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCharacterChatChannelAuth", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCharacterChatChannelAuth indicates an expected call of UpdateCharacterChatChannelAuth.
func (mr *MockChatServiceClientMockRecorder) UpdateCharacterChatChannelAuth(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCharacterChatChannelAuth", reflect.TypeOf((*MockChatServiceClient)(nil).UpdateCharacterChatChannelAuth), varargs...)
}

// MockChatServiceServer is a mock of ChatServiceServer interface.
type MockChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceServerMockRecorder
	isgomock struct{}
}

// MockChatServiceServerMockRecorder is the mock recorder for MockChatServiceServer.
type MockChatServiceServerMockRecorder struct {
	mock *MockChatServiceServer
}

// NewMockChatServiceServer creates a new mock instance.
func NewMockChatServiceServer(ctrl *gomock.Controller) *MockChatServiceServer {
	mock := &MockChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServiceServer) EXPECT() *MockChatServiceServerMockRecorder {
	return m.recorder
}

// ConnectChatChannel mocks base method.
func (m *MockChatServiceServer) ConnectChatChannel(arg0 *pb.ConnectChatChannelRequest, arg1 grpc.ServerStreamingServer[pb.ChatMessage]) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectChatChannel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectChatChannel indicates an expected call of ConnectChatChannel.
func (mr *MockChatServiceServerMockRecorder) ConnectChatChannel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectChatChannel", reflect.TypeOf((*MockChatServiceServer)(nil).ConnectChatChannel), arg0, arg1)
}

// ConnectDirectMessages mocks base method.
func (m *MockChatServiceServer) ConnectDirectMessages(arg0 *pb0.TargetId, arg1 grpc.ServerStreamingServer[pb.ChatMessage]) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectDirectMessages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectDirectMessages indicates an expected call of ConnectDirectMessages.
func (mr *MockChatServiceServerMockRecorder) ConnectDirectMessages(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectDirectMessages", reflect.TypeOf((*MockChatServiceServer)(nil).ConnectDirectMessages), arg0, arg1)
}

// CreateChatChannel mocks base method.
func (m *MockChatServiceServer) CreateChatChannel(arg0 context.Context, arg1 *pb.CreateChatChannelMessage) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChatChannel", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChatChannel indicates an expected call of CreateChatChannel.
func (mr *MockChatServiceServerMockRecorder) CreateChatChannel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChatChannel", reflect.TypeOf((*MockChatServiceServer)(nil).CreateChatChannel), arg0, arg1)
}

// DeleteChatChannel mocks base method.
func (m *MockChatServiceServer) DeleteChatChannel(arg0 context.Context, arg1 *pb0.TargetId) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChatChannel", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteChatChannel indicates an expected call of DeleteChatChannel.
func (mr *MockChatServiceServerMockRecorder) DeleteChatChannel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChatChannel", reflect.TypeOf((*MockChatServiceServer)(nil).DeleteChatChannel), arg0, arg1)
}

// EditChatChannel mocks base method.
func (m *MockChatServiceServer) EditChatChannel(arg0 context.Context, arg1 *pb.UpdateChatChannelRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditChatChannel", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditChatChannel indicates an expected call of EditChatChannel.
func (mr *MockChatServiceServerMockRecorder) EditChatChannel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditChatChannel", reflect.TypeOf((*MockChatServiceServer)(nil).EditChatChannel), arg0, arg1)
}

// GetAuthorizedChatChannels mocks base method.
func (m *MockChatServiceServer) GetAuthorizedChatChannels(arg0 context.Context, arg1 *pb0.TargetId) (*pb.ChatChannels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorizedChatChannels", arg0, arg1)
	ret0, _ := ret[0].(*pb.ChatChannels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizedChatChannels indicates an expected call of GetAuthorizedChatChannels.
func (mr *MockChatServiceServerMockRecorder) GetAuthorizedChatChannels(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizedChatChannels", reflect.TypeOf((*MockChatServiceServer)(nil).GetAuthorizedChatChannels), arg0, arg1)
}

// GetChatChannel mocks base method.
func (m *MockChatServiceServer) GetChatChannel(arg0 context.Context, arg1 *pb0.TargetId) (*pb.ChatChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatChannel", arg0, arg1)
	ret0, _ := ret[0].(*pb.ChatChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatChannel indicates an expected call of GetChatChannel.
func (mr *MockChatServiceServerMockRecorder) GetChatChannel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatChannel", reflect.TypeOf((*MockChatServiceServer)(nil).GetChatChannel), arg0, arg1)
}

// GetChatChannels mocks base method.
func (m *MockChatServiceServer) GetChatChannels(arg0 context.Context, arg1 *emptypb.Empty) (*pb.ChatChannels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatChannels", arg0, arg1)
	ret0, _ := ret[0].(*pb.ChatChannels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatChannels indicates an expected call of GetChatChannels.
func (mr *MockChatServiceServerMockRecorder) GetChatChannels(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatChannels", reflect.TypeOf((*MockChatServiceServer)(nil).GetChatChannels), arg0, arg1)
}

// SendChatChannelMessage mocks base method.
func (m *MockChatServiceServer) SendChatChannelMessage(arg0 context.Context, arg1 *pb.SendChatChannelMessageRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendChatChannelMessage", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendChatChannelMessage indicates an expected call of SendChatChannelMessage.
func (mr *MockChatServiceServerMockRecorder) SendChatChannelMessage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChatChannelMessage", reflect.TypeOf((*MockChatServiceServer)(nil).SendChatChannelMessage), arg0, arg1)
}

// SendDirectMessage mocks base method.
func (m *MockChatServiceServer) SendDirectMessage(arg0 context.Context, arg1 *pb.SendDirectMessageRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDirectMessage", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDirectMessage indicates an expected call of SendDirectMessage.
func (mr *MockChatServiceServerMockRecorder) SendDirectMessage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDirectMessage", reflect.TypeOf((*MockChatServiceServer)(nil).SendDirectMessage), arg0, arg1)
}

// SetCharacterChatChannelAuth mocks base method.
func (m *MockChatServiceServer) SetCharacterChatChannelAuth(arg0 context.Context, arg1 *pb.RequestSetCharacterSetChatChannelAuth) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharacterChatChannelAuth", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCharacterChatChannelAuth indicates an expected call of SetCharacterChatChannelAuth.
func (mr *MockChatServiceServerMockRecorder) SetCharacterChatChannelAuth(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharacterChatChannelAuth", reflect.TypeOf((*MockChatServiceServer)(nil).SetCharacterChatChannelAuth), arg0, arg1)
}

// UpdateCharacterChatChannelAuth mocks base method.
func (m *MockChatServiceServer) UpdateCharacterChatChannelAuth(arg0 context.Context, arg1 *pb.RequestUpdateCharacterSetChatChannelAuth) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCharacterChatChannelAuth", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCharacterChatChannelAuth indicates an expected call of UpdateCharacterChatChannelAuth.
func (mr *MockChatServiceServerMockRecorder) UpdateCharacterChatChannelAuth(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCharacterChatChannelAuth", reflect.TypeOf((*MockChatServiceServer)(nil).UpdateCharacterChatChannelAuth), arg0, arg1)
}

// mustEmbedUnimplementedChatServiceServer mocks base method.
func (m *MockChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatServiceServer")
}

// mustEmbedUnimplementedChatServiceServer indicates an expected call of mustEmbedUnimplementedChatServiceServer.
func (mr *MockChatServiceServerMockRecorder) mustEmbedUnimplementedChatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatServiceServer", reflect.TypeOf((*MockChatServiceServer)(nil).mustEmbedUnimplementedChatServiceServer))
}

// MockUnsafeChatServiceServer is a mock of UnsafeChatServiceServer interface.
type MockUnsafeChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeChatServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeChatServiceServerMockRecorder is the mock recorder for MockUnsafeChatServiceServer.
type MockUnsafeChatServiceServerMockRecorder struct {
	mock *MockUnsafeChatServiceServer
}

// NewMockUnsafeChatServiceServer creates a new mock instance.
func NewMockUnsafeChatServiceServer(ctrl *gomock.Controller) *MockUnsafeChatServiceServer {
	mock := &MockUnsafeChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeChatServiceServer) EXPECT() *MockUnsafeChatServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedChatServiceServer mocks base method.
func (m *MockUnsafeChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatServiceServer")
}

// mustEmbedUnimplementedChatServiceServer indicates an expected call of mustEmbedUnimplementedChatServiceServer.
func (mr *MockUnsafeChatServiceServerMockRecorder) mustEmbedUnimplementedChatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatServiceServer", reflect.TypeOf((*MockUnsafeChatServiceServer)(nil).mustEmbedUnimplementedChatServiceServer))
}
