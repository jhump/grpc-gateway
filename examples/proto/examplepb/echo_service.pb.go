// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/examplepb/echo_service.proto

package examplepb

/*
Echo Service

Echo Service API consists of a single service which returns
a message.
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Embedded represents a message embedded in SimpleMessage.
type Embedded struct {
	// Types that are valid to be assigned to Mark:
	//	*Embedded_Progress
	//	*Embedded_Note
	Mark                 isEmbedded_Mark `protobuf_oneof:"mark"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Embedded) Reset()         { *m = Embedded{} }
func (m *Embedded) String() string { return proto.CompactTextString(m) }
func (*Embedded) ProtoMessage()    {}
func (*Embedded) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_service_c2772da14984c01c, []int{0}
}
func (m *Embedded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Embedded.Unmarshal(m, b)
}
func (m *Embedded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Embedded.Marshal(b, m, deterministic)
}
func (dst *Embedded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Embedded.Merge(dst, src)
}
func (m *Embedded) XXX_Size() int {
	return xxx_messageInfo_Embedded.Size(m)
}
func (m *Embedded) XXX_DiscardUnknown() {
	xxx_messageInfo_Embedded.DiscardUnknown(m)
}

var xxx_messageInfo_Embedded proto.InternalMessageInfo

type isEmbedded_Mark interface {
	isEmbedded_Mark()
}

type Embedded_Progress struct {
	Progress int64 `protobuf:"varint,1,opt,name=progress,oneof"`
}
type Embedded_Note struct {
	Note string `protobuf:"bytes,2,opt,name=note,oneof"`
}

func (*Embedded_Progress) isEmbedded_Mark() {}
func (*Embedded_Note) isEmbedded_Mark()     {}

func (m *Embedded) GetMark() isEmbedded_Mark {
	if m != nil {
		return m.Mark
	}
	return nil
}

func (m *Embedded) GetProgress() int64 {
	if x, ok := m.GetMark().(*Embedded_Progress); ok {
		return x.Progress
	}
	return 0
}

func (m *Embedded) GetNote() string {
	if x, ok := m.GetMark().(*Embedded_Note); ok {
		return x.Note
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Embedded) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Embedded_OneofMarshaler, _Embedded_OneofUnmarshaler, _Embedded_OneofSizer, []interface{}{
		(*Embedded_Progress)(nil),
		(*Embedded_Note)(nil),
	}
}

func _Embedded_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Embedded)
	// mark
	switch x := m.Mark.(type) {
	case *Embedded_Progress:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Progress))
	case *Embedded_Note:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Note)
	case nil:
	default:
		return fmt.Errorf("Embedded.Mark has unexpected type %T", x)
	}
	return nil
}

func _Embedded_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Embedded)
	switch tag {
	case 1: // mark.progress
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mark = &Embedded_Progress{int64(x)}
		return true, err
	case 2: // mark.note
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Mark = &Embedded_Note{x}
		return true, err
	default:
		return false, nil
	}
}

func _Embedded_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Embedded)
	// mark
	switch x := m.Mark.(type) {
	case *Embedded_Progress:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Progress))
	case *Embedded_Note:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Note)))
		n += len(x.Note)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// SimpleMessage represents a simple message sent to the Echo service.
type SimpleMessage struct {
	// Id represents the message identifier.
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	// Types that are valid to be assigned to Code:
	//	*SimpleMessage_LineNum
	//	*SimpleMessage_Lang
	Code   isSimpleMessage_Code `protobuf_oneof:"code"`
	Status *Embedded            `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	// Types that are valid to be assigned to Ext:
	//	*SimpleMessage_En
	//	*SimpleMessage_No
	Ext                  isSimpleMessage_Ext `protobuf_oneof:"ext"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_service_c2772da14984c01c, []int{1}
}
func (m *SimpleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMessage.Unmarshal(m, b)
}
func (m *SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMessage.Marshal(b, m, deterministic)
}
func (dst *SimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessage.Merge(dst, src)
}
func (m *SimpleMessage) XXX_Size() int {
	return xxx_messageInfo_SimpleMessage.Size(m)
}
func (m *SimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessage proto.InternalMessageInfo

type isSimpleMessage_Code interface {
	isSimpleMessage_Code()
}
type isSimpleMessage_Ext interface {
	isSimpleMessage_Ext()
}

type SimpleMessage_LineNum struct {
	LineNum int64 `protobuf:"varint,3,opt,name=line_num,json=lineNum,oneof"`
}
type SimpleMessage_Lang struct {
	Lang string `protobuf:"bytes,4,opt,name=lang,oneof"`
}
type SimpleMessage_En struct {
	En int64 `protobuf:"varint,6,opt,name=en,oneof"`
}
type SimpleMessage_No struct {
	No *Embedded `protobuf:"bytes,7,opt,name=no,oneof"`
}

func (*SimpleMessage_LineNum) isSimpleMessage_Code() {}
func (*SimpleMessage_Lang) isSimpleMessage_Code()    {}
func (*SimpleMessage_En) isSimpleMessage_Ext()       {}
func (*SimpleMessage_No) isSimpleMessage_Ext()       {}

func (m *SimpleMessage) GetCode() isSimpleMessage_Code {
	if m != nil {
		return m.Code
	}
	return nil
}
func (m *SimpleMessage) GetExt() isSimpleMessage_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (m *SimpleMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleMessage) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *SimpleMessage) GetLineNum() int64 {
	if x, ok := m.GetCode().(*SimpleMessage_LineNum); ok {
		return x.LineNum
	}
	return 0
}

func (m *SimpleMessage) GetLang() string {
	if x, ok := m.GetCode().(*SimpleMessage_Lang); ok {
		return x.Lang
	}
	return ""
}

func (m *SimpleMessage) GetStatus() *Embedded {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SimpleMessage) GetEn() int64 {
	if x, ok := m.GetExt().(*SimpleMessage_En); ok {
		return x.En
	}
	return 0
}

func (m *SimpleMessage) GetNo() *Embedded {
	if x, ok := m.GetExt().(*SimpleMessage_No); ok {
		return x.No
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SimpleMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SimpleMessage_OneofMarshaler, _SimpleMessage_OneofUnmarshaler, _SimpleMessage_OneofSizer, []interface{}{
		(*SimpleMessage_LineNum)(nil),
		(*SimpleMessage_Lang)(nil),
		(*SimpleMessage_En)(nil),
		(*SimpleMessage_No)(nil),
	}
}

func _SimpleMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SimpleMessage)
	// code
	switch x := m.Code.(type) {
	case *SimpleMessage_LineNum:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.LineNum))
	case *SimpleMessage_Lang:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Lang)
	case nil:
	default:
		return fmt.Errorf("SimpleMessage.Code has unexpected type %T", x)
	}
	// ext
	switch x := m.Ext.(type) {
	case *SimpleMessage_En:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.En))
	case *SimpleMessage_No:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.No); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SimpleMessage.Ext has unexpected type %T", x)
	}
	return nil
}

func _SimpleMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SimpleMessage)
	switch tag {
	case 3: // code.line_num
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Code = &SimpleMessage_LineNum{int64(x)}
		return true, err
	case 4: // code.lang
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Code = &SimpleMessage_Lang{x}
		return true, err
	case 6: // ext.en
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Ext = &SimpleMessage_En{int64(x)}
		return true, err
	case 7: // ext.no
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Embedded)
		err := b.DecodeMessage(msg)
		m.Ext = &SimpleMessage_No{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SimpleMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SimpleMessage)
	// code
	switch x := m.Code.(type) {
	case *SimpleMessage_LineNum:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.LineNum))
	case *SimpleMessage_Lang:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Lang)))
		n += len(x.Lang)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// ext
	switch x := m.Ext.(type) {
	case *SimpleMessage_En:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.En))
	case *SimpleMessage_No:
		s := proto.Size(x.No)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Embedded)(nil), "grpc.gateway.examples.examplepb.Embedded")
	proto.RegisterType((*SimpleMessage)(nil), "grpc.gateway.examples.examplepb.SimpleMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.EchoService/EchoBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.EchoService/EchoDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.EchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoBody(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.EchoService/EchoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoDelete(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _EchoService_EchoBody_Handler,
		},
		{
			MethodName: "EchoDelete",
			Handler:    _EchoService_EchoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/proto/examplepb/echo_service.proto",
}

func init() {
	proto.RegisterFile("examples/proto/examplepb/echo_service.proto", fileDescriptor_echo_service_c2772da14984c01c)
}

var fileDescriptor_echo_service_c2772da14984c01c = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xed, 0x8c, 0xdd, 0x34, 0x99, 0xe8, 0xfb, 0x54, 0x8d, 0x40, 0x98, 0xb4, 0xa8, 0x91, 0xc5,
	0x22, 0x14, 0xc9, 0xa3, 0x84, 0x5d, 0x61, 0x83, 0x45, 0x51, 0x37, 0xb0, 0x70, 0x77, 0xde, 0x44,
	0x13, 0xcf, 0x95, 0x6b, 0x61, 0xcf, 0x58, 0xf6, 0xa4, 0x34, 0xb2, 0xb2, 0x41, 0x62, 0xc9, 0x8a,
	0x3d, 0x8f, 0xd0, 0x1d, 0x4f, 0xc2, 0x2b, 0xb0, 0xe7, 0x15, 0xd0, 0x4c, 0xe2, 0x48, 0xd0, 0x88,
	0xaa, 0x8b, 0xee, 0x7c, 0x7f, 0xce, 0x3d, 0xe7, 0x9e, 0xb9, 0x26, 0xcf, 0xe1, 0x8a, 0x17, 0x65,
	0x0e, 0x35, 0x2b, 0x2b, 0xa5, 0x15, 0x5b, 0x87, 0xe5, 0x8c, 0x41, 0x72, 0xa1, 0xa6, 0x35, 0x54,
	0x97, 0x59, 0x02, 0x81, 0x2d, 0xd2, 0xa3, 0xb4, 0x2a, 0x93, 0x20, 0xe5, 0x1a, 0x3e, 0xf2, 0x45,
	0xd0, 0x22, 0x83, 0x0d, 0x66, 0x70, 0x98, 0x2a, 0x95, 0xe6, 0xc0, 0x78, 0x99, 0x31, 0x2e, 0xa5,
	0xd2, 0x5c, 0x67, 0x4a, 0xd6, 0x2b, 0xb8, 0xff, 0x96, 0x74, 0x4f, 0x8b, 0x19, 0x08, 0x01, 0x82,
	0x1e, 0x92, 0x6e, 0x59, 0xa9, 0xb4, 0x82, 0xba, 0xf6, 0xd0, 0x10, 0x8d, 0x9c, 0xb3, 0x9d, 0x68,
	0x93, 0xa1, 0x0f, 0x88, 0x2b, 0x95, 0x06, 0x0f, 0x0f, 0xd1, 0xa8, 0x77, 0xb6, 0x13, 0xd9, 0x28,
	0xec, 0x10, 0xb7, 0xe0, 0xd5, 0x07, 0xff, 0x33, 0x26, 0xff, 0x9d, 0x67, 0x86, 0xf2, 0x1d, 0xd4,
	0x35, 0x4f, 0x81, 0xfe, 0x4f, 0x70, 0x26, 0xec, 0x9c, 0x5e, 0x84, 0x33, 0x41, 0xf7, 0x89, 0x23,
	0xe7, 0x85, 0x85, 0x3b, 0x91, 0xf9, 0xa4, 0x07, 0xa4, 0x9b, 0x67, 0x12, 0xa6, 0x26, 0xed, 0xac,
	0xf9, 0xf6, 0x4c, 0xe6, 0xfd, 0xbc, 0x30, 0x74, 0x39, 0x97, 0xa9, 0xe7, 0xb6, 0x74, 0x26, 0xa2,
	0xaf, 0x49, 0xa7, 0xd6, 0x5c, 0xcf, 0x6b, 0x6f, 0x77, 0x88, 0x46, 0xfd, 0xc9, 0xb3, 0xe0, 0x96,
	0xf5, 0x83, 0x76, 0xbb, 0x68, 0x0d, 0xa4, 0xfb, 0x04, 0x83, 0xf4, 0x3a, 0x96, 0x0f, 0x45, 0x18,
	0x24, 0x7d, 0x49, 0xb0, 0x54, 0xde, 0xde, 0x1d, 0x07, 0x1a, 0xb0, 0x54, 0xc6, 0x80, 0x44, 0x09,
	0x08, 0x77, 0x89, 0x03, 0x57, 0x7a, 0xf2, 0xcb, 0x25, 0xfd, 0xd3, 0xe4, 0x42, 0x9d, 0xaf, 0x1e,
	0x89, 0x7e, 0xc3, 0xc4, 0x35, 0x31, 0x0d, 0x6e, 0x1d, 0xfc, 0x87, 0x7d, 0x83, 0x3b, 0xf6, 0xfb,
	0xdf, 0xd1, 0xa7, 0x1f, 0x3f, 0xbf, 0xe2, 0x6b, 0xe4, 0x3f, 0x64, 0x97, 0xe3, 0xf6, 0x66, 0xec,
	0xc5, 0xb0, 0x26, 0x13, 0xcb, 0xf8, 0x09, 0x3d, 0xd8, 0x5a, 0x60, 0x8d, 0x9c, 0x17, 0xcb, 0xf8,
	0x29, 0xf5, 0xff, 0x51, 0x66, 0x8d, 0xb1, 0x7f, 0x19, 0x8f, 0x29, 0xfb, 0xbb, 0x6b, 0xbc, 0x6e,
	0x6b, 0x9f, 0x72, 0xc9, 0x9a, 0x95, 0xd3, 0x81, 0x39, 0x90, 0xad, 0xbc, 0x13, 0xd6, 0x48, 0xb5,
	0x2a, 0xd3, 0x6b, 0x44, 0xba, 0xc6, 0xa0, 0x50, 0x89, 0xc5, 0xbd, 0x9b, 0x14, 0x5a, 0x8f, 0x5e,
	0xdd, 0xb4, 0x68, 0x3a, 0x53, 0x62, 0x71, 0x82, 0x8e, 0xe3, 0xe1, 0x60, 0xb0, 0xb5, 0x66, 0x97,
	0x3c, 0xc1, 0x52, 0xd1, 0x2f, 0x88, 0x10, 0x23, 0xf8, 0x0d, 0xe4, 0xa0, 0xe1, 0xde, 0x25, 0x1f,
	0x59, 0xc9, 0x8f, 0x8f, 0x1f, 0xdd, 0x90, 0x25, 0xac, 0x80, 0xb0, 0x1f, 0xf7, 0x36, 0xd8, 0x59,
	0xc7, 0xfe, 0xd5, 0x2f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x6a, 0xaf, 0xec, 0x43, 0x04,
	0x00, 0x00,
}
