// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: dict/dict.proto

package dict

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

type TwoNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *TwoNumber) Reset() {
	*x = TwoNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_dict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwoNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwoNumber) ProtoMessage() {}

func (x *TwoNumber) ProtoReflect() protoreflect.Message {
	mi := &file_dict_dict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwoNumber.ProtoReflect.Descriptor instead.
func (*TwoNumber) Descriptor() ([]byte, []int) {
	return file_dict_dict_proto_rawDescGZIP(), []int{0}
}

func (x *TwoNumber) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *TwoNumber) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type ResultTwoNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultTwoNumber) Reset() {
	*x = ResultTwoNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_dict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultTwoNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultTwoNumber) ProtoMessage() {}

func (x *ResultTwoNumber) ProtoReflect() protoreflect.Message {
	mi := &file_dict_dict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultTwoNumber.ProtoReflect.Descriptor instead.
func (*ResultTwoNumber) Descriptor() ([]byte, []int) {
	return file_dict_dict_proto_rawDescGZIP(), []int{1}
}

func (x *ResultTwoNumber) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type SqrtNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *SqrtNum) Reset() {
	*x = SqrtNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_dict_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtNum) ProtoMessage() {}

func (x *SqrtNum) ProtoReflect() protoreflect.Message {
	mi := &file_dict_dict_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtNum.ProtoReflect.Descriptor instead.
func (*SqrtNum) Descriptor() ([]byte, []int) {
	return file_dict_dict_proto_rawDescGZIP(), []int{2}
}

func (x *SqrtNum) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_dict_dict_proto protoreflect.FileDescriptor

var file_dict_dict_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x64, 0x69, 0x63, 0x74, 0x22, 0x27, 0x0a, 0x09, 0x74, 0x77, 0x6f, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62,
	0x22, 0x29, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1b, 0x0a, 0x07, 0x73,
	0x71, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0xe4, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x6c,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x0f, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x74, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x1a, 0x15, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x77,
	0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x53, 0x75, 0x62,
	0x12, 0x0f, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x74, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54,
	0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04, 0x4d, 0x75,
	0x6c, 0x74, 0x12, 0x0f, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x74, 0x77, 0x6f, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03,
	0x44, 0x69, 0x76, 0x12, 0x0f, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x74, 0x77, 0x6f, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2e, 0x0a,
	0x04, 0x53, 0x71, 0x72, 0x74, 0x12, 0x0d, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x73, 0x71, 0x72,
	0x74, 0x4e, 0x75, 0x6d, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x03, 0x50, 0x6f, 0x77, 0x12, 0x0f, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x74, 0x77, 0x6f, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2f,
	0x0a, 0x03, 0x4d, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x74, 0x77, 0x6f,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_dict_dict_proto_rawDescOnce sync.Once
	file_dict_dict_proto_rawDescData = file_dict_dict_proto_rawDesc
)

func file_dict_dict_proto_rawDescGZIP() []byte {
	file_dict_dict_proto_rawDescOnce.Do(func() {
		file_dict_dict_proto_rawDescData = protoimpl.X.CompressGZIP(file_dict_dict_proto_rawDescData)
	})
	return file_dict_dict_proto_rawDescData
}

var file_dict_dict_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dict_dict_proto_goTypes = []interface{}{
	(*TwoNumber)(nil),       // 0: dict.twoNumber
	(*ResultTwoNumber)(nil), // 1: dict.resultTwoNumber
	(*SqrtNum)(nil),         // 2: dict.sqrtNum
}
var file_dict_dict_proto_depIdxs = []int32{
	0, // 0: dict.CalcService.Add:input_type -> dict.twoNumber
	0, // 1: dict.CalcService.Sub:input_type -> dict.twoNumber
	0, // 2: dict.CalcService.Mult:input_type -> dict.twoNumber
	0, // 3: dict.CalcService.Div:input_type -> dict.twoNumber
	2, // 4: dict.CalcService.Sqrt:input_type -> dict.sqrtNum
	0, // 5: dict.CalcService.Pow:input_type -> dict.twoNumber
	0, // 6: dict.CalcService.Min:input_type -> dict.twoNumber
	1, // 7: dict.CalcService.Add:output_type -> dict.resultTwoNumber
	1, // 8: dict.CalcService.Sub:output_type -> dict.resultTwoNumber
	1, // 9: dict.CalcService.Mult:output_type -> dict.resultTwoNumber
	1, // 10: dict.CalcService.Div:output_type -> dict.resultTwoNumber
	1, // 11: dict.CalcService.Sqrt:output_type -> dict.resultTwoNumber
	1, // 12: dict.CalcService.Pow:output_type -> dict.resultTwoNumber
	1, // 13: dict.CalcService.Min:output_type -> dict.resultTwoNumber
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dict_dict_proto_init() }
func file_dict_dict_proto_init() {
	if File_dict_dict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dict_dict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwoNumber); i {
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
		file_dict_dict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultTwoNumber); i {
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
		file_dict_dict_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtNum); i {
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
			RawDescriptor: file_dict_dict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dict_dict_proto_goTypes,
		DependencyIndexes: file_dict_dict_proto_depIdxs,
		MessageInfos:      file_dict_dict_proto_msgTypes,
	}.Build()
	File_dict_dict_proto = out.File
	file_dict_dict_proto_rawDesc = nil
	file_dict_dict_proto_goTypes = nil
	file_dict_dict_proto_depIdxs = nil
}
