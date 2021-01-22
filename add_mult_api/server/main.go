package main

import (
	"context"
	"fmt"
	"net"

	proto "github.com/kartik-dutt/Learning-Go/proto"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum1(), request.GetNum2()
	fmt.Println(a, b)
	res := a + b

	return &proto.Response{Ans: res}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum1(), request.GetNum2()
	res := a * b
	return &proto.Response{Ans: res}, nil
}

func main() {
	listner, _ := net.Listen("tcp", ":4040")
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(listner); e != nil {
		panic(e)
	}
}
