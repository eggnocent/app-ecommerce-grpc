// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: order/order.proto

package order

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	common "github/eggnocent/app-grpc-eccomerce/pb/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProductRequestProductItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity      int64                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequestProductItem) Reset() {
	*x = CreateProductRequestProductItem{}
	mi := &file_order_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequestProductItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequestProductItem) ProtoMessage() {}

func (x *CreateProductRequestProductItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequestProductItem.ProtoReflect.Descriptor instead.
func (*CreateProductRequestProductItem) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequestProductItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProductRequestProductItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState             `protogen:"open.v1"`
	FullName      string                             `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Address       string                             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumber   string                             `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Notes         string                             `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
	Products      []*CreateProductRequestProductItem `protobuf:"bytes,5,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateOrderRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateOrderRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateOrderRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *CreateOrderRequest) GetProducts() []*CreateProductRequestProductItem {
	if x != nil {
		return x.Products
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Base          *common.BaseResponse   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_order_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetBase() *common.BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CreateOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListOrderAdminRequest struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Pagination    *common.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderAdminRequest) Reset() {
	*x = ListOrderAdminRequest{}
	mi := &file_order_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderAdminRequest) ProtoMessage() {}

func (x *ListOrderAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderAdminRequest.ProtoReflect.Descriptor instead.
func (*ListOrderAdminRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{3}
}

func (x *ListOrderAdminRequest) GetPagination() *common.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListOrderAdminResponseItemProduct struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity      int64                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderAdminResponseItemProduct) Reset() {
	*x = ListOrderAdminResponseItemProduct{}
	mi := &file_order_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderAdminResponseItemProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderAdminResponseItemProduct) ProtoMessage() {}

func (x *ListOrderAdminResponseItemProduct) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderAdminResponseItemProduct.ProtoReflect.Descriptor instead.
func (*ListOrderAdminResponseItemProduct) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{4}
}

func (x *ListOrderAdminResponseItemProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListOrderAdminResponseItemProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListOrderAdminResponseItemProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ListOrderAdminResponseItemProduct) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ListOrderAdminResponseItem struct {
	state         protoimpl.MessageState               `protogen:"open.v1"`
	Id            string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number        string                               `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Customer      string                               `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	StatusCode    string                               `protobuf:"bytes,4,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Total         float64                              `protobuf:"fixed64,5,opt,name=total,proto3" json:"total,omitempty"`
	CreatedAt     *timestamppb.Timestamp               `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Product       []*ListOrderAdminResponseItemProduct `protobuf:"bytes,7,rep,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderAdminResponseItem) Reset() {
	*x = ListOrderAdminResponseItem{}
	mi := &file_order_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderAdminResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderAdminResponseItem) ProtoMessage() {}

func (x *ListOrderAdminResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderAdminResponseItem.ProtoReflect.Descriptor instead.
func (*ListOrderAdminResponseItem) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{5}
}

func (x *ListOrderAdminResponseItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListOrderAdminResponseItem) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *ListOrderAdminResponseItem) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *ListOrderAdminResponseItem) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

func (x *ListOrderAdminResponseItem) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListOrderAdminResponseItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListOrderAdminResponseItem) GetProduct() []*ListOrderAdminResponseItemProduct {
	if x != nil {
		return x.Product
	}
	return nil
}

type ListOrderAdminResponse struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Base          *common.BaseResponse          `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Pagination    *common.PaginationResponse    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Items         []*ListOrderAdminResponseItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderAdminResponse) Reset() {
	*x = ListOrderAdminResponse{}
	mi := &file_order_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderAdminResponse) ProtoMessage() {}

func (x *ListOrderAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderAdminResponse.ProtoReflect.Descriptor instead.
func (*ListOrderAdminResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrderAdminResponse) GetBase() *common.BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ListOrderAdminResponse) GetPagination() *common.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListOrderAdminResponse) GetItems() []*ListOrderAdminResponseItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListOrderRequest struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Pagination    *common.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderRequest) Reset() {
	*x = ListOrderRequest{}
	mi := &file_order_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderRequest) ProtoMessage() {}

func (x *ListOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderRequest.ProtoReflect.Descriptor instead.
func (*ListOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrderRequest) GetPagination() *common.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListOrderResponseItemProduct struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity      int64                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderResponseItemProduct) Reset() {
	*x = ListOrderResponseItemProduct{}
	mi := &file_order_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderResponseItemProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponseItemProduct) ProtoMessage() {}

func (x *ListOrderResponseItemProduct) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponseItemProduct.ProtoReflect.Descriptor instead.
func (*ListOrderResponseItemProduct) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{8}
}

func (x *ListOrderResponseItemProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListOrderResponseItemProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListOrderResponseItemProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ListOrderResponseItemProduct) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ListOrderResponseItem struct {
	state            protoimpl.MessageState          `protogen:"open.v1"`
	Id               string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number           string                          `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Customer         string                          `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	StatusCode       string                          `protobuf:"bytes,4,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Total            float64                         `protobuf:"fixed64,5,opt,name=total,proto3" json:"total,omitempty"`
	CreatedAt        *timestamppb.Timestamp          `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Product          []*ListOrderResponseItemProduct `protobuf:"bytes,7,rep,name=product,proto3" json:"product,omitempty"`
	XenditInvoiceUrl string                          `protobuf:"bytes,8,opt,name=xendit_invoice_url,json=xenditInvoiceUrl,proto3" json:"xendit_invoice_url,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ListOrderResponseItem) Reset() {
	*x = ListOrderResponseItem{}
	mi := &file_order_order_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponseItem) ProtoMessage() {}

func (x *ListOrderResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponseItem.ProtoReflect.Descriptor instead.
func (*ListOrderResponseItem) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{9}
}

func (x *ListOrderResponseItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListOrderResponseItem) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *ListOrderResponseItem) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *ListOrderResponseItem) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

func (x *ListOrderResponseItem) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListOrderResponseItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListOrderResponseItem) GetProduct() []*ListOrderResponseItemProduct {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ListOrderResponseItem) GetXenditInvoiceUrl() string {
	if x != nil {
		return x.XenditInvoiceUrl
	}
	return ""
}

type ListOrderResponse struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Base          *common.BaseResponse       `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Pagination    *common.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Items         []*ListOrderResponseItem   `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderResponse) Reset() {
	*x = ListOrderResponse{}
	mi := &file_order_order_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponse) ProtoMessage() {}

func (x *ListOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponse.ProtoReflect.Descriptor instead.
func (*ListOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{10}
}

func (x *ListOrderResponse) GetBase() *common.BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ListOrderResponse) GetPagination() *common.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListOrderResponse) GetItems() []*ListOrderResponseItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_order_order_proto protoreflect.FileDescriptor

const file_order_order_proto_rawDesc = "" +
	"\n" +
	"\x11order/order.proto\x12\x05order\x1a\x1bbuf/validate/validate.proto\x1a\x1acommon/base-response.proto\x1a\x17common/pagination.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"M\n" +
	"\x1fCreateProductRequestProductItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x03R\bquantity\"\xf6\x01\n" +
	"\x12CreateOrderRequest\x12'\n" +
	"\tfull_name\x18\x01 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x01\x18\xff\x01R\bfullName\x12$\n" +
	"\aaddress\x18\x02 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x01\x18\xff\x01R\aaddress\x12-\n" +
	"\fphone_number\x18\x03 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x01\x18\xff\x01R\vphoneNumber\x12\x1e\n" +
	"\x05notes\x18\x04 \x01(\tB\b\xbaH\x05r\x03\x18\xff\x01R\x05notes\x12B\n" +
	"\bproducts\x18\x05 \x03(\v2&.order.CreateProductRequestProductItemR\bproducts\"O\n" +
	"\x13CreateOrderResponse\x12(\n" +
	"\x04base\x18\x01 \x01(\v2\x14.common.BaseResponseR\x04base\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\"R\n" +
	"\x15ListOrderAdminRequest\x129\n" +
	"\n" +
	"pagination\x18\x01 \x01(\v2\x19.common.PaginationRequestR\n" +
	"pagination\"y\n" +
	"!ListOrderAdminResponseItemProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\x12\x1a\n" +
	"\bquantity\x18\x04 \x01(\x03R\bquantity\"\x96\x02\n" +
	"\x1aListOrderAdminResponseItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06number\x18\x02 \x01(\tR\x06number\x12\x1a\n" +
	"\bcustomer\x18\x03 \x01(\tR\bcustomer\x12\x1f\n" +
	"\vstatus_code\x18\x04 \x01(\tR\n" +
	"statusCode\x12\x14\n" +
	"\x05total\x18\x05 \x01(\x01R\x05total\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12B\n" +
	"\aproduct\x18\a \x03(\v2(.order.ListOrderAdminResponseItemProductR\aproduct\"\xb7\x01\n" +
	"\x16ListOrderAdminResponse\x12(\n" +
	"\x04base\x18\x01 \x01(\v2\x14.common.BaseResponseR\x04base\x12:\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x1a.common.PaginationResponseR\n" +
	"pagination\x127\n" +
	"\x05items\x18\x03 \x03(\v2!.order.ListOrderAdminResponseItemR\x05items\"M\n" +
	"\x10ListOrderRequest\x129\n" +
	"\n" +
	"pagination\x18\x01 \x01(\v2\x19.common.PaginationRequestR\n" +
	"pagination\"t\n" +
	"\x1cListOrderResponseItemProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\x12\x1a\n" +
	"\bquantity\x18\x04 \x01(\x03R\bquantity\"\xba\x02\n" +
	"\x15ListOrderResponseItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06number\x18\x02 \x01(\tR\x06number\x12\x1a\n" +
	"\bcustomer\x18\x03 \x01(\tR\bcustomer\x12\x1f\n" +
	"\vstatus_code\x18\x04 \x01(\tR\n" +
	"statusCode\x12\x14\n" +
	"\x05total\x18\x05 \x01(\x01R\x05total\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12=\n" +
	"\aproduct\x18\a \x03(\v2#.order.ListOrderResponseItemProductR\aproduct\x12,\n" +
	"\x12xendit_invoice_url\x18\b \x01(\tR\x10xenditInvoiceUrl\"\xad\x01\n" +
	"\x11ListOrderResponse\x12(\n" +
	"\x04base\x18\x01 \x01(\v2\x14.common.BaseResponseR\x04base\x12:\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x1a.common.PaginationResponseR\n" +
	"pagination\x122\n" +
	"\x05items\x18\x03 \x03(\v2\x1c.order.ListOrderResponseItemR\x05items2\xe3\x01\n" +
	"\fOrderService\x12D\n" +
	"\vCreateOrder\x12\x19.order.CreateOrderRequest\x1a\x1a.order.CreateOrderResponse\x12M\n" +
	"\x0eListOrderAdmin\x12\x1c.order.ListOrderAdminRequest\x1a\x1d.order.ListOrderAdminResponse\x12>\n" +
	"\tListOrder\x12\x17.order.ListOrderRequest\x1a\x18.order.ListOrderResponseB.Z,github/eggnocent/app-grpc-eccomerce/pb/orderb\x06proto3"

var (
	file_order_order_proto_rawDescOnce sync.Once
	file_order_order_proto_rawDescData []byte
)

func file_order_order_proto_rawDescGZIP() []byte {
	file_order_order_proto_rawDescOnce.Do(func() {
		file_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_order_proto_rawDesc), len(file_order_order_proto_rawDesc)))
	})
	return file_order_order_proto_rawDescData
}

var file_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_order_order_proto_goTypes = []any{
	(*CreateProductRequestProductItem)(nil),   // 0: order.CreateProductRequestProductItem
	(*CreateOrderRequest)(nil),                // 1: order.CreateOrderRequest
	(*CreateOrderResponse)(nil),               // 2: order.CreateOrderResponse
	(*ListOrderAdminRequest)(nil),             // 3: order.ListOrderAdminRequest
	(*ListOrderAdminResponseItemProduct)(nil), // 4: order.ListOrderAdminResponseItemProduct
	(*ListOrderAdminResponseItem)(nil),        // 5: order.ListOrderAdminResponseItem
	(*ListOrderAdminResponse)(nil),            // 6: order.ListOrderAdminResponse
	(*ListOrderRequest)(nil),                  // 7: order.ListOrderRequest
	(*ListOrderResponseItemProduct)(nil),      // 8: order.ListOrderResponseItemProduct
	(*ListOrderResponseItem)(nil),             // 9: order.ListOrderResponseItem
	(*ListOrderResponse)(nil),                 // 10: order.ListOrderResponse
	(*common.BaseResponse)(nil),               // 11: common.BaseResponse
	(*common.PaginationRequest)(nil),          // 12: common.PaginationRequest
	(*timestamppb.Timestamp)(nil),             // 13: google.protobuf.Timestamp
	(*common.PaginationResponse)(nil),         // 14: common.PaginationResponse
}
var file_order_order_proto_depIdxs = []int32{
	0,  // 0: order.CreateOrderRequest.products:type_name -> order.CreateProductRequestProductItem
	11, // 1: order.CreateOrderResponse.base:type_name -> common.BaseResponse
	12, // 2: order.ListOrderAdminRequest.pagination:type_name -> common.PaginationRequest
	13, // 3: order.ListOrderAdminResponseItem.created_at:type_name -> google.protobuf.Timestamp
	4,  // 4: order.ListOrderAdminResponseItem.product:type_name -> order.ListOrderAdminResponseItemProduct
	11, // 5: order.ListOrderAdminResponse.base:type_name -> common.BaseResponse
	14, // 6: order.ListOrderAdminResponse.pagination:type_name -> common.PaginationResponse
	5,  // 7: order.ListOrderAdminResponse.items:type_name -> order.ListOrderAdminResponseItem
	12, // 8: order.ListOrderRequest.pagination:type_name -> common.PaginationRequest
	13, // 9: order.ListOrderResponseItem.created_at:type_name -> google.protobuf.Timestamp
	8,  // 10: order.ListOrderResponseItem.product:type_name -> order.ListOrderResponseItemProduct
	11, // 11: order.ListOrderResponse.base:type_name -> common.BaseResponse
	14, // 12: order.ListOrderResponse.pagination:type_name -> common.PaginationResponse
	9,  // 13: order.ListOrderResponse.items:type_name -> order.ListOrderResponseItem
	1,  // 14: order.OrderService.CreateOrder:input_type -> order.CreateOrderRequest
	3,  // 15: order.OrderService.ListOrderAdmin:input_type -> order.ListOrderAdminRequest
	7,  // 16: order.OrderService.ListOrder:input_type -> order.ListOrderRequest
	2,  // 17: order.OrderService.CreateOrder:output_type -> order.CreateOrderResponse
	6,  // 18: order.OrderService.ListOrderAdmin:output_type -> order.ListOrderAdminResponse
	10, // 19: order.OrderService.ListOrder:output_type -> order.ListOrderResponse
	17, // [17:20] is the sub-list for method output_type
	14, // [14:17] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_order_order_proto_init() }
func file_order_order_proto_init() {
	if File_order_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_order_proto_rawDesc), len(file_order_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_order_proto_goTypes,
		DependencyIndexes: file_order_order_proto_depIdxs,
		MessageInfos:      file_order_order_proto_msgTypes,
	}.Build()
	File_order_order_proto = out.File
	file_order_order_proto_goTypes = nil
	file_order_order_proto_depIdxs = nil
}
