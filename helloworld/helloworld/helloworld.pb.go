// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld/helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld/helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloAgainClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloAgainClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/helloworld.Greeter/SayHelloAgain", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloAgainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloAgainClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloAgainClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloAgainClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(*HelloRequest, Greeter_SayHelloAgainServer) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloAgain(m, &greeterSayHelloAgainServer{stream})
}

type Greeter_SayHelloAgainServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayHelloAgainServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloAgainServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloAgain",
			Handler:       _Greeter_SayHelloAgain_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld/helloworld.proto",
}

func init() { proto.RegisterFile("helloworld/helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x22, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0x20, 0x5e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x98, 0xad, 0xa4, 0xc6, 0xc5, 0x05, 0x55, 0x53, 0x90, 0x53, 0x29, 0x24, 0xc1, 0xc5, 0x9e, 0x9b,
	0x5a, 0x5c, 0x9c, 0x98, 0x0e, 0x53, 0x04, 0xe3, 0x1a, 0x4d, 0x60, 0xe4, 0x62, 0x77, 0x2f, 0x4a,
	0x4d, 0x2d, 0x49, 0x2d, 0x12, 0xb2, 0xe3, 0xe2, 0x08, 0x4e, 0xac, 0x04, 0x6b, 0x13, 0x92, 0xd0,
	0x43, 0x72, 0x02, 0xb2, 0x6d, 0x52, 0x62, 0x58, 0x64, 0x0a, 0x72, 0x2a, 0x95, 0x18, 0x84, 0x5c,
	0xb9, 0x78, 0x61, 0xfa, 0x1d, 0xd3, 0x13, 0x33, 0xf3, 0xc8, 0x31, 0xc4, 0x80, 0xd1, 0xc9, 0x80,
	0x4b, 0x3a, 0x33, 0x5f, 0x2f, 0xbd, 0xa8, 0x20, 0x59, 0x2f, 0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27,
	0xb5, 0x18, 0x49, 0xb5, 0x13, 0x3f, 0x58, 0x79, 0x38, 0x88, 0x1d, 0x00, 0x0a, 0x9a, 0x00, 0xc6,
	0x24, 0x36, 0x70, 0x18, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x92, 0x64, 0xee, 0x11, 0x42,
	0x01, 0x00, 0x00,
}
