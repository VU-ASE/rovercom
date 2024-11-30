/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: segmentation/segmentation.proto */

#ifndef PROTOBUF_C_segmentation_2fsegmentation_2eproto__INCLUDED
#define PROTOBUF_C_segmentation_2fsegmentation_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1003000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1003003 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _ProtobufMsgs__Segment ProtobufMsgs__Segment;


/* --- enums --- */


/* --- messages --- */

/*
 * 
 * Control messages exchanged by client(s), the server and the car to send data in multiple segments
 * 
 */
struct  _ProtobufMsgs__Segment
{
  ProtobufCMessage base;
  /*
   * unique and increasing id of each the packet that this segment is part of
   */
  int64_t packetid;
  /*
   * unique and increasing id of the segment within the packet
   */
  int64_t segmentid;
  /*
   * total number of segments in the packet
   */
  int64_t totalsegments;
  /*
   * data of the segment
   */
  ProtobufCBinaryData data;
};
#define PROTOBUF_MSGS__SEGMENT__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__segment__descriptor) \
    , 0, 0, 0, {0,NULL} }


/* ProtobufMsgs__Segment methods */
void   protobuf_msgs__segment__init
                     (ProtobufMsgs__Segment         *message);
size_t protobuf_msgs__segment__get_packed_size
                     (const ProtobufMsgs__Segment   *message);
size_t protobuf_msgs__segment__pack
                     (const ProtobufMsgs__Segment   *message,
                      uint8_t             *out);
size_t protobuf_msgs__segment__pack_to_buffer
                     (const ProtobufMsgs__Segment   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__Segment *
       protobuf_msgs__segment__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__segment__free_unpacked
                     (ProtobufMsgs__Segment *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*ProtobufMsgs__Segment_Closure)
                 (const ProtobufMsgs__Segment *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor protobuf_msgs__segment__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_segmentation_2fsegmentation_2eproto__INCLUDED */