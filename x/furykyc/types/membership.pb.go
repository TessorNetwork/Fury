// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/furykyc/membership.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Membership struct {
	Owner          string     `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	TspAddress     string     `protobuf:"bytes,2,opt,name=tsp_address,json=tspAddress,proto3" json:"tsp_address,omitempty"`
	MembershipType string     `protobuf:"bytes,3,opt,name=membership_type,json=membershipType,proto3" json:"membership_type,omitempty"`
	ExpiryAt       *time.Time `protobuf:"bytes,4,opt,name=expiry_at,json=expiryAt,proto3,stdtime" json:"expiry_at,omitempty"`
}

func (m *Membership) Reset()         { *m = Membership{} }
func (m *Membership) String() string { return proto.CompactTextString(m) }
func (*Membership) ProtoMessage()    {}
func (*Membership) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7216138f4a4ea4f, []int{0}
}
func (m *Membership) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Membership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Membership.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Membership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Membership.Merge(m, src)
}
func (m *Membership) XXX_Size() int {
	return m.Size()
}
func (m *Membership) XXX_DiscardUnknown() {
	xxx_messageInfo_Membership.DiscardUnknown(m)
}

var xxx_messageInfo_Membership proto.InternalMessageInfo

func (m *Membership) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Membership) GetTspAddress() string {
	if m != nil {
		return m.TspAddress
	}
	return ""
}

func (m *Membership) GetMembershipType() string {
	if m != nil {
		return m.MembershipType
	}
	return ""
}

func (m *Membership) GetExpiryAt() *time.Time {
	if m != nil {
		return m.ExpiryAt
	}
	return nil
}

type Invite struct {
	Sender           string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	SenderMembership string `protobuf:"bytes,2,opt,name=sender_membership,json=senderMembership,proto3" json:"sender_membership,omitempty"`
	User             string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Status           uint64 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Invite) Reset()         { *m = Invite{} }
func (m *Invite) String() string { return proto.CompactTextString(m) }
func (*Invite) ProtoMessage()    {}
func (*Invite) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7216138f4a4ea4f, []int{1}
}
func (m *Invite) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Invite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Invite.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Invite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invite.Merge(m, src)
}
func (m *Invite) XXX_Size() int {
	return m.Size()
}
func (m *Invite) XXX_DiscardUnknown() {
	xxx_messageInfo_Invite.DiscardUnknown(m)
}

var xxx_messageInfo_Invite proto.InternalMessageInfo

func (m *Invite) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Invite) GetSenderMembership() string {
	if m != nil {
		return m.SenderMembership
	}
	return ""
}

func (m *Invite) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Invite) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type TrustedServiceProviders struct {
	// option (gogoproto.goproto_stringer) = false;
	// option (gogoproto.stringer)         = true;
	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (m *TrustedServiceProviders) Reset()         { *m = TrustedServiceProviders{} }
func (m *TrustedServiceProviders) String() string { return proto.CompactTextString(m) }
func (*TrustedServiceProviders) ProtoMessage()    {}
func (*TrustedServiceProviders) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7216138f4a4ea4f, []int{2}
}
func (m *TrustedServiceProviders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrustedServiceProviders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrustedServiceProviders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrustedServiceProviders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustedServiceProviders.Merge(m, src)
}
func (m *TrustedServiceProviders) XXX_Size() int {
	return m.Size()
}
func (m *TrustedServiceProviders) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustedServiceProviders.DiscardUnknown(m)
}

var xxx_messageInfo_TrustedServiceProviders proto.InternalMessageInfo

func (m *TrustedServiceProviders) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Membership)(nil), "tessornetwork.fury.furykyc.Membership")
	proto.RegisterType((*Invite)(nil), "tessornetwork.fury.furykyc.Invite")
	proto.RegisterType((*TrustedServiceProviders)(nil), "tessornetwork.fury.furykyc.TrustedServiceProviders")
}

func init() { proto.RegisterFile("fury/furykyc/membership.proto", fileDescriptor_d7216138f4a4ea4f) }

var fileDescriptor_d7216138f4a4ea4f = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xbf, 0x6e, 0xa3, 0x40,
	0x10, 0xc6, 0xbd, 0x67, 0xce, 0x3a, 0xd6, 0xd2, 0xfd, 0x59, 0x59, 0x77, 0x08, 0xdd, 0x61, 0xcb,
	0xcd, 0x59, 0x3a, 0x1d, 0x48, 0x49, 0x91, 0x2a, 0x85, 0xdd, 0x44, 0x29, 0x22, 0x45, 0xc4, 0x55,
	0x1a, 0x04, 0x66, 0x8c, 0x91, 0x03, 0x8b, 0x76, 0x07, 0xdb, 0xbc, 0x85, 0xdf, 0x23, 0x2f, 0x92,
	0xd2, 0x65, 0xba, 0x44, 0xf6, 0x8b, 0x44, 0x2c, 0xd8, 0x24, 0x0d, 0xfa, 0xe6, 0xdb, 0x6f, 0x99,
	0x1f, 0xc3, 0xd0, 0x3f, 0xf3, 0x5c, 0x14, 0x4e, 0xf9, 0x58, 0x16, 0x33, 0x27, 0x81, 0x24, 0x00,
	0x21, 0x17, 0x71, 0x66, 0x67, 0x82, 0x23, 0x67, 0x26, 0x82, 0x94, 0x5c, 0xa4, 0x80, 0x6b, 0x2e,
	0x96, 0x76, 0x99, 0xb3, 0xeb, 0xb0, 0xd9, 0x8f, 0x38, 0x8f, 0x1e, 0xc0, 0x51, 0xc9, 0x20, 0x9f,
	0x3b, 0x18, 0x27, 0x20, 0xd1, 0x4f, 0xea, 0xcb, 0x66, 0x2f, 0xe2, 0x11, 0x57, 0xd2, 0x29, 0x55,
	0xe5, 0x0e, 0x1f, 0x09, 0xa5, 0x37, 0xa7, 0x3e, 0xac, 0x47, 0x3f, 0xf3, 0x75, 0x0a, 0xc2, 0x20,
	0x03, 0x32, 0xd2, 0xdd, 0xaa, 0x60, 0x7d, 0xda, 0x45, 0x99, 0x79, 0x7e, 0x18, 0x0a, 0x90, 0xd2,
	0xf8, 0xa4, 0xce, 0x28, 0xca, 0x6c, 0x5c, 0x39, 0xec, 0x2f, 0xfd, 0xd6, 0xc0, 0x7a, 0x58, 0x64,
	0x60, 0xb4, 0x55, 0xe8, 0x6b, 0x63, 0x4f, 0x8b, 0x0c, 0xd8, 0x25, 0xd5, 0x61, 0x93, 0xc5, 0xa2,
	0xf0, 0x7c, 0x34, 0xb4, 0x01, 0x19, 0x75, 0xcf, 0x4c, 0xbb, 0x22, 0xb7, 0x8f, 0xe4, 0xf6, 0xf4,
	0x48, 0x3e, 0xd1, 0xb6, 0x2f, 0x7d, 0xe2, 0x7e, 0xa9, 0xae, 0x8c, 0x71, 0x58, 0xd0, 0xce, 0x75,
	0xba, 0x8a, 0x11, 0xd8, 0x4f, 0xda, 0x91, 0x90, 0x86, 0x27, 0xd2, 0xba, 0x62, 0xff, 0xe8, 0x8f,
	0x4a, 0x79, 0x4d, 0xe7, 0x1a, 0xf8, 0x7b, 0x75, 0xf0, 0xee, 0x6b, 0x19, 0xd5, 0x72, 0x09, 0xa2,
	0x66, 0x55, 0x5a, 0xbd, 0x18, 0x7d, 0xcc, 0xa5, 0xc2, 0xd3, 0xdc, 0xba, 0x1a, 0x5e, 0xd0, 0x5f,
	0x53, 0x91, 0x4b, 0x84, 0xf0, 0x0e, 0xc4, 0x2a, 0x9e, 0xc1, 0xad, 0xe0, 0xab, 0x38, 0x04, 0x21,
	0xd9, 0x6f, 0xaa, 0xd7, 0xa3, 0x01, 0x69, 0x90, 0x41, 0x7b, 0xa4, 0xbb, 0x8d, 0x31, 0xb9, 0x7a,
	0xda, 0x5b, 0x64, 0xb7, 0xb7, 0xc8, 0xeb, 0xde, 0x22, 0xdb, 0x83, 0xd5, 0xda, 0x1d, 0xac, 0xd6,
	0xf3, 0xc1, 0x6a, 0xdd, 0xff, 0x8f, 0x62, 0x5c, 0xe4, 0x81, 0x3d, 0xe3, 0x89, 0xf3, 0xe1, 0xcf,
	0xaa, 0x0d, 0x70, 0x36, 0xa7, 0x45, 0x28, 0x07, 0x2a, 0x83, 0x8e, 0x1a, 0xd0, 0xf9, 0x5b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x5e, 0x2c, 0xc5, 0x8f, 0x25, 0x02, 0x00, 0x00,
}

func (m *Membership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Membership) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Membership) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiryAt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ExpiryAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpiryAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintMembership(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MembershipType) > 0 {
		i -= len(m.MembershipType)
		copy(dAtA[i:], m.MembershipType)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.MembershipType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TspAddress) > 0 {
		i -= len(m.TspAddress)
		copy(dAtA[i:], m.TspAddress)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.TspAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Invite) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Invite) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Invite) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintMembership(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SenderMembership) > 0 {
		i -= len(m.SenderMembership)
		copy(dAtA[i:], m.SenderMembership)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.SenderMembership)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TrustedServiceProviders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrustedServiceProviders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrustedServiceProviders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintMembership(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMembership(dAtA []byte, offset int, v uint64) int {
	offset -= sovMembership(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Membership) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	l = len(m.TspAddress)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	l = len(m.MembershipType)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.ExpiryAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpiryAt)
		n += 1 + l + sovMembership(uint64(l))
	}
	return n
}

func (m *Invite) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	l = len(m.SenderMembership)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovMembership(uint64(m.Status))
	}
	return n
}

func (m *TrustedServiceProviders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovMembership(uint64(l))
		}
	}
	return n
}

func sovMembership(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMembership(x uint64) (n int) {
	return sovMembership(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Membership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: Membership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Membership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TspAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TspAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MembershipType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MembershipType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpiryAt == nil {
				m.ExpiryAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ExpiryAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *Invite) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: Invite: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Invite: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderMembership", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderMembership = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *TrustedServiceProviders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: TrustedServiceProviders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrustedServiceProviders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMembership
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
func skipMembership(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMembership
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
					return 0, ErrIntOverflowMembership
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
					return 0, ErrIntOverflowMembership
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
				return 0, ErrInvalidLengthMembership
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMembership
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMembership
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMembership        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMembership          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMembership = fmt.Errorf("proto: unexpected end of group")
)
