// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: services.proto

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

type ServiceInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Label         string                 `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	GrpcAddr      string                 `protobuf:"bytes,4,opt,name=grpc_addr,json=grpcAddr,proto3" json:"grpc_addr,omitempty"`
	HttpAddr      *string                `protobuf:"bytes,5,opt,name=http_addr,json=httpAddr,proto3,oneof" json:"http_addr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceInfo) Reset() {
	*x = ServiceInfo{}
	mi := &file_services_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfo) ProtoMessage() {}

func (x *ServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfo.ProtoReflect.Descriptor instead.
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ServiceInfo) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ServiceInfo) GetGrpcAddr() string {
	if x != nil {
		return x.GrpcAddr
	}
	return ""
}

func (x *ServiceInfo) GetHttpAddr() string {
	if x != nil && x.HttpAddr != nil {
		return *x.HttpAddr
	}
	return ""
}

type GetServiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Type          *string                `protobuf:"bytes,2,opt,name=type,proto3,oneof" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServiceRequest) Reset() {
	*x = GetServiceRequest{}
	mi := &file_services_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceRequest) ProtoMessage() {}

func (x *GetServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceRequest.ProtoReflect.Descriptor instead.
func (*GetServiceRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{1}
}

func (x *GetServiceRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *GetServiceRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type GetServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *ServiceInfo           `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServiceResponse) Reset() {
	*x = GetServiceResponse{}
	mi := &file_services_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceResponse) ProtoMessage() {}

func (x *GetServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceResponse.ProtoReflect.Descriptor instead.
func (*GetServiceResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{2}
}

func (x *GetServiceResponse) GetData() *ServiceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListServiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          *string                `protobuf:"bytes,1,opt,name=type,proto3,oneof" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServiceRequest) Reset() {
	*x = ListServiceRequest{}
	mi := &file_services_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceRequest) ProtoMessage() {}

func (x *ListServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceRequest.ProtoReflect.Descriptor instead.
func (*ListServiceRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{3}
}

func (x *ListServiceRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type ListServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []*ServiceInfo         `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServiceResponse) Reset() {
	*x = ListServiceResponse{}
	mi := &file_services_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceResponse) ProtoMessage() {}

func (x *ListServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceResponse.ProtoReflect.Descriptor instead.
func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{4}
}

func (x *ListServiceResponse) GetData() []*ServiceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddServiceResponse) Reset() {
	*x = AddServiceResponse{}
	mi := &file_services_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddServiceResponse) ProtoMessage() {}

func (x *AddServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddServiceResponse.ProtoReflect.Descriptor instead.
func (*AddServiceResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{5}
}

func (x *AddServiceResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type RemoveServiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveServiceRequest) Reset() {
	*x = RemoveServiceRequest{}
	mi := &file_services_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveServiceRequest) ProtoMessage() {}

func (x *RemoveServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveServiceRequest.ProtoReflect.Descriptor instead.
func (*RemoveServiceRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveServiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveServiceResponse) Reset() {
	*x = RemoveServiceResponse{}
	mi := &file_services_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveServiceResponse) ProtoMessage() {}

func (x *RemoveServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveServiceResponse.ProtoReflect.Descriptor instead.
func (*RemoveServiceResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveServiceResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type EventInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Event         string                 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventInfo) Reset() {
	*x = EventInfo{}
	mi := &file_services_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInfo) ProtoMessage() {}

func (x *EventInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventInfo.ProtoReflect.Descriptor instead.
func (*EventInfo) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{8}
}

func (x *EventInfo) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *EventInfo) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	mi := &file_services_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{9}
}

var File_services_proto protoreflect.FileDescriptor

const file_services_proto_rawDesc = "" +
	"\n" +
	"\x0eservices.proto\x12\x05proto\"\x94\x01\n" +
	"\vServiceInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x14\n" +
	"\x05label\x18\x03 \x01(\tR\x05label\x12\x1b\n" +
	"\tgrpc_addr\x18\x04 \x01(\tR\bgrpcAddr\x12 \n" +
	"\thttp_addr\x18\x05 \x01(\tH\x00R\bhttpAddr\x88\x01\x01B\f\n" +
	"\n" +
	"_http_addr\"Q\n" +
	"\x11GetServiceRequest\x12\x13\n" +
	"\x02id\x18\x01 \x01(\tH\x00R\x02id\x88\x01\x01\x12\x17\n" +
	"\x04type\x18\x02 \x01(\tH\x01R\x04type\x88\x01\x01B\x05\n" +
	"\x03_idB\a\n" +
	"\x05_type\"<\n" +
	"\x12GetServiceResponse\x12&\n" +
	"\x04data\x18\x01 \x01(\v2\x12.proto.ServiceInfoR\x04data\"6\n" +
	"\x12ListServiceRequest\x12\x17\n" +
	"\x04type\x18\x01 \x01(\tH\x00R\x04type\x88\x01\x01B\a\n" +
	"\x05_type\"=\n" +
	"\x13ListServiceResponse\x12&\n" +
	"\x04data\x18\x01 \x03(\v2\x12.proto.ServiceInfoR\x04data\"3\n" +
	"\x12AddServiceResponse\x12\x1d\n" +
	"\n" +
	"is_success\x18\x01 \x01(\bR\tisSuccess\"&\n" +
	"\x14RemoveServiceRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"6\n" +
	"\x15RemoveServiceResponse\x12\x1d\n" +
	"\n" +
	"is_success\x18\x01 \x01(\bR\tisSuccess\"5\n" +
	"\tEventInfo\x12\x14\n" +
	"\x05event\x18\x01 \x01(\tR\x05event\x12\x12\n" +
	"\x04data\x18\x02 \x01(\fR\x04data\"\x0f\n" +
	"\rEventResponse2\xe8\x02\n" +
	"\x10DirectoryService\x12C\n" +
	"\n" +
	"GetService\x12\x18.proto.GetServiceRequest\x1a\x19.proto.GetServiceResponse\"\x00\x12F\n" +
	"\vListService\x12\x19.proto.ListServiceRequest\x1a\x1a.proto.ListServiceResponse\"\x00\x12=\n" +
	"\n" +
	"AddService\x12\x12.proto.ServiceInfo\x1a\x19.proto.AddServiceResponse\"\x00\x12L\n" +
	"\rRemoveService\x12\x1b.proto.RemoveServiceRequest\x1a\x1c.proto.RemoveServiceResponse\"\x00\x12:\n" +
	"\x0eBroadcastEvent\x12\x10.proto.EventInfo\x1a\x14.proto.EventResponse\"\x00B\tZ\a.;protob\x06proto3"

var (
	file_services_proto_rawDescOnce sync.Once
	file_services_proto_rawDescData []byte
)

func file_services_proto_rawDescGZIP() []byte {
	file_services_proto_rawDescOnce.Do(func() {
		file_services_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_proto_rawDesc), len(file_services_proto_rawDesc)))
	})
	return file_services_proto_rawDescData
}

var file_services_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_proto_goTypes = []any{
	(*ServiceInfo)(nil),           // 0: proto.ServiceInfo
	(*GetServiceRequest)(nil),     // 1: proto.GetServiceRequest
	(*GetServiceResponse)(nil),    // 2: proto.GetServiceResponse
	(*ListServiceRequest)(nil),    // 3: proto.ListServiceRequest
	(*ListServiceResponse)(nil),   // 4: proto.ListServiceResponse
	(*AddServiceResponse)(nil),    // 5: proto.AddServiceResponse
	(*RemoveServiceRequest)(nil),  // 6: proto.RemoveServiceRequest
	(*RemoveServiceResponse)(nil), // 7: proto.RemoveServiceResponse
	(*EventInfo)(nil),             // 8: proto.EventInfo
	(*EventResponse)(nil),         // 9: proto.EventResponse
}
var file_services_proto_depIdxs = []int32{
	0, // 0: proto.GetServiceResponse.data:type_name -> proto.ServiceInfo
	0, // 1: proto.ListServiceResponse.data:type_name -> proto.ServiceInfo
	1, // 2: proto.DirectoryService.GetService:input_type -> proto.GetServiceRequest
	3, // 3: proto.DirectoryService.ListService:input_type -> proto.ListServiceRequest
	0, // 4: proto.DirectoryService.AddService:input_type -> proto.ServiceInfo
	6, // 5: proto.DirectoryService.RemoveService:input_type -> proto.RemoveServiceRequest
	8, // 6: proto.DirectoryService.BroadcastEvent:input_type -> proto.EventInfo
	2, // 7: proto.DirectoryService.GetService:output_type -> proto.GetServiceResponse
	4, // 8: proto.DirectoryService.ListService:output_type -> proto.ListServiceResponse
	5, // 9: proto.DirectoryService.AddService:output_type -> proto.AddServiceResponse
	7, // 10: proto.DirectoryService.RemoveService:output_type -> proto.RemoveServiceResponse
	9, // 11: proto.DirectoryService.BroadcastEvent:output_type -> proto.EventResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	file_services_proto_msgTypes[0].OneofWrappers = []any{}
	file_services_proto_msgTypes[1].OneofWrappers = []any{}
	file_services_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_proto_rawDesc), len(file_services_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
		MessageInfos:      file_services_proto_msgTypes,
	}.Build()
	File_services_proto = out.File
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
