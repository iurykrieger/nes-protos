// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protos/platformProduct.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type PlatformTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PlatformTag) Reset() {
	*x = PlatformTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_platformProduct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformTag) ProtoMessage() {}

func (x *PlatformTag) ProtoReflect() protoreflect.Message {
	mi := &file_protos_platformProduct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformTag.ProtoReflect.Descriptor instead.
func (*PlatformTag) Descriptor() ([]byte, []int) {
	return file_protos_platformProduct_proto_rawDescGZIP(), []int{0}
}

func (x *PlatformTag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatformTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PlatformSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *PlatformSpec) Reset() {
	*x = PlatformSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_platformProduct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformSpec) ProtoMessage() {}

func (x *PlatformSpec) ProtoReflect() protoreflect.Message {
	mi := &file_protos_platformProduct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformSpec.ProtoReflect.Descriptor instead.
func (*PlatformSpec) Descriptor() ([]byte, []int) {
	return file_protos_platformProduct_proto_rawDescGZIP(), []int{1}
}

func (x *PlatformSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatformSpec) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type PlatformInstallment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PlatformInstallment) Reset() {
	*x = PlatformInstallment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_platformProduct_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformInstallment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformInstallment) ProtoMessage() {}

func (x *PlatformInstallment) ProtoReflect() protoreflect.Message {
	mi := &file_protos_platformProduct_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformInstallment.ProtoReflect.Descriptor instead.
func (*PlatformInstallment) Descriptor() ([]byte, []int) {
	return file_protos_platformProduct_proto_rawDescGZIP(), []int{2}
}

func (x *PlatformInstallment) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PlatformInstallment) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type PlatformSkuProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string               `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Status      string               `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Price       float32              `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	OldPrice    float32              `protobuf:"fixed32,4,opt,name=old_price,json=oldPrice,proto3" json:"old_price,omitempty"`
	Installment *PlatformInstallment `protobuf:"bytes,5,opt,name=installment,proto3" json:"installment,omitempty"`
	EanCode     string               `protobuf:"bytes,6,opt,name=ean_code,json=eanCode,proto3" json:"ean_code,omitempty"`
	Volume      string               `protobuf:"bytes,7,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *PlatformSkuProperties) Reset() {
	*x = PlatformSkuProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_platformProduct_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformSkuProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformSkuProperties) ProtoMessage() {}

func (x *PlatformSkuProperties) ProtoReflect() protoreflect.Message {
	mi := &file_protos_platformProduct_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformSkuProperties.ProtoReflect.Descriptor instead.
func (*PlatformSkuProperties) Descriptor() ([]byte, []int) {
	return file_protos_platformProduct_proto_rawDescGZIP(), []int{3}
}

func (x *PlatformSkuProperties) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PlatformSkuProperties) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PlatformSkuProperties) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PlatformSkuProperties) GetOldPrice() float32 {
	if x != nil {
		return x.OldPrice
	}
	return 0
}

func (x *PlatformSkuProperties) GetInstallment() *PlatformInstallment {
	if x != nil {
		return x.Installment
	}
	return nil
}

func (x *PlatformSkuProperties) GetEanCode() string {
	if x != nil {
		return x.EanCode
	}
	return ""
}

func (x *PlatformSkuProperties) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

type PlatformSku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku        string                 `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Properties *PlatformSkuProperties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	Specs      map[string]string      `protobuf:"bytes,3,rep,name=specs,proto3" json:"specs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PlatformSku) Reset() {
	*x = PlatformSku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_platformProduct_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformSku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformSku) ProtoMessage() {}

func (x *PlatformSku) ProtoReflect() protoreflect.Message {
	mi := &file_protos_platformProduct_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformSku.ProtoReflect.Descriptor instead.
func (*PlatformSku) Descriptor() ([]byte, []int) {
	return file_protos_platformProduct_proto_rawDescGZIP(), []int{4}
}

func (x *PlatformSku) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *PlatformSku) GetProperties() *PlatformSkuProperties {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *PlatformSku) GetSpecs() map[string]string {
	if x != nil {
		return x.Specs
	}
	return nil
}

type PlatformCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Parents []string `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
}

func (x *PlatformCategory) Reset() {
	*x = PlatformCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_platformProduct_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformCategory) ProtoMessage() {}

func (x *PlatformCategory) ProtoReflect() protoreflect.Message {
	mi := &file_protos_platformProduct_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformCategory.ProtoReflect.Descriptor instead.
func (*PlatformCategory) Descriptor() ([]byte, []int) {
	return file_protos_platformProduct_proto_rawDescGZIP(), []int{5}
}

func (x *PlatformCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatformCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlatformCategory) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

type PlatformProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status      string                   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Name        string                   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Url         string                   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Price       float32                  `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	OldPrice    float32                  `protobuf:"fixed32,6,opt,name=old_price,json=oldPrice,proto3" json:"old_price,omitempty"`
	Images      map[string]string        `protobuf:"bytes,7,rep,name=images,proto3" json:"images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Skus        []*PlatformSku           `protobuf:"bytes,8,rep,name=skus,proto3" json:"skus,omitempty"`
	Tags        []*PlatformTag           `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	Installment *PlatformInstallment     `protobuf:"bytes,10,opt,name=installment,proto3" json:"installment,omitempty"`
	Specs       map[string]*PlatformSpec `protobuf:"bytes,11,rep,name=specs,proto3" json:"specs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Categories  []*PlatformCategory      `protobuf:"bytes,12,rep,name=categories,proto3" json:"categories,omitempty"`
	Details     map[string]*any.Any      `protobuf:"bytes,13,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description string                   `protobuf:"bytes,14,opt,name=description,proto3" json:"description,omitempty"`
	EanCode     string                   `protobuf:"bytes,15,opt,name=ean_code,json=eanCode,proto3" json:"ean_code,omitempty"`
	Brand       string                   `protobuf:"bytes,16,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *PlatformProduct) Reset() {
	*x = PlatformProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_platformProduct_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformProduct) ProtoMessage() {}

func (x *PlatformProduct) ProtoReflect() protoreflect.Message {
	mi := &file_protos_platformProduct_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformProduct.ProtoReflect.Descriptor instead.
func (*PlatformProduct) Descriptor() ([]byte, []int) {
	return file_protos_platformProduct_proto_rawDescGZIP(), []int{6}
}

func (x *PlatformProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatformProduct) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PlatformProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlatformProduct) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PlatformProduct) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PlatformProduct) GetOldPrice() float32 {
	if x != nil {
		return x.OldPrice
	}
	return 0
}

func (x *PlatformProduct) GetImages() map[string]string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *PlatformProduct) GetSkus() []*PlatformSku {
	if x != nil {
		return x.Skus
	}
	return nil
}

func (x *PlatformProduct) GetTags() []*PlatformTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PlatformProduct) GetInstallment() *PlatformInstallment {
	if x != nil {
		return x.Installment
	}
	return nil
}

func (x *PlatformProduct) GetSpecs() map[string]*PlatformSpec {
	if x != nil {
		return x.Specs
	}
	return nil
}

func (x *PlatformProduct) GetCategories() []*PlatformCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *PlatformProduct) GetDetails() map[string]*any.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *PlatformProduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PlatformProduct) GetEanCode() string {
	if x != nil {
		return x.EanCode
	}
	return ""
}

func (x *PlatformProduct) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

var File_protos_platformProduct_proto protoreflect.FileDescriptor

var file_protos_platformProduct_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x31, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x61, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x41, 0x0a, 0x13, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xe6, 0x01,
	0x0a, 0x15, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x6b, 0x75, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x61, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x53, 0x6b, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x6b,
	0x75, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x6b, 0x75, 0x2e, 0x53, 0x70, 0x65, 0x63,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73, 0x1a, 0x38, 0x0a,
	0x0a, 0x53, 0x70, 0x65, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xc4, 0x06, 0x0a, 0x0f, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x73,
	0x6b, 0x75, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x6b, 0x75, 0x52, 0x04,
	0x73, 0x6b, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3d, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x05,
	0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x73, 0x70, 0x65, 0x63, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x3e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x61, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e,
	0x0a, 0x0a, 0x53, 0x70, 0x65, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50,
	0x0a, 0x0c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x75, 0x72, 0x79, 0x6b, 0x72, 0x69, 0x65, 0x67, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x73, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_platformProduct_proto_rawDescOnce sync.Once
	file_protos_platformProduct_proto_rawDescData = file_protos_platformProduct_proto_rawDesc
)

func file_protos_platformProduct_proto_rawDescGZIP() []byte {
	file_protos_platformProduct_proto_rawDescOnce.Do(func() {
		file_protos_platformProduct_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_platformProduct_proto_rawDescData)
	})
	return file_protos_platformProduct_proto_rawDescData
}

var file_protos_platformProduct_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protos_platformProduct_proto_goTypes = []interface{}{
	(*PlatformTag)(nil),           // 0: protos.PlatformTag
	(*PlatformSpec)(nil),          // 1: protos.PlatformSpec
	(*PlatformInstallment)(nil),   // 2: protos.PlatformInstallment
	(*PlatformSkuProperties)(nil), // 3: protos.PlatformSkuProperties
	(*PlatformSku)(nil),           // 4: protos.PlatformSku
	(*PlatformCategory)(nil),      // 5: protos.PlatformCategory
	(*PlatformProduct)(nil),       // 6: protos.PlatformProduct
	nil,                           // 7: protos.PlatformSku.SpecsEntry
	nil,                           // 8: protos.PlatformProduct.ImagesEntry
	nil,                           // 9: protos.PlatformProduct.SpecsEntry
	nil,                           // 10: protos.PlatformProduct.DetailsEntry
	(*any.Any)(nil),               // 11: google.protobuf.Any
}
var file_protos_platformProduct_proto_depIdxs = []int32{
	2,  // 0: protos.PlatformSkuProperties.installment:type_name -> protos.PlatformInstallment
	3,  // 1: protos.PlatformSku.properties:type_name -> protos.PlatformSkuProperties
	7,  // 2: protos.PlatformSku.specs:type_name -> protos.PlatformSku.SpecsEntry
	8,  // 3: protos.PlatformProduct.images:type_name -> protos.PlatformProduct.ImagesEntry
	4,  // 4: protos.PlatformProduct.skus:type_name -> protos.PlatformSku
	0,  // 5: protos.PlatformProduct.tags:type_name -> protos.PlatformTag
	2,  // 6: protos.PlatformProduct.installment:type_name -> protos.PlatformInstallment
	9,  // 7: protos.PlatformProduct.specs:type_name -> protos.PlatformProduct.SpecsEntry
	5,  // 8: protos.PlatformProduct.categories:type_name -> protos.PlatformCategory
	10, // 9: protos.PlatformProduct.details:type_name -> protos.PlatformProduct.DetailsEntry
	1,  // 10: protos.PlatformProduct.SpecsEntry.value:type_name -> protos.PlatformSpec
	11, // 11: protos.PlatformProduct.DetailsEntry.value:type_name -> google.protobuf.Any
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_protos_platformProduct_proto_init() }
func file_protos_platformProduct_proto_init() {
	if File_protos_platformProduct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_platformProduct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformTag); i {
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
		file_protos_platformProduct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformSpec); i {
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
		file_protos_platformProduct_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformInstallment); i {
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
		file_protos_platformProduct_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformSkuProperties); i {
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
		file_protos_platformProduct_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformSku); i {
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
		file_protos_platformProduct_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformCategory); i {
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
		file_protos_platformProduct_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformProduct); i {
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
			RawDescriptor: file_protos_platformProduct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_platformProduct_proto_goTypes,
		DependencyIndexes: file_protos_platformProduct_proto_depIdxs,
		MessageInfos:      file_protos_platformProduct_proto_msgTypes,
	}.Build()
	File_protos_platformProduct_proto = out.File
	file_protos_platformProduct_proto_rawDesc = nil
	file_protos_platformProduct_proto_goTypes = nil
	file_protos_platformProduct_proto_depIdxs = nil
}
