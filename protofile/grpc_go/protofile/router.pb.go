// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package protofile

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	Type_UNKNOWN_TYPE Type = 0
	Type_AUTH         Type = 1
	Type_SEND         Type = 2
)

var Type_name = map[int32]string{
	0: "UNKNOWN_TYPE",
	1: "AUTH",
	2: "SEND",
}

var Type_value = map[string]int32{
	"UNKNOWN_TYPE": 0,
	"AUTH":         1,
	"SEND":         2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

type Request struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 Type     `protobuf:"varint,3,opt,name=type,proto3,enum=rpc.proto.Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_UNKNOWN_TYPE
}

type Response struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code                 int32    `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterEnum("rpc.proto.Type", Type_name, Type_value)
	proto.RegisterType((*Request)(nil), "rpc.proto.Request")
	proto.RegisterType((*Response)(nil), "rpc.proto.Response")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x40, 0xbb, 0x6d, 0xaa, 0xed, 0x50, 0x34, 0x4e, 0x2f, 0xc5, 0x53, 0x88, 0x97, 0x20, 0x12,
	0x4b, 0x3d, 0x89, 0x27, 0xb5, 0x05, 0x41, 0x88, 0xb2, 0x4d, 0x11, 0xbd, 0x94, 0x98, 0x8e, 0x35,
	0x90, 0x64, 0xd7, 0xcd, 0x46, 0xe8, 0xcf, 0x8b, 0x64, 0x53, 0xaa, 0xde, 0xcc, 0xed, 0xb1, 0xc3,
	0x7b, 0xcc, 0x2c, 0x0c, 0x94, 0x28, 0x35, 0x29, 0x5f, 0x2a, 0xa1, 0x05, 0xf6, 0x95, 0x8c, 0x6b,
	0x74, 0x43, 0xd8, 0xe7, 0xf4, 0x51, 0x52, 0xa1, 0xd1, 0x86, 0x4e, 0x99, 0xac, 0x46, 0xcc, 0x61,
	0x5e, 0x97, 0x57, 0x88, 0x08, 0x56, 0x1e, 0x65, 0x34, 0x6a, 0x3b, 0xcc, 0xeb, 0x73, 0xc3, 0x78,
	0x02, 0x96, 0xde, 0x48, 0x1a, 0x75, 0x1c, 0xe6, 0x1d, 0x4c, 0x0e, 0xfd, 0x5d, 0xca, 0x0f, 0x37,
	0x92, 0xb8, 0x19, 0xba, 0x53, 0xe8, 0x71, 0x2a, 0xa4, 0xc8, 0x0b, 0xfa, 0x67, 0x16, 0xc1, 0x8a,
	0xc5, 0xaa, 0xce, 0x76, 0xb9, 0xe1, 0xd3, 0x33, 0xb0, 0xaa, 0x26, 0xda, 0x30, 0x58, 0x04, 0xf7,
	0xc1, 0xc3, 0x53, 0xb0, 0x0c, 0x9f, 0x1f, 0x67, 0x76, 0x0b, 0x7b, 0x60, 0x5d, 0x2f, 0xc2, 0x3b,
	0x9b, 0x55, 0x34, 0x9f, 0x05, 0x53, 0xbb, 0x3d, 0xf9, 0x62, 0x00, 0x4a, 0xc6, 0xcb, 0x42, 0x2b,
	0x8a, 0x32, 0xbc, 0x84, 0x01, 0x37, 0x37, 0xcf, 0x93, 0x4c, 0xa6, 0x84, 0xf8, 0x6b, 0xd3, 0xed,
	0xc5, 0xc7, 0xc3, 0x3f, 0x6f, 0xf5, 0xbe, 0x6e, 0x0b, 0xaf, 0x76, 0x2a, 0xa9, 0x4f, 0x52, 0x0d,
	0xd4, 0x31, 0xfb, 0x91, 0x6f, 0xd3, 0x84, 0x72, 0xdd, 0x40, 0xf6, 0x2a, 0x19, 0xb6, 0xf2, 0x7b,
	0xd4, 0x4c, 0x1d, 0xb3, 0x9b, 0xe1, 0xcb, 0xd1, 0xba, 0xfa, 0x80, 0xb5, 0x38, 0x37, 0xf3, 0xb7,
	0x24, 0xa5, 0xd7, 0x3d, 0x83, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xec, 0x4b, 0x3a, 0x96,
	0x01, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcStreamClient is the client API for RpcStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcStreamClient interface {
	// 简单rpc
	RouterSimple(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// 服务端流式rpc
	RouterServer(ctx context.Context, in *Request, opts ...grpc.CallOption) (RpcStream_RouterServerClient, error)
	// 客户端流式rpc
	RouterClient(ctx context.Context, opts ...grpc.CallOption) (RpcStream_RouterClientClient, error)
	// 双向流式rpc
	RouterChat(ctx context.Context, opts ...grpc.CallOption) (RpcStream_RouterChatClient, error)
}

type rpcStreamClient struct {
	cc *grpc.ClientConn
}

func NewRpcStreamClient(cc *grpc.ClientConn) RpcStreamClient {
	return &rpcStreamClient{cc}
}

func (c *rpcStreamClient) RouterSimple(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.proto.rpc_stream/RouterSimple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcStreamClient) RouterServer(ctx context.Context, in *Request, opts ...grpc.CallOption) (RpcStream_RouterServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcStream_serviceDesc.Streams[0], "/rpc.proto.rpc_stream/RouterServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcStreamRouterServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RpcStream_RouterServerClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type rpcStreamRouterServerClient struct {
	grpc.ClientStream
}

func (x *rpcStreamRouterServerClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcStreamClient) RouterClient(ctx context.Context, opts ...grpc.CallOption) (RpcStream_RouterClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcStream_serviceDesc.Streams[1], "/rpc.proto.rpc_stream/RouterClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcStreamRouterClientClient{stream}
	return x, nil
}

type RpcStream_RouterClientClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type rpcStreamRouterClientClient struct {
	grpc.ClientStream
}

func (x *rpcStreamRouterClientClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcStreamRouterClientClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcStreamClient) RouterChat(ctx context.Context, opts ...grpc.CallOption) (RpcStream_RouterChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcStream_serviceDesc.Streams[2], "/rpc.proto.rpc_stream/RouterChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcStreamRouterChatClient{stream}
	return x, nil
}

type RpcStream_RouterChatClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type rpcStreamRouterChatClient struct {
	grpc.ClientStream
}

func (x *rpcStreamRouterChatClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcStreamRouterChatClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RpcStreamServer is the server API for RpcStream service.
type RpcStreamServer interface {
	// 简单rpc
	RouterSimple(context.Context, *Request) (*Response, error)
	// 服务端流式rpc
	RouterServer(*Request, RpcStream_RouterServerServer) error
	// 客户端流式rpc
	RouterClient(RpcStream_RouterClientServer) error
	// 双向流式rpc
	RouterChat(RpcStream_RouterChatServer) error
}

// UnimplementedRpcStreamServer can be embedded to have forward compatible implementations.
type UnimplementedRpcStreamServer struct {
}

func (*UnimplementedRpcStreamServer) RouterSimple(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouterSimple not implemented")
}
func (*UnimplementedRpcStreamServer) RouterServer(req *Request, srv RpcStream_RouterServerServer) error {
	return status.Errorf(codes.Unimplemented, "method RouterServer not implemented")
}
func (*UnimplementedRpcStreamServer) RouterClient(srv RpcStream_RouterClientServer) error {
	return status.Errorf(codes.Unimplemented, "method RouterClient not implemented")
}
func (*UnimplementedRpcStreamServer) RouterChat(srv RpcStream_RouterChatServer) error {
	return status.Errorf(codes.Unimplemented, "method RouterChat not implemented")
}

func RegisterRpcStreamServer(s *grpc.Server, srv RpcStreamServer) {
	s.RegisterService(&_RpcStream_serviceDesc, srv)
}

func _RpcStream_RouterSimple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcStreamServer).RouterSimple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.proto.rpc_stream/RouterSimple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcStreamServer).RouterSimple(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcStream_RouterServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RpcStreamServer).RouterServer(m, &rpcStreamRouterServerServer{stream})
}

type RpcStream_RouterServerServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type rpcStreamRouterServerServer struct {
	grpc.ServerStream
}

func (x *rpcStreamRouterServerServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _RpcStream_RouterClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcStreamServer).RouterClient(&rpcStreamRouterClientServer{stream})
}

type RpcStream_RouterClientServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type rpcStreamRouterClientServer struct {
	grpc.ServerStream
}

func (x *rpcStreamRouterClientServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcStreamRouterClientServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RpcStream_RouterChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcStreamServer).RouterChat(&rpcStreamRouterChatServer{stream})
}

type RpcStream_RouterChatServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type rpcStreamRouterChatServer struct {
	grpc.ServerStream
}

func (x *rpcStreamRouterChatServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcStreamRouterChatServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RpcStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.proto.rpc_stream",
	HandlerType: (*RpcStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RouterSimple",
			Handler:    _RpcStream_RouterSimple_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouterServer",
			Handler:       _RpcStream_RouterServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RouterClient",
			Handler:       _RpcStream_RouterClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouterChat",
			Handler:       _RpcStream_RouterChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "router.proto",
}