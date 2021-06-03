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
// source: event/Event.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v3 "skywalking/network/language/agent/v3"
	v33 "skywalking/network/language/profile/v3"
	v31 "skywalking/network/logging/v3"
	v32 "skywalking/network/management/v3"
	v34 "skywalking/network/servicemesh/v3"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EventType declares the supported transfer data type.
type EventType int32

const (
	EventType_CLRMetricType   EventType = 0
	EventType_JVMMetricType   EventType = 1
	EventType_MeterType       EventType = 2
	EventType_TracingType     EventType = 3
	EventType_Logging         EventType = 4
	EventType_ManagementType  EventType = 5
	EventType_ProfileType     EventType = 6
	EventType_ServiceMeshType EventType = 7
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "CLRMetricType",
		1: "JVMMetricType",
		2: "MeterType",
		3: "TracingType",
		4: "Logging",
		5: "ManagementType",
		6: "ProfileType",
		7: "ServiceMeshType",
	}
	EventType_value = map[string]int32{
		"CLRMetricType":   0,
		"JVMMetricType":   1,
		"MeterType":       2,
		"TracingType":     3,
		"Logging":         4,
		"ManagementType":  5,
		"ProfileType":     6,
		"ServiceMeshType": 7,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_Event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_event_Event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_event_Event_proto_rawDescGZIP(), []int{0}
}

// Event is the transfer unit in Satellite.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The occur time.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// unique event name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The data type.
	Type EventType `protobuf:"varint,3,opt,name=type,proto3,enum=skywalking.v3.EventType" json:"type,omitempty"`
	// Whether to send to remote. It is used in sampling.
	Remote bool `protobuf:"varint,4,opt,name=remote,proto3" json:"remote,omitempty"`
	// Additional meta-information.
	Meta map[string]string `protobuf:"bytes,5,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Transfer data.
	//
	// Types that are assignable to Data:
	//	*Event_Clr
	//	*Event_Jvm
	//	*Event_Meter
	//	*Event_Segment
	//	*Event_Log
	//	*Event_Instance
	//	*Event_Profile
	//	*Event_ServiceMesh
	Data isEvent_Data `protobuf_oneof:"data"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_Event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_Event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_Event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_CLRMetricType
}

func (x *Event) GetRemote() bool {
	if x != nil {
		return x.Remote
	}
	return false
}

func (x *Event) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (m *Event) GetData() isEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Event) GetClr() *v3.CLRMetricCollection {
	if x, ok := x.GetData().(*Event_Clr); ok {
		return x.Clr
	}
	return nil
}

func (x *Event) GetJvm() *v3.JVMMetricCollection {
	if x, ok := x.GetData().(*Event_Jvm); ok {
		return x.Jvm
	}
	return nil
}

func (x *Event) GetMeter() *v3.MeterData {
	if x, ok := x.GetData().(*Event_Meter); ok {
		return x.Meter
	}
	return nil
}

func (x *Event) GetSegment() *v3.SegmentObject {
	if x, ok := x.GetData().(*Event_Segment); ok {
		return x.Segment
	}
	return nil
}

func (x *Event) GetLog() *v31.LogData {
	if x, ok := x.GetData().(*Event_Log); ok {
		return x.Log
	}
	return nil
}

func (x *Event) GetInstance() *v32.InstanceProperties {
	if x, ok := x.GetData().(*Event_Instance); ok {
		return x.Instance
	}
	return nil
}

func (x *Event) GetProfile() *v33.ThreadSnapshot {
	if x, ok := x.GetData().(*Event_Profile); ok {
		return x.Profile
	}
	return nil
}

func (x *Event) GetServiceMesh() *v34.ServiceMeshMetric {
	if x, ok := x.GetData().(*Event_ServiceMesh); ok {
		return x.ServiceMesh
	}
	return nil
}

type isEvent_Data interface {
	isEvent_Data()
}

type Event_Clr struct {
	Clr *v3.CLRMetricCollection `protobuf:"bytes,6,opt,name=clr,proto3,oneof"`
}

type Event_Jvm struct {
	Jvm *v3.JVMMetricCollection `protobuf:"bytes,7,opt,name=jvm,proto3,oneof"`
}

type Event_Meter struct {
	Meter *v3.MeterData `protobuf:"bytes,8,opt,name=meter,proto3,oneof"`
}

type Event_Segment struct {
	Segment *v3.SegmentObject `protobuf:"bytes,9,opt,name=segment,proto3,oneof"`
}

type Event_Log struct {
	Log *v31.LogData `protobuf:"bytes,10,opt,name=log,proto3,oneof"`
}

type Event_Instance struct {
	Instance *v32.InstanceProperties `protobuf:"bytes,11,opt,name=instance,proto3,oneof"`
}

type Event_Profile struct {
	Profile *v33.ThreadSnapshot `protobuf:"bytes,12,opt,name=profile,proto3,oneof"`
}

type Event_ServiceMesh struct {
	ServiceMesh *v34.ServiceMeshMetric `protobuf:"bytes,13,opt,name=serviceMesh,proto3,oneof"`
}

func (*Event_Clr) isEvent_Data() {}

func (*Event_Jvm) isEvent_Data() {}

func (*Event_Meter) isEvent_Data() {}

func (*Event_Segment) isEvent_Data() {}

func (*Event_Log) isEvent_Data() {}

func (*Event_Instance) isEvent_Data() {}

func (*Event_Profile) isEvent_Data() {}

func (*Event_ServiceMesh) isEvent_Data() {}

var File_event_Event_proto protoreflect.FileDescriptor

var file_event_Event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x1a, 0x1e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2f, 0x43, 0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2f, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2f, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x54,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe,
	0x05, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x03, 0x63, 0x6c, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x63, 0x6c, 0x72, 0x12, 0x36, 0x0a, 0x03,
	0x6a, 0x76, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6b, 0x79, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x03, 0x6a, 0x76, 0x6d, 0x12, 0x30, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x05, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x3f, 0x0a, 0x08,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x1a, 0x37,
	0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a,
	0x98, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a,
	0x0d, 0x43, 0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x10, 0x04,
	0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x65, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x10, 0x07, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x61,
	0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_Event_proto_rawDescOnce sync.Once
	file_event_Event_proto_rawDescData = file_event_Event_proto_rawDesc
)

func file_event_Event_proto_rawDescGZIP() []byte {
	file_event_Event_proto_rawDescOnce.Do(func() {
		file_event_Event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_Event_proto_rawDescData)
	})
	return file_event_Event_proto_rawDescData
}

var file_event_Event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_Event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_event_Event_proto_goTypes = []interface{}{
	(EventType)(0),                 // 0: skywalking.v3.EventType
	(*Event)(nil),                  // 1: skywalking.v3.Event
	nil,                            // 2: skywalking.v3.Event.MetaEntry
	(*v3.CLRMetricCollection)(nil), // 3: skywalking.v3.CLRMetricCollection
	(*v3.JVMMetricCollection)(nil), // 4: skywalking.v3.JVMMetricCollection
	(*v3.MeterData)(nil),           // 5: skywalking.v3.MeterData
	(*v3.SegmentObject)(nil),       // 6: skywalking.v3.SegmentObject
	(*v31.LogData)(nil),            // 7: skywalking.v3.LogData
	(*v32.InstanceProperties)(nil), // 8: skywalking.v3.InstanceProperties
	(*v33.ThreadSnapshot)(nil),     // 9: skywalking.v3.ThreadSnapshot
	(*v34.ServiceMeshMetric)(nil),  // 10: skywalking.v3.ServiceMeshMetric
}
var file_event_Event_proto_depIdxs = []int32{
	0,  // 0: skywalking.v3.Event.type:type_name -> skywalking.v3.EventType
	2,  // 1: skywalking.v3.Event.meta:type_name -> skywalking.v3.Event.MetaEntry
	3,  // 2: skywalking.v3.Event.clr:type_name -> skywalking.v3.CLRMetricCollection
	4,  // 3: skywalking.v3.Event.jvm:type_name -> skywalking.v3.JVMMetricCollection
	5,  // 4: skywalking.v3.Event.meter:type_name -> skywalking.v3.MeterData
	6,  // 5: skywalking.v3.Event.segment:type_name -> skywalking.v3.SegmentObject
	7,  // 6: skywalking.v3.Event.log:type_name -> skywalking.v3.LogData
	8,  // 7: skywalking.v3.Event.instance:type_name -> skywalking.v3.InstanceProperties
	9,  // 8: skywalking.v3.Event.profile:type_name -> skywalking.v3.ThreadSnapshot
	10, // 9: skywalking.v3.Event.serviceMesh:type_name -> skywalking.v3.ServiceMeshMetric
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_event_Event_proto_init() }
func file_event_Event_proto_init() {
	if File_event_Event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_Event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
	file_event_Event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_Clr)(nil),
		(*Event_Jvm)(nil),
		(*Event_Meter)(nil),
		(*Event_Segment)(nil),
		(*Event_Log)(nil),
		(*Event_Instance)(nil),
		(*Event_Profile)(nil),
		(*Event_ServiceMesh)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_Event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_Event_proto_goTypes,
		DependencyIndexes: file_event_Event_proto_depIdxs,
		EnumInfos:         file_event_Event_proto_enumTypes,
		MessageInfos:      file_event_Event_proto_msgTypes,
	}.Build()
	File_event_Event_proto = out.File
	file_event_Event_proto_rawDesc = nil
	file_event_Event_proto_goTypes = nil
	file_event_Event_proto_depIdxs = nil
}
