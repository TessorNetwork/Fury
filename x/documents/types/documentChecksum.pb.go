// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/documents/documentChecksum.proto

package types

import (
	fmt "fmt"
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

type DocumentChecksum struct {
	Value     string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Algorithm string `protobuf:"bytes,2,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
}

func (m *DocumentChecksum) Reset()         { *m = DocumentChecksum{} }
func (m *DocumentChecksum) String() string { return proto.CompactTextString(m) }
func (*DocumentChecksum) ProtoMessage()    {}
func (*DocumentChecksum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9479b6167ee201f9, []int{0}
}
func (m *DocumentChecksum) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DocumentChecksum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DocumentChecksum.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DocumentChecksum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentChecksum.Merge(m, src)
}
func (m *DocumentChecksum) XXX_Size() int {
	return m.Size()
}
func (m *DocumentChecksum) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentChecksum.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentChecksum proto.InternalMessageInfo

func (m *DocumentChecksum) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DocumentChecksum) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

func init() {
	proto.RegisterType((*DocumentChecksum)(nil), "tessornetwork.fury.documents.DocumentChecksum")
}

func init() {
	proto.RegisterFile("fury/documents/documentChecksum.proto", fileDescriptor_9479b6167ee201f9)
}

var fileDescriptor_9479b6167ee201f9 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x2b, 0x2d, 0xaa,
	0xd4, 0x4f, 0xc9, 0x4f, 0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0x86, 0xb3, 0x9c, 0x33, 0x52, 0x93,
	0xb3, 0x8b, 0x4b, 0x73, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x64, 0x4a, 0x52, 0x8b, 0x8b,
	0xf3, 0x8b, 0xf2, 0x52, 0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0xf5, 0x40, 0x9a, 0xf4, 0xe0, 0x9a, 0xa4,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x0a, 0xf5, 0x41, 0x2c, 0x88, 0x1e, 0x25, 0x37, 0x2e, 0x01,
	0x17, 0x34, 0xd3, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x19, 0x2e, 0xce, 0xc4, 0x9c, 0xf4, 0xfc, 0xa2, 0xcc, 0x92,
	0x8c, 0x5c, 0x09, 0x26, 0xb0, 0x0c, 0x42, 0xc0, 0xc9, 0xf3, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xa2, 0xf4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x51, 0x1c, 0xa8, 0x0f, 0xf6, 0x55, 0x05, 0x92, 0xbf, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8,
	0xc0, 0x2e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xab, 0xfc, 0x07, 0xf6, 0x00, 0x00,
	0x00,
}

func (m *DocumentChecksum) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DocumentChecksum) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DocumentChecksum) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Algorithm) > 0 {
		i -= len(m.Algorithm)
		copy(dAtA[i:], m.Algorithm)
		i = encodeVarintDocumentChecksum(dAtA, i, uint64(len(m.Algorithm)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintDocumentChecksum(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDocumentChecksum(dAtA []byte, offset int, v uint64) int {
	offset -= sovDocumentChecksum(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DocumentChecksum) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovDocumentChecksum(uint64(l))
	}
	l = len(m.Algorithm)
	if l > 0 {
		n += 1 + l + sovDocumentChecksum(uint64(l))
	}
	return n
}

func sovDocumentChecksum(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDocumentChecksum(x uint64) (n int) {
	return sovDocumentChecksum(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DocumentChecksum) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocumentChecksum
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
			return fmt.Errorf("proto: DocumentChecksum: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DocumentChecksum: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentChecksum
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
				return ErrInvalidLengthDocumentChecksum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentChecksum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Algorithm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentChecksum
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
				return ErrInvalidLengthDocumentChecksum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentChecksum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Algorithm = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDocumentChecksum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDocumentChecksum
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
func skipDocumentChecksum(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDocumentChecksum
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
					return 0, ErrIntOverflowDocumentChecksum
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
					return 0, ErrIntOverflowDocumentChecksum
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
				return 0, ErrInvalidLengthDocumentChecksum
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDocumentChecksum
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDocumentChecksum
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDocumentChecksum        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDocumentChecksum          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDocumentChecksum = fmt.Errorf("proto: unexpected end of group")
)
