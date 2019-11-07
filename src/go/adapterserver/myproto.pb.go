// Code generated by protoc-gen-adapterserver. DO NOT EDIT.
// source: myproto.proto

package main

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

type Template struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Template) Reset()         { *m = Template{} }
func (m *Template) String() string { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()    {}
func (*Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{0}
}

func (m *Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Template.Unmarshal(m, b)
}
func (m *Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Template.Marshal(b, m, deterministic)
}
func (m *Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Template.Merge(m, src)
}
func (m *Template) XXX_Size() int {
	return xxx_messageInfo_Template.Size(m)
}
func (m *Template) XXX_DiscardUnknown() {
	xxx_messageInfo_Template.DiscardUnknown(m)
}

var xxx_messageInfo_Template proto.InternalMessageInfo

func (m *Template) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Template) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type OutputTemplate struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The number of uses for which this result can be considered valid.
	ValidUseCount        int64    `protobuf:"varint,4,opt,name=valid_use_count,json=validUseCount,proto3" json:"valid_use_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputTemplate) Reset()         { *m = OutputTemplate{} }
func (m *OutputTemplate) String() string { return proto.CompactTextString(m) }
func (*OutputTemplate) ProtoMessage()    {}
func (*OutputTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{1}
}

func (m *OutputTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputTemplate.Unmarshal(m, b)
}
func (m *OutputTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputTemplate.Marshal(b, m, deterministic)
}
func (m *OutputTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputTemplate.Merge(m, src)
}
func (m *OutputTemplate) XXX_Size() int {
	return xxx_messageInfo_OutputTemplate.Size(m)
}
func (m *OutputTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_OutputTemplate proto.InternalMessageInfo

func (m *OutputTemplate) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *OutputTemplate) GetValidUseCount() int64 {
	if m != nil {
		return m.ValidUseCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Template)(nil), "myproto.Template")
	proto.RegisterType((*OutputTemplate)(nil), "myproto.OutputTemplate")
}

func init() { proto.RegisterFile("myproto.proto", fileDescriptor_04877df457807402) }

var fileDescriptor_04877df457807402 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xad, 0x2c, 0x28,
	0xca, 0x2f, 0xc9, 0xd7, 0x03, 0x93, 0x42, 0xec, 0x50, 0xae, 0x94, 0x40, 0x6a, 0x45, 0x49, 0x6a,
	0x5e, 0x71, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x4a, 0xc9, 0x80, 0x8b, 0x23, 0x24, 0x35, 0xb7, 0x20,
	0x27, 0xb1, 0x24, 0x55, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc4, 0x14, 0x12, 0xe2, 0x62, 0x29, 0x48, 0x2c, 0xc9, 0x90, 0x60, 0x02, 0x0b, 0x81,
	0xd9, 0x4a, 0x3e, 0x5c, 0x7c, 0xfe, 0xa5, 0x25, 0x05, 0xa5, 0x25, 0x70, 0x7d, 0x42, 0x5c, 0x2c,
	0xa5, 0xc5, 0xa9, 0x45, 0x50, 0x8d, 0x60, 0xb6, 0x90, 0x1a, 0x17, 0x7f, 0x59, 0x62, 0x4e, 0x66,
	0x4a, 0x7c, 0x69, 0x71, 0x6a, 0x7c, 0x72, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x8b, 0x02, 0xa3, 0x06,
	0x73, 0x10, 0x2f, 0x58, 0x38, 0xb4, 0x38, 0xd5, 0x19, 0x24, 0x68, 0x14, 0xcc, 0x25, 0xe2, 0x91,
	0x98, 0x97, 0x92, 0x93, 0xea, 0x0b, 0x71, 0x62, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90,
	0x35, 0x17, 0x2f, 0x8a, 0xb8, 0x90, 0xa0, 0x1e, 0xcc, 0x4f, 0x30, 0x7b, 0xa5, 0xc4, 0xe1, 0x42,
	0xa8, 0x0e, 0x72, 0x92, 0xff, 0x71, 0xe9, 0xc9, 0x64, 0x26, 0x16, 0x2e, 0xf1, 0xe4, 0xfc, 0x5c,
	0xbd, 0xb4, 0xfc, 0xa2, 0xf4, 0xd4, 0xa2, 0xfc, 0xe4, 0x6c, 0xbd, 0xf4, 0xd4, 0xbc, 0xd4, 0xa2,
	0xc4, 0x92, 0xd4, 0x94, 0x24, 0x36, 0xb0, 0x36, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3,
	0xc7, 0xa1, 0xb3, 0x28, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HandleMyprotoServiceClient is the client API for HandleMyprotoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandleMyprotoServiceClient interface {
	// HandleMyproto is called by Mixer at request-time to deliver 'authorization' instances to the backend.
	HandleMyproto(ctx context.Context, in *Template, opts ...grpc.CallOption) (*OutputTemplate, error)
}

type handleMyprotoServiceClient struct {
	cc *grpc.ClientConn
}

func NewHandleMyprotoServiceClient(cc *grpc.ClientConn) HandleMyprotoServiceClient {
	return &handleMyprotoServiceClient{cc}
}

func (c *handleMyprotoServiceClient) HandleMyproto(ctx context.Context, in *Template, opts ...grpc.CallOption) (*OutputTemplate, error) {
	out := new(OutputTemplate)
	err := c.cc.Invoke(ctx, "/myproto.HandleMyprotoService/HandleMyproto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandleMyprotoServiceServer is the adapterserver API for HandleMyprotoService service.
type HandleMyprotoServiceServer interface {
	// HandleMyproto is called by Mixer at request-time to deliver 'authorization' instances to the backend.
	HandleMyproto(context.Context, *Template) (*OutputTemplate, error)
}

// UnimplementedHandleMyprotoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHandleMyprotoServiceServer struct {
}

func (*UnimplementedHandleMyprotoServiceServer) HandleMyproto(ctx context.Context, req *Template) (*OutputTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleMyproto not implemented")
}

func RegisterHandleMyprotoServiceServer(s *grpc.Server, srv HandleMyprotoServiceServer) {
	s.RegisterService(&_HandleMyprotoService_serviceDesc, srv)
}

func _HandleMyprotoService_HandleMyproto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Template)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleMyprotoServiceServer).HandleMyproto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproto.HandleMyprotoService/HandleMyproto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleMyprotoServiceServer).HandleMyproto(ctx, req.(*Template))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleMyprotoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myproto.HandleMyprotoService",
	HandlerType: (*HandleMyprotoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleMyproto",
			Handler:    _HandleMyprotoService_HandleMyproto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myproto.proto",
}