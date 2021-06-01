// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: ncbi/datasets/v1alpha1/openapi_options.proto

package v1alpha1

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type Tier int32

const (
	Tier_TIER_UNSPECIFIED      Tier = 0
	Tier_TIER_TIER3            Tier = 1
	Tier_TIER_TIER4            Tier = 2
	Tier_TIER_TIER4_DEPRECATED Tier = 3
)

// Enum value maps for Tier.
var (
	Tier_name = map[int32]string{
		0: "TIER_UNSPECIFIED",
		1: "TIER_TIER3",
		2: "TIER_TIER4",
		3: "TIER_TIER4_DEPRECATED",
	}
	Tier_value = map[string]int32{
		"TIER_UNSPECIFIED":      0,
		"TIER_TIER3":            1,
		"TIER_TIER4":            2,
		"TIER_TIER4_DEPRECATED": 3,
	}
)

func (x Tier) Enum() *Tier {
	p := new(Tier)
	*p = x
	return p
}

func (x Tier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tier) Descriptor() protoreflect.EnumDescriptor {
	return file_ncbi_datasets_v1alpha1_openapi_options_proto_enumTypes[0].Descriptor()
}

func (Tier) Type() protoreflect.EnumType {
	return &file_ncbi_datasets_v1alpha1_openapi_options_proto_enumTypes[0]
}

func (x Tier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tier.Descriptor instead.
func (Tier) EnumDescriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescGZIP(), []int{0}
}

type H1GatewayMethodOptions_Method int32

const (
	H1GatewayMethodOptions_METHOD_UNSPECIFIED H1GatewayMethodOptions_Method = 0
	H1GatewayMethodOptions_METHOD_GET         H1GatewayMethodOptions_Method = 1
	H1GatewayMethodOptions_METHOD_POST        H1GatewayMethodOptions_Method = 2
)

// Enum value maps for H1GatewayMethodOptions_Method.
var (
	H1GatewayMethodOptions_Method_name = map[int32]string{
		0: "METHOD_UNSPECIFIED",
		1: "METHOD_GET",
		2: "METHOD_POST",
	}
	H1GatewayMethodOptions_Method_value = map[string]int32{
		"METHOD_UNSPECIFIED": 0,
		"METHOD_GET":         1,
		"METHOD_POST":        2,
	}
)

func (x H1GatewayMethodOptions_Method) Enum() *H1GatewayMethodOptions_Method {
	p := new(H1GatewayMethodOptions_Method)
	*p = x
	return p
}

func (x H1GatewayMethodOptions_Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (H1GatewayMethodOptions_Method) Descriptor() protoreflect.EnumDescriptor {
	return file_ncbi_datasets_v1alpha1_openapi_options_proto_enumTypes[1].Descriptor()
}

func (H1GatewayMethodOptions_Method) Type() protoreflect.EnumType {
	return &file_ncbi_datasets_v1alpha1_openapi_options_proto_enumTypes[1]
}

func (x H1GatewayMethodOptions_Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use H1GatewayMethodOptions_Method.Descriptor instead.
func (H1GatewayMethodOptions_Method) EnumDescriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescGZIP(), []int{1, 0}
}

// Field Options related to OpenapiDB fields
type OpenapiFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elide []Tier `protobuf:"varint,1,rep,packed,name=elide,proto3,enum=ncbi.datasets.v1alpha1.Tier" json:"elide,omitempty"`
}

func (x *OpenapiFieldOptions) Reset() {
	*x = OpenapiFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenapiFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenapiFieldOptions) ProtoMessage() {}

func (x *OpenapiFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenapiFieldOptions.ProtoReflect.Descriptor instead.
func (*OpenapiFieldOptions) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescGZIP(), []int{0}
}

func (x *OpenapiFieldOptions) GetElide() []Tier {
	if x != nil {
		return x.Elide
	}
	return nil
}

// RPC method options related to the HTTP1 Gateway
type H1GatewayMethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*H1GatewayMethodOptions_Url `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	Tier Tier                          `protobuf:"varint,2,opt,name=tier,proto3,enum=ncbi.datasets.v1alpha1.Tier" json:"tier,omitempty"`
}

func (x *H1GatewayMethodOptions) Reset() {
	*x = H1GatewayMethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H1GatewayMethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H1GatewayMethodOptions) ProtoMessage() {}

func (x *H1GatewayMethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H1GatewayMethodOptions.ProtoReflect.Descriptor instead.
func (*H1GatewayMethodOptions) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescGZIP(), []int{1}
}

func (x *H1GatewayMethodOptions) GetUrls() []*H1GatewayMethodOptions_Url {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *H1GatewayMethodOptions) GetTier() Tier {
	if x != nil {
		return x.Tier
	}
	return Tier_TIER_UNSPECIFIED
}

type H1GatewayMethodOptions_Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method H1GatewayMethodOptions_Method `protobuf:"varint,1,opt,name=method,proto3,enum=ncbi.datasets.v1alpha1.H1GatewayMethodOptions_Method" json:"method,omitempty"`
	Url    string                        `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *H1GatewayMethodOptions_Url) Reset() {
	*x = H1GatewayMethodOptions_Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H1GatewayMethodOptions_Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H1GatewayMethodOptions_Url) ProtoMessage() {}

func (x *H1GatewayMethodOptions_Url) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H1GatewayMethodOptions_Url.ProtoReflect.Descriptor instead.
func (*H1GatewayMethodOptions_Url) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescGZIP(), []int{1, 0}
}

func (x *H1GatewayMethodOptions_Url) GetMethod() H1GatewayMethodOptions_Method {
	if x != nil {
		return x.Method
	}
	return H1GatewayMethodOptions_METHOD_UNSPECIFIED
}

func (x *H1GatewayMethodOptions_Url) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var file_ncbi_datasets_v1alpha1_openapi_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*OpenapiFieldOptions)(nil),
		Field:         53002,
		Name:          "ncbi.datasets.v1alpha1.openapi",
		Tag:           "bytes,53002,opt,name=openapi",
		Filename:      "ncbi/datasets/v1alpha1/openapi_options.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*H1GatewayMethodOptions)(nil),
		Field:         53001,
		Name:          "ncbi.datasets.v1alpha1.gateway",
		Tag:           "bytes,53001,opt,name=gateway",
		Filename:      "ncbi/datasets/v1alpha1/openapi_options.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional ncbi.datasets.v1alpha1.OpenapiFieldOptions openapi = 53002;
	E_Openapi = &file_ncbi_datasets_v1alpha1_openapi_options_proto_extTypes[0]
)

// Extension fields to descriptor.MethodOptions.
var (
	// optional ncbi.datasets.v1alpha1.H1GatewayMethodOptions gateway = 53001;
	E_Gateway = &file_ncbi_datasets_v1alpha1_openapi_options_proto_extTypes[1]
)

var File_ncbi_datasets_v1alpha1_openapi_options_proto protoreflect.FileDescriptor

var file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x32, 0x0a, 0x05, 0x65, 0x6c, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x52, 0x05, 0x65, 0x6c,
	0x69, 0x64, 0x65, 0x22, 0xbd, 0x02, 0x0a, 0x16, 0x48, 0x31, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46,
	0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6e,
	0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x31, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x72, 0x6c,
	0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x69,
	0x65, 0x72, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x1a, 0x66, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12,
	0x4d, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x35, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x31, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x41, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x47, 0x45, 0x54,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x50, 0x4f, 0x53,
	0x54, 0x10, 0x02, 0x2a, 0x57, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x49, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x33, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x34, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x34, 0x5f,
	0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x3a, 0x66, 0x0a, 0x07,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x3a, 0x6a, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x89, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x48, 0x31, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x42, 0x1b, 0x5a, 0x16, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescOnce sync.Once
	file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescData = file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDesc
)

func file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescGZIP() []byte {
	file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescOnce.Do(func() {
		file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescData)
	})
	return file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDescData
}

var file_ncbi_datasets_v1alpha1_openapi_options_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ncbi_datasets_v1alpha1_openapi_options_proto_goTypes = []interface{}{
	(Tier)(0),                          // 0: ncbi.datasets.v1alpha1.Tier
	(H1GatewayMethodOptions_Method)(0), // 1: ncbi.datasets.v1alpha1.H1GatewayMethodOptions.Method
	(*OpenapiFieldOptions)(nil),        // 2: ncbi.datasets.v1alpha1.OpenapiFieldOptions
	(*H1GatewayMethodOptions)(nil),     // 3: ncbi.datasets.v1alpha1.H1GatewayMethodOptions
	(*H1GatewayMethodOptions_Url)(nil), // 4: ncbi.datasets.v1alpha1.H1GatewayMethodOptions.Url
	(*descriptor.FieldOptions)(nil),    // 5: google.protobuf.FieldOptions
	(*descriptor.MethodOptions)(nil),   // 6: google.protobuf.MethodOptions
}
var file_ncbi_datasets_v1alpha1_openapi_options_proto_depIdxs = []int32{
	0, // 0: ncbi.datasets.v1alpha1.OpenapiFieldOptions.elide:type_name -> ncbi.datasets.v1alpha1.Tier
	4, // 1: ncbi.datasets.v1alpha1.H1GatewayMethodOptions.urls:type_name -> ncbi.datasets.v1alpha1.H1GatewayMethodOptions.Url
	0, // 2: ncbi.datasets.v1alpha1.H1GatewayMethodOptions.tier:type_name -> ncbi.datasets.v1alpha1.Tier
	1, // 3: ncbi.datasets.v1alpha1.H1GatewayMethodOptions.Url.method:type_name -> ncbi.datasets.v1alpha1.H1GatewayMethodOptions.Method
	5, // 4: ncbi.datasets.v1alpha1.openapi:extendee -> google.protobuf.FieldOptions
	6, // 5: ncbi.datasets.v1alpha1.gateway:extendee -> google.protobuf.MethodOptions
	2, // 6: ncbi.datasets.v1alpha1.openapi:type_name -> ncbi.datasets.v1alpha1.OpenapiFieldOptions
	3, // 7: ncbi.datasets.v1alpha1.gateway:type_name -> ncbi.datasets.v1alpha1.H1GatewayMethodOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	6, // [6:8] is the sub-list for extension type_name
	4, // [4:6] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ncbi_datasets_v1alpha1_openapi_options_proto_init() }
func file_ncbi_datasets_v1alpha1_openapi_options_proto_init() {
	if File_ncbi_datasets_v1alpha1_openapi_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenapiFieldOptions); i {
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
		file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H1GatewayMethodOptions); i {
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
		file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H1GatewayMethodOptions_Url); i {
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
			RawDescriptor: file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_ncbi_datasets_v1alpha1_openapi_options_proto_goTypes,
		DependencyIndexes: file_ncbi_datasets_v1alpha1_openapi_options_proto_depIdxs,
		EnumInfos:         file_ncbi_datasets_v1alpha1_openapi_options_proto_enumTypes,
		MessageInfos:      file_ncbi_datasets_v1alpha1_openapi_options_proto_msgTypes,
		ExtensionInfos:    file_ncbi_datasets_v1alpha1_openapi_options_proto_extTypes,
	}.Build()
	File_ncbi_datasets_v1alpha1_openapi_options_proto = out.File
	file_ncbi_datasets_v1alpha1_openapi_options_proto_rawDesc = nil
	file_ncbi_datasets_v1alpha1_openapi_options_proto_goTypes = nil
	file_ncbi_datasets_v1alpha1_openapi_options_proto_depIdxs = nil
}
