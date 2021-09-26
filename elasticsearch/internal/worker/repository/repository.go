package repository

import (
	"context"
	"github.com/matoous/go-nanoid/v2"
	"github.com/olivere/elastic/v7"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/proto"
	"log"
)

type ElasticsearchRepository struct {
	elastic *elastic.Client
}

//date integer  text
var mapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"id":{
				"type":"integer"
			},
			"author_id":{
				"type":"integer"
			},
			"title":{
				"type":"text"
			},
			"meta_title":{
				"type":"text"
			},
			"slug":{
				"type":"text"
			},
			"summary":{
				"type":"text"
			},
			"published":{
				"type":"integer"
			},
			"created_at":{
				"type":"date"
			},
			"updated_at":{
				"type":"date"
			},
			"published_at":{
				"type":"date"
			},
			"content":{
				"type":"text"
			}
		}
	}
}`


func NewElasticsearchInterface(elastic *elastic.Client) *ElasticsearchRepository{
	//ctx := context.Background()
	//createIndex, err := elastic.CreateIndex("posts").BodyString(mapping).Do(ctx)
	//if err != nil {
	//	log.Println(err)
	//}
	//if !createIndex.Acknowledged {
	//	// Not acknowledged
	//}
	return &ElasticsearchRepository{elastic:elastic}
}

func (e *ElasticsearchRepository) Insert(ctx context.Context,data  proto.Post) error{
	id, err := gonanoid.New()
	put1, err := e.elastic.Index().
		Index("posts").
		Type("post").
		Id(id).
		BodyJson(data).
		Do(ctx)
	if err != nil {
		return err
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return nil
}
