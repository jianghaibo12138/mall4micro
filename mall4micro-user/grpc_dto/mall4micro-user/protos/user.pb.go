// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: mall4micro-user/protos/user.proto

package protos

import (
	protos "github.com/jianghaibo12138/mall4micro/mall4micro-common/grpc_dto/mall4micro-common/protos"
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

type RpcUserInfoByUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *RpcUserInfoByUsernameRequest) Reset() {
	*x = RpcUserInfoByUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall4micro_user_protos_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcUserInfoByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcUserInfoByUsernameRequest) ProtoMessage() {}

func (x *RpcUserInfoByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mall4micro_user_protos_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcUserInfoByUsernameRequest.ProtoReflect.Descriptor instead.
func (*RpcUserInfoByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_mall4micro_user_protos_user_proto_rawDescGZIP(), []int{0}
}

func (x *RpcUserInfoByUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type RpcUserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply    *protos.RpcReply `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Username string           `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string           `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email    string           `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Mobile   string           `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status   int64            `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	SaltStr  string           `protobuf:"bytes,7,opt,name=salt_str,json=saltStr,proto3" json:"salt_str,omitempty"`
}

func (x *RpcUserInfoReply) Reset() {
	*x = RpcUserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mall4micro_user_protos_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcUserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcUserInfoReply) ProtoMessage() {}

func (x *RpcUserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_mall4micro_user_protos_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcUserInfoReply.ProtoReflect.Descriptor instead.
func (*RpcUserInfoReply) Descriptor() ([]byte, []int) {
	return file_mall4micro_user_protos_user_proto_rawDescGZIP(), []int{1}
}

func (x *RpcUserInfoReply) GetReply() *protos.RpcReply {
	if x != nil {
		return x.Reply
	}
	return nil
}

func (x *RpcUserInfoReply) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RpcUserInfoReply) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RpcUserInfoReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RpcUserInfoReply) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RpcUserInfoReply) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RpcUserInfoReply) GetSaltStr() string {
	if x != nil {
		return x.SaltStr
	}
	return ""
}

var File_mall4micro_user_protos_user_proto protoreflect.FileDescriptor

var file_mall4micro_user_protos_user_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x61, 0x6c, 0x6c, 0x34, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x64, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x61, 0x6c, 0x6c, 0x34, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a,
	0x0a, 0x1c, 0x52, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x10, 0x52,
	0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1f, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x61, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x61, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x32, 0x76, 0x0a, 0x18, 0x52, 0x70, 0x63,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x53, 0x72, 0x76, 0x12, 0x5a, 0x0a, 0x1c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x70, 0x63,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x53, 0x72, 0x76, 0x12, 0x21, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x52, 0x70, 0x63, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x52,
	0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x69, 0x61, 0x6e, 0x67, 0x68, 0x61, 0x69, 0x62, 0x6f, 0x31, 0x32, 0x31, 0x33, 0x38, 0x2f,
	0x6d, 0x61, 0x6c, 0x6c, 0x34, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x6d, 0x61, 0x6c, 0x6c, 0x34,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x64, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x6c, 0x6c, 0x34, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mall4micro_user_protos_user_proto_rawDescOnce sync.Once
	file_mall4micro_user_protos_user_proto_rawDescData = file_mall4micro_user_protos_user_proto_rawDesc
)

func file_mall4micro_user_protos_user_proto_rawDescGZIP() []byte {
	file_mall4micro_user_protos_user_proto_rawDescOnce.Do(func() {
		file_mall4micro_user_protos_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_mall4micro_user_protos_user_proto_rawDescData)
	})
	return file_mall4micro_user_protos_user_proto_rawDescData
}

var file_mall4micro_user_protos_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mall4micro_user_protos_user_proto_goTypes = []interface{}{
	(*RpcUserInfoByUsernameRequest)(nil), // 0: dto.RpcUserInfoByUsernameRequest
	(*RpcUserInfoReply)(nil),             // 1: dto.RpcUserInfoReply
	(*protos.RpcReply)(nil),              // 2: RpcReply
}
var file_mall4micro_user_protos_user_proto_depIdxs = []int32{
	2, // 0: dto.RpcUserInfoReply.reply:type_name -> RpcReply
	0, // 1: dto.RpcUserInfoByUsernameSrv.CallRpcUserInfoByUsernameSrv:input_type -> dto.RpcUserInfoByUsernameRequest
	1, // 2: dto.RpcUserInfoByUsernameSrv.CallRpcUserInfoByUsernameSrv:output_type -> dto.RpcUserInfoReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mall4micro_user_protos_user_proto_init() }
func file_mall4micro_user_protos_user_proto_init() {
	if File_mall4micro_user_protos_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mall4micro_user_protos_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcUserInfoByUsernameRequest); i {
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
		file_mall4micro_user_protos_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcUserInfoReply); i {
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
			RawDescriptor: file_mall4micro_user_protos_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mall4micro_user_protos_user_proto_goTypes,
		DependencyIndexes: file_mall4micro_user_protos_user_proto_depIdxs,
		MessageInfos:      file_mall4micro_user_protos_user_proto_msgTypes,
	}.Build()
	File_mall4micro_user_protos_user_proto = out.File
	file_mall4micro_user_protos_user_proto_rawDesc = nil
	file_mall4micro_user_protos_user_proto_goTypes = nil
	file_mall4micro_user_protos_user_proto_depIdxs = nil
}
