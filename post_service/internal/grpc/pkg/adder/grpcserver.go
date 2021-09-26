package adder

import (
	"context"

	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/api"
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/delivery"
)

type GRPCServer struct{
	http *delivery.Grpc
}

func (s *GRPCServer) NewGRPCServer(http *delivery.Grpc){
	s.http = http
}

func (s *GRPCServer) UpdatePost(ctx context.Context,req *api.RequestPost)(*api.ResponsePost,error){
	posts,err:=s.http.GetData()
	if err!=nil{
		return &api.ResponsePost{}, err
	}
	return &api.ResponsePost{Posts:posts}, nil
}