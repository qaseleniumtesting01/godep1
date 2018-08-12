// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encryption.proto

package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// X3DH prekey bundle
type Bundle struct {
	// Identity key
	Identity []byte `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// Signed prekey
	SignedPreKey []byte `protobuf:"bytes,2,opt,name=signed_pre_key,json=signedPreKey,proto3" json:"signed_pre_key,omitempty"`
	// Prekey signature
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bundle) Reset()         { *m = Bundle{} }
func (m *Bundle) String() string { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()    {}
func (*Bundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_encryption_a35d9f7c683dd9c2, []int{0}
}
func (m *Bundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bundle.Unmarshal(m, b)
}
func (m *Bundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bundle.Marshal(b, m, deterministic)
}
func (dst *Bundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundle.Merge(dst, src)
}
func (m *Bundle) XXX_Size() int {
	return xxx_messageInfo_Bundle.Size(m)
}
func (m *Bundle) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundle.DiscardUnknown(m)
}

var xxx_messageInfo_Bundle proto.InternalMessageInfo

func (m *Bundle) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Bundle) GetSignedPreKey() []byte {
	if m != nil {
		return m.SignedPreKey
	}
	return nil
}

func (m *Bundle) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type BundleContainer struct {
	// X3DH prekey bundle
	Bundle *Bundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
	// Private signed prekey
	PrivateSignedPreKey []byte `protobuf:"bytes,2,opt,name=private_signed_pre_key,json=privateSignedPreKey,proto3" json:"private_signed_pre_key,omitempty"`
	// Local time creation
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BundleContainer) Reset()         { *m = BundleContainer{} }
func (m *BundleContainer) String() string { return proto.CompactTextString(m) }
func (*BundleContainer) ProtoMessage()    {}
func (*BundleContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_encryption_a35d9f7c683dd9c2, []int{1}
}
func (m *BundleContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleContainer.Unmarshal(m, b)
}
func (m *BundleContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleContainer.Marshal(b, m, deterministic)
}
func (dst *BundleContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleContainer.Merge(dst, src)
}
func (m *BundleContainer) XXX_Size() int {
	return xxx_messageInfo_BundleContainer.Size(m)
}
func (m *BundleContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleContainer.DiscardUnknown(m)
}

var xxx_messageInfo_BundleContainer proto.InternalMessageInfo

func (m *BundleContainer) GetBundle() *Bundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

func (m *BundleContainer) GetPrivateSignedPreKey() []byte {
	if m != nil {
		return m.PrivateSignedPreKey
	}
	return nil
}

func (m *BundleContainer) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type DRHeader struct {
	// Current ratchet public key
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Number of the message in the sending chain
	N uint32 `protobuf:"varint,2,opt,name=n,proto3" json:"n,omitempty"`
	// Length of the previous sending chain
	Pn uint32 `protobuf:"varint,3,opt,name=pn,proto3" json:"pn,omitempty"`
	// Bundle ID
	Id                   []byte   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DRHeader) Reset()         { *m = DRHeader{} }
func (m *DRHeader) String() string { return proto.CompactTextString(m) }
func (*DRHeader) ProtoMessage()    {}
func (*DRHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_encryption_a35d9f7c683dd9c2, []int{2}
}
func (m *DRHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DRHeader.Unmarshal(m, b)
}
func (m *DRHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DRHeader.Marshal(b, m, deterministic)
}
func (dst *DRHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DRHeader.Merge(dst, src)
}
func (m *DRHeader) XXX_Size() int {
	return xxx_messageInfo_DRHeader.Size(m)
}
func (m *DRHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DRHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DRHeader proto.InternalMessageInfo

func (m *DRHeader) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *DRHeader) GetN() uint32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *DRHeader) GetPn() uint32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *DRHeader) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type DHHeader struct {
	// Compressed ephemeral public key
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DHHeader) Reset()         { *m = DHHeader{} }
func (m *DHHeader) String() string { return proto.CompactTextString(m) }
func (*DHHeader) ProtoMessage()    {}
func (*DHHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_encryption_a35d9f7c683dd9c2, []int{3}
}
func (m *DHHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHHeader.Unmarshal(m, b)
}
func (m *DHHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHHeader.Marshal(b, m, deterministic)
}
func (dst *DHHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHHeader.Merge(dst, src)
}
func (m *DHHeader) XXX_Size() int {
	return xxx_messageInfo_DHHeader.Size(m)
}
func (m *DHHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DHHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DHHeader proto.InternalMessageInfo

func (m *DHHeader) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type X3DHHeader struct {
	// Ephemeral key used
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Used bundle's signed prekey
	Id                   []byte   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *X3DHHeader) Reset()         { *m = X3DHHeader{} }
func (m *X3DHHeader) String() string { return proto.CompactTextString(m) }
func (*X3DHHeader) ProtoMessage()    {}
func (*X3DHHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_encryption_a35d9f7c683dd9c2, []int{4}
}
func (m *X3DHHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X3DHHeader.Unmarshal(m, b)
}
func (m *X3DHHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X3DHHeader.Marshal(b, m, deterministic)
}
func (dst *X3DHHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X3DHHeader.Merge(dst, src)
}
func (m *X3DHHeader) XXX_Size() int {
	return xxx_messageInfo_X3DHHeader.Size(m)
}
func (m *X3DHHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_X3DHHeader.DiscardUnknown(m)
}

var xxx_messageInfo_X3DHHeader proto.InternalMessageInfo

func (m *X3DHHeader) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *X3DHHeader) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

// Direct message value
type DirectMessageProtocol struct {
	X3DHHeader *X3DHHeader `protobuf:"bytes,1,opt,name=X3DH_header,json=X3DHHeader,proto3" json:"X3DH_header,omitempty"`
	DRHeader   *DRHeader   `protobuf:"bytes,2,opt,name=DR_header,json=DRHeader,proto3" json:"DR_header,omitempty"`
	DHHeader   *DHHeader   `protobuf:"bytes,101,opt,name=DH_header,json=DHHeader,proto3" json:"DH_header,omitempty"`
	// Encrypted payload
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectMessageProtocol) Reset()         { *m = DirectMessageProtocol{} }
func (m *DirectMessageProtocol) String() string { return proto.CompactTextString(m) }
func (*DirectMessageProtocol) ProtoMessage()    {}
func (*DirectMessageProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_encryption_a35d9f7c683dd9c2, []int{5}
}
func (m *DirectMessageProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectMessageProtocol.Unmarshal(m, b)
}
func (m *DirectMessageProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectMessageProtocol.Marshal(b, m, deterministic)
}
func (dst *DirectMessageProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectMessageProtocol.Merge(dst, src)
}
func (m *DirectMessageProtocol) XXX_Size() int {
	return xxx_messageInfo_DirectMessageProtocol.Size(m)
}
func (m *DirectMessageProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectMessageProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_DirectMessageProtocol proto.InternalMessageInfo

func (m *DirectMessageProtocol) GetX3DHHeader() *X3DHHeader {
	if m != nil {
		return m.X3DHHeader
	}
	return nil
}

func (m *DirectMessageProtocol) GetDRHeader() *DRHeader {
	if m != nil {
		return m.DRHeader
	}
	return nil
}

func (m *DirectMessageProtocol) GetDHHeader() *DHHeader {
	if m != nil {
		return m.DHHeader
	}
	return nil
}

func (m *DirectMessageProtocol) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Top-level protocol message
type ProtocolMessage struct {
	// An optional bundle is exchanged with each message
	Bundle *Bundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
	// Types that are valid to be assigned to MessageType:
	//	*ProtocolMessage_DirectMessage
	//	*ProtocolMessage_PublicMessage
	MessageType          isProtocolMessage_MessageType `protobuf_oneof:"message_type"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ProtocolMessage) Reset()         { *m = ProtocolMessage{} }
func (m *ProtocolMessage) String() string { return proto.CompactTextString(m) }
func (*ProtocolMessage) ProtoMessage()    {}
func (*ProtocolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_encryption_a35d9f7c683dd9c2, []int{6}
}
func (m *ProtocolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolMessage.Unmarshal(m, b)
}
func (m *ProtocolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolMessage.Marshal(b, m, deterministic)
}
func (dst *ProtocolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolMessage.Merge(dst, src)
}
func (m *ProtocolMessage) XXX_Size() int {
	return xxx_messageInfo_ProtocolMessage.Size(m)
}
func (m *ProtocolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolMessage proto.InternalMessageInfo

type isProtocolMessage_MessageType interface {
	isProtocolMessage_MessageType()
}

type ProtocolMessage_DirectMessage struct {
	DirectMessage *DirectMessageProtocol `protobuf:"bytes,101,opt,name=direct_message,json=directMessage,proto3,oneof"`
}
type ProtocolMessage_PublicMessage struct {
	PublicMessage []byte `protobuf:"bytes,102,opt,name=public_message,json=publicMessage,proto3,oneof"`
}

func (*ProtocolMessage_DirectMessage) isProtocolMessage_MessageType() {}
func (*ProtocolMessage_PublicMessage) isProtocolMessage_MessageType() {}

func (m *ProtocolMessage) GetMessageType() isProtocolMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *ProtocolMessage) GetBundle() *Bundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

func (m *ProtocolMessage) GetDirectMessage() *DirectMessageProtocol {
	if x, ok := m.GetMessageType().(*ProtocolMessage_DirectMessage); ok {
		return x.DirectMessage
	}
	return nil
}

func (m *ProtocolMessage) GetPublicMessage() []byte {
	if x, ok := m.GetMessageType().(*ProtocolMessage_PublicMessage); ok {
		return x.PublicMessage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProtocolMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProtocolMessage_OneofMarshaler, _ProtocolMessage_OneofUnmarshaler, _ProtocolMessage_OneofSizer, []interface{}{
		(*ProtocolMessage_DirectMessage)(nil),
		(*ProtocolMessage_PublicMessage)(nil),
	}
}

func _ProtocolMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProtocolMessage)
	// message_type
	switch x := m.MessageType.(type) {
	case *ProtocolMessage_DirectMessage:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DirectMessage); err != nil {
			return err
		}
	case *ProtocolMessage_PublicMessage:
		b.EncodeVarint(102<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.PublicMessage)
	case nil:
	default:
		return fmt.Errorf("ProtocolMessage.MessageType has unexpected type %T", x)
	}
	return nil
}

func _ProtocolMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProtocolMessage)
	switch tag {
	case 101: // message_type.direct_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DirectMessageProtocol)
		err := b.DecodeMessage(msg)
		m.MessageType = &ProtocolMessage_DirectMessage{msg}
		return true, err
	case 102: // message_type.public_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.MessageType = &ProtocolMessage_PublicMessage{x}
		return true, err
	default:
		return false, nil
	}
}

func _ProtocolMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProtocolMessage)
	// message_type
	switch x := m.MessageType.(type) {
	case *ProtocolMessage_DirectMessage:
		s := proto.Size(x.DirectMessage)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProtocolMessage_PublicMessage:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.PublicMessage)))
		n += len(x.PublicMessage)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Bundle)(nil), "chat.Bundle")
	proto.RegisterType((*BundleContainer)(nil), "chat.BundleContainer")
	proto.RegisterType((*DRHeader)(nil), "chat.DRHeader")
	proto.RegisterType((*DHHeader)(nil), "chat.DHHeader")
	proto.RegisterType((*X3DHHeader)(nil), "chat.X3DHHeader")
	proto.RegisterType((*DirectMessageProtocol)(nil), "chat.DirectMessageProtocol")
	proto.RegisterType((*ProtocolMessage)(nil), "chat.ProtocolMessage")
}

func init() { proto.RegisterFile("encryption.proto", fileDescriptor_encryption_a35d9f7c683dd9c2) }

var fileDescriptor_encryption_a35d9f7c683dd9c2 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0x4d, 0x66, 0x19, 0x67, 0x6b, 0x33, 0xd9, 0xa1, 0x45, 0x09, 0xba, 0x07, 0x09, 0x0b,
	0x0a, 0x42, 0x40, 0xe7, 0x0d, 0xd6, 0x1c, 0x82, 0x22, 0x2c, 0xed, 0xc5, 0x5b, 0xe8, 0x49, 0x97,
	0x3b, 0x8d, 0x99, 0x4e, 0xd3, 0xe9, 0x11, 0xf2, 0x0a, 0x3e, 0x90, 0x47, 0x9f, 0x4d, 0xfa, 0x4f,
	0x92, 0x55, 0x77, 0xc0, 0x5b, 0xaa, 0xea, 0xab, 0x5f, 0x7d, 0x55, 0x4d, 0x60, 0x83, 0xb2, 0xd1,
	0x83, 0x32, 0xa2, 0x93, 0x85, 0xd2, 0x9d, 0xe9, 0xc8, 0x59, 0xb3, 0x67, 0x26, 0xdf, 0xc3, 0xf2,
	0xe6, 0x28, 0x79, 0x8b, 0xe4, 0x39, 0xac, 0x04, 0x47, 0x69, 0x84, 0x19, 0xb2, 0xe8, 0x65, 0xf4,
	0x3a, 0xa1, 0x53, 0x4c, 0xae, 0x21, 0xed, 0xc5, 0x9d, 0x44, 0x5e, 0x2b, 0x8d, 0xf5, 0x37, 0x1c,
	0xb2, 0xd8, 0x29, 0x12, 0x9f, 0xbd, 0xd5, 0xf8, 0x11, 0x07, 0x72, 0x05, 0xe7, 0x36, 0x66, 0xe6,
	0xa8, 0x31, 0x5b, 0x38, 0xc1, 0x9c, 0xc8, 0x7f, 0x44, 0x70, 0xe9, 0x47, 0xbd, 0xef, 0xa4, 0x61,
	0x42, 0xa2, 0x26, 0xd7, 0xb0, 0xdc, 0xb9, 0x94, 0x9b, 0x78, 0xf1, 0x2e, 0x29, 0xac, 0xa9, 0xc2,
	0xcb, 0x68, 0xa8, 0x91, 0x2d, 0x3c, 0x53, 0x5a, 0x7c, 0x67, 0x06, 0xeb, 0x07, 0x5d, 0x3c, 0x09,
	0xd5, 0xcf, 0x7f, 0x99, 0x31, 0xe2, 0x80, 0xbd, 0x61, 0x07, 0xe5, 0xcc, 0x2c, 0xe8, 0x9c, 0xc8,
	0x3f, 0xc0, 0xaa, 0xa4, 0x15, 0x32, 0x8e, 0x9a, 0x6c, 0x60, 0x61, 0x59, 0x7e, 0x67, 0xfb, 0x49,
	0x12, 0x88, 0xa4, 0x63, 0xaf, 0x69, 0x24, 0x49, 0x0a, 0xb1, 0x92, 0x0e, 0xb1, 0xa6, 0xb1, 0x72,
	0xb1, 0xe0, 0xd9, 0x99, 0x93, 0xc7, 0x82, 0xe7, 0x57, 0xb0, 0x2a, 0xab, 0x53, 0xac, 0xbc, 0x00,
	0xf8, 0xb2, 0x3d, 0x5d, 0xff, 0x87, 0xf6, 0x2b, 0x82, 0xa7, 0xa5, 0xd0, 0xd8, 0x98, 0x4f, 0xd8,
	0xf7, 0xec, 0x0e, 0x6f, 0xed, 0x6b, 0x35, 0x5d, 0x4b, 0xde, 0xc2, 0x85, 0x25, 0xd5, 0x7b, 0x87,
	0x0a, 0x17, 0xdb, 0xf8, 0x8b, 0xcd, 0x23, 0xe8, 0xfd, 0x71, 0x6f, 0xe0, 0xbc, 0xa4, 0x63, 0x43,
	0xec, 0x1a, 0x52, 0xdf, 0x30, 0x6e, 0x4f, 0xe7, 0x3b, 0x58, 0xf1, 0x44, 0xc7, 0x3f, 0xc4, 0xd5,
	0x24, 0x1e, 0xc9, 0x19, 0x3c, 0x56, 0x6c, 0x68, 0x3b, 0xc6, 0xc3, 0x4b, 0x8f, 0x61, 0xfe, 0x33,
	0x82, 0xcb, 0xd1, 0x73, 0x58, 0xe1, 0x3f, 0xdf, 0xb9, 0x84, 0x94, 0xbb, 0xcd, 0xeb, 0x83, 0xef,
	0x0b, 0x2e, 0x5e, 0x04, 0x17, 0x0f, 0x5d, 0xa5, 0x7a, 0x44, 0xd7, 0xfc, 0x7e, 0x81, 0xbc, 0x82,
	0x54, 0x1d, 0x77, 0xad, 0x68, 0x26, 0xca, 0x57, 0x6b, 0xd0, 0x0a, 0x7d, 0x3e, 0x08, 0x6f, 0x52,
	0x48, 0x82, 0xa2, 0x36, 0x83, 0xc2, 0xdd, 0xd2, 0xfd, 0x17, 0xdb, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xd0, 0x08, 0xb6, 0x41, 0x2b, 0x03, 0x00, 0x00,
}