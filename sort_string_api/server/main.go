package main

import (
	"context"
	"net"
	"sort"
	"strings"

	service "github.com/kartik-dutt/Simple-and-Fun-GRPC-API/service"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Sort(ctx context.Context, req *service.Request) (*service.Request, error) {
	inp := req.GetInp()
	str := strings.Split(inp, "")
	sort.Strings(str)
	return &service.Request{Inp: strings.Join(str, "")}, nil
}

func main() {
	listner, _ := net.Listen("tcp", ":1040")
	defer listner.Close()
	srv := grpc.NewServer()
	service.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(listner); e != nil {
		panic(e)
	}
}
