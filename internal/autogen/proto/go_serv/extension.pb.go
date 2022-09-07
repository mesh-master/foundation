// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: go_serv/extension.proto

package go_serv

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_go_serv_extension_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60000,
		Name:          "go_serv.require_roles",
		Tag:           "bytes,60000,opt,name=require_roles",
		Filename:      "go_serv/extension.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         60100,
		Name:          "go_serv.new_insecure_session",
		Tag:           "varint,60100,opt,name=new_insecure_session",
		Filename:      "go_serv/extension.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60101,
		Name:          "go_serv.require_session",
		Tag:           "varint,60101,opt,name=require_session",
		Filename:      "go_serv/extension.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60102,
		Name:          "go_serv.optional_session",
		Tag:           "varint,60102,opt,name=optional_session",
		Filename:      "go_serv/extension.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60103,
		Name:          "go_serv.close_session",
		Tag:           "varint,60103,opt,name=close_session",
		Filename:      "go_serv/extension.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60110,
		Name:          "go_serv.copy_meta_off",
		Tag:           "varint,60110,opt,name=copy_meta_off",
		Filename:      "go_serv/extension.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60100,
		Name:          "go_serv.local_shm_ipc",
		Tag:           "varint,60100,opt,name=local_shm_ipc",
		Filename:      "go_serv/extension.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60101,
		Name:          "go_serv.enc_off",
		Tag:           "varint,60101,opt,name=enc_off",
		Filename:      "go_serv/extension.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional string require_roles = 60000;
	E_RequireRoles = &file_go_serv_extension_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// Open a new session with the specified session lifetime in seconds
	//
	// optional int32 new_insecure_session = 60100;
	E_NewInsecureSession = &file_go_serv_extension_proto_extTypes[1]
	// If set to true a session ID must be provided.
	//
	// optional bool require_session = 60101;
	E_RequireSession = &file_go_serv_extension_proto_extTypes[2]
	// If set to true a session either was created before with NetParcel.SecureSession or will be created during the call
	//
	// optional bool optional_session = 60102;
	E_OptionalSession = &file_go_serv_extension_proto_extTypes[3]
	// Ask server to close an opened session
	//
	// optional bool close_session = 60103;
	E_CloseSession = &file_go_serv_extension_proto_extTypes[4]
	//
	//
	// optional bool copy_meta_off = 60110;
	E_CopyMetaOff = &file_go_serv_extension_proto_extTypes[5]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// If set to true enables shared memory IPC for gRPC message transmission.
	// Should be used only for large messages.
	//
	// optional bool local_shm_ipc = 60100;
	E_LocalShmIpc = &file_go_serv_extension_proto_extTypes[6]
	// If set to true disables message encryption for the secure sessions.
	//
	// optional bool enc_off = 60101;
	E_EncOff = &file_go_serv_extension_proto_extTypes[7]
)

var File_go_serv_extension_proto protoreflect.FileDescriptor

var file_go_serv_extension_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x6f, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x46, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe0, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x3a, 0x52, 0x0a, 0x14,
	0x6e, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4, 0xd5, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6e, 0x65,
	0x77, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x49, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xc5, 0xd5, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x4b, 0x0a, 0x10, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xc6, 0xd5, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x45, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc7, 0xd5, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a,
	0x44, 0x0a, 0x0d, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6f, 0x66, 0x66,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xce, 0xd5, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63, 0x6f, 0x70, 0x79, 0x4d, 0x65,
	0x74, 0x61, 0x4f, 0x66, 0x66, 0x3a, 0x45, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73,
	0x68, 0x6d, 0x5f, 0x69, 0x70, 0x63, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4, 0xd5, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x68, 0x6d, 0x49, 0x70, 0x63, 0x3a, 0x3a, 0x0a, 0x07,
	0x65, 0x6e, 0x63, 0x5f, 0x6f, 0x66, 0x66, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc5, 0xd5, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x65, 0x6e, 0x63, 0x4f, 0x66, 0x66, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_go_serv_extension_proto_goTypes = []interface{}{
	(*descriptorpb.ServiceOptions)(nil), // 0: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
}
var file_go_serv_extension_proto_depIdxs = []int32{
	0, // 0: go_serv.require_roles:extendee -> google.protobuf.ServiceOptions
	1, // 1: go_serv.new_insecure_session:extendee -> google.protobuf.MethodOptions
	1, // 2: go_serv.require_session:extendee -> google.protobuf.MethodOptions
	1, // 3: go_serv.optional_session:extendee -> google.protobuf.MethodOptions
	1, // 4: go_serv.close_session:extendee -> google.protobuf.MethodOptions
	1, // 5: go_serv.copy_meta_off:extendee -> google.protobuf.MethodOptions
	2, // 6: go_serv.local_shm_ipc:extendee -> google.protobuf.MessageOptions
	2, // 7: go_serv.enc_off:extendee -> google.protobuf.MessageOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	0, // [0:8] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_serv_extension_proto_init() }
func file_go_serv_extension_proto_init() {
	if File_go_serv_extension_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_serv_extension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 8,
			NumServices:   0,
		},
		GoTypes:           file_go_serv_extension_proto_goTypes,
		DependencyIndexes: file_go_serv_extension_proto_depIdxs,
		ExtensionInfos:    file_go_serv_extension_proto_extTypes,
	}.Build()
	File_go_serv_extension_proto = out.File
	file_go_serv_extension_proto_rawDesc = nil
	file_go_serv_extension_proto_goTypes = nil
	file_go_serv_extension_proto_depIdxs = nil
}
