package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	pb "app/dict"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalcServiceServer
}

func (s *server) Add(ctx context.Context, nums *pb.TwoNumber) (*pb.ResultTwoNumber, error) {
	return &pb.ResultTwoNumber{
		Result: nums.A + nums.B,
	}, nil
}
func (s *server) Sub(ctx context.Context, nums *pb.TwoNumber) (*pb.ResultTwoNumber, error) {
	return &pb.ResultTwoNumber{
		Result: nums.A - nums.B,
	}, nil
}
func (s *server) Mult(ctx context.Context, nums *pb.TwoNumber) (*pb.ResultTwoNumber, error) {
	return &pb.ResultTwoNumber{
		Result: nums.A * nums.B,
	}, nil
}

func (s *server) Div(ctx context.Context, nums *pb.TwoNumber) (*pb.ResultTwoNumber, error) {
	return &pb.ResultTwoNumber{
		Result: nums.A / nums.B,
	}, nil
}
func (s *server) Sqrt(ctx context.Context, nums *pb.SqrtNum) (*pb.ResultTwoNumber, error) {
	result := math.Sqrt(float64(nums.Num))
	return &pb.ResultTwoNumber{
		Result: int32(result),
	}, nil
}

func (s *server) Pow(ctx context.Context, nums *pb.TwoNumber) (*pb.ResultTwoNumber, error) {
	return &pb.ResultTwoNumber{
		Result: int32(math.Pow(float64(nums.A), float64(nums.B))),
	}, nil
}
func (s *server) Min(ctx context.Context, nums *pb.TwoNumber) (*pb.ResultTwoNumber, error) {
	return &pb.ResultTwoNumber{
		Result: int32(math.Min(float64(nums.A), float64(nums.B))),
	}, nil
}
func main() {

	lis, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &server{})

	fmt.Println("Listen RPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
