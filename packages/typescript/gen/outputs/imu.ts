// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.5
//   protoc               v3.12.4
// source: outputs/imu.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "protobuf_msgs";

export interface ImuSensorOutput {
  temperature: number;
  magnetometer: ImuSensorOutput_Vector | undefined;
  gyroscope: ImuSensorOutput_Vector | undefined;
  euler: ImuSensorOutput_Vector | undefined;
  accelerometer: ImuSensorOutput_Vector | undefined;
  linearAccelerometer: ImuSensorOutput_Vector | undefined;
  velocity: ImuSensorOutput_Vector | undefined;
  speed: number;
}

export interface ImuSensorOutput_Vector {
  x: number;
  y: number;
  z: number;
}

function createBaseImuSensorOutput(): ImuSensorOutput {
  return {
    temperature: 0,
    magnetometer: undefined,
    gyroscope: undefined,
    euler: undefined,
    accelerometer: undefined,
    linearAccelerometer: undefined,
    velocity: undefined,
    speed: 0,
  };
}

export const ImuSensorOutput: MessageFns<ImuSensorOutput> = {
  encode(message: ImuSensorOutput, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.temperature !== 0) {
      writer.uint32(8).int32(message.temperature);
    }
    if (message.magnetometer !== undefined) {
      ImuSensorOutput_Vector.encode(message.magnetometer, writer.uint32(18).fork()).join();
    }
    if (message.gyroscope !== undefined) {
      ImuSensorOutput_Vector.encode(message.gyroscope, writer.uint32(26).fork()).join();
    }
    if (message.euler !== undefined) {
      ImuSensorOutput_Vector.encode(message.euler, writer.uint32(34).fork()).join();
    }
    if (message.accelerometer !== undefined) {
      ImuSensorOutput_Vector.encode(message.accelerometer, writer.uint32(42).fork()).join();
    }
    if (message.linearAccelerometer !== undefined) {
      ImuSensorOutput_Vector.encode(message.linearAccelerometer, writer.uint32(50).fork()).join();
    }
    if (message.velocity !== undefined) {
      ImuSensorOutput_Vector.encode(message.velocity, writer.uint32(58).fork()).join();
    }
    if (message.speed !== 0) {
      writer.uint32(69).float(message.speed);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ImuSensorOutput {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    const end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseImuSensorOutput();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.temperature = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.magnetometer = ImuSensorOutput_Vector.decode(reader, reader.uint32());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.gyroscope = ImuSensorOutput_Vector.decode(reader, reader.uint32());
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.euler = ImuSensorOutput_Vector.decode(reader, reader.uint32());
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.accelerometer = ImuSensorOutput_Vector.decode(reader, reader.uint32());
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.linearAccelerometer = ImuSensorOutput_Vector.decode(reader, reader.uint32());
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.velocity = ImuSensorOutput_Vector.decode(reader, reader.uint32());
          continue;
        }
        case 8: {
          if (tag !== 69) {
            break;
          }

          message.speed = reader.float();
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

  fromJSON(object: any): ImuSensorOutput {
    return {
      temperature: isSet(object.temperature) ? globalThis.Number(object.temperature) : 0,
      magnetometer: isSet(object.magnetometer) ? ImuSensorOutput_Vector.fromJSON(object.magnetometer) : undefined,
      gyroscope: isSet(object.gyroscope) ? ImuSensorOutput_Vector.fromJSON(object.gyroscope) : undefined,
      euler: isSet(object.euler) ? ImuSensorOutput_Vector.fromJSON(object.euler) : undefined,
      accelerometer: isSet(object.accelerometer) ? ImuSensorOutput_Vector.fromJSON(object.accelerometer) : undefined,
      linearAccelerometer: isSet(object.linearAccelerometer)
        ? ImuSensorOutput_Vector.fromJSON(object.linearAccelerometer)
        : undefined,
      velocity: isSet(object.velocity) ? ImuSensorOutput_Vector.fromJSON(object.velocity) : undefined,
      speed: isSet(object.speed) ? globalThis.Number(object.speed) : 0,
    };
  },

  toJSON(message: ImuSensorOutput): unknown {
    const obj: any = {};
    if (message.temperature !== 0) {
      obj.temperature = Math.round(message.temperature);
    }
    if (message.magnetometer !== undefined) {
      obj.magnetometer = ImuSensorOutput_Vector.toJSON(message.magnetometer);
    }
    if (message.gyroscope !== undefined) {
      obj.gyroscope = ImuSensorOutput_Vector.toJSON(message.gyroscope);
    }
    if (message.euler !== undefined) {
      obj.euler = ImuSensorOutput_Vector.toJSON(message.euler);
    }
    if (message.accelerometer !== undefined) {
      obj.accelerometer = ImuSensorOutput_Vector.toJSON(message.accelerometer);
    }
    if (message.linearAccelerometer !== undefined) {
      obj.linearAccelerometer = ImuSensorOutput_Vector.toJSON(message.linearAccelerometer);
    }
    if (message.velocity !== undefined) {
      obj.velocity = ImuSensorOutput_Vector.toJSON(message.velocity);
    }
    if (message.speed !== 0) {
      obj.speed = message.speed;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ImuSensorOutput>, I>>(base?: I): ImuSensorOutput {
    return ImuSensorOutput.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ImuSensorOutput>, I>>(object: I): ImuSensorOutput {
    const message = createBaseImuSensorOutput();
    message.temperature = object.temperature ?? 0;
    message.magnetometer = (object.magnetometer !== undefined && object.magnetometer !== null)
      ? ImuSensorOutput_Vector.fromPartial(object.magnetometer)
      : undefined;
    message.gyroscope = (object.gyroscope !== undefined && object.gyroscope !== null)
      ? ImuSensorOutput_Vector.fromPartial(object.gyroscope)
      : undefined;
    message.euler = (object.euler !== undefined && object.euler !== null)
      ? ImuSensorOutput_Vector.fromPartial(object.euler)
      : undefined;
    message.accelerometer = (object.accelerometer !== undefined && object.accelerometer !== null)
      ? ImuSensorOutput_Vector.fromPartial(object.accelerometer)
      : undefined;
    message.linearAccelerometer = (object.linearAccelerometer !== undefined && object.linearAccelerometer !== null)
      ? ImuSensorOutput_Vector.fromPartial(object.linearAccelerometer)
      : undefined;
    message.velocity = (object.velocity !== undefined && object.velocity !== null)
      ? ImuSensorOutput_Vector.fromPartial(object.velocity)
      : undefined;
    message.speed = object.speed ?? 0;
    return message;
  },
};

function createBaseImuSensorOutput_Vector(): ImuSensorOutput_Vector {
  return { x: 0, y: 0, z: 0 };
}

export const ImuSensorOutput_Vector: MessageFns<ImuSensorOutput_Vector> = {
  encode(message: ImuSensorOutput_Vector, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.x !== 0) {
      writer.uint32(13).float(message.x);
    }
    if (message.y !== 0) {
      writer.uint32(21).float(message.y);
    }
    if (message.z !== 0) {
      writer.uint32(29).float(message.z);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ImuSensorOutput_Vector {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    const end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseImuSensorOutput_Vector();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 13) {
            break;
          }

          message.x = reader.float();
          continue;
        }
        case 2: {
          if (tag !== 21) {
            break;
          }

          message.y = reader.float();
          continue;
        }
        case 3: {
          if (tag !== 29) {
            break;
          }

          message.z = reader.float();
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

  fromJSON(object: any): ImuSensorOutput_Vector {
    return {
      x: isSet(object.x) ? globalThis.Number(object.x) : 0,
      y: isSet(object.y) ? globalThis.Number(object.y) : 0,
      z: isSet(object.z) ? globalThis.Number(object.z) : 0,
    };
  },

  toJSON(message: ImuSensorOutput_Vector): unknown {
    const obj: any = {};
    if (message.x !== 0) {
      obj.x = message.x;
    }
    if (message.y !== 0) {
      obj.y = message.y;
    }
    if (message.z !== 0) {
      obj.z = message.z;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ImuSensorOutput_Vector>, I>>(base?: I): ImuSensorOutput_Vector {
    return ImuSensorOutput_Vector.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ImuSensorOutput_Vector>, I>>(object: I): ImuSensorOutput_Vector {
    const message = createBaseImuSensorOutput_Vector();
    message.x = object.x ?? 0;
    message.y = object.y ?? 0;
    message.z = object.z ?? 0;
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
