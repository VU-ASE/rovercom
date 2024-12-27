// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v3.12.4
// source: outputs/controller.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "protobuf_msgs";

export interface ControllerOutput {
  /** Steering angle (-1.0 to 1.0 <-> left - right) */
  steeringAngle: number;
  /** Throttle (-1.0 to 1.0 <-> full reverse - full forward) */
  leftThrottle: number;
  rightThrottle: number;
  /** Onboard lights (0.0 to 1.0 <-> off - on) */
  frontLights: boolean;
  /** Fan speed (0.0 to 1.0 <-> off - full speed) */
  fanSpeed: number;
  /** Useful for debugging */
  rawError: number;
  /** the error value after quadratic scaling, but before PID */
  scaledError: number;
}

function createBaseControllerOutput(): ControllerOutput {
  return {
    steeringAngle: 0,
    leftThrottle: 0,
    rightThrottle: 0,
    frontLights: false,
    fanSpeed: 0,
    rawError: 0,
    scaledError: 0,
  };
}

export const ControllerOutput: MessageFns<ControllerOutput> = {
  encode(message: ControllerOutput, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.steeringAngle !== 0) {
      writer.uint32(21).float(message.steeringAngle);
    }
    if (message.leftThrottle !== 0) {
      writer.uint32(29).float(message.leftThrottle);
    }
    if (message.rightThrottle !== 0) {
      writer.uint32(37).float(message.rightThrottle);
    }
    if (message.frontLights !== false) {
      writer.uint32(40).bool(message.frontLights);
    }
    if (message.fanSpeed !== 0) {
      writer.uint32(53).float(message.fanSpeed);
    }
    if (message.rawError !== 0) {
      writer.uint32(61).float(message.rawError);
    }
    if (message.scaledError !== 0) {
      writer.uint32(69).float(message.scaledError);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ControllerOutput {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseControllerOutput();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2: {
          if (tag !== 21) {
            break;
          }

          message.steeringAngle = reader.float();
          continue;
        }
        case 3: {
          if (tag !== 29) {
            break;
          }

          message.leftThrottle = reader.float();
          continue;
        }
        case 4: {
          if (tag !== 37) {
            break;
          }

          message.rightThrottle = reader.float();
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.frontLights = reader.bool();
          continue;
        }
        case 6: {
          if (tag !== 53) {
            break;
          }

          message.fanSpeed = reader.float();
          continue;
        }
        case 7: {
          if (tag !== 61) {
            break;
          }

          message.rawError = reader.float();
          continue;
        }
        case 8: {
          if (tag !== 69) {
            break;
          }

          message.scaledError = reader.float();
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

  fromJSON(object: any): ControllerOutput {
    return {
      steeringAngle: isSet(object.steeringAngle) ? globalThis.Number(object.steeringAngle) : 0,
      leftThrottle: isSet(object.leftThrottle) ? globalThis.Number(object.leftThrottle) : 0,
      rightThrottle: isSet(object.rightThrottle) ? globalThis.Number(object.rightThrottle) : 0,
      frontLights: isSet(object.frontLights) ? globalThis.Boolean(object.frontLights) : false,
      fanSpeed: isSet(object.fanSpeed) ? globalThis.Number(object.fanSpeed) : 0,
      rawError: isSet(object.rawError) ? globalThis.Number(object.rawError) : 0,
      scaledError: isSet(object.scaledError) ? globalThis.Number(object.scaledError) : 0,
    };
  },

  toJSON(message: ControllerOutput): unknown {
    const obj: any = {};
    if (message.steeringAngle !== 0) {
      obj.steeringAngle = message.steeringAngle;
    }
    if (message.leftThrottle !== 0) {
      obj.leftThrottle = message.leftThrottle;
    }
    if (message.rightThrottle !== 0) {
      obj.rightThrottle = message.rightThrottle;
    }
    if (message.frontLights !== false) {
      obj.frontLights = message.frontLights;
    }
    if (message.fanSpeed !== 0) {
      obj.fanSpeed = message.fanSpeed;
    }
    if (message.rawError !== 0) {
      obj.rawError = message.rawError;
    }
    if (message.scaledError !== 0) {
      obj.scaledError = message.scaledError;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ControllerOutput>, I>>(base?: I): ControllerOutput {
    return ControllerOutput.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ControllerOutput>, I>>(object: I): ControllerOutput {
    const message = createBaseControllerOutput();
    message.steeringAngle = object.steeringAngle ?? 0;
    message.leftThrottle = object.leftThrottle ?? 0;
    message.rightThrottle = object.rightThrottle ?? 0;
    message.frontLights = object.frontLights ?? false;
    message.fanSpeed = object.fanSpeed ?? 0;
    message.rawError = object.rawError ?? 0;
    message.scaledError = object.scaledError ?? 0;
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
