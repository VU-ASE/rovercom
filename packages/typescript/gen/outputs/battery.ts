// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.4.2
//   protoc               v3.12.4
// source: outputs/battery.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "protobuf_msgs";

export interface BatterySensorOutput {
  currentOutputVoltage: number;
  warnVoltage: number;
  killVoltage: number;
}

function createBaseBatterySensorOutput(): BatterySensorOutput {
  return { currentOutputVoltage: 0, warnVoltage: 0, killVoltage: 0 };
}

export const BatterySensorOutput: MessageFns<BatterySensorOutput> = {
  encode(message: BatterySensorOutput, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.currentOutputVoltage !== 0) {
      writer.uint32(13).float(message.currentOutputVoltage);
    }
    if (message.warnVoltage !== 0) {
      writer.uint32(21).float(message.warnVoltage);
    }
    if (message.killVoltage !== 0) {
      writer.uint32(29).float(message.killVoltage);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): BatterySensorOutput {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBatterySensorOutput();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 13) {
            break;
          }

          message.currentOutputVoltage = reader.float();
          continue;
        }
        case 2: {
          if (tag !== 21) {
            break;
          }

          message.warnVoltage = reader.float();
          continue;
        }
        case 3: {
          if (tag !== 29) {
            break;
          }

          message.killVoltage = reader.float();
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

  fromJSON(object: any): BatterySensorOutput {
    return {
      currentOutputVoltage: isSet(object.currentOutputVoltage) ? globalThis.Number(object.currentOutputVoltage) : 0,
      warnVoltage: isSet(object.warnVoltage) ? globalThis.Number(object.warnVoltage) : 0,
      killVoltage: isSet(object.killVoltage) ? globalThis.Number(object.killVoltage) : 0,
    };
  },

  toJSON(message: BatterySensorOutput): unknown {
    const obj: any = {};
    if (message.currentOutputVoltage !== 0) {
      obj.currentOutputVoltage = message.currentOutputVoltage;
    }
    if (message.warnVoltage !== 0) {
      obj.warnVoltage = message.warnVoltage;
    }
    if (message.killVoltage !== 0) {
      obj.killVoltage = message.killVoltage;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<BatterySensorOutput>, I>>(base?: I): BatterySensorOutput {
    return BatterySensorOutput.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<BatterySensorOutput>, I>>(object: I): BatterySensorOutput {
    const message = createBaseBatterySensorOutput();
    message.currentOutputVoltage = object.currentOutputVoltage ?? 0;
    message.warnVoltage = object.warnVoltage ?? 0;
    message.killVoltage = object.killVoltage ?? 0;
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
