// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/string.proto

package envoy_type_matcher

import (
	fmt "fmt"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type StringMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_Regex
	//	*StringMatcher_SafeRegex
	MatchPattern         isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StringMatcher) Reset()         { *m = StringMatcher{} }
func (m *StringMatcher) String() string { return proto.CompactTextString(m) }
func (*StringMatcher) ProtoMessage()    {}
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc62c75a0f154e3, []int{0}
}

func (m *StringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMatcher.Unmarshal(m, b)
}
func (m *StringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMatcher.Marshal(b, m, deterministic)
}
func (m *StringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatcher.Merge(m, src)
}
func (m *StringMatcher) XXX_Size() int {
	return xxx_messageInfo_StringMatcher.Size(m)
}
func (m *StringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatcher proto.InternalMessageInfo

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatcher_Regex struct {
	Regex string `protobuf:"bytes,4,opt,name=regex,proto3,oneof"`
}

type StringMatcher_SafeRegex struct {
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Prefix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Suffix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Regex) isStringMatcher_MatchPattern() {}

func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *StringMatcher) GetExact() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatcher) GetPrefix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatcher) GetSuffix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

// Deprecated: Do not use.
func (m *StringMatcher) GetRegex() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := m.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_Regex)(nil),
		(*StringMatcher_SafeRegex)(nil),
	}
}

type ListStringMatcher struct {
	Patterns             []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListStringMatcher) Reset()         { *m = ListStringMatcher{} }
func (m *ListStringMatcher) String() string { return proto.CompactTextString(m) }
func (*ListStringMatcher) ProtoMessage()    {}
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc62c75a0f154e3, []int{1}
}

func (m *ListStringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStringMatcher.Unmarshal(m, b)
}
func (m *ListStringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStringMatcher.Marshal(b, m, deterministic)
}
func (m *ListStringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringMatcher.Merge(m, src)
}
func (m *ListStringMatcher) XXX_Size() int {
	return xxx_messageInfo_ListStringMatcher.Size(m)
}
func (m *ListStringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringMatcher proto.InternalMessageInfo

func (m *ListStringMatcher) GetPatterns() []*StringMatcher {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMatcher)(nil), "envoy.type.matcher.StringMatcher")
	proto.RegisterType((*ListStringMatcher)(nil), "envoy.type.matcher.ListStringMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/string.proto", fileDescriptor_1dc62c75a0f154e3) }

var fileDescriptor_1dc62c75a0f154e3 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x97, 0x76, 0x9d, 0x5d, 0xc6, 0x60, 0x06, 0xd1, 0xb2, 0x83, 0x76, 0xf3, 0xd2, 0x53,
	0x8b, 0xf3, 0x1f, 0xe4, 0xe2, 0x40, 0x85, 0x51, 0xaf, 0xc2, 0x88, 0xdb, 0xb7, 0x59, 0xd0, 0xa6,
	0x24, 0x71, 0x74, 0x37, 0xcf, 0x1e, 0xfd, 0x35, 0xfe, 0x0e, 0x7f, 0x88, 0x07, 0x4f, 0xd2, 0x93,
	0x24, 0xa9, 0xc2, 0x58, 0x6f, 0xfd, 0x78, 0x9f, 0xf7, 0xfd, 0xde, 0xe6, 0xc3, 0x67, 0x90, 0x6f,
	0xf8, 0x36, 0x51, 0xdb, 0x02, 0x92, 0x67, 0xa6, 0x16, 0x8f, 0x20, 0x12, 0xa9, 0x44, 0x96, 0xaf,
	0xe3, 0x42, 0x70, 0xc5, 0x09, 0x31, 0x40, 0xac, 0x81, 0xb8, 0x06, 0x86, 0xa7, 0x0d, 0x26, 0x01,
	0x6b, 0x28, 0xad, 0x67, 0x78, 0x6e, 0x75, 0x96, 0xe7, 0x5c, 0x31, 0x95, 0xf1, 0x5c, 0x26, 0x4b,
	0x28, 0x04, 0x2c, 0xcc, 0x50, 0x43, 0x27, 0x1b, 0xf6, 0x94, 0x2d, 0x99, 0x82, 0xe4, 0xef, 0xc3,
	0x0a, 0xe3, 0x0a, 0xe1, 0xfe, 0x9d, 0xa9, 0x70, 0x6b, 0xb3, 0xc9, 0x31, 0xf6, 0xa0, 0x64, 0x0b,
	0x15, 0xa0, 0x10, 0x45, 0xdd, 0x69, 0x2b, 0xb5, 0x23, 0x19, 0xe1, 0x4e, 0x21, 0x60, 0x95, 0x95,
	0x81, 0xa3, 0x05, 0x7a, 0x50, 0xd1, 0xb6, 0x70, 0x42, 0x34, 0x6d, 0xa5, 0xb5, 0xa0, 0x11, 0xf9,
	0xb2, 0xd2, 0x88, 0xbb, 0x87, 0x58, 0x81, 0x44, 0xd8, 0x33, 0xe5, 0x83, 0xb6, 0x21, 0x06, 0x15,
	0xf5, 0x84, 0x1b, 0xbd, 0xfa, 0x1f, 0x5f, 0xdf, 0x9f, 0x1e, 0x0a, 0x34, 0x6a, 0x01, 0x72, 0x8d,
	0xb1, 0x64, 0x2b, 0x98, 0x5b, 0xdc, 0x0b, 0x51, 0xd4, 0x9b, 0x84, 0xf1, 0xfe, 0x03, 0xc5, 0xa9,
	0x06, 0xea, 0xf6, 0xd4, 0xaf, 0xa8, 0xf7, 0x86, 0x9c, 0x81, 0x0e, 0xea, 0x6a, 0xbf, 0x51, 0xe9,
	0x11, 0xee, 0x1b, 0x7c, 0x5e, 0x30, 0xa5, 0x40, 0xe4, 0xc4, 0xfd, 0xa1, 0x68, 0x7c, 0x8f, 0x0f,
	0x6f, 0x32, 0xa9, 0x76, 0xff, 0xff, 0x0a, 0xfb, 0x35, 0x24, 0x03, 0x14, 0xba, 0x51, 0x6f, 0x32,
	0x6a, 0xda, 0xba, 0x63, 0x32, 0x6b, 0xdf, 0x91, 0xe3, 0xa3, 0xf4, 0xdf, 0x4c, 0x2f, 0x70, 0x98,
	0x71, 0x6b, 0x2d, 0x04, 0x2f, 0xb7, 0x0d, 0x29, 0xb4, 0x67, 0x63, 0x66, 0xfa, 0x16, 0x33, 0xf4,
	0xd0, 0x31, 0x47, 0xb9, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x27, 0x1f, 0xb4, 0xc7, 0x29, 0x02,
	0x00, 0x00,
}
