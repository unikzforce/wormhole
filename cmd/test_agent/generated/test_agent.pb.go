// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: cmd/test_agent/test_agent.proto

package generated

import (
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

type WaitUntilReadyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WaitUntilReadyResponse) Reset() {
	*x = WaitUntilReadyResponse{}
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WaitUntilReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitUntilReadyResponse) ProtoMessage() {}

func (x *WaitUntilReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitUntilReadyResponse.ProtoReflect.Descriptor instead.
func (*WaitUntilReadyResponse) Descriptor() ([]byte, []int) {
	return file_cmd_test_agent_test_agent_proto_rawDescGZIP(), []int{0}
}

func (x *WaitUntilReadyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IpV4Address   string                 `protobuf:"bytes,1,opt,name=ipV4Address,proto3" json:"ipV4Address,omitempty"`
	Count         int32                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Timeout       int32                  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_cmd_test_agent_test_agent_proto_rawDescGZIP(), []int{1}
}

func (x *PingRequest) GetIpV4Address() string {
	if x != nil {
		return x.IpV4Address
	}
	return ""
}

func (x *PingRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PingRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type PingResponse struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	Success               bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	PacketsRecv           int32                  `protobuf:"varint,2,opt,name=packetsRecv,proto3" json:"packetsRecv,omitempty"`
	PacketsSent           int32                  `protobuf:"varint,3,opt,name=packetsSent,proto3" json:"packetsSent,omitempty"`
	PacketsRecvDuplicates int32                  `protobuf:"varint,4,opt,name=packetsRecvDuplicates,proto3" json:"packetsRecvDuplicates,omitempty"`
	PacketLoss            float32                `protobuf:"fixed32,5,opt,name=packetLoss,proto3" json:"packetLoss,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_cmd_test_agent_test_agent_proto_rawDescGZIP(), []int{2}
}

func (x *PingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PingResponse) GetPacketsRecv() int32 {
	if x != nil {
		return x.PacketsRecv
	}
	return 0
}

func (x *PingResponse) GetPacketsSent() int32 {
	if x != nil {
		return x.PacketsSent
	}
	return 0
}

func (x *PingResponse) GetPacketsRecvDuplicates() int32 {
	if x != nil {
		return x.PacketsRecvDuplicates
	}
	return 0
}

func (x *PingResponse) GetPacketLoss() float32 {
	if x != nil {
		return x.PacketLoss
	}
	return 0
}

type EnableSwitchAgentRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	InterfaceNames []string               `protobuf:"bytes,1,rep,name=interfaceNames,proto3" json:"interfaceNames,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EnableSwitchAgentRequest) Reset() {
	*x = EnableSwitchAgentRequest{}
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableSwitchAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableSwitchAgentRequest) ProtoMessage() {}

func (x *EnableSwitchAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableSwitchAgentRequest.ProtoReflect.Descriptor instead.
func (*EnableSwitchAgentRequest) Descriptor() ([]byte, []int) {
	return file_cmd_test_agent_test_agent_proto_rawDescGZIP(), []int{3}
}

func (x *EnableSwitchAgentRequest) GetInterfaceNames() []string {
	if x != nil {
		return x.InterfaceNames
	}
	return nil
}

type EnableSwitchAgentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resp          string                 `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	Command       string                 `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Pid           int32                  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Hostname      string                 `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnableSwitchAgentResponse) Reset() {
	*x = EnableSwitchAgentResponse{}
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableSwitchAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableSwitchAgentResponse) ProtoMessage() {}

func (x *EnableSwitchAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_test_agent_test_agent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableSwitchAgentResponse.ProtoReflect.Descriptor instead.
func (*EnableSwitchAgentResponse) Descriptor() ([]byte, []int) {
	return file_cmd_test_agent_test_agent_proto_rawDescGZIP(), []int{4}
}

func (x *EnableSwitchAgentResponse) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

func (x *EnableSwitchAgentResponse) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *EnableSwitchAgentResponse) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *EnableSwitchAgentResponse) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

var File_cmd_test_agent_test_agent_proto protoreflect.FileDescriptor

var file_cmd_test_agent_test_agent_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6d, 0x64, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x16, 0x57, 0x61, 0x69,
	0x74, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x5f, 0x0a,
	0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x70, 0x56, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x70, 0x56, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xc2,
	0x01, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x63, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x63, 0x76, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x15, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x63, 0x76, 0x44, 0x75, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x63, 0x76, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c,
	0x6f, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x18, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x77, 0x0a, 0x19, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0xc6, 0x02, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x57, 0x61, 0x69, 0x74, 0x55, 0x6e, 0x74,
	0x69, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x21, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x57, 0x61, 0x69, 0x74,
	0x55, 0x6e, 0x74, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x60, 0x0a, 0x11, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_test_agent_test_agent_proto_rawDescOnce sync.Once
	file_cmd_test_agent_test_agent_proto_rawDescData = file_cmd_test_agent_test_agent_proto_rawDesc
)

func file_cmd_test_agent_test_agent_proto_rawDescGZIP() []byte {
	file_cmd_test_agent_test_agent_proto_rawDescOnce.Do(func() {
		file_cmd_test_agent_test_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_test_agent_test_agent_proto_rawDescData)
	})
	return file_cmd_test_agent_test_agent_proto_rawDescData
}

var file_cmd_test_agent_test_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cmd_test_agent_test_agent_proto_goTypes = []any{
	(*WaitUntilReadyResponse)(nil),    // 0: generated.WaitUntilReadyResponse
	(*PingRequest)(nil),               // 1: generated.PingRequest
	(*PingResponse)(nil),              // 2: generated.PingResponse
	(*EnableSwitchAgentRequest)(nil),  // 3: generated.EnableSwitchAgentRequest
	(*EnableSwitchAgentResponse)(nil), // 4: generated.EnableSwitchAgentResponse
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_cmd_test_agent_test_agent_proto_depIdxs = []int32{
	5, // 0: generated.TestAgentService.WaitUntilReady:input_type -> google.protobuf.Empty
	1, // 1: generated.TestAgentService.Ping:input_type -> generated.PingRequest
	3, // 2: generated.TestAgentService.EnableSwitchAgent:input_type -> generated.EnableSwitchAgentRequest
	5, // 3: generated.TestAgentService.DisableSwitchAgent:input_type -> google.protobuf.Empty
	0, // 4: generated.TestAgentService.WaitUntilReady:output_type -> generated.WaitUntilReadyResponse
	2, // 5: generated.TestAgentService.Ping:output_type -> generated.PingResponse
	4, // 6: generated.TestAgentService.EnableSwitchAgent:output_type -> generated.EnableSwitchAgentResponse
	5, // 7: generated.TestAgentService.DisableSwitchAgent:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_test_agent_test_agent_proto_init() }
func file_cmd_test_agent_test_agent_proto_init() {
	if File_cmd_test_agent_test_agent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_test_agent_test_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_test_agent_test_agent_proto_goTypes,
		DependencyIndexes: file_cmd_test_agent_test_agent_proto_depIdxs,
		MessageInfos:      file_cmd_test_agent_test_agent_proto_msgTypes,
	}.Build()
	File_cmd_test_agent_test_agent_proto = out.File
	file_cmd_test_agent_test_agent_proto_rawDesc = nil
	file_cmd_test_agent_test_agent_proto_goTypes = nil
	file_cmd_test_agent_test_agent_proto_depIdxs = nil
}
