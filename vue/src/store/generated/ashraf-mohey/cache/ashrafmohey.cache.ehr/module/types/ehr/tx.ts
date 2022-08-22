/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Ehr } from "../ehr/ehr";

export const protobufPackage = "ashrafmohey.cache.ehr";

export interface MsgAddEhr {
  creator: string;
  patientId: number;
  patientPrivateKey: string;
  dataHash: string;
}

export interface MsgAddEhrResponse {
  ehr: Ehr | undefined;
}

export interface MsgSendIbcTransferEhrs {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  patientId: number;
}

export interface MsgSendIbcTransferEhrsResponse {}

const baseMsgAddEhr: object = {
  creator: "",
  patientId: 0,
  patientPrivateKey: "",
  dataHash: "",
};

export const MsgAddEhr = {
  encode(message: MsgAddEhr, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.patientId !== 0) {
      writer.uint32(16).uint64(message.patientId);
    }
    if (message.patientPrivateKey !== "") {
      writer.uint32(26).string(message.patientPrivateKey);
    }
    if (message.dataHash !== "") {
      writer.uint32(34).string(message.dataHash);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddEhr {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddEhr } as MsgAddEhr;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.patientId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.patientPrivateKey = reader.string();
          break;
        case 4:
          message.dataHash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddEhr {
    const message = { ...baseMsgAddEhr } as MsgAddEhr;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.patientId !== undefined && object.patientId !== null) {
      message.patientId = Number(object.patientId);
    } else {
      message.patientId = 0;
    }
    if (
      object.patientPrivateKey !== undefined &&
      object.patientPrivateKey !== null
    ) {
      message.patientPrivateKey = String(object.patientPrivateKey);
    } else {
      message.patientPrivateKey = "";
    }
    if (object.dataHash !== undefined && object.dataHash !== null) {
      message.dataHash = String(object.dataHash);
    } else {
      message.dataHash = "";
    }
    return message;
  },

  toJSON(message: MsgAddEhr): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.patientId !== undefined && (obj.patientId = message.patientId);
    message.patientPrivateKey !== undefined &&
      (obj.patientPrivateKey = message.patientPrivateKey);
    message.dataHash !== undefined && (obj.dataHash = message.dataHash);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddEhr>): MsgAddEhr {
    const message = { ...baseMsgAddEhr } as MsgAddEhr;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.patientId !== undefined && object.patientId !== null) {
      message.patientId = object.patientId;
    } else {
      message.patientId = 0;
    }
    if (
      object.patientPrivateKey !== undefined &&
      object.patientPrivateKey !== null
    ) {
      message.patientPrivateKey = object.patientPrivateKey;
    } else {
      message.patientPrivateKey = "";
    }
    if (object.dataHash !== undefined && object.dataHash !== null) {
      message.dataHash = object.dataHash;
    } else {
      message.dataHash = "";
    }
    return message;
  },
};

const baseMsgAddEhrResponse: object = {};

export const MsgAddEhrResponse = {
  encode(message: MsgAddEhrResponse, writer: Writer = Writer.create()): Writer {
    if (message.ehr !== undefined) {
      Ehr.encode(message.ehr, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddEhrResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddEhrResponse } as MsgAddEhrResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ehr = Ehr.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddEhrResponse {
    const message = { ...baseMsgAddEhrResponse } as MsgAddEhrResponse;
    if (object.ehr !== undefined && object.ehr !== null) {
      message.ehr = Ehr.fromJSON(object.ehr);
    } else {
      message.ehr = undefined;
    }
    return message;
  },

  toJSON(message: MsgAddEhrResponse): unknown {
    const obj: any = {};
    message.ehr !== undefined &&
      (obj.ehr = message.ehr ? Ehr.toJSON(message.ehr) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddEhrResponse>): MsgAddEhrResponse {
    const message = { ...baseMsgAddEhrResponse } as MsgAddEhrResponse;
    if (object.ehr !== undefined && object.ehr !== null) {
      message.ehr = Ehr.fromPartial(object.ehr);
    } else {
      message.ehr = undefined;
    }
    return message;
  },
};

const baseMsgSendIbcTransferEhrs: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  patientId: 0,
};

export const MsgSendIbcTransferEhrs = {
  encode(
    message: MsgSendIbcTransferEhrs,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.patientId !== 0) {
      writer.uint32(40).uint64(message.patientId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendIbcTransferEhrs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendIbcTransferEhrs } as MsgSendIbcTransferEhrs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.patientId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendIbcTransferEhrs {
    const message = { ...baseMsgSendIbcTransferEhrs } as MsgSendIbcTransferEhrs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.patientId !== undefined && object.patientId !== null) {
      message.patientId = Number(object.patientId);
    } else {
      message.patientId = 0;
    }
    return message;
  },

  toJSON(message: MsgSendIbcTransferEhrs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.patientId !== undefined && (obj.patientId = message.patientId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendIbcTransferEhrs>
  ): MsgSendIbcTransferEhrs {
    const message = { ...baseMsgSendIbcTransferEhrs } as MsgSendIbcTransferEhrs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.patientId !== undefined && object.patientId !== null) {
      message.patientId = object.patientId;
    } else {
      message.patientId = 0;
    }
    return message;
  },
};

const baseMsgSendIbcTransferEhrsResponse: object = {};

export const MsgSendIbcTransferEhrsResponse = {
  encode(
    _: MsgSendIbcTransferEhrsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendIbcTransferEhrsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendIbcTransferEhrsResponse,
    } as MsgSendIbcTransferEhrsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSendIbcTransferEhrsResponse {
    const message = {
      ...baseMsgSendIbcTransferEhrsResponse,
    } as MsgSendIbcTransferEhrsResponse;
    return message;
  },

  toJSON(_: MsgSendIbcTransferEhrsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendIbcTransferEhrsResponse>
  ): MsgSendIbcTransferEhrsResponse {
    const message = {
      ...baseMsgSendIbcTransferEhrsResponse,
    } as MsgSendIbcTransferEhrsResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddEhr(request: MsgAddEhr): Promise<MsgAddEhrResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendIbcTransferEhrs(
    request: MsgSendIbcTransferEhrs
  ): Promise<MsgSendIbcTransferEhrsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AddEhr(request: MsgAddEhr): Promise<MsgAddEhrResponse> {
    const data = MsgAddEhr.encode(request).finish();
    const promise = this.rpc.request(
      "ashrafmohey.cache.ehr.Msg",
      "AddEhr",
      data
    );
    return promise.then((data) => MsgAddEhrResponse.decode(new Reader(data)));
  }

  SendIbcTransferEhrs(
    request: MsgSendIbcTransferEhrs
  ): Promise<MsgSendIbcTransferEhrsResponse> {
    const data = MsgSendIbcTransferEhrs.encode(request).finish();
    const promise = this.rpc.request(
      "ashrafmohey.cache.ehr.Msg",
      "SendIbcTransferEhrs",
      data
    );
    return promise.then((data) =>
      MsgSendIbcTransferEhrsResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
