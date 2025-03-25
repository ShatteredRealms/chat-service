// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/chat-service/pkg/pb/chat.pb.go
//
// Generated by this command:
//
//	mockgen -source=/home/wil/dev/sro/chat-service/pkg/pb/chat.pb.go -destination=/home/wil/dev/sro/chat-service/pkg/pb/mocks/chat.pb.go
//

// Package mock_pb is a generated GoMock package.
package mock_pb

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockisChatLogRequest_OptionalSenderId is a mock of isChatLogRequest_OptionalSenderId interface.
type MockisChatLogRequest_OptionalSenderId struct {
	ctrl     *gomock.Controller
	recorder *MockisChatLogRequest_OptionalSenderIdMockRecorder
	isgomock struct{}
}

// MockisChatLogRequest_OptionalSenderIdMockRecorder is the mock recorder for MockisChatLogRequest_OptionalSenderId.
type MockisChatLogRequest_OptionalSenderIdMockRecorder struct {
	mock *MockisChatLogRequest_OptionalSenderId
}

// NewMockisChatLogRequest_OptionalSenderId creates a new mock instance.
func NewMockisChatLogRequest_OptionalSenderId(ctrl *gomock.Controller) *MockisChatLogRequest_OptionalSenderId {
	mock := &MockisChatLogRequest_OptionalSenderId{ctrl: ctrl}
	mock.recorder = &MockisChatLogRequest_OptionalSenderIdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisChatLogRequest_OptionalSenderId) EXPECT() *MockisChatLogRequest_OptionalSenderIdMockRecorder {
	return m.recorder
}

// isChatLogRequest_OptionalSenderId mocks base method.
func (m *MockisChatLogRequest_OptionalSenderId) isChatLogRequest_OptionalSenderId() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isChatLogRequest_OptionalSenderId")
}

// isChatLogRequest_OptionalSenderId indicates an expected call of isChatLogRequest_OptionalSenderId.
func (mr *MockisChatLogRequest_OptionalSenderIdMockRecorder) isChatLogRequest_OptionalSenderId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isChatLogRequest_OptionalSenderId", reflect.TypeOf((*MockisChatLogRequest_OptionalSenderId)(nil).isChatLogRequest_OptionalSenderId))
}

// MockisChatLogRequest_OptionalBefore is a mock of isChatLogRequest_OptionalBefore interface.
type MockisChatLogRequest_OptionalBefore struct {
	ctrl     *gomock.Controller
	recorder *MockisChatLogRequest_OptionalBeforeMockRecorder
	isgomock struct{}
}

// MockisChatLogRequest_OptionalBeforeMockRecorder is the mock recorder for MockisChatLogRequest_OptionalBefore.
type MockisChatLogRequest_OptionalBeforeMockRecorder struct {
	mock *MockisChatLogRequest_OptionalBefore
}

// NewMockisChatLogRequest_OptionalBefore creates a new mock instance.
func NewMockisChatLogRequest_OptionalBefore(ctrl *gomock.Controller) *MockisChatLogRequest_OptionalBefore {
	mock := &MockisChatLogRequest_OptionalBefore{ctrl: ctrl}
	mock.recorder = &MockisChatLogRequest_OptionalBeforeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisChatLogRequest_OptionalBefore) EXPECT() *MockisChatLogRequest_OptionalBeforeMockRecorder {
	return m.recorder
}

// isChatLogRequest_OptionalBefore mocks base method.
func (m *MockisChatLogRequest_OptionalBefore) isChatLogRequest_OptionalBefore() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isChatLogRequest_OptionalBefore")
}

// isChatLogRequest_OptionalBefore indicates an expected call of isChatLogRequest_OptionalBefore.
func (mr *MockisChatLogRequest_OptionalBeforeMockRecorder) isChatLogRequest_OptionalBefore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isChatLogRequest_OptionalBefore", reflect.TypeOf((*MockisChatLogRequest_OptionalBefore)(nil).isChatLogRequest_OptionalBefore))
}

// MockisChatLogRequest_OptionalAfter is a mock of isChatLogRequest_OptionalAfter interface.
type MockisChatLogRequest_OptionalAfter struct {
	ctrl     *gomock.Controller
	recorder *MockisChatLogRequest_OptionalAfterMockRecorder
	isgomock struct{}
}

// MockisChatLogRequest_OptionalAfterMockRecorder is the mock recorder for MockisChatLogRequest_OptionalAfter.
type MockisChatLogRequest_OptionalAfterMockRecorder struct {
	mock *MockisChatLogRequest_OptionalAfter
}

// NewMockisChatLogRequest_OptionalAfter creates a new mock instance.
func NewMockisChatLogRequest_OptionalAfter(ctrl *gomock.Controller) *MockisChatLogRequest_OptionalAfter {
	mock := &MockisChatLogRequest_OptionalAfter{ctrl: ctrl}
	mock.recorder = &MockisChatLogRequest_OptionalAfterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisChatLogRequest_OptionalAfter) EXPECT() *MockisChatLogRequest_OptionalAfterMockRecorder {
	return m.recorder
}

// isChatLogRequest_OptionalAfter mocks base method.
func (m *MockisChatLogRequest_OptionalAfter) isChatLogRequest_OptionalAfter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isChatLogRequest_OptionalAfter")
}

// isChatLogRequest_OptionalAfter indicates an expected call of isChatLogRequest_OptionalAfter.
func (mr *MockisChatLogRequest_OptionalAfterMockRecorder) isChatLogRequest_OptionalAfter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isChatLogRequest_OptionalAfter", reflect.TypeOf((*MockisChatLogRequest_OptionalAfter)(nil).isChatLogRequest_OptionalAfter))
}

// MockisChatLogRequest_OptionalLimit is a mock of isChatLogRequest_OptionalLimit interface.
type MockisChatLogRequest_OptionalLimit struct {
	ctrl     *gomock.Controller
	recorder *MockisChatLogRequest_OptionalLimitMockRecorder
	isgomock struct{}
}

// MockisChatLogRequest_OptionalLimitMockRecorder is the mock recorder for MockisChatLogRequest_OptionalLimit.
type MockisChatLogRequest_OptionalLimitMockRecorder struct {
	mock *MockisChatLogRequest_OptionalLimit
}

// NewMockisChatLogRequest_OptionalLimit creates a new mock instance.
func NewMockisChatLogRequest_OptionalLimit(ctrl *gomock.Controller) *MockisChatLogRequest_OptionalLimit {
	mock := &MockisChatLogRequest_OptionalLimit{ctrl: ctrl}
	mock.recorder = &MockisChatLogRequest_OptionalLimitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisChatLogRequest_OptionalLimit) EXPECT() *MockisChatLogRequest_OptionalLimitMockRecorder {
	return m.recorder
}

// isChatLogRequest_OptionalLimit mocks base method.
func (m *MockisChatLogRequest_OptionalLimit) isChatLogRequest_OptionalLimit() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isChatLogRequest_OptionalLimit")
}

// isChatLogRequest_OptionalLimit indicates an expected call of isChatLogRequest_OptionalLimit.
func (mr *MockisChatLogRequest_OptionalLimitMockRecorder) isChatLogRequest_OptionalLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isChatLogRequest_OptionalLimit", reflect.TypeOf((*MockisChatLogRequest_OptionalLimit)(nil).isChatLogRequest_OptionalLimit))
}

// MockisChatLogRequest_OptionalOffset is a mock of isChatLogRequest_OptionalOffset interface.
type MockisChatLogRequest_OptionalOffset struct {
	ctrl     *gomock.Controller
	recorder *MockisChatLogRequest_OptionalOffsetMockRecorder
	isgomock struct{}
}

// MockisChatLogRequest_OptionalOffsetMockRecorder is the mock recorder for MockisChatLogRequest_OptionalOffset.
type MockisChatLogRequest_OptionalOffsetMockRecorder struct {
	mock *MockisChatLogRequest_OptionalOffset
}

// NewMockisChatLogRequest_OptionalOffset creates a new mock instance.
func NewMockisChatLogRequest_OptionalOffset(ctrl *gomock.Controller) *MockisChatLogRequest_OptionalOffset {
	mock := &MockisChatLogRequest_OptionalOffset{ctrl: ctrl}
	mock.recorder = &MockisChatLogRequest_OptionalOffsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisChatLogRequest_OptionalOffset) EXPECT() *MockisChatLogRequest_OptionalOffsetMockRecorder {
	return m.recorder
}

// isChatLogRequest_OptionalOffset mocks base method.
func (m *MockisChatLogRequest_OptionalOffset) isChatLogRequest_OptionalOffset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isChatLogRequest_OptionalOffset")
}

// isChatLogRequest_OptionalOffset indicates an expected call of isChatLogRequest_OptionalOffset.
func (mr *MockisChatLogRequest_OptionalOffsetMockRecorder) isChatLogRequest_OptionalOffset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isChatLogRequest_OptionalOffset", reflect.TypeOf((*MockisChatLogRequest_OptionalOffset)(nil).isChatLogRequest_OptionalOffset))
}

// MockisUpdateChatChannelRequest_OptionalName is a mock of isUpdateChatChannelRequest_OptionalName interface.
type MockisUpdateChatChannelRequest_OptionalName struct {
	ctrl     *gomock.Controller
	recorder *MockisUpdateChatChannelRequest_OptionalNameMockRecorder
	isgomock struct{}
}

// MockisUpdateChatChannelRequest_OptionalNameMockRecorder is the mock recorder for MockisUpdateChatChannelRequest_OptionalName.
type MockisUpdateChatChannelRequest_OptionalNameMockRecorder struct {
	mock *MockisUpdateChatChannelRequest_OptionalName
}

// NewMockisUpdateChatChannelRequest_OptionalName creates a new mock instance.
func NewMockisUpdateChatChannelRequest_OptionalName(ctrl *gomock.Controller) *MockisUpdateChatChannelRequest_OptionalName {
	mock := &MockisUpdateChatChannelRequest_OptionalName{ctrl: ctrl}
	mock.recorder = &MockisUpdateChatChannelRequest_OptionalNameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisUpdateChatChannelRequest_OptionalName) EXPECT() *MockisUpdateChatChannelRequest_OptionalNameMockRecorder {
	return m.recorder
}

// isUpdateChatChannelRequest_OptionalName mocks base method.
func (m *MockisUpdateChatChannelRequest_OptionalName) isUpdateChatChannelRequest_OptionalName() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isUpdateChatChannelRequest_OptionalName")
}

// isUpdateChatChannelRequest_OptionalName indicates an expected call of isUpdateChatChannelRequest_OptionalName.
func (mr *MockisUpdateChatChannelRequest_OptionalNameMockRecorder) isUpdateChatChannelRequest_OptionalName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isUpdateChatChannelRequest_OptionalName", reflect.TypeOf((*MockisUpdateChatChannelRequest_OptionalName)(nil).isUpdateChatChannelRequest_OptionalName))
}

// MockisUpdateChatChannelRequest_OptionalDimension is a mock of isUpdateChatChannelRequest_OptionalDimension interface.
type MockisUpdateChatChannelRequest_OptionalDimension struct {
	ctrl     *gomock.Controller
	recorder *MockisUpdateChatChannelRequest_OptionalDimensionMockRecorder
	isgomock struct{}
}

// MockisUpdateChatChannelRequest_OptionalDimensionMockRecorder is the mock recorder for MockisUpdateChatChannelRequest_OptionalDimension.
type MockisUpdateChatChannelRequest_OptionalDimensionMockRecorder struct {
	mock *MockisUpdateChatChannelRequest_OptionalDimension
}

// NewMockisUpdateChatChannelRequest_OptionalDimension creates a new mock instance.
func NewMockisUpdateChatChannelRequest_OptionalDimension(ctrl *gomock.Controller) *MockisUpdateChatChannelRequest_OptionalDimension {
	mock := &MockisUpdateChatChannelRequest_OptionalDimension{ctrl: ctrl}
	mock.recorder = &MockisUpdateChatChannelRequest_OptionalDimensionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisUpdateChatChannelRequest_OptionalDimension) EXPECT() *MockisUpdateChatChannelRequest_OptionalDimensionMockRecorder {
	return m.recorder
}

// isUpdateChatChannelRequest_OptionalDimension mocks base method.
func (m *MockisUpdateChatChannelRequest_OptionalDimension) isUpdateChatChannelRequest_OptionalDimension() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isUpdateChatChannelRequest_OptionalDimension")
}

// isUpdateChatChannelRequest_OptionalDimension indicates an expected call of isUpdateChatChannelRequest_OptionalDimension.
func (mr *MockisUpdateChatChannelRequest_OptionalDimensionMockRecorder) isUpdateChatChannelRequest_OptionalDimension() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isUpdateChatChannelRequest_OptionalDimension", reflect.TypeOf((*MockisUpdateChatChannelRequest_OptionalDimension)(nil).isUpdateChatChannelRequest_OptionalDimension))
}

// MockisUpdateChatChannelRequest_OptionalPublic is a mock of isUpdateChatChannelRequest_OptionalPublic interface.
type MockisUpdateChatChannelRequest_OptionalPublic struct {
	ctrl     *gomock.Controller
	recorder *MockisUpdateChatChannelRequest_OptionalPublicMockRecorder
	isgomock struct{}
}

// MockisUpdateChatChannelRequest_OptionalPublicMockRecorder is the mock recorder for MockisUpdateChatChannelRequest_OptionalPublic.
type MockisUpdateChatChannelRequest_OptionalPublicMockRecorder struct {
	mock *MockisUpdateChatChannelRequest_OptionalPublic
}

// NewMockisUpdateChatChannelRequest_OptionalPublic creates a new mock instance.
func NewMockisUpdateChatChannelRequest_OptionalPublic(ctrl *gomock.Controller) *MockisUpdateChatChannelRequest_OptionalPublic {
	mock := &MockisUpdateChatChannelRequest_OptionalPublic{ctrl: ctrl}
	mock.recorder = &MockisUpdateChatChannelRequest_OptionalPublicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisUpdateChatChannelRequest_OptionalPublic) EXPECT() *MockisUpdateChatChannelRequest_OptionalPublicMockRecorder {
	return m.recorder
}

// isUpdateChatChannelRequest_OptionalPublic mocks base method.
func (m *MockisUpdateChatChannelRequest_OptionalPublic) isUpdateChatChannelRequest_OptionalPublic() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isUpdateChatChannelRequest_OptionalPublic")
}

// isUpdateChatChannelRequest_OptionalPublic indicates an expected call of isUpdateChatChannelRequest_OptionalPublic.
func (mr *MockisUpdateChatChannelRequest_OptionalPublicMockRecorder) isUpdateChatChannelRequest_OptionalPublic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isUpdateChatChannelRequest_OptionalPublic", reflect.TypeOf((*MockisUpdateChatChannelRequest_OptionalPublic)(nil).isUpdateChatChannelRequest_OptionalPublic))
}
