// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: geometry.proto

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

type QPBxGeometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*QPBxGeometry_Point
	//	*QPBxGeometry_Line
	Type isQPBxGeometry_Type `protobuf_oneof:"type"`
}

func (x *QPBxGeometry) Reset() {
	*x = QPBxGeometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxGeometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxGeometry) ProtoMessage() {}

func (x *QPBxGeometry) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxGeometry.ProtoReflect.Descriptor instead.
func (*QPBxGeometry) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{0}
}

func (m *QPBxGeometry) GetType() isQPBxGeometry_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *QPBxGeometry) GetPoint() *QPBxPoint {
	if x, ok := x.GetType().(*QPBxGeometry_Point); ok {
		return x.Point
	}
	return nil
}

func (x *QPBxGeometry) GetLine() *QPBxLine {
	if x, ok := x.GetType().(*QPBxGeometry_Line); ok {
		return x.Line
	}
	return nil
}

type isQPBxGeometry_Type interface {
	isQPBxGeometry_Type()
}

type QPBxGeometry_Point struct {
	Point *QPBxPoint `protobuf:"bytes,1,opt,name=point,proto3,oneof"`
}

type QPBxGeometry_Line struct {
	Line *QPBxLine `protobuf:"bytes,2,opt,name=line,proto3,oneof"`
}

func (*QPBxGeometry_Point) isQPBxGeometry_Type() {}

func (*QPBxGeometry_Line) isQPBxGeometry_Type() {}

type QPBxPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates []float64 `protobuf:"fixed64,1,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Options     []byte    `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *QPBxPoint) Reset() {
	*x = QPBxPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPoint) ProtoMessage() {}

func (x *QPBxPoint) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPoint.ProtoReflect.Descriptor instead.
func (*QPBxPoint) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{1}
}

func (x *QPBxPoint) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *QPBxPoint) GetOptions() []byte {
	if x != nil {
		return x.Options
	}
	return nil
}

type QPBxLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points  []*QPBxPoint `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	Options []byte       `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *QPBxLine) Reset() {
	*x = QPBxLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxLine) ProtoMessage() {}

func (x *QPBxLine) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxLine.ProtoReflect.Descriptor instead.
func (*QPBxLine) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{2}
}

func (x *QPBxLine) GetPoints() []*QPBxPoint {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *QPBxLine) GetOptions() []byte {
	if x != nil {
		return x.Options
	}
	return nil
}

var File_geometry_proto protoreflect.FileDescriptor

var file_geometry_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5b, 0x0a, 0x0c, 0x51, 0x50, 0x42, 0x78, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x4c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a,
	0x09, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x48, 0x0a, 0x08, 0x51, 0x50, 0x42, 0x78, 0x4c, 0x69,
	0x6e, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geometry_proto_rawDescOnce sync.Once
	file_geometry_proto_rawDescData = file_geometry_proto_rawDesc
)

func file_geometry_proto_rawDescGZIP() []byte {
	file_geometry_proto_rawDescOnce.Do(func() {
		file_geometry_proto_rawDescData = protoimpl.X.CompressGZIP(file_geometry_proto_rawDescData)
	})
	return file_geometry_proto_rawDescData
}

var file_geometry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_geometry_proto_goTypes = []interface{}{
	(*QPBxGeometry)(nil), // 0: QPBxGeometry
	(*QPBxPoint)(nil),    // 1: QPBxPoint
	(*QPBxLine)(nil),     // 2: QPBxLine
}
var file_geometry_proto_depIdxs = []int32{
	1, // 0: QPBxGeometry.point:type_name -> QPBxPoint
	2, // 1: QPBxGeometry.line:type_name -> QPBxLine
	1, // 2: QPBxLine.points:type_name -> QPBxPoint
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_geometry_proto_init() }
func file_geometry_proto_init() {
	if File_geometry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geometry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxGeometry); i {
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
		file_geometry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPoint); i {
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
		file_geometry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxLine); i {
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
	file_geometry_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QPBxGeometry_Point)(nil),
		(*QPBxGeometry_Line)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geometry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geometry_proto_goTypes,
		DependencyIndexes: file_geometry_proto_depIdxs,
		MessageInfos:      file_geometry_proto_msgTypes,
	}.Build()
	File_geometry_proto = out.File
	file_geometry_proto_rawDesc = nil
	file_geometry_proto_goTypes = nil
	file_geometry_proto_depIdxs = nil
}
