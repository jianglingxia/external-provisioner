// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta1/index.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The mode determines how a field is indexed.
type IndexField_Mode int32

const (
	// The mode is unspecified.
	IndexField_MODE_UNSPECIFIED IndexField_Mode = 0
	// The field's values are indexed so as to support sequencing in
	// ascending order and also query by <, >, <=, >=, and =.
	IndexField_ASCENDING IndexField_Mode = 2
	// The field's values are indexed so as to support sequencing in
	// descending order and also query by <, >, <=, >=, and =.
	IndexField_DESCENDING IndexField_Mode = 3
	// The field's array values are indexed so as to support membership using
	// ARRAY_CONTAINS queries.
	IndexField_ARRAY_CONTAINS IndexField_Mode = 4
)

var IndexField_Mode_name = map[int32]string{
	0: "MODE_UNSPECIFIED",
	2: "ASCENDING",
	3: "DESCENDING",
	4: "ARRAY_CONTAINS",
}

var IndexField_Mode_value = map[string]int32{
	"MODE_UNSPECIFIED": 0,
	"ASCENDING":        2,
	"DESCENDING":       3,
	"ARRAY_CONTAINS":   4,
}

func (x IndexField_Mode) String() string {
	return proto.EnumName(IndexField_Mode_name, int32(x))
}

func (IndexField_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_18dcb37b2ab27460, []int{0, 0}
}

// The state of an index. During index creation, an index will be in the
// `CREATING` state. If the index is created successfully, it will transition
// to the `READY` state. If the index is not able to be created, it will
// transition to the `ERROR` state.
type Index_State int32

const (
	// The state is unspecified.
	Index_STATE_UNSPECIFIED Index_State = 0
	// The index is being created.
	// There is an active long-running operation for the index.
	// The index is updated when writing a document.
	// Some index data may exist.
	Index_CREATING Index_State = 3
	// The index is ready to be used.
	// The index is updated when writing a document.
	// The index is fully populated from all stored documents it applies to.
	Index_READY Index_State = 2
	// The index was being created, but something went wrong.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing a document.
	// Some index data may exist.
	Index_ERROR Index_State = 5
)

var Index_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	3: "CREATING",
	2: "READY",
	5: "ERROR",
}

var Index_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"CREATING":          3,
	"READY":             2,
	"ERROR":             5,
}

func (x Index_State) String() string {
	return proto.EnumName(Index_State_name, int32(x))
}

func (Index_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_18dcb37b2ab27460, []int{1, 0}
}

// A field of an index.
type IndexField struct {
	// The path of the field. Must match the field path specification described
	// by [google.firestore.v1beta1.Document.fields][fields].
	// Special field path `__name__` may be used by itself or at the end of a
	// path. `__type__` may be used only at the end of path.
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// The field's mode.
	Mode                 IndexField_Mode `protobuf:"varint,2,opt,name=mode,proto3,enum=google.firestore.admin.v1beta1.IndexField_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IndexField) Reset()         { *m = IndexField{} }
func (m *IndexField) String() string { return proto.CompactTextString(m) }
func (*IndexField) ProtoMessage()    {}
func (*IndexField) Descriptor() ([]byte, []int) {
	return fileDescriptor_18dcb37b2ab27460, []int{0}
}
func (m *IndexField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexField.Unmarshal(m, b)
}
func (m *IndexField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexField.Marshal(b, m, deterministic)
}
func (m *IndexField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexField.Merge(m, src)
}
func (m *IndexField) XXX_Size() int {
	return xxx_messageInfo_IndexField.Size(m)
}
func (m *IndexField) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexField.DiscardUnknown(m)
}

var xxx_messageInfo_IndexField proto.InternalMessageInfo

func (m *IndexField) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

func (m *IndexField) GetMode() IndexField_Mode {
	if m != nil {
		return m.Mode
	}
	return IndexField_MODE_UNSPECIFIED
}

// An index definition.
type Index struct {
	// The resource name of the index.
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The collection ID to which this index applies. Required.
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// The fields to index.
	Fields []*IndexField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// The state of the index.
	// Output only.
	State                Index_State `protobuf:"varint,6,opt,name=state,proto3,enum=google.firestore.admin.v1beta1.Index_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_18dcb37b2ab27460, []int{1}
}
func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Index) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *Index) GetFields() []*IndexField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Index) GetState() Index_State {
	if m != nil {
		return m.State
	}
	return Index_STATE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*IndexField)(nil), "google.firestore.admin.v1beta1.IndexField")
	proto.RegisterType((*Index)(nil), "google.firestore.admin.v1beta1.Index")
	proto.RegisterEnum("google.firestore.admin.v1beta1.IndexField_Mode", IndexField_Mode_name, IndexField_Mode_value)
	proto.RegisterEnum("google.firestore.admin.v1beta1.Index_State", Index_State_name, Index_State_value)
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta1/index.proto", fileDescriptor_18dcb37b2ab27460)
}

var fileDescriptor_18dcb37b2ab27460 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0x69, 0x52, 0xcc, 0x75, 0xb7, 0xc4, 0x41, 0xa1, 0x88, 0x4a, 0x89, 0x3e, 0x94, 0x15,
	0x26, 0x74, 0x7d, 0xf4, 0x69, 0xf2, 0xd1, 0x92, 0x87, 0x4d, 0xcb, 0xa4, 0x0a, 0xeb, 0x4b, 0x99,
	0x6d, 0x66, 0xb3, 0x81, 0x74, 0xa6, 0x34, 0xa3, 0xf8, 0x1f, 0xfc, 0x17, 0x82, 0x2f, 0xfe, 0x0b,
	0xff, 0x99, 0xcc, 0x24, 0x74, 0x41, 0xd0, 0xed, 0xdb, 0x3d, 0x93, 0x73, 0xce, 0x3d, 0xf7, 0xe6,
	0xc2, 0x45, 0x25, 0x65, 0xd5, 0xf0, 0xf0, 0xb6, 0x3e, 0xf0, 0x56, 0xc9, 0x03, 0x0f, 0x59, 0xb9,
	0xab, 0x45, 0xf8, 0x75, 0x76, 0xc3, 0x15, 0x9b, 0x85, 0xb5, 0x28, 0xf9, 0x37, 0xbc, 0x3f, 0x48,
	0x25, 0xd1, 0xeb, 0x8e, 0x8b, 0x8f, 0x5c, 0x6c, 0xb8, 0xb8, 0xe7, 0xbe, 0x78, 0xd9, 0x7b, 0xb1,
	0x7d, 0x1d, 0x32, 0x21, 0xa4, 0x62, 0xaa, 0x96, 0xa2, 0xed, 0xd4, 0xc1, 0x6f, 0x0b, 0x20, 0xd3,
	0x6e, 0xf3, 0x9a, 0x37, 0x25, 0x7a, 0x05, 0x70, 0xab, 0x8b, 0xcd, 0x9e, 0xa9, 0xbb, 0xb1, 0x35,
	0xb1, 0xa6, 0x1e, 0xf5, 0xcc, 0xcb, 0x8a, 0xa9, 0x3b, 0x14, 0x83, 0xb3, 0x93, 0x25, 0x1f, 0xdb,
	0x13, 0x6b, 0x3a, 0xba, 0x0c, 0xf1, 0xff, 0x5b, 0xe3, 0x7b, 0x63, 0x7c, 0x25, 0x4b, 0x4e, 0x8d,
	0x38, 0x58, 0x82, 0xa3, 0x11, 0x7a, 0x06, 0xfe, 0xd5, 0x32, 0x49, 0x37, 0x1f, 0xf3, 0x62, 0x95,
	0xc6, 0xd9, 0x3c, 0x4b, 0x13, 0xff, 0x11, 0x3a, 0x07, 0x8f, 0x14, 0x71, 0x9a, 0x27, 0x59, 0xbe,
	0xf0, 0x6d, 0x34, 0x02, 0x48, 0xd2, 0x23, 0x1e, 0x20, 0x04, 0x23, 0x42, 0x29, 0xb9, 0xde, 0xc4,
	0xcb, 0x7c, 0x4d, 0xb2, 0xbc, 0xf0, 0x9d, 0xe0, 0xbb, 0x0d, 0xae, 0x69, 0x85, 0x10, 0x38, 0x82,
	0xed, 0x78, 0x1f, 0xdc, 0xd4, 0xe8, 0x0d, 0x9c, 0x6f, 0x65, 0xd3, 0xf0, 0xad, 0x1e, 0x7b, 0x53,
	0x97, 0x26, 0xbc, 0x47, 0xcf, 0xee, 0x1f, 0xb3, 0x12, 0x45, 0x30, 0x34, 0x53, 0xb6, 0xe3, 0xc1,
	0x64, 0x30, 0x7d, 0x72, 0x79, 0x71, 0xfa, 0x68, 0xb4, 0x57, 0x22, 0x02, 0x6e, 0xab, 0x98, 0xe2,
	0xe3, 0xa1, 0xd9, 0xce, 0xbb, 0x93, 0x2c, 0x70, 0xa1, 0x25, 0xb4, 0x53, 0x06, 0x11, 0xb8, 0x06,
	0xa3, 0xe7, 0xf0, 0xb4, 0x58, 0x93, 0xf5, 0xdf, 0xcb, 0x39, 0x83, 0xc7, 0x31, 0x4d, 0xc9, 0xba,
	0xdb, 0x85, 0x07, 0x2e, 0x4d, 0x49, 0x72, 0xed, 0xdb, 0xba, 0x4c, 0x29, 0x5d, 0x52, 0xdf, 0x8d,
	0x7e, 0x5a, 0x10, 0x6c, 0xe5, 0xee, 0x81, 0xee, 0x51, 0xf7, 0xd7, 0x57, 0xfa, 0x08, 0x56, 0xd6,
	0xe7, 0xb8, 0x67, 0x57, 0xb2, 0x61, 0xa2, 0xc2, 0xf2, 0x50, 0x85, 0x15, 0x17, 0xe6, 0x44, 0xc2,
	0xee, 0x13, 0xdb, 0xd7, 0xed, 0xbf, 0xee, 0xf1, 0x83, 0x41, 0x3f, 0x6c, 0x67, 0x11, 0xcf, 0x8b,
	0x5f, 0xf6, 0xdb, 0x45, 0x67, 0x16, 0x37, 0xf2, 0x4b, 0x89, 0xe7, 0xc7, 0x00, 0xc4, 0x04, 0xf8,
	0x34, 0x8b, 0xb4, 0xe6, 0x66, 0x68, 0xdc, 0xdf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xbf,
	0x61, 0xbe, 0xec, 0x02, 0x00, 0x00,
}
