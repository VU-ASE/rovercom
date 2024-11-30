// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: outputs/controller.proto

package pb_module_outputs

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

type ControllerOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Steering angle (-1.0 to 1.0 <-> left - right)
	SteeringAngle float32 `protobuf:"fixed32,2,opt,name=steeringAngle,proto3" json:"steeringAngle,omitempty"`
	// Throttle (-1.0 to 1.0 <-> full reverse - full forward)
	LeftThrottle  float32 `protobuf:"fixed32,3,opt,name=leftThrottle,proto3" json:"leftThrottle,omitempty"`
	RightThrottle float32 `protobuf:"fixed32,4,opt,name=rightThrottle,proto3" json:"rightThrottle,omitempty"`
	// Onboard lights (0.0 to 1.0 <-> off - on)
	FrontLights bool `protobuf:"varint,5,opt,name=frontLights,proto3" json:"frontLights,omitempty"`
	// Fan speed (0.0 to 1.0 <-> off - full speed)
	FanSpeed float32 `protobuf:"fixed32,6,opt,name=fanSpeed,proto3" json:"fanSpeed,omitempty"`
	// Useful for debugging
	RawError    float32 `protobuf:"fixed32,7,opt,name=rawError,proto3" json:"rawError,omitempty"`       // the error value before quadratic scaling
	ScaledError float32 `protobuf:"fixed32,8,opt,name=scaledError,proto3" json:"scaledError,omitempty"` // the error value after quadratic scaling, but before PID
}

func (x *ControllerOutput) Reset() {
	*x = ControllerOutput{}
	mi := &file_outputs_controller_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControllerOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerOutput) ProtoMessage() {}

func (x *ControllerOutput) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_controller_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerOutput.ProtoReflect.Descriptor instead.
func (*ControllerOutput) Descriptor() ([]byte, []int) {
	return file_outputs_controller_proto_rawDescGZIP(), []int{0}
}

func (x *ControllerOutput) GetSteeringAngle() float32 {
	if x != nil {
		return x.SteeringAngle
	}
	return 0
}

func (x *ControllerOutput) GetLeftThrottle() float32 {
	if x != nil {
		return x.LeftThrottle
	}
	return 0
}

func (x *ControllerOutput) GetRightThrottle() float32 {
	if x != nil {
		return x.RightThrottle
	}
	return 0
}

func (x *ControllerOutput) GetFrontLights() bool {
	if x != nil {
		return x.FrontLights
	}
	return false
}

func (x *ControllerOutput) GetFanSpeed() float32 {
	if x != nil {
		return x.FanSpeed
	}
	return 0
}

func (x *ControllerOutput) GetRawError() float32 {
	if x != nil {
		return x.RawError
	}
	return 0
}

func (x *ControllerOutput) GetScaledError() float32 {
	if x != nil {
		return x.ScaledError
	}
	return 0
}

var File_outputs_controller_proto protoreflect.FileDescriptor

var file_outputs_controller_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x22, 0xfe, 0x01, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x73, 0x74, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x74, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x41,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x66, 0x74, 0x54, 0x68, 0x72, 0x6f,
	0x74, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x6c, 0x65, 0x66, 0x74,
	0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x72, 0x69, 0x67, 0x68, 0x74, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x66, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x61, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x72, 0x61, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x17, 0x5a, 0x15, 0x61, 0x73,
	0x65, 0x2f, 0x70, 0x62, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outputs_controller_proto_rawDescOnce sync.Once
	file_outputs_controller_proto_rawDescData = file_outputs_controller_proto_rawDesc
)

func file_outputs_controller_proto_rawDescGZIP() []byte {
	file_outputs_controller_proto_rawDescOnce.Do(func() {
		file_outputs_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_outputs_controller_proto_rawDescData)
	})
	return file_outputs_controller_proto_rawDescData
}

var file_outputs_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_outputs_controller_proto_goTypes = []any{
	(*ControllerOutput)(nil), // 0: protobuf_msgs.ControllerOutput
}
var file_outputs_controller_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_outputs_controller_proto_init() }
func file_outputs_controller_proto_init() {
	if File_outputs_controller_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_outputs_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outputs_controller_proto_goTypes,
		DependencyIndexes: file_outputs_controller_proto_depIdxs,
		MessageInfos:      file_outputs_controller_proto_msgTypes,
	}.Build()
	File_outputs_controller_proto = out.File
	file_outputs_controller_proto_rawDesc = nil
	file_outputs_controller_proto_goTypes = nil
	file_outputs_controller_proto_depIdxs = nil
}
