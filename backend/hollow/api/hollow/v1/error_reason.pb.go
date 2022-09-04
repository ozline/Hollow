// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: hollow/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_NORMAL_ERROR         ErrorReason = 0 // 常规错误
	ErrorReason_PARAMS_ILLEGAL       ErrorReason = 1 // 参数不合法
	ErrorReason_INFORMATION_ILLEGAL  ErrorReason = 2 // 数据不合法
	ErrorReason_INFOMATION_NOT_FOUND ErrorReason = 3 // 数据未找到
	ErrorReason_REPEATED_OPERATION   ErrorReason = 4 // 重复操作
	ErrorReason_ITEM_EXISTED         ErrorReason = 5 // 已经存在
	ErrorReason_USER_BLOCKED         ErrorReason = 6 // 用户被封禁
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "NORMAL_ERROR",
		1: "PARAMS_ILLEGAL",
		2: "INFORMATION_ILLEGAL",
		3: "INFOMATION_NOT_FOUND",
		4: "REPEATED_OPERATION",
		5: "ITEM_EXISTED",
		6: "USER_BLOCKED",
	}
	ErrorReason_value = map[string]int32{
		"NORMAL_ERROR":         0,
		"PARAMS_ILLEGAL":       1,
		"INFORMATION_ILLEGAL":  2,
		"INFOMATION_NOT_FOUND": 3,
		"REPEATED_OPERATION":   4,
		"ITEM_EXISTED":         5,
		"USER_BLOCKED":         6,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_hollow_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_hollow_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_hollow_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_hollow_v1_error_reason_proto protoreflect.FileDescriptor

var file_hollow_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x68, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa8,
	0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x0c, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x53, 0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47,
	0x41, 0x4c, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x49, 0x4e, 0x46, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x50, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x05, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45,
	0x44, 0x10, 0x06, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x34, 0x0a, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x17, 0x68, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0xa2, 0x02, 0x0b, 0x41, 0x50, 0x49, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hollow_v1_error_reason_proto_rawDescOnce sync.Once
	file_hollow_v1_error_reason_proto_rawDescData = file_hollow_v1_error_reason_proto_rawDesc
)

func file_hollow_v1_error_reason_proto_rawDescGZIP() []byte {
	file_hollow_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_hollow_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_hollow_v1_error_reason_proto_rawDescData)
	})
	return file_hollow_v1_error_reason_proto_rawDescData
}

var file_hollow_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hollow_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: errors.v1.ErrorReason
}
var file_hollow_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hollow_v1_error_reason_proto_init() }
func file_hollow_v1_error_reason_proto_init() {
	if File_hollow_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hollow_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hollow_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_hollow_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_hollow_v1_error_reason_proto_enumTypes,
	}.Build()
	File_hollow_v1_error_reason_proto = out.File
	file_hollow_v1_error_reason_proto_rawDesc = nil
	file_hollow_v1_error_reason_proto_goTypes = nil
	file_hollow_v1_error_reason_proto_depIdxs = nil
}
