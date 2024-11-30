// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: outputs/battery.proto

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

type BatterySensorOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentOutputVoltage float32 `protobuf:"fixed32,1,opt,name=currentOutputVoltage,proto3" json:"currentOutputVoltage,omitempty"`
	WarnVoltage          float32 `protobuf:"fixed32,2,opt,name=warnVoltage,proto3" json:"warnVoltage,omitempty"`
	KillVoltage          float32 `protobuf:"fixed32,3,opt,name=killVoltage,proto3" json:"killVoltage,omitempty"`
}

func (x *BatterySensorOutput) Reset() {
	*x = BatterySensorOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outputs_battery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatterySensorOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatterySensorOutput) ProtoMessage() {}

func (x *BatterySensorOutput) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_battery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatterySensorOutput.ProtoReflect.Descriptor instead.
func (*BatterySensorOutput) Descriptor() ([]byte, []int) {
	return file_outputs_battery_proto_rawDescGZIP(), []int{0}
}

func (x *BatterySensorOutput) GetCurrentOutputVoltage() float32 {
	if x != nil {
		return x.CurrentOutputVoltage
	}
	return 0
}

func (x *BatterySensorOutput) GetWarnVoltage() float32 {
	if x != nil {
		return x.WarnVoltage
	}
	return 0
}

func (x *BatterySensorOutput) GetKillVoltage() float32 {
	if x != nil {
		return x.KillVoltage
	}
	return 0
}

var File_outputs_battery_proto protoreflect.FileDescriptor

var file_outputs_battery_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x32,
	0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x56,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x6e, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x6e, 0x56, 0x6f, 0x6c,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x56, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x56,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x62,
	0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outputs_battery_proto_rawDescOnce sync.Once
	file_outputs_battery_proto_rawDescData = file_outputs_battery_proto_rawDesc
)

func file_outputs_battery_proto_rawDescGZIP() []byte {
	file_outputs_battery_proto_rawDescOnce.Do(func() {
		file_outputs_battery_proto_rawDescData = protoimpl.X.CompressGZIP(file_outputs_battery_proto_rawDescData)
	})
	return file_outputs_battery_proto_rawDescData
}

var file_outputs_battery_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_outputs_battery_proto_goTypes = []any{
	(*BatterySensorOutput)(nil), // 0: protobuf_msgs.BatterySensorOutput
}
var file_outputs_battery_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_outputs_battery_proto_init() }
func file_outputs_battery_proto_init() {
	if File_outputs_battery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outputs_battery_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BatterySensorOutput); i {
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
			RawDescriptor: file_outputs_battery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outputs_battery_proto_goTypes,
		DependencyIndexes: file_outputs_battery_proto_depIdxs,
		MessageInfos:      file_outputs_battery_proto_msgTypes,
	}.Build()
	File_outputs_battery_proto = out.File
	file_outputs_battery_proto_rawDesc = nil
	file_outputs_battery_proto_goTypes = nil
	file_outputs_battery_proto_depIdxs = nil
}
