// Unless explicitly stated otherwise all files in this repository are licensed under the Apache License 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2020 Datadog, Inc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: ddsketch.proto

package sketchpb

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

type IndexMapping_Interpolation int32

const (
	IndexMapping_NONE      IndexMapping_Interpolation = 0
	IndexMapping_LINEAR    IndexMapping_Interpolation = 1
	IndexMapping_QUADRATIC IndexMapping_Interpolation = 2
	IndexMapping_CUBIC     IndexMapping_Interpolation = 3
)

// Enum value maps for IndexMapping_Interpolation.
var (
	IndexMapping_Interpolation_name = map[int32]string{
		0: "NONE",
		1: "LINEAR",
		2: "QUADRATIC",
		3: "CUBIC",
	}
	IndexMapping_Interpolation_value = map[string]int32{
		"NONE":      0,
		"LINEAR":    1,
		"QUADRATIC": 2,
		"CUBIC":     3,
	}
)

func (x IndexMapping_Interpolation) Enum() *IndexMapping_Interpolation {
	p := new(IndexMapping_Interpolation)
	*p = x
	return p
}

func (x IndexMapping_Interpolation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexMapping_Interpolation) Descriptor() protoreflect.EnumDescriptor {
	return file_ddsketch_proto_enumTypes[0].Descriptor()
}

func (IndexMapping_Interpolation) Type() protoreflect.EnumType {
	return &file_ddsketch_proto_enumTypes[0]
}

func (x IndexMapping_Interpolation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexMapping_Interpolation.Descriptor instead.
func (IndexMapping_Interpolation) EnumDescriptor() ([]byte, []int) {
	return file_ddsketch_proto_rawDescGZIP(), []int{1, 0}
}

// A DDSketch is essentially a histogram that partitions the range of positive values into an infinite number of
// indexed bins whose size grows exponentially. It keeps track of the number of values (or possibly floating-point
// weights) added to each bin. Negative values are partitioned like positive values, symmetrically to zero.
// The value zero as well as its close neighborhood that would be mapped to extreme bin indexes is mapped to a specific
// counter.
type DDSketch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mapping between positive values and the bin indexes they belong to.
	Mapping *IndexMapping `protobuf:"bytes,1,opt,name=mapping,proto3" json:"mapping,omitempty"`
	// The store for keeping track of positive values.
	PositiveValues *Store `protobuf:"bytes,2,opt,name=positiveValues,proto3" json:"positiveValues,omitempty"`
	// The store for keeping track of negative values. A negative value v is mapped using its positive opposite -v.
	NegativeValues *Store `protobuf:"bytes,3,opt,name=negativeValues,proto3" json:"negativeValues,omitempty"`
	// The count for the value zero and its close neighborhood (whose width depends on the mapping).
	ZeroCount float64 `protobuf:"fixed64,4,opt,name=zeroCount,proto3" json:"zeroCount,omitempty"`
}

func (x *DDSketch) Reset() {
	*x = DDSketch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddsketch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DDSketch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DDSketch) ProtoMessage() {}

func (x *DDSketch) ProtoReflect() protoreflect.Message {
	mi := &file_ddsketch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DDSketch.ProtoReflect.Descriptor instead.
func (*DDSketch) Descriptor() ([]byte, []int) {
	return file_ddsketch_proto_rawDescGZIP(), []int{0}
}

func (x *DDSketch) GetMapping() *IndexMapping {
	if x != nil {
		return x.Mapping
	}
	return nil
}

func (x *DDSketch) GetPositiveValues() *Store {
	if x != nil {
		return x.PositiveValues
	}
	return nil
}

func (x *DDSketch) GetNegativeValues() *Store {
	if x != nil {
		return x.NegativeValues
	}
	return nil
}

func (x *DDSketch) GetZeroCount() float64 {
	if x != nil {
		return x.ZeroCount
	}
	return 0
}

// How to map positive values to the bins they belong to.
type IndexMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The gamma parameter of the mapping, such that bin index that a value v belongs to is roughly equal to
	// log(v)/log(gamma).
	Gamma float64 `protobuf:"fixed64,1,opt,name=gamma,proto3" json:"gamma,omitempty"`
	// An offset that can be used to shift all bin indexes.
	IndexOffset float64 `protobuf:"fixed64,2,opt,name=indexOffset,proto3" json:"indexOffset,omitempty"`
	// To speed up the computation of the index a value belongs to, the computation of the log may be approximated using
	// the fact that the log to the base 2 of powers of 2 can be computed at a low cost from the binary representation of
	// the input value. Other values can be approximated by interpolating between successive powers of 2 (linearly,
	// quadratically or cubically).
	// NONE means that the log is to be computed exactly (no interpolation).
	Interpolation IndexMapping_Interpolation `protobuf:"varint,3,opt,name=interpolation,proto3,enum=IndexMapping_Interpolation" json:"interpolation,omitempty"`
}

func (x *IndexMapping) Reset() {
	*x = IndexMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddsketch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexMapping) ProtoMessage() {}

func (x *IndexMapping) ProtoReflect() protoreflect.Message {
	mi := &file_ddsketch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexMapping.ProtoReflect.Descriptor instead.
func (*IndexMapping) Descriptor() ([]byte, []int) {
	return file_ddsketch_proto_rawDescGZIP(), []int{1}
}

func (x *IndexMapping) GetGamma() float64 {
	if x != nil {
		return x.Gamma
	}
	return 0
}

func (x *IndexMapping) GetIndexOffset() float64 {
	if x != nil {
		return x.IndexOffset
	}
	return 0
}

func (x *IndexMapping) GetInterpolation() IndexMapping_Interpolation {
	if x != nil {
		return x.Interpolation
	}
	return IndexMapping_NONE
}

// A Store maps bin indexes to their respective counts.
// Counts can be encoded sparsely using binCounts, but also in a contiguous way using contiguousBinCounts and
// contiguousBinIndexOffset. Given that non-empty bins are in practice usually contiguous or close to one another, the
// latter contiguous encoding method is usually more efficient than the sparse one.
// Both encoding methods can be used conjointly. If a bin appears in both the sparse and the contiguous encodings, its
// count value is the sum of the counts in each encodings.
type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bin counts, encoded sparsely.
	BinCounts map[int32]float64 `protobuf:"bytes,1,rep,name=binCounts,proto3" json:"binCounts,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// The bin counts, encoded contiguously. The values of contiguousBinCounts are the counts for the bins of indexes
	// o, o+1, o+2, etc., where o is contiguousBinIndexOffset.
	ContiguousBinCounts      []float64 `protobuf:"fixed64,2,rep,packed,name=contiguousBinCounts,proto3" json:"contiguousBinCounts,omitempty"`
	ContiguousBinIndexOffset int32     `protobuf:"zigzag32,3,opt,name=contiguousBinIndexOffset,proto3" json:"contiguousBinIndexOffset,omitempty"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddsketch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_ddsketch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_ddsketch_proto_rawDescGZIP(), []int{2}
}

func (x *Store) GetBinCounts() map[int32]float64 {
	if x != nil {
		return x.BinCounts
	}
	return nil
}

func (x *Store) GetContiguousBinCounts() []float64 {
	if x != nil {
		return x.ContiguousBinCounts
	}
	return nil
}

func (x *Store) GetContiguousBinIndexOffset() int32 {
	if x != nil {
		return x.ContiguousBinIndexOffset
	}
	return 0
}

var File_ddsketch_proto protoreflect.FileDescriptor

var file_ddsketch_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x64, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb1, 0x01, 0x0a, 0x08, 0x44, 0x44, 0x53, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x12, 0x27, 0x0a,
	0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x0e, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x0e, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x7a, 0x65, 0x72, 0x6f, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x7a, 0x65, 0x72, 0x6f, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x41, 0x0a,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3f, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c,
	0x49, 0x4e, 0x45, 0x41, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x51, 0x55, 0x41, 0x44, 0x52,
	0x41, 0x54, 0x49, 0x43, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x55, 0x42, 0x49, 0x43, 0x10,
	0x03, 0x22, 0xec, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x62,
	0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x62, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x34, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x75, 0x6f, 0x75, 0x73, 0x42, 0x69,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10,
	0x01, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x75, 0x6f, 0x75, 0x73, 0x42, 0x69, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67,
	0x75, 0x6f, 0x75, 0x73, 0x42, 0x69, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67,
	0x75, 0x6f, 0x75, 0x73, 0x42, 0x69, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x1a, 0x3c, 0x0a, 0x0e, 0x42, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44,
	0x61, 0x74, 0x61, 0x44, 0x6f, 0x67, 0x2f, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2d,
	0x67, 0x6f, 0x2f, 0x64, 0x64, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x2f, 0x70, 0x62, 0x2f, 0x73,
	0x6b, 0x65, 0x74, 0x63, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ddsketch_proto_rawDescOnce sync.Once
	file_ddsketch_proto_rawDescData = file_ddsketch_proto_rawDesc
)

func file_ddsketch_proto_rawDescGZIP() []byte {
	file_ddsketch_proto_rawDescOnce.Do(func() {
		file_ddsketch_proto_rawDescData = protoimpl.X.CompressGZIP(file_ddsketch_proto_rawDescData)
	})
	return file_ddsketch_proto_rawDescData
}

var file_ddsketch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ddsketch_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ddsketch_proto_goTypes = []interface{}{
	(IndexMapping_Interpolation)(0), // 0: IndexMapping.Interpolation
	(*DDSketch)(nil),                // 1: DDSketch
	(*IndexMapping)(nil),            // 2: IndexMapping
	(*Store)(nil),                   // 3: Store
	nil,                             // 4: Store.BinCountsEntry
}
var file_ddsketch_proto_depIdxs = []int32{
	2, // 0: DDSketch.mapping:type_name -> IndexMapping
	3, // 1: DDSketch.positiveValues:type_name -> Store
	3, // 2: DDSketch.negativeValues:type_name -> Store
	0, // 3: IndexMapping.interpolation:type_name -> IndexMapping.Interpolation
	4, // 4: Store.binCounts:type_name -> Store.BinCountsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ddsketch_proto_init() }
func file_ddsketch_proto_init() {
	if File_ddsketch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ddsketch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DDSketch); i {
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
		file_ddsketch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexMapping); i {
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
		file_ddsketch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
			RawDescriptor: file_ddsketch_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ddsketch_proto_goTypes,
		DependencyIndexes: file_ddsketch_proto_depIdxs,
		EnumInfos:         file_ddsketch_proto_enumTypes,
		MessageInfos:      file_ddsketch_proto_msgTypes,
	}.Build()
	File_ddsketch_proto = out.File
	file_ddsketch_proto_rawDesc = nil
	file_ddsketch_proto_goTypes = nil
	file_ddsketch_proto_depIdxs = nil
}
