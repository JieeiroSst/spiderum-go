package elasticsearch

import "context"

type ElasticsearchUsecase interface {
	Query(ctx context.Context,name string) (interface{},error)
}