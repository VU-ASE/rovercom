// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v3.12.4
// source: segmentation/segmentation.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "protobuf_msgs";

/** Control messages exchanged by client(s), the server and the car to send data in multiple segments */
export interface Segment {
  /** unique and increasing id of each the packet that this segment is part of */
  packetId: number;
  /** unique and increasing id of the segment within the packet */
  segmentId: number;
  /** total number of segments in the packet */
  totalSegments: number;
  /** data of the segment */
  data: Uint8Array;
}

function createBaseSegment(): Segment {
  return { packetId: 0, segmentId: 0, totalSegments: 0, data: new Uint8Array(0) };
}

export const Segment: MessageFns<Segment> = {
  encode(message: Segment, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.packetId !== 0) {
      writer.uint32(8).int64(message.packetId);
    }
    if (message.segmentId !== 0) {
      writer.uint32(16).int64(message.segmentId);
    }
    if (message.totalSegments !== 0) {
      writer.uint32(24).int64(message.totalSegments);
    }
    if (message.data.length !== 0) {
      writer.uint32(34).bytes(message.data);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Segment {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSegment();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.packetId = longToNumber(reader.int64());
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.segmentId = longToNumber(reader.int64());
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.totalSegments = longToNumber(reader.int64());
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.data = reader.bytes();
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

  fromJSON(object: any): Segment {
    return {
      packetId: isSet(object.packetId) ? globalThis.Number(object.packetId) : 0,
      segmentId: isSet(object.segmentId) ? globalThis.Number(object.segmentId) : 0,
      totalSegments: isSet(object.totalSegments) ? globalThis.Number(object.totalSegments) : 0,
      data: isSet(object.data) ? bytesFromBase64(object.data) : new Uint8Array(0),
    };
  },

  toJSON(message: Segment): unknown {
    const obj: any = {};
    if (message.packetId !== 0) {
      obj.packetId = Math.round(message.packetId);
    }
    if (message.segmentId !== 0) {
      obj.segmentId = Math.round(message.segmentId);
    }
    if (message.totalSegments !== 0) {
      obj.totalSegments = Math.round(message.totalSegments);
    }
    if (message.data.length !== 0) {
      obj.data = base64FromBytes(message.data);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Segment>, I>>(base?: I): Segment {
    return Segment.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Segment>, I>>(object: I): Segment {
    const message = createBaseSegment();
    message.packetId = object.packetId ?? 0;
    message.segmentId = object.segmentId ?? 0;
    message.totalSegments = object.totalSegments ?? 0;
    message.data = object.data ?? new Uint8Array(0);
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
