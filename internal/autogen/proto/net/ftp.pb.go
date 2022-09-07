// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: net/ftp.proto

package net

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

//
type Ftp_TransferState int32

const (
	Ftp_PENDING                     Ftp_TransferState = 0
	Ftp_TRANSFER_IN_PROGRESS        Ftp_TransferState = 1
	Ftp_POST_PROCESSING_IN_PROGRESS Ftp_TransferState = 2
	Ftp_COMPLETED                   Ftp_TransferState = 3
	Ftp_SUSPENDED                   Ftp_TransferState = 4
	Ftp_FAILED                      Ftp_TransferState = 5
)

// Enum value maps for Ftp_TransferState.
var (
	Ftp_TransferState_name = map[int32]string{
		0: "PENDING",
		1: "TRANSFER_IN_PROGRESS",
		2: "POST_PROCESSING_IN_PROGRESS",
		3: "COMPLETED",
		4: "SUSPENDED",
		5: "FAILED",
	}
	Ftp_TransferState_value = map[string]int32{
		"PENDING":                     0,
		"TRANSFER_IN_PROGRESS":        1,
		"POST_PROCESSING_IN_PROGRESS": 2,
		"COMPLETED":                   3,
		"SUSPENDED":                   4,
		"FAILED":                      5,
	}
)

func (x Ftp_TransferState) Enum() *Ftp_TransferState {
	p := new(Ftp_TransferState)
	*p = x
	return p
}

func (x Ftp_TransferState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ftp_TransferState) Descriptor() protoreflect.EnumDescriptor {
	return file_net_ftp_proto_enumTypes[0].Descriptor()
}

func (Ftp_TransferState) Type() protoreflect.EnumType {
	return &file_net_ftp_proto_enumTypes[0]
}

func (x Ftp_TransferState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ftp_TransferState.Descriptor instead.
func (Ftp_TransferState) EnumDescriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 0}
}

type Ftp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ftp) Reset() {
	*x = Ftp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp) ProtoMessage() {}

func (x *Ftp) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp.ProtoReflect.Descriptor instead.
func (*Ftp) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0}
}

type Ftp_FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	RelPath string `protobuf:"bytes,1,opt,name=rel_path,json=relPath,proto3" json:"rel_path,omitempty"`
	// File size.
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// Absolute pathname of the file on client side, optional for the Web browser clients.
	AbsPath *string `protobuf:"bytes,3,opt,name=abs_path,json=absPath,proto3,oneof" json:"abs_path,omitempty"`
	// File MIME type. Empty if can't be determined.
	MimeType *string `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3,oneof" json:"mime_type,omitempty"`
	// If set to true will invoke a registered post action handler.
	// The action handlers are being mapped by file extensions.
	PostAction *bool `protobuf:"varint,6,opt,name=post_action,json=postAction,proto3,oneof" json:"post_action,omitempty"`
}

func (x *Ftp_FileInfo) Reset() {
	*x = Ftp_FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_FileInfo) ProtoMessage() {}

func (x *Ftp_FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_FileInfo.ProtoReflect.Descriptor instead.
func (*Ftp_FileInfo) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Ftp_FileInfo) GetRelPath() string {
	if x != nil {
		return x.RelPath
	}
	return ""
}

func (x *Ftp_FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Ftp_FileInfo) GetAbsPath() string {
	if x != nil && x.AbsPath != nil {
		return *x.AbsPath
	}
	return ""
}

func (x *Ftp_FileInfo) GetMimeType() string {
	if x != nil && x.MimeType != nil {
		return *x.MimeType
	}
	return ""
}

func (x *Ftp_FileInfo) GetPostAction() bool {
	if x != nil && x.PostAction != nil {
		return *x.PostAction
	}
	return false
}

//
type Ftp_FileDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handle uint64        `protobuf:"fixed64,1,opt,name=handle,proto3" json:"handle,omitempty"`
	Info   *Ftp_FileInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Ftp_FileDescriptor) Reset() {
	*x = Ftp_FileDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_FileDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_FileDescriptor) ProtoMessage() {}

func (x *Ftp_FileDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_FileDescriptor.ProtoReflect.Descriptor instead.
func (*Ftp_FileDescriptor) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Ftp_FileDescriptor) GetHandle() uint64 {
	if x != nil {
		return x.Handle
	}
	return 0
}

func (x *Ftp_FileDescriptor) GetInfo() *Ftp_FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

//
type Ftp_FileRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Ftp_FileRange) Reset() {
	*x = Ftp_FileRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_FileRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_FileRange) ProtoMessage() {}

func (x *Ftp_FileRange) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_FileRange.ProtoReflect.Descriptor instead.
func (*Ftp_FileRange) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Ftp_FileRange) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Ftp_FileRange) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

//
type Ftp_NewSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ftp_NewSession) Reset() {
	*x = Ftp_NewSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_NewSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_NewSession) ProtoMessage() {}

func (x *Ftp_NewSession) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_NewSession.ProtoReflect.Descriptor instead.
func (*Ftp_NewSession) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 3}
}

//
type Ftp_FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ftp_FileChunk) Reset() {
	*x = Ftp_FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_FileChunk) ProtoMessage() {}

func (x *Ftp_FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_FileChunk.ProtoReflect.Descriptor instead.
func (*Ftp_FileChunk) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 4}
}

//
type Ftp_PendingChunks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileHandle uint64           `protobuf:"fixed64,1,opt,name=file_handle,json=fileHandle,proto3" json:"file_handle,omitempty"`
	Ranges     []*Ftp_FileRange `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *Ftp_PendingChunks) Reset() {
	*x = Ftp_PendingChunks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_PendingChunks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_PendingChunks) ProtoMessage() {}

func (x *Ftp_PendingChunks) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_PendingChunks.ProtoReflect.Descriptor instead.
func (*Ftp_PendingChunks) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Ftp_PendingChunks) GetFileHandle() uint64 {
	if x != nil {
		return x.FileHandle
	}
	return 0
}

func (x *Ftp_PendingChunks) GetRanges() []*Ftp_FileRange {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type Ftp_Inquiry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ftp_Inquiry) Reset() {
	*x = Ftp_Inquiry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_Inquiry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_Inquiry) ProtoMessage() {}

func (x *Ftp_Inquiry) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_Inquiry.ProtoReflect.Descriptor instead.
func (*Ftp_Inquiry) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 6}
}

type Ftp_NewSession_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*Ftp_FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Temp  bool            `protobuf:"varint,2,opt,name=temp,proto3" json:"temp,omitempty"`
	// Default to zero. Specifies file permissions, location of the base directory and maximum file size.
	UploadProfile *uint32 `protobuf:"varint,3,opt,name=uploadProfile,proto3,oneof" json:"uploadProfile,omitempty"`
	Lifetime      *uint32 `protobuf:"varint,4,opt,name=lifetime,proto3,oneof" json:"lifetime,omitempty"`
}

func (x *Ftp_NewSession_Request) Reset() {
	*x = Ftp_NewSession_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_NewSession_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_NewSession_Request) ProtoMessage() {}

func (x *Ftp_NewSession_Request) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_NewSession_Request.ProtoReflect.Descriptor instead.
func (*Ftp_NewSession_Request) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 3, 0}
}

func (x *Ftp_NewSession_Request) GetFiles() []*Ftp_FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Ftp_NewSession_Request) GetTemp() bool {
	if x != nil {
		return x.Temp
	}
	return false
}

func (x *Ftp_NewSession_Request) GetUploadProfile() uint32 {
	if x != nil && x.UploadProfile != nil {
		return *x.UploadProfile
	}
	return 0
}

func (x *Ftp_NewSession_Request) GetLifetime() uint32 {
	if x != nil && x.Lifetime != nil {
		return *x.Lifetime
	}
	return 0
}

type Ftp_NewSession_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxChunkSize uint32                `protobuf:"varint,1,opt,name=max_chunk_size,json=maxChunkSize,proto3" json:"max_chunk_size,omitempty"`
	Descriptors  []*Ftp_FileDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
}

func (x *Ftp_NewSession_Response) Reset() {
	*x = Ftp_NewSession_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_NewSession_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_NewSession_Response) ProtoMessage() {}

func (x *Ftp_NewSession_Response) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_NewSession_Response.ProtoReflect.Descriptor instead.
func (*Ftp_NewSession_Response) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 3, 1}
}

func (x *Ftp_NewSession_Response) GetMaxChunkSize() uint32 {
	if x != nil {
		return x.MaxChunkSize
	}
	return 0
}

func (x *Ftp_NewSession_Response) GetDescriptors() []*Ftp_FileDescriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

type Ftp_FileChunk_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range      *Ftp_FileRange `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	FileHandle uint64         `protobuf:"fixed64,2,opt,name=file_handle,json=fileHandle,proto3" json:"file_handle,omitempty"`
	Body       []byte         `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Ftp_FileChunk_Request) Reset() {
	*x = Ftp_FileChunk_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_FileChunk_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_FileChunk_Request) ProtoMessage() {}

func (x *Ftp_FileChunk_Request) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_FileChunk_Request.ProtoReflect.Descriptor instead.
func (*Ftp_FileChunk_Request) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 4, 0}
}

func (x *Ftp_FileChunk_Request) GetRange() *Ftp_FileRange {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *Ftp_FileChunk_Request) GetFileHandle() uint64 {
	if x != nil {
		return x.FileHandle
	}
	return 0
}

func (x *Ftp_FileChunk_Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Ftp_FileChunk_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State Ftp_TransferState `protobuf:"varint,1,opt,name=state,proto3,enum=go_serv.net.Ftp_TransferState" json:"state,omitempty"`
}

func (x *Ftp_FileChunk_Response) Reset() {
	*x = Ftp_FileChunk_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_FileChunk_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_FileChunk_Response) ProtoMessage() {}

func (x *Ftp_FileChunk_Response) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_FileChunk_Response.ProtoReflect.Descriptor instead.
func (*Ftp_FileChunk_Response) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 4, 1}
}

func (x *Ftp_FileChunk_Response) GetState() Ftp_TransferState {
	if x != nil {
		return x.State
	}
	return Ftp_PENDING
}

type Ftp_Inquiry_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferSessId uint64 `protobuf:"fixed64,1,opt,name=transfer_sess_id,json=transferSessId,proto3" json:"transfer_sess_id,omitempty"`
}

func (x *Ftp_Inquiry_Request) Reset() {
	*x = Ftp_Inquiry_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_Inquiry_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_Inquiry_Request) ProtoMessage() {}

func (x *Ftp_Inquiry_Request) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_Inquiry_Request.ProtoReflect.Descriptor instead.
func (*Ftp_Inquiry_Request) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 6, 0}
}

func (x *Ftp_Inquiry_Request) GetTransferSessId() uint64 {
	if x != nil {
		return x.TransferSessId
	}
	return 0
}

type Ftp_Inquiry_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State  Ftp_TransferState     `protobuf:"varint,1,opt,name=state,proto3,enum=go_serv.net.Ftp_TransferState" json:"state,omitempty"`
	Fds    []*Ftp_FileDescriptor `protobuf:"bytes,2,rep,name=fds,proto3" json:"fds,omitempty"`
	Chunks []*Ftp_PendingChunks  `protobuf:"bytes,3,rep,name=chunks,proto3" json:"chunks,omitempty"`
}

func (x *Ftp_Inquiry_Response) Reset() {
	*x = Ftp_Inquiry_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ftp_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ftp_Inquiry_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ftp_Inquiry_Response) ProtoMessage() {}

func (x *Ftp_Inquiry_Response) ProtoReflect() protoreflect.Message {
	mi := &file_net_ftp_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ftp_Inquiry_Response.ProtoReflect.Descriptor instead.
func (*Ftp_Inquiry_Response) Descriptor() ([]byte, []int) {
	return file_net_ftp_proto_rawDescGZIP(), []int{0, 6, 1}
}

func (x *Ftp_Inquiry_Response) GetState() Ftp_TransferState {
	if x != nil {
		return x.State
	}
	return Ftp_PENDING
}

func (x *Ftp_Inquiry_Response) GetFds() []*Ftp_FileDescriptor {
	if x != nil {
		return x.Fds
	}
	return nil
}

func (x *Ftp_Inquiry_Response) GetChunks() []*Ftp_PendingChunks {
	if x != nil {
		return x.Chunks
	}
	return nil
}

var File_net_ftp_proto protoreflect.FileDescriptor

var file_net_ftp_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x2f, 0x66, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x22, 0xbd, 0x0a, 0x0a,
	0x03, 0x46, 0x74, 0x70, 0x1a, 0xcc, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x1e, 0x0a, 0x08, 0x61, 0x62, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x61, 0x62, 0x73, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01,
	0x12, 0x20, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x62, 0x73,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x57, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x33, 0x0a, 0x09,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x1a, 0xbd, 0x02, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x1a, 0xb9, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x74, 0x65, 0x6d,
	0x70, 0x12, 0x29, 0x0a, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01,
	0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x73, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x6d, 0x61, 0x78, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x41,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65,
	0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x73, 0x1a, 0xbf, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a,
	0x70, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x06, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x1a, 0x40, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x1a, 0x64, 0x0a, 0x0d, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0xec, 0x01, 0x0a, 0x07, 0x49, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x1a, 0x33, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x49, 0x64, 0x1a, 0xab, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a,
	0x03, 0x66, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x66, 0x64, 0x73,
	0x12, 0x36, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46,
	0x74, 0x70, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x46, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53,
	0x53, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x42, 0x42, 0x5a, 0x40,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x6e, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_net_ftp_proto_rawDescOnce sync.Once
	file_net_ftp_proto_rawDescData = file_net_ftp_proto_rawDesc
)

func file_net_ftp_proto_rawDescGZIP() []byte {
	file_net_ftp_proto_rawDescOnce.Do(func() {
		file_net_ftp_proto_rawDescData = protoimpl.X.CompressGZIP(file_net_ftp_proto_rawDescData)
	})
	return file_net_ftp_proto_rawDescData
}

var file_net_ftp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_net_ftp_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_net_ftp_proto_goTypes = []interface{}{
	(Ftp_TransferState)(0),          // 0: go_serv.net.Ftp.TransferState
	(*Ftp)(nil),                     // 1: go_serv.net.Ftp
	(*Ftp_FileInfo)(nil),            // 2: go_serv.net.Ftp.FileInfo
	(*Ftp_FileDescriptor)(nil),      // 3: go_serv.net.Ftp.FileDescriptor
	(*Ftp_FileRange)(nil),           // 4: go_serv.net.Ftp.FileRange
	(*Ftp_NewSession)(nil),          // 5: go_serv.net.Ftp.NewSession
	(*Ftp_FileChunk)(nil),           // 6: go_serv.net.Ftp.FileChunk
	(*Ftp_PendingChunks)(nil),       // 7: go_serv.net.Ftp.PendingChunks
	(*Ftp_Inquiry)(nil),             // 8: go_serv.net.Ftp.Inquiry
	(*Ftp_NewSession_Request)(nil),  // 9: go_serv.net.Ftp.NewSession.Request
	(*Ftp_NewSession_Response)(nil), // 10: go_serv.net.Ftp.NewSession.Response
	(*Ftp_FileChunk_Request)(nil),   // 11: go_serv.net.Ftp.FileChunk.Request
	(*Ftp_FileChunk_Response)(nil),  // 12: go_serv.net.Ftp.FileChunk.Response
	(*Ftp_Inquiry_Request)(nil),     // 13: go_serv.net.Ftp.Inquiry.Request
	(*Ftp_Inquiry_Response)(nil),    // 14: go_serv.net.Ftp.Inquiry.Response
}
var file_net_ftp_proto_depIdxs = []int32{
	2, // 0: go_serv.net.Ftp.FileDescriptor.info:type_name -> go_serv.net.Ftp.FileInfo
	4, // 1: go_serv.net.Ftp.PendingChunks.ranges:type_name -> go_serv.net.Ftp.FileRange
	2, // 2: go_serv.net.Ftp.NewSession.Request.files:type_name -> go_serv.net.Ftp.FileInfo
	3, // 3: go_serv.net.Ftp.NewSession.Response.descriptors:type_name -> go_serv.net.Ftp.FileDescriptor
	4, // 4: go_serv.net.Ftp.FileChunk.Request.range:type_name -> go_serv.net.Ftp.FileRange
	0, // 5: go_serv.net.Ftp.FileChunk.Response.state:type_name -> go_serv.net.Ftp.TransferState
	0, // 6: go_serv.net.Ftp.Inquiry.Response.state:type_name -> go_serv.net.Ftp.TransferState
	3, // 7: go_serv.net.Ftp.Inquiry.Response.fds:type_name -> go_serv.net.Ftp.FileDescriptor
	7, // 8: go_serv.net.Ftp.Inquiry.Response.chunks:type_name -> go_serv.net.Ftp.PendingChunks
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_net_ftp_proto_init() }
func file_net_ftp_proto_init() {
	if File_net_ftp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_net_ftp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp); i {
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
		file_net_ftp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_FileInfo); i {
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
		file_net_ftp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_FileDescriptor); i {
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
		file_net_ftp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_FileRange); i {
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
		file_net_ftp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_NewSession); i {
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
		file_net_ftp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_FileChunk); i {
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
		file_net_ftp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_PendingChunks); i {
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
		file_net_ftp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_Inquiry); i {
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
		file_net_ftp_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_NewSession_Request); i {
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
		file_net_ftp_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_NewSession_Response); i {
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
		file_net_ftp_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_FileChunk_Request); i {
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
		file_net_ftp_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_FileChunk_Response); i {
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
		file_net_ftp_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_Inquiry_Request); i {
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
		file_net_ftp_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ftp_Inquiry_Response); i {
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
	file_net_ftp_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_net_ftp_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_net_ftp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_net_ftp_proto_goTypes,
		DependencyIndexes: file_net_ftp_proto_depIdxs,
		EnumInfos:         file_net_ftp_proto_enumTypes,
		MessageInfos:      file_net_ftp_proto_msgTypes,
	}.Build()
	File_net_ftp_proto = out.File
	file_net_ftp_proto_rawDesc = nil
	file_net_ftp_proto_goTypes = nil
	file_net_ftp_proto_depIdxs = nil
}
