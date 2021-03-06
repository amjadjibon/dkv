// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: dkv.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KeyValueStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValueStore) Reset() {
	*x = KeyValueStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueStore) ProtoMessage() {}

func (x *KeyValueStore) ProtoReflect() protoreflect.Message {
	mi := &file_dkv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueStore.ProtoReflect.Descriptor instead.
func (*KeyValueStore) Descriptor() ([]byte, []int) {
	return file_dkv_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValueStore) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValueStore) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type KVGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KVGetRequest) Reset() {
	*x = KVGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVGetRequest) ProtoMessage() {}

func (x *KVGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVGetRequest.ProtoReflect.Descriptor instead.
func (*KVGetRequest) Descriptor() ([]byte, []int) {
	return file_dkv_proto_rawDescGZIP(), []int{1}
}

func (x *KVGetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type KVGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KVGetResponse) Reset() {
	*x = KVGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVGetResponse) ProtoMessage() {}

func (x *KVGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dkv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVGetResponse.ProtoReflect.Descriptor instead.
func (*KVGetResponse) Descriptor() ([]byte, []int) {
	return file_dkv_proto_rawDescGZIP(), []int{2}
}

func (x *KVGetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type KVPutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KVPutRequest) Reset() {
	*x = KVPutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVPutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVPutRequest) ProtoMessage() {}

func (x *KVPutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVPutRequest.ProtoReflect.Descriptor instead.
func (*KVPutRequest) Descriptor() ([]byte, []int) {
	return file_dkv_proto_rawDescGZIP(), []int{3}
}

func (x *KVPutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVPutRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type KVDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KVDeleteRequest) Reset() {
	*x = KVDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVDeleteRequest) ProtoMessage() {}

func (x *KVDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVDeleteRequest.ProtoReflect.Descriptor instead.
func (*KVDeleteRequest) Descriptor() ([]byte, []int) {
	return file_dkv_proto_rawDescGZIP(), []int{4}
}

func (x *KVDeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RaftJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId      string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeAddress string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
}

func (x *RaftJoinRequest) Reset() {
	*x = RaftJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftJoinRequest) ProtoMessage() {}

func (x *RaftJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftJoinRequest.ProtoReflect.Descriptor instead.
func (*RaftJoinRequest) Descriptor() ([]byte, []int) {
	return file_dkv_proto_rawDescGZIP(), []int{5}
}

func (x *RaftJoinRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RaftJoinRequest) GetNodeAddress() string {
	if x != nil {
		return x.NodeAddress
	}
	return ""
}

type RaftLeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *RaftLeaveRequest) Reset() {
	*x = RaftLeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftLeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftLeaveRequest) ProtoMessage() {}

func (x *RaftLeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftLeaveRequest.ProtoReflect.Descriptor instead.
func (*RaftLeaveRequest) Descriptor() ([]byte, []int) {
	return file_dkv_proto_rawDescGZIP(), []int{6}
}

func (x *RaftLeaveRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type RaftStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RaftState string `protobuf:"bytes,1,opt,name=raft_state,json=raftState,proto3" json:"raft_state,omitempty"`
}

func (x *RaftStateResponse) Reset() {
	*x = RaftStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftStateResponse) ProtoMessage() {}

func (x *RaftStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dkv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftStateResponse.ProtoReflect.Descriptor instead.
func (*RaftStateResponse) Descriptor() ([]byte, []int) {
	return file_dkv_proto_rawDescGZIP(), []int{7}
}

func (x *RaftStateResponse) GetRaftState() string {
	if x != nil {
		return x.RaftState
	}
	return ""
}

var File_dkv_proto protoreflect.FileDescriptor

var file_dkv_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a,
	0x0d, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x4b, 0x56, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x25, 0x0a, 0x0d, 0x4b, 0x56, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x36, 0x0a, 0x0c, 0x4b, 0x56, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x4b, 0x56, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x4d, 0x0a, 0x0f,
	0x52, 0x61, 0x66, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2b, 0x0a, 0x10, 0x52,
	0x61, 0x66, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x52, 0x61, 0x66, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x61, 0x66, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x32, 0xdf, 0x03, 0x0a,
	0x03, 0x44, 0x4b, 0x56, 0x12, 0x43, 0x0a, 0x05, 0x4b, 0x56, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x56, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x56, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x22, 0x04, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x45, 0x0a, 0x05, 0x4b, 0x56, 0x50,
	0x75, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x56, 0x50, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x04, 0x2f, 0x70, 0x75, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x4e, 0x0a, 0x08, 0x4b, 0x56, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x56, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x51, 0x0a, 0x08, 0x52, 0x61, 0x66, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x6a, 0x6f, 0x69, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x09, 0x52, 0x61, 0x66, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x72, 0x61, 0x66, 0x74,
	0x2f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x09, 0x52, 0x61, 0x66,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x12, 0x0c, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x27,
	0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x6a,
	0x61, 0x64, 0x6a, 0x69, 0x62, 0x6f, 0x6e, 0x2f, 0x64, 0x6b, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dkv_proto_rawDescOnce sync.Once
	file_dkv_proto_rawDescData = file_dkv_proto_rawDesc
)

func file_dkv_proto_rawDescGZIP() []byte {
	file_dkv_proto_rawDescOnce.Do(func() {
		file_dkv_proto_rawDescData = protoimpl.X.CompressGZIP(file_dkv_proto_rawDescData)
	})
	return file_dkv_proto_rawDescData
}

var file_dkv_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dkv_proto_goTypes = []interface{}{
	(*KeyValueStore)(nil),     // 0: proto.KeyValueStore
	(*KVGetRequest)(nil),      // 1: proto.KVGetRequest
	(*KVGetResponse)(nil),     // 2: proto.KVGetResponse
	(*KVPutRequest)(nil),      // 3: proto.KVPutRequest
	(*KVDeleteRequest)(nil),   // 4: proto.KVDeleteRequest
	(*RaftJoinRequest)(nil),   // 5: proto.RaftJoinRequest
	(*RaftLeaveRequest)(nil),  // 6: proto.RaftLeaveRequest
	(*RaftStateResponse)(nil), // 7: proto.RaftStateResponse
	(*emptypb.Empty)(nil),     // 8: google.protobuf.Empty
}
var file_dkv_proto_depIdxs = []int32{
	1, // 0: proto.DKV.KVGet:input_type -> proto.KVGetRequest
	3, // 1: proto.DKV.KVPut:input_type -> proto.KVPutRequest
	4, // 2: proto.DKV.KVDelete:input_type -> proto.KVDeleteRequest
	5, // 3: proto.DKV.RaftJoin:input_type -> proto.RaftJoinRequest
	6, // 4: proto.DKV.RaftLeave:input_type -> proto.RaftLeaveRequest
	8, // 5: proto.DKV.RaftState:input_type -> google.protobuf.Empty
	2, // 6: proto.DKV.KVGet:output_type -> proto.KVGetResponse
	8, // 7: proto.DKV.KVPut:output_type -> google.protobuf.Empty
	8, // 8: proto.DKV.KVDelete:output_type -> google.protobuf.Empty
	8, // 9: proto.DKV.RaftJoin:output_type -> google.protobuf.Empty
	8, // 10: proto.DKV.RaftLeave:output_type -> google.protobuf.Empty
	7, // 11: proto.DKV.RaftState:output_type -> proto.RaftStateResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dkv_proto_init() }
func file_dkv_proto_init() {
	if File_dkv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dkv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueStore); i {
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
		file_dkv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVGetRequest); i {
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
		file_dkv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVGetResponse); i {
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
		file_dkv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVPutRequest); i {
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
		file_dkv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVDeleteRequest); i {
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
		file_dkv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftJoinRequest); i {
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
		file_dkv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftLeaveRequest); i {
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
		file_dkv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftStateResponse); i {
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
			RawDescriptor: file_dkv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dkv_proto_goTypes,
		DependencyIndexes: file_dkv_proto_depIdxs,
		MessageInfos:      file_dkv_proto_msgTypes,
	}.Build()
	File_dkv_proto = out.File
	file_dkv_proto_rawDesc = nil
	file_dkv_proto_goTypes = nil
	file_dkv_proto_depIdxs = nil
}
