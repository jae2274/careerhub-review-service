// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: careerhub/review_service/crawler/crawler_grpc/review.proto

package crawler_grpc

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

type GetCrawlingTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site string `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
}

func (x *GetCrawlingTasksRequest) Reset() {
	*x = GetCrawlingTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrawlingTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrawlingTasksRequest) ProtoMessage() {}

func (x *GetCrawlingTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrawlingTasksRequest.ProtoReflect.Descriptor instead.
func (*GetCrawlingTasksRequest) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{0}
}

func (x *GetCrawlingTasksRequest) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

type GetCrawlingTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyNames []string `protobuf:"bytes,1,rep,name=companyNames,proto3" json:"companyNames,omitempty"`
}

func (x *GetCrawlingTasksResponse) Reset() {
	*x = GetCrawlingTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrawlingTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrawlingTasksResponse) ProtoMessage() {}

func (x *GetCrawlingTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrawlingTasksResponse.ProtoReflect.Descriptor instead.
func (*GetCrawlingTasksResponse) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{1}
}

func (x *GetCrawlingTasksResponse) GetCompanyNames() []string {
	if x != nil {
		return x.CompanyNames
	}
	return nil
}

type SetReviewScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site        string `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	CompanyName string `protobuf:"bytes,2,opt,name=companyName,proto3" json:"companyName,omitempty"`
	AvgScore    int32  `protobuf:"varint,3,opt,name=avgScore,proto3" json:"avgScore,omitempty"`
	ReviewCount int32  `protobuf:"varint,4,opt,name=reviewCount,proto3" json:"reviewCount,omitempty"`
}

func (x *SetReviewScoreRequest) Reset() {
	*x = SetReviewScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReviewScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReviewScoreRequest) ProtoMessage() {}

func (x *SetReviewScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReviewScoreRequest.ProtoReflect.Descriptor instead.
func (*SetReviewScoreRequest) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{2}
}

func (x *SetReviewScoreRequest) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *SetReviewScoreRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *SetReviewScoreRequest) GetAvgScore() int32 {
	if x != nil {
		return x.AvgScore
	}
	return 0
}

func (x *SetReviewScoreRequest) GetReviewCount() int32 {
	if x != nil {
		return x.ReviewCount
	}
	return 0
}

type SetNotExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site        string `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	CompanyName string `protobuf:"bytes,2,opt,name=companyName,proto3" json:"companyName,omitempty"`
}

func (x *SetNotExistRequest) Reset() {
	*x = SetNotExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNotExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNotExistRequest) ProtoMessage() {}

func (x *SetNotExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNotExistRequest.ProtoReflect.Descriptor instead.
func (*SetNotExistRequest) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{3}
}

func (x *SetNotExistRequest) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *SetNotExistRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

type GetCrawlingTargetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site string `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
}

func (x *GetCrawlingTargetsRequest) Reset() {
	*x = GetCrawlingTargetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrawlingTargetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrawlingTargetsRequest) ProtoMessage() {}

func (x *GetCrawlingTargetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrawlingTargetsRequest.ProtoReflect.Descriptor instead.
func (*GetCrawlingTargetsRequest) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{4}
}

func (x *GetCrawlingTargetsRequest) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

type GetCrawlingTargetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyNames []string `protobuf:"bytes,1,rep,name=companyNames,proto3" json:"companyNames,omitempty"`
}

func (x *GetCrawlingTargetsResponse) Reset() {
	*x = GetCrawlingTargetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrawlingTargetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrawlingTargetsResponse) ProtoMessage() {}

func (x *GetCrawlingTargetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrawlingTargetsResponse.ProtoReflect.Descriptor instead.
func (*GetCrawlingTargetsResponse) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{5}
}

func (x *GetCrawlingTargetsResponse) GetCompanyNames() []string {
	if x != nil {
		return x.CompanyNames
	}
	return nil
}

type SaveCompanyReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site        string    `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	CompanyName string    `protobuf:"bytes,2,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Reviews     []*Review `protobuf:"bytes,4,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *SaveCompanyReviewsRequest) Reset() {
	*x = SaveCompanyReviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveCompanyReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveCompanyReviewsRequest) ProtoMessage() {}

func (x *SaveCompanyReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveCompanyReviewsRequest.ProtoReflect.Descriptor instead.
func (*SaveCompanyReviewsRequest) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{6}
}

func (x *SaveCompanyReviewsRequest) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *SaveCompanyReviewsRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *SaveCompanyReviewsRequest) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score            int32  `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	Summary          string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	EmploymentStatus bool   `protobuf:"varint,3,opt,name=employmentStatus,proto3" json:"employmentStatus,omitempty"`
	ReviewUserId     string `protobuf:"bytes,4,opt,name=reviewUserId,proto3" json:"reviewUserId,omitempty"`
	JobType          string `protobuf:"bytes,5,opt,name=jobType,proto3" json:"jobType,omitempty"`
	UnixMilli        int64  `protobuf:"varint,6,opt,name=unixMilli,proto3" json:"unixMilli,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{7}
}

func (x *Review) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Review) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Review) GetEmploymentStatus() bool {
	if x != nil {
		return x.EmploymentStatus
	}
	return false
}

func (x *Review) GetReviewUserId() string {
	if x != nil {
		return x.ReviewUserId
	}
	return ""
}

func (x *Review) GetJobType() string {
	if x != nil {
		return x.JobType
	}
	return ""
}

func (x *Review) GetUnixMilli() int64 {
	if x != nil {
		return x.UnixMilli
	}
	return 0
}

type SaveCompanyReviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InsertedCount int32 `protobuf:"varint,1,opt,name=insertedCount,proto3" json:"insertedCount,omitempty"`
}

func (x *SaveCompanyReviewsResponse) Reset() {
	*x = SaveCompanyReviewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveCompanyReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveCompanyReviewsResponse) ProtoMessage() {}

func (x *SaveCompanyReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveCompanyReviewsResponse.ProtoReflect.Descriptor instead.
func (*SaveCompanyReviewsResponse) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{8}
}

func (x *SaveCompanyReviewsResponse) GetInsertedCount() int32 {
	if x != nil {
		return x.InsertedCount
	}
	return 0
}

type FinishCrawlingTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site        string `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	CompanyName string `protobuf:"bytes,2,opt,name=companyName,proto3" json:"companyName,omitempty"`
}

func (x *FinishCrawlingTaskRequest) Reset() {
	*x = FinishCrawlingTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishCrawlingTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishCrawlingTaskRequest) ProtoMessage() {}

func (x *FinishCrawlingTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishCrawlingTaskRequest.ProtoReflect.Descriptor instead.
func (*FinishCrawlingTaskRequest) Descriptor() ([]byte, []int) {
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP(), []int{9}
}

func (x *FinishCrawlingTaskRequest) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *FinishCrawlingTaskRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

var File_careerhub_review_service_crawler_crawler_grpc_review_proto protoreflect.FileDescriptor

var file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c,
	0x65, 0x72, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x22,
	0x3e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x8b, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x76, 0x67, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x61, 0x76, 0x67, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x22, 0x40, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x9a, 0x01, 0x0a,
	0x19, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x47, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x6e, 0x69, 0x78, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x78, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x22, 0x42, 0x0a, 0x1a,
	0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x51, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x32, 0x94, 0x06, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x47, 0x72,
	0x70, 0x63, 0x12, 0x93, 0x01, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x3e, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72,
	0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72,
	0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0e, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x3c, 0x2e, 0x63, 0x61, 0x72,
	0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x60, 0x0a, 0x0b, 0x73, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12,
	0x39, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x99, 0x01, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x40, 0x2e, 0x63, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x99,
	0x01, 0x0a, 0x12, 0x73, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x40, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75,
	0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72,
	0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x12, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x40, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x27, 0x5a, 0x25, 0x63, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescOnce sync.Once
	file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescData = file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDesc
)

func file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescGZIP() []byte {
	file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescOnce.Do(func() {
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescData)
	})
	return file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDescData
}

var file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_careerhub_review_service_crawler_crawler_grpc_review_proto_goTypes = []interface{}{
	(*GetCrawlingTasksRequest)(nil),    // 0: careerhub.review_service.crawler_grpc.GetCrawlingTasksRequest
	(*GetCrawlingTasksResponse)(nil),   // 1: careerhub.review_service.crawler_grpc.GetCrawlingTasksResponse
	(*SetReviewScoreRequest)(nil),      // 2: careerhub.review_service.crawler_grpc.SetReviewScoreRequest
	(*SetNotExistRequest)(nil),         // 3: careerhub.review_service.crawler_grpc.SetNotExistRequest
	(*GetCrawlingTargetsRequest)(nil),  // 4: careerhub.review_service.crawler_grpc.GetCrawlingTargetsRequest
	(*GetCrawlingTargetsResponse)(nil), // 5: careerhub.review_service.crawler_grpc.GetCrawlingTargetsResponse
	(*SaveCompanyReviewsRequest)(nil),  // 6: careerhub.review_service.crawler_grpc.SaveCompanyReviewsRequest
	(*Review)(nil),                     // 7: careerhub.review_service.crawler_grpc.Review
	(*SaveCompanyReviewsResponse)(nil), // 8: careerhub.review_service.crawler_grpc.SaveCompanyReviewsResponse
	(*FinishCrawlingTaskRequest)(nil),  // 9: careerhub.review_service.crawler_grpc.FinishCrawlingTaskRequest
	(*emptypb.Empty)(nil),              // 10: google.protobuf.Empty
}
var file_careerhub_review_service_crawler_crawler_grpc_review_proto_depIdxs = []int32{
	7,  // 0: careerhub.review_service.crawler_grpc.SaveCompanyReviewsRequest.reviews:type_name -> careerhub.review_service.crawler_grpc.Review
	0,  // 1: careerhub.review_service.crawler_grpc.ReviewGrpc.getCrawlingTasks:input_type -> careerhub.review_service.crawler_grpc.GetCrawlingTasksRequest
	2,  // 2: careerhub.review_service.crawler_grpc.ReviewGrpc.setReviewScore:input_type -> careerhub.review_service.crawler_grpc.SetReviewScoreRequest
	3,  // 3: careerhub.review_service.crawler_grpc.ReviewGrpc.setNotExist:input_type -> careerhub.review_service.crawler_grpc.SetNotExistRequest
	4,  // 4: careerhub.review_service.crawler_grpc.ReviewGrpc.getCrawlingTargets:input_type -> careerhub.review_service.crawler_grpc.GetCrawlingTargetsRequest
	6,  // 5: careerhub.review_service.crawler_grpc.ReviewGrpc.saveCompanyReviews:input_type -> careerhub.review_service.crawler_grpc.SaveCompanyReviewsRequest
	9,  // 6: careerhub.review_service.crawler_grpc.ReviewGrpc.finishCrawlingTask:input_type -> careerhub.review_service.crawler_grpc.FinishCrawlingTaskRequest
	1,  // 7: careerhub.review_service.crawler_grpc.ReviewGrpc.getCrawlingTasks:output_type -> careerhub.review_service.crawler_grpc.GetCrawlingTasksResponse
	10, // 8: careerhub.review_service.crawler_grpc.ReviewGrpc.setReviewScore:output_type -> google.protobuf.Empty
	10, // 9: careerhub.review_service.crawler_grpc.ReviewGrpc.setNotExist:output_type -> google.protobuf.Empty
	5,  // 10: careerhub.review_service.crawler_grpc.ReviewGrpc.getCrawlingTargets:output_type -> careerhub.review_service.crawler_grpc.GetCrawlingTargetsResponse
	8,  // 11: careerhub.review_service.crawler_grpc.ReviewGrpc.saveCompanyReviews:output_type -> careerhub.review_service.crawler_grpc.SaveCompanyReviewsResponse
	10, // 12: careerhub.review_service.crawler_grpc.ReviewGrpc.finishCrawlingTask:output_type -> google.protobuf.Empty
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_careerhub_review_service_crawler_crawler_grpc_review_proto_init() }
func file_careerhub_review_service_crawler_crawler_grpc_review_proto_init() {
	if File_careerhub_review_service_crawler_crawler_grpc_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrawlingTasksRequest); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrawlingTasksResponse); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetReviewScoreRequest); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNotExistRequest); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrawlingTargetsRequest); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrawlingTargetsResponse); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveCompanyReviewsRequest); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveCompanyReviewsResponse); i {
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
		file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishCrawlingTaskRequest); i {
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
			RawDescriptor: file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_careerhub_review_service_crawler_crawler_grpc_review_proto_goTypes,
		DependencyIndexes: file_careerhub_review_service_crawler_crawler_grpc_review_proto_depIdxs,
		MessageInfos:      file_careerhub_review_service_crawler_crawler_grpc_review_proto_msgTypes,
	}.Build()
	File_careerhub_review_service_crawler_crawler_grpc_review_proto = out.File
	file_careerhub_review_service_crawler_crawler_grpc_review_proto_rawDesc = nil
	file_careerhub_review_service_crawler_crawler_grpc_review_proto_goTypes = nil
	file_careerhub_review_service_crawler_crawler_grpc_review_proto_depIdxs = nil
}
