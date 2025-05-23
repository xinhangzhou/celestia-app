// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: celestia/core/v1/da/data_availability_header.proto

package da

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

// DataAvailabilityHeader contains the row and column roots of the erasure
// coded version of the data in Block.Data.
// Therefore the original Block.Data is arranged in a
// k × k matrix, which is then "extended" to a
// 2k × 2k matrix applying multiple times Reed-Solomon encoding.
// For details see Section 5.2: https://arxiv.org/abs/1809.09044
// or the Celestia specification:
// https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#availabledataheader
// Note that currently we list row and column roots in separate fields
// (different from the spec).
type DataAvailabilityHeader struct {
	// RowRoot_j = root((M_{j,1} || M_{j,2} || ... || M_{j,2k} ))
	RowRoots [][]byte `protobuf:"bytes,1,rep,name=row_roots,json=rowRoots,proto3" json:"row_roots,omitempty"`
	// ColumnRoot_j = root((M_{1,j} || M_{2,j} || ... || M_{2k,j} ))
	ColumnRoots [][]byte `protobuf:"bytes,2,rep,name=column_roots,json=columnRoots,proto3" json:"column_roots,omitempty"`
}

func (m *DataAvailabilityHeader) Reset()         { *m = DataAvailabilityHeader{} }
func (m *DataAvailabilityHeader) String() string { return proto.CompactTextString(m) }
func (*DataAvailabilityHeader) ProtoMessage()    {}
func (*DataAvailabilityHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c1f9dadbff5429, []int{0}
}
func (m *DataAvailabilityHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataAvailabilityHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataAvailabilityHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataAvailabilityHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataAvailabilityHeader.Merge(m, src)
}
func (m *DataAvailabilityHeader) XXX_Size() int {
	return m.Size()
}
func (m *DataAvailabilityHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DataAvailabilityHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DataAvailabilityHeader proto.InternalMessageInfo

func (m *DataAvailabilityHeader) GetRowRoots() [][]byte {
	if m != nil {
		return m.RowRoots
	}
	return nil
}

func (m *DataAvailabilityHeader) GetColumnRoots() [][]byte {
	if m != nil {
		return m.ColumnRoots
	}
	return nil
}

func init() {
	proto.RegisterType((*DataAvailabilityHeader)(nil), "celestia.core.v1.da.DataAvailabilityHeader")
}

func init() {
	proto.RegisterFile("celestia/core/v1/da/data_availability_header.proto", fileDescriptor_e4c1f9dadbff5429)
}

var fileDescriptor_e4c1f9dadbff5429 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0x4e, 0xcd, 0x49,
	0x2d, 0x2e, 0xc9, 0x4c, 0xd4, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0x49, 0xd4,
	0x4f, 0x49, 0x2c, 0x49, 0x8c, 0x4f, 0x2c, 0x4b, 0xcc, 0xcc, 0x49, 0x4c, 0xca, 0xcc, 0xc9, 0x2c,
	0xa9, 0x8c, 0xcf, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x86, 0xe9, 0xd1, 0x03, 0xe9, 0xd1, 0x2b, 0x33, 0xd4, 0x4b, 0x49, 0x54, 0x8a, 0xe0, 0x12, 0x73,
	0x49, 0x2c, 0x49, 0x74, 0x44, 0xd2, 0xe5, 0x01, 0xd6, 0x24, 0x24, 0xcd, 0xc5, 0x59, 0x94, 0x5f,
	0x1e, 0x5f, 0x94, 0x9f, 0x5f, 0x52, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x13, 0xc4, 0x51, 0x94,
	0x5f, 0x1e, 0x04, 0xe2, 0x0b, 0x29, 0x72, 0xf1, 0x24, 0xe7, 0xe7, 0x94, 0xe6, 0xe6, 0x41, 0xe5,
	0x99, 0xc0, 0xf2, 0xdc, 0x10, 0x31, 0xb0, 0x12, 0xa7, 0xf0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0x88, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x87, 0xb9, 0x29, 0xbf, 0x28, 0x1d, 0xce, 0xd6, 0x4d, 0x2c, 0x28, 0xd0, 0x07, 0xbb, 0x59, 0x1f,
	0x8b, 0x37, 0x93, 0xd8, 0xc0, 0x52, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x22, 0xee,
	0xe2, 0x04, 0x01, 0x00, 0x00,
}

func (m *DataAvailabilityHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataAvailabilityHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataAvailabilityHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ColumnRoots) > 0 {
		for iNdEx := len(m.ColumnRoots) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ColumnRoots[iNdEx])
			copy(dAtA[i:], m.ColumnRoots[iNdEx])
			i = encodeVarintDataAvailabilityHeader(dAtA, i, uint64(len(m.ColumnRoots[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RowRoots) > 0 {
		for iNdEx := len(m.RowRoots) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RowRoots[iNdEx])
			copy(dAtA[i:], m.RowRoots[iNdEx])
			i = encodeVarintDataAvailabilityHeader(dAtA, i, uint64(len(m.RowRoots[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataAvailabilityHeader(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataAvailabilityHeader(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataAvailabilityHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RowRoots) > 0 {
		for _, b := range m.RowRoots {
			l = len(b)
			n += 1 + l + sovDataAvailabilityHeader(uint64(l))
		}
	}
	if len(m.ColumnRoots) > 0 {
		for _, b := range m.ColumnRoots {
			l = len(b)
			n += 1 + l + sovDataAvailabilityHeader(uint64(l))
		}
	}
	return n
}

func sovDataAvailabilityHeader(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataAvailabilityHeader(x uint64) (n int) {
	return sovDataAvailabilityHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataAvailabilityHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataAvailabilityHeader
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
			return fmt.Errorf("proto: DataAvailabilityHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataAvailabilityHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RowRoots", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataAvailabilityHeader
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
				return ErrInvalidLengthDataAvailabilityHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDataAvailabilityHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RowRoots = append(m.RowRoots, make([]byte, postIndex-iNdEx))
			copy(m.RowRoots[len(m.RowRoots)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnRoots", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataAvailabilityHeader
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
				return ErrInvalidLengthDataAvailabilityHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDataAvailabilityHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ColumnRoots = append(m.ColumnRoots, make([]byte, postIndex-iNdEx))
			copy(m.ColumnRoots[len(m.ColumnRoots)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataAvailabilityHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataAvailabilityHeader
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
func skipDataAvailabilityHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataAvailabilityHeader
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
					return 0, ErrIntOverflowDataAvailabilityHeader
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
					return 0, ErrIntOverflowDataAvailabilityHeader
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
				return 0, ErrInvalidLengthDataAvailabilityHeader
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataAvailabilityHeader
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataAvailabilityHeader
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataAvailabilityHeader        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataAvailabilityHeader          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataAvailabilityHeader = fmt.Errorf("proto: unexpected end of group")
)
