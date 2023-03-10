// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.12.4
// source: metadata/service/v1/metadata.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMetadataServiceGetMetadata = "/metadata.service.v1.MetadataService/GetMetadata"

type MetadataServiceHTTPServer interface {
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)
}

func RegisterMetadataServiceHTTPServer(s *http.Server, srv MetadataServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/metadata/{id}", _MetadataService_GetMetadata0_HTTP_Handler(srv))
}

func _MetadataService_GetMetadata0_HTTP_Handler(srv MetadataServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMetadataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMetadataServiceGetMetadata)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMetadata(ctx, req.(*GetMetadataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMetadataResponse)
		return ctx.Result(200, reply)
	}
}

type MetadataServiceHTTPClient interface {
	GetMetadata(ctx context.Context, req *GetMetadataRequest, opts ...http.CallOption) (rsp *GetMetadataResponse, err error)
}

type MetadataServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewMetadataServiceHTTPClient(client *http.Client) MetadataServiceHTTPClient {
	return &MetadataServiceHTTPClientImpl{client}
}

func (c *MetadataServiceHTTPClientImpl) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...http.CallOption) (*GetMetadataResponse, error) {
	var out GetMetadataResponse
	pattern := "/metadata/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMetadataServiceGetMetadata))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
