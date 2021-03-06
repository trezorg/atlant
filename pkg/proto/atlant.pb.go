// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: pkg/proto/atlant.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type StateEnum int32

const (
	StateEnum_IN_PROGRESS StateEnum = 0
	StateEnum_SUCCESS     StateEnum = 1
	StateEnum_ERROR       StateEnum = 2
)

// Enum value maps for StateEnum.
var (
	StateEnum_name = map[int32]string{
		0: "IN_PROGRESS",
		1: "SUCCESS",
		2: "ERROR",
	}
	StateEnum_value = map[string]int32{
		"IN_PROGRESS": 0,
		"SUCCESS":     1,
		"ERROR":       2,
	}
)

func (x StateEnum) Enum() *StateEnum {
	p := new(StateEnum)
	*p = x
	return p
}

func (x StateEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_atlant_proto_enumTypes[0].Descriptor()
}

func (StateEnum) Type() protoreflect.EnumType {
	return &file_pkg_proto_atlant_proto_enumTypes[0]
}

func (x StateEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateEnum.Descriptor instead.
func (StateEnum) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{0}
}

type SortingField int32

const (
	SortingField_NAME          SortingField = 0
	SortingField_PRICE         SortingField = 1
	SortingField_UPDATED_AT    SortingField = 2
	SortingField_PRICE_CHANGES SortingField = 3
)

// Enum value maps for SortingField.
var (
	SortingField_name = map[int32]string{
		0: "NAME",
		1: "PRICE",
		2: "UPDATED_AT",
		3: "PRICE_CHANGES",
	}
	SortingField_value = map[string]int32{
		"NAME":          0,
		"PRICE":         1,
		"UPDATED_AT":    2,
		"PRICE_CHANGES": 3,
	}
)

func (x SortingField) Enum() *SortingField {
	p := new(SortingField)
	*p = x
	return p
}

func (x SortingField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingField) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_atlant_proto_enumTypes[1].Descriptor()
}

func (SortingField) Type() protoreflect.EnumType {
	return &file_pkg_proto_atlant_proto_enumTypes[1]
}

func (x SortingField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingField.Descriptor instead.
func (SortingField) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{1}
}

type SortingOrder int32

const (
	SortingOrder_ASC  SortingOrder = 0
	SortingOrder_DESC SortingOrder = 1
)

// Enum value maps for SortingOrder.
var (
	SortingOrder_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	SortingOrder_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x SortingOrder) Enum() *SortingOrder {
	p := new(SortingOrder)
	*p = x
	return p
}

func (x SortingOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_atlant_proto_enumTypes[2].Descriptor()
}

func (SortingOrder) Type() protoreflect.EnumType {
	return &file_pkg_proto_atlant_proto_enumTypes[2]
}

func (x SortingOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingOrder.Descriptor instead.
func (SortingOrder) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{2}
}

type FetchState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State         StateEnum `protobuf:"varint,1,opt,name=state,proto3,enum=proto.StateEnum" json:"state,omitempty"`
	LoadedRecords int32     `protobuf:"varint,2,opt,name=loaded_records,json=loadedRecords,proto3" json:"loaded_records,omitempty"`
}

func (x *FetchState) Reset() {
	*x = FetchState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_atlant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchState) ProtoMessage() {}

func (x *FetchState) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_atlant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchState.ProtoReflect.Descriptor instead.
func (*FetchState) Descriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{0}
}

func (x *FetchState) GetState() StateEnum {
	if x != nil {
		return x.State
	}
	return StateEnum_IN_PROGRESS
}

func (x *FetchState) GetLoadedRecords() int32 {
	if x != nil {
		return x.LoadedRecords
	}
	return 0
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url        string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	SkipHeader bool   `protobuf:"varint,2,opt,name=skip_header,json=skipHeader,proto3" json:"skip_header,omitempty"`
	Separator  int32  `protobuf:"varint,3,opt,name=separator,proto3" json:"separator,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_atlant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_atlant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{1}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FetchRequest) GetSkipHeader() bool {
	if x != nil {
		return x.SkipHeader
	}
	return false
}

func (x *FetchRequest) GetSeparator() int32 {
	if x != nil {
		return x.Separator
	}
	return 0
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Field string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_atlant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_atlant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{2}
}

func (x *Cursor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cursor) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit   uint32   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Sorting *Sorting `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Cursor  *Cursor  `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_atlant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_atlant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{3}
}

func (x *Page) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Page) GetSorting() *Sorting {
	if x != nil {
		return x.Sorting
	}
	return nil
}

func (x *Page) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

type Sorting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field SortingField `protobuf:"varint,1,opt,name=field,proto3,enum=proto.SortingField" json:"field,omitempty"`
	Order SortingOrder `protobuf:"varint,2,opt,name=order,proto3,enum=proto.SortingOrder" json:"order,omitempty"`
}

func (x *Sorting) Reset() {
	*x = Sorting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_atlant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sorting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sorting) ProtoMessage() {}

func (x *Sorting) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_atlant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sorting.ProtoReflect.Descriptor instead.
func (*Sorting) Descriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{4}
}

func (x *Sorting) GetField() SortingField {
	if x != nil {
		return x.Field
	}
	return SortingField_NAME
}

func (x *Sorting) GetOrder() SortingOrder {
	if x != nil {
		return x.Order
	}
	return SortingOrder_ASC
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price        int32                `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	UpdatedAt    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PriceChanges int32                `protobuf:"varint,4,opt,name=price_changes,json=priceChanges,proto3" json:"price_changes,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_atlant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_atlant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Product) GetPriceChanges() int32 {
	if x != nil {
		return x.PriceChanges
	}
	return 0
}

type Products struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Cursor   *Cursor    `protobuf:"bytes,2,opt,name=Cursor,proto3" json:"Cursor,omitempty"`
}

func (x *Products) Reset() {
	*x = Products{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_atlant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_atlant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_pkg_proto_atlant_proto_rawDescGZIP(), []int{6}
}

func (x *Products) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Products) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

var File_pkg_proto_atlant_proto protoreflect.FileDescriptor

var file_pkg_proto_atlant_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x6c, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5b, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x5f, 0x0a,
	0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x32,
	0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x22, 0x6d, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x28, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x22, 0x5f, 0x0a, 0x07, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x12, 0x25, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52,
	0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2a, 0x34, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x2a, 0x46, 0x0a,
	0x0c, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x49, 0x43, 0x45,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x53, 0x10, 0x03, 0x2a, 0x21, 0x0a, 0x0c, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x32, 0x6c, 0x0a, 0x0d, 0x41, 0x74, 0x6c, 0x61,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x26,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x22, 0x00, 0x42, 0x4c, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74,
	0x42, 0x0b, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x65, 0x7a,
	0x6f, 0x72, 0x67, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_atlant_proto_rawDescOnce sync.Once
	file_pkg_proto_atlant_proto_rawDescData = file_pkg_proto_atlant_proto_rawDesc
)

func file_pkg_proto_atlant_proto_rawDescGZIP() []byte {
	file_pkg_proto_atlant_proto_rawDescOnce.Do(func() {
		file_pkg_proto_atlant_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_atlant_proto_rawDescData)
	})
	return file_pkg_proto_atlant_proto_rawDescData
}

var file_pkg_proto_atlant_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_proto_atlant_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_atlant_proto_goTypes = []interface{}{
	(StateEnum)(0),              // 0: proto.StateEnum
	(SortingField)(0),           // 1: proto.SortingField
	(SortingOrder)(0),           // 2: proto.SortingOrder
	(*FetchState)(nil),          // 3: proto.FetchState
	(*FetchRequest)(nil),        // 4: proto.FetchRequest
	(*Cursor)(nil),              // 5: proto.Cursor
	(*Page)(nil),                // 6: proto.Page
	(*Sorting)(nil),             // 7: proto.Sorting
	(*Product)(nil),             // 8: proto.Product
	(*Products)(nil),            // 9: proto.Products
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_pkg_proto_atlant_proto_depIdxs = []int32{
	0,  // 0: proto.FetchState.state:type_name -> proto.StateEnum
	7,  // 1: proto.Page.sorting:type_name -> proto.Sorting
	5,  // 2: proto.Page.cursor:type_name -> proto.Cursor
	1,  // 3: proto.Sorting.field:type_name -> proto.SortingField
	2,  // 4: proto.Sorting.order:type_name -> proto.SortingOrder
	10, // 5: proto.Product.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 6: proto.Products.products:type_name -> proto.Product
	5,  // 7: proto.Products.Cursor:type_name -> proto.Cursor
	4,  // 8: proto.AtlantService.Fetch:input_type -> proto.FetchRequest
	6,  // 9: proto.AtlantService.List:input_type -> proto.Page
	3,  // 10: proto.AtlantService.Fetch:output_type -> proto.FetchState
	9,  // 11: proto.AtlantService.List:output_type -> proto.Products
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_proto_atlant_proto_init() }
func file_pkg_proto_atlant_proto_init() {
	if File_pkg_proto_atlant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_atlant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchState); i {
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
		file_pkg_proto_atlant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_pkg_proto_atlant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
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
		file_pkg_proto_atlant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_pkg_proto_atlant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sorting); i {
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
		file_pkg_proto_atlant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_pkg_proto_atlant_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Products); i {
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
			RawDescriptor: file_pkg_proto_atlant_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_atlant_proto_goTypes,
		DependencyIndexes: file_pkg_proto_atlant_proto_depIdxs,
		EnumInfos:         file_pkg_proto_atlant_proto_enumTypes,
		MessageInfos:      file_pkg_proto_atlant_proto_msgTypes,
	}.Build()
	File_pkg_proto_atlant_proto = out.File
	file_pkg_proto_atlant_proto_rawDesc = nil
	file_pkg_proto_atlant_proto_goTypes = nil
	file_pkg_proto_atlant_proto_depIdxs = nil
}
