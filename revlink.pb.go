// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: revlink.proto

package sdp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetReverseLinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account that the item belongs to
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The item that you would like to find reverse links for
	Item *Reference `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	// set to true to only return links that propagate configuration change impact
	FollowOnlyBlastPropagation bool `protobuf:"varint,3,opt,name=followOnlyBlastPropagation,proto3" json:"followOnlyBlastPropagation,omitempty"`
}

func (x *GetReverseLinksRequest) Reset() {
	*x = GetReverseLinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReverseLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReverseLinksRequest) ProtoMessage() {}

func (x *GetReverseLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_revlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReverseLinksRequest.ProtoReflect.Descriptor instead.
func (*GetReverseLinksRequest) Descriptor() ([]byte, []int) {
	return file_revlink_proto_rawDescGZIP(), []int{0}
}

func (x *GetReverseLinksRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *GetReverseLinksRequest) GetItem() *Reference {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *GetReverseLinksRequest) GetFollowOnlyBlastPropagation() bool {
	if x != nil {
		return x.FollowOnlyBlastPropagation
	}
	return false
}

type GetReverseLinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The item queries that should be executed in order to find items that link
	// to the requested item
	LinkedItemQueries []*LinkedItemQuery `protobuf:"bytes,1,rep,name=linkedItemQueries,proto3" json:"linkedItemQueries,omitempty"`
}

func (x *GetReverseLinksResponse) Reset() {
	*x = GetReverseLinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revlink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReverseLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReverseLinksResponse) ProtoMessage() {}

func (x *GetReverseLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_revlink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReverseLinksResponse.ProtoReflect.Descriptor instead.
func (*GetReverseLinksResponse) Descriptor() ([]byte, []int) {
	return file_revlink_proto_rawDescGZIP(), []int{1}
}

func (x *GetReverseLinksResponse) GetLinkedItemQueries() []*LinkedItemQuery {
	if x != nil {
		return x.LinkedItemQueries
	}
	return nil
}

type IngestGatewayResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account that the response belongs to
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The response type to ingest
	//
	// Types that are assignable to ResponseType:
	//
	//	*IngestGatewayResponseRequest_NewItem
	//	*IngestGatewayResponseRequest_NewEdge
	ResponseType isIngestGatewayResponseRequest_ResponseType `protobuf_oneof:"response_type"`
}

func (x *IngestGatewayResponseRequest) Reset() {
	*x = IngestGatewayResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revlink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestGatewayResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestGatewayResponseRequest) ProtoMessage() {}

func (x *IngestGatewayResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_revlink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestGatewayResponseRequest.ProtoReflect.Descriptor instead.
func (*IngestGatewayResponseRequest) Descriptor() ([]byte, []int) {
	return file_revlink_proto_rawDescGZIP(), []int{2}
}

func (x *IngestGatewayResponseRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (m *IngestGatewayResponseRequest) GetResponseType() isIngestGatewayResponseRequest_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return nil
}

func (x *IngestGatewayResponseRequest) GetNewItem() *Item {
	if x, ok := x.GetResponseType().(*IngestGatewayResponseRequest_NewItem); ok {
		return x.NewItem
	}
	return nil
}

func (x *IngestGatewayResponseRequest) GetNewEdge() *Edge {
	if x, ok := x.GetResponseType().(*IngestGatewayResponseRequest_NewEdge); ok {
		return x.NewEdge
	}
	return nil
}

type isIngestGatewayResponseRequest_ResponseType interface {
	isIngestGatewayResponseRequest_ResponseType()
}

type IngestGatewayResponseRequest_NewItem struct {
	NewItem *Item `protobuf:"bytes,2,opt,name=newItem,proto3,oneof"` // A new item that has been discovered
}

type IngestGatewayResponseRequest_NewEdge struct {
	NewEdge *Edge `protobuf:"bytes,3,opt,name=newEdge,proto3,oneof"` // A new edge between two items
}

func (*IngestGatewayResponseRequest_NewItem) isIngestGatewayResponseRequest_ResponseType() {}

func (*IngestGatewayResponseRequest_NewEdge) isIngestGatewayResponseRequest_ResponseType() {}

type IngestGatewayResponsesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumItemsReceived int32 `protobuf:"varint,1,opt,name=numItemsReceived,proto3" json:"numItemsReceived,omitempty"`
	NumEdgesReceived int32 `protobuf:"varint,2,opt,name=numEdgesReceived,proto3" json:"numEdgesReceived,omitempty"`
}

func (x *IngestGatewayResponsesResponse) Reset() {
	*x = IngestGatewayResponsesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revlink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestGatewayResponsesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestGatewayResponsesResponse) ProtoMessage() {}

func (x *IngestGatewayResponsesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_revlink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestGatewayResponsesResponse.ProtoReflect.Descriptor instead.
func (*IngestGatewayResponsesResponse) Descriptor() ([]byte, []int) {
	return file_revlink_proto_rawDescGZIP(), []int{3}
}

func (x *IngestGatewayResponsesResponse) GetNumItemsReceived() int32 {
	if x != nil {
		return x.NumItemsReceived
	}
	return 0
}

func (x *IngestGatewayResponsesResponse) GetNumEdgesReceived() int32 {
	if x != nil {
		return x.NumEdgesReceived
	}
	return 0
}

var File_revlink_proto protoreflect.FileDescriptor

var file_revlink_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x65, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x1a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x3e, 0x0a, 0x1a, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x6e, 0x6c, 0x79, 0x42, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x6e, 0x6c, 0x79, 0x42, 0x6c, 0x61, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x11, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x11, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x51,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x1c, 0x49, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x65, 0x77,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x45, 0x64, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07,
	0x6e, 0x65, 0x77, 0x45, 0x64, 0x67, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x78, 0x0a, 0x1e, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x75,
	0x6d, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x45, 0x64, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x6e, 0x75, 0x6d, 0x45, 0x64, 0x67, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x32, 0xd2, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x76, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x76, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x16, 0x49,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x72, 0x65, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72,
	0x65, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x74, 0x65,
	0x63, 0x68, 0x2f, 0x73, 0x64, 0x70, 0x2d, 0x67, 0x6f, 0x3b, 0x73, 0x64, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_revlink_proto_rawDescOnce sync.Once
	file_revlink_proto_rawDescData = file_revlink_proto_rawDesc
)

func file_revlink_proto_rawDescGZIP() []byte {
	file_revlink_proto_rawDescOnce.Do(func() {
		file_revlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_revlink_proto_rawDescData)
	})
	return file_revlink_proto_rawDescData
}

var file_revlink_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_revlink_proto_goTypes = []interface{}{
	(*GetReverseLinksRequest)(nil),         // 0: revlink.GetReverseLinksRequest
	(*GetReverseLinksResponse)(nil),        // 1: revlink.GetReverseLinksResponse
	(*IngestGatewayResponseRequest)(nil),   // 2: revlink.IngestGatewayResponseRequest
	(*IngestGatewayResponsesResponse)(nil), // 3: revlink.IngestGatewayResponsesResponse
	(*Reference)(nil),                      // 4: Reference
	(*LinkedItemQuery)(nil),                // 5: LinkedItemQuery
	(*Item)(nil),                           // 6: Item
	(*Edge)(nil),                           // 7: Edge
}
var file_revlink_proto_depIdxs = []int32{
	4, // 0: revlink.GetReverseLinksRequest.item:type_name -> Reference
	5, // 1: revlink.GetReverseLinksResponse.linkedItemQueries:type_name -> LinkedItemQuery
	6, // 2: revlink.IngestGatewayResponseRequest.newItem:type_name -> Item
	7, // 3: revlink.IngestGatewayResponseRequest.newEdge:type_name -> Edge
	0, // 4: revlink.RevlinkService.GetReverseLinks:input_type -> revlink.GetReverseLinksRequest
	2, // 5: revlink.RevlinkService.IngestGatewayResponses:input_type -> revlink.IngestGatewayResponseRequest
	1, // 6: revlink.RevlinkService.GetReverseLinks:output_type -> revlink.GetReverseLinksResponse
	3, // 7: revlink.RevlinkService.IngestGatewayResponses:output_type -> revlink.IngestGatewayResponsesResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_revlink_proto_init() }
func file_revlink_proto_init() {
	if File_revlink_proto != nil {
		return
	}
	file_items_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_revlink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReverseLinksRequest); i {
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
		file_revlink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReverseLinksResponse); i {
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
		file_revlink_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestGatewayResponseRequest); i {
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
		file_revlink_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestGatewayResponsesResponse); i {
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
	file_revlink_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*IngestGatewayResponseRequest_NewItem)(nil),
		(*IngestGatewayResponseRequest_NewEdge)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_revlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_revlink_proto_goTypes,
		DependencyIndexes: file_revlink_proto_depIdxs,
		MessageInfos:      file_revlink_proto_msgTypes,
	}.Build()
	File_revlink_proto = out.File
	file_revlink_proto_rawDesc = nil
	file_revlink_proto_goTypes = nil
	file_revlink_proto_depIdxs = nil
}