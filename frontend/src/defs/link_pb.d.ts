// package: links
// file: defs/link.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class NewLink extends jspb.Message {
  getShortUrl(): string;
  setShortUrl(value: string): void;

  getTargetUrl(): string;
  setTargetUrl(value: string): void;

  getPermanent(): boolean;
  setPermanent(value: boolean): void;

  getOwner(): string;
  setOwner(value: string): void;

  getProtected(): boolean;
  setProtected(value: boolean): void;

  getPrivate(): boolean;
  setPrivate(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewLink.AsObject;
  static toObject(includeInstance: boolean, msg: NewLink): NewLink.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NewLink, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewLink;
  static deserializeBinaryFromReader(message: NewLink, reader: jspb.BinaryReader): NewLink;
}

export namespace NewLink {
  export type AsObject = {
    shortUrl: string,
    targetUrl: string,
    permanent: boolean,
    owner: string,
    pb_protected: boolean,
    pb_private: boolean,
  }
}

export class Link extends jspb.Message {
  getShortUrl(): string;
  setShortUrl(value: string): void;

  getTargetUrl(): string;
  setTargetUrl(value: string): void;

  getPermanent(): boolean;
  setPermanent(value: boolean): void;

  hasMetadata(): boolean;
  clearMetadata(): void;
  getMetadata(): Link.LinksMetadata | undefined;
  setMetadata(value?: Link.LinksMetadata): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Link.AsObject;
  static toObject(includeInstance: boolean, msg: Link): Link.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Link, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Link;
  static deserializeBinaryFromReader(message: Link, reader: jspb.BinaryReader): Link;
}

export namespace Link {
  export type AsObject = {
    shortUrl: string,
    targetUrl: string,
    permanent: boolean,
    metadata?: Link.LinksMetadata.AsObject,
  }

  export class LinksMetadata extends jspb.Message {
    getOwner(): string;
    setOwner(value: string): void;

    getProtected(): boolean;
    setProtected(value: boolean): void;

    getPrivate(): boolean;
    setPrivate(value: boolean): void;

    hasCreated(): boolean;
    clearCreated(): void;
    getCreated(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCreated(value?: google_protobuf_timestamp_pb.Timestamp): void;

    hasModified(): boolean;
    clearModified(): void;
    getModified(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setModified(value?: google_protobuf_timestamp_pb.Timestamp): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LinksMetadata.AsObject;
    static toObject(includeInstance: boolean, msg: LinksMetadata): LinksMetadata.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LinksMetadata, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LinksMetadata;
    static deserializeBinaryFromReader(message: LinksMetadata, reader: jspb.BinaryReader): LinksMetadata;
  }

  export namespace LinksMetadata {
    export type AsObject = {
      owner: string,
      pb_protected: boolean,
      pb_private: boolean,
      created?: google_protobuf_timestamp_pb.Timestamp.AsObject,
      modified?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
  }
}

