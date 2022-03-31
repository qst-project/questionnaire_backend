// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QuestionnaireClient is the client API for Questionnaire service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionnaireClient interface {
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
	GetSurvey(ctx context.Context, in *GetSurveyRequest, opts ...grpc.CallOption) (*GetSurveyResponse, error)
	CreateSurvey(ctx context.Context, in *CreateSurveyRequest, opts ...grpc.CallOption) (*CreateSurveyResponse, error)
	UpdateSurvey(ctx context.Context, in *UpdateSurveyRequest, opts ...grpc.CallOption) (*UpdateSurveyResponse, error)
	DeleteSurvey(ctx context.Context, in *DeleteSurveyRequest, opts ...grpc.CallOption) (*DeleteSurveyResponse, error)
}

type questionnaireClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionnaireClient(cc grpc.ClientConnInterface) QuestionnaireClient {
	return &questionnaireClient{cc}
}

func (c *questionnaireClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/api.Questionnaire/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) GetSurvey(ctx context.Context, in *GetSurveyRequest, opts ...grpc.CallOption) (*GetSurveyResponse, error) {
	out := new(GetSurveyResponse)
	err := c.cc.Invoke(ctx, "/api.Questionnaire/GetSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) CreateSurvey(ctx context.Context, in *CreateSurveyRequest, opts ...grpc.CallOption) (*CreateSurveyResponse, error) {
	out := new(CreateSurveyResponse)
	err := c.cc.Invoke(ctx, "/api.Questionnaire/CreateSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) UpdateSurvey(ctx context.Context, in *UpdateSurveyRequest, opts ...grpc.CallOption) (*UpdateSurveyResponse, error) {
	out := new(UpdateSurveyResponse)
	err := c.cc.Invoke(ctx, "/api.Questionnaire/UpdateSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) DeleteSurvey(ctx context.Context, in *DeleteSurveyRequest, opts ...grpc.CallOption) (*DeleteSurveyResponse, error) {
	out := new(DeleteSurveyResponse)
	err := c.cc.Invoke(ctx, "/api.Questionnaire/DeleteSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionnaireServer is the server API for Questionnaire service.
// All implementations must embed UnimplementedQuestionnaireServer
// for forward compatibility
type QuestionnaireServer interface {
	Test(context.Context, *TestRequest) (*TestResponse, error)
	GetSurvey(context.Context, *GetSurveyRequest) (*GetSurveyResponse, error)
	CreateSurvey(context.Context, *CreateSurveyRequest) (*CreateSurveyResponse, error)
	UpdateSurvey(context.Context, *UpdateSurveyRequest) (*UpdateSurveyResponse, error)
	DeleteSurvey(context.Context, *DeleteSurveyRequest) (*DeleteSurveyResponse, error)
	mustEmbedUnimplementedQuestionnaireServer()
}

// UnimplementedQuestionnaireServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionnaireServer struct {
}

func (UnimplementedQuestionnaireServer) Test(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedQuestionnaireServer) GetSurvey(context.Context, *GetSurveyRequest) (*GetSurveyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSurvey not implemented")
}
func (UnimplementedQuestionnaireServer) CreateSurvey(context.Context, *CreateSurveyRequest) (*CreateSurveyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSurvey not implemented")
}
func (UnimplementedQuestionnaireServer) UpdateSurvey(context.Context, *UpdateSurveyRequest) (*UpdateSurveyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSurvey not implemented")
}
func (UnimplementedQuestionnaireServer) DeleteSurvey(context.Context, *DeleteSurveyRequest) (*DeleteSurveyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSurvey not implemented")
}
func (UnimplementedQuestionnaireServer) mustEmbedUnimplementedQuestionnaireServer() {}

// UnsafeQuestionnaireServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionnaireServer will
// result in compilation errors.
type UnsafeQuestionnaireServer interface {
	mustEmbedUnimplementedQuestionnaireServer()
}

func RegisterQuestionnaireServer(s grpc.ServiceRegistrar, srv QuestionnaireServer) {
	s.RegisterService(&Questionnaire_ServiceDesc, srv)
}

func _Questionnaire_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Questionnaire/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_GetSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSurveyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).GetSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Questionnaire/GetSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).GetSurvey(ctx, req.(*GetSurveyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_CreateSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSurveyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).CreateSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Questionnaire/CreateSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).CreateSurvey(ctx, req.(*CreateSurveyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_UpdateSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSurveyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).UpdateSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Questionnaire/UpdateSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).UpdateSurvey(ctx, req.(*UpdateSurveyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_DeleteSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSurveyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).DeleteSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Questionnaire/DeleteSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).DeleteSurvey(ctx, req.(*DeleteSurveyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Questionnaire_ServiceDesc is the grpc.ServiceDesc for Questionnaire service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Questionnaire_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Questionnaire",
	HandlerType: (*QuestionnaireServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Questionnaire_Test_Handler,
		},
		{
			MethodName: "GetSurvey",
			Handler:    _Questionnaire_GetSurvey_Handler,
		},
		{
			MethodName: "CreateSurvey",
			Handler:    _Questionnaire_CreateSurvey_Handler,
		},
		{
			MethodName: "UpdateSurvey",
			Handler:    _Questionnaire_UpdateSurvey_Handler,
		},
		{
			MethodName: "DeleteSurvey",
			Handler:    _Questionnaire_DeleteSurvey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "questionnaire.proto",
}
