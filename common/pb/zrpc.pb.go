// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: protos/zrpc.proto

package __

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

type EventWithLocationAndType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Date        string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Location    string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Banner      string `protobuf:"bytes,5,opt,name=banner,proto3" json:"banner,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EventWithLocationAndType) Reset() {
	*x = EventWithLocationAndType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventWithLocationAndType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventWithLocationAndType) ProtoMessage() {}

func (x *EventWithLocationAndType) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventWithLocationAndType.ProtoReflect.Descriptor instead.
func (*EventWithLocationAndType) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{0}
}

func (x *EventWithLocationAndType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EventWithLocationAndType) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *EventWithLocationAndType) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *EventWithLocationAndType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EventWithLocationAndType) GetBanner() string {
	if x != nil {
		return x.Banner
	}
	return ""
}

func (x *EventWithLocationAndType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ListEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	City      string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Uf        string `protobuf:"bytes,3,opt,name=uf,proto3" json:"uf,omitempty"`
	StartDate string `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate   string `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Page      int32  `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int32  `protobuf:"varint,7,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListEventsRequest) Reset() {
	*x = ListEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequest) ProtoMessage() {}

func (x *ListEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequest.ProtoReflect.Descriptor instead.
func (*ListEventsRequest) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{1}
}

func (x *ListEventsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListEventsRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ListEventsRequest) GetUf() string {
	if x != nil {
		return x.Uf
	}
	return ""
}

func (x *ListEventsRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *ListEventsRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *ListEventsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListEventsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events      []*EventWithLocationAndType `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	TotalPages  int32                       `protobuf:"varint,2,opt,name=totalPages,proto3" json:"totalPages,omitempty"`
	CurrentPage int32                       `protobuf:"varint,3,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
}

func (x *ListEventsResponse) Reset() {
	*x = ListEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsResponse) ProtoMessage() {}

func (x *ListEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsResponse.ProtoReflect.Descriptor instead.
func (*ListEventsResponse) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{2}
}

func (x *ListEventsResponse) GetEvents() []*EventWithLocationAndType {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *ListEventsResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *ListEventsResponse) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Image       string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	EventUrl    string `protobuf:"bytes,4,opt,name=eventUrl,proto3" json:"eventUrl,omitempty"`
	Remote      bool   `protobuf:"varint,5,opt,name=remote,proto3" json:"remote,omitempty"`
	Date        string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Event) GetEventUrl() string {
	if x != nil {
		return x.EventUrl
	}
	return ""
}

func (x *Event) GetRemote() bool {
	if x != nil {
		return x.Remote
	}
	return false
}

func (x *Event) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Uf    string `protobuf:"bytes,2,opt,name=uf,proto3" json:"uf,omitempty"`
	City  string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{4}
}

func (x *CreateEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *CreateEventRequest) GetUf() string {
	if x != nil {
		return x.Uf
	}
	return ""
}

func (x *CreateEventRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=eventId,proto3" json:"eventId,omitempty"`
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{5}
}

func (x *CreateEventResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type Coupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Discount int32  `protobuf:"varint,2,opt,name=discount,proto3" json:"discount,omitempty"`
	Date     string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Coupon) Reset() {
	*x = Coupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coupon) ProtoMessage() {}

func (x *Coupon) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coupon.ProtoReflect.Descriptor instead.
func (*Coupon) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{6}
}

func (x *Coupon) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Coupon) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Coupon) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CreateCouponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string  `protobuf:"bytes,1,opt,name=eventId,proto3" json:"eventId,omitempty"`
	Coupon  *Coupon `protobuf:"bytes,2,opt,name=coupon,proto3" json:"coupon,omitempty"`
}

func (x *CreateCouponRequest) Reset() {
	*x = CreateCouponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCouponRequest) ProtoMessage() {}

func (x *CreateCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCouponRequest.ProtoReflect.Descriptor instead.
func (*CreateCouponRequest) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{7}
}

func (x *CreateCouponRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *CreateCouponRequest) GetCoupon() *Coupon {
	if x != nil {
		return x.Coupon
	}
	return nil
}

type CreateCouponResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponId string `protobuf:"bytes,1,opt,name=couponId,proto3" json:"couponId,omitempty"`
}

func (x *CreateCouponResponse) Reset() {
	*x = CreateCouponResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_zrpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCouponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCouponResponse) ProtoMessage() {}

func (x *CreateCouponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_zrpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCouponResponse.ProtoReflect.Descriptor instead.
func (*CreateCouponResponse) Descriptor() ([]byte, []int) {
	return file_protos_zrpc_proto_rawDescGZIP(), []int{8}
}

func (x *CreateCouponResponse) GetCouponId() string {
	if x != nil {
		return x.CouponId
	}
	return ""
}

var File_protos_zrpc_proto protoreflect.FileDescriptor

var file_protos_zrpc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x7a, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x7a, 0x72, 0x70, 0x63, 0x22, 0xae, 0x01, 0x0a, 0x18, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb5, 0x01, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x5b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x7a, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x75, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x66, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x4c, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x55, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x7a, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x06,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xd9, 0x01, 0x0a, 0x0b, 0x5a,
	0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x7a, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x7a, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x7a, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x7a, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12,
	0x19, 0x2e, 0x7a, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x7a, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_zrpc_proto_rawDescOnce sync.Once
	file_protos_zrpc_proto_rawDescData = file_protos_zrpc_proto_rawDesc
)

func file_protos_zrpc_proto_rawDescGZIP() []byte {
	file_protos_zrpc_proto_rawDescOnce.Do(func() {
		file_protos_zrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_zrpc_proto_rawDescData)
	})
	return file_protos_zrpc_proto_rawDescData
}

var file_protos_zrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_zrpc_proto_goTypes = []interface{}{
	(*EventWithLocationAndType)(nil), // 0: zrpc.EventWithLocationAndType
	(*ListEventsRequest)(nil),        // 1: zrpc.ListEventsRequest
	(*ListEventsResponse)(nil),       // 2: zrpc.ListEventsResponse
	(*Event)(nil),                    // 3: zrpc.Event
	(*CreateEventRequest)(nil),       // 4: zrpc.CreateEventRequest
	(*CreateEventResponse)(nil),      // 5: zrpc.CreateEventResponse
	(*Coupon)(nil),                   // 6: zrpc.Coupon
	(*CreateCouponRequest)(nil),      // 7: zrpc.CreateCouponRequest
	(*CreateCouponResponse)(nil),     // 8: zrpc.CreateCouponResponse
}
var file_protos_zrpc_proto_depIdxs = []int32{
	0, // 0: zrpc.ListEventsResponse.events:type_name -> zrpc.EventWithLocationAndType
	3, // 1: zrpc.CreateEventRequest.event:type_name -> zrpc.Event
	6, // 2: zrpc.CreateCouponRequest.coupon:type_name -> zrpc.Coupon
	1, // 3: zrpc.ZrpcService.ListEvents:input_type -> zrpc.ListEventsRequest
	4, // 4: zrpc.ZrpcService.CreateEvent:input_type -> zrpc.CreateEventRequest
	7, // 5: zrpc.ZrpcService.CreateCoupon:input_type -> zrpc.CreateCouponRequest
	2, // 6: zrpc.ZrpcService.ListEvents:output_type -> zrpc.ListEventsResponse
	5, // 7: zrpc.ZrpcService.CreateEvent:output_type -> zrpc.CreateEventResponse
	8, // 8: zrpc.ZrpcService.CreateCoupon:output_type -> zrpc.CreateCouponResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_zrpc_proto_init() }
func file_protos_zrpc_proto_init() {
	if File_protos_zrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_zrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventWithLocationAndType); i {
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
		file_protos_zrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequest); i {
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
		file_protos_zrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsResponse); i {
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
		file_protos_zrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_protos_zrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_protos_zrpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_protos_zrpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coupon); i {
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
		file_protos_zrpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCouponRequest); i {
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
		file_protos_zrpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCouponResponse); i {
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
			RawDescriptor: file_protos_zrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_zrpc_proto_goTypes,
		DependencyIndexes: file_protos_zrpc_proto_depIdxs,
		MessageInfos:      file_protos_zrpc_proto_msgTypes,
	}.Build()
	File_protos_zrpc_proto = out.File
	file_protos_zrpc_proto_rawDesc = nil
	file_protos_zrpc_proto_goTypes = nil
	file_protos_zrpc_proto_depIdxs = nil
}
