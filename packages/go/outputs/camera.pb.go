// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: outputs/camera.proto

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

// Possible Objects the Imaging Service may detect
type DetectedObjects int32

const (
	DetectedObjects_FINISH_LINE        DetectedObjects = 0 // Finish_line_detected
	DetectedObjects_OFF_TRACK          DetectedObjects = 1 // Car no longer on the track
	DetectedObjects_OBSTACLE           DetectedObjects = 2 // Detected obstacle
	DetectedObjects_INTERSECTION       DetectedObjects = 3 // Detected intersection
	DetectedObjects_MISSING_LEFT_LANE  DetectedObjects = 4 // Can not find left lane
	DetectedObjects_MISSING_RIGHT_LANE DetectedObjects = 5 // Can not find right lane
	DetectedObjects_SHARP_RIGHT        DetectedObjects = 6 // 90 degree right turn
	DetectedObjects_SHARP_LEFT         DetectedObjects = 7 // 90 degree left turn
	DetectedObjects_U_TURN             DetectedObjects = 8 // Detected U turn
	DetectedObjects_S_TURN             DetectedObjects = 9 // Detected S turn (double u turn)
)

// Enum value maps for DetectedObjects.
var (
	DetectedObjects_name = map[int32]string{
		0: "FINISH_LINE",
		1: "OFF_TRACK",
		2: "OBSTACLE",
		3: "INTERSECTION",
		4: "MISSING_LEFT_LANE",
		5: "MISSING_RIGHT_LANE",
		6: "SHARP_RIGHT",
		7: "SHARP_LEFT",
		8: "U_TURN",
		9: "S_TURN",
	}
	DetectedObjects_value = map[string]int32{
		"FINISH_LINE":        0,
		"OFF_TRACK":          1,
		"OBSTACLE":           2,
		"INTERSECTION":       3,
		"MISSING_LEFT_LANE":  4,
		"MISSING_RIGHT_LANE": 5,
		"SHARP_RIGHT":        6,
		"SHARP_LEFT":         7,
		"U_TURN":             8,
		"S_TURN":             9,
	}
)

func (x DetectedObjects) Enum() *DetectedObjects {
	p := new(DetectedObjects)
	*p = x
	return p
}

func (x DetectedObjects) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DetectedObjects) Descriptor() protoreflect.EnumDescriptor {
	return file_outputs_camera_proto_enumTypes[0].Descriptor()
}

func (DetectedObjects) Type() protoreflect.EnumType {
	return &file_outputs_camera_proto_enumTypes[0]
}

func (x DetectedObjects) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DetectedObjects.Descriptor instead.
func (DetectedObjects) EnumDescriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{0}
}

// This is the message format that a camera-like service can send out. For example, the official ASE imaging service
// uses this output format. This can then be used by (for example) a controller, to determine how to steer, to stay
// on the track, or to detect obstacles, intersections, etc.
type CameraSensorOutput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Basic information, contains everything you need to know to steer and compute the middle of the track
	Resolution      *Resolution       `protobuf:"bytes,1,opt,name=resolution,proto3" json:"resolution,omitempty"`           // Resolution of the image in pixels
	HorizontalScans []*HorizontalScan `protobuf:"bytes,2,rep,name=horizontalScans,proto3" json:"horizontalScans,omitempty"` // Horizontal scans of the track, where each scan returns the track edges it finds in the image
	DetectedObjects []DetectedObjects `protobuf:"varint,3,rep,packed,name=detectedObjects,proto3,enum=protobuf_msgs.DetectedObjects" json:"detectedObjects,omitempty"`
	// Additional information that can be used to debug the image processing
	// if present, it is rendered in roverctl-web
	DebugFrame    *DebugFrame `protobuf:"bytes,4,opt,name=debugFrame,proto3" json:"debugFrame,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CameraSensorOutput) Reset() {
	*x = CameraSensorOutput{}
	mi := &file_outputs_camera_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CameraSensorOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraSensorOutput) ProtoMessage() {}

func (x *CameraSensorOutput) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_camera_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraSensorOutput.ProtoReflect.Descriptor instead.
func (*CameraSensorOutput) Descriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{0}
}

func (x *CameraSensorOutput) GetResolution() *Resolution {
	if x != nil {
		return x.Resolution
	}
	return nil
}

func (x *CameraSensorOutput) GetHorizontalScans() []*HorizontalScan {
	if x != nil {
		return x.HorizontalScans
	}
	return nil
}

func (x *CameraSensorOutput) GetDetectedObjects() []DetectedObjects {
	if x != nil {
		return x.DetectedObjects
	}
	return nil
}

func (x *CameraSensorOutput) GetDebugFrame() *DebugFrame {
	if x != nil {
		return x.DebugFrame
	}
	return nil
}

type Resolution struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Width         uint32                 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`   // Width of the image in pixels
	Height        uint32                 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"` // Height of the image in pixels
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Resolution) Reset() {
	*x = Resolution{}
	mi := &file_outputs_camera_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resolution) ProtoMessage() {}

func (x *Resolution) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_camera_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resolution.ProtoReflect.Descriptor instead.
func (*Resolution) Descriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{1}
}

func (x *Resolution) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Resolution) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type HorizontalScan struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	XLeft         uint32                 `protobuf:"varint,1,opt,name=xLeft,proto3" json:"xLeft,omitempty"`   // Leftmost point in the scan in pixels (is left edge of the track)
	XRight        uint32                 `protobuf:"varint,2,opt,name=xRight,proto3" json:"xRight,omitempty"` // Rightmost point in the scan in pixels (is right edge of the track)
	Y             uint32                 `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`           // Y coordinate of the scan in pixels
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HorizontalScan) Reset() {
	*x = HorizontalScan{}
	mi := &file_outputs_camera_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HorizontalScan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HorizontalScan) ProtoMessage() {}

func (x *HorizontalScan) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_camera_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HorizontalScan.ProtoReflect.Descriptor instead.
func (*HorizontalScan) Descriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{2}
}

func (x *HorizontalScan) GetXLeft() uint32 {
	if x != nil {
		return x.XLeft
	}
	return 0
}

func (x *HorizontalScan) GetXRight() uint32 {
	if x != nil {
		return x.XRight
	}
	return 0
}

func (x *HorizontalScan) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type DebugFrame struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// (Compressed) JPEG image of the camera output, useful for debugging
	// just JPEG bytes, that will be rendered in roverctl-web
	Jpeg []byte `protobuf:"bytes,1,opt,name=jpeg,proto3" json:"jpeg,omitempty"`
	// A "canvas" that you can "draw" on, for example by placing points, these are also rendered in roverctl-web
	Canvas        *Canvas `protobuf:"bytes,5,opt,name=canvas,proto3" json:"canvas,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DebugFrame) Reset() {
	*x = DebugFrame{}
	mi := &file_outputs_camera_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebugFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugFrame) ProtoMessage() {}

func (x *DebugFrame) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_camera_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugFrame.ProtoReflect.Descriptor instead.
func (*DebugFrame) Descriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{3}
}

func (x *DebugFrame) GetJpeg() []byte {
	if x != nil {
		return x.Jpeg
	}
	return nil
}

func (x *DebugFrame) GetCanvas() *Canvas {
	if x != nil {
		return x.Canvas
	}
	return nil
}

type CanvasObject struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Object:
	//
	//	*CanvasObject_Circle_
	Object        isCanvasObject_Object `protobuf_oneof:"object"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CanvasObject) Reset() {
	*x = CanvasObject{}
	mi := &file_outputs_camera_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CanvasObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanvasObject) ProtoMessage() {}

func (x *CanvasObject) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_camera_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanvasObject.ProtoReflect.Descriptor instead.
func (*CanvasObject) Descriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{4}
}

func (x *CanvasObject) GetObject() isCanvasObject_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *CanvasObject) GetCircle() *CanvasObject_Circle {
	if x != nil {
		if x, ok := x.Object.(*CanvasObject_Circle_); ok {
			return x.Circle
		}
	}
	return nil
}

type isCanvasObject_Object interface {
	isCanvasObject_Object()
}

type CanvasObject_Circle_ struct {
	Circle *CanvasObject_Circle `protobuf:"bytes,1,opt,name=circle,proto3,oneof"`
}

func (*CanvasObject_Circle_) isCanvasObject_Object() {}

type Canvas struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The width and height are a legacy feature, they should be the same as the resolution of the camera
	Width         uint32          `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height        uint32          `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Objects       []*CanvasObject `protobuf:"bytes,3,rep,name=objects,proto3" json:"objects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Canvas) Reset() {
	*x = Canvas{}
	mi := &file_outputs_camera_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Canvas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Canvas) ProtoMessage() {}

func (x *Canvas) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_camera_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Canvas.ProtoReflect.Descriptor instead.
func (*Canvas) Descriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{5}
}

func (x *Canvas) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Canvas) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Canvas) GetObjects() []*CanvasObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

type CanvasObject_Point struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             uint32                 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y             uint32                 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CanvasObject_Point) Reset() {
	*x = CanvasObject_Point{}
	mi := &file_outputs_camera_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CanvasObject_Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanvasObject_Point) ProtoMessage() {}

func (x *CanvasObject_Point) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_camera_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanvasObject_Point.ProtoReflect.Descriptor instead.
func (*CanvasObject_Point) Descriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{4, 0}
}

func (x *CanvasObject_Point) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *CanvasObject_Point) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type CanvasObject_Circle struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Center        *CanvasObject_Point    `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Radius        uint32                 `protobuf:"varint,2,opt,name=radius,proto3" json:"radius,omitempty"`
	Width         uint32                 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CanvasObject_Circle) Reset() {
	*x = CanvasObject_Circle{}
	mi := &file_outputs_camera_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CanvasObject_Circle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanvasObject_Circle) ProtoMessage() {}

func (x *CanvasObject_Circle) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_camera_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanvasObject_Circle.ProtoReflect.Descriptor instead.
func (*CanvasObject_Circle) Descriptor() ([]byte, []int) {
	return file_outputs_camera_proto_rawDescGZIP(), []int{4, 1}
}

func (x *CanvasObject_Circle) GetCenter() *CanvasObject_Point {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *CanvasObject_Circle) GetRadius() uint32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *CanvasObject_Circle) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

var File_outputs_camera_proto protoreflect.FileDescriptor

const file_outputs_camera_proto_rawDesc = "" +
	"\n" +
	"\x14outputs/camera.proto\x12\rprotobuf_msgs\"\x9d\x02\n" +
	"\x12CameraSensorOutput\x129\n" +
	"\n" +
	"resolution\x18\x01 \x01(\v2\x19.protobuf_msgs.ResolutionR\n" +
	"resolution\x12G\n" +
	"\x0fhorizontalScans\x18\x02 \x03(\v2\x1d.protobuf_msgs.HorizontalScanR\x0fhorizontalScans\x12H\n" +
	"\x0fdetectedObjects\x18\x03 \x03(\x0e2\x1e.protobuf_msgs.DetectedObjectsR\x0fdetectedObjects\x129\n" +
	"\n" +
	"debugFrame\x18\x04 \x01(\v2\x19.protobuf_msgs.DebugFrameR\n" +
	"debugFrame\":\n" +
	"\n" +
	"Resolution\x12\x14\n" +
	"\x05width\x18\x01 \x01(\rR\x05width\x12\x16\n" +
	"\x06height\x18\x02 \x01(\rR\x06height\"L\n" +
	"\x0eHorizontalScan\x12\x14\n" +
	"\x05xLeft\x18\x01 \x01(\rR\x05xLeft\x12\x16\n" +
	"\x06xRight\x18\x02 \x01(\rR\x06xRight\x12\f\n" +
	"\x01y\x18\x03 \x01(\rR\x01y\"O\n" +
	"\n" +
	"DebugFrame\x12\x12\n" +
	"\x04jpeg\x18\x01 \x01(\fR\x04jpeg\x12-\n" +
	"\x06canvas\x18\x05 \x01(\v2\x15.protobuf_msgs.CanvasR\x06canvas\"\xee\x01\n" +
	"\fCanvasObject\x12<\n" +
	"\x06circle\x18\x01 \x01(\v2\".protobuf_msgs.CanvasObject.CircleH\x00R\x06circle\x1a#\n" +
	"\x05Point\x12\f\n" +
	"\x01x\x18\x01 \x01(\rR\x01x\x12\f\n" +
	"\x01y\x18\x02 \x01(\rR\x01y\x1aq\n" +
	"\x06Circle\x129\n" +
	"\x06center\x18\x01 \x01(\v2!.protobuf_msgs.CanvasObject.PointR\x06center\x12\x16\n" +
	"\x06radius\x18\x02 \x01(\rR\x06radius\x12\x14\n" +
	"\x05width\x18\x03 \x01(\rR\x05widthB\b\n" +
	"\x06object\"m\n" +
	"\x06Canvas\x12\x14\n" +
	"\x05width\x18\x01 \x01(\rR\x05width\x12\x16\n" +
	"\x06height\x18\x02 \x01(\rR\x06height\x125\n" +
	"\aobjects\x18\x03 \x03(\v2\x1b.protobuf_msgs.CanvasObjectR\aobjects*\xb9\x01\n" +
	"\x0fDetectedObjects\x12\x0f\n" +
	"\vFINISH_LINE\x10\x00\x12\r\n" +
	"\tOFF_TRACK\x10\x01\x12\f\n" +
	"\bOBSTACLE\x10\x02\x12\x10\n" +
	"\fINTERSECTION\x10\x03\x12\x15\n" +
	"\x11MISSING_LEFT_LANE\x10\x04\x12\x16\n" +
	"\x12MISSING_RIGHT_LANE\x10\x05\x12\x0f\n" +
	"\vSHARP_RIGHT\x10\x06\x12\x0e\n" +
	"\n" +
	"SHARP_LEFT\x10\a\x12\n" +
	"\n" +
	"\x06U_TURN\x10\b\x12\n" +
	"\n" +
	"\x06S_TURN\x10\tB\x10Z\x0ease/pb_outputsb\x06proto3"

var (
	file_outputs_camera_proto_rawDescOnce sync.Once
	file_outputs_camera_proto_rawDescData []byte
)

func file_outputs_camera_proto_rawDescGZIP() []byte {
	file_outputs_camera_proto_rawDescOnce.Do(func() {
		file_outputs_camera_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_outputs_camera_proto_rawDesc), len(file_outputs_camera_proto_rawDesc)))
	})
	return file_outputs_camera_proto_rawDescData
}

var file_outputs_camera_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_outputs_camera_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_outputs_camera_proto_goTypes = []any{
	(DetectedObjects)(0),        // 0: protobuf_msgs.DetectedObjects
	(*CameraSensorOutput)(nil),  // 1: protobuf_msgs.CameraSensorOutput
	(*Resolution)(nil),          // 2: protobuf_msgs.Resolution
	(*HorizontalScan)(nil),      // 3: protobuf_msgs.HorizontalScan
	(*DebugFrame)(nil),          // 4: protobuf_msgs.DebugFrame
	(*CanvasObject)(nil),        // 5: protobuf_msgs.CanvasObject
	(*Canvas)(nil),              // 6: protobuf_msgs.Canvas
	(*CanvasObject_Point)(nil),  // 7: protobuf_msgs.CanvasObject.Point
	(*CanvasObject_Circle)(nil), // 8: protobuf_msgs.CanvasObject.Circle
}
var file_outputs_camera_proto_depIdxs = []int32{
	2, // 0: protobuf_msgs.CameraSensorOutput.resolution:type_name -> protobuf_msgs.Resolution
	3, // 1: protobuf_msgs.CameraSensorOutput.horizontalScans:type_name -> protobuf_msgs.HorizontalScan
	0, // 2: protobuf_msgs.CameraSensorOutput.detectedObjects:type_name -> protobuf_msgs.DetectedObjects
	4, // 3: protobuf_msgs.CameraSensorOutput.debugFrame:type_name -> protobuf_msgs.DebugFrame
	6, // 4: protobuf_msgs.DebugFrame.canvas:type_name -> protobuf_msgs.Canvas
	8, // 5: protobuf_msgs.CanvasObject.circle:type_name -> protobuf_msgs.CanvasObject.Circle
	5, // 6: protobuf_msgs.Canvas.objects:type_name -> protobuf_msgs.CanvasObject
	7, // 7: protobuf_msgs.CanvasObject.Circle.center:type_name -> protobuf_msgs.CanvasObject.Point
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_outputs_camera_proto_init() }
func file_outputs_camera_proto_init() {
	if File_outputs_camera_proto != nil {
		return
	}
	file_outputs_camera_proto_msgTypes[4].OneofWrappers = []any{
		(*CanvasObject_Circle_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_outputs_camera_proto_rawDesc), len(file_outputs_camera_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outputs_camera_proto_goTypes,
		DependencyIndexes: file_outputs_camera_proto_depIdxs,
		EnumInfos:         file_outputs_camera_proto_enumTypes,
		MessageInfos:      file_outputs_camera_proto_msgTypes,
	}.Build()
	File_outputs_camera_proto = out.File
	file_outputs_camera_proto_goTypes = nil
	file_outputs_camera_proto_depIdxs = nil
}
