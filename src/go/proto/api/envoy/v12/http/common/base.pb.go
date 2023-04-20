// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/envoy/v12/http/common/base.proto

package common

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DependencyErrorBehavior int32

const (
	DependencyErrorBehavior_UNSPECIFIED             DependencyErrorBehavior = 0
	DependencyErrorBehavior_BLOCK_INIT_ON_ANY_ERROR DependencyErrorBehavior = 1
	DependencyErrorBehavior_ALWAYS_INIT             DependencyErrorBehavior = 2
)

// Enum value maps for DependencyErrorBehavior.
var (
	DependencyErrorBehavior_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "BLOCK_INIT_ON_ANY_ERROR",
		2: "ALWAYS_INIT",
	}
	DependencyErrorBehavior_value = map[string]int32{
		"UNSPECIFIED":             0,
		"BLOCK_INIT_ON_ANY_ERROR": 1,
		"ALWAYS_INIT":             2,
	}
)

func (x DependencyErrorBehavior) Enum() *DependencyErrorBehavior {
	p := new(DependencyErrorBehavior)
	*p = x
	return p
}

func (x DependencyErrorBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DependencyErrorBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_api_envoy_v12_http_common_base_proto_enumTypes[0].Descriptor()
}

func (DependencyErrorBehavior) Type() protoreflect.EnumType {
	return &file_api_envoy_v12_http_common_base_proto_enumTypes[0]
}

func (x DependencyErrorBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DependencyErrorBehavior.Descriptor instead.
func (DependencyErrorBehavior) EnumDescriptor() ([]byte, []int) {
	return file_api_envoy_v12_http_common_base_proto_rawDescGZIP(), []int{0}
}

type HttpUri struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri     string               `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Cluster string               `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HttpUri) Reset() {
	*x = HttpUri{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_envoy_v12_http_common_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpUri) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpUri) ProtoMessage() {}

func (x *HttpUri) ProtoReflect() protoreflect.Message {
	mi := &file_api_envoy_v12_http_common_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpUri.ProtoReflect.Descriptor instead.
func (*HttpUri) Descriptor() ([]byte, []int) {
	return file_api_envoy_v12_http_common_base_proto_rawDescGZIP(), []int{0}
}

func (x *HttpUri) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *HttpUri) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *HttpUri) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TokenType:
	//	*AccessToken_RemoteToken
	TokenType isAccessToken_TokenType `protobuf_oneof:"token_type"`
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_envoy_v12_http_common_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_api_envoy_v12_http_common_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_api_envoy_v12_http_common_base_proto_rawDescGZIP(), []int{1}
}

func (m *AccessToken) GetTokenType() isAccessToken_TokenType {
	if m != nil {
		return m.TokenType
	}
	return nil
}

func (x *AccessToken) GetRemoteToken() *HttpUri {
	if x, ok := x.GetTokenType().(*AccessToken_RemoteToken); ok {
		return x.RemoteToken
	}
	return nil
}

type isAccessToken_TokenType interface {
	isAccessToken_TokenType()
}

type AccessToken_RemoteToken struct {
	RemoteToken *HttpUri `protobuf:"bytes,1,opt,name=remote_token,json=remoteToken,proto3,oneof"`
}

func (*AccessToken_RemoteToken) isAccessToken_TokenType() {}

type IamTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IamUri              *HttpUri     `protobuf:"bytes,1,opt,name=iam_uri,json=iamUri,proto3" json:"iam_uri,omitempty"`
	AccessToken         *AccessToken `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ServiceAccountEmail string       `protobuf:"bytes,3,opt,name=service_account_email,json=serviceAccountEmail,proto3" json:"service_account_email,omitempty"`
	Delegates           []string     `protobuf:"bytes,4,rep,name=delegates,proto3" json:"delegates,omitempty"`
}

func (x *IamTokenInfo) Reset() {
	*x = IamTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_envoy_v12_http_common_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamTokenInfo) ProtoMessage() {}

func (x *IamTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_envoy_v12_http_common_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamTokenInfo.ProtoReflect.Descriptor instead.
func (*IamTokenInfo) Descriptor() ([]byte, []int) {
	return file_api_envoy_v12_http_common_base_proto_rawDescGZIP(), []int{2}
}

func (x *IamTokenInfo) GetIamUri() *HttpUri {
	if x != nil {
		return x.IamUri
	}
	return nil
}

func (x *IamTokenInfo) GetAccessToken() *AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *IamTokenInfo) GetServiceAccountEmail() string {
	if x != nil {
		return x.ServiceAccountEmail
	}
	return ""
}

func (x *IamTokenInfo) GetDelegates() []string {
	if x != nil {
		return x.Delegates
	}
	return nil
}

var File_api_envoy_v12_http_common_base_proto protoreflect.FileDescriptor

var file_api_envoy_v12_http_common_base_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x32, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x65, 0x73, 0x70, 0x76, 0x32, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8e, 0x01, 0x0a, 0x07, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x12, 0x1f, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x72, 0x08,
	0x20, 0x01, 0xc0, 0x01, 0x02, 0xc8, 0x01, 0x00, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x21, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x3f, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0xaa, 0x01, 0x04, 0x08, 0x01, 0x32, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x22, 0x75, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x4d, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x73, 0x70, 0x76, 0x32, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69,
	0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x11, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8,
	0x42, 0x01, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xf4, 0x01, 0x0a, 0x0c, 0x49, 0x61, 0x6d,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x07, 0x69, 0x61, 0x6d,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x73, 0x70,
	0x76, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x32,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x55, 0x72, 0x69, 0x52, 0x06, 0x69, 0x61, 0x6d, 0x55, 0x72, 0x69, 0x12, 0x4f, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x73, 0x70, 0x76, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x0a,
	0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2a,
	0x58, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x4e, 0x59,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x4c, 0x57, 0x41,
	0x59, 0x53, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_envoy_v12_http_common_base_proto_rawDescOnce sync.Once
	file_api_envoy_v12_http_common_base_proto_rawDescData = file_api_envoy_v12_http_common_base_proto_rawDesc
)

func file_api_envoy_v12_http_common_base_proto_rawDescGZIP() []byte {
	file_api_envoy_v12_http_common_base_proto_rawDescOnce.Do(func() {
		file_api_envoy_v12_http_common_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_envoy_v12_http_common_base_proto_rawDescData)
	})
	return file_api_envoy_v12_http_common_base_proto_rawDescData
}

var file_api_envoy_v12_http_common_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_envoy_v12_http_common_base_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_envoy_v12_http_common_base_proto_goTypes = []interface{}{
	(DependencyErrorBehavior)(0), // 0: espv2.api.envoy.v12.http.common.DependencyErrorBehavior
	(*HttpUri)(nil),              // 1: espv2.api.envoy.v12.http.common.HttpUri
	(*AccessToken)(nil),          // 2: espv2.api.envoy.v12.http.common.AccessToken
	(*IamTokenInfo)(nil),         // 3: espv2.api.envoy.v12.http.common.IamTokenInfo
	(*durationpb.Duration)(nil),  // 4: google.protobuf.Duration
}
var file_api_envoy_v12_http_common_base_proto_depIdxs = []int32{
	4, // 0: espv2.api.envoy.v12.http.common.HttpUri.timeout:type_name -> google.protobuf.Duration
	1, // 1: espv2.api.envoy.v12.http.common.AccessToken.remote_token:type_name -> espv2.api.envoy.v12.http.common.HttpUri
	1, // 2: espv2.api.envoy.v12.http.common.IamTokenInfo.iam_uri:type_name -> espv2.api.envoy.v12.http.common.HttpUri
	2, // 3: espv2.api.envoy.v12.http.common.IamTokenInfo.access_token:type_name -> espv2.api.envoy.v12.http.common.AccessToken
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_envoy_v12_http_common_base_proto_init() }
func file_api_envoy_v12_http_common_base_proto_init() {
	if File_api_envoy_v12_http_common_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_envoy_v12_http_common_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpUri); i {
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
		file_api_envoy_v12_http_common_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessToken); i {
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
		file_api_envoy_v12_http_common_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamTokenInfo); i {
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
	file_api_envoy_v12_http_common_base_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AccessToken_RemoteToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_envoy_v12_http_common_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_envoy_v12_http_common_base_proto_goTypes,
		DependencyIndexes: file_api_envoy_v12_http_common_base_proto_depIdxs,
		EnumInfos:         file_api_envoy_v12_http_common_base_proto_enumTypes,
		MessageInfos:      file_api_envoy_v12_http_common_base_proto_msgTypes,
	}.Build()
	File_api_envoy_v12_http_common_base_proto = out.File
	file_api_envoy_v12_http_common_base_proto_rawDesc = nil
	file_api_envoy_v12_http_common_base_proto_goTypes = nil
	file_api_envoy_v12_http_common_base_proto_depIdxs = nil
}