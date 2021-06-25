// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: client.proto

package api

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	port "micro_services/api/v1/port"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IDMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *IDMsg) Reset() {
	*x = IDMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDMsg) ProtoMessage() {}

func (x *IDMsg) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDMsg.ProtoReflect.Descriptor instead.
func (*IDMsg) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *IDMsg) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x76, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x70, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x17, 0x0a, 0x05, 0x49, 0x44, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0x51, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x44, 0x4d, 0x73, 0x67,
	0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42, 0x08, 0x5a, 0x06, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_client_proto_goTypes = []interface{}{
	(*IDMsg)(nil),     // 0: v1.api.IDMsg
	(*port.Port)(nil), // 1: v1.pds.Port
}
var file_client_proto_depIdxs = []int32{
	0, // 0: v1.api.APIService.GetByID:input_type -> v1.api.IDMsg
	1, // 1: v1.api.APIService.GetByID:output_type -> v1.pds.Port
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDMsg); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIServiceClient interface {
	GetByID(ctx context.Context, in *IDMsg, opts ...grpc.CallOption) (*port.Port, error)
}

type aPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIServiceClient(cc grpc.ClientConnInterface) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) GetByID(ctx context.Context, in *IDMsg, opts ...grpc.CallOption) (*port.Port, error) {
	out := new(port.Port)
	err := c.cc.Invoke(ctx, "/v1.api.APIService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
type APIServiceServer interface {
	GetByID(context.Context, *IDMsg) (*port.Port, error)
}

// UnimplementedAPIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (*UnimplementedAPIServiceServer) GetByID(context.Context, *IDMsg) (*port.Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}

func RegisterAPIServiceServer(s *grpc.Server, srv APIServiceServer) {
	s.RegisterService(&_APIService_serviceDesc, srv)
}

func _APIService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.api.APIService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetByID(ctx, req.(*IDMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.api.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _APIService_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}