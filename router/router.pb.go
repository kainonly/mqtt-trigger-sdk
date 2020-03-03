// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package router

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

type PutParameter struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Queue                string   `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutParameter) Reset()         { *m = PutParameter{} }
func (m *PutParameter) String() string { return proto.CompactTextString(m) }
func (*PutParameter) ProtoMessage()    {}
func (*PutParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *PutParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutParameter.Unmarshal(m, b)
}
func (m *PutParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutParameter.Marshal(b, m, deterministic)
}
func (m *PutParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutParameter.Merge(m, src)
}
func (m *PutParameter) XXX_Size() int {
	return xxx_messageInfo_PutParameter.Size(m)
}
func (m *PutParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PutParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PutParameter proto.InternalMessageInfo

func (m *PutParameter) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *PutParameter) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

type PutResponse struct {
	Error                uint32   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *PutResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*PutParameter)(nil), "PutParameter")
	proto.RegisterType((*PutResponse)(nil), "PutResponse")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe0, 0xe2, 0x09, 0x28, 0x2d, 0x09,
	0x48, 0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0x12, 0x92, 0xe2, 0xe2, 0xc8, 0x4c, 0x49, 0xcd,
	0x2b, 0xc9, 0x2c, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x44, 0xb8,
	0x58, 0x0b, 0x4b, 0x53, 0x4b, 0x53, 0x25, 0x98, 0xc0, 0x12, 0x10, 0x8e, 0x92, 0x29, 0x17, 0x77,
	0x40, 0x69, 0x49, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x51, 0x6a, 0x51, 0x51,
	0x7e, 0x11, 0x58, 0x37, 0x6f, 0x10, 0x84, 0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x5b, 0x9c, 0x0e, 0xd5,
	0x08, 0x62, 0x1a, 0xe9, 0x71, 0xb1, 0x05, 0x81, 0x1d, 0x22, 0xa4, 0xc2, 0xc5, 0x1c, 0x50, 0x5a,
	0x22, 0xc4, 0xab, 0x87, 0xec, 0x10, 0x29, 0x1e, 0x3d, 0x24, 0x53, 0x95, 0x18, 0x92, 0xd8, 0xc0,
	0xee, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xe1, 0xef, 0xb3, 0xbf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	Put(ctx context.Context, in *PutParameter, opts ...grpc.CallOption) (*PutResponse, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Put(ctx context.Context, in *PutParameter, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/Router/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	Put(context.Context, *PutParameter) (*PutResponse, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) Put(ctx context.Context, req *PutParameter) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Put(ctx, req.(*PutParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Router_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}
