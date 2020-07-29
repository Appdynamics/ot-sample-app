package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Appdynamics/opentelemetry-ingest/gen/go/pb/appdynamics/v1"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ExportService struct {
	pb.SpanHandlerServer
}

func (a *ExportService) HandleSpans(ctx context.Context, request *pb.SpansRequest) (*pb.SpansResponse, error) {
	marshaller := &jsonpb.Marshaler{Indent: "\t"}
	s, err := marshaller.MarshalToString(request)
	if err == nil {
		log.Println(s)
	} else {
		log.Print(err.Error())
	}
	return &pb.SpansResponse{}, err
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
	pb.RegisterSpanHandlerServer(grpcServer, &ExportService{})
	panic(grpcServer.Serve(lis))
}