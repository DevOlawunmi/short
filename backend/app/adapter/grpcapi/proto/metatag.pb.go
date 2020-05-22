// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/adapter/grpcapi/proto/metatag.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetOGTagsRequest struct {
	Alias                string   `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOGTagsRequest) Reset()         { *m = GetOGTagsRequest{} }
func (m *GetOGTagsRequest) String() string { return proto.CompactTextString(m) }
func (*GetOGTagsRequest) ProtoMessage()    {}
func (*GetOGTagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac72a7e9b365c60, []int{0}
}

func (m *GetOGTagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOGTagsRequest.Unmarshal(m, b)
}
func (m *GetOGTagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOGTagsRequest.Marshal(b, m, deterministic)
}
func (m *GetOGTagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOGTagsRequest.Merge(m, src)
}
func (m *GetOGTagsRequest) XXX_Size() int {
	return xxx_messageInfo_GetOGTagsRequest.Size(m)
}
func (m *GetOGTagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOGTagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOGTagsRequest proto.InternalMessageInfo

func (m *GetOGTagsRequest) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

type GetOGTagsResponse struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl             string   `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOGTagsResponse) Reset()         { *m = GetOGTagsResponse{} }
func (m *GetOGTagsResponse) String() string { return proto.CompactTextString(m) }
func (*GetOGTagsResponse) ProtoMessage()    {}
func (*GetOGTagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac72a7e9b365c60, []int{1}
}

func (m *GetOGTagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOGTagsResponse.Unmarshal(m, b)
}
func (m *GetOGTagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOGTagsResponse.Marshal(b, m, deterministic)
}
func (m *GetOGTagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOGTagsResponse.Merge(m, src)
}
func (m *GetOGTagsResponse) XXX_Size() int {
	return xxx_messageInfo_GetOGTagsResponse.Size(m)
}
func (m *GetOGTagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOGTagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOGTagsResponse proto.InternalMessageInfo

func (m *GetOGTagsResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GetOGTagsResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetOGTagsResponse) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type GetTwitterTagsRequest struct {
	Alias                string   `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTwitterTagsRequest) Reset()         { *m = GetTwitterTagsRequest{} }
func (m *GetTwitterTagsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTwitterTagsRequest) ProtoMessage()    {}
func (*GetTwitterTagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac72a7e9b365c60, []int{2}
}

func (m *GetTwitterTagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTwitterTagsRequest.Unmarshal(m, b)
}
func (m *GetTwitterTagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTwitterTagsRequest.Marshal(b, m, deterministic)
}
func (m *GetTwitterTagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTwitterTagsRequest.Merge(m, src)
}
func (m *GetTwitterTagsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTwitterTagsRequest.Size(m)
}
func (m *GetTwitterTagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTwitterTagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTwitterTagsRequest proto.InternalMessageInfo

func (m *GetTwitterTagsRequest) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

type GetTwitterTagsResponse struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl             string   `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTwitterTagsResponse) Reset()         { *m = GetTwitterTagsResponse{} }
func (m *GetTwitterTagsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTwitterTagsResponse) ProtoMessage()    {}
func (*GetTwitterTagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac72a7e9b365c60, []int{3}
}

func (m *GetTwitterTagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTwitterTagsResponse.Unmarshal(m, b)
}
func (m *GetTwitterTagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTwitterTagsResponse.Marshal(b, m, deterministic)
}
func (m *GetTwitterTagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTwitterTagsResponse.Merge(m, src)
}
func (m *GetTwitterTagsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTwitterTagsResponse.Size(m)
}
func (m *GetTwitterTagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTwitterTagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTwitterTagsResponse proto.InternalMessageInfo

func (m *GetTwitterTagsResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GetTwitterTagsResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetTwitterTagsResponse) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type UpdateOGTagsRequest struct {
	Alias                string   `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl             string   `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateOGTagsRequest) Reset()         { *m = UpdateOGTagsRequest{} }
func (m *UpdateOGTagsRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateOGTagsRequest) ProtoMessage()    {}
func (*UpdateOGTagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac72a7e9b365c60, []int{4}
}

func (m *UpdateOGTagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOGTagsRequest.Unmarshal(m, b)
}
func (m *UpdateOGTagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOGTagsRequest.Marshal(b, m, deterministic)
}
func (m *UpdateOGTagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOGTagsRequest.Merge(m, src)
}
func (m *UpdateOGTagsRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateOGTagsRequest.Size(m)
}
func (m *UpdateOGTagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOGTagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOGTagsRequest proto.InternalMessageInfo

func (m *UpdateOGTagsRequest) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *UpdateOGTagsRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateOGTagsRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateOGTagsRequest) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type UpdateOGTagsResponse struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl             string   `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateOGTagsResponse) Reset()         { *m = UpdateOGTagsResponse{} }
func (m *UpdateOGTagsResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateOGTagsResponse) ProtoMessage()    {}
func (*UpdateOGTagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac72a7e9b365c60, []int{5}
}

func (m *UpdateOGTagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOGTagsResponse.Unmarshal(m, b)
}
func (m *UpdateOGTagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOGTagsResponse.Marshal(b, m, deterministic)
}
func (m *UpdateOGTagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOGTagsResponse.Merge(m, src)
}
func (m *UpdateOGTagsResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateOGTagsResponse.Size(m)
}
func (m *UpdateOGTagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOGTagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOGTagsResponse proto.InternalMessageInfo

func (m *UpdateOGTagsResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateOGTagsResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateOGTagsResponse) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type UpdateTwitterTagsRequest struct {
	Alias                string   `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl             string   `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTwitterTagsRequest) Reset()         { *m = UpdateTwitterTagsRequest{} }
func (m *UpdateTwitterTagsRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTwitterTagsRequest) ProtoMessage()    {}
func (*UpdateTwitterTagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac72a7e9b365c60, []int{6}
}

func (m *UpdateTwitterTagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTwitterTagsRequest.Unmarshal(m, b)
}
func (m *UpdateTwitterTagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTwitterTagsRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTwitterTagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTwitterTagsRequest.Merge(m, src)
}
func (m *UpdateTwitterTagsRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTwitterTagsRequest.Size(m)
}
func (m *UpdateTwitterTagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTwitterTagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTwitterTagsRequest proto.InternalMessageInfo

func (m *UpdateTwitterTagsRequest) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *UpdateTwitterTagsRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateTwitterTagsRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateTwitterTagsRequest) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type UpdateTwitterTagsResponse struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl             string   `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTwitterTagsResponse) Reset()         { *m = UpdateTwitterTagsResponse{} }
func (m *UpdateTwitterTagsResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTwitterTagsResponse) ProtoMessage()    {}
func (*UpdateTwitterTagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac72a7e9b365c60, []int{7}
}

func (m *UpdateTwitterTagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTwitterTagsResponse.Unmarshal(m, b)
}
func (m *UpdateTwitterTagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTwitterTagsResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTwitterTagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTwitterTagsResponse.Merge(m, src)
}
func (m *UpdateTwitterTagsResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTwitterTagsResponse.Size(m)
}
func (m *UpdateTwitterTagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTwitterTagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTwitterTagsResponse proto.InternalMessageInfo

func (m *UpdateTwitterTagsResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateTwitterTagsResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateTwitterTagsResponse) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOGTagsRequest)(nil), "proto.GetOGTagsRequest")
	proto.RegisterType((*GetOGTagsResponse)(nil), "proto.GetOGTagsResponse")
	proto.RegisterType((*GetTwitterTagsRequest)(nil), "proto.GetTwitterTagsRequest")
	proto.RegisterType((*GetTwitterTagsResponse)(nil), "proto.GetTwitterTagsResponse")
	proto.RegisterType((*UpdateOGTagsRequest)(nil), "proto.UpdateOGTagsRequest")
	proto.RegisterType((*UpdateOGTagsResponse)(nil), "proto.UpdateOGTagsResponse")
	proto.RegisterType((*UpdateTwitterTagsRequest)(nil), "proto.UpdateTwitterTagsRequest")
	proto.RegisterType((*UpdateTwitterTagsResponse)(nil), "proto.UpdateTwitterTagsResponse")
}

func init() {
	proto.RegisterFile("app/adapter/grpcapi/proto/metatag.proto", fileDescriptor_7ac72a7e9b365c60)
}

var fileDescriptor_7ac72a7e9b365c60 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x2d, 0x88, 0x91, 0xd1, 0x10, 0x59, 0x51, 0x6b, 0xd1, 0x48, 0x7a, 0x91, 0x8b, 0x34,
	0xd1, 0x17, 0xf0, 0x46, 0x3c, 0x18, 0x12, 0x84, 0xc4, 0x9b, 0x19, 0x61, 0x52, 0x37, 0x16, 0xba,
	0xee, 0x0e, 0x7a, 0xf5, 0xe2, 0xbb, 0xf9, 0x58, 0xc6, 0xad, 0xfc, 0x2f, 0xc2, 0xa5, 0xa7, 0x66,
	0xf6, 0xfb, 0x76, 0xfa, 0x9b, 0xd9, 0x19, 0xb8, 0x44, 0xa5, 0x02, 0xec, 0xa3, 0x62, 0xd2, 0x41,
	0xa8, 0x55, 0x0f, 0x95, 0x0c, 0x94, 0x8e, 0x39, 0x0e, 0x06, 0xc4, 0xc8, 0x18, 0x36, 0x6c, 0x24,
	0x0a, 0xf6, 0xe3, 0xd7, 0xe1, 0xa0, 0x49, 0xdc, 0x6a, 0x76, 0x30, 0x34, 0x6d, 0x7a, 0x1b, 0x91,
	0x61, 0x51, 0x81, 0x02, 0x46, 0x12, 0x8d, 0xeb, 0xd4, 0x9c, 0x7a, 0xb1, 0x9d, 0x04, 0xfe, 0x0b,
	0x94, 0x67, 0x9c, 0x46, 0xc5, 0x43, 0x43, 0xbf, 0x56, 0x96, 0x1c, 0xd1, 0xd8, 0x6a, 0x03, 0x51,
	0x83, 0xbd, 0x3e, 0x99, 0x9e, 0x96, 0x8a, 0x65, 0x3c, 0x74, 0x73, 0x56, 0x9b, 0x3d, 0x12, 0x55,
	0x28, 0xca, 0x01, 0x86, 0xf4, 0x34, 0xd2, 0x91, 0x9b, 0xb7, 0xfa, 0xae, 0x3d, 0xe8, 0xea, 0xc8,
	0xbf, 0x82, 0xa3, 0x26, 0x71, 0xe7, 0x43, 0x32, 0x93, 0x5e, 0x0f, 0x36, 0x80, 0xe3, 0x45, 0x7b,
	0x96, 0x74, 0x9f, 0x0e, 0x1c, 0x76, 0x55, 0x1f, 0x99, 0x36, 0xe8, 0xda, 0x14, 0x21, 0xf7, 0x0f,
	0x42, 0x7e, 0x0d, 0xc2, 0xf6, 0x02, 0xc2, 0x2b, 0x54, 0xe6, 0x09, 0xb2, 0xac, 0xf7, 0xcb, 0x01,
	0x37, 0xf9, 0xdb, 0xa6, 0x2f, 0x92, 0x4d, 0xd1, 0x0a, 0x4e, 0x53, 0x30, 0x32, 0xac, 0xfc, 0xfa,
	0x3b, 0x07, 0xa5, 0x7b, 0x62, 0xec, 0x60, 0xf8, 0x40, 0xfa, 0x5d, 0xf6, 0x48, 0xdc, 0x42, 0x71,
	0xb2, 0x04, 0xe2, 0x24, 0x59, 0xa5, 0xc6, 0xe2, 0x02, 0x79, 0xee, 0xb2, 0x90, 0x70, 0xfa, 0x5b,
	0xa2, 0x05, 0xa5, 0xf9, 0x69, 0x15, 0x67, 0x53, 0xf7, 0x72, 0x87, 0xbd, 0xf3, 0x15, 0xea, 0x24,
	0xe1, 0x1d, 0xec, 0xcf, 0x0e, 0x83, 0xf0, 0xfe, 0x2e, 0xa4, 0xcc, 0xa8, 0x57, 0x4d, 0xd5, 0x26,
	0xa9, 0x1e, 0xa1, 0xbc, 0xd4, 0x62, 0x71, 0x31, 0x77, 0x27, 0x85, 0xb0, 0xb6, 0xda, 0x30, 0xce,
	0xfc, 0xbc, 0x63, 0x2d, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x63, 0xf8, 0xce, 0x9f,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetaTagServiceClient is the client API for MetaTagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetaTagServiceClient interface {
	GetOGTags(ctx context.Context, in *GetOGTagsRequest, opts ...grpc.CallOption) (*GetOGTagsResponse, error)
	GetTwitterTags(ctx context.Context, in *GetTwitterTagsRequest, opts ...grpc.CallOption) (*GetTwitterTagsResponse, error)
	UpdateOGTags(ctx context.Context, in *UpdateOGTagsRequest, opts ...grpc.CallOption) (*UpdateOGTagsResponse, error)
	UpdateTwitterTags(ctx context.Context, in *UpdateTwitterTagsRequest, opts ...grpc.CallOption) (*UpdateTwitterTagsResponse, error)
}

type metaTagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaTagServiceClient(cc grpc.ClientConnInterface) MetaTagServiceClient {
	return &metaTagServiceClient{cc}
}

func (c *metaTagServiceClient) GetOGTags(ctx context.Context, in *GetOGTagsRequest, opts ...grpc.CallOption) (*GetOGTagsResponse, error) {
	out := new(GetOGTagsResponse)
	err := c.cc.Invoke(ctx, "/proto.MetaTagService/GetOGTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaTagServiceClient) GetTwitterTags(ctx context.Context, in *GetTwitterTagsRequest, opts ...grpc.CallOption) (*GetTwitterTagsResponse, error) {
	out := new(GetTwitterTagsResponse)
	err := c.cc.Invoke(ctx, "/proto.MetaTagService/GetTwitterTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaTagServiceClient) UpdateOGTags(ctx context.Context, in *UpdateOGTagsRequest, opts ...grpc.CallOption) (*UpdateOGTagsResponse, error) {
	out := new(UpdateOGTagsResponse)
	err := c.cc.Invoke(ctx, "/proto.MetaTagService/UpdateOGTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaTagServiceClient) UpdateTwitterTags(ctx context.Context, in *UpdateTwitterTagsRequest, opts ...grpc.CallOption) (*UpdateTwitterTagsResponse, error) {
	out := new(UpdateTwitterTagsResponse)
	err := c.cc.Invoke(ctx, "/proto.MetaTagService/UpdateTwitterTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaTagServiceServer is the server API for MetaTagService service.
type MetaTagServiceServer interface {
	GetOGTags(context.Context, *GetOGTagsRequest) (*GetOGTagsResponse, error)
	GetTwitterTags(context.Context, *GetTwitterTagsRequest) (*GetTwitterTagsResponse, error)
	UpdateOGTags(context.Context, *UpdateOGTagsRequest) (*UpdateOGTagsResponse, error)
	UpdateTwitterTags(context.Context, *UpdateTwitterTagsRequest) (*UpdateTwitterTagsResponse, error)
}

// UnimplementedMetaTagServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetaTagServiceServer struct {
}

func (*UnimplementedMetaTagServiceServer) GetOGTags(ctx context.Context, req *GetOGTagsRequest) (*GetOGTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOGTags not implemented")
}
func (*UnimplementedMetaTagServiceServer) GetTwitterTags(ctx context.Context, req *GetTwitterTagsRequest) (*GetTwitterTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTwitterTags not implemented")
}
func (*UnimplementedMetaTagServiceServer) UpdateOGTags(ctx context.Context, req *UpdateOGTagsRequest) (*UpdateOGTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOGTags not implemented")
}
func (*UnimplementedMetaTagServiceServer) UpdateTwitterTags(ctx context.Context, req *UpdateTwitterTagsRequest) (*UpdateTwitterTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTwitterTags not implemented")
}

func RegisterMetaTagServiceServer(s *grpc.Server, srv MetaTagServiceServer) {
	s.RegisterService(&_MetaTagService_serviceDesc, srv)
}

func _MetaTagService_GetOGTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOGTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaTagServiceServer).GetOGTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetaTagService/GetOGTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaTagServiceServer).GetOGTags(ctx, req.(*GetOGTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaTagService_GetTwitterTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTwitterTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaTagServiceServer).GetTwitterTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetaTagService/GetTwitterTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaTagServiceServer).GetTwitterTags(ctx, req.(*GetTwitterTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaTagService_UpdateOGTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOGTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaTagServiceServer).UpdateOGTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetaTagService/UpdateOGTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaTagServiceServer).UpdateOGTags(ctx, req.(*UpdateOGTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaTagService_UpdateTwitterTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTwitterTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaTagServiceServer).UpdateTwitterTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetaTagService/UpdateTwitterTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaTagServiceServer).UpdateTwitterTags(ctx, req.(*UpdateTwitterTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetaTagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MetaTagService",
	HandlerType: (*MetaTagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOGTags",
			Handler:    _MetaTagService_GetOGTags_Handler,
		},
		{
			MethodName: "GetTwitterTags",
			Handler:    _MetaTagService_GetTwitterTags_Handler,
		},
		{
			MethodName: "UpdateOGTags",
			Handler:    _MetaTagService_UpdateOGTags_Handler,
		},
		{
			MethodName: "UpdateTwitterTags",
			Handler:    _MetaTagService_UpdateTwitterTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/adapter/grpcapi/proto/metatag.proto",
}