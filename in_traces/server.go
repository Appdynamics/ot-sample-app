package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	pb "github.com/pavankrish123/ot-svc/gen/go/pb/appdynamics/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

type ExportService struct {
	pb.AppDServiceServer
}

func (a *ExportService) Export(ctx context.Context, request *pb.AppDExportTraceServiceRequest) (*pb.AppDExportTraceServiceResponse, error) {
	marshaller := &jsonpb.Marshaler{}
	var k strings.Builder
	err := marshaller.Marshal(&k, request)
	log.Printf(k.String())
	return &pb.AppDExportTraceServiceResponse{}, err
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
	pb.RegisterAppDServiceServer(grpcServer, &ExportService{})
	panic(grpcServer.Serve(lis))
}