package main

import (
	"context"
	"go_grpc_demo/proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const Port = "172.17.71.143:9000"

func main() {
	fromFile, err := credentials.NewClientTLSFromFile("../config/server.pem", "www.eline.com")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(Port, grpc.WithTransportCredentials(fromFile))
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
