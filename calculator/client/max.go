package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{1, 2, 5, 3, 23, 11, 80}

		for _, number := range numbers {
			log.Printf("sending number: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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
				log.Printf("error while reading server stream: %v", err)
				break
			}

			log.Printf("received a new maximum: %d\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
