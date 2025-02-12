// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: api/grpc/v1/buzz/buzz.proto

package buzz

import (
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

type SearchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int32                  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage int32                  `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_api_grpc_v1_buzz_buzz_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_buzz_buzz_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_buzz_buzz_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchRequest) GetResultPerPage() int32 {
	if x != nil {
		return x.ResultPerPage
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []string               `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	mi := &file_api_grpc_v1_buzz_buzz_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_buzz_buzz_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_buzz_buzz_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetResults() []string {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_api_grpc_v1_buzz_buzz_proto protoreflect.FileDescriptor

var file_api_grpc_v1_buzz_buzz_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75,
	0x7a, 0x7a, 0x2f, 0x62, 0x75, 0x7a, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62,
	0x75, 0x7a, 0x7a, 0x5f, 0x70, 0x62, 0x22, 0x6e, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x32, 0x4c, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x2e,
	0x62, 0x75, 0x7a, 0x7a, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x75, 0x7a, 0x7a, 0x5f, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x6f, 0x6d, 0x63, 0x61, 0x6e, 0x68, 0x61, 0x6d, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x7a, 0x7a, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_v1_buzz_buzz_proto_rawDescOnce sync.Once
	file_api_grpc_v1_buzz_buzz_proto_rawDescData = file_api_grpc_v1_buzz_buzz_proto_rawDesc
)

func file_api_grpc_v1_buzz_buzz_proto_rawDescGZIP() []byte {
	file_api_grpc_v1_buzz_buzz_proto_rawDescOnce.Do(func() {
		file_api_grpc_v1_buzz_buzz_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_v1_buzz_buzz_proto_rawDescData)
	})
	return file_api_grpc_v1_buzz_buzz_proto_rawDescData
}

var file_api_grpc_v1_buzz_buzz_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grpc_v1_buzz_buzz_proto_goTypes = []any{
	(*SearchRequest)(nil),  // 0: buzz_pb.SearchRequest
	(*SearchResponse)(nil), // 1: buzz_pb.SearchResponse
}
var file_api_grpc_v1_buzz_buzz_proto_depIdxs = []int32{
	0, // 0: buzz_pb.SearchService.Search:input_type -> buzz_pb.SearchRequest
	1, // 1: buzz_pb.SearchService.Search:output_type -> buzz_pb.SearchResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_v1_buzz_buzz_proto_init() }
func file_api_grpc_v1_buzz_buzz_proto_init() {
	if File_api_grpc_v1_buzz_buzz_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_grpc_v1_buzz_buzz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_v1_buzz_buzz_proto_goTypes,
		DependencyIndexes: file_api_grpc_v1_buzz_buzz_proto_depIdxs,
		MessageInfos:      file_api_grpc_v1_buzz_buzz_proto_msgTypes,
	}.Build()
	File_api_grpc_v1_buzz_buzz_proto = out.File
	file_api_grpc_v1_buzz_buzz_proto_rawDesc = nil
	file_api_grpc_v1_buzz_buzz_proto_goTypes = nil
	file_api_grpc_v1_buzz_buzz_proto_depIdxs = nil
}
