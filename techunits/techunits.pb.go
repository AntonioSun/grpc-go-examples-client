// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: techunits/techunits.proto

package SamplePackage

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

type EntryListRequest_Status int32

const (
	EntryListRequest_PUBLISHED EntryListRequest_Status = 0
	EntryListRequest_DRAFT     EntryListRequest_Status = 1
)

// Enum value maps for EntryListRequest_Status.
var (
	EntryListRequest_Status_name = map[int32]string{
		0: "PUBLISHED",
		1: "DRAFT",
	}
	EntryListRequest_Status_value = map[string]int32{
		"PUBLISHED": 0,
		"DRAFT":     1,
	}
)

func (x EntryListRequest_Status) Enum() *EntryListRequest_Status {
	p := new(EntryListRequest_Status)
	*p = x
	return p
}

func (x EntryListRequest_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryListRequest_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_techunits_techunits_proto_enumTypes[0].Descriptor()
}

func (EntryListRequest_Status) Type() protoreflect.EnumType {
	return &file_techunits_techunits_proto_enumTypes[0]
}

func (x EntryListRequest_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryListRequest_Status.Descriptor instead.
func (EntryListRequest_Status) EnumDescriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{7, 0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{0}
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_techunits_techunits_proto_rawDescGZIP(), []int{1}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Field_01  string `protobuf:"bytes,2,opt,name=field_01,json=field01,proto3" json:"field_01,omitempty"`
	Field_02  string `protobuf:"bytes,3,opt,name=field_02,json=field02,proto3" json:"field_02,omitempty"`
	Field_03  string `protobuf:"bytes,4,opt,name=field_03,json=field03,proto3" json:"field_03,omitempty"`
	Field_04  string `protobuf:"bytes,5,opt,name=field_04,json=field04,proto3" json:"field_04,omitempty"`
	Status    string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedOn int32  `protobuf:"varint,7,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_techunits_techunits_proto_rawDescGZIP(), []int{2}
}

func (x *PingResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PingResponse) GetField_01() string {
	if x != nil {
		return x.Field_01
	}
	return ""
}

func (x *PingResponse) GetField_02() string {
	if x != nil {
		return x.Field_02
	}
	return ""
}

func (x *PingResponse) GetField_03() string {
	if x != nil {
		return x.Field_03
	}
	return ""
}

func (x *PingResponse) GetField_04() string {
	if x != nil {
		return x.Field_04
	}
	return ""
}

func (x *PingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PingResponse) GetCreatedOn() int32 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

type EntryCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EntryCreateRequest) Reset() {
	*x = EntryCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryCreateRequest) ProtoMessage() {}

func (x *EntryCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryCreateRequest.ProtoReflect.Descriptor instead.
func (*EntryCreateRequest) Descriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{3}
}

func (x *EntryCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EntryCreateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *EntryCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EntryUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId     string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EntryUpdateRequest) Reset() {
	*x = EntryUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryUpdateRequest) ProtoMessage() {}

func (x *EntryUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryUpdateRequest.ProtoReflect.Descriptor instead.
func (*EntryUpdateRequest) Descriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{4}
}

func (x *EntryUpdateRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

func (x *EntryUpdateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EntryUpdateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *EntryUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EntryFileUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*EntryFileUploadRequest_EntryId
	//	*EntryFileUploadRequest_Chunk
	Data isEntryFileUploadRequest_Data `protobuf_oneof:"data"`
}

func (x *EntryFileUploadRequest) Reset() {
	*x = EntryFileUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryFileUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryFileUploadRequest) ProtoMessage() {}

func (x *EntryFileUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryFileUploadRequest.ProtoReflect.Descriptor instead.
func (*EntryFileUploadRequest) Descriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{5}
}

func (m *EntryFileUploadRequest) GetData() isEntryFileUploadRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *EntryFileUploadRequest) GetEntryId() string {
	if x, ok := x.GetData().(*EntryFileUploadRequest_EntryId); ok {
		return x.EntryId
	}
	return ""
}

func (x *EntryFileUploadRequest) GetChunk() []byte {
	if x, ok := x.GetData().(*EntryFileUploadRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isEntryFileUploadRequest_Data interface {
	isEntryFileUploadRequest_Data()
}

type EntryFileUploadRequest_EntryId struct {
	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3,oneof"`
}

type EntryFileUploadRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*EntryFileUploadRequest_EntryId) isEntryFileUploadRequest_Data() {}

func (*EntryFileUploadRequest_Chunk) isEntryFileUploadRequest_Data() {}

type EntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	FileUrl     string `protobuf:"bytes,5,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	CreatedOn   int32  `protobuf:"varint,6,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	ModifiedOn  int32  `protobuf:"varint,7,opt,name=modified_on,json=modifiedOn,proto3" json:"modified_on,omitempty"`
}

func (x *EntryResponse) Reset() {
	*x = EntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryResponse) ProtoMessage() {}

func (x *EntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryResponse.ProtoReflect.Descriptor instead.
func (*EntryResponse) Descriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{6}
}

func (x *EntryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EntryResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EntryResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *EntryResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EntryResponse) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *EntryResponse) GetCreatedOn() int32 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

func (x *EntryResponse) GetModifiedOn() int32 {
	if x != nil {
		return x.ModifiedOn
	}
	return 0
}

type EntryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32                   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page   int32                   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Status EntryListRequest_Status `protobuf:"varint,3,opt,name=status,proto3,enum=SamplePackage.EntryListRequest_Status" json:"status,omitempty"`
}

func (x *EntryListRequest) Reset() {
	*x = EntryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryListRequest) ProtoMessage() {}

func (x *EntryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryListRequest.ProtoReflect.Descriptor instead.
func (*EntryListRequest) Descriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{7}
}

func (x *EntryListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *EntryListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *EntryListRequest) GetStatus() EntryListRequest_Status {
	if x != nil {
		return x.Status
	}
	return EntryListRequest_PUBLISHED
}

type EntryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*EntryResponse `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *EntryListResponse) Reset() {
	*x = EntryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryListResponse) ProtoMessage() {}

func (x *EntryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryListResponse.ProtoReflect.Descriptor instead.
func (*EntryListResponse) Descriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{8}
}

func (x *EntryListResponse) GetEntries() []*EntryResponse {
	if x != nil {
		return x.Entries
	}
	return nil
}

type EntryDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *EntryDetailRequest) Reset() {
	*x = EntryDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techunits_techunits_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryDetailRequest) ProtoMessage() {}

func (x *EntryDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_techunits_techunits_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryDetailRequest.ProtoReflect.Descriptor instead.
func (*EntryDetailRequest) Descriptor() ([]byte, []int) {
	return file_techunits_techunits_proto_rawDescGZIP(), []int{9}
}

func (x *EntryDetailRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

var File_techunits_techunits_proto protoreflect.FileDescriptor

var file_techunits_techunits_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x63, 0x68, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x63, 0x68,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xc1, 0x01, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x30, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x30, 0x31, 0x12, 0x19,
	0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x30, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x30, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x30, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x30, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x30, 0x34,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x30, 0x34, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x60, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x16, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc6, 0x01, 0x0a,
	0x0d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x5f, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d,
	0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x01, 0x22, 0x4b, 0x0a, 0x11, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x32, 0xbd, 0x06, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x1a, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a,
	0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x21, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x50, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x25,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x52, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x51, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x21, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x21, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x64, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x3b, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_techunits_techunits_proto_rawDescOnce sync.Once
	file_techunits_techunits_proto_rawDescData = file_techunits_techunits_proto_rawDesc
)

func file_techunits_techunits_proto_rawDescGZIP() []byte {
	file_techunits_techunits_proto_rawDescOnce.Do(func() {
		file_techunits_techunits_proto_rawDescData = protoimpl.X.CompressGZIP(file_techunits_techunits_proto_rawDescData)
	})
	return file_techunits_techunits_proto_rawDescData
}

var file_techunits_techunits_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_techunits_techunits_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_techunits_techunits_proto_goTypes = []interface{}{
	(EntryListRequest_Status)(0),   // 0: SamplePackage.EntryListRequest.Status
	(*Empty)(nil),                  // 1: SamplePackage.Empty
	(*PingRequest)(nil),            // 2: SamplePackage.PingRequest
	(*PingResponse)(nil),           // 3: SamplePackage.PingResponse
	(*EntryCreateRequest)(nil),     // 4: SamplePackage.EntryCreateRequest
	(*EntryUpdateRequest)(nil),     // 5: SamplePackage.EntryUpdateRequest
	(*EntryFileUploadRequest)(nil), // 6: SamplePackage.EntryFileUploadRequest
	(*EntryResponse)(nil),          // 7: SamplePackage.EntryResponse
	(*EntryListRequest)(nil),       // 8: SamplePackage.EntryListRequest
	(*EntryListResponse)(nil),      // 9: SamplePackage.EntryListResponse
	(*EntryDetailRequest)(nil),     // 10: SamplePackage.EntryDetailRequest
}
var file_techunits_techunits_proto_depIdxs = []int32{
	0,  // 0: SamplePackage.EntryListRequest.status:type_name -> SamplePackage.EntryListRequest.Status
	7,  // 1: SamplePackage.EntryListResponse.entries:type_name -> SamplePackage.EntryResponse
	2,  // 2: SamplePackage.SampleService.ping:input_type -> SamplePackage.PingRequest
	4,  // 3: SamplePackage.SampleService.createEntry:input_type -> SamplePackage.EntryCreateRequest
	4,  // 4: SamplePackage.SampleService.createBulkEntries:input_type -> SamplePackage.EntryCreateRequest
	5,  // 5: SamplePackage.SampleService.updateEntry:input_type -> SamplePackage.EntryUpdateRequest
	6,  // 6: SamplePackage.SampleService.uploadEntryFile:input_type -> SamplePackage.EntryFileUploadRequest
	8,  // 7: SamplePackage.SampleService.listEntries:input_type -> SamplePackage.EntryListRequest
	8,  // 8: SamplePackage.SampleService.streamEntries:input_type -> SamplePackage.EntryListRequest
	10, // 9: SamplePackage.SampleService.getEntryInfo:input_type -> SamplePackage.EntryDetailRequest
	10, // 10: SamplePackage.SampleService.deleteEntry:input_type -> SamplePackage.EntryDetailRequest
	1,  // 11: SamplePackage.SampleService.notDefinedSample:input_type -> SamplePackage.Empty
	3,  // 12: SamplePackage.SampleService.ping:output_type -> SamplePackage.PingResponse
	7,  // 13: SamplePackage.SampleService.createEntry:output_type -> SamplePackage.EntryResponse
	7,  // 14: SamplePackage.SampleService.createBulkEntries:output_type -> SamplePackage.EntryResponse
	7,  // 15: SamplePackage.SampleService.updateEntry:output_type -> SamplePackage.EntryResponse
	7,  // 16: SamplePackage.SampleService.uploadEntryFile:output_type -> SamplePackage.EntryResponse
	9,  // 17: SamplePackage.SampleService.listEntries:output_type -> SamplePackage.EntryListResponse
	7,  // 18: SamplePackage.SampleService.streamEntries:output_type -> SamplePackage.EntryResponse
	7,  // 19: SamplePackage.SampleService.getEntryInfo:output_type -> SamplePackage.EntryResponse
	7,  // 20: SamplePackage.SampleService.deleteEntry:output_type -> SamplePackage.EntryResponse
	1,  // 21: SamplePackage.SampleService.notDefinedSample:output_type -> SamplePackage.Empty
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_techunits_techunits_proto_init() }
func file_techunits_techunits_proto_init() {
	if File_techunits_techunits_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_techunits_techunits_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_techunits_techunits_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_techunits_techunits_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_techunits_techunits_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryCreateRequest); i {
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
		file_techunits_techunits_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryUpdateRequest); i {
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
		file_techunits_techunits_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryFileUploadRequest); i {
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
		file_techunits_techunits_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryResponse); i {
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
		file_techunits_techunits_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryListRequest); i {
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
		file_techunits_techunits_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryListResponse); i {
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
		file_techunits_techunits_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryDetailRequest); i {
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
	file_techunits_techunits_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*EntryFileUploadRequest_EntryId)(nil),
		(*EntryFileUploadRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_techunits_techunits_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_techunits_techunits_proto_goTypes,
		DependencyIndexes: file_techunits_techunits_proto_depIdxs,
		EnumInfos:         file_techunits_techunits_proto_enumTypes,
		MessageInfos:      file_techunits_techunits_proto_msgTypes,
	}.Build()
	File_techunits_techunits_proto = out.File
	file_techunits_techunits_proto_rawDesc = nil
	file_techunits_techunits_proto_goTypes = nil
	file_techunits_techunits_proto_depIdxs = nil
}
