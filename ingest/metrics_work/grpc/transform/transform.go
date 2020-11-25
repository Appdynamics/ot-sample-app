package transform

import (
	"fmt"
	cpb "ingest/internal/opentelemetry-proto-gen/common/v1"
	mv "ingest/internal/opentelemetry-proto-gen/metrics/v1"
	"strings"
)

const (
	GaugeInt64                  = "GAUGE_INT64"
	GaugeDouble                 = "GAUGE_DOUBLE"
	GaugeCumulativeInt          = "CUMULATIVE_INT64"
	GaugeCumulativeDouble       = "CUMULATIVE_DOUBLE"
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
	Value          int64
}


type MetricNameProcessorMap map[string]string

var (
	metricsNameProcessor = MetricNameProcessorMap{
		"system.memory.usage": "system.memory.usage.state",
		"system.network.io": "system.network.io.direction",
		"system.network.errors": "system.network.errors.direction",
		"system.network.dropped_packets": "system.network.dropped_packets.direction",
		"system.network.packets": "system.network.packets.direction",
	}
)

// TODO use a regex
func (m MetricNameProcessorMap) Processor(rdl *RDLPoint) (*RDLPoint, bool) {
	var changed bool
	if temp, ok := metricsNameProcessor[rdl.Descriptor.Name]; ok {
		labels := strings.Split(temp, ".")
		s, _  := rdl.ResourceLabels[labels[len(labels)-1]]
		labels = append(labels[:len(labels)-1], s)
		rdl.Descriptor.Name = strings.Join(labels, ".")
		changed = true
	}
	return rdl, changed
}

func stringify(val *cpb.AnyValue) string {
	switch val.Value.(type) {
	case *cpb.AnyValue_BoolValue:
		return fmt.Sprintf("%t", val.GetBoolValue())
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
	descriptor *DescriptorInfo, labels []*cpb.StringKeyValue, value int64, timeStamp uint64) *RDLPoint {
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
	for _, dp := range mx.GetDoubleGauge().GetDataPoints() {
		descriptor := descriptor(mx, GaugeDouble)
		info := library(libraryInfo)
		point, _ := metricsNameProcessor.Processor(dataPointToRDL(resourceInfo, info, descriptor, dp.GetLabels(),
			int64(dp.GetValue()), dp.GetTimeUnixNano()))
		rdl <- point
	}
	return nil
}

func GaugeIntToRDL(mx *mv.Metric, resourceInfo []*cpb.KeyValue, libraryInfo *cpb.InstrumentationLibrary, rdl chan<- *RDLPoint) error {
	for _, dp := range mx.GetIntGauge().GetDataPoints() {
		descriptor := descriptor(mx, GaugeInt64)
		info := library(libraryInfo)
		point, _ := metricsNameProcessor.Processor(dataPointToRDL(resourceInfo, info, descriptor, dp.GetLabels(),
			dp.GetValue(), dp.GetTimeUnixNano()))
		rdl <- point
	}
	return nil
}

func CumulativeIntToRDL(mx *mv.Metric, resourceInfo []*cpb.KeyValue, libraryInfo *cpb.InstrumentationLibrary, rdl chan<- *RDLPoint) error {
	for _, dp := range mx.GetIntSum().GetDataPoints() {
		descriptor := descriptor(mx, GaugeCumulativeInt)
		info := library(libraryInfo)
		descriptor.Temporality = mx.GetIntSum().GetAggregationTemporality().String()
		descriptor.IsMonotonic = mx.GetIntSum().GetIsMonotonic()
		point, _ := metricsNameProcessor.Processor(dataPointToRDL(resourceInfo, info, descriptor, dp.GetLabels(),
			dp.GetValue(), dp.GetTimeUnixNano()))
		rdl <- point
	}
	return nil
}

func CumulativeDoubleToRDL(mx *mv.Metric, resourceInfo []*cpb.KeyValue, libraryInfo *cpb.InstrumentationLibrary, rdl chan<- *RDLPoint) error {
	for _, dp := range mx.GetDoubleSum().GetDataPoints() {
		descriptor := descriptor(mx, GaugeCumulativeDouble)
		info := library(libraryInfo)
		descriptor.Temporality = mx.GetDoubleSum().GetAggregationTemporality().String()
		descriptor.IsMonotonic = mx.GetDoubleSum().GetIsMonotonic()
		point, _ := metricsNameProcessor.Processor(dataPointToRDL(resourceInfo, info, descriptor, dp.GetLabels(),
			int64(dp.GetValue()), dp.GetTimeUnixNano()))
		rdl <- point
	}
	return nil
}
