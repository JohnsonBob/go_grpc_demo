package main

import (
	"go_grpc_demo/proto/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

const Port = ":9000"

type SearchService struct {
}

func (s SearchService) Search(c context.Context, request *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{Response: request.GetRequest() + " Server"}, nil
}

func main() {
	tlsFromFile, err := credentials.NewServerTLSFromFile("../config/server.pem", "../config/server.key")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(grpc.Creds(tlsFromFile))
	proto.RegisterSearchServiceServer(server, SearchService{})
	listen, err := net.Listen("tcp", Port)
	if err != nil {
		panic(err)
	}
	_ = server.Serve(listen)
}
