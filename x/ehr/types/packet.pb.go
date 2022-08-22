// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ehr/packet.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type EhrPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*EhrPacketData_NoData
	//	*EhrPacketData_IbcTransferEhrsPacket
	Packet isEhrPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *EhrPacketData) Reset()         { *m = EhrPacketData{} }
func (m *EhrPacketData) String() string { return proto.CompactTextString(m) }
func (*EhrPacketData) ProtoMessage()    {}
func (*EhrPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a1144014c8875d5, []int{0}
}
func (m *EhrPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EhrPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EhrPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EhrPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EhrPacketData.Merge(m, src)
}
func (m *EhrPacketData) XXX_Size() int {
	return m.Size()
}
func (m *EhrPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_EhrPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_EhrPacketData proto.InternalMessageInfo

type isEhrPacketData_Packet interface {
	isEhrPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type EhrPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type EhrPacketData_IbcTransferEhrsPacket struct {
	IbcTransferEhrsPacket *IbcTransferEhrsPacketData `protobuf:"bytes,2,opt,name=ibcTransferEhrsPacket,proto3,oneof" json:"ibcTransferEhrsPacket,omitempty"`
}

func (*EhrPacketData_NoData) isEhrPacketData_Packet()                {}
func (*EhrPacketData_IbcTransferEhrsPacket) isEhrPacketData_Packet() {}

func (m *EhrPacketData) GetPacket() isEhrPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *EhrPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*EhrPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *EhrPacketData) GetIbcTransferEhrsPacket() *IbcTransferEhrsPacketData {
	if x, ok := m.GetPacket().(*EhrPacketData_IbcTransferEhrsPacket); ok {
		return x.IbcTransferEhrsPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EhrPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EhrPacketData_NoData)(nil),
		(*EhrPacketData_IbcTransferEhrsPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a1144014c8875d5, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// IbcTransferEhrsPacketData defines a struct for the packet payload
type IbcTransferEhrsPacketData struct {
	Creator             string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PatientId           uint64 `protobuf:"varint,2,opt,name=patientId,proto3" json:"patientId,omitempty"`
	OrganizationAddress string `protobuf:"bytes,3,opt,name=organizationAddress,proto3" json:"organizationAddress,omitempty"`
	PendingTransferUrl  string `protobuf:"bytes,4,opt,name=pendingTransferUrl,proto3" json:"pendingTransferUrl,omitempty"`
}

func (m *IbcTransferEhrsPacketData) Reset()         { *m = IbcTransferEhrsPacketData{} }
func (m *IbcTransferEhrsPacketData) String() string { return proto.CompactTextString(m) }
func (*IbcTransferEhrsPacketData) ProtoMessage()    {}
func (*IbcTransferEhrsPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a1144014c8875d5, []int{2}
}
func (m *IbcTransferEhrsPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IbcTransferEhrsPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IbcTransferEhrsPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IbcTransferEhrsPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IbcTransferEhrsPacketData.Merge(m, src)
}
func (m *IbcTransferEhrsPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IbcTransferEhrsPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IbcTransferEhrsPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IbcTransferEhrsPacketData proto.InternalMessageInfo

func (m *IbcTransferEhrsPacketData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *IbcTransferEhrsPacketData) GetPatientId() uint64 {
	if m != nil {
		return m.PatientId
	}
	return 0
}

func (m *IbcTransferEhrsPacketData) GetOrganizationAddress() string {
	if m != nil {
		return m.OrganizationAddress
	}
	return ""
}

func (m *IbcTransferEhrsPacketData) GetPendingTransferUrl() string {
	if m != nil {
		return m.PendingTransferUrl
	}
	return ""
}

// IbcTransferEhrsPacketAck defines a struct for the packet acknowledgment
type IbcTransferEhrsPacketAck struct {
	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (m *IbcTransferEhrsPacketAck) Reset()         { *m = IbcTransferEhrsPacketAck{} }
func (m *IbcTransferEhrsPacketAck) String() string { return proto.CompactTextString(m) }
func (*IbcTransferEhrsPacketAck) ProtoMessage()    {}
func (*IbcTransferEhrsPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a1144014c8875d5, []int{3}
}
func (m *IbcTransferEhrsPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IbcTransferEhrsPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IbcTransferEhrsPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IbcTransferEhrsPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IbcTransferEhrsPacketAck.Merge(m, src)
}
func (m *IbcTransferEhrsPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *IbcTransferEhrsPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_IbcTransferEhrsPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_IbcTransferEhrsPacketAck proto.InternalMessageInfo

func (m *IbcTransferEhrsPacketAck) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*EhrPacketData)(nil), "ashrafmohey.cache.ehr.EhrPacketData")
	proto.RegisterType((*NoData)(nil), "ashrafmohey.cache.ehr.NoData")
	proto.RegisterType((*IbcTransferEhrsPacketData)(nil), "ashrafmohey.cache.ehr.IbcTransferEhrsPacketData")
	proto.RegisterType((*IbcTransferEhrsPacketAck)(nil), "ashrafmohey.cache.ehr.IbcTransferEhrsPacketAck")
}

func init() { proto.RegisterFile("ehr/packet.proto", fileDescriptor_3a1144014c8875d5) }

var fileDescriptor_3a1144014c8875d5 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x37, 0xb6, 0xac, 0x6d, 0x44, 0x28, 0x91, 0x42, 0x04, 0x5d, 0xca, 0x9e, 0x2a, 0x68,
	0xb6, 0xe8, 0xc1, 0x73, 0xab, 0x05, 0x7b, 0x11, 0x59, 0xf4, 0xe2, 0x2d, 0xcd, 0xa6, 0x4d, 0xa8,
	0x4d, 0x96, 0x24, 0x82, 0xf5, 0x29, 0x7c, 0x19, 0x6f, 0x3e, 0x80, 0xc7, 0x1e, 0x3d, 0x4a, 0xfb,
	0x22, 0xb2, 0xd9, 0x16, 0x2f, 0xdb, 0xdb, 0x24, 0x33, 0xff, 0xc7, 0x3f, 0x3f, 0x03, 0x5b, 0x5c,
	0x98, 0x24, 0xa7, 0x6c, 0xc6, 0x1d, 0xc9, 0x8d, 0x76, 0x1a, 0xb5, 0xa9, 0x15, 0x86, 0x4e, 0xe6,
	0x5a, 0xf0, 0x05, 0x61, 0x94, 0x09, 0x4e, 0xb8, 0x30, 0xf1, 0x17, 0x80, 0x87, 0x43, 0x61, 0x1e,
	0xfc, 0xe8, 0x2d, 0x75, 0x14, 0x5d, 0xc3, 0x50, 0xe9, 0xa2, 0xc2, 0xa0, 0x03, 0xba, 0x07, 0x97,
	0xa7, 0xa4, 0x52, 0x49, 0xee, 0xfd, 0xd0, 0x5d, 0x90, 0x6e, 0xc6, 0x91, 0x80, 0x6d, 0x39, 0x66,
	0x8f, 0x86, 0x2a, 0x3b, 0xe1, 0x66, 0x28, 0x8c, 0x2d, 0xa9, 0x78, 0xcf, 0x73, 0x7a, 0x3b, 0x38,
	0xa3, 0x2a, 0xcd, 0x06, 0x5d, 0x0d, 0x1c, 0x34, 0x60, 0x58, 0xee, 0x16, 0x37, 0x60, 0x58, 0xfa,
	0x88, 0x3f, 0x01, 0x3c, 0xde, 0x89, 0x42, 0x18, 0xee, 0x33, 0xc3, 0xa9, 0xd3, 0xc6, 0x6f, 0xd5,
	0x4c, 0xb7, 0x4f, 0x74, 0x02, 0x9b, 0x39, 0x75, 0x92, 0x2b, 0x37, 0xca, 0xbc, 0xd3, 0x7a, 0xfa,
	0xff, 0x81, 0x7a, 0xf0, 0x48, 0x9b, 0x29, 0x55, 0xf2, 0x9d, 0x3a, 0xa9, 0x55, 0x3f, 0xcb, 0x0c,
	0xb7, 0x16, 0xd7, 0x3c, 0xa3, 0xaa, 0x85, 0x08, 0x44, 0x39, 0x57, 0x99, 0x54, 0xd3, 0xad, 0x95,
	0x27, 0xf3, 0x82, 0xeb, 0x5e, 0x50, 0xd1, 0x89, 0xcf, 0x21, 0xae, 0xb4, 0xdd, 0x67, 0x33, 0xd4,
	0x82, 0x35, 0x99, 0x59, 0x0c, 0x3a, 0xb5, 0x6e, 0x3d, 0x2d, 0xca, 0xc1, 0xcd, 0xf7, 0x2a, 0x02,
	0xcb, 0x55, 0x04, 0x7e, 0x57, 0x11, 0xf8, 0x58, 0x47, 0xc1, 0x72, 0x1d, 0x05, 0x3f, 0xeb, 0x28,
	0x78, 0x3e, 0x9b, 0x4a, 0x27, 0x5e, 0xc7, 0x84, 0xe9, 0x79, 0x52, 0x06, 0x7d, 0xe1, 0x93, 0x4e,
	0x7c, 0xd2, 0xc9, 0x5b, 0x52, 0x5c, 0x84, 0x5b, 0xe4, 0xdc, 0x8e, 0x43, 0x7f, 0x11, 0x57, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0xbb, 0xd6, 0x6e, 0x25, 0x02, 0x00, 0x00,
}

func (m *EhrPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EhrPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EhrPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *EhrPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EhrPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *EhrPacketData_IbcTransferEhrsPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EhrPacketData_IbcTransferEhrsPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IbcTransferEhrsPacket != nil {
		{
			size, err := m.IbcTransferEhrsPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *IbcTransferEhrsPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IbcTransferEhrsPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IbcTransferEhrsPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PendingTransferUrl) > 0 {
		i -= len(m.PendingTransferUrl)
		copy(dAtA[i:], m.PendingTransferUrl)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.PendingTransferUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OrganizationAddress) > 0 {
		i -= len(m.OrganizationAddress)
		copy(dAtA[i:], m.OrganizationAddress)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.OrganizationAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PatientId != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.PatientId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IbcTransferEhrsPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IbcTransferEhrsPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IbcTransferEhrsPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		dAtA4 := make([]byte, len(m.Ids)*10)
		var j3 int
		for _, num := range m.Ids {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintPacket(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EhrPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *EhrPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *EhrPacketData_IbcTransferEhrsPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IbcTransferEhrsPacket != nil {
		l = m.IbcTransferEhrsPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *IbcTransferEhrsPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.PatientId != 0 {
		n += 1 + sovPacket(uint64(m.PatientId))
	}
	l = len(m.OrganizationAddress)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.PendingTransferUrl)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *IbcTransferEhrsPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovPacket(uint64(e))
		}
		n += 1 + sovPacket(uint64(l)) + l
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EhrPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: EhrPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EhrPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &EhrPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcTransferEhrsPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IbcTransferEhrsPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &EhrPacketData_IbcTransferEhrsPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *IbcTransferEhrsPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: IbcTransferEhrsPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IbcTransferEhrsPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PatientId", wireType)
			}
			m.PatientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PatientId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrganizationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrganizationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingTransferUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingTransferUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *IbcTransferEhrsPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: IbcTransferEhrsPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IbcTransferEhrsPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPacket
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPacket
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPacket
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPacket
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPacket
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
