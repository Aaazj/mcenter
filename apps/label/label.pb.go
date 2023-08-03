// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: mcenter/apps/label/pb/label.proto

package label

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

// A label selector operator is the set of operators that can be used in a selector requirement.
type OPERATOR int32

const (
	OPERATOR_IN     OPERATOR = 0
	OPERATOR_NOT_IN OPERATOR = 1
)

// Enum value maps for OPERATOR.
var (
	OPERATOR_name = map[int32]string{
		0: "IN",
		1: "NOT_IN",
	}
	OPERATOR_value = map[string]int32{
		"IN":     0,
		"NOT_IN": 1,
	}
)

func (x OPERATOR) Enum() *OPERATOR {
	p := new(OPERATOR)
	*p = x
	return p
}

func (x OPERATOR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OPERATOR) Descriptor() protoreflect.EnumDescriptor {
	return file_mcenter_apps_label_pb_label_proto_enumTypes[0].Descriptor()
}

func (OPERATOR) Type() protoreflect.EnumType {
	return &file_mcenter_apps_label_pb_label_proto_enumTypes[0]
}

func (x OPERATOR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OPERATOR.Descriptor instead.
func (OPERATOR) EnumDescriptor() ([]byte, []int) {
	return file_mcenter_apps_label_pb_label_proto_rawDescGZIP(), []int{0}
}

// A label selector is a label query over a set of resources. The result of matchLabels and
// matchExpressions are ANDed. An empty label selector matches all objects. A null
// label selector matches no objects.
// +structType=atomic
type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
	// map is equivalent to an element of matchExpressions, whose key field is "key", the
	// operator is "In", and the values array contains only "value". The requirements are ANDed.
	// @gotags: bson:"match_lablels" json:"match_lablels"
	MatchLablels map[string]string `protobuf:"bytes,1,rep,name=match_lablels,json=matchLablels,proto3" json:"match_lablels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"match_lablels"`
	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	// @gotags: bson:"match_lablels" json:"match_expressions"
	MatchExpressions []*LabelRequirement `protobuf:"bytes,2,rep,name=match_expressions,json=matchExpressions,proto3" json:"match_expressions" bson:"match_lablels"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_label_pb_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_label_pb_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_label_pb_label_proto_rawDescGZIP(), []int{0}
}

func (x *Selector) GetMatchLablels() map[string]string {
	if x != nil {
		return x.MatchLablels
	}
	return nil
}

func (x *Selector) GetMatchExpressions() []*LabelRequirement {
	if x != nil {
		return x.MatchExpressions
	}
	return nil
}

// A label selector requirement is a selector that contains values, a key, and an operator that
// relates the key and values.
type LabelRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is the label key that the selector applies to.
	// @gotags: bson:"match_lablels" json:"match_lablels"
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"match_lablels" bson:"match_lablels"`
	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	// @gotags: bson:"operator" json:"operator"
	Operator OPERATOR `protobuf:"varint,2,opt,name=operator,proto3,enum=mcenter.label.OPERATOR" json:"operator" bson:"operator"`
	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic
	// @gotags: bson:"values" json:"values"
	Values []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values" bson:"values"`
}

func (x *LabelRequirement) Reset() {
	*x = LabelRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_label_pb_label_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelRequirement) ProtoMessage() {}

func (x *LabelRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_label_pb_label_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelRequirement.ProtoReflect.Descriptor instead.
func (*LabelRequirement) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_label_pb_label_proto_rawDescGZIP(), []int{1}
}

func (x *LabelRequirement) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LabelRequirement) GetOperator() OPERATOR {
	if x != nil {
		return x.Operator
	}
	return OPERATOR_IN
}

func (x *LabelRequirement) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_mcenter_apps_label_pb_label_proto protoreflect.FileDescriptor

var file_mcenter_apps_label_pb_label_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x22, 0xe9, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x4e, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x6c, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x6c, 0x65, 0x6c, 0x73, 0x12,
	0x4c, 0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3f, 0x0a,
	0x11, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x6c, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x71,
	0x0a, 0x10, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x52,
	0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x2a, 0x1e, 0x0a, 0x08, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x12, 0x06, 0x0a,
	0x02, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x10,
	0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x61, 0x61, 0x7a, 0x6a, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_label_pb_label_proto_rawDescOnce sync.Once
	file_mcenter_apps_label_pb_label_proto_rawDescData = file_mcenter_apps_label_pb_label_proto_rawDesc
)

func file_mcenter_apps_label_pb_label_proto_rawDescGZIP() []byte {
	file_mcenter_apps_label_pb_label_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_label_pb_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_label_pb_label_proto_rawDescData)
	})
	return file_mcenter_apps_label_pb_label_proto_rawDescData
}

var file_mcenter_apps_label_pb_label_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcenter_apps_label_pb_label_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mcenter_apps_label_pb_label_proto_goTypes = []interface{}{
	(OPERATOR)(0),            // 0: mcenter.label.OPERATOR
	(*Selector)(nil),         // 1: mcenter.label.Selector
	(*LabelRequirement)(nil), // 2: mcenter.label.LabelRequirement
	nil,                      // 3: mcenter.label.Selector.MatchLablelsEntry
}
var file_mcenter_apps_label_pb_label_proto_depIdxs = []int32{
	3, // 0: mcenter.label.Selector.match_lablels:type_name -> mcenter.label.Selector.MatchLablelsEntry
	2, // 1: mcenter.label.Selector.match_expressions:type_name -> mcenter.label.LabelRequirement
	0, // 2: mcenter.label.LabelRequirement.operator:type_name -> mcenter.label.OPERATOR
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mcenter_apps_label_pb_label_proto_init() }
func file_mcenter_apps_label_pb_label_proto_init() {
	if File_mcenter_apps_label_pb_label_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_label_pb_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selector); i {
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
		file_mcenter_apps_label_pb_label_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelRequirement); i {
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
			RawDescriptor: file_mcenter_apps_label_pb_label_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_label_pb_label_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_label_pb_label_proto_depIdxs,
		EnumInfos:         file_mcenter_apps_label_pb_label_proto_enumTypes,
		MessageInfos:      file_mcenter_apps_label_pb_label_proto_msgTypes,
	}.Build()
	File_mcenter_apps_label_pb_label_proto = out.File
	file_mcenter_apps_label_pb_label_proto_rawDesc = nil
	file_mcenter_apps_label_pb_label_proto_goTypes = nil
	file_mcenter_apps_label_pb_label_proto_depIdxs = nil
}
