// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: to_favorite.proto

package services

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

//获赞数
type UpdateTotalFavoritedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Count  int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`                 // 增加的数量
	Type   int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`                   // 1是增加，2是减少
}

func (x *UpdateTotalFavoritedRequest) Reset() {
	*x = UpdateTotalFavoritedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTotalFavoritedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTotalFavoritedRequest) ProtoMessage() {}

func (x *UpdateTotalFavoritedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_to_favorite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTotalFavoritedRequest.ProtoReflect.Descriptor instead.
func (*UpdateTotalFavoritedRequest) Descriptor() ([]byte, []int) {
	return file_to_favorite_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTotalFavoritedRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateTotalFavoritedRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *UpdateTotalFavoritedRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type UpdateTotalFavoritedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` //响应，成功是0，失败是其他值
}

func (x *UpdateTotalFavoritedResponse) Reset() {
	*x = UpdateTotalFavoritedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_favorite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTotalFavoritedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTotalFavoritedResponse) ProtoMessage() {}

func (x *UpdateTotalFavoritedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_to_favorite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTotalFavoritedResponse.ProtoReflect.Descriptor instead.
func (*UpdateTotalFavoritedResponse) Descriptor() ([]byte, []int) {
	return file_to_favorite_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTotalFavoritedResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

//喜欢的视频数
type UpdateFavoriteCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Count  int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`                 // 增加的数量
	Type   int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`                   // 1是增加，2是减少
}

func (x *UpdateFavoriteCountRequest) Reset() {
	*x = UpdateFavoriteCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_favorite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFavoriteCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFavoriteCountRequest) ProtoMessage() {}

func (x *UpdateFavoriteCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_to_favorite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFavoriteCountRequest.ProtoReflect.Descriptor instead.
func (*UpdateFavoriteCountRequest) Descriptor() ([]byte, []int) {
	return file_to_favorite_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFavoriteCountRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateFavoriteCountRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *UpdateFavoriteCountRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type UpdateFavoriteCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` //响应，成功是0，失败是其他值
}

func (x *UpdateFavoriteCountResponse) Reset() {
	*x = UpdateFavoriteCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_favorite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFavoriteCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFavoriteCountResponse) ProtoMessage() {}

func (x *UpdateFavoriteCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_to_favorite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFavoriteCountResponse.ProtoReflect.Descriptor instead.
func (*UpdateFavoriteCountResponse) Descriptor() ([]byte, []int) {
	return file_to_favorite_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFavoriteCountResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_to_favorite_proto protoreflect.FileDescriptor

var file_to_favorite_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x6f, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x61, 0x0a,
	0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x40, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x60, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x3f, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xe2, 0x01, 0x0a, 0x11, 0x54, 0x6f, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x64, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_to_favorite_proto_rawDescOnce sync.Once
	file_to_favorite_proto_rawDescData = file_to_favorite_proto_rawDesc
)

func file_to_favorite_proto_rawDescGZIP() []byte {
	file_to_favorite_proto_rawDescOnce.Do(func() {
		file_to_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_to_favorite_proto_rawDescData)
	})
	return file_to_favorite_proto_rawDescData
}

var file_to_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_to_favorite_proto_goTypes = []interface{}{
	(*UpdateTotalFavoritedRequest)(nil),  // 0: service.UpdateTotalFavorited_request
	(*UpdateTotalFavoritedResponse)(nil), // 1: service.UpdateTotalFavorited_response
	(*UpdateFavoriteCountRequest)(nil),   // 2: service.UpdateFavoriteCount_request
	(*UpdateFavoriteCountResponse)(nil),  // 3: service.UpdateFavoriteCount_response
}
var file_to_favorite_proto_depIdxs = []int32{
	0, // 0: service.ToFavoriteService.UpdateTotalFavorited:input_type -> service.UpdateTotalFavorited_request
	2, // 1: service.ToFavoriteService.UpdateFavoriteCount:input_type -> service.UpdateFavoriteCount_request
	1, // 2: service.ToFavoriteService.UpdateTotalFavorited:output_type -> service.UpdateTotalFavorited_response
	3, // 3: service.ToFavoriteService.UpdateFavoriteCount:output_type -> service.UpdateFavoriteCount_response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_to_favorite_proto_init() }
func file_to_favorite_proto_init() {
	if File_to_favorite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_to_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTotalFavoritedRequest); i {
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
		file_to_favorite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTotalFavoritedResponse); i {
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
		file_to_favorite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFavoriteCountRequest); i {
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
		file_to_favorite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFavoriteCountResponse); i {
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
			RawDescriptor: file_to_favorite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_to_favorite_proto_goTypes,
		DependencyIndexes: file_to_favorite_proto_depIdxs,
		MessageInfos:      file_to_favorite_proto_msgTypes,
	}.Build()
	File_to_favorite_proto = out.File
	file_to_favorite_proto_rawDesc = nil
	file_to_favorite_proto_goTypes = nil
	file_to_favorite_proto_depIdxs = nil
}
