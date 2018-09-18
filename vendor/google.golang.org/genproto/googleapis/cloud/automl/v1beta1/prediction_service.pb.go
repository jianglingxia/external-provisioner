// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/prediction_service.proto

package automl

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [PredictionService.Predict][google.cloud.automl.v1beta1.PredictionService.Predict].
type PredictRequest struct {
	// Name of the model requested to serve the prediction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	// Payload to perform a prediction on. The payload must match the
	// problem type that the model was trained to solve.
	Payload *ExamplePayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific parameters, any string must be up to 25000
	// characters long.
	//
	// *  For Image Classification:
	//
	//    `score_threshold` - (float) A value from 0.0 to 1.0. When the model
	//     makes predictions for an
	//     image, it will only produce results that have at least this confidence
	//     score threshold. The default is 0.5.
	Params               map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictRequest) Reset()         { *m = PredictRequest{} }
func (m *PredictRequest) String() string { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()    {}
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59a9dba5da3c687d, []int{0}
}
func (m *PredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictRequest.Unmarshal(m, b)
}
func (m *PredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictRequest.Marshal(b, m, deterministic)
}
func (m *PredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictRequest.Merge(m, src)
}
func (m *PredictRequest) XXX_Size() int {
	return xxx_messageInfo_PredictRequest.Size(m)
}
func (m *PredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictRequest proto.InternalMessageInfo

func (m *PredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictRequest) GetPayload() *ExamplePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Response message for [PredictionService.Predict][google.cloud.automl.v1beta1.PredictionService.Predict].
//
// Currently, this is only
// used to return an image recognition prediction result. More prediction
// output metadata might be introduced in the future.
type PredictResponse struct {
	// Prediction result.
	Payload []*AnnotationPayload `protobuf:"bytes,1,rep,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific prediction response metadata.
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictResponse) Reset()         { *m = PredictResponse{} }
func (m *PredictResponse) String() string { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()    {}
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59a9dba5da3c687d, []int{1}
}
func (m *PredictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictResponse.Unmarshal(m, b)
}
func (m *PredictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictResponse.Marshal(b, m, deterministic)
}
func (m *PredictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictResponse.Merge(m, src)
}
func (m *PredictResponse) XXX_Size() int {
	return xxx_messageInfo_PredictResponse.Size(m)
}
func (m *PredictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictResponse proto.InternalMessageInfo

func (m *PredictResponse) GetPayload() []*AnnotationPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "google.cloud.automl.v1beta1.PredictRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1beta1.PredictRequest.ParamsEntry")
	proto.RegisterType((*PredictResponse)(nil), "google.cloud.automl.v1beta1.PredictResponse")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1beta1.PredictResponse.MetadataEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Perform a prediction.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
}

type predictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictionServiceClient(cc *grpc.ClientConn) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.automl.v1beta1.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Perform a prediction.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.automl.v1beta1.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.automl.v1beta1.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/automl/v1beta1/prediction_service.proto",
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/prediction_service.proto", fileDescriptor_59a9dba5da3c687d)
}

var fileDescriptor_59a9dba5da3c687d = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x99, 0x56, 0x77, 0xdd, 0x29, 0xfe, 0x1b, 0x44, 0x4a, 0x14, 0x2c, 0xbd, 0x2a, 0xdd,
	0x65, 0x86, 0xad, 0xc2, 0x6a, 0x57, 0x2f, 0xb6, 0x50, 0xf0, 0x46, 0x0c, 0x11, 0xbc, 0xf0, 0xa6,
	0x9c, 0x4d, 0x0e, 0x21, 0x3a, 0xc9, 0x8c, 0x99, 0x49, 0xb1, 0x88, 0x5e, 0xf8, 0x0a, 0xbe, 0x84,
	0x4f, 0xe1, 0x4b, 0xf8, 0x0a, 0x3e, 0x80, 0x77, 0xde, 0x4a, 0x32, 0x63, 0xd6, 0x65, 0x25, 0x6c,
	0xef, 0x32, 0x7f, 0x7e, 0xdf, 0xf9, 0xbe, 0xcc, 0x39, 0xf4, 0x51, 0xaa, 0x54, 0x2a, 0x51, 0xc4,
	0x52, 0x55, 0x89, 0x80, 0xca, 0xaa, 0x5c, 0x8a, 0xf5, 0xe1, 0x29, 0x5a, 0x38, 0x14, 0xba, 0xc4,
	0x24, 0x8b, 0x6d, 0xa6, 0x8a, 0x95, 0xc1, 0x72, 0x9d, 0xc5, 0xc8, 0x75, 0xa9, 0xac, 0x62, 0xf7,
	0x1c, 0xc5, 0x1b, 0x8a, 0x3b, 0x8a, 0x7b, 0x2a, 0xb8, 0xef, 0x25, 0x41, 0x67, 0x02, 0x8a, 0x42,
	0x59, 0xa8, 0x15, 0x8c, 0x43, 0x83, 0xce, 0x82, 0x67, 0xd7, 0x57, 0x1a, 0x36, 0x52, 0x41, 0xe2,
	0xa9, 0x83, 0x2e, 0x2a, 0x01, 0x0b, 0xab, 0xcc, 0x62, 0xee, 0x6b, 0x8c, 0x7f, 0x13, 0x7a, 0x23,
	0x74, 0xde, 0x23, 0x7c, 0x5f, 0xa1, 0xb1, 0x8c, 0xd1, 0x2b, 0x05, 0xe4, 0x38, 0x24, 0x23, 0x32,
	0xd9, 0x8b, 0x9a, 0x6f, 0xb6, 0xa4, 0xbb, 0xbe, 0xca, 0xb0, 0x37, 0x22, 0x93, 0xc1, 0x6c, 0x9f,
	0x77, 0xe4, 0xe2, 0xcb, 0x0f, 0x90, 0x6b, 0x89, 0xa1, 0x43, 0xa2, 0xbf, 0x2c, 0x7b, 0x49, 0x77,
	0x34, 0x94, 0x90, 0x9b, 0x61, 0x7f, 0xd4, 0x9f, 0x0c, 0x66, 0x47, 0x9d, 0x2a, 0xe7, 0x7d, 0xf1,
	0xb0, 0x21, 0x97, 0x85, 0x2d, 0x37, 0x91, 0x97, 0x09, 0x9e, 0xd0, 0xc1, 0x3f, 0xdb, 0xec, 0x16,
	0xed, 0xbf, 0xc3, 0x8d, 0x77, 0x5e, 0x7f, 0xb2, 0x3b, 0xf4, 0xea, 0x1a, 0x64, 0x85, 0x8d, 0xed,
	0xbd, 0xc8, 0x2d, 0xe6, 0xbd, 0xc7, 0x64, 0xfc, 0x8b, 0xd0, 0x9b, 0x6d, 0x05, 0xa3, 0x55, 0x61,
	0x90, 0x3d, 0x3f, 0x8b, 0x49, 0x1a, 0x83, 0xbc, 0xd3, 0xe0, 0x49, 0xfb, 0x06, 0x17, 0x92, 0xbe,
	0xa6, 0xd7, 0x72, 0xb4, 0x50, 0xff, 0xef, 0x61, 0xaf, 0x91, 0x9a, 0x5f, 0x2e, 0xab, 0x73, 0xc2,
	0x5f, 0x78, 0xd8, 0xc5, 0x6d, 0xb5, 0x82, 0x63, 0x7a, 0xfd, 0xdc, 0xd1, 0x36, 0x91, 0x67, 0xdf,
	0x09, 0xbd, 0x1d, 0xb6, 0x8d, 0xfa, 0xca, 0xf5, 0x29, 0xfb, 0x46, 0xe8, 0xae, 0xdf, 0x65, 0xfb,
	0x5b, 0x3c, 0x48, 0x70, 0xb0, 0x4d, 0xa2, 0xf1, 0xe2, 0xcb, 0x8f, 0x9f, 0x5f, 0x7b, 0x4f, 0xc7,
	0x47, 0x6d, 0x33, 0x7e, 0xac, 0x5b, 0xeb, 0x99, 0x2e, 0xd5, 0x5b, 0x8c, 0xad, 0x11, 0x53, 0x21,
	0x55, 0xec, 0x06, 0x40, 0x4c, 0x45, 0xae, 0x12, 0x94, 0x46, 0x4c, 0x3f, 0xcd, 0xfd, 0x68, 0xcd,
	0xc9, 0x74, 0xf1, 0x99, 0x3e, 0x88, 0x55, 0xde, 0x55, 0x76, 0x71, 0xf7, 0x42, 0xc0, 0xb0, 0x6e,
	0xf4, 0x90, 0xbc, 0x39, 0xf1, 0x58, 0xaa, 0x24, 0x14, 0x29, 0x57, 0x65, 0x2a, 0x52, 0x2c, 0x9a,
	0x31, 0x10, 0xee, 0x08, 0x74, 0x66, 0xfe, 0x3b, 0x37, 0xc7, 0x6e, 0x79, 0xba, 0xd3, 0xdc, 0x7e,
	0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xb3, 0x24, 0x52, 0x0b, 0x04, 0x00, 0x00,
}
