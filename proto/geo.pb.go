// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: qwibi/geo.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QPBxGeoObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid        string           `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Geometry   *QPBxGeometry    `protobuf:"bytes,2,opt,name=geometry,proto3" json:"geometry,omitempty"`
	Properties *structpb.Struct `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QPBxGeoObject) Reset() {
	*x = QPBxGeoObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_geo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxGeoObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxGeoObject) ProtoMessage() {}

func (x *QPBxGeoObject) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_geo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxGeoObject.ProtoReflect.Descriptor instead.
func (*QPBxGeoObject) Descriptor() ([]byte, []int) {
	return file_qwibi_geo_proto_rawDescGZIP(), []int{0}
}

func (x *QPBxGeoObject) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *QPBxGeoObject) GetGeometry() *QPBxGeometry {
	if x != nil {
		return x.Geometry
	}
	return nil
}

func (x *QPBxGeoObject) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

var File_qwibi_geo_proto protoreflect.FileDescriptor

var file_qwibi_geo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0d, 0x51, 0x50, 0x42, 0x78, 0x47, 0x65,
	0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x67, 0x65, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x51, 0x50,
	0x42, 0x78, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qwibi_geo_proto_rawDescOnce sync.Once
	file_qwibi_geo_proto_rawDescData = file_qwibi_geo_proto_rawDesc
)

func file_qwibi_geo_proto_rawDescGZIP() []byte {
	file_qwibi_geo_proto_rawDescOnce.Do(func() {
		file_qwibi_geo_proto_rawDescData = protoimpl.X.CompressGZIP(file_qwibi_geo_proto_rawDescData)
	})
	return file_qwibi_geo_proto_rawDescData
}

var file_qwibi_geo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_qwibi_geo_proto_goTypes = []interface{}{
	(*QPBxGeoObject)(nil),   // 0: QPBxGeoObject
	(*QPBxGeometry)(nil),    // 1: QPBxGeometry
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
}
var file_qwibi_geo_proto_depIdxs = []int32{
	1, // 0: QPBxGeoObject.geometry:type_name -> QPBxGeometry
	2, // 1: QPBxGeoObject.properties:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_qwibi_geo_proto_init() }
func file_qwibi_geo_proto_init() {
	if File_qwibi_geo_proto != nil {
		return
	}
	file_qwibi_geometry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_qwibi_geo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxGeoObject); i {
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
			RawDescriptor: file_qwibi_geo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qwibi_geo_proto_goTypes,
		DependencyIndexes: file_qwibi_geo_proto_depIdxs,
		MessageInfos:      file_qwibi_geo_proto_msgTypes,
	}.Build()
	File_qwibi_geo_proto = out.File
	file_qwibi_geo_proto_rawDesc = nil
	file_qwibi_geo_proto_goTypes = nil
	file_qwibi_geo_proto_depIdxs = nil
}
