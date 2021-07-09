// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.2
// source: api/applications/applications.proto

package applications

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type AddApplicationRequest_DeploymentType int32

const (
	AddApplicationRequest_KUSTOMIZE AddApplicationRequest_DeploymentType = 0
	AddApplicationRequest_HELM      AddApplicationRequest_DeploymentType = 1
)

// Enum value maps for AddApplicationRequest_DeploymentType.
var (
	AddApplicationRequest_DeploymentType_name = map[int32]string{
		0: "KUSTOMIZE",
		1: "HELM",
	}
	AddApplicationRequest_DeploymentType_value = map[string]int32{
		"KUSTOMIZE": 0,
		"HELM":      1,
	}
)

func (x AddApplicationRequest_DeploymentType) Enum() *AddApplicationRequest_DeploymentType {
	p := new(AddApplicationRequest_DeploymentType)
	*p = x
	return p
}

func (x AddApplicationRequest_DeploymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddApplicationRequest_DeploymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_applications_applications_proto_enumTypes[0].Descriptor()
}

func (AddApplicationRequest_DeploymentType) Type() protoreflect.EnumType {
	return &file_api_applications_applications_proto_enumTypes[0]
}

func (x AddApplicationRequest_DeploymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddApplicationRequest_DeploymentType.Descriptor instead.
func (AddApplicationRequest_DeploymentType) EnumDescriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{5, 0}
}

type AddApplicationRequest_SourceType int32

const (
	AddApplicationRequest_GIT   AddApplicationRequest_SourceType = 0
	AddApplicationRequest_CHART AddApplicationRequest_SourceType = 1
)

// Enum value maps for AddApplicationRequest_SourceType.
var (
	AddApplicationRequest_SourceType_name = map[int32]string{
		0: "GIT",
		1: "CHART",
	}
	AddApplicationRequest_SourceType_value = map[string]int32{
		"GIT":   0,
		"CHART": 1,
	}
)

func (x AddApplicationRequest_SourceType) Enum() *AddApplicationRequest_SourceType {
	p := new(AddApplicationRequest_SourceType)
	*p = x
	return p
}

func (x AddApplicationRequest_SourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddApplicationRequest_SourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_applications_applications_proto_enumTypes[1].Descriptor()
}

func (AddApplicationRequest_SourceType) Type() protoreflect.EnumType {
	return &file_api_applications_applications_proto_enumTypes[1]
}

func (x AddApplicationRequest_SourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddApplicationRequest_SourceType.Descriptor instead.
func (AddApplicationRequest_SourceType) EnumDescriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{5, 1}
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // The name of the application
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"` // The file path where the k8s yaml files for this application are stored.
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`   // The git repository URL for this application
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_applications_applications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_api_applications_applications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Application) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListApplicationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"` // The namespace to look for applications
}

func (x *ListApplicationsRequest) Reset() {
	*x = ListApplicationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_applications_applications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationsRequest) ProtoMessage() {}

func (x *ListApplicationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_applications_applications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationsRequest.ProtoReflect.Descriptor instead.
func (*ListApplicationsRequest) Descriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{1}
}

func (x *ListApplicationsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ListApplicationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"` // A list of applications
}

func (x *ListApplicationsResponse) Reset() {
	*x = ListApplicationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_applications_applications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationsResponse) ProtoMessage() {}

func (x *ListApplicationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_applications_applications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationsResponse.ProtoReflect.Descriptor instead.
func (*ListApplicationsResponse) Descriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{2}
}

func (x *ListApplicationsResponse) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

type GetApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`           // The name of an application
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"` // The kubernetes namespace of the application. Default is `wego-system`
}

func (x *GetApplicationRequest) Reset() {
	*x = GetApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_applications_applications_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplicationRequest) ProtoMessage() {}

func (x *GetApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_applications_applications_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplicationRequest.ProtoReflect.Descriptor instead.
func (*GetApplicationRequest) Descriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{3}
}

func (x *GetApplicationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetApplicationRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GetApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *GetApplicationResponse) Reset() {
	*x = GetApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_applications_applications_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplicationResponse) ProtoMessage() {}

func (x *GetApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_applications_applications_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplicationResponse.ProtoReflect.Descriptor instead.
func (*GetApplicationResponse) Descriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{4}
}

func (x *GetApplicationResponse) GetApplication() *Application {
	if x != nil {
		return x.Application
	}
	return nil
}

type AddApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string                               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url            string                               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Path           string                               `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Branch         string                               `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	DeploymentType AddApplicationRequest_DeploymentType `protobuf:"varint,5,opt,name=deployment_type,json=deploymentType,proto3,enum=wego_server.v1.AddApplicationRequest_DeploymentType" json:"deployment_type,omitempty"`
	Chart          string                               `protobuf:"bytes,6,opt,name=chart,proto3" json:"chart,omitempty"`
	SourceType     AddApplicationRequest_SourceType     `protobuf:"varint,7,opt,name=source_type,json=sourceType,proto3,enum=wego_server.v1.AddApplicationRequest_SourceType" json:"source_type,omitempty"`
	AppConfigUrl   string                               `protobuf:"bytes,8,opt,name=app_config_url,json=appConfigUrl,proto3" json:"app_config_url,omitempty"`
	Namespace      string                               `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
	DryRun         bool                                 `protobuf:"varint,10,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	AutoMerge      bool                                 `protobuf:"varint,11,opt,name=auto_merge,json=autoMerge,proto3" json:"auto_merge,omitempty"`
}

func (x *AddApplicationRequest) Reset() {
	*x = AddApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_applications_applications_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddApplicationRequest) ProtoMessage() {}

func (x *AddApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_applications_applications_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddApplicationRequest.ProtoReflect.Descriptor instead.
func (*AddApplicationRequest) Descriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{5}
}

func (x *AddApplicationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddApplicationRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddApplicationRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AddApplicationRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *AddApplicationRequest) GetDeploymentType() AddApplicationRequest_DeploymentType {
	if x != nil {
		return x.DeploymentType
	}
	return AddApplicationRequest_KUSTOMIZE
}

func (x *AddApplicationRequest) GetChart() string {
	if x != nil {
		return x.Chart
	}
	return ""
}

func (x *AddApplicationRequest) GetSourceType() AddApplicationRequest_SourceType {
	if x != nil {
		return x.SourceType
	}
	return AddApplicationRequest_GIT
}

func (x *AddApplicationRequest) GetAppConfigUrl() string {
	if x != nil {
		return x.AppConfigUrl
	}
	return ""
}

func (x *AddApplicationRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AddApplicationRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *AddApplicationRequest) GetAutoMerge() bool {
	if x != nil {
		return x.AutoMerge
	}
	return false
}

type AddApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *AddApplicationResponse) Reset() {
	*x = AddApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_applications_applications_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddApplicationResponse) ProtoMessage() {}

func (x *AddApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_applications_applications_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddApplicationResponse.ProtoReflect.Descriptor instead.
func (*AddApplicationResponse) Descriptor() ([]byte, []int) {
	return file_api_applications_applications_proto_rawDescGZIP(), []int{6}
}

func (x *AddApplicationResponse) GetApplication() *Application {
	if x != nil {
		return x.Application
	}
	return nil
}

var File_api_applications_applications_proto protoreflect.FileDescriptor

var file_api_applications_applications_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x65, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x37, 0x0a, 0x17,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x5b, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x65, 0x67, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x49, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x57, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77,
	0x65, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfa, 0x03, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x12, 0x5d, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x77, 0x65,
	0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x12, 0x51, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x77,
	0x65, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x70,
	0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x72, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x75, 0x74,
	0x6f, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x49, 0x5a, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x4c, 0x4d, 0x10,
	0x01, 0x22, 0x20, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x47, 0x49, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x48, 0x41, 0x52,
	0x54, 0x10, 0x01, 0x22, 0x57, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x65, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x8d, 0x03, 0x0a,
	0x0c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x7f, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x27, 0x2e, 0x77, 0x65, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x77, 0x65, 0x67,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x80,
	0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x2e, 0x77, 0x65, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x77, 0x65, 0x67, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x12, 0x79, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x77, 0x65, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x77, 0x65, 0x67,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0xce, 0x01, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2d, 0x67, 0x69, 0x74,
	0x6f, 0x70, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x92, 0x41, 0x8e, 0x01, 0x12,
	0x68, 0x0a, 0x15, 0x57, 0x65, 0x47, 0x6f, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x20, 0x41, 0x50, 0x49, 0x12, 0x4a, 0x54, 0x68, 0x65, 0x20, 0x57, 0x65,
	0x47, 0x6f, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x20, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x57, 0x65, 0x61, 0x76, 0x65,
	0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_applications_applications_proto_rawDescOnce sync.Once
	file_api_applications_applications_proto_rawDescData = file_api_applications_applications_proto_rawDesc
)

func file_api_applications_applications_proto_rawDescGZIP() []byte {
	file_api_applications_applications_proto_rawDescOnce.Do(func() {
		file_api_applications_applications_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_applications_applications_proto_rawDescData)
	})
	return file_api_applications_applications_proto_rawDescData
}

var file_api_applications_applications_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_applications_applications_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_applications_applications_proto_goTypes = []interface{}{
	(AddApplicationRequest_DeploymentType)(0), // 0: wego_server.v1.AddApplicationRequest.DeploymentType
	(AddApplicationRequest_SourceType)(0),     // 1: wego_server.v1.AddApplicationRequest.SourceType
	(*Application)(nil),                       // 2: wego_server.v1.Application
	(*ListApplicationsRequest)(nil),           // 3: wego_server.v1.ListApplicationsRequest
	(*ListApplicationsResponse)(nil),          // 4: wego_server.v1.ListApplicationsResponse
	(*GetApplicationRequest)(nil),             // 5: wego_server.v1.GetApplicationRequest
	(*GetApplicationResponse)(nil),            // 6: wego_server.v1.GetApplicationResponse
	(*AddApplicationRequest)(nil),             // 7: wego_server.v1.AddApplicationRequest
	(*AddApplicationResponse)(nil),            // 8: wego_server.v1.AddApplicationResponse
}
var file_api_applications_applications_proto_depIdxs = []int32{
	2, // 0: wego_server.v1.ListApplicationsResponse.applications:type_name -> wego_server.v1.Application
	2, // 1: wego_server.v1.GetApplicationResponse.application:type_name -> wego_server.v1.Application
	0, // 2: wego_server.v1.AddApplicationRequest.deployment_type:type_name -> wego_server.v1.AddApplicationRequest.DeploymentType
	1, // 3: wego_server.v1.AddApplicationRequest.source_type:type_name -> wego_server.v1.AddApplicationRequest.SourceType
	2, // 4: wego_server.v1.AddApplicationResponse.application:type_name -> wego_server.v1.Application
	3, // 5: wego_server.v1.Applications.ListApplications:input_type -> wego_server.v1.ListApplicationsRequest
	5, // 6: wego_server.v1.Applications.GetApplication:input_type -> wego_server.v1.GetApplicationRequest
	7, // 7: wego_server.v1.Applications.AddApplication:input_type -> wego_server.v1.AddApplicationRequest
	4, // 8: wego_server.v1.Applications.ListApplications:output_type -> wego_server.v1.ListApplicationsResponse
	6, // 9: wego_server.v1.Applications.GetApplication:output_type -> wego_server.v1.GetApplicationResponse
	8, // 10: wego_server.v1.Applications.AddApplication:output_type -> wego_server.v1.AddApplicationResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_applications_applications_proto_init() }
func file_api_applications_applications_proto_init() {
	if File_api_applications_applications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_applications_applications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_api_applications_applications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationsRequest); i {
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
		file_api_applications_applications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationsResponse); i {
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
		file_api_applications_applications_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplicationRequest); i {
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
		file_api_applications_applications_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplicationResponse); i {
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
		file_api_applications_applications_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddApplicationRequest); i {
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
		file_api_applications_applications_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddApplicationResponse); i {
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
			RawDescriptor: file_api_applications_applications_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_applications_applications_proto_goTypes,
		DependencyIndexes: file_api_applications_applications_proto_depIdxs,
		EnumInfos:         file_api_applications_applications_proto_enumTypes,
		MessageInfos:      file_api_applications_applications_proto_msgTypes,
	}.Build()
	File_api_applications_applications_proto = out.File
	file_api_applications_applications_proto_rawDesc = nil
	file_api_applications_applications_proto_goTypes = nil
	file_api_applications_applications_proto_depIdxs = nil
}
