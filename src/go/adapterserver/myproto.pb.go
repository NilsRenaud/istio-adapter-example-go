// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: myproto/myproto.proto

package main

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/api/mixer/adapter/model/v1beta1"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Template struct {
	Key  []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Path []byte `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (m *Template) Reset()      { *m = Template{} }
func (*Template) ProtoMessage() {}
func (*Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_501d76cd23c80327, []int{0}
}
func (m *Template) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Template.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Template.Merge(m, src)
}
func (m *Template) XXX_Size() int {
	return m.Size()
}
func (m *Template) XXX_DiscardUnknown() {
	xxx_messageInfo_Template.DiscardUnknown(m)
}

var xxx_messageInfo_Template proto.InternalMessageInfo

func (m *Template) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Template) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type OutputTemplate struct {
	User []byte `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The number of uses for which this result can be considered valid.
	ValidUseCount int64 `protobuf:"varint,4,opt,name=valid_use_count,json=validUseCount,proto3" json:"valid_use_count,omitempty"`
}

func (m *OutputTemplate) Reset()      { *m = OutputTemplate{} }
func (*OutputTemplate) ProtoMessage() {}
func (*OutputTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_501d76cd23c80327, []int{1}
}
func (m *OutputTemplate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutputTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutputTemplate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutputTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputTemplate.Merge(m, src)
}
func (m *OutputTemplate) XXX_Size() int {
	return m.Size()
}
func (m *OutputTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_OutputTemplate proto.InternalMessageInfo

func (m *OutputTemplate) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
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

func init() { proto.RegisterFile("myproto/myproto.proto", fileDescriptor_501d76cd23c80327) }

var fileDescriptor_501d76cd23c80327 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xad, 0x2c, 0x28,
	0xca, 0x2f, 0xc9, 0xd7, 0x87, 0xd2, 0x7a, 0x60, 0x52, 0x88, 0x1d, 0xca, 0x95, 0x12, 0x48, 0xad,
	0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0x86, 0x48, 0x29, 0x19, 0x70, 0x71, 0x84, 0xa4,
	0xe6, 0x16, 0xe4, 0x24, 0x96, 0xa4, 0x0a, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0xf0, 0x04, 0x81, 0x98, 0x42, 0x42, 0x5c, 0x2c, 0x05, 0x89, 0x25, 0x19, 0x12, 0x4c,
	0x60, 0x21, 0x30, 0x5b, 0xc9, 0x87, 0x8b, 0xcf, 0xbf, 0xb4, 0xa4, 0xa0, 0xb4, 0x04, 0xae, 0x4f,
	0x88, 0x8b, 0xa5, 0xb4, 0x38, 0xb5, 0x08, 0xaa, 0x11, 0xcc, 0x16, 0x52, 0xe3, 0xe2, 0x2f, 0x4b,
	0xcc, 0xc9, 0x4c, 0x89, 0x2f, 0x2d, 0x4e, 0x8d, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x51,
	0x60, 0xd4, 0x60, 0x0e, 0xe2, 0x05, 0x0b, 0x87, 0x16, 0xa7, 0x3a, 0x83, 0x04, 0x8d, 0x82, 0xb9,
	0x44, 0x3c, 0x12, 0xf3, 0x52, 0x72, 0x52, 0x7d, 0x21, 0x4e, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x15, 0xb2, 0xe6, 0xe2, 0x45, 0x11, 0x17, 0x12, 0xd4, 0x83, 0xf9, 0x09, 0x66, 0xaf, 0x94,
	0x38, 0x5c, 0x08, 0xd5, 0x41, 0x4e, 0xc1, 0x17, 0x1e, 0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7, 0xf0,
	0xe1, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x17, 0x8f, 0xe4, 0x18, 0x3e, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x7e, 0x5c,
	0x7a, 0x32, 0x99, 0x89, 0x85, 0x4b, 0x3c, 0x39, 0x3f, 0x57, 0x2f, 0x2d, 0xbf, 0x28, 0x3d, 0xb5,
	0x28, 0x3f, 0x39, 0x5b, 0x2f, 0x3d, 0x35, 0x2f, 0xb5, 0x28, 0xb1, 0x24, 0x35, 0x25, 0x89, 0x0d,
	0x6c, 0x95, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x64, 0x15, 0x5b, 0xab, 0x64, 0x01, 0x00, 0x00,
}

func (this *Template) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Template)
	if !ok {
		that2, ok := that.(Template)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Key, that1.Key) {
		return false
	}
	if !bytes.Equal(this.Path, that1.Path) {
		return false
	}
	return true
}
func (this *OutputTemplate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OutputTemplate)
	if !ok {
		that2, ok := that.(OutputTemplate)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.User, that1.User) {
		return false
	}
	if this.ValidUseCount != that1.ValidUseCount {
		return false
	}
	return true
}
func (this *Template) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&myproto.Template{")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Path: "+fmt.Sprintf("%#v", this.Path)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *OutputTemplate) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&myproto.OutputTemplate{")
	s = append(s, "User: "+fmt.Sprintf("%#v", this.User)+",\n")
	s = append(s, "ValidUseCount: "+fmt.Sprintf("%#v", this.ValidUseCount)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMyproto(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
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

// HandleMyprotoServiceServer is the server API for HandleMyprotoService service.
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
	Metadata: "myproto/myproto.proto",
}

func (m *Template) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Template) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Template) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintMyproto(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintMyproto(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OutputTemplate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutputTemplate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutputTemplate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValidUseCount != 0 {
		i = encodeVarintMyproto(dAtA, i, uint64(m.ValidUseCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintMyproto(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMyproto(dAtA []byte, offset int, v uint64) int {
	offset -= sovMyproto(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Template) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMyproto(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovMyproto(uint64(l))
	}
	return n
}

func (m *OutputTemplate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovMyproto(uint64(l))
	}
	if m.ValidUseCount != 0 {
		n += 1 + sovMyproto(uint64(m.ValidUseCount))
	}
	return n
}

func sovMyproto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMyproto(x uint64) (n int) {
	return sovMyproto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Template) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Template{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`}`,
	}, "")
	return s
}
func (this *OutputTemplate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&OutputTemplate{`,
		`User:` + fmt.Sprintf("%v", this.User) + `,`,
		`ValidUseCount:` + fmt.Sprintf("%v", this.ValidUseCount) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMyproto(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Template) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMyproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Template: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Template: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMyproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMyproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMyproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMyproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMyproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMyproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = append(m.Path[:0], dAtA[iNdEx:postIndex]...)
			if m.Path == nil {
				m.Path = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMyproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMyproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMyproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OutputTemplate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMyproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutputTemplate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutputTemplate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMyproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMyproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMyproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = append(m.User[:0], dAtA[iNdEx:postIndex]...)
			if m.User == nil {
				m.User = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUseCount", wireType)
			}
			m.ValidUseCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMyproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidUseCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMyproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMyproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMyproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMyproto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMyproto
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMyproto
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMyproto
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMyproto
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMyproto
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMyproto
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMyproto        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMyproto          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMyproto = fmt.Errorf("proto: unexpected end of group")
)
