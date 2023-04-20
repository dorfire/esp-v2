// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/envoy/v12/http/service_control/requirement.proto

package service_control

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ApiKeyLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Key:
	//	*ApiKeyLocation_Query
	//	*ApiKeyLocation_Header
	//	*ApiKeyLocation_Cookie
	Key isApiKeyLocation_Key `protobuf_oneof:"key"`
}

func (x *ApiKeyLocation) Reset() {
	*x = ApiKeyLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyLocation) ProtoMessage() {}

func (x *ApiKeyLocation) ProtoReflect() protoreflect.Message {
	mi := &file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyLocation.ProtoReflect.Descriptor instead.
func (*ApiKeyLocation) Descriptor() ([]byte, []int) {
	return file_api_envoy_v12_http_service_control_requirement_proto_rawDescGZIP(), []int{0}
}

func (m *ApiKeyLocation) GetKey() isApiKeyLocation_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (x *ApiKeyLocation) GetQuery() string {
	if x, ok := x.GetKey().(*ApiKeyLocation_Query); ok {
		return x.Query
	}
	return ""
}

func (x *ApiKeyLocation) GetHeader() string {
	if x, ok := x.GetKey().(*ApiKeyLocation_Header); ok {
		return x.Header
	}
	return ""
}

func (x *ApiKeyLocation) GetCookie() string {
	if x, ok := x.GetKey().(*ApiKeyLocation_Cookie); ok {
		return x.Cookie
	}
	return ""
}

type isApiKeyLocation_Key interface {
	isApiKeyLocation_Key()
}

type ApiKeyLocation_Query struct {
	Query string `protobuf:"bytes,1,opt,name=query,proto3,oneof"`
}

type ApiKeyLocation_Header struct {
	Header string `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

type ApiKeyLocation_Cookie struct {
	Cookie string `protobuf:"bytes,3,opt,name=cookie,proto3,oneof"`
}

func (*ApiKeyLocation_Query) isApiKeyLocation_Key() {}

func (*ApiKeyLocation_Header) isApiKeyLocation_Key() {}

func (*ApiKeyLocation_Cookie) isApiKeyLocation_Key() {}

type ApiKeyRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations          []*ApiKeyLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	AllowWithoutApiKey bool              `protobuf:"varint,2,opt,name=allow_without_api_key,json=allowWithoutApiKey,proto3" json:"allow_without_api_key,omitempty"`
}

func (x *ApiKeyRequirement) Reset() {
	*x = ApiKeyRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyRequirement) ProtoMessage() {}

func (x *ApiKeyRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyRequirement.ProtoReflect.Descriptor instead.
func (*ApiKeyRequirement) Descriptor() ([]byte, []int) {
	return file_api_envoy_v12_http_service_control_requirement_proto_rawDescGZIP(), []int{1}
}

func (x *ApiKeyRequirement) GetLocations() []*ApiKeyLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *ApiKeyRequirement) GetAllowWithoutApiKey() bool {
	if x != nil {
		return x.AllowWithoutApiKey
	}
	return false
}

type MetricCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cost int64  `protobuf:"varint,2,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *MetricCost) Reset() {
	*x = MetricCost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricCost) ProtoMessage() {}

func (x *MetricCost) ProtoReflect() protoreflect.Message {
	mi := &file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricCost.ProtoReflect.Descriptor instead.
func (*MetricCost) Descriptor() ([]byte, []int) {
	return file_api_envoy_v12_http_service_control_requirement_proto_rawDescGZIP(), []int{2}
}

func (x *MetricCost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricCost) GetCost() int64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type Requirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName        string             `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	OperationName      string             `protobuf:"bytes,2,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	ApiKey             *ApiKeyRequirement `protobuf:"bytes,3,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	CustomLabels       []string           `protobuf:"bytes,4,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty"`
	ApiName            string             `protobuf:"bytes,5,opt,name=api_name,json=apiName,proto3" json:"api_name,omitempty"`
	ApiVersion         string             `protobuf:"bytes,6,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	SkipServiceControl bool               `protobuf:"varint,7,opt,name=skip_service_control,json=skipServiceControl,proto3" json:"skip_service_control,omitempty"`
	MetricCosts        []*MetricCost      `protobuf:"bytes,8,rep,name=metric_costs,json=metricCosts,proto3" json:"metric_costs,omitempty"`
}

func (x *Requirement) Reset() {
	*x = Requirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Requirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Requirement) ProtoMessage() {}

func (x *Requirement) ProtoReflect() protoreflect.Message {
	mi := &file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Requirement.ProtoReflect.Descriptor instead.
func (*Requirement) Descriptor() ([]byte, []int) {
	return file_api_envoy_v12_http_service_control_requirement_proto_rawDescGZIP(), []int{3}
}

func (x *Requirement) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Requirement) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

func (x *Requirement) GetApiKey() *ApiKeyRequirement {
	if x != nil {
		return x.ApiKey
	}
	return nil
}

func (x *Requirement) GetCustomLabels() []string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *Requirement) GetApiName() string {
	if x != nil {
		return x.ApiName
	}
	return ""
}

func (x *Requirement) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Requirement) GetSkipServiceControl() bool {
	if x != nil {
		return x.SkipServiceControl
	}
	return false
}

func (x *Requirement) GetMetricCosts() []*MetricCost {
	if x != nil {
		return x.MetricCosts
	}
	return nil
}

var File_api_envoy_v12_http_service_control_requirement_proto protoreflect.FileDescriptor

var file_api_envoy_v12_http_service_control_requirement_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x32, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x65, 0x73, 0x70, 0x76, 0x32, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0e, 0x41, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x15,
	0x72, 0x13, 0x20, 0x01, 0x32, 0x0f, 0x5e, 0x5b, 0x5e, 0x3f, 0x26, 0x23, 0x5c, 0x72, 0x5c, 0x6e,
	0x5c, 0x30, 0x5d, 0x2b, 0x24, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x24,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x20, 0x01, 0xc0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x20, 0x01, 0xc0, 0x01, 0x02,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x42, 0x0a, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x9e, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x65, 0x73, 0x70, 0x76, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75,
	0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x34, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0xab, 0x03,
	0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x61, 0x70, 0x69,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x65, 0x73, 0x70,
	0x76, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x32,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x30, 0x0a, 0x14, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x73, 0x6b, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x12, 0x57, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x73,
	0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x73, 0x70, 0x76, 0x32,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x0b,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_envoy_v12_http_service_control_requirement_proto_rawDescOnce sync.Once
	file_api_envoy_v12_http_service_control_requirement_proto_rawDescData = file_api_envoy_v12_http_service_control_requirement_proto_rawDesc
)

func file_api_envoy_v12_http_service_control_requirement_proto_rawDescGZIP() []byte {
	file_api_envoy_v12_http_service_control_requirement_proto_rawDescOnce.Do(func() {
		file_api_envoy_v12_http_service_control_requirement_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_envoy_v12_http_service_control_requirement_proto_rawDescData)
	})
	return file_api_envoy_v12_http_service_control_requirement_proto_rawDescData
}

var file_api_envoy_v12_http_service_control_requirement_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_envoy_v12_http_service_control_requirement_proto_goTypes = []interface{}{
	(*ApiKeyLocation)(nil),    // 0: espv2.api.envoy.v12.http.service_control.ApiKeyLocation
	(*ApiKeyRequirement)(nil), // 1: espv2.api.envoy.v12.http.service_control.ApiKeyRequirement
	(*MetricCost)(nil),        // 2: espv2.api.envoy.v12.http.service_control.MetricCost
	(*Requirement)(nil),       // 3: espv2.api.envoy.v12.http.service_control.Requirement
}
var file_api_envoy_v12_http_service_control_requirement_proto_depIdxs = []int32{
	0, // 0: espv2.api.envoy.v12.http.service_control.ApiKeyRequirement.locations:type_name -> espv2.api.envoy.v12.http.service_control.ApiKeyLocation
	1, // 1: espv2.api.envoy.v12.http.service_control.Requirement.api_key:type_name -> espv2.api.envoy.v12.http.service_control.ApiKeyRequirement
	2, // 2: espv2.api.envoy.v12.http.service_control.Requirement.metric_costs:type_name -> espv2.api.envoy.v12.http.service_control.MetricCost
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_envoy_v12_http_service_control_requirement_proto_init() }
func file_api_envoy_v12_http_service_control_requirement_proto_init() {
	if File_api_envoy_v12_http_service_control_requirement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyLocation); i {
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
		file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyRequirement); i {
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
		file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricCost); i {
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
		file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Requirement); i {
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
	file_api_envoy_v12_http_service_control_requirement_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ApiKeyLocation_Query)(nil),
		(*ApiKeyLocation_Header)(nil),
		(*ApiKeyLocation_Cookie)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_envoy_v12_http_service_control_requirement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_envoy_v12_http_service_control_requirement_proto_goTypes,
		DependencyIndexes: file_api_envoy_v12_http_service_control_requirement_proto_depIdxs,
		MessageInfos:      file_api_envoy_v12_http_service_control_requirement_proto_msgTypes,
	}.Build()
	File_api_envoy_v12_http_service_control_requirement_proto = out.File
	file_api_envoy_v12_http_service_control_requirement_proto_rawDesc = nil
	file_api_envoy_v12_http_service_control_requirement_proto_goTypes = nil
	file_api_envoy_v12_http_service_control_requirement_proto_depIdxs = nil
}