package main

import (
	"context"
	"io"
	"log"

	pb "grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes function was invoked")

	req := &pb.PrimeRequest{
		Number: 1000,
	}

	streams, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := streams.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}
}
