// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: filter_message.proto

package go_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Filter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MaxPriceUsd   float64                `protobuf:"fixed64,1,opt,name=max_price_usd,json=maxPriceUsd,proto3" json:"max_price_usd,omitempty"`
	MinCpuCores   uint32                 `protobuf:"varint,2,opt,name=min_cpu_cores,json=minCpuCores,proto3" json:"min_cpu_cores,omitempty"`
	MinCpuGhz     float64                `protobuf:"fixed64,3,opt,name=min_cpu_ghz,json=minCpuGhz,proto3" json:"min_cpu_ghz,omitempty"`
	MinRam        *Memory                `protobuf:"bytes,4,opt,name=min_ram,json=minRam,proto3" json:"min_ram,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_filter_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_filter_message_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetMaxPriceUsd() float64 {
	if x != nil {
		return x.MaxPriceUsd
	}
	return 0
}

func (x *Filter) GetMinCpuCores() uint32 {
	if x != nil {
		return x.MinCpuCores
	}
	return 0
}

func (x *Filter) GetMinCpuGhz() float64 {
	if x != nil {
		return x.MinCpuGhz
	}
	return 0
}

func (x *Filter) GetMinRam() *Memory {
	if x != nil {
		return x.MinRam
	}
	return nil
}

var File_filter_message_proto protoreflect.FileDescriptor

var file_filter_message_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x73, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6d,
	0x69, 0x6e, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x43, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x67, 0x68, 0x7a, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x43, 0x70, 0x75, 0x47, 0x68, 0x7a, 0x12,
	0x20, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x52, 0x61,
	0x6d, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_filter_message_proto_rawDescOnce sync.Once
	file_filter_message_proto_rawDescData []byte
)

func file_filter_message_proto_rawDescGZIP() []byte {
	file_filter_message_proto_rawDescOnce.Do(func() {
		file_filter_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_filter_message_proto_rawDesc), len(file_filter_message_proto_rawDesc)))
	})
	return file_filter_message_proto_rawDescData
}

var file_filter_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_filter_message_proto_goTypes = []any{
	(*Filter)(nil), // 0: Filter
	(*Memory)(nil), // 1: Memory
}
var file_filter_message_proto_depIdxs = []int32{
	1, // 0: Filter.min_ram:type_name -> Memory
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_filter_message_proto_init() }
func file_filter_message_proto_init() {
	if File_filter_message_proto != nil {
		return
	}
	file_memory_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_filter_message_proto_rawDesc), len(file_filter_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filter_message_proto_goTypes,
		DependencyIndexes: file_filter_message_proto_depIdxs,
		MessageInfos:      file_filter_message_proto_msgTypes,
	}.Build()
	File_filter_message_proto = out.File
	file_filter_message_proto_goTypes = nil
	file_filter_message_proto_depIdxs = nil
}
