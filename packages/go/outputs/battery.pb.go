// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: outputs/battery.proto

package pb_outputs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatterySensorOutput struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	CurrentOutputVoltage float32                `protobuf:"fixed32,1,opt,name=currentOutputVoltage,proto3" json:"currentOutputVoltage,omitempty"`
	WarnVoltage          float32                `protobuf:"fixed32,2,opt,name=warnVoltage,proto3" json:"warnVoltage,omitempty"`
	KillVoltage          float32                `protobuf:"fixed32,3,opt,name=killVoltage,proto3" json:"killVoltage,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *BatterySensorOutput) Reset() {
	*x = BatterySensorOutput{}
	mi := &file_outputs_battery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatterySensorOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatterySensorOutput) ProtoMessage() {}

func (x *BatterySensorOutput) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_battery_proto_msgTypes[0]
	if x != nil {
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

const file_outputs_battery_proto_rawDesc = "" +
	"\n" +
	"\x15outputs/battery.proto\x12\rprotobuf_msgs\"\x8d\x01\n" +
	"\x13BatterySensorOutput\x122\n" +
	"\x14currentOutputVoltage\x18\x01 \x01(\x02R\x14currentOutputVoltage\x12 \n" +
	"\vwarnVoltage\x18\x02 \x01(\x02R\vwarnVoltage\x12 \n" +
	"\vkillVoltage\x18\x03 \x01(\x02R\vkillVoltageB\x10Z\x0ease/pb_outputsb\x06proto3"

var (
	file_outputs_battery_proto_rawDescOnce sync.Once
	file_outputs_battery_proto_rawDescData []byte
)

func file_outputs_battery_proto_rawDescGZIP() []byte {
	file_outputs_battery_proto_rawDescOnce.Do(func() {
		file_outputs_battery_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_outputs_battery_proto_rawDesc), len(file_outputs_battery_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_outputs_battery_proto_rawDesc), len(file_outputs_battery_proto_rawDesc)),
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
	file_outputs_battery_proto_goTypes = nil
	file_outputs_battery_proto_depIdxs = nil
}
