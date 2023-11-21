// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: c1/connector/v2/annotation_resource_tree.proto

package v2

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

type ChildResourceType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceTypeId string `protobuf:"bytes,1,opt,name=resource_type_id,json=resourceTypeId,proto3" json:"resource_type_id,omitempty"`
}

func (x *ChildResourceType) Reset() {
	*x = ChildResourceType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c1_connector_v2_annotation_resource_tree_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildResourceType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildResourceType) ProtoMessage() {}

func (x *ChildResourceType) ProtoReflect() protoreflect.Message {
	mi := &file_c1_connector_v2_annotation_resource_tree_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildResourceType.ProtoReflect.Descriptor instead.
func (*ChildResourceType) Descriptor() ([]byte, []int) {
	return file_c1_connector_v2_annotation_resource_tree_proto_rawDescGZIP(), []int{0}
}

func (x *ChildResourceType) GetResourceTypeId() string {
	if x != nil {
		return x.ResourceTypeId
	}
	return ""
}

type SkipEntitlementsAndGrants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SkipEntitlementsAndGrants) Reset() {
	*x = SkipEntitlementsAndGrants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c1_connector_v2_annotation_resource_tree_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkipEntitlementsAndGrants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkipEntitlementsAndGrants) ProtoMessage() {}

func (x *SkipEntitlementsAndGrants) ProtoReflect() protoreflect.Message {
	mi := &file_c1_connector_v2_annotation_resource_tree_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkipEntitlementsAndGrants.ProtoReflect.Descriptor instead.
func (*SkipEntitlementsAndGrants) Descriptor() ([]byte, []int) {
	return file_c1_connector_v2_annotation_resource_tree_proto_rawDescGZIP(), []int{1}
}

var File_c1_connector_v2_annotation_resource_tree_proto protoreflect.FileDescriptor

var file_c1_connector_v2_annotation_resource_tree_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x63, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76,
	0x32, 0x22, 0x3d, 0x0a, 0x11, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x22, 0x1b, 0x0a, 0x19, 0x53, 0x6b, 0x69, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x41, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x64,
	0x75, 0x63, 0x74, 0x6f, 0x72, 0x6f, 0x6e, 0x65, 0x2f, 0x62, 0x61, 0x74, 0x6f, 0x6e, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_c1_connector_v2_annotation_resource_tree_proto_rawDescOnce sync.Once
	file_c1_connector_v2_annotation_resource_tree_proto_rawDescData = file_c1_connector_v2_annotation_resource_tree_proto_rawDesc
)

func file_c1_connector_v2_annotation_resource_tree_proto_rawDescGZIP() []byte {
	file_c1_connector_v2_annotation_resource_tree_proto_rawDescOnce.Do(func() {
		file_c1_connector_v2_annotation_resource_tree_proto_rawDescData = protoimpl.X.CompressGZIP(file_c1_connector_v2_annotation_resource_tree_proto_rawDescData)
	})
	return file_c1_connector_v2_annotation_resource_tree_proto_rawDescData
}

var file_c1_connector_v2_annotation_resource_tree_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_c1_connector_v2_annotation_resource_tree_proto_goTypes = []interface{}{
	(*ChildResourceType)(nil),         // 0: c1.connector.v2.ChildResourceType
	(*SkipEntitlementsAndGrants)(nil), // 1: c1.connector.v2.SkipEntitlementsAndGrants
}
var file_c1_connector_v2_annotation_resource_tree_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_c1_connector_v2_annotation_resource_tree_proto_init() }
func file_c1_connector_v2_annotation_resource_tree_proto_init() {
	if File_c1_connector_v2_annotation_resource_tree_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_c1_connector_v2_annotation_resource_tree_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildResourceType); i {
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
		file_c1_connector_v2_annotation_resource_tree_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkipEntitlementsAndGrants); i {
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
			RawDescriptor: file_c1_connector_v2_annotation_resource_tree_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_c1_connector_v2_annotation_resource_tree_proto_goTypes,
		DependencyIndexes: file_c1_connector_v2_annotation_resource_tree_proto_depIdxs,
		MessageInfos:      file_c1_connector_v2_annotation_resource_tree_proto_msgTypes,
	}.Build()
	File_c1_connector_v2_annotation_resource_tree_proto = out.File
	file_c1_connector_v2_annotation_resource_tree_proto_rawDesc = nil
	file_c1_connector_v2_annotation_resource_tree_proto_goTypes = nil
	file_c1_connector_v2_annotation_resource_tree_proto_depIdxs = nil
}