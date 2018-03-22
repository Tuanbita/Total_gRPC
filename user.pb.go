// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserRequest
	UserResponse
	UserCondition
	UserSummary
	UserMessage
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRequest struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UserResponse struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age    uint32 `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Gender uint32 `protobuf:"varint,3,opt,name=gender" json:"gender,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserResponse) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

type UserCondition struct {
	Gender uint32 `protobuf:"varint,1,opt,name=gender" json:"gender,omitempty"`
}

func (m *UserCondition) Reset()                    { *m = UserCondition{} }
func (m *UserCondition) String() string            { return proto.CompactTextString(m) }
func (*UserCondition) ProtoMessage()               {}
func (*UserCondition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserCondition) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

type UserSummary struct {
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Total       uint32 `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *UserSummary) Reset()                    { *m = UserSummary{} }
func (m *UserSummary) String() string            { return proto.CompactTextString(m) }
func (*UserSummary) ProtoMessage()               {}
func (*UserSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserSummary) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UserSummary) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type UserMessage struct {
	Dem  int32  `protobuf:"varint,1,opt,name=dem" json:"dem,omitempty"`
	Talk string `protobuf:"bytes,2,opt,name=talk" json:"talk,omitempty"`
}

func (m *UserMessage) Reset()                    { *m = UserMessage{} }
func (m *UserMessage) String() string            { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()               {}
func (*UserMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserMessage) GetDem() int32 {
	if m != nil {
		return m.Dem
	}
	return 0
}

func (m *UserMessage) GetTalk() string {
	if m != nil {
		return m.Talk
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "UserRequest")
	proto.RegisterType((*UserResponse)(nil), "UserResponse")
	proto.RegisterType((*UserCondition)(nil), "UserCondition")
	proto.RegisterType((*UserSummary)(nil), "UserSummary")
	proto.RegisterType((*UserMessage)(nil), "UserMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	QueryUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ListUser(ctx context.Context, in *UserCondition, opts ...grpc.CallOption) (User_ListUserClient, error)
	SendUser(ctx context.Context, opts ...grpc.CallOption) (User_SendUserClient, error)
	Chat(ctx context.Context, opts ...grpc.CallOption) (User_ChatClient, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) QueryUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/User/QueryUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUser(ctx context.Context, in *UserCondition, opts ...grpc.CallOption) (User_ListUserClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_User_serviceDesc.Streams[0], c.cc, "/User/ListUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &userListUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type User_ListUserClient interface {
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type userListUserClient struct {
	grpc.ClientStream
}

func (x *userListUserClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userClient) SendUser(ctx context.Context, opts ...grpc.CallOption) (User_SendUserClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_User_serviceDesc.Streams[1], c.cc, "/User/SendUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &userSendUserClient{stream}
	return x, nil
}

type User_SendUserClient interface {
	Send(*UserRequest) error
	CloseAndRecv() (*UserSummary, error)
	grpc.ClientStream
}

type userSendUserClient struct {
	grpc.ClientStream
}

func (x *userSendUserClient) Send(m *UserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userSendUserClient) CloseAndRecv() (*UserSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UserSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userClient) Chat(ctx context.Context, opts ...grpc.CallOption) (User_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_User_serviceDesc.Streams[2], c.cc, "/User/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &userChatClient{stream}
	return x, nil
}

type User_ChatClient interface {
	Send(*UserMessage) error
	Recv() (*UserMessage, error)
	grpc.ClientStream
}

type userChatClient struct {
	grpc.ClientStream
}

func (x *userChatClient) Send(m *UserMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userChatClient) Recv() (*UserMessage, error) {
	m := new(UserMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for User service

type UserServer interface {
	QueryUser(context.Context, *UserRequest) (*UserResponse, error)
	ListUser(*UserCondition, User_ListUserServer) error
	SendUser(User_SendUserServer) error
	Chat(User_ChatServer) error
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_QueryUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).QueryUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/QueryUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).QueryUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserCondition)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServer).ListUser(m, &userListUserServer{stream})
}

type User_ListUserServer interface {
	Send(*UserResponse) error
	grpc.ServerStream
}

type userListUserServer struct {
	grpc.ServerStream
}

func (x *userListUserServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _User_SendUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServer).SendUser(&userSendUserServer{stream})
}

type User_SendUserServer interface {
	SendAndClose(*UserSummary) error
	Recv() (*UserRequest, error)
	grpc.ServerStream
}

type userSendUserServer struct {
	grpc.ServerStream
}

func (x *userSendUserServer) SendAndClose(m *UserSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userSendUserServer) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _User_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServer).Chat(&userChatServer{stream})
}

type User_ChatServer interface {
	Send(*UserMessage) error
	Recv() (*UserMessage, error)
	grpc.ServerStream
}

type userChatServer struct {
	grpc.ServerStream
}

func (x *userChatServer) Send(m *UserMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userChatServer) Recv() (*UserMessage, error) {
	m := new(UserMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryUser",
			Handler:    _User_QueryUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUser",
			Handler:       _User_ListUser_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendUser",
			Handler:       _User_SendUser_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _User_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0xa3, 0xdf, 0x49, 0x88, 0x27, 0xf1, 0x4f, 0x10, 0xa5, 0x98, 0x6c, 0x6a, 0xb4, 0xa9,
	0x09, 0xd4, 0x84, 0xe6, 0x08, 0xa1, 0xbb, 0x74, 0x51, 0x85, 0x1e, 0x40, 0xad, 0x86, 0xd4, 0x34,
	0x96, 0x5c, 0x49, 0x5e, 0xe4, 0x64, 0xbd, 0x5e, 0x91, 0xac, 0x85, 0x4d, 0xbb, 0x7b, 0x4f, 0x7c,
	0x6f, 0x3c, 0x6f, 0x0c, 0xd0, 0x59, 0x34, 0x55, 0x6b, 0xb4, 0xd3, 0xec, 0x0e, 0x96, 0xaf, 0x16,
	0x0d, 0xc7, 0xaf, 0x0e, 0xad, 0xa3, 0x6b, 0x48, 0xba, 0x5a, 0xe6, 0xa4, 0x20, 0x65, 0xc2, 0xbd,
	0x64, 0x47, 0x58, 0xf5, 0x80, 0x6d, 0xb5, 0xb2, 0x48, 0x29, 0x4c, 0x95, 0x68, 0x30, 0x20, 0x29,
	0x0f, 0xda, 0xa7, 0xc4, 0x19, 0xf3, 0x7f, 0x05, 0x29, 0x33, 0xee, 0x25, 0xbd, 0x85, 0xf9, 0x19,
	0x95, 0x44, 0x93, 0x27, 0xe1, 0x31, 0x3a, 0x76, 0x0f, 0x99, 0x9f, 0x76, 0xd0, 0x4a, 0xd6, 0xae,
	0xd6, 0x6a, 0x00, 0x92, 0x11, 0xf8, 0xd4, 0xef, 0x75, 0xea, 0x9a, 0x46, 0x98, 0x2b, 0x2d, 0x60,
	0x29, 0xd1, 0xbe, 0x9b, 0xba, 0xf5, 0xa9, 0xf8, 0xf1, 0xe1, 0x13, 0xbd, 0x81, 0x99, 0xd3, 0x4e,
	0x5c, 0xe2, 0x16, 0xbd, 0x61, 0xfb, 0x7e, 0xcc, 0x33, 0x5a, 0xeb, 0xd7, 0x5a, 0x43, 0x22, 0xb1,
	0x09, 0xf1, 0x19, 0xf7, 0xd2, 0xd7, 0x71, 0xe2, 0xf2, 0x19, 0x52, 0x29, 0x0f, 0xfa, 0xf1, 0x9b,
	0xc0, 0xd4, 0xa7, 0xe8, 0x16, 0xd2, 0x97, 0x0e, 0xcd, 0x35, 0x98, 0x55, 0x35, 0x38, 0xd4, 0x26,
	0xab, 0x86, 0x57, 0x61, 0x13, 0xfa, 0x00, 0x8b, 0x63, 0x6d, 0x5d, 0x40, 0xff, 0x57, 0xa3, 0x92,
	0xbf, 0xe0, 0x1d, 0xa1, 0x5b, 0x58, 0x9c, 0x50, 0xc9, 0x3f, 0x26, 0xf7, 0x2e, 0x16, 0x67, 0x93,
	0x92, 0xd0, 0x12, 0xa6, 0x87, 0x0f, 0xe1, 0x22, 0x17, 0xbb, 0x6c, 0x46, 0xce, 0x73, 0x3b, 0xf2,
	0x36, 0x0f, 0x3f, 0x75, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x20, 0x8e, 0x6b, 0x47, 0xe2, 0x01,
	0x00, 0x00,
}
