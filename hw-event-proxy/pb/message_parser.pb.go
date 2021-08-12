// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: message_parser.proto

package pb

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

type ParserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId   string   `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	MessageArgs []string `protobuf:"bytes,2,rep,name=message_args,json=messageArgs,proto3" json:"message_args,omitempty"`
}

func (x *ParserRequest) Reset() {
	*x = ParserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_parser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParserRequest) ProtoMessage() {}

func (x *ParserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_parser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParserRequest.ProtoReflect.Descriptor instead.
func (*ParserRequest) Descriptor() ([]byte, []int) {
	return file_message_parser_proto_rawDescGZIP(), []int{0}
}

func (x *ParserRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *ParserRequest) GetMessageArgs() []string {
	if x != nil {
		return x.MessageArgs
	}
	return nil
}

type ParserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Severity   string `protobuf:"bytes,2,opt,name=severity,proto3" json:"severity,omitempty"`
	Resolution string `protobuf:"bytes,3,opt,name=resolution,proto3" json:"resolution,omitempty"`
}

func (x *ParserResponse) Reset() {
	*x = ParserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_parser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParserResponse) ProtoMessage() {}

func (x *ParserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_parser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParserResponse.ProtoReflect.Descriptor instead.
func (*ParserResponse) Descriptor() ([]byte, []int) {
	return file_message_parser_proto_rawDescGZIP(), []int{1}
}

func (x *ParserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ParserResponse) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *ParserResponse) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

var File_message_parser_proto protoreflect.FileDescriptor

var file_message_parser_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x51, 0x0a, 0x0d, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x22, 0x66, 0x0a,
	0x0e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x41, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74, 0x2d, 0x63, 0x6e,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x68, 0x77, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_parser_proto_rawDescOnce sync.Once
	file_message_parser_proto_rawDescData = file_message_parser_proto_rawDesc
)

func file_message_parser_proto_rawDescGZIP() []byte {
	file_message_parser_proto_rawDescOnce.Do(func() {
		file_message_parser_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_parser_proto_rawDescData)
	})
	return file_message_parser_proto_rawDescData
}

var file_message_parser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_parser_proto_goTypes = []interface{}{
	(*ParserRequest)(nil),  // 0: pb.ParserRequest
	(*ParserResponse)(nil), // 1: pb.ParserResponse
}
var file_message_parser_proto_depIdxs = []int32{
	0, // 0: pb.MessageParser.Parse:input_type -> pb.ParserRequest
	1, // 1: pb.MessageParser.Parse:output_type -> pb.ParserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_parser_proto_init() }
func file_message_parser_proto_init() {
	if File_message_parser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_parser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParserRequest); i {
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
		file_message_parser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParserResponse); i {
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
			RawDescriptor: file_message_parser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_parser_proto_goTypes,
		DependencyIndexes: file_message_parser_proto_depIdxs,
		MessageInfos:      file_message_parser_proto_msgTypes,
	}.Build()
	File_message_parser_proto = out.File
	file_message_parser_proto_rawDesc = nil
	file_message_parser_proto_goTypes = nil
	file_message_parser_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageParserClient is the client API for MessageParser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageParserClient interface {
	Parse(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error)
}

type messageParserClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageParserClient(cc grpc.ClientConnInterface) MessageParserClient {
	return &messageParserClient{cc}
}

func (c *messageParserClient) Parse(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error) {
	out := new(ParserResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageParser/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageParserServer is the server API for MessageParser service.
type MessageParserServer interface {
	Parse(context.Context, *ParserRequest) (*ParserResponse, error)
}

// UnimplementedMessageParserServer can be embedded to have forward compatible implementations.
type UnimplementedMessageParserServer struct {
}

func (*UnimplementedMessageParserServer) Parse(context.Context, *ParserRequest) (*ParserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}

func RegisterMessageParserServer(s *grpc.Server, srv MessageParserServer) {
	s.RegisterService(&_MessageParser_serviceDesc, srv)
}

func _MessageParser_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageParserServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageParser/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageParserServer).Parse(ctx, req.(*ParserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageParser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageParser",
	HandlerType: (*MessageParserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _MessageParser_Parse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_parser.proto",
}