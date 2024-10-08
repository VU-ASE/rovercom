/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: core/servicediscovery.proto */

#ifndef PROTOBUF_C_core_2fservicediscovery_2eproto__INCLUDED
#define PROTOBUF_C_core_2fservicediscovery_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1003000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1003003 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _ProtobufMsgs__ServiceIdentifier ProtobufMsgs__ServiceIdentifier;
typedef struct _ProtobufMsgs__ServiceEndpoint ProtobufMsgs__ServiceEndpoint;
typedef struct _ProtobufMsgs__ServiceOption ProtobufMsgs__ServiceOption;
typedef struct _ProtobufMsgs__ServiceDependency ProtobufMsgs__ServiceDependency;
typedef struct _ProtobufMsgs__Service ProtobufMsgs__Service;
typedef struct _ProtobufMsgs__ServiceInformationRequest ProtobufMsgs__ServiceInformationRequest;
typedef struct _ProtobufMsgs__ServiceOrder ProtobufMsgs__ServiceOrder;
typedef struct _ProtobufMsgs__ServiceListRequest ProtobufMsgs__ServiceListRequest;
typedef struct _ProtobufMsgs__ServiceList ProtobufMsgs__ServiceList;
typedef struct _ProtobufMsgs__ServiceStatusUpdate ProtobufMsgs__ServiceStatusUpdate;


/* --- enums --- */

typedef enum _ProtobufMsgs__ServiceOption__Type {
  PROTOBUF_MSGS__SERVICE_OPTION__TYPE__STRING = 0,
  PROTOBUF_MSGS__SERVICE_OPTION__TYPE__INT = 1,
  PROTOBUF_MSGS__SERVICE_OPTION__TYPE__FLOAT = 2
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(PROTOBUF_MSGS__SERVICE_OPTION__TYPE)
} ProtobufMsgs__ServiceOption__Type;
typedef enum _ProtobufMsgs__ServiceOrder__OrderType {
  /*
   * will attempt a graceful shutdown
   */
  PROTOBUF_MSGS__SERVICE_ORDER__ORDER_TYPE__STOP = 0,
  /*
   * will kill the service immediately
   */
  PROTOBUF_MSGS__SERVICE_ORDER__ORDER_TYPE__KILL = 1,
  /*
   * will kill the service immediately and restart
   */
  PROTOBUF_MSGS__SERVICE_ORDER__ORDER_TYPE__FORCE_RESTART = 2
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(PROTOBUF_MSGS__SERVICE_ORDER__ORDER_TYPE)
} ProtobufMsgs__ServiceOrder__OrderType;
typedef enum _ProtobufMsgs__ServiceStatus {
  PROTOBUF_MSGS__SERVICE_STATUS__UNKNOWN = 0,
  /*
   * Registered, but not running yet (probably waiting for dependencies)
   */
  PROTOBUF_MSGS__SERVICE_STATUS__REGISTERED = 1,
  /*
   * Currently running (after registration)
   */
  PROTOBUF_MSGS__SERVICE_STATUS__RUNNING = 2,
  /*
   * Stopped gracefully
   */
  PROTOBUF_MSGS__SERVICE_STATUS__STOPPED = 3,
  /*
   * Not registered yet (useful if you are waiting for this dependency)
   */
  PROTOBUF_MSGS__SERVICE_STATUS__NOT_REGISTERED = 4
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(PROTOBUF_MSGS__SERVICE_STATUS)
} ProtobufMsgs__ServiceStatus;

/* --- messages --- */

/*
 * Used to identify a service within the system
 */
struct  _ProtobufMsgs__ServiceIdentifier
{
  ProtobufCMessage base;
  char *name;
  int32_t pid;
};
#define PROTOBUF_MSGS__SERVICE_IDENTIFIER__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_identifier__descriptor) \
    , (char *)protobuf_c_empty_string, 0 }


/*
 * An endpoint that is made available by a service
 */
struct  _ProtobufMsgs__ServiceEndpoint
{
  ProtobufCMessage base;
  /*
   * the identifier to select this endpoint
   */
  char *name;
  /*
   * the zmq address to connect to
   */
  char *address;
};
#define PROTOBUF_MSGS__SERVICE_ENDPOINT__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_endpoint__descriptor) \
    , (char *)protobuf_c_empty_string, (char *)protobuf_c_empty_string }


/*
 * The options that can be set for a service
 */
struct  _ProtobufMsgs__ServiceOption
{
  ProtobufCMessage base;
  char *name;
  ProtobufMsgs__ServiceOption__Type type;
  protobuf_c_boolean mutable_;
  /*
   * should be set and checked based on the type
   */
  char *string_default;
  int32_t int_default;
  float float_default;
};
#define PROTOBUF_MSGS__SERVICE_OPTION__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_option__descriptor) \
    , (char *)protobuf_c_empty_string, PROTOBUF_MSGS__SERVICE_OPTION__TYPE__STRING, 0, (char *)protobuf_c_empty_string, 0, 0 }


/*
 * The core knows the dependencies of each service, to build a dependency graph
 */
struct  _ProtobufMsgs__ServiceDependency
{
  ProtobufCMessage base;
  char *servicename;
  char *outputname;
};
#define PROTOBUF_MSGS__SERVICE_DEPENDENCY__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_dependency__descriptor) \
    , (char *)protobuf_c_empty_string, (char *)protobuf_c_empty_string }


/*
 * A description of a service, sent by a service to register itself or broadcasted by the core
 */
struct  _ProtobufMsgs__Service
{
  ProtobufCMessage base;
  ProtobufMsgs__ServiceIdentifier *identifier;
  size_t n_endpoints;
  ProtobufMsgs__ServiceEndpoint **endpoints;
  size_t n_options;
  ProtobufMsgs__ServiceOption **options;
  size_t n_dependencies;
  ProtobufMsgs__ServiceDependency **dependencies;
  ProtobufMsgs__ServiceStatus status;
  int64_t registeredat;
};
#define PROTOBUF_MSGS__SERVICE__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service__descriptor) \
    , NULL, 0,NULL, 0,NULL, 0,NULL, PROTOBUF_MSGS__SERVICE_STATUS__UNKNOWN, 0 }


/*
 * When a service requests information about other services, it sends an InformationRequest message
 */
struct  _ProtobufMsgs__ServiceInformationRequest
{
  ProtobufCMessage base;
  ProtobufMsgs__ServiceIdentifier *requested;
};
#define PROTOBUF_MSGS__SERVICE_INFORMATION_REQUEST__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_information_request__descriptor) \
    , NULL }


/*
 * The core can order services to stop/restart by sending a service order
 */
struct  _ProtobufMsgs__ServiceOrder
{
  ProtobufCMessage base;
  /*
   * The service this order is for
   */
  ProtobufMsgs__ServiceIdentifier *service;
  ProtobufMsgs__ServiceOrder__OrderType order;
};
#define PROTOBUF_MSGS__SERVICE_ORDER__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_order__descriptor) \
    , NULL, PROTOBUF_MSGS__SERVICE_ORDER__ORDER_TYPE__STOP }


/*
 * When a service wants to fetch all services, it sends a ServiceListRequest
 */
struct  _ProtobufMsgs__ServiceListRequest
{
  ProtobufCMessage base;
  ProtobufMsgs__ServiceIdentifier *requested;
};
#define PROTOBUF_MSGS__SERVICE_LIST_REQUEST__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_list_request__descriptor) \
    , NULL }


/*
 * When the core sends a list of services, it sends a ServiceList
 */
struct  _ProtobufMsgs__ServiceList
{
  ProtobufCMessage base;
  size_t n_services;
  ProtobufMsgs__Service **services;
};
#define PROTOBUF_MSGS__SERVICE_LIST__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_list__descriptor) \
    , 0,NULL }


/*
 * This will inform the core of a status update
 */
struct  _ProtobufMsgs__ServiceStatusUpdate
{
  ProtobufCMessage base;
  ProtobufMsgs__ServiceIdentifier *service;
  ProtobufMsgs__ServiceStatus status;
};
#define PROTOBUF_MSGS__SERVICE_STATUS_UPDATE__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&protobuf_msgs__service_status_update__descriptor) \
    , NULL, PROTOBUF_MSGS__SERVICE_STATUS__UNKNOWN }


/* ProtobufMsgs__ServiceIdentifier methods */
void   protobuf_msgs__service_identifier__init
                     (ProtobufMsgs__ServiceIdentifier         *message);
size_t protobuf_msgs__service_identifier__get_packed_size
                     (const ProtobufMsgs__ServiceIdentifier   *message);
size_t protobuf_msgs__service_identifier__pack
                     (const ProtobufMsgs__ServiceIdentifier   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_identifier__pack_to_buffer
                     (const ProtobufMsgs__ServiceIdentifier   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceIdentifier *
       protobuf_msgs__service_identifier__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_identifier__free_unpacked
                     (ProtobufMsgs__ServiceIdentifier *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__ServiceEndpoint methods */
void   protobuf_msgs__service_endpoint__init
                     (ProtobufMsgs__ServiceEndpoint         *message);
size_t protobuf_msgs__service_endpoint__get_packed_size
                     (const ProtobufMsgs__ServiceEndpoint   *message);
size_t protobuf_msgs__service_endpoint__pack
                     (const ProtobufMsgs__ServiceEndpoint   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_endpoint__pack_to_buffer
                     (const ProtobufMsgs__ServiceEndpoint   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceEndpoint *
       protobuf_msgs__service_endpoint__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_endpoint__free_unpacked
                     (ProtobufMsgs__ServiceEndpoint *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__ServiceOption methods */
void   protobuf_msgs__service_option__init
                     (ProtobufMsgs__ServiceOption         *message);
size_t protobuf_msgs__service_option__get_packed_size
                     (const ProtobufMsgs__ServiceOption   *message);
size_t protobuf_msgs__service_option__pack
                     (const ProtobufMsgs__ServiceOption   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_option__pack_to_buffer
                     (const ProtobufMsgs__ServiceOption   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceOption *
       protobuf_msgs__service_option__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_option__free_unpacked
                     (ProtobufMsgs__ServiceOption *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__ServiceDependency methods */
void   protobuf_msgs__service_dependency__init
                     (ProtobufMsgs__ServiceDependency         *message);
size_t protobuf_msgs__service_dependency__get_packed_size
                     (const ProtobufMsgs__ServiceDependency   *message);
size_t protobuf_msgs__service_dependency__pack
                     (const ProtobufMsgs__ServiceDependency   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_dependency__pack_to_buffer
                     (const ProtobufMsgs__ServiceDependency   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceDependency *
       protobuf_msgs__service_dependency__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_dependency__free_unpacked
                     (ProtobufMsgs__ServiceDependency *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__Service methods */
void   protobuf_msgs__service__init
                     (ProtobufMsgs__Service         *message);
size_t protobuf_msgs__service__get_packed_size
                     (const ProtobufMsgs__Service   *message);
size_t protobuf_msgs__service__pack
                     (const ProtobufMsgs__Service   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service__pack_to_buffer
                     (const ProtobufMsgs__Service   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__Service *
       protobuf_msgs__service__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service__free_unpacked
                     (ProtobufMsgs__Service *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__ServiceInformationRequest methods */
void   protobuf_msgs__service_information_request__init
                     (ProtobufMsgs__ServiceInformationRequest         *message);
size_t protobuf_msgs__service_information_request__get_packed_size
                     (const ProtobufMsgs__ServiceInformationRequest   *message);
size_t protobuf_msgs__service_information_request__pack
                     (const ProtobufMsgs__ServiceInformationRequest   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_information_request__pack_to_buffer
                     (const ProtobufMsgs__ServiceInformationRequest   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceInformationRequest *
       protobuf_msgs__service_information_request__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_information_request__free_unpacked
                     (ProtobufMsgs__ServiceInformationRequest *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__ServiceOrder methods */
void   protobuf_msgs__service_order__init
                     (ProtobufMsgs__ServiceOrder         *message);
size_t protobuf_msgs__service_order__get_packed_size
                     (const ProtobufMsgs__ServiceOrder   *message);
size_t protobuf_msgs__service_order__pack
                     (const ProtobufMsgs__ServiceOrder   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_order__pack_to_buffer
                     (const ProtobufMsgs__ServiceOrder   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceOrder *
       protobuf_msgs__service_order__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_order__free_unpacked
                     (ProtobufMsgs__ServiceOrder *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__ServiceListRequest methods */
void   protobuf_msgs__service_list_request__init
                     (ProtobufMsgs__ServiceListRequest         *message);
size_t protobuf_msgs__service_list_request__get_packed_size
                     (const ProtobufMsgs__ServiceListRequest   *message);
size_t protobuf_msgs__service_list_request__pack
                     (const ProtobufMsgs__ServiceListRequest   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_list_request__pack_to_buffer
                     (const ProtobufMsgs__ServiceListRequest   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceListRequest *
       protobuf_msgs__service_list_request__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_list_request__free_unpacked
                     (ProtobufMsgs__ServiceListRequest *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__ServiceList methods */
void   protobuf_msgs__service_list__init
                     (ProtobufMsgs__ServiceList         *message);
size_t protobuf_msgs__service_list__get_packed_size
                     (const ProtobufMsgs__ServiceList   *message);
size_t protobuf_msgs__service_list__pack
                     (const ProtobufMsgs__ServiceList   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_list__pack_to_buffer
                     (const ProtobufMsgs__ServiceList   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceList *
       protobuf_msgs__service_list__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_list__free_unpacked
                     (ProtobufMsgs__ServiceList *message,
                      ProtobufCAllocator *allocator);
/* ProtobufMsgs__ServiceStatusUpdate methods */
void   protobuf_msgs__service_status_update__init
                     (ProtobufMsgs__ServiceStatusUpdate         *message);
size_t protobuf_msgs__service_status_update__get_packed_size
                     (const ProtobufMsgs__ServiceStatusUpdate   *message);
size_t protobuf_msgs__service_status_update__pack
                     (const ProtobufMsgs__ServiceStatusUpdate   *message,
                      uint8_t             *out);
size_t protobuf_msgs__service_status_update__pack_to_buffer
                     (const ProtobufMsgs__ServiceStatusUpdate   *message,
                      ProtobufCBuffer     *buffer);
ProtobufMsgs__ServiceStatusUpdate *
       protobuf_msgs__service_status_update__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   protobuf_msgs__service_status_update__free_unpacked
                     (ProtobufMsgs__ServiceStatusUpdate *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*ProtobufMsgs__ServiceIdentifier_Closure)
                 (const ProtobufMsgs__ServiceIdentifier *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__ServiceEndpoint_Closure)
                 (const ProtobufMsgs__ServiceEndpoint *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__ServiceOption_Closure)
                 (const ProtobufMsgs__ServiceOption *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__ServiceDependency_Closure)
                 (const ProtobufMsgs__ServiceDependency *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__Service_Closure)
                 (const ProtobufMsgs__Service *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__ServiceInformationRequest_Closure)
                 (const ProtobufMsgs__ServiceInformationRequest *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__ServiceOrder_Closure)
                 (const ProtobufMsgs__ServiceOrder *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__ServiceListRequest_Closure)
                 (const ProtobufMsgs__ServiceListRequest *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__ServiceList_Closure)
                 (const ProtobufMsgs__ServiceList *message,
                  void *closure_data);
typedef void (*ProtobufMsgs__ServiceStatusUpdate_Closure)
                 (const ProtobufMsgs__ServiceStatusUpdate *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCEnumDescriptor    protobuf_msgs__service_status__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_identifier__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_endpoint__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_option__descriptor;
extern const ProtobufCEnumDescriptor    protobuf_msgs__service_option__type__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_dependency__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_information_request__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_order__descriptor;
extern const ProtobufCEnumDescriptor    protobuf_msgs__service_order__order_type__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_list_request__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_list__descriptor;
extern const ProtobufCMessageDescriptor protobuf_msgs__service_status_update__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_core_2fservicediscovery_2eproto__INCLUDED */
