package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	pb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	"log"
	"net"
)

type ExportService struct {
	pb.TraceServiceServer
}

func (a *ExportService) Export(ctx context.Context, request *pb.ExportTraceServiceRequest) (*pb.ExportTraceServiceResponse, error) {
	marshaller := &jsonpb.Marshaler{Indent: "\t"}
	s, err := marshaller.MarshalToString(request)
	if err == nil {
		log.Println(s)
	} else {
		log.Print(err.Error())
	}
	return &pb.ExportTraceServiceResponse{}, err
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
	pb.RegisterTraceServiceServer(grpcServer, &ExportService{})
	panic(grpcServer.Serve(lis))
}
