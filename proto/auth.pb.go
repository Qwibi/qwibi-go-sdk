// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: qwibi/auth.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type QPBxAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Auth:
	//	*QPBxAuthRequest_Anonym
	//	*QPBxAuthRequest_Basic
	Auth isQPBxAuthRequest_Auth `protobuf_oneof:"auth"`
}

func (x *QPBxAuthRequest) Reset() {
	*x = QPBxAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxAuthRequest) ProtoMessage() {}

func (x *QPBxAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxAuthRequest.ProtoReflect.Descriptor instead.
func (*QPBxAuthRequest) Descriptor() ([]byte, []int) {
	return file_qwibi_auth_proto_rawDescGZIP(), []int{0}
}

func (m *QPBxAuthRequest) GetAuth() isQPBxAuthRequest_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (x *QPBxAuthRequest) GetAnonym() *QPBxAnonymAuth {
	if x, ok := x.GetAuth().(*QPBxAuthRequest_Anonym); ok {
		return x.Anonym
	}
	return nil
}

func (x *QPBxAuthRequest) GetBasic() *QPBxBasicAuth {
	if x, ok := x.GetAuth().(*QPBxAuthRequest_Basic); ok {
		return x.Basic
	}
	return nil
}

type isQPBxAuthRequest_Auth interface {
	isQPBxAuthRequest_Auth()
}

type QPBxAuthRequest_Anonym struct {
	Anonym *QPBxAnonymAuth `protobuf:"bytes,1,opt,name=anonym,proto3,oneof"`
}

type QPBxAuthRequest_Basic struct {
	Basic *QPBxBasicAuth `protobuf:"bytes,2,opt,name=basic,proto3,oneof"`
}

func (*QPBxAuthRequest_Anonym) isQPBxAuthRequest_Auth() {}

func (*QPBxAuthRequest_Basic) isQPBxAuthRequest_Auth() {}

type QPBxAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *QPBxAuthResponse) Reset() {
	*x = QPBxAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxAuthResponse) ProtoMessage() {}

func (x *QPBxAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxAuthResponse.ProtoReflect.Descriptor instead.
func (*QPBxAuthResponse) Descriptor() ([]byte, []int) {
	return file_qwibi_auth_proto_rawDescGZIP(), []int{1}
}

func (x *QPBxAuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type QPBxAnonymAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QPBxAnonymAuth) Reset() {
	*x = QPBxAnonymAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxAnonymAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxAnonymAuth) ProtoMessage() {}

func (x *QPBxAnonymAuth) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxAnonymAuth.ProtoReflect.Descriptor instead.
func (*QPBxAnonymAuth) Descriptor() ([]byte, []int) {
	return file_qwibi_auth_proto_rawDescGZIP(), []int{2}
}

type QPBxBasicAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *QPBxBasicAuth) Reset() {
	*x = QPBxBasicAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxBasicAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxBasicAuth) ProtoMessage() {}

func (x *QPBxBasicAuth) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxBasicAuth.ProtoReflect.Descriptor instead.
func (*QPBxBasicAuth) Descriptor() ([]byte, []int) {
	return file_qwibi_auth_proto_rawDescGZIP(), []int{3}
}

func (x *QPBxBasicAuth) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *QPBxBasicAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_qwibi_auth_proto protoreflect.FileDescriptor

var file_qwibi_auth_proto_rawDesc = []byte{
	0x0a, 0x10, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0f, 0x51, 0x50, 0x42, 0x78, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x41, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x12, 0x26, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x48,
	0x00, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x42, 0x06, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x22, 0x28, 0x0a, 0x10, 0x51, 0x50, 0x42, 0x78, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x51, 0x50,
	0x42, 0x78, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x22, 0x41, 0x0a, 0x0d,
	0x51, 0x50, 0x42, 0x78, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42,
	0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qwibi_auth_proto_rawDescOnce sync.Once
	file_qwibi_auth_proto_rawDescData = file_qwibi_auth_proto_rawDesc
)

func file_qwibi_auth_proto_rawDescGZIP() []byte {
	file_qwibi_auth_proto_rawDescOnce.Do(func() {
		file_qwibi_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_qwibi_auth_proto_rawDescData)
	})
	return file_qwibi_auth_proto_rawDescData
}

var file_qwibi_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_qwibi_auth_proto_goTypes = []interface{}{
	(*QPBxAuthRequest)(nil),  // 0: QPBxAuthRequest
	(*QPBxAuthResponse)(nil), // 1: QPBxAuthResponse
	(*QPBxAnonymAuth)(nil),   // 2: QPBxAnonymAuth
	(*QPBxBasicAuth)(nil),    // 3: QPBxBasicAuth
}
var file_qwibi_auth_proto_depIdxs = []int32{
	2, // 0: QPBxAuthRequest.anonym:type_name -> QPBxAnonymAuth
	3, // 1: QPBxAuthRequest.basic:type_name -> QPBxBasicAuth
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_qwibi_auth_proto_init() }
func file_qwibi_auth_proto_init() {
	if File_qwibi_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qwibi_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxAuthRequest); i {
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
		file_qwibi_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxAuthResponse); i {
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
		file_qwibi_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxAnonymAuth); i {
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
		file_qwibi_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxBasicAuth); i {
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
	file_qwibi_auth_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QPBxAuthRequest_Anonym)(nil),
		(*QPBxAuthRequest_Basic)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qwibi_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qwibi_auth_proto_goTypes,
		DependencyIndexes: file_qwibi_auth_proto_depIdxs,
		MessageInfos:      file_qwibi_auth_proto_msgTypes,
	}.Build()
	File_qwibi_auth_proto = out.File
	file_qwibi_auth_proto_rawDesc = nil
	file_qwibi_auth_proto_goTypes = nil
	file_qwibi_auth_proto_depIdxs = nil
}
