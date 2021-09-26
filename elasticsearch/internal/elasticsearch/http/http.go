package http

import (
	"context"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/elasticsearch/usecase"
)

type Http struct {
	usecase *usecase.ElasticsearchUsecase
}

func NewHttp(usecase *usecase.ElasticsearchUsecase) *Http {
	return &Http{usecase:usecase}
}

func (h *Http) Query(ctx context.Context,name string) (interface{},error){
	return h.usecase.Query(ctx,name)
}