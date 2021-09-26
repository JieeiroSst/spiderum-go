package elasticsearch

import (
	"context"
	"github.com/olivere/elastic/v7"
	"sync"
)

var (
	instance Elasticsearch
	once sync.Once
)

type Elasticsearch struct {
	elastic *elastic.Client
}

type ElasticsearchInterface interface {
	Insert(ctx context.Context,data interface{}) error
	Query(ctx context.Context,name string) (interface{},error)
}

func GetElasticsearchConnInstance(dns string) *Elasticsearch{
	once.Do(func() {
		client, err := elastic.NewClient(elastic.SetURL(dns))
		if err != nil {
			panic(err)
		}
		instance=Elasticsearch{elastic:client}
	})
	return &instance
}

func NewGetElasticsearchConn(dns string) *elastic.Client {
	return GetElasticsearchConnInstance(dns).elastic
}
