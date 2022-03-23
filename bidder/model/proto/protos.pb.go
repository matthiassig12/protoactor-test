// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model/proto/protos.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

type Auction_State int32

const (
	IDLE    Auction_State = 0
	RUNNING Auction_State = 1
	CLOSED  Auction_State = 2
)

var Auction_State_name = map[int32]string{
	0: "IDLE",
	1: "RUNNING",
	2: "CLOSED",
}

var Auction_State_value = map[string]int32{
	"IDLE":    0,
	"RUNNING": 1,
	"CLOSED":  2,
}

func (Auction_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a77cee963bfd712, []int{0, 0}
}

type Auction struct {
	Id    string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State Auction_State `protobuf:"varint,2,opt,name=state,proto3,enum=model.Auction_State" json:"state,omitempty"`
	Lots  []*Lot        `protobuf:"bytes,3,rep,name=lots,proto3" json:"lots,omitempty"`
}

func (m *Auction) Reset()      { *m = Auction{} }
func (*Auction) ProtoMessage() {}
func (*Auction) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a77cee963bfd712, []int{0}
}
func (m *Auction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auction.Merge(m, src)
}
func (m *Auction) XXX_Size() int {
	return m.Size()
}
func (m *Auction) XXX_DiscardUnknown() {
	xxx_messageInfo_Auction.DiscardUnknown(m)
}

var xxx_messageInfo_Auction proto.InternalMessageInfo

func (m *Auction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Auction) GetState() Auction_State {
	if m != nil {
		return m.State
	}
	return IDLE
}

func (m *Auction) GetLots() []*Lot {
	if m != nil {
		return m.Lots
	}
	return nil
}

type Lot struct {
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number     int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	CurrentBid int64  `protobuf:"varint,3,opt,name=current_bid,json=currentBid,proto3" json:"current_bid,omitempty"`
	Leader     string `protobuf:"bytes,4,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (m *Lot) Reset()      { *m = Lot{} }
func (*Lot) ProtoMessage() {}
func (*Lot) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a77cee963bfd712, []int{1}
}
func (m *Lot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lot.Merge(m, src)
}
func (m *Lot) XXX_Size() int {
	return m.Size()
}
func (m *Lot) XXX_DiscardUnknown() {
	xxx_messageInfo_Lot.DiscardUnknown(m)
}

var xxx_messageInfo_Lot proto.InternalMessageInfo

func (m *Lot) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Lot) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Lot) GetCurrentBid() int64 {
	if m != nil {
		return m.CurrentBid
	}
	return 0
}

func (m *Lot) GetLeader() string {
	if m != nil {
		return m.Leader
	}
	return ""
}

func init() {
	proto.RegisterEnum("model.Auction_State", Auction_State_name, Auction_State_value)
	proto.RegisterType((*Auction)(nil), "model.Auction")
	proto.RegisterType((*Lot)(nil), "model.Lot")
}

func init() { proto.RegisterFile("model/proto/protos.proto", fileDescriptor_6a77cee963bfd712) }

var fileDescriptor_6a77cee963bfd712 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0xe7, 0x4f, 0x9a, 0x56, 0xff, 0x42, 0x29, 0x83, 0xc8, 0xac, 0x7e, 0x43, 0x57, 0xa1,
	0x8b, 0x08, 0xd5, 0x0b, 0x58, 0x5b, 0xa4, 0x10, 0x2a, 0x4c, 0x71, 0x2d, 0x6d, 0x33, 0x42, 0xa0,
	0xed, 0x48, 0x3a, 0xdd, 0x7b, 0x04, 0xbd, 0x85, 0x47, 0x71, 0x99, 0x65, 0x97, 0x66, 0xb2, 0x71,
	0xd9, 0x23, 0x88, 0x93, 0xec, 0xdc, 0xcc, 0xf0, 0xde, 0xfb, 0xde, 0xf0, 0x18, 0x14, 0x5b, 0x9d,
	0xaa, 0xcd, 0xf5, 0x6b, 0xae, 0x8d, 0xae, 0xcf, 0x7d, 0xec, 0x2e, 0x1e, 0xb8, 0x64, 0xf0, 0x01,
	0xd8, 0xb9, 0x3b, 0xac, 0x4d, 0xa6, 0x77, 0xbc, 0x87, 0x5e, 0x96, 0x0a, 0x08, 0x21, 0x3a, 0x97,
	0x5e, 0x96, 0xf2, 0x21, 0x06, 0x7b, 0xb3, 0x34, 0x4a, 0x78, 0x21, 0x44, 0xbd, 0xd1, 0x45, 0xec,
	0x2a, 0x71, 0x83, 0xc7, 0x8b, 0xbf, 0x4c, 0xd6, 0x08, 0x27, 0x6c, 0x6d, 0xb4, 0xd9, 0x0b, 0x3f,
	0xf4, 0xa3, 0xee, 0x08, 0x1b, 0x34, 0xd1, 0x46, 0x3a, 0x7f, 0x30, 0xc4, 0xc0, 0xf1, 0xfc, 0x0c,
	0x5b, 0xb3, 0x49, 0x32, 0xed, 0x33, 0xde, 0xc5, 0x8e, 0x7c, 0x9a, 0xcf, 0x67, 0xf3, 0x87, 0x3e,
	0x70, 0xc4, 0xf6, 0x7d, 0xf2, 0xb8, 0x98, 0x4e, 0xfa, 0xde, 0xe0, 0x05, 0xfd, 0x44, 0x9b, 0x7f,
	0x73, 0x2e, 0xb1, 0xbd, 0x3b, 0x6c, 0x57, 0x2a, 0x77, 0x7b, 0x02, 0xd9, 0x28, 0x7e, 0x85, 0xdd,
	0xf5, 0x21, 0xcf, 0xd5, 0xce, 0x3c, 0xaf, 0xb2, 0x54, 0xf8, 0x21, 0x44, 0xbe, 0xc4, 0xc6, 0x1a,
	0xd7, 0xc5, 0x8d, 0x5a, 0xa6, 0x2a, 0x17, 0x2d, 0xf7, 0x58, 0xa3, 0xc6, 0xb7, 0x45, 0x49, 0xec,
	0x58, 0x12, 0x3b, 0x95, 0x04, 0x6f, 0x96, 0xe0, 0xd3, 0x12, 0x7c, 0x59, 0x82, 0xc2, 0x12, 0x7c,
	0x5b, 0x82, 0x1f, 0x4b, 0xec, 0x64, 0x09, 0xde, 0x2b, 0x62, 0x45, 0x45, 0xec, 0x58, 0x11, 0x5b,
	0xb5, 0xdd, 0xff, 0xdd, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xf9, 0x96, 0xce, 0x5b, 0x01,
	0x00, 0x00,
}

func (x Auction_State) String() string {
	s, ok := Auction_State_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Auction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Auction)
	if !ok {
		that2, ok := that.(Auction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if len(this.Lots) != len(that1.Lots) {
		return false
	}
	for i := range this.Lots {
		if !this.Lots[i].Equal(that1.Lots[i]) {
			return false
		}
	}
	return true
}
func (this *Lot) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Lot)
	if !ok {
		that2, ok := that.(Lot)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Number != that1.Number {
		return false
	}
	if this.CurrentBid != that1.CurrentBid {
		return false
	}
	if this.Leader != that1.Leader {
		return false
	}
	return true
}
func (this *Auction) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&model.Auction{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	if this.Lots != nil {
		s = append(s, "Lots: "+fmt.Sprintf("%#v", this.Lots)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Lot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&model.Lot{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Number: "+fmt.Sprintf("%#v", this.Number)+",\n")
	s = append(s, "CurrentBid: "+fmt.Sprintf("%#v", this.CurrentBid)+",\n")
	s = append(s, "Leader: "+fmt.Sprintf("%#v", this.Leader)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtos(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Auction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Lots) > 0 {
		for iNdEx := len(m.Lots) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Lots[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProtos(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.State != 0 {
		i = encodeVarintProtos(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Lot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Leader) > 0 {
		i -= len(m.Leader)
		copy(dAtA[i:], m.Leader)
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Leader)))
		i--
		dAtA[i] = 0x22
	}
	if m.CurrentBid != 0 {
		i = encodeVarintProtos(dAtA, i, uint64(m.CurrentBid))
		i--
		dAtA[i] = 0x18
	}
	if m.Number != 0 {
		i = encodeVarintProtos(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtos(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtos(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Auction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovProtos(uint64(m.State))
	}
	if len(m.Lots) > 0 {
		for _, e := range m.Lots {
			l = e.Size()
			n += 1 + l + sovProtos(uint64(l))
		}
	}
	return n
}

func (m *Lot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	if m.Number != 0 {
		n += 1 + sovProtos(uint64(m.Number))
	}
	if m.CurrentBid != 0 {
		n += 1 + sovProtos(uint64(m.CurrentBid))
	}
	l = len(m.Leader)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func sovProtos(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtos(x uint64) (n int) {
	return sovProtos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Auction) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForLots := "[]*Lot{"
	for _, f := range this.Lots {
		repeatedStringForLots += strings.Replace(f.String(), "Lot", "Lot", 1) + ","
	}
	repeatedStringForLots += "}"
	s := strings.Join([]string{`&Auction{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`Lots:` + repeatedStringForLots + `,`,
		`}`,
	}, "")
	return s
}
func (this *Lot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Lot{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`CurrentBid:` + fmt.Sprintf("%v", this.CurrentBid) + `,`,
		`Leader:` + fmt.Sprintf("%v", this.Leader) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProtos(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Auction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
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
			return fmt.Errorf("proto: Auction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
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
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Auction_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lots", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
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
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lots = append(m.Lots, &Lot{})
			if err := m.Lots[len(m.Lots)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtos
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
func (m *Lot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
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
			return fmt.Errorf("proto: Lot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
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
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentBid", wireType)
			}
			m.CurrentBid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentBid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
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
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtos
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
func skipProtos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtos
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
					return 0, ErrIntOverflowProtos
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
					return 0, ErrIntOverflowProtos
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
				return 0, ErrInvalidLengthProtos
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtos
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtos
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtos        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtos          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProtos = fmt.Errorf("proto: unexpected end of group")
)
