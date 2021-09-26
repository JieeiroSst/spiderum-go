package interface_worker

import (
	"context"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/proto"
)

type ElasticsearchRepository interface {
	Insert(ctx context.Context,data proto.Post) error
}