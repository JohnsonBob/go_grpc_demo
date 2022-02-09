package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go_grpc_demo/proto/proto"
	"go_grpc_demo/server/interceptor"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

const Port = ":9000"

type SearchService struct {
}

func (s SearchService) Search(c context.Context, request *proto.SearchRequest) (*proto.SearchResponse, error) {
	var a = 10
	var b = 10
	b = b * 0
	a = a / (8 * b)
	return &proto.SearchResponse{Response: request.GetRequest() + " Server"}, nil
}

func main() {
	tlsFromFile, err := credentials.NewServerTLSFromFile("../config/server.pem", "../config/server.key")
	if err != nil {
		panic(err)
	}
	options := []grpc.ServerOption{
		grpc.Creds(tlsFromFile),
		grpc_middleware.WithUnaryServerChain(
			interceptor.RecoveryInterceptor,
			interceptor.LoggingInterceptor,
		),
	}

	server := grpc.NewServer(options...)
	proto.RegisterSearchServiceServer(server, SearchService{})
	listen, err := net.Listen("tcp", Port)
	if err != nil {
		panic(err)
	}
	_ = server.Serve(listen)
}
