package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	mpb "ingest/internal/opentelemetry-proto-gen/collector/metrics/v1"
	tpb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	"log"
	"net"
)

type TraceExportService struct {
	tpb.TraceServiceServer
}

func (a *TraceExportService) Export(ctx context.Context, request *tpb.ExportTraceServiceRequest) (*tpb.ExportTraceServiceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Print("No meta data found")
	} else {
		token, ok := md["access-key"]
		if ok {
			log.Println("Token Info:")
			log.Println(token)
		}
	}

	marshaller := &jsonpb.Marshaler{Indent: "\t"}
	s, err := marshaller.MarshalToString(request)
	if err == nil {
		log.Println()
		log.Println(s)
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
		log.Println(s)
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
