// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bar "github.com/tokopedia/gripmock/example/multi-package/bar"

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

type Response struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_8f5f45fceee2f0f1, []int{0}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "hello.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GripmockClient is the client API for Gripmock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GripmockClient interface {
	Greet(ctx context.Context, in *bar.Bar, opts ...grpc.CallOption) (*Response, error)
}

type gripmockClient struct {
	cc *grpc.ClientConn
}

func NewGripmockClient(cc *grpc.ClientConn) GripmockClient {
	return &gripmockClient{cc}
}

func (c *gripmockClient) Greet(ctx context.Context, in *bar.Bar, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/hello.Gripmock/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripmockServer is the server API for Gripmock service.
type GripmockServer interface {
	Greet(context.Context, *bar.Bar) (*Response, error)
}

func RegisterGripmockServer(s *grpc.Server, srv GripmockServer) {
	s.RegisterService(&_Gripmock_serviceDesc, srv)
}

func _Gripmock_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bar.Bar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripmockServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Gripmock/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripmockServer).Greet(ctx, req.(*bar.Bar))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gripmock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Gripmock",
	HandlerType: (*GripmockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Gripmock_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_8f5f45fceee2f0f1) }

var fileDescriptor_hello_8f5f45fceee2f0f1 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x78, 0x93, 0x12, 0x8b,
	0xf4, 0x93, 0x12, 0x8b, 0x20, 0xa2, 0x4a, 0x6a, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x45, 0x50, 0xb6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x9c, 0x6f, 0xa4, 0xc7, 0xc5, 0xe1, 0x5e, 0x94, 0x59, 0x90, 0x9b, 0x9f, 0x9c, 0x2d, 0xa4, 0xc4,
	0xc5, 0xea, 0x5e, 0x94, 0x9a, 0x5a, 0x22, 0xc4, 0xa1, 0x07, 0x32, 0xc8, 0x29, 0xb1, 0x48, 0x8a,
	0x5f, 0x0f, 0x62, 0x15, 0xcc, 0xac, 0x24, 0x36, 0xb0, 0xf1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x72, 0x3b, 0x00, 0x6e, 0x83, 0x00, 0x00, 0x00,
}