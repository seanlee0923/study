package dltlsdml_com

import (
	"dltlsdml.com/grpc/dltlsdml.com/dltlsdml"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	dltlsdml.UnimplementedDltlsdmlServer
}

func (s *server) SayHello(ctx context.Context, req *dltlsdml.HelloRequest) (*dltlsdml.HelloReply, error) {
	return &dltlsdml.HelloReply{Message: "Hello " + req.Name, Address: req.Address}, nil
}

func TestGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	dltlsdml.RegisterDltlsdmlServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
