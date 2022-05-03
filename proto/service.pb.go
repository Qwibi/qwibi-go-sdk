// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: qwibi/service.proto

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

/// Auth
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
		mi := &file_qwibi_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxAuthRequest) ProtoMessage() {}

func (x *QPBxAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_service_proto_msgTypes[0]
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
	return file_qwibi_service_proto_rawDescGZIP(), []int{0}
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
		mi := &file_qwibi_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxAuthResponse) ProtoMessage() {}

func (x *QPBxAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_service_proto_msgTypes[1]
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
	return file_qwibi_service_proto_rawDescGZIP(), []int{1}
}

func (x *QPBxAuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

/// Join
type QPBxJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid string `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
}

func (x *QPBxJoinRequest) Reset() {
	*x = QPBxJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxJoinRequest) ProtoMessage() {}

func (x *QPBxJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxJoinRequest.ProtoReflect.Descriptor instead.
func (*QPBxJoinRequest) Descriptor() ([]byte, []int) {
	return file_qwibi_service_proto_rawDescGZIP(), []int{2}
}

func (x *QPBxJoinRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

type QPBxJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *QPBxGeoObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *QPBxJoinResponse) Reset() {
	*x = QPBxJoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxJoinResponse) ProtoMessage() {}

func (x *QPBxJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxJoinResponse.ProtoReflect.Descriptor instead.
func (*QPBxJoinResponse) Descriptor() ([]byte, []int) {
	return file_qwibi_service_proto_rawDescGZIP(), []int{3}
}

func (x *QPBxJoinResponse) GetObject() *QPBxGeoObject {
	if x != nil {
		return x.Object
	}
	return nil
}

/// Post
type QPBxPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *QPBxGeoObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *QPBxPostRequest) Reset() {
	*x = QPBxPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPostRequest) ProtoMessage() {}

func (x *QPBxPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPostRequest.ProtoReflect.Descriptor instead.
func (*QPBxPostRequest) Descriptor() ([]byte, []int) {
	return file_qwibi_service_proto_rawDescGZIP(), []int{4}
}

func (x *QPBxPostRequest) GetObject() *QPBxGeoObject {
	if x != nil {
		return x.Object
	}
	return nil
}

type QPBxPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *QPBxGeoObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *QPBxPostResponse) Reset() {
	*x = QPBxPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxPostResponse) ProtoMessage() {}

func (x *QPBxPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxPostResponse.ProtoReflect.Descriptor instead.
func (*QPBxPostResponse) Descriptor() ([]byte, []int) {
	return file_qwibi_service_proto_rawDescGZIP(), []int{5}
}

func (x *QPBxPostResponse) GetObject() *QPBxGeoObject {
	if x != nil {
		return x.Object
	}
	return nil
}

/// Stream
type QPBxStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QPBxStreamRequest) Reset() {
	*x = QPBxStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxStreamRequest) ProtoMessage() {}

func (x *QPBxStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxStreamRequest.ProtoReflect.Descriptor instead.
func (*QPBxStreamRequest) Descriptor() ([]byte, []int) {
	return file_qwibi_service_proto_rawDescGZIP(), []int{6}
}

type QPBxStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QPBxStreamResponse) Reset() {
	*x = QPBxStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qwibi_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QPBxStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QPBxStreamResponse) ProtoMessage() {}

func (x *QPBxStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qwibi_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QPBxStreamResponse.ProtoReflect.Descriptor instead.
func (*QPBxStreamResponse) Descriptor() ([]byte, []int) {
	return file_qwibi_service_proto_rawDescGZIP(), []int{7}
}

var File_qwibi_service_proto protoreflect.FileDescriptor

var file_qwibi_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x71, 0x77, 0x69, 0x62, 0x69, 0x2f, 0x67,
	0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0f, 0x51, 0x50, 0x42, 0x78,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x51, 0x50,
	0x42, 0x78, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x06,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x42, 0x06,
	0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x28, 0x0a, 0x10, 0x51, 0x50, 0x42, 0x78, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x23, 0x0a, 0x0f, 0x51, 0x50, 0x42, 0x78, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x67, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x51, 0x50, 0x42, 0x78, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x51, 0x50, 0x42, 0x78,
	0x47, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x39, 0x0a, 0x0f, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x47, 0x65, 0x6f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x0a, 0x10,
	0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x47, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x51, 0x50, 0x42, 0x78,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a,
	0x12, 0x51, 0x50, 0x42, 0x78, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xc7, 0x01, 0x0a, 0x07, 0x51, 0x50, 0x42, 0x78, 0x41, 0x70, 0x69, 0x12,
	0x2b, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x10, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x51, 0x50, 0x42, 0x78,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x10, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x12, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x51, 0x50, 0x42, 0x78, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qwibi_service_proto_rawDescOnce sync.Once
	file_qwibi_service_proto_rawDescData = file_qwibi_service_proto_rawDesc
)

func file_qwibi_service_proto_rawDescGZIP() []byte {
	file_qwibi_service_proto_rawDescOnce.Do(func() {
		file_qwibi_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_qwibi_service_proto_rawDescData)
	})
	return file_qwibi_service_proto_rawDescData
}

var file_qwibi_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_qwibi_service_proto_goTypes = []interface{}{
	(*QPBxAuthRequest)(nil),    // 0: QPBxAuthRequest
	(*QPBxAuthResponse)(nil),   // 1: QPBxAuthResponse
	(*QPBxJoinRequest)(nil),    // 2: QPBxJoinRequest
	(*QPBxJoinResponse)(nil),   // 3: QPBxJoinResponse
	(*QPBxPostRequest)(nil),    // 4: QPBxPostRequest
	(*QPBxPostResponse)(nil),   // 5: QPBxPostResponse
	(*QPBxStreamRequest)(nil),  // 6: QPBxStreamRequest
	(*QPBxStreamResponse)(nil), // 7: QPBxStreamResponse
	(*QPBxAnonymAuth)(nil),     // 8: QPBxAnonymAuth
	(*QPBxBasicAuth)(nil),      // 9: QPBxBasicAuth
	(*QPBxGeoObject)(nil),      // 10: QPBxGeoObject
}
var file_qwibi_service_proto_depIdxs = []int32{
	8,  // 0: QPBxAuthRequest.anonym:type_name -> QPBxAnonymAuth
	9,  // 1: QPBxAuthRequest.basic:type_name -> QPBxBasicAuth
	10, // 2: QPBxJoinResponse.object:type_name -> QPBxGeoObject
	10, // 3: QPBxPostRequest.object:type_name -> QPBxGeoObject
	10, // 4: QPBxPostResponse.object:type_name -> QPBxGeoObject
	0,  // 5: QPBxApi.Auth:input_type -> QPBxAuthRequest
	2,  // 6: QPBxApi.Join:input_type -> QPBxJoinRequest
	4,  // 7: QPBxApi.Post:input_type -> QPBxPostRequest
	6,  // 8: QPBxApi.Stream:input_type -> QPBxStreamRequest
	1,  // 9: QPBxApi.Auth:output_type -> QPBxAuthResponse
	3,  // 10: QPBxApi.Join:output_type -> QPBxJoinResponse
	5,  // 11: QPBxApi.Post:output_type -> QPBxPostResponse
	7,  // 12: QPBxApi.Stream:output_type -> QPBxStreamResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_qwibi_service_proto_init() }
func file_qwibi_service_proto_init() {
	if File_qwibi_service_proto != nil {
		return
	}
	file_qwibi_auth_proto_init()
	file_qwibi_geo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_qwibi_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_qwibi_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_qwibi_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxJoinRequest); i {
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
		file_qwibi_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxJoinResponse); i {
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
		file_qwibi_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPostRequest); i {
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
		file_qwibi_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxPostResponse); i {
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
		file_qwibi_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxStreamRequest); i {
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
		file_qwibi_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QPBxStreamResponse); i {
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
	file_qwibi_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QPBxAuthRequest_Anonym)(nil),
		(*QPBxAuthRequest_Basic)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qwibi_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qwibi_service_proto_goTypes,
		DependencyIndexes: file_qwibi_service_proto_depIdxs,
		MessageInfos:      file_qwibi_service_proto_msgTypes,
	}.Build()
	File_qwibi_service_proto = out.File
	file_qwibi_service_proto_rawDesc = nil
	file_qwibi_service_proto_goTypes = nil
	file_qwibi_service_proto_depIdxs = nil
}
