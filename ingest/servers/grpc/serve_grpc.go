package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	mpb "ingest/internal/opentelemetry-proto-gen/collector/metrics/v1"
	tpb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	"os"

	"log"
	"net"
)

var (
	rdb *redis.Client
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

func (b *MetricExportService) Export(ctx context.Context, request *mpb.ExportMetricsServiceRequest) (*mpb.ExportMetricsServiceResponse, error) {
	marshaller := &jsonpb.Marshaler{Indent: "\t"}
	s, err := marshaller.MarshalToString(request)
	if err != nil {
		log.Print(err.Error())
	}

	err = rdb.Publish(ctx, os.Getenv("REDIS_METRICS_CHANNEL"), s).Err()
	if err != nil {
		log.Print(err.Error())
	}
	log.Print("success!")



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
