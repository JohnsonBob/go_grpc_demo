package bean

import (
	"go_grpc_demo/proto/proto"
	"golang.org/x/net/context"
)

type SearchService struct {
	Auth *Auth
}

func (s SearchService) Search(ctx context.Context, request *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{Response: request.GetRequest() + " Server"}, nil
}
