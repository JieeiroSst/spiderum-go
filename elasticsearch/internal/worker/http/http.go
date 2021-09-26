package http

import (
	"context"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/interfaceWorker"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/proto"
)

type Http struct {
	usecase interface_worker.ElasticsearchUsecase
}

func NewHttp(usecase interface_worker.ElasticsearchUsecase)*Http{
	return &Http{usecase:usecase}
}

func (h *Http) InsertPost(ctx context.Context,data proto.Post) error{
	return h.usecase.Insert(ctx,data)
}