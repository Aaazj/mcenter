// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: mcenter/apps/device/pb/device.proto

package device

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

// Endpoint Service's features
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"_id" json:"id" validate:"required,lte=64"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id" validate:"required,lte=64"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at,omitempty"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty" bson:"create_at"`
	// 创建日期
	// @gotags: bson:",create_date" json:"create_date,omitempty"
	CreateDate string `protobuf:"bytes,3,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty" bson:",create_date"`
	// 设备分配的域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 设备分配的空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 设备条目信息
	// @gotags: bson:",inline" json:"entry" validate:"required"
	Entry *DeviceEntry `protobuf:"bytes,6,opt,name=entry,proto3" json:"entry" bson:",inline" validate:"required"`
	// 剩余时间
	// @gotags: bson:",remaining_days" json:"remaining_days,omitempty"
	RemainingDays int64 `protobuf:"varint,7,opt,name=remaining_days,json=remainingDays,proto3" json:"remaining_days,omitempty" bson:",remaining_days"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Device) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *Device) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Device) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Device) GetEntry() *DeviceEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *Device) GetRemainingDays() int64 {
	if x != nil {
		return x.RemainingDays
	}
	return 0
}

// Entry 条目
type DeviceEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// 设备分配时间
	// @gotags: bson:"allocate_days" json:"allocate_days"
	AllocateDays int64 `protobuf:"varint,2,opt,name=AllocateDays,proto3" json:"allocate_days" bson:"allocate_days"`
}

func (x *DeviceEntry) Reset() {
	*x = DeviceEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceEntry) ProtoMessage() {}

func (x *DeviceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceEntry.ProtoReflect.Descriptor instead.
func (*DeviceEntry) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceEntry) GetAllocateDays() int64 {
	if x != nil {
		return x.AllocateDays
	}
	return 0
}

type DeviceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Device `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *DeviceSet) Reset() {
	*x = DeviceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_device_pb_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceSet) ProtoMessage() {}

func (x *DeviceSet) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_device_pb_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceSet.ProtoReflect.Descriptor instead.
func (*DeviceSet) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_device_pb_device_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DeviceSet) GetItems() []*Device {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_mcenter_apps_device_pb_device_proto protoreflect.FileDescriptor

var file_mcenter_apps_device_pb_device_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0xe6, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x73, 0x22, 0x45,
	0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x44, 0x61, 0x79,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x79, 0x73, 0x22, 0x4f, 0x0a, 0x09, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x61, 0x61, 0x7a, 0x6a, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_device_pb_device_proto_rawDescOnce sync.Once
	file_mcenter_apps_device_pb_device_proto_rawDescData = file_mcenter_apps_device_pb_device_proto_rawDesc
)

func file_mcenter_apps_device_pb_device_proto_rawDescGZIP() []byte {
	file_mcenter_apps_device_pb_device_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_device_pb_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_device_pb_device_proto_rawDescData)
	})
	return file_mcenter_apps_device_pb_device_proto_rawDescData
}

var file_mcenter_apps_device_pb_device_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mcenter_apps_device_pb_device_proto_goTypes = []interface{}{
	(*Device)(nil),      // 0: mcenter.device.Device
	(*DeviceEntry)(nil), // 1: mcenter.device.DeviceEntry
	(*DeviceSet)(nil),   // 2: mcenter.device.DeviceSet
}
var file_mcenter_apps_device_pb_device_proto_depIdxs = []int32{
	1, // 0: mcenter.device.Device.entry:type_name -> mcenter.device.DeviceEntry
	0, // 1: mcenter.device.DeviceSet.items:type_name -> mcenter.device.Device
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mcenter_apps_device_pb_device_proto_init() }
func file_mcenter_apps_device_pb_device_proto_init() {
	if File_mcenter_apps_device_pb_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_device_pb_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_mcenter_apps_device_pb_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceEntry); i {
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
		file_mcenter_apps_device_pb_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceSet); i {
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
			RawDescriptor: file_mcenter_apps_device_pb_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_device_pb_device_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_device_pb_device_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_device_pb_device_proto_msgTypes,
	}.Build()
	File_mcenter_apps_device_pb_device_proto = out.File
	file_mcenter_apps_device_pb_device_proto_rawDesc = nil
	file_mcenter_apps_device_pb_device_proto_goTypes = nil
	file_mcenter_apps_device_pb_device_proto_depIdxs = nil
}