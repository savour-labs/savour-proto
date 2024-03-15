// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: savour_rpc/airdrop.proto

package airdrop

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

type DppLinkPointsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Address       string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DppLinkPointsReq) Reset() {
	*x = DppLinkPointsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savour_rpc_airdrop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DppLinkPointsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DppLinkPointsReq) ProtoMessage() {}

func (x *DppLinkPointsReq) ProtoReflect() protoreflect.Message {
	mi := &file_savour_rpc_airdrop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DppLinkPointsReq.ProtoReflect.Descriptor instead.
func (*DppLinkPointsReq) Descriptor() ([]byte, []int) {
	return file_savour_rpc_airdrop_proto_rawDescGZIP(), []int{0}
}

func (x *DppLinkPointsReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *DppLinkPointsReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DppLinkPointsReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DppLinkPointsRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DppLinkPointsRep) Reset() {
	*x = DppLinkPointsRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savour_rpc_airdrop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DppLinkPointsRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DppLinkPointsRep) ProtoMessage() {}

func (x *DppLinkPointsRep) ProtoReflect() protoreflect.Message {
	mi := &file_savour_rpc_airdrop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DppLinkPointsRep.ProtoReflect.Descriptor instead.
func (*DppLinkPointsRep) Descriptor() ([]byte, []int) {
	return file_savour_rpc_airdrop_proto_rawDescGZIP(), []int{1}
}

func (x *DppLinkPointsRep) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DppLinkPointsRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_savour_rpc_airdrop_proto protoreflect.FileDescriptor

var file_savour_rpc_airdrop_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x69, 0x72,
	0x64, 0x72, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x22, 0x67, 0x0a, 0x10, 0x44, 0x70, 0x70, 0x4c, 0x69,
	0x6e, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x38, 0x0a, 0x10, 0x44, 0x70, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x87, 0x01, 0x0a, 0x0e, 0x41,
	0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a,
	0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x70, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x69, 0x72, 0x64, 0x72,
	0x6f, 0x70, 0x2e, 0x44, 0x70, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73,
	0x61, 0x76, 0x6f, 0x75, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f,
	0x70, 0x2e, 0x44, 0x70, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x70, 0x22, 0x00, 0x42, 0x27, 0x0a, 0x14, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x61,
	0x76, 0x6f, 0x75, 0x72, 0x2e, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x5a, 0x0f, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_savour_rpc_airdrop_proto_rawDescOnce sync.Once
	file_savour_rpc_airdrop_proto_rawDescData = file_savour_rpc_airdrop_proto_rawDesc
)

func file_savour_rpc_airdrop_proto_rawDescGZIP() []byte {
	file_savour_rpc_airdrop_proto_rawDescOnce.Do(func() {
		file_savour_rpc_airdrop_proto_rawDescData = protoimpl.X.CompressGZIP(file_savour_rpc_airdrop_proto_rawDescData)
	})
	return file_savour_rpc_airdrop_proto_rawDescData
}

var file_savour_rpc_airdrop_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_savour_rpc_airdrop_proto_goTypes = []interface{}{
	(*DppLinkPointsReq)(nil), // 0: services.savour_rpc.airdrop.DppLinkPointsReq
	(*DppLinkPointsRep)(nil), // 1: services.savour_rpc.airdrop.DppLinkPointsRep
}
var file_savour_rpc_airdrop_proto_depIdxs = []int32{
	0, // 0: services.savour_rpc.airdrop.AirdropService.submitDppLinkPoints:input_type -> services.savour_rpc.airdrop.DppLinkPointsReq
	1, // 1: services.savour_rpc.airdrop.AirdropService.submitDppLinkPoints:output_type -> services.savour_rpc.airdrop.DppLinkPointsRep
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_savour_rpc_airdrop_proto_init() }
func file_savour_rpc_airdrop_proto_init() {
	if File_savour_rpc_airdrop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_savour_rpc_airdrop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DppLinkPointsReq); i {
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
		file_savour_rpc_airdrop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DppLinkPointsRep); i {
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
			RawDescriptor: file_savour_rpc_airdrop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_savour_rpc_airdrop_proto_goTypes,
		DependencyIndexes: file_savour_rpc_airdrop_proto_depIdxs,
		MessageInfos:      file_savour_rpc_airdrop_proto_msgTypes,
	}.Build()
	File_savour_rpc_airdrop_proto = out.File
	file_savour_rpc_airdrop_proto_rawDesc = nil
	file_savour_rpc_airdrop_proto_goTypes = nil
	file_savour_rpc_airdrop_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AirdropServiceClient is the client API for AirdropService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AirdropServiceClient interface {
	SubmitDppLinkPoints(ctx context.Context, in *DppLinkPointsReq, opts ...grpc.CallOption) (*DppLinkPointsRep, error)
}

type airdropServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAirdropServiceClient(cc grpc.ClientConnInterface) AirdropServiceClient {
	return &airdropServiceClient{cc}
}

func (c *airdropServiceClient) SubmitDppLinkPoints(ctx context.Context, in *DppLinkPointsReq, opts ...grpc.CallOption) (*DppLinkPointsRep, error) {
	out := new(DppLinkPointsRep)
	err := c.cc.Invoke(ctx, "/services.savour_rpc.airdrop.AirdropService/submitDppLinkPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AirdropServiceServer is the server API for AirdropService service.
type AirdropServiceServer interface {
	SubmitDppLinkPoints(context.Context, *DppLinkPointsReq) (*DppLinkPointsRep, error)
}

// UnimplementedAirdropServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAirdropServiceServer struct {
}

func (*UnimplementedAirdropServiceServer) SubmitDppLinkPoints(context.Context, *DppLinkPointsReq) (*DppLinkPointsRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDppLinkPoints not implemented")
}

func RegisterAirdropServiceServer(s *grpc.Server, srv AirdropServiceServer) {
	s.RegisterService(&_AirdropService_serviceDesc, srv)
}

func _AirdropService_SubmitDppLinkPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DppLinkPointsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirdropServiceServer).SubmitDppLinkPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.savour_rpc.airdrop.AirdropService/SubmitDppLinkPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirdropServiceServer).SubmitDppLinkPoints(ctx, req.(*DppLinkPointsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AirdropService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.savour_rpc.airdrop.AirdropService",
	HandlerType: (*AirdropServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "submitDppLinkPoints",
			Handler:    _AirdropService_SubmitDppLinkPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "savour_rpc/airdrop.proto",
}
