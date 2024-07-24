package dltlsdml_com

import (
	"context"
	"dltlsdml.com/grpc/dltlsdml.com/dltlsdml"
	"google.golang.org/grpc"
	"log"
	"time"
)

func TestGrpcClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := dltlsdml.NewDltlsdmlClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &dltlsdml.HelloRequest{Name: "이신의", Address: "집"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s, Address: %s", r.GetMessage(), r.GetAddress())
}
