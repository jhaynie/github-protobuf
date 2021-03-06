// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/cloud/ml/v1beta1/project_service.proto
// DO NOT EDIT!

package google_cloud_ml_v1beta1 // import "google.golang.org/genproto/googleapis/cloud/ml/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Requests service account information associated with a project.
type GetConfigRequest struct {
	// Required. The project name.
	//
	// Authorization: requires `Viewer` role on the specified project.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetConfigRequest) Reset()                    { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()               {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// Returns service account information associated with a project.
type GetConfigResponse struct {
	// The service account Cloud ML uses to access resources in the project.
	ServiceAccount string `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
	// The project number for `service_account`.
	ServiceAccountProject int64 `protobuf:"varint,2,opt,name=service_account_project,json=serviceAccountProject" json:"service_account_project,omitempty"`
}

func (m *GetConfigResponse) Reset()                    { *m = GetConfigResponse{} }
func (m *GetConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*GetConfigResponse) ProtoMessage()               {}
func (*GetConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func init() {
	proto.RegisterType((*GetConfigRequest)(nil), "google.cloud.ml.v1beta1.GetConfigRequest")
	proto.RegisterType((*GetConfigResponse)(nil), "google.cloud.ml.v1beta1.GetConfigResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProjectManagementService service

type ProjectManagementServiceClient interface {
	// Get the service account information associated with your project. You need
	// this information in order to grant the service account persmissions for
	// the Google Cloud Storage location where you put your model training code
	// for training the model with Google Cloud Machine Learning.
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
}

type projectManagementServiceClient struct {
	cc *grpc.ClientConn
}

func NewProjectManagementServiceClient(cc *grpc.ClientConn) ProjectManagementServiceClient {
	return &projectManagementServiceClient{cc}
}

func (c *projectManagementServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.ProjectManagementService/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProjectManagementService service

type ProjectManagementServiceServer interface {
	// Get the service account information associated with your project. You need
	// this information in order to grant the service account persmissions for
	// the Google Cloud Storage location where you put your model training code
	// for training the model with Google Cloud Machine Learning.
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
}

func RegisterProjectManagementServiceServer(s *grpc.Server, srv ProjectManagementServiceServer) {
	s.RegisterService(&_ProjectManagementService_serviceDesc, srv)
}

func _ProjectManagementService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectManagementServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.ProjectManagementService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectManagementServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.ml.v1beta1.ProjectManagementService",
	HandlerType: (*ProjectManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ProjectManagementService_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google.golang.org/genproto/googleapis/cloud/ml/v1beta1/project_service.proto",
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/cloud/ml/v1beta1/project_service.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4a, 0x43, 0x31,
	0x10, 0xe6, 0x55, 0x11, 0x9a, 0x85, 0x3f, 0x11, 0x69, 0xe9, 0xc6, 0x52, 0xa4, 0xd6, 0x22, 0x09,
	0x55, 0x10, 0x11, 0x5c, 0x58, 0x17, 0x22, 0x28, 0x94, 0x7a, 0x80, 0x92, 0xc6, 0x31, 0x44, 0xde,
	0xcb, 0xc4, 0x97, 0xbc, 0x6e, 0xc4, 0x8d, 0x27, 0x10, 0x3c, 0x87, 0xa7, 0xf1, 0x0a, 0x1e, 0x44,
	0xfa, 0x92, 0x16, 0xad, 0x08, 0xdd, 0x0d, 0x33, 0xdf, 0xf7, 0xcd, 0x7c, 0xf3, 0x91, 0x1b, 0x85,
	0xa8, 0x52, 0x60, 0x0a, 0x53, 0x61, 0x14, 0xc3, 0x5c, 0x71, 0x05, 0xc6, 0xe6, 0xe8, 0x91, 0x87,
	0x91, 0xb0, 0xda, 0x71, 0x99, 0x62, 0x71, 0xcf, 0xb3, 0x94, 0x4f, 0x7a, 0x63, 0xf0, 0xa2, 0xc7,
	0x6d, 0x8e, 0x8f, 0x20, 0xfd, 0xc8, 0x41, 0x3e, 0xd1, 0x12, 0x58, 0xc9, 0xa0, 0xb5, 0xa8, 0x56,
	0xc2, 0x59, 0x96, 0xb2, 0x08, 0x6f, 0x5c, 0x2f, 0xb7, 0x46, 0x58, 0xcd, 0xa3, 0xa2, 0x44, 0xf3,
	0xa0, 0x15, 0x17, 0xc6, 0xa0, 0x17, 0x5e, 0xa3, 0x71, 0x61, 0x47, 0xab, 0x4d, 0x36, 0xaf, 0xc0,
	0x5f, 0x96, 0xe3, 0x21, 0x3c, 0x15, 0xe0, 0x3c, 0xa5, 0x64, 0xd5, 0x88, 0x0c, 0xea, 0x49, 0x33,
	0xe9, 0x54, 0x87, 0x65, 0xdd, 0xf2, 0x64, 0xeb, 0x07, 0xce, 0x59, 0x34, 0x0e, 0xe8, 0x3e, 0xd9,
	0x88, 0xfa, 0x23, 0x21, 0x25, 0x16, 0xc6, 0x47, 0xce, 0x7a, 0x6c, 0x5f, 0x84, 0x2e, 0x3d, 0x21,
	0xb5, 0x05, 0xe0, 0x28, 0x5a, 0xae, 0x57, 0x9a, 0x49, 0x67, 0x65, 0xb8, 0xf3, 0x9b, 0x30, 0x08,
	0xc3, 0xa3, 0x8f, 0x84, 0xd4, 0x63, 0x7d, 0x2b, 0x8c, 0x50, 0x90, 0x81, 0xf1, 0x77, 0x01, 0x4a,
	0xdf, 0x12, 0x52, 0x9d, 0xdf, 0x44, 0x0f, 0xd8, 0x3f, 0xdf, 0x62, 0x8b, 0xfe, 0x1a, 0xdd, 0x65,
	0xa0, 0xc1, 0x62, 0xeb, 0xf0, 0xf5, 0xf3, 0xeb, 0xbd, 0xd2, 0xa6, 0x7b, 0xf3, 0xac, 0x9e, 0xa7,
	0xff, 0x38, 0x8f, 0xe7, 0x3b, 0xde, 0x7d, 0x39, 0x53, 0x33, 0x56, 0xff, 0x94, 0xec, 0x4a, 0xcc,
	0xfe, 0xc8, 0x0b, 0xab, 0x67, 0x2b, 0xfa, 0xdb, 0xd1, 0x4f, 0x74, 0x31, 0x98, 0xa6, 0x30, 0x48,
	0xc6, 0x6b, 0x65, 0x1c, 0xc7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01, 0x64, 0xc1, 0x94, 0x42,
	0x02, 0x00, 0x00,
}
