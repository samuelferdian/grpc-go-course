package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "test1"},
		{FirstName: "test2"},
		{FirstName: "test3"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("error while receiving : %v\n", err)
				break
			}

			log.Printf("received: %v\n", res)
		}
		close(waitc)
	}()

	<-waitc
}
