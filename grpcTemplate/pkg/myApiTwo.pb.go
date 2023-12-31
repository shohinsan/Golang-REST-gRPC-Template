// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: myApiTwo.proto

package pkg

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ThirdGetRpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ThirdGetRpc) Reset() {
	*x = ThirdGetRpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myApiTwo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThirdGetRpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThirdGetRpc) ProtoMessage() {}

func (x *ThirdGetRpc) ProtoReflect() protoreflect.Message {
	mi := &file_myApiTwo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThirdGetRpc.ProtoReflect.Descriptor instead.
func (*ThirdGetRpc) Descriptor() ([]byte, []int) {
	return file_myApiTwo_proto_rawDescGZIP(), []int{0}
}

func (x *ThirdGetRpc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FourthGetRpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *FourthGetRpc) Reset() {
	*x = FourthGetRpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myApiTwo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FourthGetRpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FourthGetRpc) ProtoMessage() {}

func (x *FourthGetRpc) ProtoReflect() protoreflect.Message {
	mi := &file_myApiTwo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FourthGetRpc.ProtoReflect.Descriptor instead.
func (*FourthGetRpc) Descriptor() ([]byte, []int) {
	return file_myApiTwo_proto_rawDescGZIP(), []int{1}
}

func (x *FourthGetRpc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FourthGetRpc) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_myApiTwo_proto protoreflect.FileDescriptor

var file_myApiTwo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x54, 0x77, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0b,
	0x54, 0x68, 0x69, 0x72, 0x64, 0x47, 0x65, 0x74, 0x52, 0x70, 0x63, 0x12, 0x2b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xba, 0x48, 0x14, 0xba, 0x01,
	0x11, 0x1a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x27, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x0c, 0x46, 0x6f, 0x75,
	0x72, 0x74, 0x68, 0x47, 0x65, 0x74, 0x52, 0x70, 0x63, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xba, 0x48, 0x15, 0xba, 0x01, 0x12, 0x1a,
	0x10, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x66, 0x6f, 0x75, 0x72, 0x74, 0x68,
	0x27, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0x48, 0x05, 0xb2, 0x01, 0x02, 0x38,
	0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xa9, 0x01, 0x0a,
	0x0b, 0x4d, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x12, 0x4b, 0x0a, 0x0b,
	0x54, 0x68, 0x69, 0x72, 0x64, 0x47, 0x65, 0x74, 0x52, 0x70, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x08, 0x12, 0x06, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x46, 0x6f, 0x75,
	0x72, 0x74, 0x68, 0x47, 0x65, 0x74, 0x52, 0x70, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12,
	0x07, 0x2f, 0x66, 0x6f, 0x75, 0x72, 0x74, 0x68, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x70, 0x6b, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_myApiTwo_proto_rawDescOnce sync.Once
	file_myApiTwo_proto_rawDescData = file_myApiTwo_proto_rawDesc
)

func file_myApiTwo_proto_rawDescGZIP() []byte {
	file_myApiTwo_proto_rawDescOnce.Do(func() {
		file_myApiTwo_proto_rawDescData = protoimpl.X.CompressGZIP(file_myApiTwo_proto_rawDescData)
	})
	return file_myApiTwo_proto_rawDescData
}

var file_myApiTwo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_myApiTwo_proto_goTypes = []interface{}{
	(*ThirdGetRpc)(nil),           // 0: routes.ThirdGetRpc
	(*FourthGetRpc)(nil),          // 1: routes.FourthGetRpc
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
	(*httpbody.HttpBody)(nil),     // 4: google.api.HttpBody
}
var file_myApiTwo_proto_depIdxs = []int32{
	2, // 0: routes.FourthGetRpc.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: routes.MySecondApi.ThirdGetRpc:input_type -> google.protobuf.Empty
	3, // 2: routes.MySecondApi.FourthGetRpc:input_type -> google.protobuf.Empty
	4, // 3: routes.MySecondApi.ThirdGetRpc:output_type -> google.api.HttpBody
	4, // 4: routes.MySecondApi.FourthGetRpc:output_type -> google.api.HttpBody
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_myApiTwo_proto_init() }
func file_myApiTwo_proto_init() {
	if File_myApiTwo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_myApiTwo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThirdGetRpc); i {
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
		file_myApiTwo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FourthGetRpc); i {
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
			RawDescriptor: file_myApiTwo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_myApiTwo_proto_goTypes,
		DependencyIndexes: file_myApiTwo_proto_depIdxs,
		MessageInfos:      file_myApiTwo_proto_msgTypes,
	}.Build()
	File_myApiTwo_proto = out.File
	file_myApiTwo_proto_rawDesc = nil
	file_myApiTwo_proto_goTypes = nil
	file_myApiTwo_proto_depIdxs = nil
}
