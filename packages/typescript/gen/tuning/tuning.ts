// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.5.0
//   protoc               v3.12.4
// source: tuning/tuning.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "protobuf_msgs";

export interface TuningState {
  /** increased when the state changes, used to prevent unnecessary updates (so *not* in milliseconds since epoch) */
  timestamp: number;
  dynamicParameters: TuningState_Parameter[];
}

export interface TuningState_Parameter {
  float?: TuningState_Parameter_FloatParameter | undefined;
  int?: TuningState_Parameter_IntParameter | undefined;
  string?: TuningState_Parameter_StringParameter | undefined;
}

/**
 * note: it may seem weird to not extract the key from the oneof, but this is so that the parser can easily determine the type of the parameter
 * extracting it to a separate field on the same level as oneof would make it ambiguous
 */
export interface TuningState_Parameter_FloatParameter {
  key: string;
  value: number;
}

export interface TuningState_Parameter_IntParameter {
  key: string;
  value: number;
}

export interface TuningState_Parameter_StringParameter {
  key: string;
  value: string;
}

function createBaseTuningState(): TuningState {
  return { timestamp: 0, dynamicParameters: [] };
}

export const TuningState: MessageFns<TuningState> = {
  encode(message: TuningState, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.timestamp !== 0) {
      writer.uint32(8).uint64(message.timestamp);
    }
    for (const v of message.dynamicParameters) {
      TuningState_Parameter.encode(v!, writer.uint32(18).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): TuningState {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTuningState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.timestamp = longToNumber(reader.uint64());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.dynamicParameters.push(TuningState_Parameter.decode(reader, reader.uint32()));
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

  fromJSON(object: any): TuningState {
    return {
      timestamp: isSet(object.timestamp) ? globalThis.Number(object.timestamp) : 0,
      dynamicParameters: globalThis.Array.isArray(object?.dynamicParameters)
        ? object.dynamicParameters.map((e: any) => TuningState_Parameter.fromJSON(e))
        : [],
    };
  },

  toJSON(message: TuningState): unknown {
    const obj: any = {};
    if (message.timestamp !== 0) {
      obj.timestamp = Math.round(message.timestamp);
    }
    if (message.dynamicParameters?.length) {
      obj.dynamicParameters = message.dynamicParameters.map((e) => TuningState_Parameter.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TuningState>, I>>(base?: I): TuningState {
    return TuningState.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TuningState>, I>>(object: I): TuningState {
    const message = createBaseTuningState();
    message.timestamp = object.timestamp ?? 0;
    message.dynamicParameters = object.dynamicParameters?.map((e) => TuningState_Parameter.fromPartial(e)) || [];
    return message;
  },
};

function createBaseTuningState_Parameter(): TuningState_Parameter {
  return { float: undefined, int: undefined, string: undefined };
}

export const TuningState_Parameter: MessageFns<TuningState_Parameter> = {
  encode(message: TuningState_Parameter, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.float !== undefined) {
      TuningState_Parameter_FloatParameter.encode(message.float, writer.uint32(10).fork()).join();
    }
    if (message.int !== undefined) {
      TuningState_Parameter_IntParameter.encode(message.int, writer.uint32(18).fork()).join();
    }
    if (message.string !== undefined) {
      TuningState_Parameter_StringParameter.encode(message.string, writer.uint32(26).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): TuningState_Parameter {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTuningState_Parameter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.float = TuningState_Parameter_FloatParameter.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.int = TuningState_Parameter_IntParameter.decode(reader, reader.uint32());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.string = TuningState_Parameter_StringParameter.decode(reader, reader.uint32());
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

  fromJSON(object: any): TuningState_Parameter {
    return {
      float: isSet(object.float) ? TuningState_Parameter_FloatParameter.fromJSON(object.float) : undefined,
      int: isSet(object.int) ? TuningState_Parameter_IntParameter.fromJSON(object.int) : undefined,
      string: isSet(object.string) ? TuningState_Parameter_StringParameter.fromJSON(object.string) : undefined,
    };
  },

  toJSON(message: TuningState_Parameter): unknown {
    const obj: any = {};
    if (message.float !== undefined) {
      obj.float = TuningState_Parameter_FloatParameter.toJSON(message.float);
    }
    if (message.int !== undefined) {
      obj.int = TuningState_Parameter_IntParameter.toJSON(message.int);
    }
    if (message.string !== undefined) {
      obj.string = TuningState_Parameter_StringParameter.toJSON(message.string);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TuningState_Parameter>, I>>(base?: I): TuningState_Parameter {
    return TuningState_Parameter.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TuningState_Parameter>, I>>(object: I): TuningState_Parameter {
    const message = createBaseTuningState_Parameter();
    message.float = (object.float !== undefined && object.float !== null)
      ? TuningState_Parameter_FloatParameter.fromPartial(object.float)
      : undefined;
    message.int = (object.int !== undefined && object.int !== null)
      ? TuningState_Parameter_IntParameter.fromPartial(object.int)
      : undefined;
    message.string = (object.string !== undefined && object.string !== null)
      ? TuningState_Parameter_StringParameter.fromPartial(object.string)
      : undefined;
    return message;
  },
};

function createBaseTuningState_Parameter_FloatParameter(): TuningState_Parameter_FloatParameter {
  return { key: "", value: 0 };
}

export const TuningState_Parameter_FloatParameter: MessageFns<TuningState_Parameter_FloatParameter> = {
  encode(message: TuningState_Parameter_FloatParameter, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== 0) {
      writer.uint32(21).float(message.value);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): TuningState_Parameter_FloatParameter {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTuningState_Parameter_FloatParameter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 21) {
            break;
          }

          message.value = reader.float();
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

  fromJSON(object: any): TuningState_Parameter_FloatParameter {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.Number(object.value) : 0,
    };
  },

  toJSON(message: TuningState_Parameter_FloatParameter): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== 0) {
      obj.value = message.value;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TuningState_Parameter_FloatParameter>, I>>(
    base?: I,
  ): TuningState_Parameter_FloatParameter {
    return TuningState_Parameter_FloatParameter.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TuningState_Parameter_FloatParameter>, I>>(
    object: I,
  ): TuningState_Parameter_FloatParameter {
    const message = createBaseTuningState_Parameter_FloatParameter();
    message.key = object.key ?? "";
    message.value = object.value ?? 0;
    return message;
  },
};

function createBaseTuningState_Parameter_IntParameter(): TuningState_Parameter_IntParameter {
  return { key: "", value: 0 };
}

export const TuningState_Parameter_IntParameter: MessageFns<TuningState_Parameter_IntParameter> = {
  encode(message: TuningState_Parameter_IntParameter, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== 0) {
      writer.uint32(16).int64(message.value);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): TuningState_Parameter_IntParameter {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTuningState_Parameter_IntParameter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.value = longToNumber(reader.int64());
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

  fromJSON(object: any): TuningState_Parameter_IntParameter {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.Number(object.value) : 0,
    };
  },

  toJSON(message: TuningState_Parameter_IntParameter): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== 0) {
      obj.value = Math.round(message.value);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TuningState_Parameter_IntParameter>, I>>(
    base?: I,
  ): TuningState_Parameter_IntParameter {
    return TuningState_Parameter_IntParameter.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TuningState_Parameter_IntParameter>, I>>(
    object: I,
  ): TuningState_Parameter_IntParameter {
    const message = createBaseTuningState_Parameter_IntParameter();
    message.key = object.key ?? "";
    message.value = object.value ?? 0;
    return message;
  },
};

function createBaseTuningState_Parameter_StringParameter(): TuningState_Parameter_StringParameter {
  return { key: "", value: "" };
}

export const TuningState_Parameter_StringParameter: MessageFns<TuningState_Parameter_StringParameter> = {
  encode(message: TuningState_Parameter_StringParameter, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): TuningState_Parameter_StringParameter {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTuningState_Parameter_StringParameter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.value = reader.string();
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

  fromJSON(object: any): TuningState_Parameter_StringParameter {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.String(object.value) : "",
    };
  },

  toJSON(message: TuningState_Parameter_StringParameter): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== "") {
      obj.value = message.value;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TuningState_Parameter_StringParameter>, I>>(
    base?: I,
  ): TuningState_Parameter_StringParameter {
    return TuningState_Parameter_StringParameter.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TuningState_Parameter_StringParameter>, I>>(
    object: I,
  ): TuningState_Parameter_StringParameter {
    const message = createBaseTuningState_Parameter_StringParameter();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
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
