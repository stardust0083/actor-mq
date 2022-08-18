// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: actor.proto

package actor

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

//user messages
type States int32

const (
	Empty      States = 0
	Restarting States = 1
	Stopping   States = 2
	Stopped    States = 3
	PoisonPill States = 4
	Started    States = 5
)

var States_name = map[int32]string{
	0: "Empty",
	1: "Restarting",
	2: "Stopping",
	3: "Stopped",
	4: "PoisonPill",
	5: "Started",
}

var States_value = map[string]int32{
	"Empty":      0,
	"Restarting": 1,
	"Stopping":   2,
	"Stopped":    3,
	"PoisonPill": 4,
	"Started":    5,
}

func (States) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_93a2698287ded216, []int{0}
}

type Directive int32

const (
	EmptyDirective    Directive = 0
	ResumeDirective   Directive = 1
	RestartDirective  Directive = 2
	StopDirective     Directive = 3
	EscalateDirective Directive = 4
)

var Directive_name = map[int32]string{
	0: "EmptyDirective",
	1: "ResumeDirective",
	2: "RestartDirective",
	3: "StopDirective",
	4: "EscalateDirective",
}

var Directive_value = map[string]int32{
	"EmptyDirective":    0,
	"ResumeDirective":   1,
	"RestartDirective":  2,
	"StopDirective":     3,
	"EscalateDirective": 4,
}

func (Directive) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_93a2698287ded216, []int{1}
}

type PID struct {
	Host string `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *PID) Reset()      { *m = PID{} }
func (*PID) ProtoMessage() {}
func (*PID) Descriptor() ([]byte, []int) {
	return fileDescriptor_93a2698287ded216, []int{0}
}
func (m *PID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PID.Merge(m, src)
}
func (m *PID) XXX_Size() int {
	return m.Size()
}
func (m *PID) XXX_DiscardUnknown() {
	xxx_messageInfo_PID.DiscardUnknown(m)
}

var xxx_messageInfo_PID proto.InternalMessageInfo

func (m *PID) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StateMsg struct {
	State States `protobuf:"varint,1,opt,name=State,proto3,enum=actor.States" json:"State,omitempty"`
}

func (m *StateMsg) Reset()      { *m = StateMsg{} }
func (*StateMsg) ProtoMessage() {}
func (*StateMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_93a2698287ded216, []int{1}
}
func (m *StateMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateMsg.Merge(m, src)
}
func (m *StateMsg) XXX_Size() int {
	return m.Size()
}
func (m *StateMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StateMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StateMsg proto.InternalMessageInfo

func (m *StateMsg) GetState() States {
	if m != nil {
		return m.State
	}
	return Empty
}

func init() {
	proto.RegisterEnum("actor.States", States_name, States_value)
	proto.RegisterEnum("actor.Directive", Directive_name, Directive_value)
	proto.RegisterType((*PID)(nil), "actor.PID")
	proto.RegisterType((*StateMsg)(nil), "actor.StateMsg")
}

func init() { proto.RegisterFile("actor.proto", fileDescriptor_93a2698287ded216) }

var fileDescriptor_93a2698287ded216 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4e, 0x02, 0x41,
	0x18, 0x85, 0x67, 0x16, 0x16, 0xe5, 0x47, 0xd6, 0xe1, 0x57, 0x13, 0xaa, 0x3f, 0x06, 0x1b, 0xa5,
	0xc0, 0x44, 0x3d, 0x81, 0x81, 0x44, 0x0a, 0x13, 0xb2, 0x96, 0x56, 0x23, 0x3b, 0x21, 0x9b, 0x00,
	0xb3, 0xd9, 0x19, 0x4d, 0xec, 0x3c, 0x82, 0xc7, 0xf0, 0x28, 0x96, 0x94, 0x94, 0x32, 0x34, 0x96,
	0x1c, 0xc1, 0xec, 0xac, 0x71, 0xed, 0xfe, 0xf7, 0xbd, 0x79, 0xef, 0x25, 0x03, 0x2d, 0x39, 0xb5,
	0x3a, 0x1f, 0x64, 0xb9, 0xb6, 0x1a, 0x43, 0x2f, 0x7a, 0x17, 0x50, 0x9b, 0x8c, 0x87, 0x88, 0x50,
	0xbf, 0xd3, 0xc6, 0x76, 0xf9, 0x29, 0x3f, 0x6f, 0xc6, 0xfe, 0xc6, 0x08, 0x82, 0x71, 0xd2, 0x0d,
	0x3c, 0x09, 0xc6, 0x49, 0xef, 0x12, 0xf6, 0x1f, 0xac, 0xb4, 0xea, 0xde, 0xcc, 0xf0, 0x0c, 0x42,
	0x7f, 0xfb, 0x40, 0x74, 0xd5, 0x1e, 0x94, 0xd5, 0x9e, 0x99, 0xb8, 0xf4, 0xfa, 0x8f, 0xd0, 0x28,
	0x01, 0x36, 0x21, 0x1c, 0x2d, 0x32, 0xfb, 0x2a, 0x18, 0x46, 0x00, 0xb1, 0x32, 0x56, 0xe6, 0x36,
	0x5d, 0xce, 0x04, 0xc7, 0x83, 0xa2, 0x55, 0x67, 0x59, 0xa1, 0x02, 0x6c, 0xc1, 0x9e, 0x57, 0x2a,
	0x11, 0xb5, 0xe2, 0xe9, 0x44, 0xa7, 0x46, 0x2f, 0x27, 0xe9, 0x7c, 0x2e, 0xea, 0xa5, 0x29, 0x73,
	0xab, 0x12, 0x11, 0xf6, 0x2d, 0x34, 0x87, 0x69, 0xae, 0xa6, 0x36, 0x7d, 0x51, 0x88, 0x10, 0xf9,
	0xfe, 0x3f, 0x22, 0x18, 0x1e, 0xc1, 0x61, 0xac, 0xcc, 0xf3, 0x42, 0x55, 0x90, 0xe3, 0x31, 0x88,
	0xdf, 0xf5, 0x8a, 0x06, 0xd8, 0x81, 0x76, 0xb1, 0x5a, 0xa1, 0x1a, 0x9e, 0x40, 0x67, 0x64, 0xa6,
	0x72, 0x2e, 0xed, 0xbf, 0x7c, 0xfd, 0xf6, 0x66, 0xb5, 0x21, 0xb6, 0xde, 0x10, 0xdb, 0x6d, 0x88,
	0xbf, 0x39, 0xe2, 0x1f, 0x8e, 0xf8, 0xa7, 0x23, 0xbe, 0x72, 0xc4, 0xbf, 0x1c, 0xf1, 0x6f, 0x47,
	0x6c, 0xe7, 0x88, 0xbf, 0x6f, 0x89, 0xad, 0xb6, 0xc4, 0xd6, 0x5b, 0x62, 0x4f, 0x0d, 0xff, 0xe5,
	0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0x7b, 0xd8, 0xc3, 0x81, 0x01, 0x00, 0x00,
}

func (x States) String() string {
	s, ok := States_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x Directive) String() string {
	s, ok := Directive_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *PID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PID)
	if !ok {
		that2, ok := that.(PID)
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
	if this.Host != that1.Host {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	return true
}
func (this *StateMsg) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StateMsg)
	if !ok {
		that2, ok := that.(StateMsg)
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
	if this.State != that1.State {
		return false
	}
	return true
}
func (this *PID) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&actor.PID{")
	s = append(s, "Host: "+fmt.Sprintf("%#v", this.Host)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StateMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&actor.StateMsg{")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringActor(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintActor(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintActor(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StateMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintActor(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintActor(dAtA []byte, offset int, v uint64) int {
	offset -= sovActor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovActor(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovActor(uint64(l))
	}
	return n
}

func (m *StateMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovActor(uint64(m.State))
	}
	return n
}

func sovActor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActor(x uint64) (n int) {
	return sovActor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PID) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PID{`,
		`Host:` + fmt.Sprintf("%v", this.Host) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StateMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StateMsg{`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringActor(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActor
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
			return fmt.Errorf("proto: PID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActor
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
				return ErrInvalidLengthActor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActor
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
				return ErrInvalidLengthActor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActor
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
func (m *StateMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActor
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
			return fmt.Errorf("proto: StateMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= States(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActor
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
func skipActor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActor
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
					return 0, ErrIntOverflowActor
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
					return 0, ErrIntOverflowActor
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
				return 0, ErrInvalidLengthActor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActor = fmt.Errorf("proto: unexpected end of group")
)
