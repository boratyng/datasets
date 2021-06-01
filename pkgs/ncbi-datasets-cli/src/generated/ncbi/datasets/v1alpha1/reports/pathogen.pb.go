// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: ncbi/datasets/v1alpha1/reports/pathogen.proto

package reports

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

type PathogenFastaDefline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organism *Organism `protobuf:"bytes,1,opt,name=organism,proto3" json:"organism,omitempty"`
	// begin-end
	ElementRange  string `protobuf:"bytes,2,opt,name=element_range,json=elementRange,proto3" json:"element_range,omitempty"`
	ElementName   string `protobuf:"bytes,3,opt,name=element_name,json=elementName,proto3" json:"element_name,omitempty"`
	ElementSymbol string `protobuf:"bytes,4,opt,name=element_symbol,json=elementSymbol,proto3" json:"element_symbol,omitempty"`
	Contig        string `protobuf:"bytes,5,opt,name=contig,proto3" json:"contig,omitempty"`
}

func (x *PathogenFastaDefline) Reset() {
	*x = PathogenFastaDefline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_reports_pathogen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathogenFastaDefline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathogenFastaDefline) ProtoMessage() {}

func (x *PathogenFastaDefline) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_reports_pathogen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathogenFastaDefline.ProtoReflect.Descriptor instead.
func (*PathogenFastaDefline) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDescGZIP(), []int{0}
}

func (x *PathogenFastaDefline) GetOrganism() *Organism {
	if x != nil {
		return x.Organism
	}
	return nil
}

func (x *PathogenFastaDefline) GetElementRange() string {
	if x != nil {
		return x.ElementRange
	}
	return ""
}

func (x *PathogenFastaDefline) GetElementName() string {
	if x != nil {
		return x.ElementName
	}
	return ""
}

func (x *PathogenFastaDefline) GetElementSymbol() string {
	if x != nil {
		return x.ElementSymbol
	}
	return ""
}

func (x *PathogenFastaDefline) GetContig() string {
	if x != nil {
		return x.Contig
	}
	return ""
}

var File_ncbi_datasets_v1alpha1_reports_pathogen_proto protoreflect.FileDescriptor

var file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2f, 0x70, 0x61, 0x74, 0x68, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a,
	0x2b, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6e, 0x63,
	0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x04, 0x0a, 0x14, 0x50,
	0x61, 0x74, 0x68, 0x6f, 0x67, 0x65, 0x6e, 0x46, 0x61, 0x73, 0x74, 0x61, 0x44, 0x65, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x52,
	0x08, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x78, 0x0a, 0x0d, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x53, 0xca, 0xf3, 0x18, 0x1f, 0x08, 0x01, 0x12, 0x1b, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x20,
	0x74, 0x6f, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0xd2, 0xf3, 0x18, 0x2c, 0x0a, 0x0d, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x20,
	0x74, 0x6f, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x6d, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4a, 0xca, 0xf3, 0x18, 0x1b, 0x08,
	0x02, 0x12, 0x17, 0x54, 0x68, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xd2, 0xf3, 0x18, 0x27, 0x0a, 0x0c,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x54, 0x68,
	0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x77, 0x0a, 0x0e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x50, 0xca, 0xf3, 0x18, 0x1d,
	0x08, 0x03, 0x12, 0x19, 0x54, 0x68, 0x65, 0x20, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xd2, 0xf3, 0x18,
	0x2b, 0x0a, 0x0e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x19, 0x54, 0x68, 0x65, 0x20, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x58, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x40, 0xca, 0xf3, 0x18,
	0x14, 0x08, 0x04, 0x12, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x20, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0xd2, 0xf3, 0x18, 0x24, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x69,
	0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x43, 0x6f, 0x6e,
	0x74, 0x69, 0x67, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x67, 0x42, 0x23, 0x5a, 0x1e, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDescOnce sync.Once
	file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDescData = file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDesc
)

func file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDescGZIP() []byte {
	file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDescOnce.Do(func() {
		file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDescData = protoimpl.X.CompressGZIP(file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDescData)
	})
	return file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDescData
}

var file_ncbi_datasets_v1alpha1_reports_pathogen_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ncbi_datasets_v1alpha1_reports_pathogen_proto_goTypes = []interface{}{
	(*PathogenFastaDefline)(nil), // 0: ncbi.datasets.v1alpha1.reports.PathogenFastaDefline
	(*Organism)(nil),             // 1: ncbi.datasets.v1alpha1.reports.Organism
}
var file_ncbi_datasets_v1alpha1_reports_pathogen_proto_depIdxs = []int32{
	1, // 0: ncbi.datasets.v1alpha1.reports.PathogenFastaDefline.organism:type_name -> ncbi.datasets.v1alpha1.reports.Organism
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ncbi_datasets_v1alpha1_reports_pathogen_proto_init() }
func file_ncbi_datasets_v1alpha1_reports_pathogen_proto_init() {
	if File_ncbi_datasets_v1alpha1_reports_pathogen_proto != nil {
		return
	}
	file_ncbi_datasets_v1alpha1_reports_common_proto_init()
	file_ncbi_datasets_v1alpha1_reports_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ncbi_datasets_v1alpha1_reports_pathogen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathogenFastaDefline); i {
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
			RawDescriptor: file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ncbi_datasets_v1alpha1_reports_pathogen_proto_goTypes,
		DependencyIndexes: file_ncbi_datasets_v1alpha1_reports_pathogen_proto_depIdxs,
		MessageInfos:      file_ncbi_datasets_v1alpha1_reports_pathogen_proto_msgTypes,
	}.Build()
	File_ncbi_datasets_v1alpha1_reports_pathogen_proto = out.File
	file_ncbi_datasets_v1alpha1_reports_pathogen_proto_rawDesc = nil
	file_ncbi_datasets_v1alpha1_reports_pathogen_proto_goTypes = nil
	file_ncbi_datasets_v1alpha1_reports_pathogen_proto_depIdxs = nil
}
