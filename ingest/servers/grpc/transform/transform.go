package transform

import (
	"fmt"
	cpb "ingest/internal/opentelemetry-proto-gen/common/v1"
	mv "ingest/internal/opentelemetry-proto-gen/metrics/v1"
)

const (
	GaugeInt64                  = "GAUGE_INT64"
	GaugeDouble                 = "GAUGE_DOUBLE"
	GaugeDistribution           = "GAUGE_DISTRIBUTION"
	GaugeCumulativeInt          = "CUMULATIVE_INT64"
	GaugeCumulativeDouble       = "CUMULATIVE_DOUBLE"
	GaugeCumulativeDistribution = "CUMULATIVE_DISTRIBUTION"
)

type Labels map[string]string

type SourceInfo struct {
	Name    string
	Version string
}

type DescriptorInfo struct {
	Name        string
	Description string
	Unit        string
	Type        string
	Temporality string
	IsMonotonic bool
}

type RDLPoint struct {
	Source         *SourceInfo
	ResourceLabels Labels
	Descriptor     *DescriptorInfo
	TimeStamp      uint64
	Value          float64
}

type OTDataPoint interface {
	GetLabels() []*cpb.StringKeyValue
	GetTimeUnixNano() uint64
}

func stringify(val *cpb.AnyValue) string {
	switch val.Value.(type) {
	case *cpb.AnyValue_BoolValue:
		return fmt.Sprintf("%s", val.GetBoolValue())
	case *cpb.AnyValue_DoubleValue:
		return fmt.Sprintf("%f", val.GetDoubleValue())
	case *cpb.AnyValue_IntValue:
		return fmt.Sprintf("%d", val.GetIntValue())
	case *cpb.AnyValue_StringValue:
		return val.GetStringValue()
	default:
		return ""
	}
}

func dataPointToRDL(resourceAttrs []*cpb.KeyValue, info *SourceInfo,
	descriptor *DescriptorInfo, labels []*cpb.StringKeyValue, value float64, timeStamp uint64) *RDLPoint {
	rdp := &RDLPoint{
		Value:      value,
		TimeStamp:  timeStamp,
		Source:     info,
		Descriptor: descriptor,
	}
	rdp.ResourceLabels = make(Labels)
	for _, kv := range labels {
		rdp.ResourceLabels[kv.GetKey()] = kv.GetValue()
	}
	for _, kv := range resourceAttrs {
		rdp.ResourceLabels[kv.GetKey()] = stringify(kv.GetValue())
	}
	return rdp
}

func descriptor(mx *mv.Metric, t string) *DescriptorInfo {
	return &DescriptorInfo{
		Name:        mx.GetName(),
		Description: mx.GetDescription(),
		Unit:        mx.GetUnit(),
		Type:        t,
	}
}

func library(libraryInfo *cpb.InstrumentationLibrary) *SourceInfo {
	return &SourceInfo{
		Name:    libraryInfo.GetName(),
		Version: libraryInfo.GetVersion(),
	}
}

func GaugeDoubleToRDL(mx *mv.Metric, resourceInfo []*cpb.KeyValue, libraryInfo *cpb.InstrumentationLibrary, rdl chan<- *RDLPoint) error {
	descriptor := descriptor(mx, GaugeDouble)
	info := library(libraryInfo)
	for _, dp := range mx.GetDoubleGauge().GetDataPoints() {
		rdl <- dataPointToRDL(resourceInfo, info, descriptor, dp.GetLabels(), dp.GetValue(), dp.GetTimeUnixNano())
	}
	return nil
}

func GaugeIntToRDL(mx *mv.Metric, resourceInfo []*cpb.KeyValue, libraryInfo *cpb.InstrumentationLibrary, rdl chan<- *RDLPoint) error {
	descriptor := descriptor(mx, GaugeInt64)
	info := library(libraryInfo)
	for _, dp := range mx.GetIntGauge().GetDataPoints() {
		rdl <- dataPointToRDL(resourceInfo, info, descriptor, dp.GetLabels(), float64(dp.GetValue()), dp.GetTimeUnixNano())
	}
	return nil
}

func CumulativeIntToRDL(mx *mv.Metric, resourceInfo []*cpb.KeyValue, libraryInfo *cpb.InstrumentationLibrary, rdl chan<- *RDLPoint) error {
	descriptor := descriptor(mx, GaugeCumulativeInt)
	info := library(libraryInfo)
	for _, dp := range mx.GetIntSum().GetDataPoints() {
		descriptor.Temporality = mx.GetIntSum().GetAggregationTemporality().String()
		descriptor.IsMonotonic = mx.GetIntSum().GetIsMonotonic()
		rdl <- dataPointToRDL(resourceInfo, info, descriptor, dp.GetLabels(), float64(dp.GetValue()), dp.GetTimeUnixNano())
	}
	return nil
}

func CumulativeDoubleToRDL(mx *mv.Metric, resourceInfo []*cpb.KeyValue, libraryInfo *cpb.InstrumentationLibrary, rdl chan<- *RDLPoint) error {
	descriptor := descriptor(mx, GaugeCumulativeDouble)
	info := library(libraryInfo)
	for _, dp := range mx.GetDoubleSum().GetDataPoints() {
		descriptor.Temporality = mx.GetDoubleSum().GetAggregationTemporality().String()
		descriptor.IsMonotonic = mx.GetDoubleSum().GetIsMonotonic()
		rdl <- dataPointToRDL(resourceInfo, info, descriptor, dp.GetLabels(), dp.GetValue(), dp.GetTimeUnixNano())
	}
	return nil
}
