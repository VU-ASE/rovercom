// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.4.2
//   protoc               v3.12.4
// source: outputs/laptime.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "protobuf_msgs";

export interface LapTimeOutput {
  /** Time of complete lap in ms */
  lapTime: number;
  /** Start of the lap */
  lapStartTime: number;
}

function createBaseLapTimeOutput(): LapTimeOutput {
  return { lapTime: 0, lapStartTime: 0 };
}

export const LapTimeOutput: MessageFns<LapTimeOutput> = {
  encode(message: LapTimeOutput, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.lapTime !== 0) {
      writer.uint32(8).uint64(message.lapTime);
    }
    if (message.lapStartTime !== 0) {
      writer.uint32(16).uint64(message.lapStartTime);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): LapTimeOutput {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLapTimeOutput();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.lapTime = longToNumber(reader.uint64());
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.lapStartTime = longToNumber(reader.uint64());
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

  fromJSON(object: any): LapTimeOutput {
    return {
      lapTime: isSet(object.lapTime) ? globalThis.Number(object.lapTime) : 0,
      lapStartTime: isSet(object.lapStartTime) ? globalThis.Number(object.lapStartTime) : 0,
    };
  },

  toJSON(message: LapTimeOutput): unknown {
    const obj: any = {};
    if (message.lapTime !== 0) {
      obj.lapTime = Math.round(message.lapTime);
    }
    if (message.lapStartTime !== 0) {
      obj.lapStartTime = Math.round(message.lapStartTime);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<LapTimeOutput>, I>>(base?: I): LapTimeOutput {
    return LapTimeOutput.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<LapTimeOutput>, I>>(object: I): LapTimeOutput {
    const message = createBaseLapTimeOutput();
    message.lapTime = object.lapTime ?? 0;
    message.lapStartTime = object.lapStartTime ?? 0;
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