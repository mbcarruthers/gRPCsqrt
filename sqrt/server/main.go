package main

import (
	pb "github.com/mbcarruthers/gRPCsqrt/sqrt/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	addr = "0.0.0.0:5051"
	tcp  = "tcp"
)

type Server struct {
	pb.SqrtServiceServer
}

func main() {
	lis, err := net.Listen(tcp, addr)
	if err != nil {
		log.Fatalf("Error listening to %s:%s\n%s",
			tcp, addr,
			err.Error())
	}
	log.Printf("[Server] now listening at %s",
		addr)

	s := grpc.NewServer()
	pb.RegisterSqrtServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error listening at grpc server\n%s\n",
			err.Error())
	}
}
