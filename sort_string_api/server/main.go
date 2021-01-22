package main

import (
	"context"
	"sort"
)

type server struct{}

func (s *server) Sort(ctx *context.Context, req *service.Request) (*service.Request, error) {
	inp := req.GetInp()
	sort.Slice(inp, func(i, j int) bool {
		return inp[i] < inp[j]
	})

	return &service.Request{Inp: inp}, nil
}
