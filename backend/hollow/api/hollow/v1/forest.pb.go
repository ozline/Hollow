// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: hollow/v1/forest.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PushLeafRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid  int64  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushLeafRequest) Reset() {
	*x = PushLeafRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hollow_v1_forest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushLeafRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushLeafRequest) ProtoMessage() {}

func (x *PushLeafRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hollow_v1_forest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushLeafRequest.ProtoReflect.Descriptor instead.
func (*PushLeafRequest) Descriptor() ([]byte, []int) {
	return file_hollow_v1_forest_proto_rawDescGZIP(), []int{0}
}

func (x *PushLeafRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *PushLeafRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetLeafsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Pagesize int64 `protobuf:"varint,2,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
}

func (x *GetLeafsRequest) Reset() {
	*x = GetLeafsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hollow_v1_forest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeafsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeafsRequest) ProtoMessage() {}

func (x *GetLeafsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hollow_v1_forest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeafsRequest.ProtoReflect.Descriptor instead.
func (*GetLeafsRequest) Descriptor() ([]byte, []int) {
	return file_hollow_v1_forest_proto_rawDescGZIP(), []int{1}
}

func (x *GetLeafsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetLeafsRequest) GetPagesize() int64 {
	if x != nil {
		return x.Pagesize
	}
	return 0
}

type Leaf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner    int64  `protobuf:"varint,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	CreateAt int64  `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	Status   int64  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Leaf) Reset() {
	*x = Leaf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hollow_v1_forest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Leaf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leaf) ProtoMessage() {}

func (x *Leaf) ProtoReflect() protoreflect.Message {
	mi := &file_hollow_v1_forest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leaf.ProtoReflect.Descriptor instead.
func (*Leaf) Descriptor() ([]byte, []int) {
	return file_hollow_v1_forest_proto_rawDescGZIP(), []int{2}
}

func (x *Leaf) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Leaf) GetOwner() int64 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *Leaf) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Leaf) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Leaf) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type MultipleTodoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Leaf `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *MultipleTodoReply) Reset() {
	*x = MultipleTodoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hollow_v1_forest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleTodoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleTodoReply) ProtoMessage() {}

func (x *MultipleTodoReply) ProtoReflect() protoreflect.Message {
	mi := &file_hollow_v1_forest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleTodoReply.ProtoReflect.Descriptor instead.
func (*MultipleTodoReply) Descriptor() ([]byte, []int) {
	return file_hollow_v1_forest_proto_rawDescGZIP(), []int{3}
}

func (x *MultipleTodoReply) GetList() []*Leaf {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *MultipleTodoReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type PushLeafReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // Leaf data = 3;
}

func (x *PushLeafReply) Reset() {
	*x = PushLeafReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hollow_v1_forest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushLeafReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushLeafReply) ProtoMessage() {}

func (x *PushLeafReply) ProtoReflect() protoreflect.Message {
	mi := &file_hollow_v1_forest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushLeafReply.ProtoReflect.Descriptor instead.
func (*PushLeafReply) Descriptor() ([]byte, []int) {
	return file_hollow_v1_forest_proto_rawDescGZIP(), []int{4}
}

func (x *PushLeafReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PushLeafReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetLeafsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string             `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *MultipleTodoReply `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetLeafsReply) Reset() {
	*x = GetLeafsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hollow_v1_forest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeafsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeafsReply) ProtoMessage() {}

func (x *GetLeafsReply) ProtoReflect() protoreflect.Message {
	mi := &file_hollow_v1_forest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeafsReply.ProtoReflect.Descriptor instead.
func (*GetLeafsReply) Descriptor() ([]byte, []int) {
	return file_hollow_v1_forest_proto_rawDescGZIP(), []int{5}
}

func (x *GetLeafsReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetLeafsReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetLeafsReply) GetData() *MultipleTodoReply {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_hollow_v1_forest_proto protoreflect.FileDescriptor

var file_hollow_v1_forest_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x72, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x43, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61,
	0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x7b, 0x0a, 0x04, 0x4c, 0x65, 0x61,
	0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4e, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x6f, 0x72, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x35, 0x0a, 0x0d, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x65,
	0x61, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x67, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x66, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xbd, 0x01, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x59, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1a, 0x2e, 0x66, 0x6f, 0x72,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x65, 0x61, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x6f, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x57, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x65, 0x61, 0x66, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x2f,
	0x61, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x42, 0x44, 0x0a, 0x18, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0d, 0x46, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x17, 0x68, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hollow_v1_forest_proto_rawDescOnce sync.Once
	file_hollow_v1_forest_proto_rawDescData = file_hollow_v1_forest_proto_rawDesc
)

func file_hollow_v1_forest_proto_rawDescGZIP() []byte {
	file_hollow_v1_forest_proto_rawDescOnce.Do(func() {
		file_hollow_v1_forest_proto_rawDescData = protoimpl.X.CompressGZIP(file_hollow_v1_forest_proto_rawDescData)
	})
	return file_hollow_v1_forest_proto_rawDescData
}

var file_hollow_v1_forest_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hollow_v1_forest_proto_goTypes = []interface{}{
	(*PushLeafRequest)(nil),   // 0: forest.v1.PushLeafRequest
	(*GetLeafsRequest)(nil),   // 1: forest.v1.GetLeafsRequest
	(*Leaf)(nil),              // 2: forest.v1.Leaf
	(*MultipleTodoReply)(nil), // 3: forest.v1.MultipleTodoReply
	(*PushLeafReply)(nil),     // 4: forest.v1.PushLeafReply
	(*GetLeafsReply)(nil),     // 5: forest.v1.GetLeafsReply
}
var file_hollow_v1_forest_proto_depIdxs = []int32{
	2, // 0: forest.v1.MultipleTodoReply.list:type_name -> forest.v1.Leaf
	3, // 1: forest.v1.GetLeafsReply.data:type_name -> forest.v1.MultipleTodoReply
	0, // 2: forest.v1.Forests.Push:input_type -> forest.v1.PushLeafRequest
	1, // 3: forest.v1.Forests.Get:input_type -> forest.v1.GetLeafsRequest
	4, // 4: forest.v1.Forests.Push:output_type -> forest.v1.PushLeafReply
	5, // 5: forest.v1.Forests.Get:output_type -> forest.v1.GetLeafsReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hollow_v1_forest_proto_init() }
func file_hollow_v1_forest_proto_init() {
	if File_hollow_v1_forest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hollow_v1_forest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushLeafRequest); i {
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
		file_hollow_v1_forest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeafsRequest); i {
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
		file_hollow_v1_forest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Leaf); i {
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
		file_hollow_v1_forest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleTodoReply); i {
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
		file_hollow_v1_forest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushLeafReply); i {
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
		file_hollow_v1_forest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeafsReply); i {
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
			RawDescriptor: file_hollow_v1_forest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hollow_v1_forest_proto_goTypes,
		DependencyIndexes: file_hollow_v1_forest_proto_depIdxs,
		MessageInfos:      file_hollow_v1_forest_proto_msgTypes,
	}.Build()
	File_hollow_v1_forest_proto = out.File
	file_hollow_v1_forest_proto_rawDesc = nil
	file_hollow_v1_forest_proto_goTypes = nil
	file_hollow_v1_forest_proto_depIdxs = nil
}
