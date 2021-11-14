// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/dummy_data_service.proto

package Magedatamonk

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

type Username struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Username) Reset() {
	*x = Username{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dummy_data_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Username) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Username) ProtoMessage() {}

func (x *Username) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dummy_data_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Username.ProtoReflect.Descriptor instead.
func (*Username) Descriptor() ([]byte, []int) {
	return file_proto_dummy_data_service_proto_rawDescGZIP(), []int{0}
}

func (x *Username) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Class string `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	Roll  int64  `protobuf:"varint,3,opt,name=roll,proto3" json:"roll,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dummy_data_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dummy_data_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_dummy_data_service_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *User) GetRoll() int64 {
	if x != nil {
		return x.Roll
	}
	return 0
}

var File_proto_dummy_data_service_proto protoreflect.FileDescriptor

var file_proto_dummy_data_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x32, 0x57, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x63,
	0x6b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x18, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x68, 0x75, 0x62, 0x68, 0x61, 0x6d, 0x74, 0x68, 0x61, 0x6b, 0x61, 0x72, 0x2f,
	0x4d, 0x61, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x6e, 0x6b, 0x3b, 0x4d, 0x61,
	0x67, 0x65, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_dummy_data_service_proto_rawDescOnce sync.Once
	file_proto_dummy_data_service_proto_rawDescData = file_proto_dummy_data_service_proto_rawDesc
)

func file_proto_dummy_data_service_proto_rawDescGZIP() []byte {
	file_proto_dummy_data_service_proto_rawDescOnce.Do(func() {
		file_proto_dummy_data_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dummy_data_service_proto_rawDescData)
	})
	return file_proto_dummy_data_service_proto_rawDescData
}

var file_proto_dummy_data_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_dummy_data_service_proto_goTypes = []interface{}{
	(*Username)(nil), // 0: dummy_data_service.Username
	(*User)(nil),     // 1: dummy_data_service.User
}
var file_proto_dummy_data_service_proto_depIdxs = []int32{
	0, // 0: dummy_data_service.UserName.GetMockUserData:input_type -> dummy_data_service.Username
	1, // 1: dummy_data_service.UserName.GetMockUserData:output_type -> dummy_data_service.User
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_dummy_data_service_proto_init() }
func file_proto_dummy_data_service_proto_init() {
	if File_proto_dummy_data_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dummy_data_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Username); i {
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
		file_proto_dummy_data_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_proto_dummy_data_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dummy_data_service_proto_goTypes,
		DependencyIndexes: file_proto_dummy_data_service_proto_depIdxs,
		MessageInfos:      file_proto_dummy_data_service_proto_msgTypes,
	}.Build()
	File_proto_dummy_data_service_proto = out.File
	file_proto_dummy_data_service_proto_rawDesc = nil
	file_proto_dummy_data_service_proto_goTypes = nil
	file_proto_dummy_data_service_proto_depIdxs = nil
}
