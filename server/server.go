package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go_grpc_demo/proto/proto"
	"go_grpc_demo/server/interceptor"
	"go_grpc_demo/server/interceptor/bean"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

const Port = ":9000"

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
			interceptor.AuthInterceptor,
		),
	}

	server := grpc.NewServer(options...)
	proto.RegisterSearchServiceServer(server, bean.SearchService{Auth: &bean.Auth{AppKey: "eddycjy", AppSecret: "20181005"}})
	listen, err := net.Listen("tcp", Port)
	if err != nil {
		panic(err)
	}
	_ = server.Serve(listen)
}
