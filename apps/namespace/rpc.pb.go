// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: mcenter/apps/namespace/pb/rpc.proto

package namespace

import (
	request "github.com/infraboard/mcube/http/request"
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

// QueryNamespaceRequest 查询应用列表
type QueryNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 空间的名称
	// @gotags: json:"name"
	Name []string `protobuf:"bytes,3,rep,name=name,proto3" json:"name"`
	// 命名空间的id列表
	// @gotags: json:"ids"
	Ids []string `protobuf:"bytes,4,rep,name=ids,proto3" json:"ids"`
	// 是否查询子空间
	// @gotags: json:"with_sub"
	WithSub bool `protobuf:"varint,5,opt,name=with_sub,json=withSub,proto3" json:"with_sub"`
	// 用户加入的空间
	// @gotags: json:"username"
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username"`
}

func (x *QueryNamespaceRequest) Reset() {
	*x = QueryNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryNamespaceRequest) ProtoMessage() {}

func (x *QueryNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryNamespaceRequest.ProtoReflect.Descriptor instead.
func (*QueryNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryNamespaceRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryNamespaceRequest) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *QueryNamespaceRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *QueryNamespaceRequest) GetWithSub() bool {
	if x != nil {
		return x.WithSub
	}
	return false
}

func (x *QueryNamespaceRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// DescriptNamespaceRequest 查询应用详情
type DescriptNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *DescriptNamespaceRequest) Reset() {
	*x = DescriptNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptNamespaceRequest) ProtoMessage() {}

func (x *DescriptNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DescriptNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescriptNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DescriptNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// DeleteNamespaceRequest todo
type DeleteNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *DeleteNamespaceRequest) Reset() {
	*x = DeleteNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRequest) ProtoMessage() {}

func (x *DeleteNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DeleteNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_mcenter_apps_namespace_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_namespace_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x1a, 0x0f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc4, 0x01, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x75, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x77, 0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0xee, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x71, 0x0a, 0x0e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x33,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x74, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_namespace_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_namespace_pb_rpc_proto_rawDescData = file_mcenter_apps_namespace_pb_rpc_proto_rawDesc
)

func file_mcenter_apps_namespace_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_namespace_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_namespace_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_namespace_pb_rpc_proto_rawDescData)
	})
	return file_mcenter_apps_namespace_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_namespace_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mcenter_apps_namespace_pb_rpc_proto_goTypes = []interface{}{
	(*QueryNamespaceRequest)(nil),    // 0: infraboard.mcenter.namespace.QueryNamespaceRequest
	(*DescriptNamespaceRequest)(nil), // 1: infraboard.mcenter.namespace.DescriptNamespaceRequest
	(*DeleteNamespaceRequest)(nil),   // 2: infraboard.mcenter.namespace.DeleteNamespaceRequest
	(*request.PageRequest)(nil),      // 3: infraboard.mcube.page.PageRequest
	(*NamespaceSet)(nil),             // 4: infraboard.mcenter.namespace.NamespaceSet
	(*Namespace)(nil),                // 5: infraboard.mcenter.namespace.Namespace
}
var file_mcenter_apps_namespace_pb_rpc_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.namespace.QueryNamespaceRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 1: infraboard.mcenter.namespace.RPC.QueryNamespace:input_type -> infraboard.mcenter.namespace.QueryNamespaceRequest
	1, // 2: infraboard.mcenter.namespace.RPC.DescribeNamespace:input_type -> infraboard.mcenter.namespace.DescriptNamespaceRequest
	4, // 3: infraboard.mcenter.namespace.RPC.QueryNamespace:output_type -> infraboard.mcenter.namespace.NamespaceSet
	5, // 4: infraboard.mcenter.namespace.RPC.DescribeNamespace:output_type -> infraboard.mcenter.namespace.Namespace
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mcenter_apps_namespace_pb_rpc_proto_init() }
func file_mcenter_apps_namespace_pb_rpc_proto_init() {
	if File_mcenter_apps_namespace_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_namespace_pb_namespace_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryNamespaceRequest); i {
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
		file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptNamespaceRequest); i {
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
		file_mcenter_apps_namespace_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRequest); i {
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
			RawDescriptor: file_mcenter_apps_namespace_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_namespace_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_namespace_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_namespace_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_namespace_pb_rpc_proto = out.File
	file_mcenter_apps_namespace_pb_rpc_proto_rawDesc = nil
	file_mcenter_apps_namespace_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_namespace_pb_rpc_proto_depIdxs = nil
}
