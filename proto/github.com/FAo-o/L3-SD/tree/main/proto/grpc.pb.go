// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: grpc.proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

type MunitionInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MunicionAT int32 `protobuf:"varint,1,opt,name=municionAT,proto3" json:"municionAT,omitempty"`
	MunicionMP int32 `protobuf:"varint,2,opt,name=municionMP,proto3" json:"municionMP,omitempty"`
	MaxAT      int32 `protobuf:"varint,3,opt,name=maxAT,proto3" json:"maxAT,omitempty"`
	MaxMP      int32 `protobuf:"varint,4,opt,name=maxMP,proto3" json:"maxMP,omitempty"`
}

func (x *MunitionInfoResponse) Reset() {
	*x = MunitionInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MunitionInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MunitionInfoResponse) ProtoMessage() {}

func (x *MunitionInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MunitionInfoResponse.ProtoReflect.Descriptor instead.
func (*MunitionInfoResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *MunitionInfoResponse) GetMunicionAT() int32 {
	if x != nil {
		return x.MunicionAT
	}
	return 0
}

func (x *MunitionInfoResponse) GetMunicionMP() int32 {
	if x != nil {
		return x.MunicionMP
	}
	return 0
}

func (x *MunitionInfoResponse) GetMaxAT() int32 {
	if x != nil {
		return x.MaxAT
	}
	return 0
}

func (x *MunitionInfoResponse) GetMaxMP() int32 {
	if x != nil {
		return x.MaxMP
	}
	return 0
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x82, 0x01, 0x0a, 0x14,
	0x4d, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x6e,
	0x41, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x69,
	0x6f, 0x6e, 0x41, 0x54, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x6e,
	0x4d, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x69,
	0x6f, 0x6e, 0x4d, 0x50, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x78, 0x41, 0x54, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x41, 0x54, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61,
	0x78, 0x4d, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x4d, 0x50,
	0x32, 0x4d, 0x0a, 0x0f, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x75, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x41,
	0x6f, 0x2d, 0x6f, 0x2f, 0x4c, 0x33, 0x2d, 0x53, 0x44, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_proto_goTypes = []interface{}{
	(*Empty)(nil),                // 0: grpc.Empty
	(*MunitionInfoResponse)(nil), // 1: grpc.MunitionInfoResponse
}
var file_grpc_proto_depIdxs = []int32{
	0, // 0: grpc.WishListService.GetMunitionInfo:input_type -> grpc.Empty
	1, // 1: grpc.WishListService.GetMunitionInfo:output_type -> grpc.MunitionInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MunitionInfoResponse); i {
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
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}
