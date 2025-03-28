// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: outputs/wrapper.proto

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

type SensorOutput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Every sensor has a unique ID to support multiple sensors of the same type
	SensorId uint32 `protobuf:"varint,1,opt,name=sensorId,proto3" json:"sensorId,omitempty"`
	// Add a timestamp to the output to make debugging, logging and synchronisation easier
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Report an error if the sensor is not working correctly (controller can decide to ignore or stop the car)
	// 0 = no error, any other value = error
	Status uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// Add the output here to make it available to the receiver
	//
	// Types that are valid to be assigned to SensorOutput:
	//
	//	*SensorOutput_CameraOutput
	//	*SensorOutput_DistanceOutput
	//	*SensorOutput_SpeedOutput
	//	*SensorOutput_ControllerOutput
	//	*SensorOutput_ImuOutput
	//	*SensorOutput_BatteryOutput
	//	*SensorOutput_RpmOuput
	//	*SensorOutput_LuxOutput
	//	*SensorOutput_LaptimeOutput
	SensorOutput  isSensorOutput_SensorOutput `protobuf_oneof:"sensorOutput"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SensorOutput) Reset() {
	*x = SensorOutput{}
	mi := &file_outputs_wrapper_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SensorOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorOutput) ProtoMessage() {}

func (x *SensorOutput) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_wrapper_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorOutput.ProtoReflect.Descriptor instead.
func (*SensorOutput) Descriptor() ([]byte, []int) {
	return file_outputs_wrapper_proto_rawDescGZIP(), []int{0}
}

func (x *SensorOutput) GetSensorId() uint32 {
	if x != nil {
		return x.SensorId
	}
	return 0
}

func (x *SensorOutput) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *SensorOutput) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SensorOutput) GetSensorOutput() isSensorOutput_SensorOutput {
	if x != nil {
		return x.SensorOutput
	}
	return nil
}

func (x *SensorOutput) GetCameraOutput() *CameraSensorOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_CameraOutput); ok {
			return x.CameraOutput
		}
	}
	return nil
}

func (x *SensorOutput) GetDistanceOutput() *DistanceSensorOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_DistanceOutput); ok {
			return x.DistanceOutput
		}
	}
	return nil
}

func (x *SensorOutput) GetSpeedOutput() *SpeedSensorOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_SpeedOutput); ok {
			return x.SpeedOutput
		}
	}
	return nil
}

func (x *SensorOutput) GetControllerOutput() *ControllerOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_ControllerOutput); ok {
			return x.ControllerOutput
		}
	}
	return nil
}

func (x *SensorOutput) GetImuOutput() *ImuSensorOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_ImuOutput); ok {
			return x.ImuOutput
		}
	}
	return nil
}

func (x *SensorOutput) GetBatteryOutput() *BatterySensorOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_BatteryOutput); ok {
			return x.BatteryOutput
		}
	}
	return nil
}

func (x *SensorOutput) GetRpmOuput() *RpmSensorOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_RpmOuput); ok {
			return x.RpmOuput
		}
	}
	return nil
}

func (x *SensorOutput) GetLuxOutput() *LuxSensorOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_LuxOutput); ok {
			return x.LuxOutput
		}
	}
	return nil
}

func (x *SensorOutput) GetLaptimeOutput() *LapTimeOutput {
	if x != nil {
		if x, ok := x.SensorOutput.(*SensorOutput_LaptimeOutput); ok {
			return x.LaptimeOutput
		}
	}
	return nil
}

type isSensorOutput_SensorOutput interface {
	isSensorOutput_SensorOutput()
}

type SensorOutput_CameraOutput struct {
	CameraOutput *CameraSensorOutput `protobuf:"bytes,4,opt,name=cameraOutput,proto3,oneof"`
}

type SensorOutput_DistanceOutput struct {
	DistanceOutput *DistanceSensorOutput `protobuf:"bytes,5,opt,name=distanceOutput,proto3,oneof"`
}

type SensorOutput_SpeedOutput struct {
	SpeedOutput *SpeedSensorOutput `protobuf:"bytes,6,opt,name=speedOutput,proto3,oneof"`
}

type SensorOutput_ControllerOutput struct {
	ControllerOutput *ControllerOutput `protobuf:"bytes,7,opt,name=controllerOutput,proto3,oneof"`
}

type SensorOutput_ImuOutput struct {
	ImuOutput *ImuSensorOutput `protobuf:"bytes,8,opt,name=imuOutput,proto3,oneof"`
}

type SensorOutput_BatteryOutput struct {
	BatteryOutput *BatterySensorOutput `protobuf:"bytes,9,opt,name=batteryOutput,proto3,oneof"`
}

type SensorOutput_RpmOuput struct {
	RpmOuput *RpmSensorOutput `protobuf:"bytes,10,opt,name=rpmOuput,proto3,oneof"`
}

type SensorOutput_LuxOutput struct {
	LuxOutput *LuxSensorOutput `protobuf:"bytes,11,opt,name=luxOutput,proto3,oneof"`
}

type SensorOutput_LaptimeOutput struct {
	LaptimeOutput *LapTimeOutput `protobuf:"bytes,12,opt,name=laptimeOutput,proto3,oneof"`
}

func (*SensorOutput_CameraOutput) isSensorOutput_SensorOutput() {}

func (*SensorOutput_DistanceOutput) isSensorOutput_SensorOutput() {}

func (*SensorOutput_SpeedOutput) isSensorOutput_SensorOutput() {}

func (*SensorOutput_ControllerOutput) isSensorOutput_SensorOutput() {}

func (*SensorOutput_ImuOutput) isSensorOutput_SensorOutput() {}

func (*SensorOutput_BatteryOutput) isSensorOutput_SensorOutput() {}

func (*SensorOutput_RpmOuput) isSensorOutput_SensorOutput() {}

func (*SensorOutput_LuxOutput) isSensorOutput_SensorOutput() {}

func (*SensorOutput_LaptimeOutput) isSensorOutput_SensorOutput() {}

var File_outputs_wrapper_proto protoreflect.FileDescriptor

const file_outputs_wrapper_proto_rawDesc = "" +
	"\n" +
	"\x15outputs/wrapper.proto\x12\rprotobuf_msgs\x1a\x14outputs/camera.proto\x1a\x16outputs/distance.proto\x1a\x13outputs/speed.proto\x1a\x18outputs/controller.proto\x1a\x11outputs/imu.proto\x1a\x15outputs/battery.proto\x1a\x11outputs/rpm.proto\x1a\x11outputs/lux.proto\x1a\x15outputs/laptime.proto\"\xed\x05\n" +
	"\fSensorOutput\x12\x1a\n" +
	"\bsensorId\x18\x01 \x01(\rR\bsensorId\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\x04R\ttimestamp\x12\x16\n" +
	"\x06status\x18\x03 \x01(\rR\x06status\x12G\n" +
	"\fcameraOutput\x18\x04 \x01(\v2!.protobuf_msgs.CameraSensorOutputH\x00R\fcameraOutput\x12M\n" +
	"\x0edistanceOutput\x18\x05 \x01(\v2#.protobuf_msgs.DistanceSensorOutputH\x00R\x0edistanceOutput\x12D\n" +
	"\vspeedOutput\x18\x06 \x01(\v2 .protobuf_msgs.SpeedSensorOutputH\x00R\vspeedOutput\x12M\n" +
	"\x10controllerOutput\x18\a \x01(\v2\x1f.protobuf_msgs.ControllerOutputH\x00R\x10controllerOutput\x12>\n" +
	"\timuOutput\x18\b \x01(\v2\x1e.protobuf_msgs.ImuSensorOutputH\x00R\timuOutput\x12J\n" +
	"\rbatteryOutput\x18\t \x01(\v2\".protobuf_msgs.BatterySensorOutputH\x00R\rbatteryOutput\x12<\n" +
	"\brpmOuput\x18\n" +
	" \x01(\v2\x1e.protobuf_msgs.RpmSensorOutputH\x00R\brpmOuput\x12>\n" +
	"\tluxOutput\x18\v \x01(\v2\x1e.protobuf_msgs.LuxSensorOutputH\x00R\tluxOutput\x12D\n" +
	"\rlaptimeOutput\x18\f \x01(\v2\x1c.protobuf_msgs.LapTimeOutputH\x00R\rlaptimeOutputB\x0e\n" +
	"\fsensorOutputB\x10Z\x0ease/pb_outputsb\x06proto3"

var (
	file_outputs_wrapper_proto_rawDescOnce sync.Once
	file_outputs_wrapper_proto_rawDescData []byte
)

func file_outputs_wrapper_proto_rawDescGZIP() []byte {
	file_outputs_wrapper_proto_rawDescOnce.Do(func() {
		file_outputs_wrapper_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_outputs_wrapper_proto_rawDesc), len(file_outputs_wrapper_proto_rawDesc)))
	})
	return file_outputs_wrapper_proto_rawDescData
}

var file_outputs_wrapper_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_outputs_wrapper_proto_goTypes = []any{
	(*SensorOutput)(nil),         // 0: protobuf_msgs.SensorOutput
	(*CameraSensorOutput)(nil),   // 1: protobuf_msgs.CameraSensorOutput
	(*DistanceSensorOutput)(nil), // 2: protobuf_msgs.DistanceSensorOutput
	(*SpeedSensorOutput)(nil),    // 3: protobuf_msgs.SpeedSensorOutput
	(*ControllerOutput)(nil),     // 4: protobuf_msgs.ControllerOutput
	(*ImuSensorOutput)(nil),      // 5: protobuf_msgs.ImuSensorOutput
	(*BatterySensorOutput)(nil),  // 6: protobuf_msgs.BatterySensorOutput
	(*RpmSensorOutput)(nil),      // 7: protobuf_msgs.RpmSensorOutput
	(*LuxSensorOutput)(nil),      // 8: protobuf_msgs.LuxSensorOutput
	(*LapTimeOutput)(nil),        // 9: protobuf_msgs.LapTimeOutput
}
var file_outputs_wrapper_proto_depIdxs = []int32{
	1, // 0: protobuf_msgs.SensorOutput.cameraOutput:type_name -> protobuf_msgs.CameraSensorOutput
	2, // 1: protobuf_msgs.SensorOutput.distanceOutput:type_name -> protobuf_msgs.DistanceSensorOutput
	3, // 2: protobuf_msgs.SensorOutput.speedOutput:type_name -> protobuf_msgs.SpeedSensorOutput
	4, // 3: protobuf_msgs.SensorOutput.controllerOutput:type_name -> protobuf_msgs.ControllerOutput
	5, // 4: protobuf_msgs.SensorOutput.imuOutput:type_name -> protobuf_msgs.ImuSensorOutput
	6, // 5: protobuf_msgs.SensorOutput.batteryOutput:type_name -> protobuf_msgs.BatterySensorOutput
	7, // 6: protobuf_msgs.SensorOutput.rpmOuput:type_name -> protobuf_msgs.RpmSensorOutput
	8, // 7: protobuf_msgs.SensorOutput.luxOutput:type_name -> protobuf_msgs.LuxSensorOutput
	9, // 8: protobuf_msgs.SensorOutput.laptimeOutput:type_name -> protobuf_msgs.LapTimeOutput
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_outputs_wrapper_proto_init() }
func file_outputs_wrapper_proto_init() {
	if File_outputs_wrapper_proto != nil {
		return
	}
	file_outputs_camera_proto_init()
	file_outputs_distance_proto_init()
	file_outputs_speed_proto_init()
	file_outputs_controller_proto_init()
	file_outputs_imu_proto_init()
	file_outputs_battery_proto_init()
	file_outputs_rpm_proto_init()
	file_outputs_lux_proto_init()
	file_outputs_laptime_proto_init()
	file_outputs_wrapper_proto_msgTypes[0].OneofWrappers = []any{
		(*SensorOutput_CameraOutput)(nil),
		(*SensorOutput_DistanceOutput)(nil),
		(*SensorOutput_SpeedOutput)(nil),
		(*SensorOutput_ControllerOutput)(nil),
		(*SensorOutput_ImuOutput)(nil),
		(*SensorOutput_BatteryOutput)(nil),
		(*SensorOutput_RpmOuput)(nil),
		(*SensorOutput_LuxOutput)(nil),
		(*SensorOutput_LaptimeOutput)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_outputs_wrapper_proto_rawDesc), len(file_outputs_wrapper_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outputs_wrapper_proto_goTypes,
		DependencyIndexes: file_outputs_wrapper_proto_depIdxs,
		MessageInfos:      file_outputs_wrapper_proto_msgTypes,
	}.Build()
	File_outputs_wrapper_proto = out.File
	file_outputs_wrapper_proto_goTypes = nil
	file_outputs_wrapper_proto_depIdxs = nil
}
