// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: command.proto

package proto

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

type CommandInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Tags   []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CommandInfo) Reset() {
	*x = CommandInfo{}
	mi := &file_command_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandInfo) ProtoMessage() {}

func (x *CommandInfo) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandInfo.ProtoReflect.Descriptor instead.
func (*CommandInfo) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

func (x *CommandInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommandInfo) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CommandInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CommandLookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *CommandLookupRequest) Reset() {
	*x = CommandLookupRequest{}
	mi := &file_command_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandLookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandLookupRequest) ProtoMessage() {}

func (x *CommandLookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandLookupRequest.ProtoReflect.Descriptor instead.
func (*CommandLookupRequest) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{1}
}

func (x *CommandLookupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommandLookupRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type AddCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *AddCommandResponse) Reset() {
	*x = AddCommandResponse{}
	mi := &file_command_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommandResponse) ProtoMessage() {}

func (x *AddCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommandResponse.ProtoReflect.Descriptor instead.
func (*AddCommandResponse) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{2}
}

func (x *AddCommandResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type RemoveCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *RemoveCommandResponse) Reset() {
	*x = RemoveCommandResponse{}
	mi := &file_command_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCommandResponse) ProtoMessage() {}

func (x *RemoveCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCommandResponse.ProtoReflect.Descriptor instead.
func (*RemoveCommandResponse) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveCommandResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type CommandArgument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
}

func (x *CommandArgument) Reset() {
	*x = CommandArgument{}
	mi := &file_command_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandArgument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandArgument) ProtoMessage() {}

func (x *CommandArgument) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandArgument.ProtoReflect.Descriptor instead.
func (*CommandArgument) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{4}
}

func (x *CommandArgument) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CommandArgument) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CommandArgument) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CommandReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDelivered bool   `protobuf:"varint,1,opt,name=is_delivered,json=isDelivered,proto3" json:"is_delivered,omitempty"`
	Status      int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Payload     []byte `protobuf:"bytes,4,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
}

func (x *CommandReturn) Reset() {
	*x = CommandReturn{}
	mi := &file_command_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandReturn) ProtoMessage() {}

func (x *CommandReturn) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandReturn.ProtoReflect.Descriptor instead.
func (*CommandReturn) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{5}
}

func (x *CommandReturn) GetIsDelivered() bool {
	if x != nil {
		return x.IsDelivered
	}
	return false
}

func (x *CommandReturn) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CommandReturn) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *CommandReturn) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x22, 0x3e, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x22, 0x33, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x6e,
	0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x98,
	0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xa8, 0x02, 0x0a, 0x11, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x3d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x11, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_command_proto_goTypes = []any{
	(*CommandInfo)(nil),           // 0: proto.CommandInfo
	(*CommandLookupRequest)(nil),  // 1: proto.CommandLookupRequest
	(*AddCommandResponse)(nil),    // 2: proto.AddCommandResponse
	(*RemoveCommandResponse)(nil), // 3: proto.RemoveCommandResponse
	(*CommandArgument)(nil),       // 4: proto.CommandArgument
	(*CommandReturn)(nil),         // 5: proto.CommandReturn
}
var file_command_proto_depIdxs = []int32{
	0, // 0: proto.CommandController.AddCommand:input_type -> proto.CommandInfo
	1, // 1: proto.CommandController.RemoveCommand:input_type -> proto.CommandLookupRequest
	4, // 2: proto.CommandController.SendCommand:input_type -> proto.CommandArgument
	4, // 3: proto.CommandController.SendStreamCommand:input_type -> proto.CommandArgument
	2, // 4: proto.CommandController.AddCommand:output_type -> proto.AddCommandResponse
	3, // 5: proto.CommandController.RemoveCommand:output_type -> proto.RemoveCommandResponse
	5, // 6: proto.CommandController.SendCommand:output_type -> proto.CommandReturn
	5, // 7: proto.CommandController.SendStreamCommand:output_type -> proto.CommandReturn
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	file_command_proto_msgTypes[4].OneofWrappers = []any{}
	file_command_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		MessageInfos:      file_command_proto_msgTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}
