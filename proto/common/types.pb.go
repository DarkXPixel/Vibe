// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: common/types.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserId) Reset() {
	*x = UserId{}
	mi := &file_common_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_common_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_common_types_proto_rawDescGZIP(), []int{0}
}

func (x *UserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Timestamp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Unix          int64                  `protobuf:"varint,1,opt,name=unix,proto3" json:"unix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	mi := &file_common_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_common_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_common_types_proto_rawDescGZIP(), []int{1}
}

func (x *Timestamp) GetUnix() int64 {
	if x != nil {
		return x.Unix
	}
	return 0
}

var File_common_types_proto protoreflect.FileDescriptor

const file_common_types_proto_rawDesc = "" +
	"\n" +
	"\x12common/types.proto\x12\x06common\"\x18\n" +
	"\x06UserId\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x1f\n" +
	"\tTimestamp\x12\x12\n" +
	"\x04unix\x18\x01 \x01(\x03R\x04unixB)Z'github.com/DarkXPixel/Vibe/proto/commonb\x06proto3"

var (
	file_common_types_proto_rawDescOnce sync.Once
	file_common_types_proto_rawDescData []byte
)

func file_common_types_proto_rawDescGZIP() []byte {
	file_common_types_proto_rawDescOnce.Do(func() {
		file_common_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_types_proto_rawDesc), len(file_common_types_proto_rawDesc)))
	})
	return file_common_types_proto_rawDescData
}

var file_common_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_types_proto_goTypes = []any{
	(*UserId)(nil),    // 0: common.UserId
	(*Timestamp)(nil), // 1: common.Timestamp
}
var file_common_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_types_proto_init() }
func file_common_types_proto_init() {
	if File_common_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_types_proto_rawDesc), len(file_common_types_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_types_proto_goTypes,
		DependencyIndexes: file_common_types_proto_depIdxs,
		MessageInfos:      file_common_types_proto_msgTypes,
	}.Build()
	File_common_types_proto = out.File
	file_common_types_proto_goTypes = nil
	file_common_types_proto_depIdxs = nil
}
