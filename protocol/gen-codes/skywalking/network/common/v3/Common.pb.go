//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.1
// source: common/Common.proto

package v3

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

// In most cases, detect point should be `server` or `client`.
// Even in service mesh, this means `server`/`client` side sidecar
// `proxy` is reserved only.
type DetectPoint int32

const (
	DetectPoint_client DetectPoint = 0
	DetectPoint_server DetectPoint = 1
	DetectPoint_proxy  DetectPoint = 2
)

// Enum value maps for DetectPoint.
var (
	DetectPoint_name = map[int32]string{
		0: "client",
		1: "server",
		2: "proxy",
	}
	DetectPoint_value = map[string]int32{
		"client": 0,
		"server": 1,
		"proxy":  2,
	}
)

func (x DetectPoint) Enum() *DetectPoint {
	p := new(DetectPoint)
	*p = x
	return p
}

func (x DetectPoint) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DetectPoint) Descriptor() protoreflect.EnumDescriptor {
	return file_common_Common_proto_enumTypes[0].Descriptor()
}

func (DetectPoint) Type() protoreflect.EnumType {
	return &file_common_Common_proto_enumTypes[0]
}

func (x DetectPoint) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DetectPoint.Descriptor instead.
func (DetectPoint) EnumDescriptor() ([]byte, []int) {
	return file_common_Common_proto_rawDescGZIP(), []int{0}
}

type KeyStringValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyStringValuePair) Reset() {
	*x = KeyStringValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_Common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyStringValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyStringValuePair) ProtoMessage() {}

func (x *KeyStringValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_common_Common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyStringValuePair.ProtoReflect.Descriptor instead.
func (*KeyStringValuePair) Descriptor() ([]byte, []int) {
	return file_common_Common_proto_rawDescGZIP(), []int{0}
}

func (x *KeyStringValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyStringValuePair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CPU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsagePercent float64 `protobuf:"fixed64,2,opt,name=usagePercent,proto3" json:"usagePercent,omitempty"`
}

func (x *CPU) Reset() {
	*x = CPU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_Common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPU) ProtoMessage() {}

func (x *CPU) ProtoReflect() protoreflect.Message {
	mi := &file_common_Common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPU.ProtoReflect.Descriptor instead.
func (*CPU) Descriptor() ([]byte, []int) {
	return file_common_Common_proto_rawDescGZIP(), []int{1}
}

func (x *CPU) GetUsagePercent() float64 {
	if x != nil {
		return x.UsagePercent
	}
	return 0
}

type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []*Command `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_Common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_common_Common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_common_Common_proto_rawDescGZIP(), []int{2}
}

func (x *Commands) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string                `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args    []*KeyStringValuePair `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_Common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_common_Common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_common_Common_proto_rawDescGZIP(), []int{3}
}

func (x *Command) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Command) GetArgs() []*KeyStringValuePair {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_common_Common_proto protoreflect.FileDescriptor

var file_common_Common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x22, 0x3c, 0x0a, 0x12, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x29, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0c, 0x75, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a,
	0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x5a, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x2a, 0x30, 0x0a, 0x0b, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x10, 0x02, 0x42, 0x6d, 0x0a, 0x2b, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x50, 0x01, 0x5a, 0x1c, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0xaa, 0x02, 0x1d, 0x53, 0x6b, 0x79,
	0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_Common_proto_rawDescOnce sync.Once
	file_common_Common_proto_rawDescData = file_common_Common_proto_rawDesc
)

func file_common_Common_proto_rawDescGZIP() []byte {
	file_common_Common_proto_rawDescOnce.Do(func() {
		file_common_Common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_Common_proto_rawDescData)
	})
	return file_common_Common_proto_rawDescData
}

var file_common_Common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_Common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_Common_proto_goTypes = []interface{}{
	(DetectPoint)(0),           // 0: skywalking.v3.DetectPoint
	(*KeyStringValuePair)(nil), // 1: skywalking.v3.KeyStringValuePair
	(*CPU)(nil),                // 2: skywalking.v3.CPU
	(*Commands)(nil),           // 3: skywalking.v3.Commands
	(*Command)(nil),            // 4: skywalking.v3.Command
}
var file_common_Common_proto_depIdxs = []int32{
	4, // 0: skywalking.v3.Commands.commands:type_name -> skywalking.v3.Command
	1, // 1: skywalking.v3.Command.args:type_name -> skywalking.v3.KeyStringValuePair
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_Common_proto_init() }
func file_common_Common_proto_init() {
	if File_common_Common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_Common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyStringValuePair); i {
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
		file_common_Common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPU); i {
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
		file_common_Common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commands); i {
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
		file_common_Common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
			RawDescriptor: file_common_Common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_Common_proto_goTypes,
		DependencyIndexes: file_common_Common_proto_depIdxs,
		EnumInfos:         file_common_Common_proto_enumTypes,
		MessageInfos:      file_common_Common_proto_msgTypes,
	}.Build()
	File_common_Common_proto = out.File
	file_common_Common_proto_rawDesc = nil
	file_common_Common_proto_goTypes = nil
	file_common_Common_proto_depIdxs = nil
}
