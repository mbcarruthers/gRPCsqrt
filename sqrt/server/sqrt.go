package main

import (
	"context"
	"fmt"
	pb "github.com/mbcarruthers/gRPCsqrt/sqrt/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Server) Sqrt(ctx context.Context,
	req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("sqrt(server) was invoked with %v\n", req)

	if req.Number <= 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			fmt.Sprintf("Received negative number %d", req.Number))
	} else {
		val := float64(req.Number)
		sqrt := math.Sqrt(val)
		return &pb.SqrtResponse{
			Result: sqrt,
		}, nil
	}
}
