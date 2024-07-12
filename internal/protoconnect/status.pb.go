// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: status.proto

package protoconnect

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusEvent_Type int32

const (
	// The status has not been specified.
	StatusEvent_TYPE_UNSPECIFIED StatusEvent_Type = 0
	// An instance has parsed a config and is now attempting to run a pipeline.
	StatusEvent_TYPE_INITIALISING StatusEvent_Type = 1
	// An instance is running and is connected to all inputs and outputs.
	StatusEvent_TYPE_CONNECTION_HEALTHY StatusEvent_Type = 2
	// An instance is running but is not connected to all inputs and outputs.
	StatusEvent_TYPE_CONNECTION_ERROR StatusEvent_Type = 3
	// An instance is in the process of exiting and will no longer sent status events.
	StatusEvent_TYPE_EXITING StatusEvent_Type = 4
)

// Enum value maps for StatusEvent_Type.
var (
	StatusEvent_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_INITIALISING",
		2: "TYPE_CONNECTION_HEALTHY",
		3: "TYPE_CONNECTION_ERROR",
		4: "TYPE_EXITING",
	}
	StatusEvent_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":        0,
		"TYPE_INITIALISING":       1,
		"TYPE_CONNECTION_HEALTHY": 2,
		"TYPE_CONNECTION_ERROR":   3,
		"TYPE_EXITING":            4,
	}
)

func (x StatusEvent_Type) Enum() *StatusEvent_Type {
	p := new(StatusEvent_Type)
	*p = x
	return p
}

func (x StatusEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_status_proto_enumTypes[0].Descriptor()
}

func (StatusEvent_Type) Type() protoreflect.EnumType {
	return &file_status_proto_enumTypes[0]
}

func (x StatusEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusEvent_Type.Descriptor instead.
func (StatusEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{2, 0}
}

// ConnectionError describes a specific connection failure.
type ConnectionError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`   // The error message.
	Path    []string `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`         // The path of the connector in the config.
	Label   *string  `protobuf:"bytes,3,opt,name=label,proto3,oneof" json:"label,omitempty"` // An optional label given to the connector.
}

func (x *ConnectionError) Reset() {
	*x = ConnectionError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionError) ProtoMessage() {}

func (x *ConnectionError) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionError.ProtoReflect.Descriptor instead.
func (*ConnectionError) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConnectionError) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *ConnectionError) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

// ExitError describes an error encountered that caused the instance to exit.
type ExitError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // The error message.
}

func (x *ExitError) Reset() {
	*x = ExitError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitError) ProtoMessage() {}

func (x *ExitError) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitError.ProtoReflect.Descriptor instead.
func (*ExitError) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{1}
}

func (x *ExitError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// StatusEvent describes the current state of an individual connect instance,
// which is self-reported periodically.
type StatusEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             StatusEvent_Type   `protobuf:"varint,1,opt,name=type,proto3,enum=redpanda.api.connect.v1alpha1.StatusEvent_Type" json:"type,omitempty"` // The type of the event.
	PipelineId       string             `protobuf:"bytes,2,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`                        // The identifier of the running pipeline.
	InstanceId       string             `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`                        // The unique identifier of the connect instance.
	Timestamp        int64              `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                           // The time this event was emitted.
	ConnectionErrors []*ConnectionError `protobuf:"bytes,5,rep,name=connection_errors,json=connectionErrors,proto3" json:"connection_errors,omitempty"`      // Zero or more connection errors.
	ExitError        *ExitError         `protobuf:"bytes,6,opt,name=exit_error,json=exitError,proto3,oneof" json:"exit_error,omitempty"`                     // An optional exit error.
}

func (x *StatusEvent) Reset() {
	*x = StatusEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusEvent) ProtoMessage() {}

func (x *StatusEvent) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusEvent.ProtoReflect.Descriptor instead.
func (*StatusEvent) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{2}
}

func (x *StatusEvent) GetType() StatusEvent_Type {
	if x != nil {
		return x.Type
	}
	return StatusEvent_TYPE_UNSPECIFIED
}

func (x *StatusEvent) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

func (x *StatusEvent) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *StatusEvent) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *StatusEvent) GetConnectionErrors() []*ConnectionError {
	if x != nil {
		return x.ConnectionErrors
	}
	return nil
}

func (x *StatusEvent) GetExitError() *ExitError {
	if x != nil {
		return x.ExitError
	}
	return nil
}

var File_status_proto protoreflect.FileDescriptor

var file_status_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x64, 0x0a,
	0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x19,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0x25, 0x0a, 0x09, 0x45, 0x78, 0x69, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xeb, 0x03, 0x0a, 0x0b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x5b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0a,
	0x65, 0x78, 0x69, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x78, 0x69, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78,
	0x69, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x22, 0x7d, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x1b, 0x0a, 0x17, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x58, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78,
	0x69, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x17, 0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_status_proto_rawDescOnce sync.Once
	file_status_proto_rawDescData = file_status_proto_rawDesc
)

func file_status_proto_rawDescGZIP() []byte {
	file_status_proto_rawDescOnce.Do(func() {
		file_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_status_proto_rawDescData)
	})
	return file_status_proto_rawDescData
}

var file_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_status_proto_goTypes = []any{
	(StatusEvent_Type)(0),   // 0: redpanda.api.connect.v1alpha1.StatusEvent.Type
	(*ConnectionError)(nil), // 1: redpanda.api.connect.v1alpha1.ConnectionError
	(*ExitError)(nil),       // 2: redpanda.api.connect.v1alpha1.ExitError
	(*StatusEvent)(nil),     // 3: redpanda.api.connect.v1alpha1.StatusEvent
}
var file_status_proto_depIdxs = []int32{
	0, // 0: redpanda.api.connect.v1alpha1.StatusEvent.type:type_name -> redpanda.api.connect.v1alpha1.StatusEvent.Type
	1, // 1: redpanda.api.connect.v1alpha1.StatusEvent.connection_errors:type_name -> redpanda.api.connect.v1alpha1.ConnectionError
	2, // 2: redpanda.api.connect.v1alpha1.StatusEvent.exit_error:type_name -> redpanda.api.connect.v1alpha1.ExitError
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_status_proto_init() }
func file_status_proto_init() {
	if File_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_status_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectionError); i {
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
		file_status_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExitError); i {
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
		file_status_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StatusEvent); i {
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
	file_status_proto_msgTypes[0].OneofWrappers = []any{}
	file_status_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_status_proto_goTypes,
		DependencyIndexes: file_status_proto_depIdxs,
		EnumInfos:         file_status_proto_enumTypes,
		MessageInfos:      file_status_proto_msgTypes,
	}.Build()
	File_status_proto = out.File
	file_status_proto_rawDesc = nil
	file_status_proto_goTypes = nil
	file_status_proto_depIdxs = nil
}
