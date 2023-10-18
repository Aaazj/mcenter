// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: mcenter/apps/device/pb/rpc.proto

package device

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

type QueryDeviceByNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 设备分配的空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 设备分配的域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" bson:"domain"`
}

func (x *QueryDeviceByNamespaceRequest) Reset() {
	*x = QueryDeviceByNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDeviceByNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDeviceByNamespaceRequest) ProtoMessage() {}

func (x *QueryDeviceByNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDeviceByNamespaceRequest.ProtoReflect.Descriptor instead.
func (*QueryDeviceByNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryDeviceByNamespaceRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryDeviceByNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type DeviceByNamespaceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"items" json:"items"
	Items []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *DeviceByNamespaceSet) Reset() {
	*x = DeviceByNamespaceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceByNamespaceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceByNamespaceSet) ProtoMessage() {}

func (x *DeviceByNamespaceSet) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceByNamespaceSet.ProtoReflect.Descriptor instead.
func (*DeviceByNamespaceSet) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceByNamespaceSet) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

type AllocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 设备分配的域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 设备分配的空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: json:"entries"
	Entries []*DeviceEntry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries"`
}

func (x *AllocationRequest) Reset() {
	*x = AllocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationRequest) ProtoMessage() {}

func (x *AllocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationRequest.ProtoReflect.Descriptor instead.
func (*AllocationRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *AllocationRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *AllocationRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AllocationRequest) GetEntries() []*DeviceEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type AllocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"message"
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *AllocationResponse) Reset() {
	*x = AllocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationResponse) ProtoMessage() {}

func (x *AllocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationResponse.ProtoReflect.Descriptor instead.
func (*AllocationResponse) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *AllocationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// DescribeEndpointRequest todo
type DescribeDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *DescribeDeviceRequest) Reset() {
	*x = DescribeDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDeviceRequest) ProtoMessage() {}

func (x *DescribeDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDeviceRequest.ProtoReflect.Descriptor instead.
func (*DescribeDeviceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeDeviceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 设备分配的空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 设备分配的域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain" bson:"domain"`
}

func (x *QueryDeviceRequest) Reset() {
	*x = QueryDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDeviceRequest) ProtoMessage() {}

func (x *QueryDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDeviceRequest.ProtoReflect.Descriptor instead.
func (*QueryDeviceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *QueryDeviceRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryDeviceRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryDeviceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type ReleaseDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"names"
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names"`
}

func (x *ReleaseDevicesRequest) Reset() {
	*x = ReleaseDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseDevicesRequest) ProtoMessage() {}

func (x *ReleaseDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseDevicesRequest.ProtoReflect.Descriptor instead.
func (*ReleaseDevicesRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *ReleaseDevicesRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type ValidateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *ValidateDeviceRequest) Reset() {
	*x = ValidateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateDeviceRequest) ProtoMessage() {}

func (x *ValidateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateDeviceRequest.ProtoReflect.Descriptor instead.
func (*ValidateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateDeviceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Entry 条目
type DeviceRenewalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源名称
	// @gotags: bson:"name" json:"name" validate:"required"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name" validate:"required"`
	// 设备分配时间
	// @gotags: bson:"renewal_time" json:"renewal_time" validate:"required"
	RenewalTime int64 `protobuf:"varint,2,opt,name=RenewalTime,proto3" json:"renewal_time" bson:"renewal_time" validate:"required"`
}

func (x *DeviceRenewalRequest) Reset() {
	*x = DeviceRenewalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRenewalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRenewalRequest) ProtoMessage() {}

func (x *DeviceRenewalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRenewalRequest.ProtoReflect.Descriptor instead.
func (*DeviceRenewalRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *DeviceRenewalRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceRenewalRequest) GetRenewalTime() int64 {
	if x != nil {
		return x.RenewalTime
	}
	return 0
}

var File_mcenter_apps_device_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_device_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x0f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22,
	0x2c, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x80, 0x01,
	0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x2e, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x2b, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x82, 0x01,
	0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x22, 0x2d, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0x2b, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4c,
	0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xd9, 0x04, 0x0a,
	0x03, 0x52, 0x50, 0x43, 0x12, 0x4f, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x50, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x52, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x4f, 0x0a, 0x0e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x16, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x2d, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x4d, 0x0a, 0x0d, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x61, 0x61, 0x7a, 0x6a, 0x2f, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_device_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_device_pb_rpc_proto_rawDescData = file_mcenter_apps_device_pb_rpc_proto_rawDesc
)

func file_mcenter_apps_device_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_device_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_device_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_device_pb_rpc_proto_rawDescData)
	})
	return file_mcenter_apps_device_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_device_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mcenter_apps_device_pb_rpc_proto_goTypes = []interface{}{
	(*QueryDeviceByNamespaceRequest)(nil), // 0: mcenter.device.QueryDeviceByNamespaceRequest
	(*DeviceByNamespaceSet)(nil),          // 1: mcenter.device.DeviceByNamespaceSet
	(*AllocationRequest)(nil),             // 2: mcenter.device.AllocationRequest
	(*AllocationResponse)(nil),            // 3: mcenter.device.AllocationResponse
	(*DescribeDeviceRequest)(nil),         // 4: mcenter.device.DescribeDeviceRequest
	(*QueryDeviceRequest)(nil),            // 5: mcenter.device.QueryDeviceRequest
	(*ReleaseDevicesRequest)(nil),         // 6: mcenter.device.ReleaseDevicesRequest
	(*ValidateDeviceRequest)(nil),         // 7: mcenter.device.ValidateDeviceRequest
	(*DeviceRenewalRequest)(nil),          // 8: mcenter.device.DeviceRenewalRequest
	(*DeviceEntry)(nil),                   // 9: mcenter.device.DeviceEntry
	(*request.PageRequest)(nil),           // 10: infraboard.mcube.page.PageRequest
	(*Device)(nil),                        // 11: mcenter.device.Device
	(*DeviceSet)(nil),                     // 12: mcenter.device.DeviceSet
}
var file_mcenter_apps_device_pb_rpc_proto_depIdxs = []int32{
	9,  // 0: mcenter.device.AllocationRequest.entries:type_name -> mcenter.device.DeviceEntry
	10, // 1: mcenter.device.QueryDeviceRequest.page:type_name -> infraboard.mcube.page.PageRequest
	4,  // 2: mcenter.device.RPC.DescribeDevice:input_type -> mcenter.device.DescribeDeviceRequest
	5,  // 3: mcenter.device.RPC.QueryDevice:input_type -> mcenter.device.QueryDeviceRequest
	2,  // 4: mcenter.device.RPC.AllocationDevice:input_type -> mcenter.device.AllocationRequest
	6,  // 5: mcenter.device.RPC.ReleaseDevices:input_type -> mcenter.device.ReleaseDevicesRequest
	7,  // 6: mcenter.device.RPC.ValidateDevice:input_type -> mcenter.device.ValidateDeviceRequest
	0,  // 7: mcenter.device.RPC.QueryDeviceByNamespace:input_type -> mcenter.device.QueryDeviceByNamespaceRequest
	8,  // 8: mcenter.device.RPC.RenewalDevice:input_type -> mcenter.device.DeviceRenewalRequest
	11, // 9: mcenter.device.RPC.DescribeDevice:output_type -> mcenter.device.Device
	12, // 10: mcenter.device.RPC.QueryDevice:output_type -> mcenter.device.DeviceSet
	12, // 11: mcenter.device.RPC.AllocationDevice:output_type -> mcenter.device.DeviceSet
	12, // 12: mcenter.device.RPC.ReleaseDevices:output_type -> mcenter.device.DeviceSet
	11, // 13: mcenter.device.RPC.ValidateDevice:output_type -> mcenter.device.Device
	1,  // 14: mcenter.device.RPC.QueryDeviceByNamespace:output_type -> mcenter.device.DeviceByNamespaceSet
	11, // 15: mcenter.device.RPC.RenewalDevice:output_type -> mcenter.device.Device
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_mcenter_apps_device_pb_rpc_proto_init() }
func file_mcenter_apps_device_pb_rpc_proto_init() {
	if File_mcenter_apps_device_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_device_pb_device_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDeviceByNamespaceRequest); i {
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
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceByNamespaceSet); i {
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
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationRequest); i {
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
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationResponse); i {
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
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDeviceRequest); i {
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
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDeviceRequest); i {
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
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseDevicesRequest); i {
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
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateDeviceRequest); i {
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
		file_mcenter_apps_device_pb_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRenewalRequest); i {
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
			RawDescriptor: file_mcenter_apps_device_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_device_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_device_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_device_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_device_pb_rpc_proto = out.File
	file_mcenter_apps_device_pb_rpc_proto_rawDesc = nil
	file_mcenter_apps_device_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_device_pb_rpc_proto_depIdxs = nil
}
