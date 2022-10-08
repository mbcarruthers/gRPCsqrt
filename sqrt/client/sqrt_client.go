package main

import (
	"context"
	pb "github.com/mbcarruthers/gRPCsqrt/sqrt/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func sendValue(c pb.SqrtServiceClient, n int32) {
	req := &pb.SqrtRequest{
		Number: n,
	}
	if res, err := c.Sqrt(context.Background(), req); err != nil {
		e, ok := status.FromError(err) // var ok will relay if err is
		if ok {                        // a grpc error
			log.Printf("grpc-error:[%s]::\n%s\n",
				e.Code(), e.Message())
			if e.Code() == codes.InvalidArgument {
				log.Println("Sent a negative number")
			}
		} else {
			log.Fatalf("non-grpc error %s\n",
				err.Error())
		}
	} else {
		log.Printf("Sqrt %f\n", res.Result)
	}
}
