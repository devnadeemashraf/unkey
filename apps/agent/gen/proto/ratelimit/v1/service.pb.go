// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: proto/ratelimit/v1/service.proto

package ratelimitv1

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

type LivenessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LivenessRequest) Reset() {
	*x = LivenessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessRequest) ProtoMessage() {}

func (x *LivenessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessRequest.ProtoReflect.Descriptor instead.
func (*LivenessRequest) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{0}
}

type LivenessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LivenessResponse) Reset() {
	*x = LivenessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessResponse) ProtoMessage() {}

func (x *LivenessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessResponse.ProtoReflect.Descriptor instead.
func (*LivenessResponse) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *LivenessResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RatelimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Limit      int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Duration   int64  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Cost       int64  `protobuf:"varint,4,opt,name=cost,proto3" json:"cost,omitempty"`
	// A name for the ratelimit, used for debugging
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RatelimitRequest) Reset() {
	*x = RatelimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatelimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatelimitRequest) ProtoMessage() {}

func (x *RatelimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatelimitRequest.ProtoReflect.Descriptor instead.
func (*RatelimitRequest) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *RatelimitRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *RatelimitRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RatelimitRequest) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *RatelimitRequest) GetCost() int64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *RatelimitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RatelimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit     int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Remaining int64 `protobuf:"varint,2,opt,name=remaining,proto3" json:"remaining,omitempty"`
	Reset_    int64 `protobuf:"varint,3,opt,name=reset,proto3" json:"reset,omitempty"`
	Success   bool  `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
	Current   int64 `protobuf:"varint,5,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *RatelimitResponse) Reset() {
	*x = RatelimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatelimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatelimitResponse) ProtoMessage() {}

func (x *RatelimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatelimitResponse.ProtoReflect.Descriptor instead.
func (*RatelimitResponse) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *RatelimitResponse) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RatelimitResponse) GetRemaining() int64 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

func (x *RatelimitResponse) GetReset_() int64 {
	if x != nil {
		return x.Reset_
	}
	return 0
}

func (x *RatelimitResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RatelimitResponse) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

type RatelimitMultiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ratelimits []*RatelimitRequest `protobuf:"bytes,1,rep,name=ratelimits,proto3" json:"ratelimits,omitempty"`
}

func (x *RatelimitMultiRequest) Reset() {
	*x = RatelimitMultiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatelimitMultiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatelimitMultiRequest) ProtoMessage() {}

func (x *RatelimitMultiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatelimitMultiRequest.ProtoReflect.Descriptor instead.
func (*RatelimitMultiRequest) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *RatelimitMultiRequest) GetRatelimits() []*RatelimitRequest {
	if x != nil {
		return x.Ratelimits
	}
	return nil
}

type RatelimitMultiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ratelimits []*RatelimitResponse `protobuf:"bytes,1,rep,name=ratelimits,proto3" json:"ratelimits,omitempty"`
}

func (x *RatelimitMultiResponse) Reset() {
	*x = RatelimitMultiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatelimitMultiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatelimitMultiResponse) ProtoMessage() {}

func (x *RatelimitMultiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatelimitMultiResponse.ProtoReflect.Descriptor instead.
func (*RatelimitMultiResponse) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *RatelimitMultiResponse) GetRatelimits() []*RatelimitResponse {
	if x != nil {
		return x.Ratelimits
	}
	return nil
}

type PushPullEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Limit      int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Duration   int64  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Cost       int64  `protobuf:"varint,4,opt,name=cost,proto3" json:"cost,omitempty"`
	// used for metrics
	Time int64 `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *PushPullEvent) Reset() {
	*x = PushPullEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushPullEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushPullEvent) ProtoMessage() {}

func (x *PushPullEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushPullEvent.ProtoReflect.Descriptor instead.
func (*PushPullEvent) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *PushPullEvent) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *PushPullEvent) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PushPullEvent) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *PushPullEvent) GetCost() int64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *PushPullEvent) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type PushPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*PushPullEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *PushPullRequest) Reset() {
	*x = PushPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushPullRequest) ProtoMessage() {}

func (x *PushPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushPullRequest.ProtoReflect.Descriptor instead.
func (*PushPullRequest) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *PushPullRequest) GetEvents() []*PushPullEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type PushPullUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Current    int64  `protobuf:"varint,2,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *PushPullUpdate) Reset() {
	*x = PushPullUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushPullUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushPullUpdate) ProtoMessage() {}

func (x *PushPullUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushPullUpdate.ProtoReflect.Descriptor instead.
func (*PushPullUpdate) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *PushPullUpdate) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *PushPullUpdate) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

type PushPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updates []*PushPullUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
}

func (x *PushPullResponse) Reset() {
	*x = PushPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ratelimit_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushPullResponse) ProtoMessage() {}

func (x *PushPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ratelimit_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushPullResponse.ProtoReflect.Descriptor instead.
func (*PushPullResponse) Descriptor() ([]byte, []int) {
	return file_proto_ratelimit_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *PushPullResponse) GetUpdates() []*PushPullUpdate {
	if x != nil {
		return x.Updates
	}
	return nil
}

var File_proto_ratelimit_v1_service_proto protoreflect.FileDescriptor

var file_proto_ratelimit_v1_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x10, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x8c, 0x01, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x91,
	0x01, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x22, 0x57, 0x0a, 0x15, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0a, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x59, 0x0a, 0x16, 0x52,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x73, 0x68, 0x50,
	0x75, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x46, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x4a, 0x0a, 0x0e, 0x50, 0x75,
	0x73, 0x68, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x50, 0x75,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x50,
	0x75, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x32, 0xdb, 0x02, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x23, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x50, 0x75, 0x6c, 0x6c, 0x12,
	0x1d, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75,
	0x6e, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_ratelimit_v1_service_proto_rawDescOnce sync.Once
	file_proto_ratelimit_v1_service_proto_rawDescData = file_proto_ratelimit_v1_service_proto_rawDesc
)

func file_proto_ratelimit_v1_service_proto_rawDescGZIP() []byte {
	file_proto_ratelimit_v1_service_proto_rawDescOnce.Do(func() {
		file_proto_ratelimit_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ratelimit_v1_service_proto_rawDescData)
	})
	return file_proto_ratelimit_v1_service_proto_rawDescData
}

var file_proto_ratelimit_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_ratelimit_v1_service_proto_goTypes = []any{
	(*LivenessRequest)(nil),        // 0: ratelimit.v1.LivenessRequest
	(*LivenessResponse)(nil),       // 1: ratelimit.v1.LivenessResponse
	(*RatelimitRequest)(nil),       // 2: ratelimit.v1.RatelimitRequest
	(*RatelimitResponse)(nil),      // 3: ratelimit.v1.RatelimitResponse
	(*RatelimitMultiRequest)(nil),  // 4: ratelimit.v1.RatelimitMultiRequest
	(*RatelimitMultiResponse)(nil), // 5: ratelimit.v1.RatelimitMultiResponse
	(*PushPullEvent)(nil),          // 6: ratelimit.v1.PushPullEvent
	(*PushPullRequest)(nil),        // 7: ratelimit.v1.PushPullRequest
	(*PushPullUpdate)(nil),         // 8: ratelimit.v1.PushPullUpdate
	(*PushPullResponse)(nil),       // 9: ratelimit.v1.PushPullResponse
}
var file_proto_ratelimit_v1_service_proto_depIdxs = []int32{
	2, // 0: ratelimit.v1.RatelimitMultiRequest.ratelimits:type_name -> ratelimit.v1.RatelimitRequest
	3, // 1: ratelimit.v1.RatelimitMultiResponse.ratelimits:type_name -> ratelimit.v1.RatelimitResponse
	6, // 2: ratelimit.v1.PushPullRequest.events:type_name -> ratelimit.v1.PushPullEvent
	8, // 3: ratelimit.v1.PushPullResponse.updates:type_name -> ratelimit.v1.PushPullUpdate
	0, // 4: ratelimit.v1.RatelimitService.Liveness:input_type -> ratelimit.v1.LivenessRequest
	2, // 5: ratelimit.v1.RatelimitService.Ratelimit:input_type -> ratelimit.v1.RatelimitRequest
	4, // 6: ratelimit.v1.RatelimitService.MultiRatelimit:input_type -> ratelimit.v1.RatelimitMultiRequest
	7, // 7: ratelimit.v1.RatelimitService.PushPull:input_type -> ratelimit.v1.PushPullRequest
	1, // 8: ratelimit.v1.RatelimitService.Liveness:output_type -> ratelimit.v1.LivenessResponse
	3, // 9: ratelimit.v1.RatelimitService.Ratelimit:output_type -> ratelimit.v1.RatelimitResponse
	5, // 10: ratelimit.v1.RatelimitService.MultiRatelimit:output_type -> ratelimit.v1.RatelimitMultiResponse
	9, // 11: ratelimit.v1.RatelimitService.PushPull:output_type -> ratelimit.v1.PushPullResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ratelimit_v1_service_proto_init() }
func file_proto_ratelimit_v1_service_proto_init() {
	if File_proto_ratelimit_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ratelimit_v1_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LivenessRequest); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LivenessResponse); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RatelimitRequest); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RatelimitResponse); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RatelimitMultiRequest); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RatelimitMultiResponse); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PushPullEvent); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PushPullRequest); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*PushPullUpdate); i {
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
		file_proto_ratelimit_v1_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*PushPullResponse); i {
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
			RawDescriptor: file_proto_ratelimit_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ratelimit_v1_service_proto_goTypes,
		DependencyIndexes: file_proto_ratelimit_v1_service_proto_depIdxs,
		MessageInfos:      file_proto_ratelimit_v1_service_proto_msgTypes,
	}.Build()
	File_proto_ratelimit_v1_service_proto = out.File
	file_proto_ratelimit_v1_service_proto_rawDesc = nil
	file_proto_ratelimit_v1_service_proto_goTypes = nil
	file_proto_ratelimit_v1_service_proto_depIdxs = nil
}
