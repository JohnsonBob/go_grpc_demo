package interceptor

import (
	"go_grpc_demo/server/interceptor/bean"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	server := info.Server
	switch server.(type) {
	case bean.SearchService:
		var searchService = server.(bean.SearchService)
		if err := searchService.Auth.Check(ctx); err != nil {
			return resp, err
		} else {
			return handler(ctx, req)
		}
	default:
		return handler(ctx, req)
	}
}
