package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	pg "grpc-go-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pg.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "test1"},
		{FirstName: "test2"},
		{FirstName: "test3"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %v\n", res.Result)
}
