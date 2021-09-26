package repository

import (
	"context"
	"github.com/olivere/elastic/v7"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/proto"
	"reflect"
)

type Elasticsearch struct {
	elastic *elastic.Client
}

func NewElasticsearchInterface(elastic *elastic.Client) *Elasticsearch{
	return &Elasticsearch{elastic:elastic}
}

func (e *Elasticsearch) Query(ctx context.Context,name string) (interface{},error){
	query := elastic.NewMatchBoolPrefixQuery("title", name)
	result, err := e.elastic.Search().
		Index("posts").
		Type("post").
		Query(query).
		Pretty(true).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	var posts []proto.Post
	var post proto.Post
	for _, item := range result.Each(reflect.TypeOf(post)) {
		t := item.(proto.Post)
		post = proto.Post{
			Id:          t.Id,
			AuthorId:    t.AuthorId,
			Title:       t.Title,
			MetaTitle:   t.MetaTitle,
			Slug:        t.Slug,
			Summary:     t.Summary,
			Published:   t.Published,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			PublishedAt: t.PublishedAt,
			Content:     t.Content,
		}
		posts= append(posts,post)
	}
	return posts,nil
}