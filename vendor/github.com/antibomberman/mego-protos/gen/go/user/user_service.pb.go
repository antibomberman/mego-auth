// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: user_service.proto

package user

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

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xee,
	0x05, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2c, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0b, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x08, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x7c, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0c, 0x5a,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_user_service_proto_goTypes = []any{
	(*FindUserRequest)(nil),             // 0: user.FindUserRequest
	(*Id)(nil),                          // 1: user.Id
	(*Email)(nil),                       // 2: user.Email
	(*Phone)(nil),                       // 3: user.Phone
	(*Token)(nil),                       // 4: user.Token
	(*CreateUserRequest)(nil),           // 5: user.CreateUserRequest
	(*UpdateUserRequest)(nil),           // 6: user.UpdateUserRequest
	(*UpdateProfileRequest)(nil),        // 7: user.UpdateProfileRequest
	(*UpdateLangRequest)(nil),           // 8: user.UpdateLangRequest
	(*UpdateThemeRequest)(nil),          // 9: user.UpdateThemeRequest
	(*UpdateEmailRequest)(nil),          // 10: user.UpdateEmailRequest
	(*UpdateEmailSendCodeRequest)(nil),  // 11: user.UpdateEmailSendCodeRequest
	(*Empty)(nil),                       // 12: user.Empty
	(*FindUserResponse)(nil),            // 13: user.FindUserResponse
	(*UserDetails)(nil),                 // 14: user.UserDetails
	(*UpdateLangResponse)(nil),          // 15: user.UpdateLangResponse
	(*UpdateThemeResponse)(nil),         // 16: user.UpdateThemeResponse
	(*UpdateEmailResponse)(nil),         // 17: user.UpdateEmailResponse
	(*UpdateEmailSendCodeResponse)(nil), // 18: user.UpdateEmailSendCodeResponse
}
var file_user_service_proto_depIdxs = []int32{
	0,  // 0: user.UserService.Find:input_type -> user.FindUserRequest
	1,  // 1: user.UserService.GetById:input_type -> user.Id
	2,  // 2: user.UserService.GetByEmail:input_type -> user.Email
	3,  // 3: user.UserService.GetByPhone:input_type -> user.Phone
	4,  // 4: user.UserService.GetByToken:input_type -> user.Token
	5,  // 5: user.UserService.Create:input_type -> user.CreateUserRequest
	6,  // 6: user.UserService.Update:input_type -> user.UpdateUserRequest
	1,  // 7: user.UserService.Delete:input_type -> user.Id
	7,  // 8: user.UserService.UpdateProfile:input_type -> user.UpdateProfileRequest
	8,  // 9: user.UserService.UpdateLang:input_type -> user.UpdateLangRequest
	9,  // 10: user.UserService.UpdateTheme:input_type -> user.UpdateThemeRequest
	10, // 11: user.UserService.UpdateEmail:input_type -> user.UpdateEmailRequest
	11, // 12: user.UserService.UpdateEmailSendCode:input_type -> user.UpdateEmailSendCodeRequest
	12, // 13: user.UserRoleService.Roles:input_type -> user.Empty
	12, // 14: user.UserRoleService.Create:input_type -> user.Empty
	12, // 15: user.UserRoleService.Delete:input_type -> user.Empty
	13, // 16: user.UserService.Find:output_type -> user.FindUserResponse
	14, // 17: user.UserService.GetById:output_type -> user.UserDetails
	14, // 18: user.UserService.GetByEmail:output_type -> user.UserDetails
	14, // 19: user.UserService.GetByPhone:output_type -> user.UserDetails
	14, // 20: user.UserService.GetByToken:output_type -> user.UserDetails
	14, // 21: user.UserService.Create:output_type -> user.UserDetails
	14, // 22: user.UserService.Update:output_type -> user.UserDetails
	14, // 23: user.UserService.Delete:output_type -> user.UserDetails
	14, // 24: user.UserService.UpdateProfile:output_type -> user.UserDetails
	15, // 25: user.UserService.UpdateLang:output_type -> user.UpdateLangResponse
	16, // 26: user.UserService.UpdateTheme:output_type -> user.UpdateThemeResponse
	17, // 27: user.UserService.UpdateEmail:output_type -> user.UpdateEmailResponse
	18, // 28: user.UserService.UpdateEmailSendCode:output_type -> user.UpdateEmailSendCodeResponse
	12, // 29: user.UserRoleService.Roles:output_type -> user.Empty
	12, // 30: user.UserRoleService.Create:output_type -> user.Empty
	12, // 31: user.UserRoleService.Delete:output_type -> user.Empty
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	file_user_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}