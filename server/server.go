package main

import (
	"go_grpc_demo/proto/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

const Port = ":9000"

type SearchService struct {
}

func (s SearchService) Search(c context.Context, request *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{Response: request.GetRequest() + " Server"}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterSearchServiceServer(server, SearchService{})
	listen, err := net.Listen("tcp", Port)
	if err != nil {
		panic(err)
	}
	_ = server.Serve(listen)
}
