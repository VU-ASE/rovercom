// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: outputs/energy.proto

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

type EnergySensorOutput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The measure supply voltage measured in volts (V).
	SupplyVoltage float32 `protobuf:"fixed32,1,opt,name=supplyVoltage,proto3" json:"supplyVoltage,omitempty"`
	// The current amp draw in amperes (A) of the power supply source.
	CurrentAmps float32 `protobuf:"fixed32,2,opt,name=currentAmps,proto3" json:"currentAmps,omitempty"`
	// The current power output of the Rover measured in watts (W).
	PowerWatts    float32 `protobuf:"fixed32,3,opt,name=powerWatts,proto3" json:"powerWatts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnergySensorOutput) Reset() {
	*x = EnergySensorOutput{}
	mi := &file_outputs_energy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnergySensorOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnergySensorOutput) ProtoMessage() {}

func (x *EnergySensorOutput) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_energy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnergySensorOutput.ProtoReflect.Descriptor instead.
func (*EnergySensorOutput) Descriptor() ([]byte, []int) {
	return file_outputs_energy_proto_rawDescGZIP(), []int{0}
}

func (x *EnergySensorOutput) GetSupplyVoltage() float32 {
	if x != nil {
		return x.SupplyVoltage
	}
	return 0
}

func (x *EnergySensorOutput) GetCurrentAmps() float32 {
	if x != nil {
		return x.CurrentAmps
	}
	return 0
}

func (x *EnergySensorOutput) GetPowerWatts() float32 {
	if x != nil {
		return x.PowerWatts
	}
	return 0
}

var File_outputs_energy_proto protoreflect.FileDescriptor

const file_outputs_energy_proto_rawDesc = "" +
	"\n" +
	"\x14outputs/energy.proto\x12\rprotobuf_msgs\"|\n" +
	"\x12EnergySensorOutput\x12$\n" +
	"\rsupplyVoltage\x18\x01 \x01(\x02R\rsupplyVoltage\x12 \n" +
	"\vcurrentAmps\x18\x02 \x01(\x02R\vcurrentAmps\x12\x1e\n" +
	"\n" +
	"powerWatts\x18\x03 \x01(\x02R\n" +
	"powerWattsB\x10Z\x0ease/pb_outputsb\x06proto3"

var (
	file_outputs_energy_proto_rawDescOnce sync.Once
	file_outputs_energy_proto_rawDescData []byte
)

func file_outputs_energy_proto_rawDescGZIP() []byte {
	file_outputs_energy_proto_rawDescOnce.Do(func() {
		file_outputs_energy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_outputs_energy_proto_rawDesc), len(file_outputs_energy_proto_rawDesc)))
	})
	return file_outputs_energy_proto_rawDescData
}

var file_outputs_energy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_outputs_energy_proto_goTypes = []any{
	(*EnergySensorOutput)(nil), // 0: protobuf_msgs.EnergySensorOutput
}
var file_outputs_energy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_outputs_energy_proto_init() }
func file_outputs_energy_proto_init() {
	if File_outputs_energy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_outputs_energy_proto_rawDesc), len(file_outputs_energy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outputs_energy_proto_goTypes,
		DependencyIndexes: file_outputs_energy_proto_depIdxs,
		MessageInfos:      file_outputs_energy_proto_msgTypes,
	}.Build()
	File_outputs_energy_proto = out.File
	file_outputs_energy_proto_goTypes = nil
	file_outputs_energy_proto_depIdxs = nil
}
