// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: auth_service.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_service_proto protoreflect.FileDescriptor

var file_auth_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x12, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcc,
	0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x13,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_auth_service_proto_goTypes = []any{
	(*LoginByEmailRequest)(nil),          // 0: auth.LoginByEmailRequest
	(*LoginByEmailSendCodeRequest)(nil),  // 1: auth.LoginByEmailSendCodeRequest
	(*LogoutRequest)(nil),                // 2: auth.LogoutRequest
	(*CheckRequest)(nil),                 // 3: auth.CheckRequest
	(*ParseRequest)(nil),                 // 4: auth.ParseRequest
	(*LoginByEmailResponse)(nil),         // 5: auth.LoginByEmailResponse
	(*LoginByEmailSendCodeResponse)(nil), // 6: auth.LoginByEmailSendCodeResponse
	(*LogoutResponse)(nil),               // 7: auth.LogoutResponse
	(*CheckResponse)(nil),                // 8: auth.CheckResponse
	(*ParseResponse)(nil),                // 9: auth.ParseResponse
}
var file_auth_service_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.LoginByEmail:input_type -> auth.LoginByEmailRequest
	1, // 1: auth.AuthService.LoginByEmailSendCode:input_type -> auth.LoginByEmailSendCodeRequest
	2, // 2: auth.AuthService.Logout:input_type -> auth.LogoutRequest
	3, // 3: auth.AuthService.Check:input_type -> auth.CheckRequest
	4, // 4: auth.AuthService.Parse:input_type -> auth.ParseRequest
	5, // 5: auth.AuthService.LoginByEmail:output_type -> auth.LoginByEmailResponse
	6, // 6: auth.AuthService.LoginByEmailSendCode:output_type -> auth.LoginByEmailSendCodeResponse
	7, // 7: auth.AuthService.Logout:output_type -> auth.LogoutResponse
	8, // 8: auth.AuthService.Check:output_type -> auth.CheckResponse
	9, // 9: auth.AuthService.Parse:output_type -> auth.ParseResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_service_proto_init() }
func file_auth_service_proto_init() {
	if File_auth_service_proto != nil {
		return
	}
	file_auth_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_service_proto_depIdxs,
	}.Build()
	File_auth_service_proto = out.File
	file_auth_service_proto_rawDesc = nil
	file_auth_service_proto_goTypes = nil
	file_auth_service_proto_depIdxs = nil
}