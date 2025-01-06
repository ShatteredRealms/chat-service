// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: sro/chat/chat.proto

package pb

import (
	context "context"
	pb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChatService_ConnectChatChannel_FullMethodName             = "/sro.chat.ChatService/ConnectChatChannel"
	ChatService_ConnectDirectMessages_FullMethodName          = "/sro.chat.ChatService/ConnectDirectMessages"
	ChatService_SendChatChannelMessage_FullMethodName         = "/sro.chat.ChatService/SendChatChannelMessage"
	ChatService_SendDirectMessage_FullMethodName              = "/sro.chat.ChatService/SendDirectMessage"
	ChatService_GetChatChannels_FullMethodName                = "/sro.chat.ChatService/GetChatChannels"
	ChatService_GetChatChannel_FullMethodName                 = "/sro.chat.ChatService/GetChatChannel"
	ChatService_CreateChatChannel_FullMethodName              = "/sro.chat.ChatService/CreateChatChannel"
	ChatService_DeleteChatChannel_FullMethodName              = "/sro.chat.ChatService/DeleteChatChannel"
	ChatService_EditChatChannel_FullMethodName                = "/sro.chat.ChatService/EditChatChannel"
	ChatService_GetAuthorizedChatChannels_FullMethodName      = "/sro.chat.ChatService/GetAuthorizedChatChannels"
	ChatService_SetCharacterChatChannelAuth_FullMethodName    = "/sro.chat.ChatService/SetCharacterChatChannelAuth"
	ChatService_UpdateCharacterChatChannelAuth_FullMethodName = "/sro.chat.ChatService/UpdateCharacterChatChannelAuth"
	ChatService_BanCharacterFromChatChannel_FullMethodName    = "/sro.chat.ChatService/BanCharacterFromChatChannel"
	ChatService_GetChatLogs_FullMethodName                    = "/sro.chat.ChatService/GetChatLogs"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	ConnectChatChannel(ctx context.Context, in *ConnectChatChannelRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatMessage], error)
	ConnectDirectMessages(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatMessage], error)
	SendChatChannelMessage(ctx context.Context, in *SendChatChannelMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendDirectMessage(ctx context.Context, in *SendDirectMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// TODO: Have request allow for filtering
	GetChatChannels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChatChannels, error)
	GetChatChannel(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*ChatChannel, error)
	CreateChatChannel(ctx context.Context, in *CreateChatChannelMessage, opts ...grpc.CallOption) (*ChatChannel, error)
	DeleteChatChannel(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*ChatChannel, error)
	EditChatChannel(ctx context.Context, in *UpdateChatChannelRequest, opts ...grpc.CallOption) (*ChatChannel, error)
	GetAuthorizedChatChannels(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*ChatChannels, error)
	// Sets the character chat channels to the given list of channels
	SetCharacterChatChannelAuth(ctx context.Context, in *RequestSetCharacterSetChatChannelAuth, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// If add is true, adds the given channels to the character's chat channels,
	// otherwise removes them
	UpdateCharacterChatChannelAuth(ctx context.Context, in *RequestUpdateCharacterSetChatChannelAuth, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BanCharacterFromChatChannel(ctx context.Context, in *BanRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetChatLogs(ctx context.Context, in *ChatLogRequest, opts ...grpc.CallOption) (*ChatLogs, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ConnectChatChannel(ctx context.Context, in *ConnectChatChannelRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_ConnectChatChannel_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ConnectChatChannelRequest, ChatMessage]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ConnectChatChannelClient = grpc.ServerStreamingClient[ChatMessage]

func (c *chatServiceClient) ConnectDirectMessages(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_ConnectDirectMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[pb.TargetId, ChatMessage]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ConnectDirectMessagesClient = grpc.ServerStreamingClient[ChatMessage]

func (c *chatServiceClient) SendChatChannelMessage(ctx context.Context, in *SendChatChannelMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_SendChatChannelMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendDirectMessage(ctx context.Context, in *SendDirectMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_SendDirectMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatChannels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChatChannels, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatChannels)
	err := c.cc.Invoke(ctx, ChatService_GetChatChannels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatChannel(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*ChatChannel, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatChannel)
	err := c.cc.Invoke(ctx, ChatService_GetChatChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateChatChannel(ctx context.Context, in *CreateChatChannelMessage, opts ...grpc.CallOption) (*ChatChannel, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatChannel)
	err := c.cc.Invoke(ctx, ChatService_CreateChatChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteChatChannel(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*ChatChannel, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatChannel)
	err := c.cc.Invoke(ctx, ChatService_DeleteChatChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EditChatChannel(ctx context.Context, in *UpdateChatChannelRequest, opts ...grpc.CallOption) (*ChatChannel, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatChannel)
	err := c.cc.Invoke(ctx, ChatService_EditChatChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetAuthorizedChatChannels(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*ChatChannels, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatChannels)
	err := c.cc.Invoke(ctx, ChatService_GetAuthorizedChatChannels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SetCharacterChatChannelAuth(ctx context.Context, in *RequestSetCharacterSetChatChannelAuth, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_SetCharacterChatChannelAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateCharacterChatChannelAuth(ctx context.Context, in *RequestUpdateCharacterSetChatChannelAuth, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_UpdateCharacterChatChannelAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) BanCharacterFromChatChannel(ctx context.Context, in *BanRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_BanCharacterFromChatChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatLogs(ctx context.Context, in *ChatLogRequest, opts ...grpc.CallOption) (*ChatLogs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatLogs)
	err := c.cc.Invoke(ctx, ChatService_GetChatLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility.
type ChatServiceServer interface {
	ConnectChatChannel(*ConnectChatChannelRequest, grpc.ServerStreamingServer[ChatMessage]) error
	ConnectDirectMessages(*pb.TargetId, grpc.ServerStreamingServer[ChatMessage]) error
	SendChatChannelMessage(context.Context, *SendChatChannelMessageRequest) (*emptypb.Empty, error)
	SendDirectMessage(context.Context, *SendDirectMessageRequest) (*emptypb.Empty, error)
	// TODO: Have request allow for filtering
	GetChatChannels(context.Context, *emptypb.Empty) (*ChatChannels, error)
	GetChatChannel(context.Context, *pb.TargetId) (*ChatChannel, error)
	CreateChatChannel(context.Context, *CreateChatChannelMessage) (*ChatChannel, error)
	DeleteChatChannel(context.Context, *pb.TargetId) (*ChatChannel, error)
	EditChatChannel(context.Context, *UpdateChatChannelRequest) (*ChatChannel, error)
	GetAuthorizedChatChannels(context.Context, *pb.TargetId) (*ChatChannels, error)
	// Sets the character chat channels to the given list of channels
	SetCharacterChatChannelAuth(context.Context, *RequestSetCharacterSetChatChannelAuth) (*emptypb.Empty, error)
	// If add is true, adds the given channels to the character's chat channels,
	// otherwise removes them
	UpdateCharacterChatChannelAuth(context.Context, *RequestUpdateCharacterSetChatChannelAuth) (*emptypb.Empty, error)
	BanCharacterFromChatChannel(context.Context, *BanRequest) (*emptypb.Empty, error)
	GetChatLogs(context.Context, *ChatLogRequest) (*ChatLogs, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) ConnectChatChannel(*ConnectChatChannelRequest, grpc.ServerStreamingServer[ChatMessage]) error {
	return status.Errorf(codes.Unimplemented, "method ConnectChatChannel not implemented")
}
func (UnimplementedChatServiceServer) ConnectDirectMessages(*pb.TargetId, grpc.ServerStreamingServer[ChatMessage]) error {
	return status.Errorf(codes.Unimplemented, "method ConnectDirectMessages not implemented")
}
func (UnimplementedChatServiceServer) SendChatChannelMessage(context.Context, *SendChatChannelMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChatChannelMessage not implemented")
}
func (UnimplementedChatServiceServer) SendDirectMessage(context.Context, *SendDirectMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDirectMessage not implemented")
}
func (UnimplementedChatServiceServer) GetChatChannels(context.Context, *emptypb.Empty) (*ChatChannels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatChannels not implemented")
}
func (UnimplementedChatServiceServer) GetChatChannel(context.Context, *pb.TargetId) (*ChatChannel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatChannel not implemented")
}
func (UnimplementedChatServiceServer) CreateChatChannel(context.Context, *CreateChatChannelMessage) (*ChatChannel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatChannel not implemented")
}
func (UnimplementedChatServiceServer) DeleteChatChannel(context.Context, *pb.TargetId) (*ChatChannel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChatChannel not implemented")
}
func (UnimplementedChatServiceServer) EditChatChannel(context.Context, *UpdateChatChannelRequest) (*ChatChannel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditChatChannel not implemented")
}
func (UnimplementedChatServiceServer) GetAuthorizedChatChannels(context.Context, *pb.TargetId) (*ChatChannels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorizedChatChannels not implemented")
}
func (UnimplementedChatServiceServer) SetCharacterChatChannelAuth(context.Context, *RequestSetCharacterSetChatChannelAuth) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCharacterChatChannelAuth not implemented")
}
func (UnimplementedChatServiceServer) UpdateCharacterChatChannelAuth(context.Context, *RequestUpdateCharacterSetChatChannelAuth) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCharacterChatChannelAuth not implemented")
}
func (UnimplementedChatServiceServer) BanCharacterFromChatChannel(context.Context, *BanRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanCharacterFromChatChannel not implemented")
}
func (UnimplementedChatServiceServer) GetChatLogs(context.Context, *ChatLogRequest) (*ChatLogs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatLogs not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}
func (UnimplementedChatServiceServer) testEmbeddedByValue()                     {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_ConnectChatChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectChatChannelRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ConnectChatChannel(m, &grpc.GenericServerStream[ConnectChatChannelRequest, ChatMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ConnectChatChannelServer = grpc.ServerStreamingServer[ChatMessage]

func _ChatService_ConnectDirectMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pb.TargetId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ConnectDirectMessages(m, &grpc.GenericServerStream[pb.TargetId, ChatMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ConnectDirectMessagesServer = grpc.ServerStreamingServer[ChatMessage]

func _ChatService_SendChatChannelMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendChatChannelMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendChatChannelMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendChatChannelMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendChatChannelMessage(ctx, req.(*SendChatChannelMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendDirectMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDirectMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendDirectMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendDirectMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendDirectMessage(ctx, req.(*SendDirectMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatChannels(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TargetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatChannel(ctx, req.(*pb.TargetId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateChatChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatChannelMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChatChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateChatChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChatChannel(ctx, req.(*CreateChatChannelMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteChatChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TargetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteChatChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteChatChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteChatChannel(ctx, req.(*pb.TargetId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EditChatChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChatChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EditChatChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_EditChatChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EditChatChannel(ctx, req.(*UpdateChatChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetAuthorizedChatChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TargetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetAuthorizedChatChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetAuthorizedChatChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetAuthorizedChatChannels(ctx, req.(*pb.TargetId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SetCharacterChatChannelAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSetCharacterSetChatChannelAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SetCharacterChatChannelAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SetCharacterChatChannelAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SetCharacterChatChannelAuth(ctx, req.(*RequestSetCharacterSetChatChannelAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateCharacterChatChannelAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateCharacterSetChatChannelAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateCharacterChatChannelAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateCharacterChatChannelAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateCharacterChatChannelAuth(ctx, req.(*RequestUpdateCharacterSetChatChannelAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_BanCharacterFromChatChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).BanCharacterFromChatChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_BanCharacterFromChatChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).BanCharacterFromChatChannel(ctx, req.(*BanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatLogs(ctx, req.(*ChatLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sro.chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChatChannelMessage",
			Handler:    _ChatService_SendChatChannelMessage_Handler,
		},
		{
			MethodName: "SendDirectMessage",
			Handler:    _ChatService_SendDirectMessage_Handler,
		},
		{
			MethodName: "GetChatChannels",
			Handler:    _ChatService_GetChatChannels_Handler,
		},
		{
			MethodName: "GetChatChannel",
			Handler:    _ChatService_GetChatChannel_Handler,
		},
		{
			MethodName: "CreateChatChannel",
			Handler:    _ChatService_CreateChatChannel_Handler,
		},
		{
			MethodName: "DeleteChatChannel",
			Handler:    _ChatService_DeleteChatChannel_Handler,
		},
		{
			MethodName: "EditChatChannel",
			Handler:    _ChatService_EditChatChannel_Handler,
		},
		{
			MethodName: "GetAuthorizedChatChannels",
			Handler:    _ChatService_GetAuthorizedChatChannels_Handler,
		},
		{
			MethodName: "SetCharacterChatChannelAuth",
			Handler:    _ChatService_SetCharacterChatChannelAuth_Handler,
		},
		{
			MethodName: "UpdateCharacterChatChannelAuth",
			Handler:    _ChatService_UpdateCharacterChatChannelAuth_Handler,
		},
		{
			MethodName: "BanCharacterFromChatChannel",
			Handler:    _ChatService_BanCharacterFromChatChannel_Handler,
		},
		{
			MethodName: "GetChatLogs",
			Handler:    _ChatService_GetChatLogs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectChatChannel",
			Handler:       _ChatService_ConnectChatChannel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ConnectDirectMessages",
			Handler:       _ChatService_ConnectDirectMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sro/chat/chat.proto",
}
