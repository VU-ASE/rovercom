// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v3.12.4
// source: outputs/wrapper.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { BatterySensorOutput } from "./battery";
import { CameraSensorOutput } from "./camera";
import { ControllerOutput } from "./controller";
import { DistanceSensorOutput } from "./distance";
import { ImuSensorOutput } from "./imu";
import { LapTimeOutput } from "./laptime";
import { LuxSensorOutput } from "./lux";
import { RpmSensorOutput } from "./rpm";
import { SpeedSensorOutput } from "./speed";

export const protobufPackage = "protobuf_msgs";

export interface SensorOutput {
  /** Every sensor has a unique ID to support multiple sensors of the same type */
  sensorId: number;
  /** Add a timestamp to the output to make debugging, logging and synchronisation easier */
  timestamp: number;
  /**
   * Report an error if the sensor is not working correctly (controller can decide to ignore or stop the car)
   * 0 = no error, any other value = error
   */
  status: number;
  cameraOutput?: CameraSensorOutput | undefined;
  distanceOutput?: DistanceSensorOutput | undefined;
  speedOutput?: SpeedSensorOutput | undefined;
  controllerOutput?: ControllerOutput | undefined;
  imuOutput?: ImuSensorOutput | undefined;
  batteryOutput?: BatterySensorOutput | undefined;
  rpmOuput?: RpmSensorOutput | undefined;
  luxOutput?: LuxSensorOutput | undefined;
  laptimeOutput?: LapTimeOutput | undefined;
}

function createBaseSensorOutput(): SensorOutput {
  return {
    sensorId: 0,
    timestamp: 0,
    status: 0,
    cameraOutput: undefined,
    distanceOutput: undefined,
    speedOutput: undefined,
    controllerOutput: undefined,
    imuOutput: undefined,
    batteryOutput: undefined,
    rpmOuput: undefined,
    luxOutput: undefined,
    laptimeOutput: undefined,
  };
}

export const SensorOutput: MessageFns<SensorOutput> = {
  encode(message: SensorOutput, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.sensorId !== 0) {
      writer.uint32(8).uint32(message.sensorId);
    }
    if (message.timestamp !== 0) {
      writer.uint32(16).uint64(message.timestamp);
    }
    if (message.status !== 0) {
      writer.uint32(24).uint32(message.status);
    }
    if (message.cameraOutput !== undefined) {
      CameraSensorOutput.encode(message.cameraOutput, writer.uint32(34).fork()).join();
    }
    if (message.distanceOutput !== undefined) {
      DistanceSensorOutput.encode(message.distanceOutput, writer.uint32(42).fork()).join();
    }
    if (message.speedOutput !== undefined) {
      SpeedSensorOutput.encode(message.speedOutput, writer.uint32(50).fork()).join();
    }
    if (message.controllerOutput !== undefined) {
      ControllerOutput.encode(message.controllerOutput, writer.uint32(58).fork()).join();
    }
    if (message.imuOutput !== undefined) {
      ImuSensorOutput.encode(message.imuOutput, writer.uint32(66).fork()).join();
    }
    if (message.batteryOutput !== undefined) {
      BatterySensorOutput.encode(message.batteryOutput, writer.uint32(74).fork()).join();
    }
    if (message.rpmOuput !== undefined) {
      RpmSensorOutput.encode(message.rpmOuput, writer.uint32(82).fork()).join();
    }
    if (message.luxOutput !== undefined) {
      LuxSensorOutput.encode(message.luxOutput, writer.uint32(90).fork()).join();
    }
    if (message.laptimeOutput !== undefined) {
      LapTimeOutput.encode(message.laptimeOutput, writer.uint32(98).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SensorOutput {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSensorOutput();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.sensorId = reader.uint32();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.timestamp = longToNumber(reader.uint64());
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.status = reader.uint32();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.cameraOutput = CameraSensorOutput.decode(reader, reader.uint32());
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.distanceOutput = DistanceSensorOutput.decode(reader, reader.uint32());
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.speedOutput = SpeedSensorOutput.decode(reader, reader.uint32());
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.controllerOutput = ControllerOutput.decode(reader, reader.uint32());
          continue;
        }
        case 8: {
          if (tag !== 66) {
            break;
          }

          message.imuOutput = ImuSensorOutput.decode(reader, reader.uint32());
          continue;
        }
        case 9: {
          if (tag !== 74) {
            break;
          }

          message.batteryOutput = BatterySensorOutput.decode(reader, reader.uint32());
          continue;
        }
        case 10: {
          if (tag !== 82) {
            break;
          }

          message.rpmOuput = RpmSensorOutput.decode(reader, reader.uint32());
          continue;
        }
        case 11: {
          if (tag !== 90) {
            break;
          }

          message.luxOutput = LuxSensorOutput.decode(reader, reader.uint32());
          continue;
        }
        case 12: {
          if (tag !== 98) {
            break;
          }

          message.laptimeOutput = LapTimeOutput.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SensorOutput {
    return {
      sensorId: isSet(object.sensorId) ? globalThis.Number(object.sensorId) : 0,
      timestamp: isSet(object.timestamp) ? globalThis.Number(object.timestamp) : 0,
      status: isSet(object.status) ? globalThis.Number(object.status) : 0,
      cameraOutput: isSet(object.cameraOutput) ? CameraSensorOutput.fromJSON(object.cameraOutput) : undefined,
      distanceOutput: isSet(object.distanceOutput) ? DistanceSensorOutput.fromJSON(object.distanceOutput) : undefined,
      speedOutput: isSet(object.speedOutput) ? SpeedSensorOutput.fromJSON(object.speedOutput) : undefined,
      controllerOutput: isSet(object.controllerOutput) ? ControllerOutput.fromJSON(object.controllerOutput) : undefined,
      imuOutput: isSet(object.imuOutput) ? ImuSensorOutput.fromJSON(object.imuOutput) : undefined,
      batteryOutput: isSet(object.batteryOutput) ? BatterySensorOutput.fromJSON(object.batteryOutput) : undefined,
      rpmOuput: isSet(object.rpmOuput) ? RpmSensorOutput.fromJSON(object.rpmOuput) : undefined,
      luxOutput: isSet(object.luxOutput) ? LuxSensorOutput.fromJSON(object.luxOutput) : undefined,
      laptimeOutput: isSet(object.laptimeOutput) ? LapTimeOutput.fromJSON(object.laptimeOutput) : undefined,
    };
  },

  toJSON(message: SensorOutput): unknown {
    const obj: any = {};
    if (message.sensorId !== 0) {
      obj.sensorId = Math.round(message.sensorId);
    }
    if (message.timestamp !== 0) {
      obj.timestamp = Math.round(message.timestamp);
    }
    if (message.status !== 0) {
      obj.status = Math.round(message.status);
    }
    if (message.cameraOutput !== undefined) {
      obj.cameraOutput = CameraSensorOutput.toJSON(message.cameraOutput);
    }
    if (message.distanceOutput !== undefined) {
      obj.distanceOutput = DistanceSensorOutput.toJSON(message.distanceOutput);
    }
    if (message.speedOutput !== undefined) {
      obj.speedOutput = SpeedSensorOutput.toJSON(message.speedOutput);
    }
    if (message.controllerOutput !== undefined) {
      obj.controllerOutput = ControllerOutput.toJSON(message.controllerOutput);
    }
    if (message.imuOutput !== undefined) {
      obj.imuOutput = ImuSensorOutput.toJSON(message.imuOutput);
    }
    if (message.batteryOutput !== undefined) {
      obj.batteryOutput = BatterySensorOutput.toJSON(message.batteryOutput);
    }
    if (message.rpmOuput !== undefined) {
      obj.rpmOuput = RpmSensorOutput.toJSON(message.rpmOuput);
    }
    if (message.luxOutput !== undefined) {
      obj.luxOutput = LuxSensorOutput.toJSON(message.luxOutput);
    }
    if (message.laptimeOutput !== undefined) {
      obj.laptimeOutput = LapTimeOutput.toJSON(message.laptimeOutput);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SensorOutput>, I>>(base?: I): SensorOutput {
    return SensorOutput.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SensorOutput>, I>>(object: I): SensorOutput {
    const message = createBaseSensorOutput();
    message.sensorId = object.sensorId ?? 0;
    message.timestamp = object.timestamp ?? 0;
    message.status = object.status ?? 0;
    message.cameraOutput = (object.cameraOutput !== undefined && object.cameraOutput !== null)
      ? CameraSensorOutput.fromPartial(object.cameraOutput)
      : undefined;
    message.distanceOutput = (object.distanceOutput !== undefined && object.distanceOutput !== null)
      ? DistanceSensorOutput.fromPartial(object.distanceOutput)
      : undefined;
    message.speedOutput = (object.speedOutput !== undefined && object.speedOutput !== null)
      ? SpeedSensorOutput.fromPartial(object.speedOutput)
      : undefined;
    message.controllerOutput = (object.controllerOutput !== undefined && object.controllerOutput !== null)
      ? ControllerOutput.fromPartial(object.controllerOutput)
      : undefined;
    message.imuOutput = (object.imuOutput !== undefined && object.imuOutput !== null)
      ? ImuSensorOutput.fromPartial(object.imuOutput)
      : undefined;
    message.batteryOutput = (object.batteryOutput !== undefined && object.batteryOutput !== null)
      ? BatterySensorOutput.fromPartial(object.batteryOutput)
      : undefined;
    message.rpmOuput = (object.rpmOuput !== undefined && object.rpmOuput !== null)
      ? RpmSensorOutput.fromPartial(object.rpmOuput)
      : undefined;
    message.luxOutput = (object.luxOutput !== undefined && object.luxOutput !== null)
      ? LuxSensorOutput.fromPartial(object.luxOutput)
      : undefined;
    message.laptimeOutput = (object.laptimeOutput !== undefined && object.laptimeOutput !== null)
      ? LapTimeOutput.fromPartial(object.laptimeOutput)
      : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(int64: { toString(): string }): number {
  const num = globalThis.Number(int64.toString());
  if (num > globalThis.Number.MAX_SAFE_INTEGER) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  if (num < globalThis.Number.MIN_SAFE_INTEGER) {
    throw new globalThis.Error("Value is smaller than Number.MIN_SAFE_INTEGER");
  }
  return num;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
