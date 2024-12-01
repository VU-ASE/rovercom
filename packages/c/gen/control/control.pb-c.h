/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: control/control.proto */

#ifndef PROTOBUF_C_control_2fcontrol_2eproto__INCLUDED
#define PROTOBUF_C_control_2fcontrol_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1003000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1003003 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _ProtobufMsgs__ConnectionState ProtobufMsgs__ConnectionState;


/* --- enums --- */


/* --- messages --- */

/*
 * Tell a client if a given client/rover is connected or not
 */
struct  _ProtobufMsgs__ConnectionState
{
  ProtobufCMessage base;
  char *client;
  protobuf_c_boolean connected;
  /*
   * the offset between the system time of the client/rover and the server
   */
  int64_t timestampoffset;
};
#define PROTOBUF_MSGS__CONNECTION_STATE__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__connection_state__descriptor) \
    , (char *)protobuf_c_empty_string, 0, 0 }


/* ProtobufMsgs__ConnectionState methods */
void   protobuf_msgs__connection_state__init
                     (ProtobufMsgs__ConnectionState         *message);
size_t protobuf_msgs__connection_state__get_packed_size
                     (const ProtobufMsgs__ConnectionState   *message);
size_t protobuf_msgs__connection_state__pack
                     (const ProtobufMsgs__ConnectionState   *message,
                      uint8_t             *out);
size_t protobuf_msgs__connection_state__pack_to_buffer
                     (const ProtobufMsgs__ConnectionState   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ConnectionState *
       protobuf_msgs__connection_state__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__connection_state__free_unpacked
                     (ProtobufMsgs__ConnectionState *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*ProtobufMsgs__ConnectionState_Closure)
                 (const ProtobufMsgs__ConnectionState *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor protobuf_msgs__connection_state__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_control_2fcontrol_2eproto__INCLUDED */
