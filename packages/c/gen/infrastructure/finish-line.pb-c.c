/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: infrastructure/finish-line.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "infrastructure/finish-line.pb-c.h"
void   protobuf_msgs__finish_line_event__init
                     (ProtobufMsgs__FinishLineEvent         *message)
{
  static const ProtobufMsgs__FinishLineEvent init_value = PROTOBUF_MSGS__FINISH_LINE_EVENT__INIT;
  *message = init_value;
}
size_t protobuf_msgs__finish_line_event__get_packed_size
                     (const ProtobufMsgs__FinishLineEvent *message)
{
  assert(message->base.descriptor == &protobuf_msgs__finish_line_event__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t protobuf_msgs__finish_line_event__pack
                     (const ProtobufMsgs__FinishLineEvent *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &protobuf_msgs__finish_line_event__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t protobuf_msgs__finish_line_event__pack_to_buffer
                     (const ProtobufMsgs__FinishLineEvent *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &protobuf_msgs__finish_line_event__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
ProtobufMsgs__FinishLineEvent *
       protobuf_msgs__finish_line_event__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (ProtobufMsgs__FinishLineEvent *)
     protobuf_c_message_unpack (&protobuf_msgs__finish_line_event__descriptor,
                                allocator, len, data);
}
void   protobuf_msgs__finish_line_event__free_unpacked
                     (ProtobufMsgs__FinishLineEvent *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &protobuf_msgs__finish_line_event__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor protobuf_msgs__finish_line_event__field_descriptors[1] =
{
  {
    "timestamp",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(ProtobufMsgs__FinishLineEvent, timestamp),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned protobuf_msgs__finish_line_event__field_indices_by_name[] = {
  0,   /* field[0] = timestamp */
};
static const ProtobufCIntRange protobuf_msgs__finish_line_event__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor protobuf_msgs__finish_line_event__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "protobuf_msgs.FinishLineEvent",
  "FinishLineEvent",
  "ProtobufMsgs__FinishLineEvent",
  "protobuf_msgs",
  sizeof(ProtobufMsgs__FinishLineEvent),
  1,
  protobuf_msgs__finish_line_event__field_descriptors,
  protobuf_msgs__finish_line_event__field_indices_by_name,
  1,  protobuf_msgs__finish_line_event__number_ranges,
  (ProtobufCMessageInit) protobuf_msgs__finish_line_event__init,
  NULL,NULL,NULL    /* reserved[123] */
};
