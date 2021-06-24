// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pds.proto

package portdomain

import (
	context "context"
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

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	City        string    `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	Country     string    `protobuf:"bytes,5,opt,name=Country,proto3" json:"Country,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,6,rep,packed,name=Coordinates,proto3" json:"Coordinates,omitempty"`
	Province    string    `protobuf:"bytes,7,opt,name=Province,proto3" json:"Province,omitempty"`
	Timezone    string    `protobuf:"bytes,8,opt,name=Timezone,proto3" json:"Timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,9,rep,name=Unlocs,proto3" json:"Unlocs,omitempty"`
	Code        string    `protobuf:"bytes,10,opt,name=Code,proto3" json:"Code,omitempty"`
	Regions     []string  `protobuf:"bytes,11,rep,name=Regions,proto3" json:"Regions,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_pds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_pds_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

type PortSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of ports received.
	PortCount int32 `protobuf:"varint,1,opt,name=port_count,json=portCount,proto3" json:"port_count,omitempty"`
}

func (x *PortSummary) Reset() {
	*x = PortSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortSummary) ProtoMessage() {}

func (x *PortSummary) ProtoReflect() protoreflect.Message {
	mi := &file_pds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortSummary.ProtoReflect.Descriptor instead.
func (*PortSummary) Descriptor() ([]byte, []int) {
	return file_pds_proto_rawDescGZIP(), []int{1}
}

func (x *PortSummary) GetPortCount() int32 {
	if x != nil {
		return x.PortCount
	}
	return 0
}

type IDMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *IDMsg) Reset() {
	*x = IDMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDMsg) ProtoMessage() {}

func (x *IDMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pds_proto_msgTypes[2]
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
	return file_pds_proto_rawDescGZIP(), []int{2}
}

func (x *IDMsg) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_pds_proto protoreflect.FileDescriptor

var file_pds_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2c, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x44, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0x4a, 0x0a,
	0x09, 0x50, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x05, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x1a,
	0x0c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x28, 0x01, 0x12,
	0x18, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x06, 0x2e, 0x49, 0x44, 0x4d,
	0x73, 0x67, 0x1a, 0x05, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pds_proto_rawDescOnce sync.Once
	file_pds_proto_rawDescData = file_pds_proto_rawDesc
)

func file_pds_proto_rawDescGZIP() []byte {
	file_pds_proto_rawDescOnce.Do(func() {
		file_pds_proto_rawDescData = protoimpl.X.CompressGZIP(file_pds_proto_rawDescData)
	})
	return file_pds_proto_rawDescData
}

var file_pds_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pds_proto_goTypes = []interface{}{
	(*Port)(nil),        // 0: Port
	(*PortSummary)(nil), // 1: PortSummary
	(*IDMsg)(nil),       // 2: IDMsg
}
var file_pds_proto_depIdxs = []int32{
	0, // 0: PDService.RecordPort:input_type -> Port
	2, // 1: PDService.GetByID:input_type -> IDMsg
	1, // 2: PDService.RecordPort:output_type -> PortSummary
	0, // 3: PDService.GetByID:output_type -> Port
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pds_proto_init() }
func file_pds_proto_init() {
	if File_pds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_pds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortSummary); i {
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
		file_pds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pds_proto_goTypes,
		DependencyIndexes: file_pds_proto_depIdxs,
		MessageInfos:      file_pds_proto_msgTypes,
	}.Build()
	File_pds_proto = out.File
	file_pds_proto_rawDesc = nil
	file_pds_proto_goTypes = nil
	file_pds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PDServiceClient is the client API for PDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PDServiceClient interface {
	RecordPort(ctx context.Context, opts ...grpc.CallOption) (PDService_RecordPortClient, error)
	GetByID(ctx context.Context, in *IDMsg, opts ...grpc.CallOption) (*Port, error)
}

type pDServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPDServiceClient(cc grpc.ClientConnInterface) PDServiceClient {
	return &pDServiceClient{cc}
}

func (c *pDServiceClient) RecordPort(ctx context.Context, opts ...grpc.CallOption) (PDService_RecordPortClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PDService_serviceDesc.Streams[0], "/PDService/RecordPort", opts...)
	if err != nil {
		return nil, err
	}
	x := &pDServiceRecordPortClient{stream}
	return x, nil
}

type PDService_RecordPortClient interface {
	Send(*Port) error
	CloseAndRecv() (*PortSummary, error)
	grpc.ClientStream
}

type pDServiceRecordPortClient struct {
	grpc.ClientStream
}

func (x *pDServiceRecordPortClient) Send(m *Port) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pDServiceRecordPortClient) CloseAndRecv() (*PortSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PortSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pDServiceClient) GetByID(ctx context.Context, in *IDMsg, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/PDService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PDServiceServer is the server API for PDService service.
type PDServiceServer interface {
	RecordPort(PDService_RecordPortServer) error
	GetByID(context.Context, *IDMsg) (*Port, error)
}

// UnimplementedPDServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPDServiceServer struct {
}

func (*UnimplementedPDServiceServer) RecordPort(PDService_RecordPortServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordPort not implemented")
}
func (*UnimplementedPDServiceServer) GetByID(context.Context, *IDMsg) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}

func RegisterPDServiceServer(s *grpc.Server, srv PDServiceServer) {
	s.RegisterService(&_PDService_serviceDesc, srv)
}

func _PDService_RecordPort_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PDServiceServer).RecordPort(&pDServiceRecordPortServer{stream})
}

type PDService_RecordPortServer interface {
	SendAndClose(*PortSummary) error
	Recv() (*Port, error)
	grpc.ServerStream
}

type pDServiceRecordPortServer struct {
	grpc.ServerStream
}

func (x *pDServiceRecordPortServer) SendAndClose(m *PortSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pDServiceRecordPortServer) Recv() (*Port, error) {
	m := new(Port)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PDService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PDService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDServiceServer).GetByID(ctx, req.(*IDMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _PDService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PDService",
	HandlerType: (*PDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _PDService_GetByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordPort",
			Handler:       _PDService_RecordPort_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pds.proto",
}
