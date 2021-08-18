// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: flattenproto_test.proto

package flattenproto

import (
	proto "github.com/golang/protobuf/proto"
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

type EE int32

const (
	EE_EE_NONE   EE = 0
	EE_EE_FIRST  EE = 1
	EE_EE_SECOND EE = 2
)

// Enum value maps for EE.
var (
	EE_name = map[int32]string{
		0: "EE_NONE",
		1: "EE_FIRST",
		2: "EE_SECOND",
	}
	EE_value = map[string]int32{
		"EE_NONE":   0,
		"EE_FIRST":  1,
		"EE_SECOND": 2,
	}
)

func (x EE) Enum() *EE {
	p := new(EE)
	*p = x
	return p
}

func (x EE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EE) Descriptor() protoreflect.EnumDescriptor {
	return file_flattenproto_test_proto_enumTypes[0].Descriptor()
}

func (EE) Type() protoreflect.EnumType {
	return &file_flattenproto_test_proto_enumTypes[0]
}

func (x EE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EE.Descriptor instead.
func (EE) EnumDescriptor() ([]byte, []int) {
	return file_flattenproto_test_proto_rawDescGZIP(), []int{0}
}

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X2 string `protobuf:"bytes,2,opt,name=x2,proto3" json:"x2,omitempty"`
	X3 int32  `protobuf:"varint,3,opt,name=x3,proto3" json:"x3,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flattenproto_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_flattenproto_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_flattenproto_test_proto_rawDescGZIP(), []int{0}
}

func (x *Bar) GetX2() string {
	if x != nil {
		return x.X2
	}
	return ""
}

func (x *Bar) GetX3() int32 {
	if x != nil {
		return x.X3
	}
	return 0
}

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X1    []string          `protobuf:"bytes,1,rep,name=x1,proto3" json:"x1,omitempty"`
	X2    string            `protobuf:"bytes,2,opt,name=x2,proto3" json:"x2,omitempty"`
	X3    int32             `protobuf:"varint,3,opt,name=x3,proto3" json:"x3,omitempty"`
	X4    map[string]string `protobuf:"bytes,4,rep,name=x4,proto3" json:"x4,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	X5Xyz string            `protobuf:"bytes,5,opt,name=x5_xyz,json=x5Xyz,proto3" json:"x5_xyz,omitempty"`
	Ee    EE                `protobuf:"varint,6,opt,name=ee,proto3,enum=flattenproto.EE" json:"ee,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flattenproto_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_flattenproto_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_flattenproto_test_proto_rawDescGZIP(), []int{1}
}

func (x *Foo) GetX1() []string {
	if x != nil {
		return x.X1
	}
	return nil
}

func (x *Foo) GetX2() string {
	if x != nil {
		return x.X2
	}
	return ""
}

func (x *Foo) GetX3() int32 {
	if x != nil {
		return x.X3
	}
	return 0
}

func (x *Foo) GetX4() map[string]string {
	if x != nil {
		return x.X4
	}
	return nil
}

func (x *Foo) GetX5Xyz() string {
	if x != nil {
		return x.X5Xyz
	}
	return ""
}

func (x *Foo) GetEe() EE {
	if x != nil {
		return x.Ee
	}
	return EE_EE_NONE
}

var File_flattenproto_test_proto protoreflect.FileDescriptor

var file_flattenproto_test_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x6c, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x66, 0x6c, 0x61, 0x74, 0x74,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x78, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x78, 0x32, 0x12, 0x0e,
	0x0a, 0x02, 0x78, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x78, 0x33, 0x22, 0xd0,
	0x01, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x78, 0x31, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x02, 0x78, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x78, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x78, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x78, 0x33, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x78, 0x33, 0x12, 0x29, 0x0a, 0x02, 0x78, 0x34, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x6f, 0x6f, 0x2e, 0x58, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x78,
	0x34, 0x12, 0x15, 0x0a, 0x06, 0x78, 0x35, 0x5f, 0x78, 0x79, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x78, 0x35, 0x58, 0x79, 0x7a, 0x12, 0x20, 0x0a, 0x02, 0x65, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x66, 0x6c, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x45, 0x52, 0x02, 0x65, 0x65, 0x1a, 0x35, 0x0a, 0x07, 0x58, 0x34,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x2a, 0x2e, 0x0a, 0x02, 0x45, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10,
	0x02, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x66, 0x6c, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flattenproto_test_proto_rawDescOnce sync.Once
	file_flattenproto_test_proto_rawDescData = file_flattenproto_test_proto_rawDesc
)

func file_flattenproto_test_proto_rawDescGZIP() []byte {
	file_flattenproto_test_proto_rawDescOnce.Do(func() {
		file_flattenproto_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_flattenproto_test_proto_rawDescData)
	})
	return file_flattenproto_test_proto_rawDescData
}

var file_flattenproto_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flattenproto_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_flattenproto_test_proto_goTypes = []interface{}{
	(EE)(0),     // 0: flattenproto.EE
	(*Bar)(nil), // 1: flattenproto.Bar
	(*Foo)(nil), // 2: flattenproto.Foo
	nil,         // 3: flattenproto.Foo.X4Entry
}
var file_flattenproto_test_proto_depIdxs = []int32{
	3, // 0: flattenproto.Foo.x4:type_name -> flattenproto.Foo.X4Entry
	0, // 1: flattenproto.Foo.ee:type_name -> flattenproto.EE
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_flattenproto_test_proto_init() }
func file_flattenproto_test_proto_init() {
	if File_flattenproto_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flattenproto_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
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
		file_flattenproto_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
			RawDescriptor: file_flattenproto_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flattenproto_test_proto_goTypes,
		DependencyIndexes: file_flattenproto_test_proto_depIdxs,
		EnumInfos:         file_flattenproto_test_proto_enumTypes,
		MessageInfos:      file_flattenproto_test_proto_msgTypes,
	}.Build()
	File_flattenproto_test_proto = out.File
	file_flattenproto_test_proto_rawDesc = nil
	file_flattenproto_test_proto_goTypes = nil
	file_flattenproto_test_proto_depIdxs = nil
}
