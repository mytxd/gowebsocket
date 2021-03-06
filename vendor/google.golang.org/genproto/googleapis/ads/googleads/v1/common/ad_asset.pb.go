// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/ad_asset.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A text asset used inside an ad.
type AdTextAsset struct {
	// Asset text.
	Text *wrappers.StringValue `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The pinned field of the asset. This restricts the asset to only serve
	// within this field. Multiple assets can be pinned to the same field. An
	// asset that is unpinned or pinned to a different field will not serve in a
	// field where some other asset has been pinned.
	PinnedField          enums.ServedAssetFieldTypeEnum_ServedAssetFieldType `protobuf:"varint,2,opt,name=pinned_field,json=pinnedField,proto3,enum=google.ads.googleads.v1.enums.ServedAssetFieldTypeEnum_ServedAssetFieldType" json:"pinned_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *AdTextAsset) Reset()         { *m = AdTextAsset{} }
func (m *AdTextAsset) String() string { return proto.CompactTextString(m) }
func (*AdTextAsset) ProtoMessage()    {}
func (*AdTextAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1e7a8010a4d26a5, []int{0}
}

func (m *AdTextAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdTextAsset.Unmarshal(m, b)
}
func (m *AdTextAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdTextAsset.Marshal(b, m, deterministic)
}
func (m *AdTextAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdTextAsset.Merge(m, src)
}
func (m *AdTextAsset) XXX_Size() int {
	return xxx_messageInfo_AdTextAsset.Size(m)
}
func (m *AdTextAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_AdTextAsset.DiscardUnknown(m)
}

var xxx_messageInfo_AdTextAsset proto.InternalMessageInfo

func (m *AdTextAsset) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *AdTextAsset) GetPinnedField() enums.ServedAssetFieldTypeEnum_ServedAssetFieldType {
	if m != nil {
		return m.PinnedField
	}
	return enums.ServedAssetFieldTypeEnum_UNSPECIFIED
}

// An image asset used inside an ad.
type AdImageAsset struct {
	// The Asset resource name of this image.
	Asset                *wrappers.StringValue `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdImageAsset) Reset()         { *m = AdImageAsset{} }
func (m *AdImageAsset) String() string { return proto.CompactTextString(m) }
func (*AdImageAsset) ProtoMessage()    {}
func (*AdImageAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1e7a8010a4d26a5, []int{1}
}

func (m *AdImageAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdImageAsset.Unmarshal(m, b)
}
func (m *AdImageAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdImageAsset.Marshal(b, m, deterministic)
}
func (m *AdImageAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdImageAsset.Merge(m, src)
}
func (m *AdImageAsset) XXX_Size() int {
	return xxx_messageInfo_AdImageAsset.Size(m)
}
func (m *AdImageAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_AdImageAsset.DiscardUnknown(m)
}

var xxx_messageInfo_AdImageAsset proto.InternalMessageInfo

func (m *AdImageAsset) GetAsset() *wrappers.StringValue {
	if m != nil {
		return m.Asset
	}
	return nil
}

// A video asset used inside an ad.
type AdVideoAsset struct {
	// The Asset resource name of this video.
	Asset                *wrappers.StringValue `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdVideoAsset) Reset()         { *m = AdVideoAsset{} }
func (m *AdVideoAsset) String() string { return proto.CompactTextString(m) }
func (*AdVideoAsset) ProtoMessage()    {}
func (*AdVideoAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1e7a8010a4d26a5, []int{2}
}

func (m *AdVideoAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdVideoAsset.Unmarshal(m, b)
}
func (m *AdVideoAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdVideoAsset.Marshal(b, m, deterministic)
}
func (m *AdVideoAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdVideoAsset.Merge(m, src)
}
func (m *AdVideoAsset) XXX_Size() int {
	return xxx_messageInfo_AdVideoAsset.Size(m)
}
func (m *AdVideoAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_AdVideoAsset.DiscardUnknown(m)
}

var xxx_messageInfo_AdVideoAsset proto.InternalMessageInfo

func (m *AdVideoAsset) GetAsset() *wrappers.StringValue {
	if m != nil {
		return m.Asset
	}
	return nil
}

// A media bundle asset used inside an ad.
type AdMediaBundleAsset struct {
	// The Asset resource name of this media bundle.
	Asset                *wrappers.StringValue `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdMediaBundleAsset) Reset()         { *m = AdMediaBundleAsset{} }
func (m *AdMediaBundleAsset) String() string { return proto.CompactTextString(m) }
func (*AdMediaBundleAsset) ProtoMessage()    {}
func (*AdMediaBundleAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1e7a8010a4d26a5, []int{3}
}

func (m *AdMediaBundleAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdMediaBundleAsset.Unmarshal(m, b)
}
func (m *AdMediaBundleAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdMediaBundleAsset.Marshal(b, m, deterministic)
}
func (m *AdMediaBundleAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdMediaBundleAsset.Merge(m, src)
}
func (m *AdMediaBundleAsset) XXX_Size() int {
	return xxx_messageInfo_AdMediaBundleAsset.Size(m)
}
func (m *AdMediaBundleAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_AdMediaBundleAsset.DiscardUnknown(m)
}

var xxx_messageInfo_AdMediaBundleAsset proto.InternalMessageInfo

func (m *AdMediaBundleAsset) GetAsset() *wrappers.StringValue {
	if m != nil {
		return m.Asset
	}
	return nil
}

func init() {
	proto.RegisterType((*AdTextAsset)(nil), "google.ads.googleads.v1.common.AdTextAsset")
	proto.RegisterType((*AdImageAsset)(nil), "google.ads.googleads.v1.common.AdImageAsset")
	proto.RegisterType((*AdVideoAsset)(nil), "google.ads.googleads.v1.common.AdVideoAsset")
	proto.RegisterType((*AdMediaBundleAsset)(nil), "google.ads.googleads.v1.common.AdMediaBundleAsset")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/ad_asset.proto", fileDescriptor_a1e7a8010a4d26a5)
}

var fileDescriptor_a1e7a8010a4d26a5 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x6a, 0xdc, 0x30,
	0x10, 0xc7, 0xf1, 0xf6, 0xe3, 0xa0, 0x5d, 0x7a, 0xf0, 0x29, 0x84, 0x10, 0x82, 0x4f, 0xb9, 0x54,
	0xaa, 0xb7, 0x37, 0xe5, 0x24, 0xf7, 0x23, 0x2d, 0xb4, 0x10, 0x92, 0xe0, 0x43, 0x31, 0x2c, 0x4a,
	0x34, 0x11, 0x02, 0x5b, 0x12, 0x96, 0xbc, 0x4d, 0x5e, 0xa7, 0xc7, 0x9e, 0xfa, 0x1c, 0x7d, 0x91,
	0x42, 0x9f, 0xa2, 0x58, 0x63, 0xef, 0xa9, 0x5b, 0x4a, 0x4e, 0x1e, 0x6b, 0x7e, 0xf3, 0x9f, 0xff,
	0x0c, 0x43, 0x5e, 0x6a, 0xe7, 0x74, 0x0b, 0x4c, 0xaa, 0xc0, 0x30, 0x1c, 0xa3, 0x6d, 0xc9, 0x6e,
	0x5d, 0xd7, 0x39, 0xcb, 0xa4, 0xda, 0xc8, 0x10, 0x20, 0x52, 0xdf, 0xbb, 0xe8, 0xf2, 0x63, 0x64,
	0xa8, 0x54, 0x81, 0xee, 0x70, 0xba, 0x2d, 0x29, 0xe2, 0x87, 0x67, 0xfb, 0xe4, 0xc0, 0x0e, 0x5d,
	0x60, 0x01, 0xfa, 0x2d, 0x4c, 0x8a, 0x9b, 0x3b, 0x03, 0xad, 0xda, 0xc4, 0x07, 0x0f, 0x28, 0x7e,
	0x78, 0x34, 0x17, 0x7b, 0xc3, 0xa4, 0xb5, 0x2e, 0xca, 0x68, 0x9c, 0x0d, 0x53, 0x76, 0x6a, 0xcd,
	0xd2, 0xdf, 0xcd, 0x70, 0xc7, 0xbe, 0xf6, 0xd2, 0x7b, 0xe8, 0xa7, 0x7c, 0xf1, 0x23, 0x23, 0x4b,
	0xa1, 0xae, 0xe1, 0x3e, 0x8a, 0x51, 0x3e, 0x7f, 0x45, 0x9e, 0x46, 0xb8, 0x8f, 0x07, 0xd9, 0x49,
	0x76, 0xba, 0x5c, 0x1f, 0x4d, 0x76, 0xe9, 0x5c, 0x4e, 0xaf, 0x62, 0x6f, 0xac, 0xae, 0x65, 0x3b,
	0xc0, 0x65, 0x22, 0x73, 0x47, 0x56, 0xde, 0x58, 0x0b, 0x0a, 0xad, 0x1d, 0x2c, 0x4e, 0xb2, 0xd3,
	0x17, 0xeb, 0x4f, 0x74, 0xdf, 0xcc, 0x69, 0x26, 0x7a, 0x95, 0x66, 0x4a, 0x3d, 0xdf, 0x8f, 0x65,
	0xd7, 0x0f, 0x1e, 0xde, 0xd9, 0xa1, 0xfb, 0x6b, 0xe2, 0x72, 0x89, 0x1d, 0xd2, 0x43, 0x51, 0x91,
	0x95, 0x50, 0x1f, 0x3b, 0xa9, 0x01, 0x2d, 0xaf, 0xc9, 0xb3, 0xb4, 0x9a, 0xff, 0xf2, 0x8c, 0x28,
	0x6a, 0xd4, 0x46, 0x81, 0x7b, 0xbc, 0xc6, 0x07, 0x92, 0x0b, 0xf5, 0x19, 0x94, 0x91, 0xd5, 0x60,
	0x55, 0xfb, 0x78, 0x37, 0xd5, 0xaf, 0x8c, 0x14, 0xb7, 0xae, 0xa3, 0xff, 0x3e, 0x93, 0x6a, 0x25,
	0x70, 0x2f, 0x17, 0xa3, 0xd4, 0x45, 0xf6, 0xe5, 0xed, 0xc4, 0x6b, 0xd7, 0x4a, 0xab, 0xa9, 0xeb,
	0x35, 0xd3, 0x60, 0x53, 0xa3, 0xf9, 0x8c, 0xbc, 0x09, 0xfb, 0x8e, 0xf4, 0x0c, 0x3f, 0xdf, 0x16,
	0x4f, 0xce, 0x85, 0xf8, 0xbe, 0x38, 0x3e, 0x47, 0x31, 0xa1, 0x02, 0xc5, 0x70, 0x8c, 0xea, 0x92,
	0xbe, 0x49, 0xd8, 0xcf, 0x19, 0x68, 0x84, 0x0a, 0xcd, 0x0e, 0x68, 0xea, 0xb2, 0x41, 0xe0, 0xf7,
	0xa2, 0xc0, 0x57, 0xce, 0x85, 0x0a, 0x9c, 0xef, 0x10, 0xce, 0xeb, 0x92, 0x73, 0x84, 0x6e, 0x9e,
	0x27, 0x77, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x84, 0x22, 0x05, 0x41, 0x03, 0x00,
	0x00,
}
