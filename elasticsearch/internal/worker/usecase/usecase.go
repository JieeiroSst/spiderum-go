package usecase

import (
	"context"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/interfaceWorker"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/proto"
)

type ElasticsearchUsecase struct {
	repo interface_worker.ElasticsearchRepository
}

func NewElasticsearchUsecase(repo interface_worker.ElasticsearchRepository)*ElasticsearchUsecase{
	return &ElasticsearchUsecase{repo:repo}
}

func (e *ElasticsearchUsecase) Insert(ctx context.Context,data proto.Post) error{
	return e.repo.Insert(ctx,data)
}