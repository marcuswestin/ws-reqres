// Code generated by protoc-gen-go.
// source: birect.proto
// DO NOT EDIT!

/*
Package wire is a generated protocol buffer package.

It is generated from these files:
	birect.proto

It has these top-level messages:
	Wrapper
	Message
	Request
	Response
*/
package wire

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DataType int32

const (
	DataType__     DataType = 0
	DataType_JSON  DataType = 1
	DataType_Proto DataType = 2
)

var DataType_name = map[int32]string{
	0: "_",
	1: "JSON",
	2: "Proto",
}
var DataType_value = map[string]int32{
	"_":     0,
	"JSON":  1,
	"Proto": 2,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Wrapper struct {
	// Types that are valid to be assigned to Content:
	//	*Wrapper_Message
	//	*Wrapper_Request
	//	*Wrapper_Response
	Content isWrapper_Content `protobuf_oneof:"content"`
}

func (m *Wrapper) Reset()                    { *m = Wrapper{} }
func (m *Wrapper) String() string            { return proto.CompactTextString(m) }
func (*Wrapper) ProtoMessage()               {}
func (*Wrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isWrapper_Content interface {
	isWrapper_Content()
}

type Wrapper_Message struct {
	Message *Message `protobuf:"bytes,1,opt,name=message,oneof"`
}
type Wrapper_Request struct {
	Request *Request `protobuf:"bytes,2,opt,name=request,oneof"`
}
type Wrapper_Response struct {
	Response *Response `protobuf:"bytes,3,opt,name=response,oneof"`
}

func (*Wrapper_Message) isWrapper_Content()  {}
func (*Wrapper_Request) isWrapper_Content()  {}
func (*Wrapper_Response) isWrapper_Content() {}

func (m *Wrapper) GetContent() isWrapper_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Wrapper) GetMessage() *Message {
	if x, ok := m.GetContent().(*Wrapper_Message); ok {
		return x.Message
	}
	return nil
}

func (m *Wrapper) GetRequest() *Request {
	if x, ok := m.GetContent().(*Wrapper_Request); ok {
		return x.Request
	}
	return nil
}

func (m *Wrapper) GetResponse() *Response {
	if x, ok := m.GetContent().(*Wrapper_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Wrapper) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Wrapper_OneofMarshaler, _Wrapper_OneofUnmarshaler, _Wrapper_OneofSizer, []interface{}{
		(*Wrapper_Message)(nil),
		(*Wrapper_Request)(nil),
		(*Wrapper_Response)(nil),
	}
}

func _Wrapper_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Wrapper)
	// content
	switch x := m.Content.(type) {
	case *Wrapper_Message:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Message); err != nil {
			return err
		}
	case *Wrapper_Request:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Request); err != nil {
			return err
		}
	case *Wrapper_Response:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Response); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Wrapper.Content has unexpected type %T", x)
	}
	return nil
}

func _Wrapper_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Wrapper)
	switch tag {
	case 1: // content.message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message)
		err := b.DecodeMessage(msg)
		m.Content = &Wrapper_Message{msg}
		return true, err
	case 2: // content.request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request)
		err := b.DecodeMessage(msg)
		m.Content = &Wrapper_Request{msg}
		return true, err
	case 3: // content.response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Response)
		err := b.DecodeMessage(msg)
		m.Content = &Wrapper_Response{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Wrapper_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Wrapper)
	// content
	switch x := m.Content.(type) {
	case *Wrapper_Message:
		s := proto.Size(x.Message)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Wrapper_Request:
		s := proto.Size(x.Request)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Wrapper_Response:
		s := proto.Size(x.Response)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Message struct {
	Type DataType `protobuf:"varint,1,opt,name=type,enum=wire.DataType" json:"type,omitempty"`
	Name string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Data []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Request struct {
	Type  DataType `protobuf:"varint,1,opt,name=type,enum=wire.DataType" json:"type,omitempty"`
	ReqId uint32   `protobuf:"varint,2,opt,name=req_id" json:"req_id,omitempty"`
	Name  string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Data  []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Type    DataType `protobuf:"varint,1,opt,name=type,enum=wire.DataType" json:"type,omitempty"`
	ReqId   uint32   `protobuf:"varint,2,opt,name=req_id" json:"req_id,omitempty"`
	IsError bool     `protobuf:"varint,3,opt,name=isError" json:"isError,omitempty"`
	// 3 left out
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Wrapper)(nil), "wire.Wrapper")
	proto.RegisterType((*Message)(nil), "wire.Message")
	proto.RegisterType((*Request)(nil), "wire.Request")
	proto.RegisterType((*Response)(nil), "wire.Response")
	proto.RegisterEnum("wire.DataType", DataType_name, DataType_value)
}

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0xbb, 0xed, 0xb6, 0xc9, 0x8e, 0x6d, 0x59, 0x02, 0xc2, 0x1e, 0x65, 0x0f, 0xa2, 0x22,
	0x7b, 0xd0, 0x37, 0x10, 0x05, 0x15, 0xfc, 0x43, 0x54, 0x3c, 0x96, 0xb4, 0x3b, 0xc8, 0x82, 0xdd,
	0xc4, 0x24, 0x22, 0xbe, 0x89, 0x8f, 0x6b, 0x92, 0xdd, 0x94, 0x1e, 0x7a, 0x10, 0xbc, 0x65, 0x7e,
	0xf3, 0x65, 0xbe, 0x09, 0x04, 0xa6, 0xcb, 0x46, 0xe3, 0xca, 0x56, 0x4a, 0x4b, 0x2b, 0x59, 0xfa,
	0xe5, 0xaa, 0xf2, 0x27, 0x01, 0xf2, 0xaa, 0x85, 0x52, 0xa8, 0xd9, 0x31, 0x90, 0x35, 0x1a, 0x23,
	0xde, 0xb0, 0x48, 0x0e, 0x92, 0xa3, 0xbd, 0xb3, 0x59, 0xe5, 0x99, 0xea, 0xae, 0x0b, 0xaf, 0x07,
	0x3c, 0xf6, 0x3d, 0xaa, 0xf1, 0xe3, 0x13, 0x8d, 0x2d, 0x86, 0xdb, 0x28, 0xef, 0x42, 0x8f, 0xf6,
	0x7d, 0x76, 0x0a, 0x54, 0xa3, 0x51, 0xb2, 0x35, 0x58, 0x8c, 0x02, 0x3b, 0x8f, 0x6c, 0x97, 0x3a,
	0x78, 0x43, 0x5c, 0x64, 0x40, 0x56, 0xb2, 0xb5, 0xd8, 0xda, 0xf2, 0x05, 0x48, 0x6f, 0x66, 0x25,
	0xa4, 0xf6, 0x5b, 0x75, 0x6b, 0xcd, 0xe3, 0xfd, 0x4b, 0x61, 0xc5, 0xb3, 0x4b, 0x79, 0xe8, 0x31,
	0x06, 0x69, 0x2b, 0xd6, 0x18, 0xf6, 0xc9, 0x78, 0x38, 0xfb, 0xac, 0x76, 0x54, 0xf0, 0x4e, 0x79,
	0x38, 0x97, 0xef, 0x40, 0xfa, 0x2d, 0xff, 0x34, 0x76, 0x1f, 0x26, 0xee, 0x25, 0x8b, 0xa6, 0x0e,
	0x83, 0x67, 0x7c, 0xec, 0xaa, 0x9b, 0x7a, 0x63, 0x1b, 0xed, 0xb0, 0xa5, 0x5b, 0x36, 0x03, 0x34,
	0xbe, 0xf3, 0x3f, 0xba, 0x02, 0x48, 0x63, 0xae, 0xb4, 0x96, 0x3a, 0x18, 0x29, 0x8f, 0xe5, 0x2e,
	0xe9, 0xc9, 0x21, 0xd0, 0x38, 0x96, 0x8d, 0x21, 0x59, 0xe4, 0x03, 0x46, 0x21, 0xbd, 0x7d, 0x7a,
	0xb8, 0xcf, 0x13, 0x96, 0xc1, 0xf8, 0xd1, 0x7f, 0x80, 0x7c, 0xb8, 0x9c, 0x84, 0x9f, 0x70, 0xfe,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x9a, 0x03, 0xe0, 0x19, 0x02, 0x00, 0x00,
}
