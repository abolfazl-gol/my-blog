// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/post.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	math "math"
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

type Post struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BlogId               int64    `protobuf:"varint,2,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Slug                 string   `protobuf:"bytes,6,opt,name=slug,proto3" json:"slug,omitempty"`
	Likes                int32    `protobuf:"varint,7,opt,name=likes,proto3" json:"likes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetLikes() int32 {
	if m != nil {
		return m.Likes
	}
	return 0
}

type ListPostsRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	SortBy               string   `protobuf:"bytes,2,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	OrderBy              string   `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	BlogId               int64    `protobuf:"varint,4,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsRequest) Reset()         { *m = ListPostsRequest{} }
func (m *ListPostsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPostsRequest) ProtoMessage()    {}
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{1}
}

func (m *ListPostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsRequest.Unmarshal(m, b)
}
func (m *ListPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsRequest.Marshal(b, m, deterministic)
}
func (m *ListPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsRequest.Merge(m, src)
}
func (m *ListPostsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPostsRequest.Size(m)
}
func (m *ListPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsRequest proto.InternalMessageInfo

func (m *ListPostsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPostsRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *ListPostsRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListPostsRequest) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

type ListPostsResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	NextPage             int32    `protobuf:"varint,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsResponse) Reset()         { *m = ListPostsResponse{} }
func (m *ListPostsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPostsResponse) ProtoMessage()    {}
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{2}
}

func (m *ListPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsResponse.Unmarshal(m, b)
}
func (m *ListPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsResponse.Marshal(b, m, deterministic)
}
func (m *ListPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsResponse.Merge(m, src)
}
func (m *ListPostsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPostsResponse.Size(m)
}
func (m *ListPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsResponse proto.InternalMessageInfo

func (m *ListPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *ListPostsResponse) GetNextPage() int32 {
	if m != nil {
		return m.NextPage
	}
	return 0
}

type GetPostRequest struct {
	PostId               int64    `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	BlogId               int64    `protobuf:"varint,2,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostRequest) Reset()         { *m = GetPostRequest{} }
func (m *GetPostRequest) String() string { return proto.CompactTextString(m) }
func (*GetPostRequest) ProtoMessage()    {}
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{3}
}

func (m *GetPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPostRequest.Unmarshal(m, b)
}
func (m *GetPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPostRequest.Marshal(b, m, deterministic)
}
func (m *GetPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostRequest.Merge(m, src)
}
func (m *GetPostRequest) XXX_Size() int {
	return xxx_messageInfo_GetPostRequest.Size(m)
}
func (m *GetPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostRequest proto.InternalMessageInfo

func (m *GetPostRequest) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *GetPostRequest) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

type CreatePostRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{4}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePostRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type UpdatePostRequest struct {
	Post                 *Post                 `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdatePostRequest) Reset()         { *m = UpdatePostRequest{} }
func (m *UpdatePostRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePostRequest) ProtoMessage()    {}
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{5}
}

func (m *UpdatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostRequest.Unmarshal(m, b)
}
func (m *UpdatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostRequest.Merge(m, src)
}
func (m *UpdatePostRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePostRequest.Size(m)
}
func (m *UpdatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostRequest proto.InternalMessageInfo

func (m *UpdatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *UpdatePostRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeletePostRequest struct {
	PostId               int64    `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRequest) Reset()         { *m = DeletePostRequest{} }
func (m *DeletePostRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePostRequest) ProtoMessage()    {}
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{6}
}

func (m *DeletePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRequest.Unmarshal(m, b)
}
func (m *DeletePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRequest.Marshal(b, m, deterministic)
}
func (m *DeletePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRequest.Merge(m, src)
}
func (m *DeletePostRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePostRequest.Size(m)
}
func (m *DeletePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRequest proto.InternalMessageInfo

func (m *DeletePostRequest) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*Post)(nil), "proto.Post")
	proto.RegisterType((*ListPostsRequest)(nil), "proto.ListPostsRequest")
	proto.RegisterType((*ListPostsResponse)(nil), "proto.ListPostsResponse")
	proto.RegisterType((*GetPostRequest)(nil), "proto.GetPostRequest")
	proto.RegisterType((*CreatePostRequest)(nil), "proto.CreatePostRequest")
	proto.RegisterType((*UpdatePostRequest)(nil), "proto.UpdatePostRequest")
	proto.RegisterType((*DeletePostRequest)(nil), "proto.DeletePostRequest")
}

func init() { proto.RegisterFile("proto/post.proto", fileDescriptor_5e57bba97c118b83) }

var fileDescriptor_5e57bba97c118b83 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x25, 0x99, 0x7c, 0xec, 0x54, 0x60, 0xd9, 0x34, 0xc2, 0xc6, 0xf5, 0x60, 0xcc, 0x29, 0x07,
	0xc9, 0xc0, 0x7a, 0x14, 0x2f, 0xa3, 0x28, 0x0b, 0x0a, 0x4b, 0x8b, 0xe7, 0x90, 0xd0, 0xb5, 0x21,
	0x4c, 0xdc, 0xce, 0xa6, 0x3b, 0x60, 0xce, 0xfe, 0x71, 0xa9, 0x6a, 0x76, 0x8c, 0x83, 0x73, 0xea,
	0x7a, 0x55, 0xfd, 0xaa, 0x5e, 0x3d, 0x0a, 0xae, 0xc6, 0x49, 0x5b, 0xbd, 0x1b, 0xb5, 0xb1, 0x15,
	0x87, 0x22, 0xe4, 0xe7, 0x26, 0xef, 0xb4, 0xee, 0x06, 0xdc, 0x31, 0x6a, 0xe7, 0x87, 0xdd, 0x43,
	0x8f, 0x83, 0xaa, 0x7f, 0x36, 0xe6, 0xe0, 0x3e, 0x16, 0xbf, 0x3d, 0x08, 0xee, 0xb5, 0xb1, 0xe2,
	0x12, 0xfc, 0x5e, 0x65, 0x5e, 0xee, 0x95, 0x1b, 0xe9, 0xf7, 0x4a, 0x5c, 0x43, 0xdc, 0x0e, 0xba,
	0xab, 0x7b, 0x95, 0xf9, 0x9c, 0x8c, 0x08, 0xde, 0x29, 0xf1, 0x02, 0x42, 0xdb, 0xdb, 0x01, 0xb3,
	0x20, 0xf7, 0xca, 0xad, 0x74, 0x40, 0x08, 0x08, 0x5a, 0xad, 0x96, 0x2c, 0xe4, 0x24, 0xc7, 0x94,
	0x33, 0xc3, 0xdc, 0x65, 0x91, 0xcb, 0x51, 0x4c, 0xec, 0xa1, 0x3f, 0xa0, 0xc9, 0xe2, 0xdc, 0x2b,
	0x43, 0xe9, 0x40, 0x61, 0xe0, 0xea, 0x6b, 0x6f, 0x2c, 0x09, 0x31, 0x12, 0x9f, 0x66, 0x34, 0x96,
	0xd8, 0x63, 0xd3, 0x21, 0x4b, 0x0a, 0x25, 0xc7, 0x24, 0xca, 0xe8, 0xc9, 0xd6, 0xed, 0xc2, 0xa2,
	0xb6, 0x32, 0x22, 0xb8, 0x5f, 0xc4, 0x4b, 0xb8, 0xd0, 0x93, 0xc2, 0x89, 0x2a, 0x1b, 0xae, 0xc4,
	0x8c, 0xf7, 0xcb, 0x7a, 0x91, 0x60, 0xbd, 0x48, 0xf1, 0x1d, 0xd2, 0xd5, 0x50, 0x33, 0xea, 0x47,
	0x83, 0xe2, 0x0d, 0x84, 0x64, 0xa3, 0xc9, 0xbc, 0x7c, 0x53, 0x26, 0xb7, 0x89, 0xb3, 0xa9, 0xa2,
	0x4f, 0xd2, 0x55, 0xc4, 0x2b, 0xd8, 0x3e, 0xe2, 0x2f, 0x5b, 0xb3, 0x3a, 0x9f, 0xd5, 0x5d, 0x50,
	0xe2, 0xbe, 0xe9, 0xb0, 0xd8, 0xc3, 0xe5, 0x17, 0xe4, 0x9e, 0xcf, 0x7b, 0x5c, 0x43, 0x4c, 0xbc,
	0xfa, 0xe8, 0x6e, 0x44, 0xf0, 0xee, 0xbc, 0xc3, 0xc5, 0x07, 0x48, 0x3f, 0x4e, 0xd8, 0x58, 0x5c,
	0xb7, 0x39, 0xda, 0xee, 0xfd, 0xcf, 0x76, 0xff, 0xaf, 0xed, 0xc5, 0x13, 0xa4, 0x3f, 0x46, 0x75,
	0x42, 0x7f, 0x0d, 0x01, 0x8d, 0x65, 0xf6, 0xc9, 0x5a, 0x5c, 0x10, 0xef, 0x21, 0x99, 0x99, 0xc5,
	0xd7, 0xc1, 0x0d, 0x93, 0xdb, 0x9b, 0xca, 0x1d, 0x50, 0xf5, 0x7c, 0x40, 0xd5, 0x67, 0x3a, 0xa0,
	0x6f, 0x8d, 0x39, 0x48, 0x70, 0xdf, 0x29, 0x2e, 0xde, 0x42, 0xfa, 0x09, 0x07, 0xfc, 0x77, 0xe4,
	0xb9, 0xc5, 0xdb, 0x88, 0xbb, 0xbd, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x00, 0xbc, 0xfd, 0x07,
	0xb7, 0x02, 0x00, 0x00,
}
