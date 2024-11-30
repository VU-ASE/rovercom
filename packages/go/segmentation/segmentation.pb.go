// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: segmentation/segmentation.proto

package pb_segmentation

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

// Control messages exchanged by client(s), the server and the car to send data in multiple segments
type Segment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketId      int64  `protobuf:"varint,1,opt,name=packetId,proto3" json:"packetId,omitempty"`           // unique and increasing id of each the packet that this segment is part of
	SegmentId     int64  `protobuf:"varint,2,opt,name=segmentId,proto3" json:"segmentId,omitempty"`         // unique and increasing id of the segment within the packet
	TotalSegments int64  `protobuf:"varint,3,opt,name=totalSegments,proto3" json:"totalSegments,omitempty"` // total number of segments in the packet
	Data          []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                    // data of the segment
}

func (x *Segment) Reset() {
	*x = Segment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segmentation_segmentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Segment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Segment) ProtoMessage() {}

func (x *Segment) ProtoReflect() protoreflect.Message {
	mi := &file_segmentation_segmentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Segment.ProtoReflect.Descriptor instead.
func (*Segment) Descriptor() ([]byte, []int) {
	return file_segmentation_segmentation_proto_rawDescGZIP(), []int{0}
}

func (x *Segment) GetPacketId() int64 {
	if x != nil {
		return x.PacketId
	}
	return 0
}

func (x *Segment) GetSegmentId() int64 {
	if x != nil {
		return x.SegmentId
	}
	return 0
}

func (x *Segment) GetTotalSegments() int64 {
	if x != nil {
		return x.TotalSegments
	}
	return 0
}

func (x *Segment) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_segmentation_segmentation_proto protoreflect.FileDescriptor

var file_segmentation_segmentation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73,
	0x22, 0x7d, 0x0a, 0x07, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x15, 0x5a, 0x13, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x62, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_segmentation_segmentation_proto_rawDescOnce sync.Once
	file_segmentation_segmentation_proto_rawDescData = file_segmentation_segmentation_proto_rawDesc
)

func file_segmentation_segmentation_proto_rawDescGZIP() []byte {
	file_segmentation_segmentation_proto_rawDescOnce.Do(func() {
		file_segmentation_segmentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_segmentation_segmentation_proto_rawDescData)
	})
	return file_segmentation_segmentation_proto_rawDescData
}

var file_segmentation_segmentation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_segmentation_segmentation_proto_goTypes = []any{
	(*Segment)(nil), // 0: protobuf_msgs.Segment
}
var file_segmentation_segmentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_segmentation_segmentation_proto_init() }
func file_segmentation_segmentation_proto_init() {
	if File_segmentation_segmentation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_segmentation_segmentation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Segment); i {
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
			RawDescriptor: file_segmentation_segmentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_segmentation_segmentation_proto_goTypes,
		DependencyIndexes: file_segmentation_segmentation_proto_depIdxs,
		MessageInfos:      file_segmentation_segmentation_proto_msgTypes,
	}.Build()
	File_segmentation_segmentation_proto = out.File
	file_segmentation_segmentation_proto_rawDesc = nil
	file_segmentation_segmentation_proto_goTypes = nil
	file_segmentation_segmentation_proto_depIdxs = nil
}
