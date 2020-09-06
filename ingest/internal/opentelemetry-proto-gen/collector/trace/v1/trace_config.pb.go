// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/trace/v1/trace_config.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// How spans should be sampled:
// - Always off
// - Always on
// - Always follow the parent Span's decision (off if no parent).
type ConstantSampler_ConstantDecision int32

const (
	ConstantSampler_ALWAYS_OFF    ConstantSampler_ConstantDecision = 0
	ConstantSampler_ALWAYS_ON     ConstantSampler_ConstantDecision = 1
	ConstantSampler_ALWAYS_PARENT ConstantSampler_ConstantDecision = 2
)

var ConstantSampler_ConstantDecision_name = map[int32]string{
	0: "ALWAYS_OFF",
	1: "ALWAYS_ON",
	2: "ALWAYS_PARENT",
}

var ConstantSampler_ConstantDecision_value = map[string]int32{
	"ALWAYS_OFF":    0,
	"ALWAYS_ON":     1,
	"ALWAYS_PARENT": 2,
}

func (x ConstantSampler_ConstantDecision) String() string {
	return proto.EnumName(ConstantSampler_ConstantDecision_name, int32(x))
}

func (ConstantSampler_ConstantDecision) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{1, 0}
}

// Global configuration of the trace service. All fields must be specified, or
// the default (zero) values will be used for each type.
type TraceConfig struct {
	// The global default sampler used to make decisions on span sampling.
	//
	// Types that are valid to be assigned to Sampler:
	//	*TraceConfig_ConstantSampler
	//	*TraceConfig_ProbabilitySampler
	//	*TraceConfig_RateLimitingSampler
	Sampler isTraceConfig_Sampler `protobuf_oneof:"sampler"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributes int64 `protobuf:"varint,4,opt,name=max_number_of_attributes,json=maxNumberOfAttributes,proto3" json:"max_number_of_attributes,omitempty"`
	// The global default max number of annotation events per span.
	MaxNumberOfTimedEvents int64 `protobuf:"varint,5,opt,name=max_number_of_timed_events,json=maxNumberOfTimedEvents,proto3" json:"max_number_of_timed_events,omitempty"`
	// The global default max number of attributes per timed event.
	MaxNumberOfAttributesPerTimedEvent int64 `protobuf:"varint,6,opt,name=max_number_of_attributes_per_timed_event,json=maxNumberOfAttributesPerTimedEvent,proto3" json:"max_number_of_attributes_per_timed_event,omitempty"`
	// The global default max number of link entries per span.
	MaxNumberOfLinks int64 `protobuf:"varint,7,opt,name=max_number_of_links,json=maxNumberOfLinks,proto3" json:"max_number_of_links,omitempty"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributesPerLink int64    `protobuf:"varint,8,opt,name=max_number_of_attributes_per_link,json=maxNumberOfAttributesPerLink,proto3" json:"max_number_of_attributes_per_link,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{0}
}
func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceConfig.Unmarshal(m, b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
}
func (m *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(m, src)
}
func (m *TraceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceConfig.Size(m)
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

type isTraceConfig_Sampler interface {
	isTraceConfig_Sampler()
}

type TraceConfig_ConstantSampler struct {
	ConstantSampler *ConstantSampler `protobuf:"bytes,1,opt,name=constant_sampler,json=constantSampler,proto3,oneof" json:"constant_sampler,omitempty"`
}
type TraceConfig_ProbabilitySampler struct {
	ProbabilitySampler *ProbabilitySampler `protobuf:"bytes,2,opt,name=probability_sampler,json=probabilitySampler,proto3,oneof" json:"probability_sampler,omitempty"`
}
type TraceConfig_RateLimitingSampler struct {
	RateLimitingSampler *RateLimitingSampler `protobuf:"bytes,3,opt,name=rate_limiting_sampler,json=rateLimitingSampler,proto3,oneof" json:"rate_limiting_sampler,omitempty"`
}

func (*TraceConfig_ConstantSampler) isTraceConfig_Sampler()     {}
func (*TraceConfig_ProbabilitySampler) isTraceConfig_Sampler()  {}
func (*TraceConfig_RateLimitingSampler) isTraceConfig_Sampler() {}

func (m *TraceConfig) GetSampler() isTraceConfig_Sampler {
	if m != nil {
		return m.Sampler
	}
	return nil
}

func (m *TraceConfig) GetConstantSampler() *ConstantSampler {
	if x, ok := m.GetSampler().(*TraceConfig_ConstantSampler); ok {
		return x.ConstantSampler
	}
	return nil
}

func (m *TraceConfig) GetProbabilitySampler() *ProbabilitySampler {
	if x, ok := m.GetSampler().(*TraceConfig_ProbabilitySampler); ok {
		return x.ProbabilitySampler
	}
	return nil
}

func (m *TraceConfig) GetRateLimitingSampler() *RateLimitingSampler {
	if x, ok := m.GetSampler().(*TraceConfig_RateLimitingSampler); ok {
		return x.RateLimitingSampler
	}
	return nil
}

func (m *TraceConfig) GetMaxNumberOfAttributes() int64 {
	if m != nil {
		return m.MaxNumberOfAttributes
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfTimedEvents() int64 {
	if m != nil {
		return m.MaxNumberOfTimedEvents
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAttributesPerTimedEvent() int64 {
	if m != nil {
		return m.MaxNumberOfAttributesPerTimedEvent
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfLinks() int64 {
	if m != nil {
		return m.MaxNumberOfLinks
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAttributesPerLink() int64 {
	if m != nil {
		return m.MaxNumberOfAttributesPerLink
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TraceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TraceConfig_ConstantSampler)(nil),
		(*TraceConfig_ProbabilitySampler)(nil),
		(*TraceConfig_RateLimitingSampler)(nil),
	}
}

// Sampler that always makes a constant decision on span sampling.
type ConstantSampler struct {
	Decision             ConstantSampler_ConstantDecision `protobuf:"varint,1,opt,name=decision,proto3,enum=opentelemetry.proto.trace.v1.ConstantSampler_ConstantDecision" json:"decision,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ConstantSampler) Reset()         { *m = ConstantSampler{} }
func (m *ConstantSampler) String() string { return proto.CompactTextString(m) }
func (*ConstantSampler) ProtoMessage()    {}
func (*ConstantSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{1}
}
func (m *ConstantSampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConstantSampler.Unmarshal(m, b)
}
func (m *ConstantSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConstantSampler.Marshal(b, m, deterministic)
}
func (m *ConstantSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConstantSampler.Merge(m, src)
}
func (m *ConstantSampler) XXX_Size() int {
	return xxx_messageInfo_ConstantSampler.Size(m)
}
func (m *ConstantSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ConstantSampler.DiscardUnknown(m)
}

var xxx_messageInfo_ConstantSampler proto.InternalMessageInfo

func (m *ConstantSampler) GetDecision() ConstantSampler_ConstantDecision {
	if m != nil {
		return m.Decision
	}
	return ConstantSampler_ALWAYS_OFF
}

// Sampler that tries to uniformly sample traces with a given probability.
// The probability of sampling a trace is equal to that of the specified probability.
type ProbabilitySampler struct {
	// The desired probability of sampling. Must be within [0.0, 1.0].
	SamplingProbability  float64  `protobuf:"fixed64,1,opt,name=samplingProbability,proto3" json:"samplingProbability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbabilitySampler) Reset()         { *m = ProbabilitySampler{} }
func (m *ProbabilitySampler) String() string { return proto.CompactTextString(m) }
func (*ProbabilitySampler) ProtoMessage()    {}
func (*ProbabilitySampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{2}
}
func (m *ProbabilitySampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbabilitySampler.Unmarshal(m, b)
}
func (m *ProbabilitySampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbabilitySampler.Marshal(b, m, deterministic)
}
func (m *ProbabilitySampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbabilitySampler.Merge(m, src)
}
func (m *ProbabilitySampler) XXX_Size() int {
	return xxx_messageInfo_ProbabilitySampler.Size(m)
}
func (m *ProbabilitySampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbabilitySampler.DiscardUnknown(m)
}

var xxx_messageInfo_ProbabilitySampler proto.InternalMessageInfo

func (m *ProbabilitySampler) GetSamplingProbability() float64 {
	if m != nil {
		return m.SamplingProbability
	}
	return 0
}

// Sampler that tries to sample with a rate per time window.
type RateLimitingSampler struct {
	// Rate per second.
	Qps                  int64    `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitingSampler) Reset()         { *m = RateLimitingSampler{} }
func (m *RateLimitingSampler) String() string { return proto.CompactTextString(m) }
func (*RateLimitingSampler) ProtoMessage()    {}
func (*RateLimitingSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{3}
}
func (m *RateLimitingSampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitingSampler.Unmarshal(m, b)
}
func (m *RateLimitingSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitingSampler.Marshal(b, m, deterministic)
}
func (m *RateLimitingSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitingSampler.Merge(m, src)
}
func (m *RateLimitingSampler) XXX_Size() int {
	return xxx_messageInfo_RateLimitingSampler.Size(m)
}
func (m *RateLimitingSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitingSampler.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitingSampler proto.InternalMessageInfo

func (m *RateLimitingSampler) GetQps() int64 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func init() {
	proto.RegisterEnum("opentelemetry.proto.trace.v1.ConstantSampler_ConstantDecision", ConstantSampler_ConstantDecision_name, ConstantSampler_ConstantDecision_value)
	proto.RegisterType((*TraceConfig)(nil), "opentelemetry.proto.trace.v1.TraceConfig")
	proto.RegisterType((*ConstantSampler)(nil), "opentelemetry.proto.trace.v1.ConstantSampler")
	proto.RegisterType((*ProbabilitySampler)(nil), "opentelemetry.proto.trace.v1.ProbabilitySampler")
	proto.RegisterType((*RateLimitingSampler)(nil), "opentelemetry.proto.trace.v1.RateLimitingSampler")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/trace/v1/trace_config.proto", fileDescriptor_5936aa8fa6443e6f)
}

var fileDescriptor_5936aa8fa6443e6f = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x4e, 0xdb, 0x30,
	0x14, 0xc6, 0x09, 0x1d, 0xff, 0x0e, 0x02, 0x32, 0x57, 0x4c, 0xd1, 0x84, 0x34, 0x96, 0x9b, 0x71,
	0xd3, 0x86, 0xb2, 0x8b, 0x49, 0x5c, 0x4c, 0x6a, 0x81, 0x6e, 0x17, 0x55, 0x89, 0x42, 0xa5, 0x69,
	0xbd, 0x89, 0x9c, 0x70, 0x1a, 0x59, 0x4b, 0xec, 0xe0, 0x98, 0x0a, 0x1e, 0x60, 0x4f, 0xb4, 0x17,
	0x9c, 0xe2, 0x66, 0x69, 0x42, 0x4b, 0xa4, 0xdd, 0xf9, 0x9c, 0x2f, 0xdf, 0xef, 0xb3, 0x93, 0x13,
	0x83, 0x23, 0x52, 0xe4, 0x0a, 0x63, 0x4c, 0x50, 0xc9, 0x67, 0x27, 0x95, 0x42, 0x09, 0x47, 0x49,
	0x1a, 0xa2, 0x33, 0xef, 0x2d, 0x16, 0x7e, 0x28, 0xf8, 0x8c, 0x45, 0x5d, 0xad, 0x91, 0x93, 0x9a,
	0x61, 0xd1, 0xec, 0xea, 0xe7, 0xba, 0xf3, 0x9e, 0xfd, 0x7b, 0x0b, 0xf6, 0x27, 0x79, 0x71, 0xa5,
	0x3d, 0x64, 0x0a, 0x66, 0x28, 0x78, 0xa6, 0x28, 0x57, 0x7e, 0x46, 0x93, 0x34, 0x46, 0x69, 0x19,
	0xa7, 0xc6, 0xd9, 0xfe, 0x45, 0xa7, 0xdb, 0x04, 0xea, 0x5e, 0x15, 0xae, 0xbb, 0x85, 0xe9, 0xfb,
	0x86, 0x77, 0x14, 0xd6, 0x5b, 0x24, 0x84, 0x76, 0x2a, 0x45, 0x40, 0x03, 0x16, 0x33, 0xf5, 0x5c,
	0xe2, 0x37, 0x35, 0xfe, 0xbc, 0x19, 0xef, 0x2e, 0x8d, 0xcb, 0x04, 0x92, 0xae, 0x74, 0x49, 0x04,
	0xc7, 0x92, 0x2a, 0xf4, 0x63, 0x96, 0x30, 0xc5, 0x78, 0x54, 0xc6, 0xb4, 0x74, 0x4c, 0xaf, 0x39,
	0xc6, 0xa3, 0x0a, 0x47, 0x85, 0x73, 0x99, 0xd3, 0x96, 0xab, 0x6d, 0xf2, 0x05, 0xac, 0x84, 0x3e,
	0xf9, 0xfc, 0x31, 0x09, 0x50, 0xfa, 0x62, 0xe6, 0x53, 0xa5, 0x24, 0x0b, 0x1e, 0x15, 0x66, 0xd6,
	0x9b, 0x53, 0xe3, 0xac, 0xe5, 0x1d, 0x27, 0xf4, 0x69, 0xac, 0xe5, 0xdb, 0x59, 0xbf, 0x14, 0xc9,
	0x25, 0xbc, 0xaf, 0x1b, 0x15, 0x4b, 0xf0, 0xde, 0xc7, 0x39, 0x72, 0x95, 0x59, 0x5b, 0xda, 0xfa,
	0xae, 0x62, 0x9d, 0xe4, 0xf2, 0x8d, 0x56, 0xc9, 0x04, 0xce, 0x5e, 0x0b, 0xf5, 0x53, 0x94, 0x55,
	0x94, 0xb5, 0xad, 0x49, 0xf6, 0xda, 0x4d, 0xb8, 0x28, 0x97, 0x58, 0xd2, 0x81, 0x76, 0x9d, 0x1a,
	0x33, 0xfe, 0x2b, 0xb3, 0x76, 0x34, 0xc0, 0xac, 0x00, 0x46, 0x79, 0x9f, 0x7c, 0x83, 0x8f, 0x8d,
	0x9b, 0xc8, 0xdd, 0xd6, 0xae, 0x36, 0x9f, 0xbc, 0x96, 0x9e, 0x93, 0x06, 0x7b, 0xb0, 0x53, 0x7c,
	0x1d, 0xfb, 0x8f, 0x01, 0x47, 0x2f, 0x46, 0x88, 0x4c, 0x61, 0xf7, 0x1e, 0x43, 0x96, 0x31, 0xc1,
	0xf5, 0x0c, 0x1e, 0x5e, 0x7c, 0xfd, 0xaf, 0x19, 0x2c, 0xeb, 0xeb, 0x82, 0xe2, 0x95, 0x3c, 0xfb,
	0x1a, 0xcc, 0x97, 0x2a, 0x39, 0x04, 0xe8, 0x8f, 0x7e, 0xf4, 0x7f, 0xde, 0xf9, 0xb7, 0xc3, 0xa1,
	0xb9, 0x41, 0x0e, 0x60, 0xef, 0x5f, 0x3d, 0x36, 0x0d, 0xf2, 0x16, 0x0e, 0x8a, 0xd2, 0xed, 0x7b,
	0x37, 0xe3, 0x89, 0xb9, 0x69, 0x0f, 0x81, 0xac, 0x0e, 0x26, 0x39, 0x87, 0xb6, 0x3e, 0x16, 0xe3,
	0x51, 0x45, 0xd5, 0x47, 0x30, 0xbc, 0x75, 0x92, 0xfd, 0x09, 0xda, 0x6b, 0x26, 0x8f, 0x98, 0xd0,
	0x7a, 0x48, 0x33, 0x6d, 0x6c, 0x79, 0xf9, 0x72, 0xf0, 0x00, 0x1f, 0x98, 0x68, 0x7c, 0x09, 0x03,
	0xb3, 0xf2, 0x3b, 0xbb, 0xb9, 0xe4, 0x1a, 0xd3, 0x4b, 0xc6, 0x23, 0xcc, 0x94, 0xc3, 0xb8, 0x42,
	0xc9, 0x69, 0x5c, 0xbf, 0x44, 0x3a, 0x9a, 0xd0, 0x89, 0x90, 0x3b, 0xa1, 0x88, 0x63, 0x0c, 0x95,
	0x90, 0xe5, 0x95, 0x12, 0x6c, 0x6b, 0xf5, 0xf3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x67,
	0x65, 0x9a, 0x79, 0x04, 0x00, 0x00,
}
