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
// source: logging/Logging.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v3 "skywalking/network/common/v3"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Log data is collected through file scratcher of agent.
// Natively, Satellite provides various ways to collect logs.
type LogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [Optional] The timestamp of the log, in millisecond.
	// If not set, OAP server would use the received timestamp as log's timestamp, or relies on the OAP server analyzer.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// [Required] **Service**. Represents a set/group of workloads which provide the same behaviours for incoming requests.
	//
	// The logic name represents the service. This would show as a separate node in the topology.
	// The metrics analyzed from the spans, would be aggregated for this entity as the service level.
	//
	// If this is not the first element of the streaming, use the previous not-null name as the service name.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// [Optional] **Service Instance**. Each individual workload in the Service group is known as an instance. Like `pods` in Kubernetes, it
	// doesn't need to be a single OS process, however, if you are using instrument agents, an instance is actually a real OS process.
	//
	// The logic name represents the service instance. This would show as a separate node in the instance relationship.
	// The metrics analyzed from the spans, would be aggregated for this entity as the service instance level.
	ServiceInstance string `protobuf:"bytes,3,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	// [Optional] **Endpoint**. A path in a service for incoming requests, such as an HTTP URI path or a gRPC service class + method signature.
	//
	// The logic name represents the endpoint, which logs belong.
	Endpoint string `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// [Required] The content of the log.
	Body *LogDataBody `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// [Optional] Logs with trace context
	TraceContext *TraceContext `protobuf:"bytes,6,opt,name=traceContext,proto3" json:"traceContext,omitempty"`
	// [Optional] The available tags. OAP server could provide search/analysis capabilities base on these.
	Tags *LogTags `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *LogData) Reset() {
	*x = LogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_Logging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogData) ProtoMessage() {}

func (x *LogData) ProtoReflect() protoreflect.Message {
	mi := &file_logging_Logging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogData.ProtoReflect.Descriptor instead.
func (*LogData) Descriptor() ([]byte, []int) {
	return file_logging_Logging_proto_rawDescGZIP(), []int{0}
}

func (x *LogData) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LogData) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *LogData) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

func (x *LogData) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *LogData) GetBody() *LogDataBody {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *LogData) GetTraceContext() *TraceContext {
	if x != nil {
		return x.TraceContext
	}
	return nil
}

func (x *LogData) GetTags() *LogTags {
	if x != nil {
		return x.Tags
	}
	return nil
}

// The content of the log data
type LogDataBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A type to match analyzer(s) at the OAP server.
	// The data could be analysis at the client side, but could be partial
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Content with extendable format.
	//
	// Types that are assignable to Content:
	//	*LogDataBody_Text
	//	*LogDataBody_Json
	//	*LogDataBody_Yaml
	Content isLogDataBody_Content `protobuf_oneof:"content"`
}

func (x *LogDataBody) Reset() {
	*x = LogDataBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_Logging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogDataBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogDataBody) ProtoMessage() {}

func (x *LogDataBody) ProtoReflect() protoreflect.Message {
	mi := &file_logging_Logging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogDataBody.ProtoReflect.Descriptor instead.
func (*LogDataBody) Descriptor() ([]byte, []int) {
	return file_logging_Logging_proto_rawDescGZIP(), []int{1}
}

func (x *LogDataBody) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (m *LogDataBody) GetContent() isLogDataBody_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *LogDataBody) GetText() *TextLog {
	if x, ok := x.GetContent().(*LogDataBody_Text); ok {
		return x.Text
	}
	return nil
}

func (x *LogDataBody) GetJson() *JSONLog {
	if x, ok := x.GetContent().(*LogDataBody_Json); ok {
		return x.Json
	}
	return nil
}

func (x *LogDataBody) GetYaml() *YAMLLog {
	if x, ok := x.GetContent().(*LogDataBody_Yaml); ok {
		return x.Yaml
	}
	return nil
}

type isLogDataBody_Content interface {
	isLogDataBody_Content()
}

type LogDataBody_Text struct {
	Text *TextLog `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

type LogDataBody_Json struct {
	Json *JSONLog `protobuf:"bytes,3,opt,name=json,proto3,oneof"`
}

type LogDataBody_Yaml struct {
	Yaml *YAMLLog `protobuf:"bytes,4,opt,name=yaml,proto3,oneof"`
}

func (*LogDataBody_Text) isLogDataBody_Content() {}

func (*LogDataBody_Json) isLogDataBody_Content() {}

func (*LogDataBody_Yaml) isLogDataBody_Content() {}

// Literal text log, typically requires regex or split mechanism to filter meaningful info.
type TextLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextLog) Reset() {
	*x = TextLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_Logging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextLog) ProtoMessage() {}

func (x *TextLog) ProtoReflect() protoreflect.Message {
	mi := &file_logging_Logging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextLog.ProtoReflect.Descriptor instead.
func (*TextLog) Descriptor() ([]byte, []int) {
	return file_logging_Logging_proto_rawDescGZIP(), []int{2}
}

func (x *TextLog) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// JSON formatted log. The json field represents the string could be formatted as a JSON object.
type JSONLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *JSONLog) Reset() {
	*x = JSONLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_Logging_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONLog) ProtoMessage() {}

func (x *JSONLog) ProtoReflect() protoreflect.Message {
	mi := &file_logging_Logging_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONLog.ProtoReflect.Descriptor instead.
func (*JSONLog) Descriptor() ([]byte, []int) {
	return file_logging_Logging_proto_rawDescGZIP(), []int{3}
}

func (x *JSONLog) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

// YAML formatted log. The yaml field represents the string could be formatted as a YAML map.
type YAMLLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaml string `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *YAMLLog) Reset() {
	*x = YAMLLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_Logging_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YAMLLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YAMLLog) ProtoMessage() {}

func (x *YAMLLog) ProtoReflect() protoreflect.Message {
	mi := &file_logging_Logging_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YAMLLog.ProtoReflect.Descriptor instead.
func (*YAMLLog) Descriptor() ([]byte, []int) {
	return file_logging_Logging_proto_rawDescGZIP(), []int{4}
}

func (x *YAMLLog) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

// Logs with trace context, represent agent system has injects context(IDs) into log text.
type TraceContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [Optional] A string id represents the whole trace.
	TraceId string `protobuf:"bytes,1,opt,name=traceId,proto3" json:"traceId,omitempty"`
	// [Optional] A unique id represents this segment. Other segments could use this id to reference as a child segment.
	TraceSegmentId string `protobuf:"bytes,2,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	// [Optional] The number id of the span. Should be unique in the whole segment.
	// Starting at 0.
	SpanId int32 `protobuf:"varint,3,opt,name=spanId,proto3" json:"spanId,omitempty"`
}

func (x *TraceContext) Reset() {
	*x = TraceContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_Logging_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceContext) ProtoMessage() {}

func (x *TraceContext) ProtoReflect() protoreflect.Message {
	mi := &file_logging_Logging_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceContext.ProtoReflect.Descriptor instead.
func (*TraceContext) Descriptor() ([]byte, []int) {
	return file_logging_Logging_proto_rawDescGZIP(), []int{5}
}

func (x *TraceContext) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *TraceContext) GetTraceSegmentId() string {
	if x != nil {
		return x.TraceSegmentId
	}
	return ""
}

func (x *TraceContext) GetSpanId() int32 {
	if x != nil {
		return x.SpanId
	}
	return 0
}

type LogTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String key, String value pair.
	Data []*v3.KeyStringValuePair `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *LogTags) Reset() {
	*x = LogTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_Logging_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogTags) ProtoMessage() {}

func (x *LogTags) ProtoReflect() protoreflect.Message {
	mi := &file_logging_Logging_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogTags.ProtoReflect.Descriptor instead.
func (*LogTags) Descriptor() ([]byte, []int) {
	return file_logging_Logging_proto_rawDescGZIP(), []int{6}
}

func (x *LogTags) GetData() []*v3.KeyStringValuePair {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_logging_Logging_proto protoreflect.FileDescriptor

var file_logging_Logging_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x07,
	0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6f, 0x64, 0x79, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x54, 0x61, 0x67, 0x73, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x4a, 0x53, 0x4f, 0x4e, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x73,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x59, 0x41, 0x4d, 0x4c, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c,
	0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x1d, 0x0a, 0x07, 0x54,
	0x65, 0x78, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x1d, 0x0a, 0x07, 0x4a, 0x53,
	0x4f, 0x4e, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x07, 0x59, 0x41, 0x4d,
	0x4c, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x22, 0x68, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70,
	0x61, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e,
	0x49, 0x64, 0x22, 0x40, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x54, 0x61, 0x67, 0x73, 0x12, 0x35, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4b, 0x65, 0x79, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x52, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x28, 0x01, 0x42, 0x6f, 0x0a, 0x2c, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x50, 0x01, 0x5a, 0x1d, 0x73, 0x6b, 0x79, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0xaa, 0x02, 0x1d, 0x53, 0x6b, 0x79, 0x57,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_logging_Logging_proto_rawDescOnce sync.Once
	file_logging_Logging_proto_rawDescData = file_logging_Logging_proto_rawDesc
)

func file_logging_Logging_proto_rawDescGZIP() []byte {
	file_logging_Logging_proto_rawDescOnce.Do(func() {
		file_logging_Logging_proto_rawDescData = protoimpl.X.CompressGZIP(file_logging_Logging_proto_rawDescData)
	})
	return file_logging_Logging_proto_rawDescData
}

var file_logging_Logging_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_logging_Logging_proto_goTypes = []interface{}{
	(*LogData)(nil),               // 0: skywalking.v3.LogData
	(*LogDataBody)(nil),           // 1: skywalking.v3.LogDataBody
	(*TextLog)(nil),               // 2: skywalking.v3.TextLog
	(*JSONLog)(nil),               // 3: skywalking.v3.JSONLog
	(*YAMLLog)(nil),               // 4: skywalking.v3.YAMLLog
	(*TraceContext)(nil),          // 5: skywalking.v3.TraceContext
	(*LogTags)(nil),               // 6: skywalking.v3.LogTags
	(*v3.KeyStringValuePair)(nil), // 7: skywalking.v3.KeyStringValuePair
	(*v3.Commands)(nil),           // 8: skywalking.v3.Commands
}
var file_logging_Logging_proto_depIdxs = []int32{
	1, // 0: skywalking.v3.LogData.body:type_name -> skywalking.v3.LogDataBody
	5, // 1: skywalking.v3.LogData.traceContext:type_name -> skywalking.v3.TraceContext
	6, // 2: skywalking.v3.LogData.tags:type_name -> skywalking.v3.LogTags
	2, // 3: skywalking.v3.LogDataBody.text:type_name -> skywalking.v3.TextLog
	3, // 4: skywalking.v3.LogDataBody.json:type_name -> skywalking.v3.JSONLog
	4, // 5: skywalking.v3.LogDataBody.yaml:type_name -> skywalking.v3.YAMLLog
	7, // 6: skywalking.v3.LogTags.data:type_name -> skywalking.v3.KeyStringValuePair
	0, // 7: skywalking.v3.LogReportService.collect:input_type -> skywalking.v3.LogData
	8, // 8: skywalking.v3.LogReportService.collect:output_type -> skywalking.v3.Commands
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_logging_Logging_proto_init() }
func file_logging_Logging_proto_init() {
	if File_logging_Logging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logging_Logging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogData); i {
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
		file_logging_Logging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogDataBody); i {
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
		file_logging_Logging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextLog); i {
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
		file_logging_Logging_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONLog); i {
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
		file_logging_Logging_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YAMLLog); i {
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
		file_logging_Logging_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceContext); i {
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
		file_logging_Logging_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogTags); i {
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
	file_logging_Logging_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*LogDataBody_Text)(nil),
		(*LogDataBody_Json)(nil),
		(*LogDataBody_Yaml)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logging_Logging_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logging_Logging_proto_goTypes,
		DependencyIndexes: file_logging_Logging_proto_depIdxs,
		MessageInfos:      file_logging_Logging_proto_msgTypes,
	}.Build()
	File_logging_Logging_proto = out.File
	file_logging_Logging_proto_rawDesc = nil
	file_logging_Logging_proto_goTypes = nil
	file_logging_Logging_proto_depIdxs = nil
}
