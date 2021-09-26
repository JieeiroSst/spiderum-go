package usecase

import (
	"context"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/elasticsearch"
)

type ElasticsearchUsecase struct {
	repo elasticsearch.ElasticsearchRepository
}

func NewElasticsearchUsecase(repo elasticsearch.ElasticsearchRepository) *ElasticsearchUsecase{
	return &ElasticsearchUsecase{repo:repo}
}

func (e *ElasticsearchUsecase) Query(ctx context.Context,name string) (interface{},error){
	return e.repo.Query(ctx,name)
}