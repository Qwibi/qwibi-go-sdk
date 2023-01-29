// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: qwibi/feature.proto

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

type QPBxLayerFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Properties []byte `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QPBxLayerFeature) Reset() {
	*x = QPBxLayerFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_feature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxLayerFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxLayerFeature) ProtoMessage() {}

func (x *QPBxLayerFeature) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_feature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxLayerFeature.ProtoReflect.Descriptor instead.
func (*QPBxLayerFeature) Descriptor() ([]byte, []int) {
	return file_qwibi_feature_proto_rawDescGZIP(), []int{0}
}

func (x *QPBxLayerFeature) GetProperties() []byte {
	if x != nil {
		return x.Properties
	}
	return nil
}

type QPBxPointFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Geometry   *QPBxPoint `protobuf:"bytes,1,opt,name=geometry,proto3" json:"geometry,omitempty"`
	Properties []byte     `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QPBxPointFeature) Reset() {
	*x = QPBxPointFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_feature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPointFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPointFeature) ProtoMessage() {}

func (x *QPBxPointFeature) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_feature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPointFeature.ProtoReflect.Descriptor instead.
func (*QPBxPointFeature) Descriptor() ([]byte, []int) {
	return file_qwibi_feature_proto_rawDescGZIP(), []int{1}
}

func (x *QPBxPointFeature) GetGeometry() *QPBxPoint {
	if x != nil {
		return x.Geometry
	}
	return nil
}

func (x *QPBxPointFeature) GetProperties() []byte {
	if x != nil {
		return x.Properties
	}
	return nil
}

type QPBxLineFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Geometry   *QPBxLine `protobuf:"bytes,1,opt,name=geometry,proto3" json:"geometry,omitempty"`
	Properties []byte    `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QPBxLineFeature) Reset() {
	*x = QPBxLineFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_feature_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxLineFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxLineFeature) ProtoMessage() {}

func (x *QPBxLineFeature) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_feature_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxLineFeature.ProtoReflect.Descriptor instead.
func (*QPBxLineFeature) Descriptor() ([]byte, []int) {
	return file_qwibi_feature_proto_rawDescGZIP(), []int{2}
}

func (x *QPBxLineFeature) GetGeometry() *QPBxLine {
	if x != nil {
		return x.Geometry
	}
	return nil
}

func (x *QPBxLineFeature) GetProperties() []byte {
	if x != nil {
		return x.Properties
	}
	return nil
}

var File_qwibi_feature_proto protoreflect.FileDescriptor

var file_qwibi_feature_proto_rawDesc = []byte{
	0x0a, 0x13, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x67, 0x65, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x10, 0x51,
	0x50, 0x42, 0x78, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x5a, 0x0a, 0x10, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x0f, 0x51,
	0x50, 0x42, 0x78, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x25,
	0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x67, 0x65, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qwibi_feature_proto_rawDescOnce sync.Once
	file_qwibi_feature_proto_rawDescData = file_qwibi_feature_proto_rawDesc
)

func file_qwibi_feature_proto_rawDescGZIP() []byte {
	file_qwibi_feature_proto_rawDescOnce.Do(func() {
		file_qwibi_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_qwibi_feature_proto_rawDescData)
	})
	return file_qwibi_feature_proto_rawDescData
}

var file_qwibi_feature_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_qwibi_feature_proto_goTypes = []interface{}{
	(*QPBxLayerFeature)(nil), // 0: QPBxLayerFeature
	(*QPBxPointFeature)(nil), // 1: QPBxPointFeature
	(*QPBxLineFeature)(nil),  // 2: QPBxLineFeature
	(*QPBxPoint)(nil),        // 3: QPBxPoint
	(*QPBxLine)(nil),         // 4: QPBxLine
}
var file_qwibi_feature_proto_depIdxs = []int32{
	3, // 0: QPBxPointFeature.geometry:type_name -> QPBxPoint
	4, // 1: QPBxLineFeature.geometry:type_name -> QPBxLine
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_qwibi_feature_proto_init() }
func file_qwibi_feature_proto_init() {
	if File_qwibi_feature_proto != nil {
		return
	}
	file_qwibi_geometry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_qwibi_feature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxLayerFeature); i {
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
		file_qwibi_feature_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPointFeature); i {
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
		file_qwibi_feature_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxLineFeature); i {
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
			RawDescriptor: file_qwibi_feature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qwibi_feature_proto_goTypes,
		DependencyIndexes: file_qwibi_feature_proto_depIdxs,
		MessageInfos:      file_qwibi_feature_proto_msgTypes,
	}.Build()
	File_qwibi_feature_proto = out.File
	file_qwibi_feature_proto_rawDesc = nil
	file_qwibi_feature_proto_goTypes = nil
	file_qwibi_feature_proto_depIdxs = nil
}
