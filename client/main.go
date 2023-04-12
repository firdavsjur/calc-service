package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "app/dict"
)

func main() {

	conn, err := grpc.Dial("localhost:9001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("error to connect: %+v", err)
	}

	c := pb.NewCalcServiceClient(conn)
	var nums = &pb.TwoNumber{
		A: 2,
		B: 3,
	}
	resp1, err := c.Add(context.Background(), nums)
	if err != nil {
		log.Fatal(err)
	}
	resp2, err := c.Sub(context.Background(), nums)
	if err != nil {
		log.Fatal(err)
	}
	resp3, err := c.Mult(context.Background(), nums)
	if err != nil {
		log.Fatal(err)
	}
	resp4, err := c.Div(context.Background(), nums)
	if err != nil {
		log.Fatal(err)
	}
	resp5, err := c.Pow(context.Background(), nums)
	if err != nil {
		log.Fatal(err)
	}
	resp6, err := c.Min(context.Background(), nums)
	if err != nil {
		log.Fatal(err)
	}
	resp7, err := c.Sqrt(context.Background(), &pb.SqrtNum{
		Num: 9,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("A,B = ", nums.A, nums.B)
	fmt.Printf("Add: %v\nSub: %v\nMult: %v\nDiv: %b\nPow: %v\nMin: %v\nSqrt: %v\n", resp1.Result, resp2.Result, resp3.Result, resp4.Result, resp5.Result, resp6.Result, resp7.Result)

}
