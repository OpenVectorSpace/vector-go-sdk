// Copyright (c) 2018 Anki, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License in the file LICENSE.txt or at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Add a decorator for streamed message handling in app

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.4
// source: extensions.proto

package vectorpb

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var file_extensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60000,
		Name:          "Anki.Vector.external_interface.streamed",
		Tag:           "varint,60000,opt,name=streamed",
		Filename:      "extensions.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional bool streamed = 60000;
	E_Streamed = &file_extensions_proto_extTypes[0]
)

var File_extensions_proto protoreflect.FileDescriptor

var file_extensions_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x41, 0x6e, 0x6b, 0x69, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3d, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xe0, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x64, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x2d,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_extensions_proto_goTypes = []interface{}{
	(*descriptor.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
}
var file_extensions_proto_depIdxs = []int32{
	0, // 0: Anki.Vector.external_interface.streamed:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_extensions_proto_init() }
func file_extensions_proto_init() {
	if File_extensions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_extensions_proto_goTypes,
		DependencyIndexes: file_extensions_proto_depIdxs,
		ExtensionInfos:    file_extensions_proto_extTypes,
	}.Build()
	File_extensions_proto = out.File
	file_extensions_proto_rawDesc = nil
	file_extensions_proto_goTypes = nil
	file_extensions_proto_depIdxs = nil
}