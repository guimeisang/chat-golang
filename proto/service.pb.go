package proto

import (
	"context"
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_40d27f16416a858b, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_40d27f16416a858b, []int{1}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type Connect struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_40d27f16416a858b, []int{2}
}
func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (dst *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(dst, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Connect) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type Close struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_40d27f16416a858b, []int{3}
}
func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (dst *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(dst, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Message)(nil), "proto.Message")
	proto.RegisterType((*Connect)(nil), "proto.Connect")
	proto.RegisterType((*Close)(nil), "proto.Close")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BroadcastClient is the client API for Broadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BroadcastClient interface {
	CreateStream(ctx context.Context, in *Connect, opts ...grpc.CallOption) (Broadcast_CreateStreamClient, error)
	BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error)
}

type broadcastClient struct {
	cc *grpc.ClientConn
}

func NewBroadcastClient(cc *grpc.ClientConn) BroadcastClient {
	return &broadcastClient{cc}
}

func (c *broadcastClient) CreateStream(ctx context.Context, in *Connect, opts ...grpc.CallOption) (Broadcast_CreateStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Broadcast_serviceDesc.Streams[0], "/proto.Broadcast/CreateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &broadcastCreateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broadcast_CreateStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type broadcastCreateStreamClient struct {
	grpc.ClientStream
}

func (x *broadcastCreateStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *broadcastClient) BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error) {
	out := new(Close)
	err := c.cc.Invoke(ctx, "/proto.Broadcast/BroadcastMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServer is the server API for Broadcast service.
type BroadcastServer interface {
	CreateStream(*Connect, Broadcast_CreateStreamServer) error
	BroadcastMessage(context.Context, *Message) (*Close, error)
}

func RegisterBroadcastServer(s *grpc.Server, srv BroadcastServer) {
	s.RegisterService(&_Broadcast_serviceDesc, srv)
}

func _Broadcast_CreateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Connect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BroadcastServer).CreateStream(m, &broadcastCreateStreamServer{stream})
}

type Broadcast_CreateStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type broadcastCreateStreamServer struct {
	grpc.ServerStream
}

func (x *broadcastCreateStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Broadcast_BroadcastMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).BroadcastMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Broadcast/BroadcastMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).BroadcastMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broadcast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Broadcast",
	HandlerType: (*BroadcastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastMessage",
			Handler:    _Broadcast_BroadcastMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateStream",
			Handler:       _Broadcast_CreateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_40d27f16416a858b) }

var fileDescriptor_service_40d27f16416a858b = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x4c, 0x1b, 0x33, 0xad, 0x45, 0xe6, 0x20, 0xa1, 0x08, 0x4a, 0x4e, 0xe2, 0xa1,
	0x84, 0xfa, 0x06, 0xcd, 0xd9, 0x83, 0x11, 0x1f, 0x60, 0xdc, 0x0c, 0xb2, 0x60, 0x76, 0xcb, 0xce,
	0xd8, 0xe7, 0x97, 0x6e, 0x37, 0x15, 0xe9, 0x69, 0xf7, 0x37, 0x7f, 0xbe, 0xf9, 0x3e, 0xb8, 0x11,
	0x0e, 0x07, 0x6b, 0x78, 0xb3, 0x0f, 0x5e, 0x3d, 0xce, 0xe2, 0xd3, 0x3c, 0x43, 0xf1, 0x21, 0x1c,
	0x70, 0x05, 0xb9, 0x1d, 0xea, 0xec, 0x31, 0x7b, 0xaa, 0xfa, 0xdc, 0x0e, 0x88, 0x50, 0x38, 0x1a,
	0xb9, 0xce, 0x63, 0x25, 0xfe, 0x9b, 0x37, 0x28, 0x5f, 0x59, 0x84, 0xbe, 0xf8, 0x62, 0xbc, 0x86,
	0xd2, 0x78, 0xa7, 0xec, 0x34, 0x6d, 0x4c, 0x88, 0xf7, 0x50, 0xa9, 0x1d, 0x59, 0x94, 0xc6, 0x7d,
	0x7d, 0x15, 0x7b, 0x7f, 0x85, 0x66, 0x07, 0x65, 0xe7, 0x9d, 0x63, 0xa3, 0xf8, 0x00, 0xc5, 0x8f,
	0x70, 0x88, 0xa2, 0x8b, 0xed, 0xe2, 0x64, 0x73, 0x73, 0x34, 0xd7, 0xc7, 0x06, 0xde, 0xc1, 0x9c,
	0x8c, 0xda, 0xc3, 0xc9, 0xd4, 0x75, 0x9f, 0xa8, 0x29, 0x61, 0xd6, 0x7d, 0x7b, 0xe1, 0xad, 0x87,
	0x6a, 0x17, 0x3c, 0x0d, 0x86, 0x44, 0xb1, 0x85, 0x65, 0x17, 0x98, 0x94, 0xdf, 0x35, 0x30, 0x8d,
	0xb8, 0x4a, 0x82, 0xe9, 0xdc, 0x7a, 0xe2, 0x94, 0xa8, 0xcd, 0xb0, 0x85, 0xdb, 0xf3, 0xfa, 0x39,
	0xe7, 0xff, 0xa9, 0xf5, 0x72, 0x52, 0x39, 0x1e, 0xfc, 0x9c, 0x47, 0x78, 0xf9, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x3f, 0x75, 0x4e, 0x77, 0x5b, 0x01, 0x00, 0x00,
}
