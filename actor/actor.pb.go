// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: actor.proto

package actor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//user messages
type States int32

const (
	States_Empty      States = 0
	States_Restarting States = 1
	States_Stopping   States = 2
	States_Stopped    States = 3
	States_PoisonPill States = 4
	States_Started    States = 5
)

// Enum value maps for States.
var (
	States_name = map[int32]string{
		0: "Empty",
		1: "Restarting",
		2: "Stopping",
		3: "Stopped",
		4: "PoisonPill",
		5: "Started",
	}
	States_value = map[string]int32{
		"Empty":      0,
		"Restarting": 1,
		"Stopping":   2,
		"Stopped":    3,
		"PoisonPill": 4,
		"Started":    5,
	}
)

func (x States) Enum() *States {
	p := new(States)
	*p = x
	return p
}

func (x States) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (States) Descriptor() protoreflect.EnumDescriptor {
	return file_actor_proto_enumTypes[0].Descriptor()
}

func (States) Type() protoreflect.EnumType {
	return &file_actor_proto_enumTypes[0]
}

func (x States) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use States.Descriptor instead.
func (States) EnumDescriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{0}
}

type PID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *PID) Reset() {
	*x = PID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PID) ProtoMessage() {}

func (x *PID) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PID.ProtoReflect.Descriptor instead.
func (*PID) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{0}
}

func (x *PID) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_actor_proto protoreflect.FileDescriptor

var file_actor_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x22, 0x29, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x2a,
	0x5b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x50, 0x6f, 0x69, 0x73, 0x6f, 0x6e, 0x50, 0x69, 0x6c, 0x6c, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x05, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2e, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_actor_proto_rawDescOnce sync.Once
	file_actor_proto_rawDescData = file_actor_proto_rawDesc
)

func file_actor_proto_rawDescGZIP() []byte {
	file_actor_proto_rawDescOnce.Do(func() {
		file_actor_proto_rawDescData = protoimpl.X.CompressGZIP(file_actor_proto_rawDescData)
	})
	return file_actor_proto_rawDescData
}

var file_actor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_actor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_actor_proto_goTypes = []interface{}{
	(States)(0), // 0: actor.States
	(*PID)(nil), // 1: actor.PID
}
var file_actor_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_actor_proto_init() }
func file_actor_proto_init() {
	if File_actor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_actor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_actor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_actor_proto_goTypes,
		DependencyIndexes: file_actor_proto_depIdxs,
		EnumInfos:         file_actor_proto_enumTypes,
		MessageInfos:      file_actor_proto_msgTypes,
	}.Build()
	File_actor_proto = out.File
	file_actor_proto_rawDesc = nil
	file_actor_proto_goTypes = nil
	file_actor_proto_depIdxs = nil
}
