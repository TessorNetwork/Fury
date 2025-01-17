// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/vbr/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the vbr module's genesis state.
type GenesisState struct {
	PoolAmount github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=poolAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"poolAmount" yaml:"pool_amount"`
	Params     Params                                      `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2da6b8c2d2d00ea5, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetPoolAmount() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.PoolAmount
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tessornetwork.fury.vbr.GenesisState")
}

func init() { proto.RegisterFile("fury/vbr/genesis.proto", fileDescriptor_2da6b8c2d2d00ea5) }

var fileDescriptor_2da6b8c2d2d00ea5 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4e, 0xf3, 0x30,
	0x1c, 0xc5, 0x63, 0x7d, 0x55, 0x87, 0xf4, 0x9b, 0x22, 0xa8, 0xaa, 0x0a, 0xb9, 0x55, 0xa7, 0x22,
	0x84, 0xad, 0x96, 0x0d, 0xb1, 0x90, 0x22, 0x21, 0x36, 0x54, 0x36, 0x16, 0x64, 0x07, 0x13, 0xa2,
	0x36, 0xf9, 0x47, 0xb6, 0x13, 0xc8, 0x09, 0x58, 0x39, 0x07, 0x27, 0xe9, 0xd8, 0x09, 0x31, 0x15,
	0x94, 0xdc, 0x80, 0x13, 0xa0, 0xd8, 0x01, 0x15, 0x89, 0xc9, 0x96, 0xfc, 0xde, 0xef, 0x3d, 0x3f,
	0xb7, 0x7b, 0x97, 0xc9, 0x82, 0xe6, 0x5c, 0xd2, 0x50, 0x24, 0x42, 0x45, 0x8a, 0xa4, 0x12, 0x34,
	0x78, 0x5d, 0x2d, 0x94, 0x02, 0x99, 0x08, 0xfd, 0x00, 0x72, 0x41, 0x6a, 0x15, 0xc9, 0xb9, 0xec,
	0xef, 0x84, 0x10, 0x82, 0x91, 0xd0, 0xfa, 0x66, 0xd5, 0xfd, 0xdd, 0x1f, 0x4a, 0xca, 0x24, 0x8b,
	0x1b, 0x48, 0x1f, 0x07, 0xa0, 0x62, 0x50, 0x94, 0x33, 0x25, 0x68, 0x3e, 0xe1, 0x42, 0xb3, 0x09,
	0x0d, 0x20, 0x4a, 0xec, 0xfb, 0xe8, 0x15, 0xb9, 0xff, 0xcf, 0x6d, 0xec, 0x95, 0x66, 0x5a, 0x78,
	0x4f, 0xc8, 0x75, 0x53, 0x80, 0xe5, 0x69, 0x0c, 0x59, 0xa2, 0x7b, 0x68, 0xf8, 0x6f, 0xdc, 0x99,
	0xee, 0x11, 0x8b, 0x21, 0x35, 0x86, 0x34, 0x18, 0x72, 0x26, 0x82, 0x19, 0x44, 0x89, 0x7f, 0xb1,
	0xda, 0x0c, 0x9c, 0xcf, 0xcd, 0xc0, 0x2b, 0x58, 0xbc, 0x3c, 0x1e, 0xd5, 0xee, 0x1b, 0x66, 0xec,
	0xa3, 0x97, 0xf7, 0xc1, 0x41, 0x18, 0xe9, 0xfb, 0x8c, 0x93, 0x00, 0x62, 0xda, 0x94, 0xb1, 0xc7,
	0xa1, 0xba, 0x5d, 0x50, 0x5d, 0xa4, 0x42, 0x7d, 0x93, 0xd4, 0x7c, 0x2b, 0xda, 0x3b, 0x71, 0xdb,
	0xf6, 0x2b, 0xbd, 0xd6, 0x10, 0x8d, 0x3b, 0x53, 0x4c, 0xfe, 0x1e, 0x84, 0x5c, 0x1a, 0x95, 0xdf,
	0xaa, 0x6b, 0xcc, 0x1b, 0x8f, 0x3f, 0x5b, 0x95, 0x18, 0xad, 0x4b, 0x8c, 0x3e, 0x4a, 0x8c, 0x9e,
	0x2b, 0xec, 0xac, 0x2b, 0xec, 0xbc, 0x55, 0xd8, 0xb9, 0xde, 0xdf, 0x2a, 0xf4, 0x8b, 0x48, 0xcd,
	0x84, 0x8f, 0x66, 0x44, 0xd3, 0x8b, 0xb7, 0xcd, 0x48, 0x47, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xc8, 0x82, 0xa5, 0x63, 0xa3, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.PoolAmount) > 0 {
		for iNdEx := len(m.PoolAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PoolAmount) > 0 {
		for _, e := range m.PoolAmount {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAmount = append(m.PoolAmount, types.DecCoin{})
			if err := m.PoolAmount[len(m.PoolAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
