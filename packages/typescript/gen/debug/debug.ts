// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.5
//   protoc               v3.12.4
// source: debug/debug.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "protobuf_msgs";

/** Used to identify a service within the pipeline */
export interface ServiceIdentifier {
  name: string;
  pid: number;
}

/** An endpoint that is made available by a service */
export interface ServiceEndpoint {
  /** the identifier to select this endpoint */
  name: string;
  /** the zmq address to connect to */
  address: string;
}

/**
 * When the transceivers picks up a SensorOutput from a service, it will wrap it in a ServiceMessage message,
 * so that the receiver can determine from which process the message originated
 */
export interface DebugOutput {
  /** The service that sent the message */
  service:
    | ServiceIdentifier
    | undefined;
  /** The endpoint in this service that sent the message */
  endpoint:
    | ServiceEndpoint
    | undefined;
  /** Time in milliseconds since epoch when the message was sent */
  sentAt: number;
  message: Uint8Array;
}

function createBaseServiceIdentifier(): ServiceIdentifier {
  return { name: "", pid: 0 };
}

export const ServiceIdentifier: MessageFns<ServiceIdentifier> = {
  encode(message: ServiceIdentifier, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ServiceIdentifier {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    const end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseServiceIdentifier();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.pid = reader.int32();
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

  fromJSON(object: any): ServiceIdentifier {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      pid: isSet(object.pid) ? globalThis.Number(object.pid) : 0,
    };
  },

  toJSON(message: ServiceIdentifier): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.pid !== 0) {
      obj.pid = Math.round(message.pid);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ServiceIdentifier>, I>>(base?: I): ServiceIdentifier {
    return ServiceIdentifier.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ServiceIdentifier>, I>>(object: I): ServiceIdentifier {
    const message = createBaseServiceIdentifier();
    message.name = object.name ?? "";
    message.pid = object.pid ?? 0;
    return message;
  },
};

function createBaseServiceEndpoint(): ServiceEndpoint {
  return { name: "", address: "" };
}

export const ServiceEndpoint: MessageFns<ServiceEndpoint> = {
  encode(message: ServiceEndpoint, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ServiceEndpoint {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    const end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseServiceEndpoint();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.address = reader.string();
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

  fromJSON(object: any): ServiceEndpoint {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      address: isSet(object.address) ? globalThis.String(object.address) : "",
    };
  },

  toJSON(message: ServiceEndpoint): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.address !== "") {
      obj.address = message.address;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ServiceEndpoint>, I>>(base?: I): ServiceEndpoint {
    return ServiceEndpoint.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ServiceEndpoint>, I>>(object: I): ServiceEndpoint {
    const message = createBaseServiceEndpoint();
    message.name = object.name ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseDebugOutput(): DebugOutput {
  return { service: undefined, endpoint: undefined, sentAt: 0, message: new Uint8Array(0) };
}

export const DebugOutput: MessageFns<DebugOutput> = {
  encode(message: DebugOutput, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.service !== undefined) {
      ServiceIdentifier.encode(message.service, writer.uint32(10).fork()).join();
    }
    if (message.endpoint !== undefined) {
      ServiceEndpoint.encode(message.endpoint, writer.uint32(18).fork()).join();
    }
    if (message.sentAt !== 0) {
      writer.uint32(32).int64(message.sentAt);
    }
    if (message.message.length !== 0) {
      writer.uint32(26).bytes(message.message);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DebugOutput {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    const end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDebugOutput();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.service = ServiceIdentifier.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.endpoint = ServiceEndpoint.decode(reader, reader.uint32());
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.sentAt = longToNumber(reader.int64());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.message = reader.bytes();
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

  fromJSON(object: any): DebugOutput {
    return {
      service: isSet(object.service) ? ServiceIdentifier.fromJSON(object.service) : undefined,
      endpoint: isSet(object.endpoint) ? ServiceEndpoint.fromJSON(object.endpoint) : undefined,
      sentAt: isSet(object.sentAt) ? globalThis.Number(object.sentAt) : 0,
      message: isSet(object.message) ? bytesFromBase64(object.message) : new Uint8Array(0),
    };
  },

  toJSON(message: DebugOutput): unknown {
    const obj: any = {};
    if (message.service !== undefined) {
      obj.service = ServiceIdentifier.toJSON(message.service);
    }
    if (message.endpoint !== undefined) {
      obj.endpoint = ServiceEndpoint.toJSON(message.endpoint);
    }
    if (message.sentAt !== 0) {
      obj.sentAt = Math.round(message.sentAt);
    }
    if (message.message.length !== 0) {
      obj.message = base64FromBytes(message.message);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DebugOutput>, I>>(base?: I): DebugOutput {
    return DebugOutput.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DebugOutput>, I>>(object: I): DebugOutput {
    const message = createBaseDebugOutput();
    message.service = (object.service !== undefined && object.service !== null)
      ? ServiceIdentifier.fromPartial(object.service)
      : undefined;
    message.endpoint = (object.endpoint !== undefined && object.endpoint !== null)
      ? ServiceEndpoint.fromPartial(object.endpoint)
      : undefined;
    message.sentAt = object.sentAt ?? 0;
    message.message = object.message ?? new Uint8Array(0);
    return message;
  },
};

function bytesFromBase64(b64: string): Uint8Array {
  if ((globalThis as any).Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if ((globalThis as any).Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(globalThis.String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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
