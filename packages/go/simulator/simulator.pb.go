// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: simulator/simulator.proto

package pb_simulator_messages

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

// Possible Sim Requests. Useful for interfaces with Gym
type SimStatus int32

const (
	SimStatus_SIM_PAUSED    SimStatus = 0 // Simulator is paused
	SimStatus_SIM_REQ_STEP  SimStatus = 1 // Request single step
	SimStatus_SIM_REQ_RESET SimStatus = 2 // Request hard reset
)

// Enum value maps for SimStatus.
var (
	SimStatus_name = map[int32]string{
		0: "SIM_PAUSED",
		1: "SIM_REQ_STEP",
		2: "SIM_REQ_RESET",
	}
	SimStatus_value = map[string]int32{
		"SIM_PAUSED":    0,
		"SIM_REQ_STEP":  1,
		"SIM_REQ_RESET": 2,
	}
)

func (x SimStatus) Enum() *SimStatus {
	p := new(SimStatus)
	*p = x
	return p
}

func (x SimStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SimStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_simulator_simulator_proto_enumTypes[0].Descriptor()
}

func (SimStatus) Type() protoreflect.EnumType {
	return &file_simulator_simulator_proto_enumTypes[0]
}

func (x SimStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SimStatus.Descriptor instead.
func (SimStatus) EnumDescriptor() ([]byte, []int) {
	return file_simulator_simulator_proto_rawDescGZIP(), []int{0}
}

// Simulator sensor outputs.
type SimulatorImageOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Width         uint32                 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height        uint32                 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Pixels        []byte                 `protobuf:"bytes,4,opt,name=pixels,proto3" json:"pixels,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SimulatorImageOutput) Reset() {
	*x = SimulatorImageOutput{}
	mi := &file_simulator_simulator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimulatorImageOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorImageOutput) ProtoMessage() {}

func (x *SimulatorImageOutput) ProtoReflect() protoreflect.Message {
	mi := &file_simulator_simulator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorImageOutput.ProtoReflect.Descriptor instead.
func (*SimulatorImageOutput) Descriptor() ([]byte, []int) {
	return file_simulator_simulator_proto_rawDescGZIP(), []int{0}
}

func (x *SimulatorImageOutput) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *SimulatorImageOutput) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *SimulatorImageOutput) GetPixels() []byte {
	if x != nil {
		return x.Pixels
	}
	return nil
}

// Generic state of Simulator
type SimulatorState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Speed         float32                `protobuf:"fixed32,1,opt,name=speed,proto3" json:"speed,omitempty"`                                              // Meters per second
	WheelOffTrack []bool                 `protobuf:"varint,2,rep,packed,name=wheel_off_track,json=wheelOffTrack,proto3" json:"wheel_off_track,omitempty"` // [0] = FL, [1] = FR, [2] = RL, [3] = RR
	Image         *SimulatorImageOutput  `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`                                                // Image frame
	Pos           []float32              `protobuf:"fixed32,4,rep,packed,name=pos,proto3" json:"pos,omitempty"`                                           // [0] = x, [1] = y, [2] = z
	IsDrifting    bool                   `protobuf:"varint,5,opt,name=isDrifting,proto3" json:"isDrifting,omitempty"`                                     // =false when not drifting
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SimulatorState) Reset() {
	*x = SimulatorState{}
	mi := &file_simulator_simulator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimulatorState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorState) ProtoMessage() {}

func (x *SimulatorState) ProtoReflect() protoreflect.Message {
	mi := &file_simulator_simulator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorState.ProtoReflect.Descriptor instead.
func (*SimulatorState) Descriptor() ([]byte, []int) {
	return file_simulator_simulator_proto_rawDescGZIP(), []int{1}
}

func (x *SimulatorState) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *SimulatorState) GetWheelOffTrack() []bool {
	if x != nil {
		return x.WheelOffTrack
	}
	return nil
}

func (x *SimulatorState) GetImage() *SimulatorImageOutput {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *SimulatorState) GetPos() []float32 {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *SimulatorState) GetIsDrifting() bool {
	if x != nil {
		return x.IsDrifting
	}
	return false
}

var File_simulator_simulator_proto protoreflect.FileDescriptor

var file_simulator_simulator_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x22, 0x5c, 0x0a, 0x14, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x68, 0x65, 0x65, 0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x08, 0x52, 0x0d, 0x77, 0x68, 0x65, 0x65,
	0x6c, 0x4f, 0x66, 0x66, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x44, 0x72, 0x69, 0x66,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x72,
	0x69, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x2a, 0x40, 0x0a, 0x09, 0x53, 0x69, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x49, 0x4d, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x49, 0x4d, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x53,
	0x54, 0x45, 0x50, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x49, 0x4d, 0x5f, 0x52, 0x45, 0x51,
	0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x02, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x73, 0x65, 0x2f,
	0x70, 0x62, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simulator_simulator_proto_rawDescOnce sync.Once
	file_simulator_simulator_proto_rawDescData = file_simulator_simulator_proto_rawDesc
)

func file_simulator_simulator_proto_rawDescGZIP() []byte {
	file_simulator_simulator_proto_rawDescOnce.Do(func() {
		file_simulator_simulator_proto_rawDescData = protoimpl.X.CompressGZIP(file_simulator_simulator_proto_rawDescData)
	})
	return file_simulator_simulator_proto_rawDescData
}

var file_simulator_simulator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_simulator_simulator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_simulator_simulator_proto_goTypes = []any{
	(SimStatus)(0),               // 0: protobuf_msgs.SimStatus
	(*SimulatorImageOutput)(nil), // 1: protobuf_msgs.SimulatorImageOutput
	(*SimulatorState)(nil),       // 2: protobuf_msgs.SimulatorState
}
var file_simulator_simulator_proto_depIdxs = []int32{
	1, // 0: protobuf_msgs.SimulatorState.image:type_name -> protobuf_msgs.SimulatorImageOutput
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_simulator_simulator_proto_init() }
func file_simulator_simulator_proto_init() {
	if File_simulator_simulator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_simulator_simulator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_simulator_simulator_proto_goTypes,
		DependencyIndexes: file_simulator_simulator_proto_depIdxs,
		EnumInfos:         file_simulator_simulator_proto_enumTypes,
		MessageInfos:      file_simulator_simulator_proto_msgTypes,
	}.Build()
	File_simulator_simulator_proto = out.File
	file_simulator_simulator_proto_rawDesc = nil
	file_simulator_simulator_proto_goTypes = nil
	file_simulator_simulator_proto_depIdxs = nil
}
