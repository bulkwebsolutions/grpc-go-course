package main

import (
	
	"log"
	"fmt"

	"grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)



func main() {
	fmt.Println("Hello I'm client fucker")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println("Created client: %f", c)

	
}