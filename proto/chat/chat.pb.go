// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: chat/chat.proto

package chat

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

type SendMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	SenderId      string                 `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	mi := &file_chat_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *SendMessageRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *SendMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GetChatMessagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChatMessagesRequest) Reset() {
	*x = GetChatMessagesRequest{}
	mi := &file_chat_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChatMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatMessagesRequest) ProtoMessage() {}

func (x *GetChatMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetChatMessagesRequest) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetChatMessagesRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *GetChatMessagesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetChatMessagesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type StreamMessagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamMessagesRequest) Reset() {
	*x = StreamMessagesRequest{}
	mi := &file_chat_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessagesRequest) ProtoMessage() {}

func (x *StreamMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessagesRequest.ProtoReflect.Descriptor instead.
func (*StreamMessagesRequest) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *StreamMessagesRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChatId        string                 `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	SenderId      string                 `protobuf:"bytes,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp     int64                  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_chat_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ChatHistory struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*Message             `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatHistory) Reset() {
	*x = ChatHistory{}
	mi := &file_chat_chat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatHistory) ProtoMessage() {}

func (x *ChatHistory) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatHistory.ProtoReflect.Descriptor instead.
func (*ChatHistory) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatHistory) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_chat_chat_proto protoreflect.FileDescriptor

const file_chat_chat_proto_rawDesc = "" +
	"\n" +
	"\x0fchat/chat.proto\x12\x04chat\"d\n" +
	"\x12SendMessageRequest\x12\x17\n" +
	"\achat_id\x18\x01 \x01(\tR\x06chatId\x12\x1b\n" +
	"\tsender_id\x18\x02 \x01(\tR\bsenderId\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\"_\n" +
	"\x16GetChatMessagesRequest\x12\x17\n" +
	"\achat_id\x18\x01 \x01(\tR\x06chatId\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x03 \x01(\x05R\x06offset\"0\n" +
	"\x15StreamMessagesRequest\x12\x17\n" +
	"\achat_id\x18\x01 \x01(\tR\x06chatId\"\x87\x01\n" +
	"\aMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\achat_id\x18\x02 \x01(\tR\x06chatId\x12\x1b\n" +
	"\tsender_id\x18\x03 \x01(\tR\bsenderId\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\x12\x1c\n" +
	"\ttimestamp\x18\x05 \x01(\x03R\ttimestamp\"8\n" +
	"\vChatHistory\x12)\n" +
	"\bmessages\x18\x01 \x03(\v2\r.chat.MessageR\bmessages2\xc9\x01\n" +
	"\vChatService\x126\n" +
	"\vSendMessage\x12\x18.chat.SendMessageRequest\x1a\r.chat.Message\x12B\n" +
	"\x0fGetChatMessages\x12\x1c.chat.GetChatMessagesRequest\x1a\x11.chat.ChatHistory\x12>\n" +
	"\x0eStreamMessages\x12\x1b.chat.StreamMessagesRequest\x1a\r.chat.Message0\x01B'Z%github.com/DarkXPixel/Vibe/proto/chatb\x06proto3"

var (
	file_chat_chat_proto_rawDescOnce sync.Once
	file_chat_chat_proto_rawDescData []byte
)

func file_chat_chat_proto_rawDescGZIP() []byte {
	file_chat_chat_proto_rawDescOnce.Do(func() {
		file_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chat_chat_proto_rawDesc), len(file_chat_chat_proto_rawDesc)))
	})
	return file_chat_chat_proto_rawDescData
}

var file_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chat_chat_proto_goTypes = []any{
	(*SendMessageRequest)(nil),     // 0: chat.SendMessageRequest
	(*GetChatMessagesRequest)(nil), // 1: chat.GetChatMessagesRequest
	(*StreamMessagesRequest)(nil),  // 2: chat.StreamMessagesRequest
	(*Message)(nil),                // 3: chat.Message
	(*ChatHistory)(nil),            // 4: chat.ChatHistory
}
var file_chat_chat_proto_depIdxs = []int32{
	3, // 0: chat.ChatHistory.messages:type_name -> chat.Message
	0, // 1: chat.ChatService.SendMessage:input_type -> chat.SendMessageRequest
	1, // 2: chat.ChatService.GetChatMessages:input_type -> chat.GetChatMessagesRequest
	2, // 3: chat.ChatService.StreamMessages:input_type -> chat.StreamMessagesRequest
	3, // 4: chat.ChatService.SendMessage:output_type -> chat.Message
	4, // 5: chat.ChatService.GetChatMessages:output_type -> chat.ChatHistory
	3, // 6: chat.ChatService.StreamMessages:output_type -> chat.Message
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_chat_proto_init() }
func file_chat_chat_proto_init() {
	if File_chat_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chat_chat_proto_rawDesc), len(file_chat_chat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_chat_proto_goTypes,
		DependencyIndexes: file_chat_chat_proto_depIdxs,
		MessageInfos:      file_chat_chat_proto_msgTypes,
	}.Build()
	File_chat_chat_proto = out.File
	file_chat_chat_proto_goTypes = nil
	file_chat_chat_proto_depIdxs = nil
}
