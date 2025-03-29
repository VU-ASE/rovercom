// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v3.12.4
// source: outputs/rpm.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "protobuf_msgs";

export interface RpmSensorOutput {
  leftRpm: number;
  leftAngle: number;
  rightRpm: number;
  rightAngle: number;
}

function createBaseRpmSensorOutput(): RpmSensorOutput {
  return { leftRpm: 0, leftAngle: 0, rightRpm: 0, rightAngle: 0 };
}

export const RpmSensorOutput: MessageFns<RpmSensorOutput> = {
  encode(message: RpmSensorOutput, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.leftRpm !== 0) {
      writer.uint32(13).float(message.leftRpm);
    }
    if (message.leftAngle !== 0) {
      writer.uint32(21).float(message.leftAngle);
    }
    if (message.rightRpm !== 0) {
      writer.uint32(29).float(message.rightRpm);
    }
    if (message.rightAngle !== 0) {
      writer.uint32(37).float(message.rightAngle);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): RpmSensorOutput {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRpmSensorOutput();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 13) {
            break;
          }

          message.leftRpm = reader.float();
          continue;
        }
        case 2: {
          if (tag !== 21) {
            break;
          }

          message.leftAngle = reader.float();
          continue;
        }
        case 3: {
          if (tag !== 29) {
            break;
          }

          message.rightRpm = reader.float();
          continue;
        }
        case 4: {
          if (tag !== 37) {
            break;
          }

          message.rightAngle = reader.float();
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

  fromJSON(object: any): RpmSensorOutput {
    return {
      leftRpm: isSet(object.leftRpm) ? globalThis.Number(object.leftRpm) : 0,
      leftAngle: isSet(object.leftAngle) ? globalThis.Number(object.leftAngle) : 0,
      rightRpm: isSet(object.rightRpm) ? globalThis.Number(object.rightRpm) : 0,
      rightAngle: isSet(object.rightAngle) ? globalThis.Number(object.rightAngle) : 0,
    };
  },

  toJSON(message: RpmSensorOutput): unknown {
    const obj: any = {};
    if (message.leftRpm !== 0) {
      obj.leftRpm = message.leftRpm;
    }
    if (message.leftAngle !== 0) {
      obj.leftAngle = message.leftAngle;
    }
    if (message.rightRpm !== 0) {
      obj.rightRpm = message.rightRpm;
    }
    if (message.rightAngle !== 0) {
      obj.rightAngle = message.rightAngle;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RpmSensorOutput>, I>>(base?: I): RpmSensorOutput {
    return RpmSensorOutput.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RpmSensorOutput>, I>>(object: I): RpmSensorOutput {
    const message = createBaseRpmSensorOutput();
    message.leftRpm = object.leftRpm ?? 0;
    message.leftAngle = object.leftAngle ?? 0;
    message.rightRpm = object.rightRpm ?? 0;
    message.rightAngle = object.rightAngle ?? 0;
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
