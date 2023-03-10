// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: moon/ibank/transaction.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type TxnStatus int32

const (
	TxnStatus_TXN_PENDING TxnStatus = 0
	TxnStatus_TXN_SENT    TxnStatus = 1
	TxnStatus_TXN_EXPIRED TxnStatus = 2
)

var TxnStatus_name = map[int32]string{
	0: "TXN_PENDING",
	1: "TXN_SENT",
	2: "TXN_EXPIRED",
}

var TxnStatus_value = map[string]int32{
	"TXN_PENDING": 0,
	"TXN_SENT":    1,
	"TXN_EXPIRED": 2,
}

func (x TxnStatus) String() string {
	return proto.EnumName(TxnStatus_name, int32(x))
}

func (TxnStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d43bfa2a2dd3ead2, []int{0}
}

type Transaction struct {
	Id       uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender   string    `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	SentAt   time.Time `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3,stdtime" json:"sent_at"`
	Receiver string    `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// If sent_at is equal to received_at, transaction have not been performed
	ReceivedAt time.Time    `protobuf:"bytes,5,opt,name=received_at,json=receivedAt,proto3,stdtime" json:"received_at"`
	Coins      []types.Coin `protobuf:"bytes,6,rep,name=coins,proto3" json:"coins"`
	Status     TxnStatus    `protobuf:"varint,7,opt,name=status,proto3,enum=moon.ibank.TxnStatus" json:"status,omitempty"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d43bfa2a2dd3ead2, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Transaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Transaction) GetSentAt() time.Time {
	if m != nil {
		return m.SentAt
	}
	return time.Time{}
}

func (m *Transaction) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Transaction) GetReceivedAt() time.Time {
	if m != nil {
		return m.ReceivedAt
	}
	return time.Time{}
}

func (m *Transaction) GetCoins() []types.Coin {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Transaction) GetStatus() TxnStatus {
	if m != nil {
		return m.Status
	}
	return TxnStatus_TXN_PENDING
}

func init() {
	proto.RegisterEnum("moon.ibank.TxnStatus", TxnStatus_name, TxnStatus_value)
	proto.RegisterType((*Transaction)(nil), "moon.ibank.Transaction")
}

func init() { proto.RegisterFile("moon/ibank/transaction.proto", fileDescriptor_d43bfa2a2dd3ead2) }

var fileDescriptor_d43bfa2a2dd3ead2 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x8e, 0xb3, 0x2e, 0xeb, 0x1c, 0x34, 0x2a, 0x0b, 0x50, 0x88, 0x50, 0x1a, 0x71, 0x8a, 0x10,
	0xd8, 0x5a, 0x11, 0x27, 0xc4, 0x61, 0x65, 0x11, 0xda, 0x25, 0x9a, 0xb2, 0x1c, 0x26, 0x2e, 0x93,
	0x93, 0x98, 0xc8, 0x82, 0xd8, 0x55, 0xec, 0x4d, 0xe5, 0x5f, 0xf4, 0x0f, 0x71, 0xef, 0xb1, 0x47,
	0x4e, 0x80, 0xda, 0x3f, 0x82, 0xe2, 0x24, 0xed, 0x79, 0xb7, 0xf7, 0xde, 0xf7, 0xbd, 0xef, 0x93,
	0x3f, 0x3f, 0xf8, 0xaa, 0x96, 0x52, 0x10, 0x9e, 0x53, 0xf1, 0x9d, 0xe8, 0x86, 0x0a, 0x45, 0x0b,
	0xcd, 0xa5, 0xc0, 0x8b, 0x46, 0x6a, 0x89, 0x60, 0x8b, 0x62, 0x83, 0xfa, 0xcf, 0x2a, 0x59, 0x49,
	0x33, 0x26, 0x6d, 0xd5, 0x31, 0xfc, 0xa0, 0x90, 0xaa, 0x96, 0x8a, 0xe4, 0x54, 0x31, 0xf2, 0x70,
	0x9e, 0x33, 0x4d, 0xcf, 0x49, 0x21, 0x79, 0xaf, 0xe0, 0x4f, 0x2b, 0x29, 0xab, 0x1f, 0x8c, 0x98,
	0x2e, 0xbf, 0xff, 0x46, 0x34, 0xaf, 0x99, 0xd2, 0xb4, 0x5e, 0x74, 0x84, 0xd7, 0xbf, 0x6c, 0xe8,
	0x66, 0x07, 0x63, 0x74, 0x06, 0x6d, 0x5e, 0x7a, 0x20, 0x04, 0xd1, 0x28, 0xb5, 0x79, 0x89, 0x5e,
	0x40, 0x47, 0x31, 0x51, 0xb2, 0xc6, 0xb3, 0x43, 0x10, 0x9d, 0xa6, 0x7d, 0x87, 0x3e, 0xc1, 0x13,
	0xc5, 0x84, 0xbe, 0xa3, 0xda, 0x3b, 0x0a, 0x41, 0xe4, 0xce, 0x7c, 0xdc, 0x59, 0xe1, 0xc1, 0x0a,
	0x67, 0x83, 0xd5, 0x7c, 0xbc, 0xfe, 0x33, 0xb5, 0x56, 0x7f, 0xa7, 0xc0, 0xac, 0xeb, 0x0b, 0x8d,
	0x7c, 0x38, 0x6e, 0x58, 0xc1, 0xf8, 0x03, 0x6b, 0xbc, 0x91, 0x11, 0xde, 0xf7, 0x28, 0x86, 0x6e,
	0x5f, 0x97, 0xad, 0xfc, 0xf1, 0x23, 0xe4, 0xe1, 0xb0, 0x78, 0xa1, 0xd1, 0x07, 0x78, 0xdc, 0x06,
	0xa1, 0x3c, 0x27, 0x3c, 0x8a, 0xdc, 0xd9, 0x4b, 0xdc, 0x45, 0x85, 0xdb, 0xa8, 0x70, 0x1f, 0x15,
	0xfe, 0x2c, 0xb9, 0x98, 0x8f, 0xda, 0xfd, 0xb4, 0x63, 0xa3, 0x77, 0xd0, 0x51, 0x9a, 0xea, 0x7b,
	0xe5, 0x9d, 0x84, 0x20, 0x3a, 0x9b, 0x3d, 0xc7, 0x87, 0x4f, 0xc0, 0xd9, 0x52, 0xdc, 0x18, 0x30,
	0xed, 0x49, 0x6f, 0x3e, 0xc2, 0xd3, 0xfd, 0x10, 0x3d, 0x85, 0x6e, 0x76, 0x9b, 0xdc, 0x5d, 0xc7,
	0xc9, 0xe5, 0x55, 0xf2, 0x65, 0x62, 0xa1, 0x27, 0x70, 0xdc, 0x0e, 0x6e, 0xe2, 0x24, 0x9b, 0x80,
	0x01, 0x8e, 0x6f, 0xaf, 0xaf, 0xd2, 0xf8, 0x72, 0x62, 0xcf, 0xdf, 0xae, 0xb7, 0x01, 0xd8, 0x6c,
	0x03, 0xf0, 0x6f, 0x1b, 0x80, 0xd5, 0x2e, 0xb0, 0x36, 0xbb, 0xc0, 0xfa, 0xbd, 0x0b, 0xac, 0xaf,
	0xc8, 0xdc, 0xc5, 0x72, 0xb8, 0x8c, 0x9f, 0x0b, 0xa6, 0x72, 0xc7, 0x3c, 0xfd, 0xfd, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x17, 0x87, 0x40, 0x4c, 0x34, 0x02, 0x00, 0x00,
}

func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTransaction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ReceivedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ReceivedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTransaction(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x22
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.SentAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.SentAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTransaction(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTransaction(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.SentAt)
	n += 1 + l + sovTransaction(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ReceivedAt)
	n += 1 + l + sovTransaction(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTransaction(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTransaction(uint64(m.Status))
	}
	return n
}

func sovTransaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.SentAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ReceivedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= TxnStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransaction
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
func skipTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
				return 0, ErrInvalidLengthTransaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransaction = fmt.Errorf("proto: unexpected end of group")
)
