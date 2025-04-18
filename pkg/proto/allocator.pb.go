// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: allocator.proto

package proto

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

type AllocMqRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocMqRequest) Reset() {
	*x = AllocMqRequest{}
	mi := &file_allocator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocMqRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocMqRequest) ProtoMessage() {}

func (x *AllocMqRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allocator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocMqRequest.ProtoReflect.Descriptor instead.
func (*AllocMqRequest) Descriptor() ([]byte, []int) {
	return file_allocator_proto_rawDescGZIP(), []int{0}
}

type AllocMqResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocMqResponse) Reset() {
	*x = AllocMqResponse{}
	mi := &file_allocator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocMqResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocMqResponse) ProtoMessage() {}

func (x *AllocMqResponse) ProtoReflect() protoreflect.Message {
	mi := &file_allocator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocMqResponse.ProtoReflect.Descriptor instead.
func (*AllocMqResponse) Descriptor() ([]byte, []int) {
	return file_allocator_proto_rawDescGZIP(), []int{1}
}

func (x *AllocMqResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *AllocMqResponse) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type AllocKvRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocKvRequest) Reset() {
	*x = AllocKvRequest{}
	mi := &file_allocator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocKvRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocKvRequest) ProtoMessage() {}

func (x *AllocKvRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allocator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocKvRequest.ProtoReflect.Descriptor instead.
func (*AllocKvRequest) Descriptor() ([]byte, []int) {
	return file_allocator_proto_rawDescGZIP(), []int{2}
}

type AllocKvResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Endpoints     []string               `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocKvResponse) Reset() {
	*x = AllocKvResponse{}
	mi := &file_allocator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocKvResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocKvResponse) ProtoMessage() {}

func (x *AllocKvResponse) ProtoReflect() protoreflect.Message {
	mi := &file_allocator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocKvResponse.ProtoReflect.Descriptor instead.
func (*AllocKvResponse) Descriptor() ([]byte, []int) {
	return file_allocator_proto_rawDescGZIP(), []int{3}
}

func (x *AllocKvResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *AllocKvResponse) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type AllocCacheRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Db            int32                  `protobuf:"varint,1,opt,name=db,proto3" json:"db,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocCacheRequest) Reset() {
	*x = AllocCacheRequest{}
	mi := &file_allocator_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocCacheRequest) ProtoMessage() {}

func (x *AllocCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allocator_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocCacheRequest.ProtoReflect.Descriptor instead.
func (*AllocCacheRequest) Descriptor() ([]byte, []int) {
	return file_allocator_proto_rawDescGZIP(), []int{4}
}

func (x *AllocCacheRequest) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

type AllocCacheResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Db            int32                  `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocCacheResponse) Reset() {
	*x = AllocCacheResponse{}
	mi := &file_allocator_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocCacheResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocCacheResponse) ProtoMessage() {}

func (x *AllocCacheResponse) ProtoReflect() protoreflect.Message {
	mi := &file_allocator_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocCacheResponse.ProtoReflect.Descriptor instead.
func (*AllocCacheResponse) Descriptor() ([]byte, []int) {
	return file_allocator_proto_rawDescGZIP(), []int{5}
}

func (x *AllocCacheResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *AllocCacheResponse) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *AllocCacheResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AllocCacheResponse) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

var File_allocator_proto protoreflect.FileDescriptor

const file_allocator_proto_rawDesc = "" +
	"\n" +
	"\x0fallocator.proto\x12\x05proto\"\x10\n" +
	"\x0eAllocMqRequest\"D\n" +
	"\x0fAllocMqResponse\x12\x1d\n" +
	"\n" +
	"is_success\x18\x01 \x01(\bR\tisSuccess\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\"\x10\n" +
	"\x0eAllocKvRequest\"N\n" +
	"\x0fAllocKvResponse\x12\x1d\n" +
	"\n" +
	"is_success\x18\x01 \x01(\bR\tisSuccess\x12\x1c\n" +
	"\tendpoints\x18\x02 \x03(\tR\tendpoints\"#\n" +
	"\x11AllocCacheRequest\x12\x0e\n" +
	"\x02db\x18\x01 \x01(\x05R\x02db\"s\n" +
	"\x12AllocCacheResponse\x12\x1d\n" +
	"\n" +
	"is_success\x18\x01 \x01(\bR\tisSuccess\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x0e\n" +
	"\x02db\x18\x04 \x01(\x05R\x02db2\xd9\x01\n" +
	"\x10AllocatorService\x12D\n" +
	"\x11AllocMessageQueue\x12\x15.proto.AllocMqRequest\x1a\x16.proto.AllocMqResponse\"\x00\x12:\n" +
	"\aAllocKv\x12\x15.proto.AllocKvRequest\x1a\x16.proto.AllocKvResponse\"\x00\x12C\n" +
	"\n" +
	"AllocCache\x12\x18.proto.AllocCacheRequest\x1a\x19.proto.AllocCacheResponse\"\x00B\tZ\a.;protob\x06proto3"

var (
	file_allocator_proto_rawDescOnce sync.Once
	file_allocator_proto_rawDescData []byte
)

func file_allocator_proto_rawDescGZIP() []byte {
	file_allocator_proto_rawDescOnce.Do(func() {
		file_allocator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_allocator_proto_rawDesc), len(file_allocator_proto_rawDesc)))
	})
	return file_allocator_proto_rawDescData
}

var file_allocator_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_allocator_proto_goTypes = []any{
	(*AllocMqRequest)(nil),     // 0: proto.AllocMqRequest
	(*AllocMqResponse)(nil),    // 1: proto.AllocMqResponse
	(*AllocKvRequest)(nil),     // 2: proto.AllocKvRequest
	(*AllocKvResponse)(nil),    // 3: proto.AllocKvResponse
	(*AllocCacheRequest)(nil),  // 4: proto.AllocCacheRequest
	(*AllocCacheResponse)(nil), // 5: proto.AllocCacheResponse
}
var file_allocator_proto_depIdxs = []int32{
	0, // 0: proto.AllocatorService.AllocMessageQueue:input_type -> proto.AllocMqRequest
	2, // 1: proto.AllocatorService.AllocKv:input_type -> proto.AllocKvRequest
	4, // 2: proto.AllocatorService.AllocCache:input_type -> proto.AllocCacheRequest
	1, // 3: proto.AllocatorService.AllocMessageQueue:output_type -> proto.AllocMqResponse
	3, // 4: proto.AllocatorService.AllocKv:output_type -> proto.AllocKvResponse
	5, // 5: proto.AllocatorService.AllocCache:output_type -> proto.AllocCacheResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_allocator_proto_init() }
func file_allocator_proto_init() {
	if File_allocator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_allocator_proto_rawDesc), len(file_allocator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_allocator_proto_goTypes,
		DependencyIndexes: file_allocator_proto_depIdxs,
		MessageInfos:      file_allocator_proto_msgTypes,
	}.Build()
	File_allocator_proto = out.File
	file_allocator_proto_goTypes = nil
	file_allocator_proto_depIdxs = nil
}
