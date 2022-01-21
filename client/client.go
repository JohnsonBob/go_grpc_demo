package main

import (
	"context"
	"go_grpc_demo/proto/proto"
	"google.golang.org/grpc"
	"log"
)

const Port = "ddns.lingdian.site:9000"

func main() {
	conn, err := grpc.Dial(Port, grpc.WithInsecure())
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
