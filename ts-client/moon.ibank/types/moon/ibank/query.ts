/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import { Transaction } from "./transaction";

export const protobufPackage = "moon.ibank";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetTransactionRequest {
  id: number;
}

export interface QueryGetTransactionResponse {
  Transaction: Transaction | undefined;
}

export interface QueryAllTransactionRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTransactionResponse {
  Transaction: Transaction[];
  pagination: PageResponse | undefined;
}

export interface QueryShowIncomingRequest {
  receiver: string;
  pending: boolean;
  pagination: PageRequest | undefined;
}

export interface QueryShowIncomingResponse {
  Transaction: Transaction[];
  pagination: PageResponse | undefined;
}

export interface QueryShowOutgoingRequest {
  sender: string;
  pending: boolean;
  pagination: PageRequest | undefined;
}

export interface QueryShowOutgoingResponse {
  Transaction: Transaction[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetTransactionRequest(): QueryGetTransactionRequest {
  return { id: 0 };
}

export const QueryGetTransactionRequest = {
  encode(message: QueryGetTransactionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTransactionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTransactionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTransactionRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetTransactionRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTransactionRequest>, I>>(object: I): QueryGetTransactionRequest {
    const message = createBaseQueryGetTransactionRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetTransactionResponse(): QueryGetTransactionResponse {
  return { Transaction: undefined };
}

export const QueryGetTransactionResponse = {
  encode(message: QueryGetTransactionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Transaction !== undefined) {
      Transaction.encode(message.Transaction, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTransactionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTransactionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Transaction = Transaction.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTransactionResponse {
    return { Transaction: isSet(object.Transaction) ? Transaction.fromJSON(object.Transaction) : undefined };
  },

  toJSON(message: QueryGetTransactionResponse): unknown {
    const obj: any = {};
    message.Transaction !== undefined
      && (obj.Transaction = message.Transaction ? Transaction.toJSON(message.Transaction) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTransactionResponse>, I>>(object: I): QueryGetTransactionResponse {
    const message = createBaseQueryGetTransactionResponse();
    message.Transaction = (object.Transaction !== undefined && object.Transaction !== null)
      ? Transaction.fromPartial(object.Transaction)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTransactionRequest(): QueryAllTransactionRequest {
  return { pagination: undefined };
}

export const QueryAllTransactionRequest = {
  encode(message: QueryAllTransactionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTransactionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTransactionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTransactionRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllTransactionRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllTransactionRequest>, I>>(object: I): QueryAllTransactionRequest {
    const message = createBaseQueryAllTransactionRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTransactionResponse(): QueryAllTransactionResponse {
  return { Transaction: [], pagination: undefined };
}

export const QueryAllTransactionResponse = {
  encode(message: QueryAllTransactionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Transaction) {
      Transaction.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTransactionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTransactionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Transaction.push(Transaction.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTransactionResponse {
    return {
      Transaction: Array.isArray(object?.Transaction)
        ? object.Transaction.map((e: any) => Transaction.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllTransactionResponse): unknown {
    const obj: any = {};
    if (message.Transaction) {
      obj.Transaction = message.Transaction.map((e) => e ? Transaction.toJSON(e) : undefined);
    } else {
      obj.Transaction = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllTransactionResponse>, I>>(object: I): QueryAllTransactionResponse {
    const message = createBaseQueryAllTransactionResponse();
    message.Transaction = object.Transaction?.map((e) => Transaction.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryShowIncomingRequest(): QueryShowIncomingRequest {
  return { receiver: "", pending: false, pagination: undefined };
}

export const QueryShowIncomingRequest = {
  encode(message: QueryShowIncomingRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.receiver !== "") {
      writer.uint32(10).string(message.receiver);
    }
    if (message.pending === true) {
      writer.uint32(16).bool(message.pending);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryShowIncomingRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryShowIncomingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.receiver = reader.string();
          break;
        case 2:
          message.pending = reader.bool();
          break;
        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryShowIncomingRequest {
    return {
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      pending: isSet(object.pending) ? Boolean(object.pending) : false,
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryShowIncomingRequest): unknown {
    const obj: any = {};
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.pending !== undefined && (obj.pending = message.pending);
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowIncomingRequest>, I>>(object: I): QueryShowIncomingRequest {
    const message = createBaseQueryShowIncomingRequest();
    message.receiver = object.receiver ?? "";
    message.pending = object.pending ?? false;
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryShowIncomingResponse(): QueryShowIncomingResponse {
  return { Transaction: [], pagination: undefined };
}

export const QueryShowIncomingResponse = {
  encode(message: QueryShowIncomingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Transaction) {
      Transaction.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryShowIncomingResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryShowIncomingResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Transaction.push(Transaction.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryShowIncomingResponse {
    return {
      Transaction: Array.isArray(object?.Transaction)
        ? object.Transaction.map((e: any) => Transaction.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryShowIncomingResponse): unknown {
    const obj: any = {};
    if (message.Transaction) {
      obj.Transaction = message.Transaction.map((e) => e ? Transaction.toJSON(e) : undefined);
    } else {
      obj.Transaction = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowIncomingResponse>, I>>(object: I): QueryShowIncomingResponse {
    const message = createBaseQueryShowIncomingResponse();
    message.Transaction = object.Transaction?.map((e) => Transaction.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryShowOutgoingRequest(): QueryShowOutgoingRequest {
  return { sender: "", pending: false, pagination: undefined };
}

export const QueryShowOutgoingRequest = {
  encode(message: QueryShowOutgoingRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.pending === true) {
      writer.uint32(16).bool(message.pending);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryShowOutgoingRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryShowOutgoingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.pending = reader.bool();
          break;
        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryShowOutgoingRequest {
    return {
      sender: isSet(object.sender) ? String(object.sender) : "",
      pending: isSet(object.pending) ? Boolean(object.pending) : false,
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryShowOutgoingRequest): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.pending !== undefined && (obj.pending = message.pending);
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowOutgoingRequest>, I>>(object: I): QueryShowOutgoingRequest {
    const message = createBaseQueryShowOutgoingRequest();
    message.sender = object.sender ?? "";
    message.pending = object.pending ?? false;
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryShowOutgoingResponse(): QueryShowOutgoingResponse {
  return { Transaction: [], pagination: undefined };
}

export const QueryShowOutgoingResponse = {
  encode(message: QueryShowOutgoingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Transaction) {
      Transaction.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryShowOutgoingResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryShowOutgoingResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Transaction.push(Transaction.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryShowOutgoingResponse {
    return {
      Transaction: Array.isArray(object?.Transaction)
        ? object.Transaction.map((e: any) => Transaction.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryShowOutgoingResponse): unknown {
    const obj: any = {};
    if (message.Transaction) {
      obj.Transaction = message.Transaction.map((e) => e ? Transaction.toJSON(e) : undefined);
    } else {
      obj.Transaction = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowOutgoingResponse>, I>>(object: I): QueryShowOutgoingResponse {
    const message = createBaseQueryShowOutgoingResponse();
    message.Transaction = object.Transaction?.map((e) => Transaction.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Transaction by id. */
  Transaction(request: QueryGetTransactionRequest): Promise<QueryGetTransactionResponse>;
  /** Queries a list of Transaction items. */
  TransactionAll(request: QueryAllTransactionRequest): Promise<QueryAllTransactionResponse>;
  /** Queries a list of ShowIncoming items. */
  ShowIncoming(request: QueryShowIncomingRequest): Promise<QueryShowIncomingResponse>;
  /** Queries a list of ShowOutgoing items. */
  ShowOutgoing(request: QueryShowOutgoingRequest): Promise<QueryShowOutgoingResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Transaction = this.Transaction.bind(this);
    this.TransactionAll = this.TransactionAll.bind(this);
    this.ShowIncoming = this.ShowIncoming.bind(this);
    this.ShowOutgoing = this.ShowOutgoing.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("moon.ibank.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Transaction(request: QueryGetTransactionRequest): Promise<QueryGetTransactionResponse> {
    const data = QueryGetTransactionRequest.encode(request).finish();
    const promise = this.rpc.request("moon.ibank.Query", "Transaction", data);
    return promise.then((data) => QueryGetTransactionResponse.decode(new _m0.Reader(data)));
  }

  TransactionAll(request: QueryAllTransactionRequest): Promise<QueryAllTransactionResponse> {
    const data = QueryAllTransactionRequest.encode(request).finish();
    const promise = this.rpc.request("moon.ibank.Query", "TransactionAll", data);
    return promise.then((data) => QueryAllTransactionResponse.decode(new _m0.Reader(data)));
  }

  ShowIncoming(request: QueryShowIncomingRequest): Promise<QueryShowIncomingResponse> {
    const data = QueryShowIncomingRequest.encode(request).finish();
    const promise = this.rpc.request("moon.ibank.Query", "ShowIncoming", data);
    return promise.then((data) => QueryShowIncomingResponse.decode(new _m0.Reader(data)));
  }

  ShowOutgoing(request: QueryShowOutgoingRequest): Promise<QueryShowOutgoingResponse> {
    const data = QueryShowOutgoingRequest.encode(request).finish();
    const promise = this.rpc.request("moon.ibank.Query", "ShowOutgoing", data);
    return promise.then((data) => QueryShowOutgoingResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
