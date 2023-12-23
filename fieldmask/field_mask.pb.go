// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v4.24.3
// source: field_mask.proto

package fieldmask

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Visibility int32

const (
	Visibility_VISIBILITY_UNSPECIFIED Visibility = 0
	Visibility_USER                   Visibility = 1
	Visibility_ADMIN                  Visibility = 2
	Visibility_USER_ADMIN             Visibility = 3
	Visibility_INNER                  Visibility = 4
	Visibility_USER_INNER             Visibility = 5
	Visibility_ADMIN_INNER            Visibility = 6
	Visibility_PUBLIC                 Visibility = 7
)

// Enum value maps for Visibility.
var (
	Visibility_name = map[int32]string{
		0: "VISIBILITY_UNSPECIFIED",
		1: "USER",
		2: "ADMIN",
		3: "USER_ADMIN",
		4: "INNER",
		5: "USER_INNER",
		6: "ADMIN_INNER",
		7: "PUBLIC",
	}
	Visibility_value = map[string]int32{
		"VISIBILITY_UNSPECIFIED": 0,
		"USER":                   1,
		"ADMIN":                  2,
		"USER_ADMIN":             3,
		"INNER":                  4,
		"USER_INNER":             5,
		"ADMIN_INNER":            6,
		"PUBLIC":                 7,
	}
)

func (x Visibility) Enum() *Visibility {
	p := new(Visibility)
	*p = x
	return p
}

func (x Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_field_mask_proto_enumTypes[0].Descriptor()
}

func (Visibility) Type() protoreflect.EnumType {
	return &file_field_mask_proto_enumTypes[0]
}

func (x Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Visibility.Descriptor instead.
func (Visibility) EnumDescriptor() ([]byte, []int) {
	return file_field_mask_proto_rawDescGZIP(), []int{0}
}

var file_field_mask_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Visibility)(nil),
		Field:         3430001,
		Name:          "hyfco.fieldmask.fv",
		Tag:           "varint,3430001,opt,name=fv,enum=hyfco.fieldmask.Visibility",
		Filename:      "field_mask.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional hyfco.fieldmask.Visibility fv = 3430001;
	E_Fv = &file_field_mask_proto_extTypes[0]
)

var File_field_mask_proto protoreflect.FileDescriptor

var file_field_mask_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x68, 0x79, 0x66, 0x63, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d,
	0x61, 0x73, 0x6b, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x85, 0x01, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41,
	0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x10,
	0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x10,
	0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x49, 0x4e, 0x4e, 0x45, 0x52,
	0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x07, 0x3a, 0x4d,
	0x0a, 0x02, 0x66, 0x76, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf1, 0xac, 0xd1, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x68,
	0x79, 0x66, 0x63, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x02, 0x66, 0x76, 0x42, 0x1c, 0x5a,
	0x1a, 0x68, 0x79, 0x66, 0x63, 0x6f, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b,
	0x3b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_field_mask_proto_rawDescOnce sync.Once
	file_field_mask_proto_rawDescData = file_field_mask_proto_rawDesc
)

func file_field_mask_proto_rawDescGZIP() []byte {
	file_field_mask_proto_rawDescOnce.Do(func() {
		file_field_mask_proto_rawDescData = protoimpl.X.CompressGZIP(file_field_mask_proto_rawDescData)
	})
	return file_field_mask_proto_rawDescData
}

var file_field_mask_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_field_mask_proto_goTypes = []interface{}{
	(Visibility)(0),                   // 0: hyfco.fieldmask.Visibility
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_field_mask_proto_depIdxs = []int32{
	1, // 0: hyfco.fieldmask.fv:extendee -> google.protobuf.FieldOptions
	0, // 1: hyfco.fieldmask.fv:type_name -> hyfco.fieldmask.Visibility
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_field_mask_proto_init() }
func file_field_mask_proto_init() {
	if File_field_mask_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_field_mask_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_field_mask_proto_goTypes,
		DependencyIndexes: file_field_mask_proto_depIdxs,
		EnumInfos:         file_field_mask_proto_enumTypes,
		ExtensionInfos:    file_field_mask_proto_extTypes,
	}.Build()
	File_field_mask_proto = out.File
	file_field_mask_proto_rawDesc = nil
	file_field_mask_proto_goTypes = nil
	file_field_mask_proto_depIdxs = nil
}