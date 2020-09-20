package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	mpb "ingest/internal/opentelemetry-proto-gen/collector/metrics/v1"
	tpb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	"io"
	"log"
	"net"
	"os"
)

type TraceExportService struct {
	tpb.TraceServiceServer
}



func (a *TraceExportService) Export(_ context.Context, request *tpb.ExportTraceServiceRequest) (*tpb.ExportTraceServiceResponse, error) {
	marshaller := &jsonpb.Marshaler{Indent: "\t"}
	s, err := marshaller.MarshalToString(request)
	if err == nil {
		if _, err := io.WriteString(os.Stdout, s); err != nil {
			fmt.Print(err)
		}
	} else {
		log.Print(err.Error())
	}
	return &tpb.ExportTraceServiceResponse{}, err
}

type MetricExportService struct {
	mpb.MetricsServiceServer
}

func (b *MetricExportService) Export(_ context.Context, request *mpb.ExportMetricsServiceRequest) (*mpb.ExportMetricsServiceResponse, error) {
	marshaller := &jsonpb.Marshaler{Indent: "\t"}
	s, err := marshaller.MarshalToString(request)
	if err == nil {
		if _, err := io.WriteString(os.Stdout, s); err != nil {
			fmt.Print(err)
		}

	} else {
		log.Print(err.Error())
	}
	return &mpb.ExportMetricsServiceResponse{}, err
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
