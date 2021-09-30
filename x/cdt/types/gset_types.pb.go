// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cross/cdt/gset_types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type GSetValue struct {
	Elements [][]byte `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (m *GSetValue) Reset()         { *m = GSetValue{} }
func (m *GSetValue) String() string { return proto.CompactTextString(m) }
func (*GSetValue) ProtoMessage()    {}
func (*GSetValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f1afaab50c2ccf6, []int{0}
}
func (m *GSetValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GSetValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GSetValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GSetValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GSetValue.Merge(m, src)
}
func (m *GSetValue) XXX_Size() int {
	return m.Size()
}
func (m *GSetValue) XXX_DiscardUnknown() {
	xxx_messageInfo_GSetValue.DiscardUnknown(m)
}

var xxx_messageInfo_GSetValue proto.InternalMessageInfo

type GSetOP struct {
	K   []byte   `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	Add [][]byte `protobuf:"bytes,2,rep,name=add,proto3" json:"add,omitempty"`
}

func (m *GSetOP) Reset()         { *m = GSetOP{} }
func (m *GSetOP) String() string { return proto.CompactTextString(m) }
func (*GSetOP) ProtoMessage()    {}
func (*GSetOP) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f1afaab50c2ccf6, []int{1}
}
func (m *GSetOP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GSetOP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GSetOP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GSetOP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GSetOP.Merge(m, src)
}
func (m *GSetOP) XXX_Size() int {
	return m.Size()
}
func (m *GSetOP) XXX_DiscardUnknown() {
	xxx_messageInfo_GSetOP.DiscardUnknown(m)
}

var xxx_messageInfo_GSetOP proto.InternalMessageInfo

type GSetCDTState struct {
	AddSet [][]byte `protobuf:"bytes,1,rep,name=add_set,json=addSet,proto3" json:"add_set,omitempty"`
}

func (m *GSetCDTState) Reset()         { *m = GSetCDTState{} }
func (m *GSetCDTState) String() string { return proto.CompactTextString(m) }
func (*GSetCDTState) ProtoMessage()    {}
func (*GSetCDTState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f1afaab50c2ccf6, []int{2}
}
func (m *GSetCDTState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GSetCDTState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GSetCDTState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GSetCDTState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GSetCDTState.Merge(m, src)
}
func (m *GSetCDTState) XXX_Size() int {
	return m.Size()
}
func (m *GSetCDTState) XXX_DiscardUnknown() {
	xxx_messageInfo_GSetCDTState.DiscardUnknown(m)
}

var xxx_messageInfo_GSetCDTState proto.InternalMessageInfo

type GSetInverseOP struct {
	K   []byte   `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	Add [][]byte `protobuf:"bytes,2,rep,name=add,proto3" json:"add,omitempty"`
}

func (m *GSetInverseOP) Reset()         { *m = GSetInverseOP{} }
func (m *GSetInverseOP) String() string { return proto.CompactTextString(m) }
func (*GSetInverseOP) ProtoMessage()    {}
func (*GSetInverseOP) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f1afaab50c2ccf6, []int{3}
}
func (m *GSetInverseOP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GSetInverseOP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GSetInverseOP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GSetInverseOP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GSetInverseOP.Merge(m, src)
}
func (m *GSetInverseOP) XXX_Size() int {
	return m.Size()
}
func (m *GSetInverseOP) XXX_DiscardUnknown() {
	xxx_messageInfo_GSetInverseOP.DiscardUnknown(m)
}

var xxx_messageInfo_GSetInverseOP proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GSetValue)(nil), "cross.cdt.GSetValue")
	proto.RegisterType((*GSetOP)(nil), "cross.cdt.GSetOP")
	proto.RegisterType((*GSetCDTState)(nil), "cross.cdt.GSetCDTState")
	proto.RegisterType((*GSetInverseOP)(nil), "cross.cdt.GSetInverseOP")
}

func init() { proto.RegisterFile("cross/cdt/gset_types.proto", fileDescriptor_5f1afaab50c2ccf6) }

var fileDescriptor_5f1afaab50c2ccf6 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x73, 0x16, 0xa2, 0x3d, 0x22, 0x48, 0x10, 0x8c, 0x19, 0x8e, 0x92, 0xa5, 0x59, 0x9a,
	0x1b, 0xfc, 0x07, 0x2a, 0x88, 0x38, 0x28, 0x8d, 0x38, 0xb8, 0x94, 0x4b, 0xee, 0xf3, 0x5a, 0x9a,
	0xe6, 0x4a, 0xee, 0x8b, 0xd8, 0x7f, 0xe1, 0xcf, 0xea, 0xd8, 0xd1, 0x51, 0x93, 0x3f, 0x22, 0x77,
	0x15, 0xd7, 0x6e, 0xdf, 0xfb, 0xbd, 0x0f, 0xef, 0xf0, 0xd0, 0xb8, 0x6c, 0xb4, 0x31, 0xbc, 0x94,
	0xc8, 0x95, 0x01, 0x9c, 0xe1, 0x66, 0x0d, 0x26, 0x5b, 0x37, 0x1a, 0x75, 0x38, 0x74, 0x5d, 0x56,
	0x4a, 0x8c, 0xcf, 0x95, 0x56, 0xda, 0x7d, 0xb9, 0xbd, 0xf6, 0x40, 0x7c, 0xa9, 0xb4, 0x56, 0x15,
	0x70, 0x97, 0x8a, 0xf6, 0x8d, 0x8b, 0x7a, 0xb3, 0xaf, 0x92, 0x31, 0x1d, 0xde, 0xe5, 0x80, 0x2f,
	0xa2, 0x6a, 0x21, 0x8c, 0xe9, 0x09, 0x54, 0xb0, 0x82, 0x1a, 0x4d, 0x44, 0x46, 0x83, 0x34, 0x98,
	0xfe, 0xe7, 0x24, 0xa5, 0xbe, 0x05, 0x1f, 0x9f, 0xc2, 0x80, 0x92, 0x65, 0x44, 0x46, 0x24, 0x0d,
	0xa6, 0x64, 0x19, 0x9e, 0xd1, 0x81, 0x90, 0x32, 0x3a, 0x72, 0xb8, 0x3d, 0x93, 0x31, 0x0d, 0x2c,
	0x79, 0x73, 0xfb, 0x9c, 0xa3, 0x40, 0x08, 0x2f, 0xe8, 0xb1, 0x90, 0x72, 0x66, 0x00, 0xff, 0x46,
	0x7d, 0x21, 0x65, 0x0e, 0x98, 0x70, 0x7a, 0x6a, 0xc1, 0xfb, 0xfa, 0x1d, 0x1a, 0x03, 0x87, 0x97,
	0xaf, 0x1f, 0xb6, 0x3f, 0xcc, 0xdb, 0x76, 0x8c, 0xec, 0x3a, 0x46, 0xbe, 0x3b, 0x46, 0x3e, 0x7b,
	0xe6, 0xed, 0x7a, 0xe6, 0x7d, 0xf5, 0xcc, 0x7b, 0x9d, 0xa8, 0x05, 0xce, 0xdb, 0x22, 0x2b, 0xf5,
	0x8a, 0x4b, 0x81, 0xa2, 0x9c, 0x8b, 0x45, 0x5d, 0x89, 0x82, 0x3b, 0x3d, 0x13, 0xab, 0xee, 0xc3,
	0x09, 0x74, 0xee, 0x0a, 0xdf, 0x09, 0xb8, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x3d, 0x3a,
	0x56, 0x5a, 0x01, 0x00, 0x00,
}

func (m *GSetValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GSetValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GSetValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Elements) > 0 {
		for iNdEx := len(m.Elements) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Elements[iNdEx])
			copy(dAtA[i:], m.Elements[iNdEx])
			i = encodeVarintGsetTypes(dAtA, i, uint64(len(m.Elements[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GSetOP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GSetOP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GSetOP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Add) > 0 {
		for iNdEx := len(m.Add) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Add[iNdEx])
			copy(dAtA[i:], m.Add[iNdEx])
			i = encodeVarintGsetTypes(dAtA, i, uint64(len(m.Add[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.K) > 0 {
		i -= len(m.K)
		copy(dAtA[i:], m.K)
		i = encodeVarintGsetTypes(dAtA, i, uint64(len(m.K)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GSetCDTState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GSetCDTState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GSetCDTState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AddSet) > 0 {
		for iNdEx := len(m.AddSet) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AddSet[iNdEx])
			copy(dAtA[i:], m.AddSet[iNdEx])
			i = encodeVarintGsetTypes(dAtA, i, uint64(len(m.AddSet[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GSetInverseOP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GSetInverseOP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GSetInverseOP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Add) > 0 {
		for iNdEx := len(m.Add) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Add[iNdEx])
			copy(dAtA[i:], m.Add[iNdEx])
			i = encodeVarintGsetTypes(dAtA, i, uint64(len(m.Add[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.K) > 0 {
		i -= len(m.K)
		copy(dAtA[i:], m.K)
		i = encodeVarintGsetTypes(dAtA, i, uint64(len(m.K)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGsetTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovGsetTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GSetValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Elements) > 0 {
		for _, b := range m.Elements {
			l = len(b)
			n += 1 + l + sovGsetTypes(uint64(l))
		}
	}
	return n
}

func (m *GSetOP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.K)
	if l > 0 {
		n += 1 + l + sovGsetTypes(uint64(l))
	}
	if len(m.Add) > 0 {
		for _, b := range m.Add {
			l = len(b)
			n += 1 + l + sovGsetTypes(uint64(l))
		}
	}
	return n
}

func (m *GSetCDTState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AddSet) > 0 {
		for _, b := range m.AddSet {
			l = len(b)
			n += 1 + l + sovGsetTypes(uint64(l))
		}
	}
	return n
}

func (m *GSetInverseOP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.K)
	if l > 0 {
		n += 1 + l + sovGsetTypes(uint64(l))
	}
	if len(m.Add) > 0 {
		for _, b := range m.Add {
			l = len(b)
			n += 1 + l + sovGsetTypes(uint64(l))
		}
	}
	return n
}

func sovGsetTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGsetTypes(x uint64) (n int) {
	return sovGsetTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GSetValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGsetTypes
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
			return fmt.Errorf("proto: GSetValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GSetValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elements", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGsetTypes
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
				return ErrInvalidLengthGsetTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elements = append(m.Elements, make([]byte, postIndex-iNdEx))
			copy(m.Elements[len(m.Elements)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGsetTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGsetTypes
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
func (m *GSetOP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGsetTypes
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
			return fmt.Errorf("proto: GSetOP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GSetOP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field K", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGsetTypes
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
				return ErrInvalidLengthGsetTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.K = append(m.K[:0], dAtA[iNdEx:postIndex]...)
			if m.K == nil {
				m.K = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Add", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGsetTypes
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
				return ErrInvalidLengthGsetTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Add = append(m.Add, make([]byte, postIndex-iNdEx))
			copy(m.Add[len(m.Add)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGsetTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGsetTypes
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
func (m *GSetCDTState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGsetTypes
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
			return fmt.Errorf("proto: GSetCDTState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GSetCDTState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddSet", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGsetTypes
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
				return ErrInvalidLengthGsetTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddSet = append(m.AddSet, make([]byte, postIndex-iNdEx))
			copy(m.AddSet[len(m.AddSet)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGsetTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGsetTypes
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
func (m *GSetInverseOP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGsetTypes
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
			return fmt.Errorf("proto: GSetInverseOP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GSetInverseOP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field K", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGsetTypes
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
				return ErrInvalidLengthGsetTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.K = append(m.K[:0], dAtA[iNdEx:postIndex]...)
			if m.K == nil {
				m.K = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Add", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGsetTypes
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
				return ErrInvalidLengthGsetTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Add = append(m.Add, make([]byte, postIndex-iNdEx))
			copy(m.Add[len(m.Add)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGsetTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGsetTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGsetTypes
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
func skipGsetTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGsetTypes
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
					return 0, ErrIntOverflowGsetTypes
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
					return 0, ErrIntOverflowGsetTypes
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
				return 0, ErrInvalidLengthGsetTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGsetTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGsetTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGsetTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGsetTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGsetTypes = fmt.Errorf("proto: unexpected end of group")
)
