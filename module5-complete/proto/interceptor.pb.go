// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: proto/interceptor.proto

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

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interceptor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interceptor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_interceptor_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interceptor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interceptor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_interceptor_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LongRunningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LongRunningRequest) Reset() {
	*x = LongRunningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interceptor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongRunningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongRunningRequest) ProtoMessage() {}

func (x *LongRunningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interceptor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongRunningRequest.ProtoReflect.Descriptor instead.
func (*LongRunningRequest) Descriptor() ([]byte, []int) {
	return file_proto_interceptor_proto_rawDescGZIP(), []int{2}
}

type LongRunningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LongRunningResponse) Reset() {
	*x = LongRunningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interceptor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongRunningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongRunningResponse) ProtoMessage() {}

func (x *LongRunningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interceptor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongRunningResponse.ProtoReflect.Descriptor instead.
func (*LongRunningResponse) Descriptor() ([]byte, []int) {
	return file_proto_interceptor_proto_rawDescGZIP(), []int{3}
}

type ProtectedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProtectedRequest) Reset() {
	*x = ProtectedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interceptor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtectedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtectedRequest) ProtoMessage() {}

func (x *ProtectedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interceptor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtectedRequest.ProtoReflect.Descriptor instead.
func (*ProtectedRequest) Descriptor() ([]byte, []int) {
	return file_proto_interceptor_proto_rawDescGZIP(), []int{4}
}

type ProtectedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ProtectedResponse) Reset() {
	*x = ProtectedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_interceptor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtectedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtectedResponse) ProtoMessage() {}

func (x *ProtectedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_interceptor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtectedResponse.ProtoReflect.Descriptor instead.
func (*ProtectedResponse) Descriptor() ([]byte, []int) {
	return file_proto_interceptor_proto_rawDescGZIP(), []int{5}
}

func (x *ProtectedResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_proto_interceptor_proto protoreflect.FileDescriptor

var file_proto_interceptor_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a,
	0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4c,
	0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x11,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xfb, 0x01, 0x0a, 0x12, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1c, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x4c, 0x6f,
	0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x73, 0x68, 0x65, 0x70, 0x34, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_interceptor_proto_rawDescOnce sync.Once
	file_proto_interceptor_proto_rawDescData = file_proto_interceptor_proto_rawDesc
)

func file_proto_interceptor_proto_rawDescGZIP() []byte {
	file_proto_interceptor_proto_rawDescOnce.Do(func() {
		file_proto_interceptor_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_interceptor_proto_rawDescData)
	})
	return file_proto_interceptor_proto_rawDescData
}

var file_proto_interceptor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_interceptor_proto_goTypes = []interface{}{
	(*SayHelloRequest)(nil),     // 0: interceptor.SayHelloRequest
	(*SayHelloResponse)(nil),    // 1: interceptor.SayHelloResponse
	(*LongRunningRequest)(nil),  // 2: interceptor.LongRunningRequest
	(*LongRunningResponse)(nil), // 3: interceptor.LongRunningResponse
	(*ProtectedRequest)(nil),    // 4: interceptor.ProtectedRequest
	(*ProtectedResponse)(nil),   // 5: interceptor.ProtectedResponse
}
var file_proto_interceptor_proto_depIdxs = []int32{
	0, // 0: interceptor.InterceptorService.SayHello:input_type -> interceptor.SayHelloRequest
	2, // 1: interceptor.InterceptorService.LongRunning:input_type -> interceptor.LongRunningRequest
	4, // 2: interceptor.InterceptorService.Protected:input_type -> interceptor.ProtectedRequest
	1, // 3: interceptor.InterceptorService.SayHello:output_type -> interceptor.SayHelloResponse
	3, // 4: interceptor.InterceptorService.LongRunning:output_type -> interceptor.LongRunningResponse
	5, // 5: interceptor.InterceptorService.Protected:output_type -> interceptor.ProtectedResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_interceptor_proto_init() }
func file_proto_interceptor_proto_init() {
	if File_proto_interceptor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_interceptor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_proto_interceptor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
		file_proto_interceptor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongRunningRequest); i {
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
		file_proto_interceptor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongRunningResponse); i {
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
		file_proto_interceptor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtectedRequest); i {
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
		file_proto_interceptor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtectedResponse); i {
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
			RawDescriptor: file_proto_interceptor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_interceptor_proto_goTypes,
		DependencyIndexes: file_proto_interceptor_proto_depIdxs,
		MessageInfos:      file_proto_interceptor_proto_msgTypes,
	}.Build()
	File_proto_interceptor_proto = out.File
	file_proto_interceptor_proto_rawDesc = nil
	file_proto_interceptor_proto_goTypes = nil
	file_proto_interceptor_proto_depIdxs = nil
}
