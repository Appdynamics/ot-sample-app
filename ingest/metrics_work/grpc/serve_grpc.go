package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	mpb "ingest/internal/opentelemetry-proto-gen/collector/metrics/v1"
	tpb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	mv "ingest/internal/opentelemetry-proto-gen/metrics/v1"
	"ingest/servers/grpc/transform"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"text/template"
)


type TraceExportService struct {
	tpb.TraceServiceServer
}

func (a *TraceExportService) Export(_ context.Context, request *tpb.ExportTraceServiceRequest) (*tpb.ExportTraceServiceResponse, error) {
	marshaller := &jsonpb.Marshaler{Indent: "\t"}
	s, err := marshaller.MarshalToString(request)
	if err != nil {
		log.Print(err.Error())
	}
	log.Println(s)
	return &tpb.ExportTraceServiceResponse{}, err
}

type MetricExportService struct {
	mpb.MetricsServiceServer
	messages chan<- *transform.RDLPoint
}

var (
	RdlTemplateStr = "<Metric(name={{.Descriptor.Name}} type={{.Descriptor.Type}} unit={{.Descriptor.Unit}} description={{.Descriptor.Description}} temporality={{.Descriptor.Temporality}}) " +
		" Value(timestamp={{.TimeStamp}} value={{.Value}})" +
		" Labels({{ range $key, $val := .ResourceLabels}}{{$key}}={{$val}} {{end}}) Source(source={{.Source.Name}}, version={{.Source.Version}})>"
	RdlTemplate    = template.Must(template.New("Template").Parse(RdlTemplateStr))
)

func (b *MetricExportService) Export(ctx context.Context, request *mpb.ExportMetricsServiceRequest) (*mpb.ExportMetricsServiceResponse, error) {

	marshaller := jsonpb.Marshaler{Indent: " "}
	file, err := os.OpenFile("json_ex.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}

	err = marshaller.Marshal(file, request)
	if err != nil {
		fmt.Println(err)
	}

	for _, rm := range request.GetResourceMetrics() {
		resource := rm.GetResource().GetAttributes()
		for _, im := range rm.GetInstrumentationLibraryMetrics() {
			for _, mx := range im.GetMetrics() {
				switch mx.GetData().(type) {
				case *mv.Metric_DoubleGauge:
					if err := transform.GaugeDoubleToRDL(mx, resource, im.GetInstrumentationLibrary(), b.messages); err != nil {
						log.Println(err)
					}

				case *mv.Metric_IntGauge:
					if err := transform.GaugeIntToRDL(mx, resource, im.GetInstrumentationLibrary(), b.messages); err != nil {
						log.Println(err)
					}

				case *mv.Metric_DoubleSum:
					if err := transform.CumulativeDoubleToRDL(mx, resource, im.GetInstrumentationLibrary(), b.messages); err != nil {
						log.Println(err)
					}

				case *mv.Metric_IntSum:
					if err := transform.CumulativeIntToRDL(mx, resource, im.GetInstrumentationLibrary(), b.messages); err != nil {
						log.Println(err)
					}
				}
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

	messages := make(chan *transform.RDLPoint, 50)


	go func() {
		log.Printf("Starting main routine\n")
		var wg sync.WaitGroup
		process := func (messages chan *transform.RDLPoint, w io.Writer, wg *sync.WaitGroup) {
			for msg := range messages {
				err := RdlTemplate.Execute(os.Stdout, msg)
				fmt.Print("\n\n")
				if err != nil {
					fmt.Println(err)
				}
			}
			wg.Done()
		}
		for i := 0; i < 1; i++ {
			log.Printf("Starting routine %d\n", i)
			wg.Add(1)
			go process(messages, os.Stdout, &wg)
		}
		wg.Wait()
	}()

	grpcServer := grpc.NewServer(opts...)
	tpb.RegisterTraceServiceServer(grpcServer, &TraceExportService{})
	mpb.RegisterMetricsServiceServer(grpcServer, &MetricExportService{messages: messages})
	panic(grpcServer.Serve(lis))
}
