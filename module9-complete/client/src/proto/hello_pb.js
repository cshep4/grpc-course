// @generated by protoc-gen-es v1.10.0
// @generated from file hello.proto (package hello, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message hello.SayHelloRequest
 */
export const SayHelloRequest = /*@__PURE__*/ proto3.makeMessageType(
  "hello.SayHelloRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message hello.SayHelloResponse
 */
export const SayHelloResponse = /*@__PURE__*/ proto3.makeMessageType(
  "hello.SayHelloResponse",
  () => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);
