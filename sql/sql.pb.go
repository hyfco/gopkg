// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v4.24.3
// source: sql.proto

package sql

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimaryKey *Key    `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey,proto3,oneof" json:"primary_key,omitempty"`
	UniqueKeys []*Key  `protobuf:"bytes,2,rep,name=unique_keys,json=uniqueKeys,proto3" json:"unique_keys,omitempty"`
	Keys       []*Key  `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
	Table      *string `protobuf:"bytes,4,opt,name=table,proto3,oneof" json:"table,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_sql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_sql_proto_rawDescGZIP(), []int{0}
}

func (x *Table) GetPrimaryKey() *Key {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

func (x *Table) GetUniqueKeys() []*Key {
	if x != nil {
		return x.UniqueKeys
	}
	return nil
}

func (x *Table) GetKeys() []*Key {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Table) GetTable() string {
	if x != nil && x.Table != nil {
		return *x.Table
	}
	return ""
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F []string `protobuf:"bytes,1,rep,name=f,proto3" json:"f,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_sql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_sql_proto_rawDescGZIP(), []int{1}
}

func (x *Key) GetF() []string {
	if x != nil {
		return x.F
	}
	return nil
}

var file_sql_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Table)(nil),
		Field:         82100006,
		Name:          "hyfco.sql.table",
		Tag:           "bytes,82100006,opt,name=table",
		Filename:      "sql.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional hyfco.sql.Table table = 82100006;
	E_Table = &file_sql_proto_extTypes[0]
)

var File_sql_proto protoreflect.FileDescriptor

var file_sql_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x79, 0x66,
	0x63, 0x6f, 0x2e, 0x73, 0x71, 0x6c, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x05, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x79, 0x66, 0x63, 0x6f, 0x2e,
	0x73, 0x71, 0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0b, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x68, 0x79, 0x66, 0x63, 0x6f, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x79, 0x66, 0x63, 0x6f, 0x2e,
	0x73, 0x71, 0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x22, 0x13, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x66, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x01, 0x66, 0x3a, 0x4d, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xa6, 0xfe, 0x92, 0x27, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x79, 0x66,
	0x63, 0x6f, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x68, 0x79, 0x66, 0x63, 0x6f, 0x2f,
	0x67, 0x6f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x71, 0x6c, 0x3b, 0x73, 0x71, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sql_proto_rawDescOnce sync.Once
	file_sql_proto_rawDescData = file_sql_proto_rawDesc
)

func file_sql_proto_rawDescGZIP() []byte {
	file_sql_proto_rawDescOnce.Do(func() {
		file_sql_proto_rawDescData = protoimpl.X.CompressGZIP(file_sql_proto_rawDescData)
	})
	return file_sql_proto_rawDescData
}

var file_sql_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sql_proto_goTypes = []interface{}{
	(*Table)(nil),                       // 0: hyfco.sql.Table
	(*Key)(nil),                         // 1: hyfco.sql.Key
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
}
var file_sql_proto_depIdxs = []int32{
	1, // 0: hyfco.sql.Table.primary_key:type_name -> hyfco.sql.Key
	1, // 1: hyfco.sql.Table.unique_keys:type_name -> hyfco.sql.Key
	1, // 2: hyfco.sql.Table.keys:type_name -> hyfco.sql.Key
	2, // 3: hyfco.sql.table:extendee -> google.protobuf.MessageOptions
	0, // 4: hyfco.sql.table:type_name -> hyfco.sql.Table
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sql_proto_init() }
func file_sql_proto_init() {
	if File_sql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_sql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
	file_sql_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_sql_proto_goTypes,
		DependencyIndexes: file_sql_proto_depIdxs,
		MessageInfos:      file_sql_proto_msgTypes,
		ExtensionInfos:    file_sql_proto_extTypes,
	}.Build()
	File_sql_proto = out.File
	file_sql_proto_rawDesc = nil
	file_sql_proto_goTypes = nil
	file_sql_proto_depIdxs = nil
}
