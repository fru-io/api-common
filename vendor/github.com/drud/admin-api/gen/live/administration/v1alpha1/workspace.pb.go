// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.0
// source: live/administration/v1alpha1/workspace.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListWorkspaceRequest_ListWorkspaceScope int32

const (
	ListWorkspaceRequest_DEVELOPER ListWorkspaceRequest_ListWorkspaceScope = 0
	ListWorkspaceRequest_ADMIN     ListWorkspaceRequest_ListWorkspaceScope = 1
)

// Enum value maps for ListWorkspaceRequest_ListWorkspaceScope.
var (
	ListWorkspaceRequest_ListWorkspaceScope_name = map[int32]string{
		0: "DEVELOPER",
		1: "ADMIN",
	}
	ListWorkspaceRequest_ListWorkspaceScope_value = map[string]int32{
		"DEVELOPER": 0,
		"ADMIN":     1,
	}
)

func (x ListWorkspaceRequest_ListWorkspaceScope) Enum() *ListWorkspaceRequest_ListWorkspaceScope {
	p := new(ListWorkspaceRequest_ListWorkspaceScope)
	*p = x
	return p
}

func (x ListWorkspaceRequest_ListWorkspaceScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListWorkspaceRequest_ListWorkspaceScope) Descriptor() protoreflect.EnumDescriptor {
	return file_live_administration_v1alpha1_workspace_proto_enumTypes[0].Descriptor()
}

func (ListWorkspaceRequest_ListWorkspaceScope) Type() protoreflect.EnumType {
	return &file_live_administration_v1alpha1_workspace_proto_enumTypes[0]
}

func (x ListWorkspaceRequest_ListWorkspaceScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListWorkspaceRequest_ListWorkspaceScope.Descriptor instead.
func (ListWorkspaceRequest_ListWorkspaceScope) EnumDescriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{1, 0}
}

type Workspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Admins       []string `protobuf:"bytes,2,rep,name=admins,proto3" json:"admins,omitempty"`
	Developers   []string `protobuf:"bytes,3,rep,name=developers,proto3" json:"developers,omitempty"`
	Subscription string   `protobuf:"bytes,4,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *Workspace) Reset() {
	*x = Workspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace) ProtoMessage() {}

func (x *Workspace) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace.ProtoReflect.Descriptor instead.
func (*Workspace) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{0}
}

func (x *Workspace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workspace) GetAdmins() []string {
	if x != nil {
		return x.Admins
	}
	return nil
}

func (x *Workspace) GetDevelopers() []string {
	if x != nil {
		return x.Developers
	}
	return nil
}

func (x *Workspace) GetSubscription() string {
	if x != nil {
		return x.Subscription
	}
	return ""
}

type ListWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope ListWorkspaceRequest_ListWorkspaceScope `protobuf:"varint,1,opt,name=Scope,proto3,enum=ddev.administration.v1alpha1.ListWorkspaceRequest_ListWorkspaceScope" json:"Scope,omitempty"`
}

func (x *ListWorkspaceRequest) Reset() {
	*x = ListWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspaceRequest) ProtoMessage() {}

func (x *ListWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*ListWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{1}
}

func (x *ListWorkspaceRequest) GetScope() ListWorkspaceRequest_ListWorkspaceScope {
	if x != nil {
		return x.Scope
	}
	return ListWorkspaceRequest_DEVELOPER
}

type ListWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaces []*Workspace `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty"`
}

func (x *ListWorkspaceResponse) Reset() {
	*x = ListWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspaceResponse) ProtoMessage() {}

func (x *ListWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*ListWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{2}
}

func (x *ListWorkspaceResponse) GetWorkspaces() []*Workspace {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

type AddWorkspaceAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace string `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AddWorkspaceAdminRequest) Reset() {
	*x = AddWorkspaceAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWorkspaceAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWorkspaceAdminRequest) ProtoMessage() {}

func (x *AddWorkspaceAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWorkspaceAdminRequest.ProtoReflect.Descriptor instead.
func (*AddWorkspaceAdminRequest) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{3}
}

func (x *AddWorkspaceAdminRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *AddWorkspaceAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AddWorkspaceAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *AddWorkspaceAdminResponse) Reset() {
	*x = AddWorkspaceAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWorkspaceAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWorkspaceAdminResponse) ProtoMessage() {}

func (x *AddWorkspaceAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWorkspaceAdminResponse.ProtoReflect.Descriptor instead.
func (*AddWorkspaceAdminResponse) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{4}
}

func (x *AddWorkspaceAdminResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

type AddWorkspaceDeveloperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace string `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AddWorkspaceDeveloperRequest) Reset() {
	*x = AddWorkspaceDeveloperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWorkspaceDeveloperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWorkspaceDeveloperRequest) ProtoMessage() {}

func (x *AddWorkspaceDeveloperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWorkspaceDeveloperRequest.ProtoReflect.Descriptor instead.
func (*AddWorkspaceDeveloperRequest) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{5}
}

func (x *AddWorkspaceDeveloperRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *AddWorkspaceDeveloperRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AddWorkspaceDeveloperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *AddWorkspaceDeveloperResponse) Reset() {
	*x = AddWorkspaceDeveloperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWorkspaceDeveloperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWorkspaceDeveloperResponse) ProtoMessage() {}

func (x *AddWorkspaceDeveloperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWorkspaceDeveloperResponse.ProtoReflect.Descriptor instead.
func (*AddWorkspaceDeveloperResponse) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{6}
}

func (x *AddWorkspaceDeveloperResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

type DeleteWorkspaceAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace string `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DeleteWorkspaceAdminRequest) Reset() {
	*x = DeleteWorkspaceAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceAdminRequest) ProtoMessage() {}

func (x *DeleteWorkspaceAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceAdminRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceAdminRequest) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteWorkspaceAdminRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *DeleteWorkspaceAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DeleteWorkspaceAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *DeleteWorkspaceAdminResponse) Reset() {
	*x = DeleteWorkspaceAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceAdminResponse) ProtoMessage() {}

func (x *DeleteWorkspaceAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceAdminResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceAdminResponse) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteWorkspaceAdminResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

type DeleteWorkspaceDeveloperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace string `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DeleteWorkspaceDeveloperRequest) Reset() {
	*x = DeleteWorkspaceDeveloperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceDeveloperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceDeveloperRequest) ProtoMessage() {}

func (x *DeleteWorkspaceDeveloperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceDeveloperRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceDeveloperRequest) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteWorkspaceDeveloperRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *DeleteWorkspaceDeveloperRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DeleteWorkspaceDeveloperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *DeleteWorkspaceDeveloperResponse) Reset() {
	*x = DeleteWorkspaceDeveloperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceDeveloperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceDeveloperResponse) ProtoMessage() {}

func (x *DeleteWorkspaceDeveloperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_live_administration_v1alpha1_workspace_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceDeveloperResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceDeveloperResponse) Descriptor() ([]byte, []int) {
	return file_live_administration_v1alpha1_workspace_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteWorkspaceDeveloperResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

var File_live_administration_v1alpha1_workspace_proto protoreflect.FileDescriptor

var file_live_administration_v1alpha1_workspace_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x64, 0x64, 0x65, 0x76, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x7b, 0x0a, 0x09,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x5b, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x45, 0x2e, 0x64, 0x64, 0x65, 0x76, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x22,
	0x2e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50,
	0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x22,
	0x60, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64,
	0x64, 0x65, 0x76, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x22, 0x4e, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x62, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45,
	0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x64, 0x64, 0x65, 0x76, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x52, 0x0a, 0x1c, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x66, 0x0a, 0x1d, 0x41, 0x64, 0x64,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x64, 0x64, 0x65, 0x76, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x51, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x65, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x64, 0x65, 0x76, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x55, 0x0a, 0x1f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x69, 0x0a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x64, 0x65, 0x76,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x73, 0x0a,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x75, 0x64, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_live_administration_v1alpha1_workspace_proto_rawDescOnce sync.Once
	file_live_administration_v1alpha1_workspace_proto_rawDescData = file_live_administration_v1alpha1_workspace_proto_rawDesc
)

func file_live_administration_v1alpha1_workspace_proto_rawDescGZIP() []byte {
	file_live_administration_v1alpha1_workspace_proto_rawDescOnce.Do(func() {
		file_live_administration_v1alpha1_workspace_proto_rawDescData = protoimpl.X.CompressGZIP(file_live_administration_v1alpha1_workspace_proto_rawDescData)
	})
	return file_live_administration_v1alpha1_workspace_proto_rawDescData
}

var file_live_administration_v1alpha1_workspace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_live_administration_v1alpha1_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_live_administration_v1alpha1_workspace_proto_goTypes = []interface{}{
	(ListWorkspaceRequest_ListWorkspaceScope)(0), // 0: ddev.administration.v1alpha1.ListWorkspaceRequest.ListWorkspaceScope
	(*Workspace)(nil),                        // 1: ddev.administration.v1alpha1.Workspace
	(*ListWorkspaceRequest)(nil),             // 2: ddev.administration.v1alpha1.ListWorkspaceRequest
	(*ListWorkspaceResponse)(nil),            // 3: ddev.administration.v1alpha1.ListWorkspaceResponse
	(*AddWorkspaceAdminRequest)(nil),         // 4: ddev.administration.v1alpha1.AddWorkspaceAdminRequest
	(*AddWorkspaceAdminResponse)(nil),        // 5: ddev.administration.v1alpha1.AddWorkspaceAdminResponse
	(*AddWorkspaceDeveloperRequest)(nil),     // 6: ddev.administration.v1alpha1.AddWorkspaceDeveloperRequest
	(*AddWorkspaceDeveloperResponse)(nil),    // 7: ddev.administration.v1alpha1.AddWorkspaceDeveloperResponse
	(*DeleteWorkspaceAdminRequest)(nil),      // 8: ddev.administration.v1alpha1.DeleteWorkspaceAdminRequest
	(*DeleteWorkspaceAdminResponse)(nil),     // 9: ddev.administration.v1alpha1.DeleteWorkspaceAdminResponse
	(*DeleteWorkspaceDeveloperRequest)(nil),  // 10: ddev.administration.v1alpha1.DeleteWorkspaceDeveloperRequest
	(*DeleteWorkspaceDeveloperResponse)(nil), // 11: ddev.administration.v1alpha1.DeleteWorkspaceDeveloperResponse
}
var file_live_administration_v1alpha1_workspace_proto_depIdxs = []int32{
	0, // 0: ddev.administration.v1alpha1.ListWorkspaceRequest.Scope:type_name -> ddev.administration.v1alpha1.ListWorkspaceRequest.ListWorkspaceScope
	1, // 1: ddev.administration.v1alpha1.ListWorkspaceResponse.workspaces:type_name -> ddev.administration.v1alpha1.Workspace
	1, // 2: ddev.administration.v1alpha1.AddWorkspaceAdminResponse.workspace:type_name -> ddev.administration.v1alpha1.Workspace
	1, // 3: ddev.administration.v1alpha1.AddWorkspaceDeveloperResponse.workspace:type_name -> ddev.administration.v1alpha1.Workspace
	1, // 4: ddev.administration.v1alpha1.DeleteWorkspaceAdminResponse.workspace:type_name -> ddev.administration.v1alpha1.Workspace
	1, // 5: ddev.administration.v1alpha1.DeleteWorkspaceDeveloperResponse.workspace:type_name -> ddev.administration.v1alpha1.Workspace
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_live_administration_v1alpha1_workspace_proto_init() }
func file_live_administration_v1alpha1_workspace_proto_init() {
	if File_live_administration_v1alpha1_workspace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_live_administration_v1alpha1_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkspaceRequest); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkspaceResponse); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWorkspaceAdminRequest); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWorkspaceAdminResponse); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWorkspaceDeveloperRequest); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWorkspaceDeveloperResponse); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceAdminRequest); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceAdminResponse); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceDeveloperRequest); i {
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
		file_live_administration_v1alpha1_workspace_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceDeveloperResponse); i {
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
			RawDescriptor: file_live_administration_v1alpha1_workspace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_live_administration_v1alpha1_workspace_proto_goTypes,
		DependencyIndexes: file_live_administration_v1alpha1_workspace_proto_depIdxs,
		EnumInfos:         file_live_administration_v1alpha1_workspace_proto_enumTypes,
		MessageInfos:      file_live_administration_v1alpha1_workspace_proto_msgTypes,
	}.Build()
	File_live_administration_v1alpha1_workspace_proto = out.File
	file_live_administration_v1alpha1_workspace_proto_rawDesc = nil
	file_live_administration_v1alpha1_workspace_proto_goTypes = nil
	file_live_administration_v1alpha1_workspace_proto_depIdxs = nil
}
