package main

import (
	pb "github.com/mbcarruthers/gRPCsqrt/sqrt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	addr = "localhost:5051"
)

func main() {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dialing %s\n %s",
			addr, err.Error())
	}
	defer conn.Close()

	clientConn := pb.NewSqrtServiceClient(conn)

	sendValue(clientConn, 0)

}
