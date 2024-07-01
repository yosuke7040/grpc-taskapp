// @generated by protoc-gen-connect-query v0.4.1 with parameter "target=ts"
// @generated from file rpc/auth/v1/auth.proto (package rpc.auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { createQueryService } from "@bufbuild/connect-query";
import { MethodKind } from "@bufbuild/protobuf";
import { LoginRequest, LoginResponse } from "./auth_pb.js";

export const typeName = "rpc.auth.v1.AuthService";

/**
 * @generated from rpc rpc.auth.v1.AuthService.Login
 */
export const login = createQueryService({
  service: {
    methods: {
      login: {
        name: "Login",
        kind: MethodKind.Unary,
        I: LoginRequest,
        O: LoginResponse,
      },
    },
    typeName: "rpc.auth.v1.AuthService",
  },
}).login;