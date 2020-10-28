package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	mpb "ingest/internal/opentelemetry-proto-gen/collector/metrics/v1"
	tpb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	cpb "ingest/internal/opentelemetry-proto-gen/common/v1"
	mv "ingest/internal/opentelemetry-proto-gen/metrics/v1"
	"log"
	"net"
	"os"
	"text/template"
)

var (
	rdb *redis.Client
)

const (
	GaugeInt64                  = "GAUGE_INT64"
	GaugeDouble                 = "GAUGE_DOUBLE"
	GaugeDistribution           = "GAUGE_DISTRIBUTION"
	GaugeCumulativeInt          = "CUMULATIVE_INT64"
	GaugeCumulativeDouble       = "CUMULATIVE_DOUBLE"
	GaugeCumulativeDistribution = "CUMULATIVE_DISTRIBUTION"
)

var (
	RdlTemplateStr = `<Metric(name={{.Descriptor.Name}} type={{.Descriptor.Type}} unit={{.Descriptor.Unit}} description={{.Descriptor.Description}}) Value(timestamp={{.TimeStamp}} value={{.Value}}) Labels({{ range $key, $val := .Labels}}{{$key}}={{$val}} {{end}}) Source(source={{.Source.Name}}, version={{.Source.Version}})>`
	RdlTemplate = template.Must(template.New("Template").Parse(RdlTemplateStr))
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

type TraceExportService struct {
	tpb.TraceServiceServer
}

func (a *TraceExportService) Export(ctx context.Context, request *tpb.ExportTraceServiceRequest) (*tpb.ExportTraceServiceResponse, error) {
	marshaller := &jsonpb.Marshaler{Indent: "\t"}
	s, err := marshaller.MarshalToString(request)
	if err != nil {
		log.Print(err.Error())
	}

	err = rdb.Publish(ctx, os.Getenv("REDIS_TRACES_CHANNEL"), s).Err()
	if err != nil {
		log.Print(err.Error())
	}

	return &tpb.ExportTraceServiceResponse{}, err
}

type MetricExportService struct {
	mpb.MetricsServiceServer
}

func stringify(val *cpb.AnyValue) string {
	switch val.Value.(type) {
	case *cpb.AnyValue_BoolValue:
		return fmt.Sprintf("%", val.GetBoolValue())
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

type RDL []*RDLPoint

type SourceInfo struct {
	Name    string
	Version string
}

type DescriptorInfo struct {
	Name        string
	Description string
	Unit        string
	Type        string
}

type RDLPoint struct {
	Source     *SourceInfo
	Labels     map[string]string
	Descriptor *DescriptorInfo
	TimeStamp  uint64
	Value      float64
}

func (b *MetricExportService) Export(ctx context.Context, request *mpb.ExportMetricsServiceRequest) (*mpb.ExportMetricsServiceResponse, error) {
	out := make(RDL, 0)
	labels := map[string]string{}
	source := &SourceInfo{
		"OpenTelemetry Collector",
		"v0.12",
	}
	for _, rm := range request.GetResourceMetrics() {
		_ = proto.MarshalText(os.Stdout, rm)
		fmt.Println("\n")
		resource := rm.GetResource()
		for _, kv := range resource.GetAttributes() {
			labels[kv.GetKey()] = stringify(kv.GetValue())
		}
		for _, im := range rm.GetInstrumentationLibraryMetrics() {
			for _, mx := range im.GetMetrics() {
				md := &DescriptorInfo{mx.GetName(), mx.GetDescription(), mx.GetUnit(), ""}
				rdp := &RDLPoint{}
				switch mx.GetData().(type) {
				case *mv.Metric_DoubleGauge:
					dps := mx.GetDoubleGauge()
					md.Type = GaugeDouble
					for _, dp := range dps.GetDataPoints() {
						for _, kv := range dp.Labels {
							labels[kv.GetKey()] = kv.GetValue()
						}
						rdp.Labels = labels
						rdp.Value = dp.GetValue()
						rdp.TimeStamp = dp.GetTimeUnixNano()
						rdp.Source = source
						rdp.Descriptor = md
					}
				case *mv.Metric_IntGauge:
					dps := mx.GetIntGauge()
					md.Type = GaugeInt64
					for _, dp := range dps.GetDataPoints() {
						for _, kv := range dp.Labels {
							labels[kv.GetKey()] = kv.GetValue()
						}
						rdp.Labels = labels
						rdp.Value = float64(dp.GetValue())
						rdp.TimeStamp = dp.GetTimeUnixNano()
						rdp.Source = source
						rdp.Descriptor = md
					}
				case *mv.Metric_DoubleSum:
					dps := mx.GetDoubleSum()
					md.Type = GaugeCumulativeDouble
					for _, dp := range dps.GetDataPoints() {
						for _, kv := range dp.Labels {
							labels[kv.GetKey()] = kv.GetValue()
						}
						rdp.Labels = labels
						rdp.Value = dp.GetValue()
						rdp.TimeStamp = dp.GetTimeUnixNano()
						rdp.Source = source
						rdp.Descriptor = md
					}
				case *mv.Metric_IntSum:
					dps := mx.GetIntSum()
					md.Type = GaugeCumulativeDouble
					for _, dp := range dps.GetDataPoints() {
						for _, kv := range dp.Labels {
							labels[kv.GetKey()] = kv.GetValue()
						}
						rdp.Labels = labels
						rdp.Value = float64(dp.GetValue())
						rdp.TimeStamp = dp.GetTimeUnixNano()
						rdp.Source = source
						rdp.Descriptor = md
					}
				}



				err := RdlTemplate.Execute(os.Stdout, rdp)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("\n")
				out = append(out, rdp)
			}
		}
	}
	return &mpb.ExportMetricsServiceResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 5000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server Listening on :%v", 5000)
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	tpb.RegisterTraceServiceServer(grpcServer, &TraceExportService{})
	mpb.RegisterMetricsServiceServer(grpcServer, &MetricExportService{})
	panic(grpcServer.Serve(lis))
}
