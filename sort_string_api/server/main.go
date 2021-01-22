package main

import (
	"context"
	"net"
	"sort"

	service "github.com/kartik-dutt/Simple-and-Fun-GRPC-API/service"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Sort(ctx *context.Context, req *service.Request) (*service.Request, error) {
	inp := req.GetInp()
	sort.Slice(inp, func(i, j int) bool {
		return inp[i] < inp[j]
	})

	return &service.Request{Inp: inp}, nil
}

func main() {
	listner, _ := net.Listen("tcp", ":4040")
	srv := grpc.NewServer()
	service.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(listner); e != nil {
		panic(e)
	}
}
