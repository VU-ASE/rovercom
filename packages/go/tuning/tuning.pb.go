// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: tuning/tuning.proto

package pb_tuning

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

type TuningState struct {
	state             protoimpl.MessageState   `protogen:"open.v1"`
	Timestamp         uint64                   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // increased when the state changes, used to prevent unnecessary updates (so *not* in milliseconds since epoch)
	DynamicParameters []*TuningState_Parameter `protobuf:"bytes,2,rep,name=dynamicParameters,proto3" json:"dynamicParameters,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *TuningState) Reset() {
	*x = TuningState{}
	mi := &file_tuning_tuning_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TuningState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState) ProtoMessage() {}

func (x *TuningState) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_tuning_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState.ProtoReflect.Descriptor instead.
func (*TuningState) Descriptor() ([]byte, []int) {
	return file_tuning_tuning_proto_rawDescGZIP(), []int{0}
}

func (x *TuningState) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TuningState) GetDynamicParameters() []*TuningState_Parameter {
	if x != nil {
		return x.DynamicParameters
	}
	return nil
}

type TuningState_Parameter struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Parameter:
	//
	//	*TuningState_Parameter_Number
	//	*TuningState_Parameter_String_
	Parameter     isTuningState_Parameter_Parameter `protobuf_oneof:"parameter"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TuningState_Parameter) Reset() {
	*x = TuningState_Parameter{}
	mi := &file_tuning_tuning_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TuningState_Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState_Parameter) ProtoMessage() {}

func (x *TuningState_Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_tuning_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState_Parameter.ProtoReflect.Descriptor instead.
func (*TuningState_Parameter) Descriptor() ([]byte, []int) {
	return file_tuning_tuning_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TuningState_Parameter) GetParameter() isTuningState_Parameter_Parameter {
	if x != nil {
		return x.Parameter
	}
	return nil
}

func (x *TuningState_Parameter) GetNumber() *TuningState_Parameter_NumberParameter {
	if x != nil {
		if x, ok := x.Parameter.(*TuningState_Parameter_Number); ok {
			return x.Number
		}
	}
	return nil
}

func (x *TuningState_Parameter) GetString_() *TuningState_Parameter_StringParameter {
	if x != nil {
		if x, ok := x.Parameter.(*TuningState_Parameter_String_); ok {
			return x.String_
		}
	}
	return nil
}

type isTuningState_Parameter_Parameter interface {
	isTuningState_Parameter_Parameter()
}

type TuningState_Parameter_Number struct {
	Number *TuningState_Parameter_NumberParameter `protobuf:"bytes,1,opt,name=number,proto3,oneof"`
}

type TuningState_Parameter_String_ struct {
	String_ *TuningState_Parameter_StringParameter `protobuf:"bytes,3,opt,name=string,proto3,oneof"`
}

func (*TuningState_Parameter_Number) isTuningState_Parameter_Parameter() {}

func (*TuningState_Parameter_String_) isTuningState_Parameter_Parameter() {}

// note: it may seem weird to not extract the key from the oneof, but this is so that the parser can easily determine the type of the parameter
// extracting it to a separate field on the same level as oneof would make it ambiguous
type TuningState_Parameter_NumberParameter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TuningState_Parameter_NumberParameter) Reset() {
	*x = TuningState_Parameter_NumberParameter{}
	mi := &file_tuning_tuning_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TuningState_Parameter_NumberParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState_Parameter_NumberParameter) ProtoMessage() {}

func (x *TuningState_Parameter_NumberParameter) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_tuning_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState_Parameter_NumberParameter.ProtoReflect.Descriptor instead.
func (*TuningState_Parameter_NumberParameter) Descriptor() ([]byte, []int) {
	return file_tuning_tuning_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *TuningState_Parameter_NumberParameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TuningState_Parameter_NumberParameter) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TuningState_Parameter_StringParameter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TuningState_Parameter_StringParameter) Reset() {
	*x = TuningState_Parameter_StringParameter{}
	mi := &file_tuning_tuning_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TuningState_Parameter_StringParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState_Parameter_StringParameter) ProtoMessage() {}

func (x *TuningState_Parameter_StringParameter) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_tuning_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState_Parameter_StringParameter.ProtoReflect.Descriptor instead.
func (*TuningState_Parameter_StringParameter) Descriptor() ([]byte, []int) {
	return file_tuning_tuning_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *TuningState_Parameter_StringParameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TuningState_Parameter_StringParameter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_tuning_tuning_proto protoreflect.FileDescriptor

const file_tuning_tuning_proto_rawDesc = "" +
	"\n" +
	"\x13tuning/tuning.proto\x12\rprotobuf_msgs\"\xb0\x03\n" +
	"\vTuningState\x12\x1c\n" +
	"\ttimestamp\x18\x01 \x01(\x04R\ttimestamp\x12R\n" +
	"\x11dynamicParameters\x18\x02 \x03(\v2$.protobuf_msgs.TuningState.ParameterR\x11dynamicParameters\x1a\xae\x02\n" +
	"\tParameter\x12N\n" +
	"\x06number\x18\x01 \x01(\v24.protobuf_msgs.TuningState.Parameter.NumberParameterH\x00R\x06number\x12N\n" +
	"\x06string\x18\x03 \x01(\v24.protobuf_msgs.TuningState.Parameter.StringParameterH\x00R\x06string\x1a9\n" +
	"\x0fNumberParameter\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x02R\x05value\x1a9\n" +
	"\x0fStringParameter\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05valueB\v\n" +
	"\tparameterB\x0fZ\rase/pb_tuningb\x06proto3"

var (
	file_tuning_tuning_proto_rawDescOnce sync.Once
	file_tuning_tuning_proto_rawDescData []byte
)

func file_tuning_tuning_proto_rawDescGZIP() []byte {
	file_tuning_tuning_proto_rawDescOnce.Do(func() {
		file_tuning_tuning_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tuning_tuning_proto_rawDesc), len(file_tuning_tuning_proto_rawDesc)))
	})
	return file_tuning_tuning_proto_rawDescData
}

var file_tuning_tuning_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tuning_tuning_proto_goTypes = []any{
	(*TuningState)(nil),                           // 0: protobuf_msgs.TuningState
	(*TuningState_Parameter)(nil),                 // 1: protobuf_msgs.TuningState.Parameter
	(*TuningState_Parameter_NumberParameter)(nil), // 2: protobuf_msgs.TuningState.Parameter.NumberParameter
	(*TuningState_Parameter_StringParameter)(nil), // 3: protobuf_msgs.TuningState.Parameter.StringParameter
}
var file_tuning_tuning_proto_depIdxs = []int32{
	1, // 0: protobuf_msgs.TuningState.dynamicParameters:type_name -> protobuf_msgs.TuningState.Parameter
	2, // 1: protobuf_msgs.TuningState.Parameter.number:type_name -> protobuf_msgs.TuningState.Parameter.NumberParameter
	3, // 2: protobuf_msgs.TuningState.Parameter.string:type_name -> protobuf_msgs.TuningState.Parameter.StringParameter
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tuning_tuning_proto_init() }
func file_tuning_tuning_proto_init() {
	if File_tuning_tuning_proto != nil {
		return
	}
	file_tuning_tuning_proto_msgTypes[1].OneofWrappers = []any{
		(*TuningState_Parameter_Number)(nil),
		(*TuningState_Parameter_String_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tuning_tuning_proto_rawDesc), len(file_tuning_tuning_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tuning_tuning_proto_goTypes,
		DependencyIndexes: file_tuning_tuning_proto_depIdxs,
		MessageInfos:      file_tuning_tuning_proto_msgTypes,
	}.Build()
	File_tuning_tuning_proto = out.File
	file_tuning_tuning_proto_goTypes = nil
	file_tuning_tuning_proto_depIdxs = nil
}
