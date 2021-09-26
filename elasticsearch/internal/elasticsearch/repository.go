package elasticsearch

import "context"

type ElasticsearchRepository interface {
	Query(ctx context.Context,name string) (interface{},error)
}