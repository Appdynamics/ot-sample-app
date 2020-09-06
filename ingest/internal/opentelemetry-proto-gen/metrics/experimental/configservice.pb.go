// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/metrics/experimental/configservice.proto

package experimental

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "ingest/internal/opentelemetry-proto-gen/resource/v1"
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

type MetricConfigRequest struct {
	// Required. The resource for which configuration should be returned.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// Optional. The value of MetricConfigResponse.fingerprint for the last
	// configuration that the caller received and successfully applied.
	LastKnownFingerprint []byte   `protobuf:"bytes,2,opt,name=last_known_fingerprint,json=lastKnownFingerprint,proto3" json:"last_known_fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricConfigRequest) Reset()         { *m = MetricConfigRequest{} }
func (m *MetricConfigRequest) String() string { return proto.CompactTextString(m) }
func (*MetricConfigRequest) ProtoMessage()    {}
func (*MetricConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b5d4ea55caf90b, []int{0}
}
func (m *MetricConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricConfigRequest.Unmarshal(m, b)
}
func (m *MetricConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricConfigRequest.Marshal(b, m, deterministic)
}
func (m *MetricConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricConfigRequest.Merge(m, src)
}
func (m *MetricConfigRequest) XXX_Size() int {
	return xxx_messageInfo_MetricConfigRequest.Size(m)
}
func (m *MetricConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricConfigRequest proto.InternalMessageInfo

func (m *MetricConfigRequest) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *MetricConfigRequest) GetLastKnownFingerprint() []byte {
	if m != nil {
		return m.LastKnownFingerprint
	}
	return nil
}

type MetricConfigResponse struct {
	// Optional. The fingerprint associated with this MetricConfigResponse. Each
	// change in configs yields a different fingerprint. The resource SHOULD copy
	// this value to MetricConfigRequest.last_known_fingerprint for the next
	// configuration request. If there are no changes between fingerprint and
	// MetricConfigRequest.last_known_fingerprint, then all other fields besides
	// fingerprint in the response are optional, or the same as the last update if
	// present.
	//
	// The exact mechanics of generating the fingerprint is up to the
	// implementation. However, a fingerprint must be deterministically determined
	// by the configurations -- the same configuration will generate the same
	// fingerprint on any instance of an implementation. Hence using a timestamp is
	// unacceptable, but a deterministic hash is fine.
	Fingerprint []byte `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// A single metric may match multiple schedules. In such cases, the schedule
	// that specifies the smallest period is applied.
	//
	// Note, for optimization purposes, it is recommended to use as few schedules
	// as possible to capture all required metric updates. Where you can be
	// conservative, do take full advantage of the inclusion/exclusion patterns to
	// capture as much of your targeted metrics.
	Schedules []*MetricConfigResponse_Schedule `protobuf:"bytes,2,rep,name=schedules,proto3" json:"schedules,omitempty"`
	// Optional. The client is suggested to wait this long (in seconds) before
	// pinging the configuration service again.
	SuggestedWaitTimeSec int32    `protobuf:"varint,3,opt,name=suggested_wait_time_sec,json=suggestedWaitTimeSec,proto3" json:"suggested_wait_time_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricConfigResponse) Reset()         { *m = MetricConfigResponse{} }
func (m *MetricConfigResponse) String() string { return proto.CompactTextString(m) }
func (*MetricConfigResponse) ProtoMessage()    {}
func (*MetricConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b5d4ea55caf90b, []int{1}
}
func (m *MetricConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricConfigResponse.Unmarshal(m, b)
}
func (m *MetricConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricConfigResponse.Marshal(b, m, deterministic)
}
func (m *MetricConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricConfigResponse.Merge(m, src)
}
func (m *MetricConfigResponse) XXX_Size() int {
	return xxx_messageInfo_MetricConfigResponse.Size(m)
}
func (m *MetricConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetricConfigResponse proto.InternalMessageInfo

func (m *MetricConfigResponse) GetFingerprint() []byte {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *MetricConfigResponse) GetSchedules() []*MetricConfigResponse_Schedule {
	if m != nil {
		return m.Schedules
	}
	return nil
}

func (m *MetricConfigResponse) GetSuggestedWaitTimeSec() int32 {
	if m != nil {
		return m.SuggestedWaitTimeSec
	}
	return 0
}

// A Schedule is used to apply a particular scheduling configuration to
// a metric. If a metric name matches a schedule's patterns, then the metric
// adopts the configuration specified by the schedule.
type MetricConfigResponse_Schedule struct {
	// Metrics with names that match a rule in the inclusion_patterns are
	// targeted by this schedule. Metrics that match the exclusion_patterns
	// are not targeted for this schedule, even if they match an inclusion
	// pattern.
	ExclusionPatterns []*MetricConfigResponse_Schedule_Pattern `protobuf:"bytes,1,rep,name=exclusion_patterns,json=exclusionPatterns,proto3" json:"exclusion_patterns,omitempty"`
	InclusionPatterns []*MetricConfigResponse_Schedule_Pattern `protobuf:"bytes,2,rep,name=inclusion_patterns,json=inclusionPatterns,proto3" json:"inclusion_patterns,omitempty"`
	// Describes the collection period for each metric in seconds.
	// A period of 0 means to not export.
	PeriodSec            int32    `protobuf:"varint,3,opt,name=period_sec,json=periodSec,proto3" json:"period_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricConfigResponse_Schedule) Reset()         { *m = MetricConfigResponse_Schedule{} }
func (m *MetricConfigResponse_Schedule) String() string { return proto.CompactTextString(m) }
func (*MetricConfigResponse_Schedule) ProtoMessage()    {}
func (*MetricConfigResponse_Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b5d4ea55caf90b, []int{1, 0}
}
func (m *MetricConfigResponse_Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricConfigResponse_Schedule.Unmarshal(m, b)
}
func (m *MetricConfigResponse_Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricConfigResponse_Schedule.Marshal(b, m, deterministic)
}
func (m *MetricConfigResponse_Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricConfigResponse_Schedule.Merge(m, src)
}
func (m *MetricConfigResponse_Schedule) XXX_Size() int {
	return xxx_messageInfo_MetricConfigResponse_Schedule.Size(m)
}
func (m *MetricConfigResponse_Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricConfigResponse_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_MetricConfigResponse_Schedule proto.InternalMessageInfo

func (m *MetricConfigResponse_Schedule) GetExclusionPatterns() []*MetricConfigResponse_Schedule_Pattern {
	if m != nil {
		return m.ExclusionPatterns
	}
	return nil
}

func (m *MetricConfigResponse_Schedule) GetInclusionPatterns() []*MetricConfigResponse_Schedule_Pattern {
	if m != nil {
		return m.InclusionPatterns
	}
	return nil
}

func (m *MetricConfigResponse_Schedule) GetPeriodSec() int32 {
	if m != nil {
		return m.PeriodSec
	}
	return 0
}

// A light-weight pattern that can match 1 or more
// metrics, for which this schedule will apply. The string is used to
// match against metric names. It should not exceed 100k characters.
type MetricConfigResponse_Schedule_Pattern struct {
	// Types that are valid to be assigned to Match:
	//	*MetricConfigResponse_Schedule_Pattern_Equals
	//	*MetricConfigResponse_Schedule_Pattern_StartsWith
	Match                isMetricConfigResponse_Schedule_Pattern_Match `protobuf_oneof:"match"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *MetricConfigResponse_Schedule_Pattern) Reset()         { *m = MetricConfigResponse_Schedule_Pattern{} }
func (m *MetricConfigResponse_Schedule_Pattern) String() string { return proto.CompactTextString(m) }
func (*MetricConfigResponse_Schedule_Pattern) ProtoMessage()    {}
func (*MetricConfigResponse_Schedule_Pattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b5d4ea55caf90b, []int{1, 0, 0}
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.Unmarshal(m, b)
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.Marshal(b, m, deterministic)
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.Merge(m, src)
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_Size() int {
	return xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.Size(m)
}
func (m *MetricConfigResponse_Schedule_Pattern) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricConfigResponse_Schedule_Pattern.DiscardUnknown(m)
}

var xxx_messageInfo_MetricConfigResponse_Schedule_Pattern proto.InternalMessageInfo

type isMetricConfigResponse_Schedule_Pattern_Match interface {
	isMetricConfigResponse_Schedule_Pattern_Match()
}

type MetricConfigResponse_Schedule_Pattern_Equals struct {
	Equals string `protobuf:"bytes,1,opt,name=equals,proto3,oneof" json:"equals,omitempty"`
}
type MetricConfigResponse_Schedule_Pattern_StartsWith struct {
	StartsWith string `protobuf:"bytes,2,opt,name=starts_with,json=startsWith,proto3,oneof" json:"starts_with,omitempty"`
}

func (*MetricConfigResponse_Schedule_Pattern_Equals) isMetricConfigResponse_Schedule_Pattern_Match() {
}
func (*MetricConfigResponse_Schedule_Pattern_StartsWith) isMetricConfigResponse_Schedule_Pattern_Match() {
}

func (m *MetricConfigResponse_Schedule_Pattern) GetMatch() isMetricConfigResponse_Schedule_Pattern_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *MetricConfigResponse_Schedule_Pattern) GetEquals() string {
	if x, ok := m.GetMatch().(*MetricConfigResponse_Schedule_Pattern_Equals); ok {
		return x.Equals
	}
	return ""
}

func (m *MetricConfigResponse_Schedule_Pattern) GetStartsWith() string {
	if x, ok := m.GetMatch().(*MetricConfigResponse_Schedule_Pattern_StartsWith); ok {
		return x.StartsWith
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetricConfigResponse_Schedule_Pattern) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetricConfigResponse_Schedule_Pattern_Equals)(nil),
		(*MetricConfigResponse_Schedule_Pattern_StartsWith)(nil),
	}
}

func init() {
	proto.RegisterType((*MetricConfigRequest)(nil), "opentelemetry.proto.metrics.experimental.MetricConfigRequest")
	proto.RegisterType((*MetricConfigResponse)(nil), "opentelemetry.proto.metrics.experimental.MetricConfigResponse")
	proto.RegisterType((*MetricConfigResponse_Schedule)(nil), "opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule")
	proto.RegisterType((*MetricConfigResponse_Schedule_Pattern)(nil), "opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule.Pattern")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/metrics/experimental/configservice.proto", fileDescriptor_79b5d4ea55caf90b)
}

var fileDescriptor_79b5d4ea55caf90b = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xd9, 0x94, 0x7e, 0x4d, 0x2a, 0x21, 0x96, 0x08, 0xac, 0x48, 0x48, 0xa1, 0xa7, 0x20,
	0xd4, 0xb5, 0x1a, 0xe0, 0x56, 0x38, 0x04, 0x41, 0x91, 0x10, 0x6a, 0xe4, 0x20, 0x55, 0xe2, 0x62,
	0x19, 0x67, 0x9a, 0xac, 0x70, 0x76, 0xdd, 0xdd, 0x71, 0x52, 0x2e, 0x5c, 0xb9, 0x22, 0xde, 0x80,
	0x67, 0xe2, 0x6d, 0x38, 0xa1, 0x5d, 0xbb, 0xae, 0x23, 0x72, 0xa8, 0xf8, 0xb8, 0x8d, 0xe7, 0x3f,
	0xf3, 0xfb, 0x8f, 0xc7, 0xde, 0x85, 0x23, 0x9d, 0xa3, 0x22, 0xcc, 0x70, 0x8e, 0x64, 0x3e, 0x85,
	0xb9, 0xd1, 0xa4, 0x43, 0x17, 0xcb, 0xd4, 0x86, 0x78, 0x91, 0xa3, 0x91, 0x73, 0x54, 0x94, 0x64,
	0x61, 0xaa, 0xd5, 0x99, 0x9c, 0x5a, 0x34, 0x0b, 0x99, 0xa2, 0xf0, 0x85, 0xbc, 0xbf, 0xd2, 0x5d,
	0x26, 0x45, 0xd5, 0x2d, 0x9a, 0xdd, 0x5d, 0xb1, 0xce, 0xc7, 0xa0, 0xd5, 0x85, 0x49, 0x31, 0x5c,
	0x1c, 0xd6, 0x71, 0x09, 0xd9, 0xff, 0xc6, 0xe0, 0xce, 0x5b, 0x0f, 0x7a, 0xe1, 0x7d, 0x23, 0x3c,
	0x2f, 0xd0, 0x12, 0x7f, 0x09, 0x3b, 0x97, 0x95, 0x01, 0xeb, 0xb1, 0x7e, 0x7b, 0xf0, 0x50, 0xac,
	0x1b, 0xa2, 0xc6, 0x2d, 0x0e, 0x45, 0x54, 0xc5, 0x51, 0xdd, 0xca, 0x9f, 0xc0, 0xdd, 0x2c, 0xb1,
	0x14, 0x7f, 0x54, 0x7a, 0xa9, 0xe2, 0x33, 0xa9, 0xa6, 0x68, 0x72, 0x23, 0x15, 0x05, 0xad, 0x1e,
	0xeb, 0xef, 0x45, 0x1d, 0xa7, 0xbe, 0x71, 0xe2, 0xab, 0x2b, 0x6d, 0xff, 0xc7, 0x4d, 0xe8, 0xac,
	0x0e, 0x65, 0x73, 0xad, 0x2c, 0xf2, 0x1e, 0xb4, 0x9b, 0x0c, 0xe6, 0x19, 0xcd, 0x14, 0x47, 0xd8,
	0xb5, 0xe9, 0x0c, 0x27, 0x45, 0x86, 0x36, 0x68, 0xf5, 0x36, 0xfa, 0xed, 0xc1, 0xb1, 0xb8, 0xee,
	0xf6, 0xc4, 0x3a, 0x53, 0x31, 0xae, 0x78, 0xd1, 0x15, 0x99, 0x3f, 0x85, 0x7b, 0xb6, 0x98, 0x4e,
	0xd1, 0x12, 0x4e, 0xe2, 0x65, 0x22, 0x29, 0x26, 0x39, 0xc7, 0xd8, 0x62, 0x1a, 0x6c, 0xf4, 0x58,
	0x7f, 0x33, 0xea, 0xd4, 0xf2, 0x69, 0x22, 0xe9, 0x9d, 0x9c, 0xe3, 0x18, 0xd3, 0xee, 0xcf, 0x16,
	0xec, 0x5c, 0xe2, 0xf8, 0x67, 0xe0, 0x78, 0x91, 0x66, 0x85, 0x95, 0x5a, 0xc5, 0x79, 0x42, 0x84,
	0x46, 0xd9, 0x80, 0xf9, 0x99, 0x4f, 0xfe, 0xd1, 0xcc, 0x62, 0x54, 0x72, 0xa3, 0xdb, 0xb5, 0x55,
	0x95, 0xb1, 0xce, 0x5f, 0xaa, 0xdf, 0xfc, 0x5b, 0xff, 0xc9, 0xbf, 0xb6, 0xaa, 0xfd, 0xef, 0x03,
	0x38, 0x8c, 0x9e, 0x34, 0xd6, 0xb6, 0x5b, 0x66, 0xdc, 0xae, 0x4e, 0x60, 0xbb, 0x2a, 0xe5, 0x01,
	0x6c, 0xe1, 0x79, 0x91, 0x64, 0xd6, 0x7f, 0xf1, 0xdd, 0xd7, 0x37, 0xa2, 0xea, 0x99, 0x3f, 0x80,
	0xb6, 0xa5, 0xc4, 0x90, 0x8d, 0x97, 0x92, 0x66, 0xfe, 0xa7, 0x72, 0x32, 0x94, 0xc9, 0x53, 0x49,
	0xb3, 0xe1, 0x36, 0x6c, 0xce, 0x13, 0x4a, 0x67, 0x83, 0xef, 0x0c, 0xf6, 0x9a, 0xc3, 0xf2, 0xaf,
	0x0c, 0x6e, 0x1d, 0x23, 0xad, 0xe4, 0x9e, 0xfd, 0xe9, 0x8b, 0xfb, 0x63, 0xd3, 0x7d, 0xfe, 0x77,
	0x7b, 0x1b, 0x7e, 0x61, 0xf0, 0x48, 0xea, 0x6b, 0x43, 0x86, 0x41, 0x93, 0x32, 0x2e, 0xef, 0x8c,
	0x91, 0x2b, 0x1f, 0xb1, 0xf7, 0x47, 0xee, 0x58, 0x58, 0x0a, 0xa5, 0x72, 0x4b, 0x4c, 0xb2, 0x70,
	0x85, 0x7a, 0xe0, 0xa9, 0x07, 0x53, 0x54, 0x6b, 0xaf, 0xa1, 0x0f, 0x5b, 0x5e, 0x7f, 0xfc, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x25, 0x65, 0x3e, 0xa0, 0xb9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricConfigClient is the client API for MetricConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricConfigClient interface {
	GetMetricConfig(ctx context.Context, in *MetricConfigRequest, opts ...grpc.CallOption) (*MetricConfigResponse, error)
}

type metricConfigClient struct {
	cc *grpc.ClientConn
}

func NewMetricConfigClient(cc *grpc.ClientConn) MetricConfigClient {
	return &metricConfigClient{cc}
}

func (c *metricConfigClient) GetMetricConfig(ctx context.Context, in *MetricConfigRequest, opts ...grpc.CallOption) (*MetricConfigResponse, error) {
	out := new(MetricConfigResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.metrics.experimental.MetricConfig/GetMetricConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricConfigServer is the server API for MetricConfig service.
type MetricConfigServer interface {
	GetMetricConfig(context.Context, *MetricConfigRequest) (*MetricConfigResponse, error)
}

// UnimplementedMetricConfigServer can be embedded to have forward compatible implementations.
type UnimplementedMetricConfigServer struct {
}

func (*UnimplementedMetricConfigServer) GetMetricConfig(ctx context.Context, req *MetricConfigRequest) (*MetricConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricConfig not implemented")
}

func RegisterMetricConfigServer(s *grpc.Server, srv MetricConfigServer) {
	s.RegisterService(&_MetricConfig_serviceDesc, srv)
}

func _MetricConfig_GetMetricConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricConfigServer).GetMetricConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.metrics.experimental.MetricConfig/GetMetricConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricConfigServer).GetMetricConfig(ctx, req.(*MetricConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.metrics.experimental.MetricConfig",
	HandlerType: (*MetricConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetricConfig",
			Handler:    _MetricConfig_GetMetricConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/metrics/experimental/configservice.proto",
}
