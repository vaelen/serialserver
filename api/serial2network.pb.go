// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/serial2network.proto

package serial2network

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("api/serial2network.proto", fileDescriptor_54162d2de807133e) }

var fileDescriptor_54162d2de807133e = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2c, 0xc8, 0xd4,
	0x2f, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0x31, 0xca, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92,
	0x4a, 0xd3, 0xf4, 0xcb, 0x8b, 0x12, 0x0b, 0x0a, 0x52, 0x8b, 0x8a, 0x21, 0xf2, 0x46, 0x01, 0x5c,
	0x6c, 0xc1, 0x60, 0x7d, 0x42, 0x6e, 0x5c, 0x2c, 0xfe, 0x05, 0xa9, 0x79, 0x42, 0xd2, 0x7a, 0x10,
	0x2d, 0x7a, 0x30, 0x2d, 0x7a, 0x4e, 0x95, 0x25, 0xa9, 0xc5, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x52,
	0xf8, 0x24, 0x95, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x52, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x38, 0xd0, 0x9c, 0x88, 0x94, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SerialClient is the client API for Serial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SerialClient interface {
	// Sends and receives data from the serial port
	Open(ctx context.Context, opts ...grpc.CallOption) (Serial_OpenClient, error)
}

type serialClient struct {
	cc *grpc.ClientConn
}

func NewSerialClient(cc *grpc.ClientConn) SerialClient {
	return &serialClient{cc}
}

func (c *serialClient) Open(ctx context.Context, opts ...grpc.CallOption) (Serial_OpenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Serial_serviceDesc.Streams[0], "/Serial/Open", opts...)
	if err != nil {
		return nil, err
	}
	x := &serialOpenClient{stream}
	return x, nil
}

type Serial_OpenClient interface {
	Send(*wrappers.BytesValue) error
	Recv() (*wrappers.BytesValue, error)
	grpc.ClientStream
}

type serialOpenClient struct {
	grpc.ClientStream
}

func (x *serialOpenClient) Send(m *wrappers.BytesValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serialOpenClient) Recv() (*wrappers.BytesValue, error) {
	m := new(wrappers.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SerialServer is the server API for Serial service.
type SerialServer interface {
	// Sends and receives data from the serial port
	Open(Serial_OpenServer) error
}

func RegisterSerialServer(s *grpc.Server, srv SerialServer) {
	s.RegisterService(&_Serial_serviceDesc, srv)
}

func _Serial_Open_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SerialServer).Open(&serialOpenServer{stream})
}

type Serial_OpenServer interface {
	Send(*wrappers.BytesValue) error
	Recv() (*wrappers.BytesValue, error)
	grpc.ServerStream
}

type serialOpenServer struct {
	grpc.ServerStream
}

func (x *serialOpenServer) Send(m *wrappers.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serialOpenServer) Recv() (*wrappers.BytesValue, error) {
	m := new(wrappers.BytesValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Serial_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Serial",
	HandlerType: (*SerialServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Open",
			Handler:       _Serial_Open_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/serial2network.proto",
}