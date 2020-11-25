package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	tpb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	"log"
	"net"
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
	panic(grpcServer.Serve(lis))
}
