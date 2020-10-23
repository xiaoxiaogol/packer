// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: yandex/cloud/certificatemanager/v1/certificate_content_service.proto

package certificatemanager

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetCertificateContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the certificate
	CertificateId string `protobuf:"bytes,1,opt,name=certificate_id,json=certificateId,proto3" json:"certificate_id,omitempty"`
	// PEM-encoded certificate chain content of the certificate.
	CertificateChain []string `protobuf:"bytes,3,rep,name=certificate_chain,json=certificateChain,proto3" json:"certificate_chain,omitempty"`
	// PEM-encoded private key content of the certificate.
	PrivateKey string `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
}

func (x *GetCertificateContentResponse) Reset() {
	*x = GetCertificateContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertificateContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertificateContentResponse) ProtoMessage() {}

func (x *GetCertificateContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertificateContentResponse.ProtoReflect.Descriptor instead.
func (*GetCertificateContentResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCertificateContentResponse) GetCertificateId() string {
	if x != nil {
		return x.CertificateId
	}
	return ""
}

func (x *GetCertificateContentResponse) GetCertificateChain() []string {
	if x != nil {
		return x.CertificateChain
	}
	return nil
}

func (x *GetCertificateContentResponse) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

type GetCertificateContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the certificate to download content.
	CertificateId string `protobuf:"bytes,1,opt,name=certificate_id,json=certificateId,proto3" json:"certificate_id,omitempty"`
}

func (x *GetCertificateContentRequest) Reset() {
	*x = GetCertificateContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertificateContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertificateContentRequest) ProtoMessage() {}

func (x *GetCertificateContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertificateContentRequest.ProtoReflect.Descriptor instead.
func (*GetCertificateContentRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCertificateContentRequest) GetCertificateId() string {
	if x != nil {
		return x.CertificateId
	}
	return ""
}

var File_yandex_cloud_certificatemanager_v1_certificate_content_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDesc = []byte{
	0x0a, 0x44, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x2b, 0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x45,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x32, 0xf2, 0x01, 0x0a, 0x19, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xd4, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x40, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x48, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x42, 0x12, 0x40, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a,
	0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x83, 0x01, 0x0a, 0x26, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescData = file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDesc
)

func file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescData)
	})
	return file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDescData
}

var file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_goTypes = []interface{}{
	(*GetCertificateContentResponse)(nil), // 0: yandex.cloud.certificatemanager.v1.GetCertificateContentResponse
	(*GetCertificateContentRequest)(nil),  // 1: yandex.cloud.certificatemanager.v1.GetCertificateContentRequest
}
var file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.certificatemanager.v1.CertificateContentService.Get:input_type -> yandex.cloud.certificatemanager.v1.GetCertificateContentRequest
	0, // 1: yandex.cloud.certificatemanager.v1.CertificateContentService.Get:output_type -> yandex.cloud.certificatemanager.v1.GetCertificateContentResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_init() }
func file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_init() {
	if File_yandex_cloud_certificatemanager_v1_certificate_content_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertificateContentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertificateContentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_certificatemanager_v1_certificate_content_service_proto = out.File
	file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_rawDesc = nil
	file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_goTypes = nil
	file_yandex_cloud_certificatemanager_v1_certificate_content_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CertificateContentServiceClient is the client API for CertificateContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateContentServiceClient interface {
	// Returns chain and private key of the specified certificate
	Get(ctx context.Context, in *GetCertificateContentRequest, opts ...grpc.CallOption) (*GetCertificateContentResponse, error)
}

type certificateContentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateContentServiceClient(cc grpc.ClientConnInterface) CertificateContentServiceClient {
	return &certificateContentServiceClient{cc}
}

func (c *certificateContentServiceClient) Get(ctx context.Context, in *GetCertificateContentRequest, opts ...grpc.CallOption) (*GetCertificateContentResponse, error) {
	out := new(GetCertificateContentResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.certificatemanager.v1.CertificateContentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateContentServiceServer is the server API for CertificateContentService service.
type CertificateContentServiceServer interface {
	// Returns chain and private key of the specified certificate
	Get(context.Context, *GetCertificateContentRequest) (*GetCertificateContentResponse, error)
}

// UnimplementedCertificateContentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateContentServiceServer struct {
}

func (*UnimplementedCertificateContentServiceServer) Get(context.Context, *GetCertificateContentRequest) (*GetCertificateContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterCertificateContentServiceServer(s *grpc.Server, srv CertificateContentServiceServer) {
	s.RegisterService(&_CertificateContentService_serviceDesc, srv)
}

func _CertificateContentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateContentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.certificatemanager.v1.CertificateContentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateContentServiceServer).Get(ctx, req.(*GetCertificateContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.certificatemanager.v1.CertificateContentService",
	HandlerType: (*CertificateContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CertificateContentService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/certificatemanager/v1/certificate_content_service.proto",
}
