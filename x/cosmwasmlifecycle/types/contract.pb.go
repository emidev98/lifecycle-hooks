// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmwasmlifecycle/contract.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Contract defines the parameters for the module.
type Contract struct {
	// Amount of strikes that the contract has at the moment.
	Strikes int64 `protobuf:"varint,1,opt,name=strikes,proto3" json:"strikes,omitempty"`
	// Contract's execution type
	ExecutionType ExecutionType `protobuf:"varint,2,opt,name=execution_type,json=executionType,proto3,enum=cosmwasmlifecycle.ExecutionType" json:"execution_type,omitempty"`
	// Collateral deposited to the contract.
	Deposit github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=deposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"deposit"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_59daeaab22a14127, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetStrikes() int64 {
	if m != nil {
		return m.Strikes
	}
	return 0
}

func (m *Contract) GetExecutionType() ExecutionType {
	if m != nil {
		return m.ExecutionType
	}
	return ExecutionType_BEGIN_AND_END_BLOCK
}

func init() {
	proto.RegisterType((*Contract)(nil), "cosmwasmlifecycle.Contract")
}

func init() { proto.RegisterFile("cosmwasmlifecycle/contract.proto", fileDescriptor_59daeaab22a14127) }

var fileDescriptor_59daeaab22a14127 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xc1, 0x4e, 0xc2, 0x40,
	0x14, 0xec, 0x4a, 0x22, 0xa6, 0x46, 0x12, 0x1b, 0x0f, 0x85, 0xc3, 0xd2, 0x78, 0x50, 0x2e, 0xec,
	0x06, 0xbc, 0xe8, 0xcd, 0x40, 0x8c, 0x77, 0xe2, 0x89, 0x0b, 0x69, 0xb7, 0x4f, 0xdc, 0x40, 0xfb,
	0x1a, 0x76, 0x41, 0xf8, 0x0b, 0x3f, 0x8b, 0x23, 0x89, 0x17, 0xe3, 0x81, 0x18, 0xfa, 0x23, 0xa6,
	0xed, 0x36, 0x62, 0x7a, 0xda, 0x9d, 0xcc, 0xbc, 0xf7, 0x66, 0xc6, 0xf6, 0x04, 0xaa, 0xe8, 0xdd,
	0x57, 0xd1, 0x5c, 0xbe, 0x82, 0xd8, 0x88, 0x39, 0x70, 0x81, 0xb1, 0x5e, 0xf8, 0x42, 0xb3, 0x64,
	0x81, 0x1a, 0x9d, 0xcb, 0x8a, 0xa2, 0x75, 0x35, 0xc5, 0x29, 0xe6, 0x2c, 0xcf, 0x7e, 0x85, 0xb0,
	0xd5, 0xcc, 0x84, 0xa8, 0x26, 0x05, 0x51, 0x00, 0x43, 0xd1, 0x02, 0xf1, 0xc0, 0x57, 0xc0, 0x57,
	0xbd, 0x00, 0xb4, 0xdf, 0xe3, 0x02, 0x65, 0x6c, 0xf8, 0x9b, 0xaa, 0x0b, 0x58, 0x83, 0x58, 0x6a,
	0x89, 0xf1, 0x44, 0x6f, 0x12, 0x28, 0x74, 0xd7, 0x9f, 0xc4, 0x3e, 0x1b, 0x1a, 0x7b, 0x8e, 0x6b,
	0xd7, 0x95, 0x5e, 0xc8, 0x19, 0x28, 0x97, 0x78, 0xa4, 0x53, 0x1b, 0x95, 0xd0, 0x79, 0xb6, 0x1b,
	0xff, 0xc7, 0xdd, 0x13, 0x8f, 0x74, 0x1a, 0x7d, 0x8f, 0x55, 0xee, 0xb0, 0xa7, 0x52, 0xf8, 0xb2,
	0x49, 0x60, 0x74, 0x01, 0xc7, 0xd0, 0x09, 0xed, 0x7a, 0x08, 0x09, 0x2a, 0xa9, 0xdd, 0x9a, 0x47,
	0x3a, 0xe7, 0xfd, 0x26, 0x33, 0xb9, 0xb2, 0x24, 0xcc, 0x24, 0x61, 0x43, 0x94, 0xf1, 0x80, 0x6f,
	0xf7, 0x6d, 0xeb, 0x7b, 0xdf, 0xbe, 0x9d, 0x4a, 0xfd, 0xb6, 0x0c, 0x98, 0xc0, 0xc8, 0x94, 0x60,
	0x9e, 0xae, 0x0a, 0x67, 0x3c, 0x73, 0xa3, 0xf2, 0x81, 0x51, 0xb9, 0x7a, 0x30, 0xde, 0x1e, 0x28,
	0xd9, 0x1d, 0x28, 0xf9, 0x39, 0x50, 0xf2, 0x91, 0x52, 0x6b, 0x97, 0x52, 0xeb, 0x2b, 0xa5, 0xd6,
	0xf8, 0xf1, 0x68, 0x17, 0x44, 0x32, 0x84, 0xd5, 0xc3, 0x3d, 0x2f, 0x33, 0x74, 0xff, 0xca, 0x5a,
	0xf3, 0x6a, 0x81, 0xf9, 0xa5, 0xe0, 0x34, 0x2f, 0xee, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0xba,
	0xec, 0x75, 0x06, 0xe8, 0x01, 0x00, 0x00,
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Deposit.Size()
		i -= size
		if _, err := m.Deposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintContract(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.ExecutionType != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.ExecutionType))
		i--
		dAtA[i] = 0x10
	}
	if m.Strikes != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.Strikes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Strikes != 0 {
		n += 1 + sovContract(uint64(m.Strikes))
	}
	if m.ExecutionType != 0 {
		n += 1 + sovContract(uint64(m.ExecutionType))
	}
	l = m.Deposit.Size()
	n += 1 + l + sovContract(uint64(l))
	return n
}

func sovContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContract(x uint64) (n int) {
	return sovContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strikes", wireType)
			}
			m.Strikes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Strikes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionType", wireType)
			}
			m.ExecutionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionType |= ExecutionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContract
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
func skipContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
				return 0, ErrInvalidLengthContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContract = fmt.Errorf("proto: unexpected end of group")
)