package main

import (
	"context"
	"go_grpc_demo/proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const Port = "127.0.0.1:9000"

type Auth struct {
	AppKey    string
	AppSecret string
}

func main() {
	fromFile, err := credentials.NewClientTLSFromFile("../config/server.pem", "www.eline.com")
	if err != nil {
		panic(err)
	}
	auth := Auth{
		AppKey:    "eddycjy",
		AppSecret: "20181005",
	}
	conn, err := grpc.Dial(Port, grpc.WithTransportCredentials(fromFile), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()
	client := proto.NewSearchServiceClient(conn)
	search, err := client.Search(context.Background(), &proto.SearchRequest{Request: "gRPC"})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", search.GetResponse())
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return true
}
