// main.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "simple-grpc/proto/service/v1"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	aa := &pb.Server{ServerName: fmt.Sprintf("%d", rand.Int())}
	srv := grpc.NewServer()
	pb.RegisterServiceServer(srv, aa)

	fmt.Println("gRPC server is listening on port 50051...")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
