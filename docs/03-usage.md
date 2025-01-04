# Usage

## In Go

Import one of the generated `pb_*` packages, which correspond to the message folder collections and prefix it:

```go
import (
	pb_outputs "github.com/VU-ASE/rovercom/packages/go/outputs"
)
```

Then, instantiate an object like so:

```go
msg := pb_outputs.SensorOutput{
    Timestamp: uint64(time.Now().UnixMilli()), // milliseconds since epoch
    Status:    0,                              // all is well
    SensorId:  1,                              // this is the first and only sensor we have
    SensorOutput: &pb_outputs.SensorOutput_ControllerOutput{
        ControllerOutput: &pb_outputs.ControllerOutput{
            SteeringAngle: steerPosition,
            LeftThrottle:  0,
            RightThrottle: 0,
            FanSpeed:      0,
            FrontLights:   false,
        },
    },
}
```

## For C

Include the header files for the objects that you want to instantiate:

```c
#include "./outputs/wrapper.pb-c.h"
#include "./outputs/camera.pb-c.h"
```

Then, instantiate an object like so:

```c
// Initialize the message that we want to send to the controller
ProtobufMsgs__SensorOutput actuator_msg = PROTOBUF_MSGS__SENSOR_OUTPUT__INIT;
// Set the message fields
msg->timestamp = current_time_millis();   // milliseconds since epoch
msg->status = 0;                          // all is well
msg->sensorid = 1;                        // this is the first and only sensor we have
// Set the oneof field contents
ProtobufMsgs__ControllerOutput controller_output = PROTOBUF_MSGS__CONTROLLER_OUTPUT__INIT;
controller_output.steeringangle = steerPosition;
controller_output.leftthrottle = 0;
controller_output.rightthrottle = 0;
controller_output.fanspeed = 0;
controller_output.frontlights = false;
// Set the oneof field (union)
msg->controlleroutput = &controller_output;
msg->sensor_output_case = PROTOBUF_MSGS__SENSOR_OUTPUT__SENSOR_OUTPUT_CONTROLLER_OUTPUT;
```

## For TypeScirpt

Import the objects that you want to instantiate:

```typescript
import { SensorOutput } from 'ase-rovercom/gen/outputs/wrapper';
import { ControllerOutput } from 'ase-rovercom/gen/outputs/controller';
```

Then, instantiate an object like so:

```typescript
const controllerOutput : ControllerOutput = {
    steeringAngle: 0,
    leftThrottle: 0,
    rightThrottle: 0,
    fanSpeed: 0,
    frontLights: false,
    rawError: 0,
    scaledError: 0,
}

const sensorOutput: SensorOutput = {
    timestamp: 0,
    sensorId: 1,
    status: 0,
    controllerOutput: controllerOutput,
}
```