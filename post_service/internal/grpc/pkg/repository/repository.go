package repository

import (
	"github.com/golang/protobuf/ptypes"
	"gitlab.com/Spide_IT/spide_it/internal/grpc/pkg/api"
	"gitlab.com/Spide_IT/spide_it/internal/post/model"
	"gorm.io/gorm"
)

type GrpcRepository struct {
	db *gorm.DB
}

func NewGrpcRepository(db *gorm.DB)*GrpcRepository{
	return &GrpcRepository{db:db}
}

func (grpc *GrpcRepository) GetData() ([]*api.Post,error){
	var postAll []model.Posts
	grpc.db.Find(&postAll)
	var posts []*api.Post
	for _,post:=range postAll{
		createdAt,_:=ptypes.TimestampProto(post.CreatedAt)
		updatedAt,_:=ptypes.TimestampProto(post.UpdatedAt)
		publishedAt,_:=ptypes.TimestampProto(post.PublishedAt)
		data:=api.Post{
			Id:          int32(post.Id),
			AuthorId:    int32(post.AuthorId),
			Title:       post.Title,
			MetaTitle:   post.MetaTitle,
			Slug:        post.Slug,
			Summary:     post.Summary,
			Published:   int32(post.Published),
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			PublishedAt: publishedAt,
			Content:     post.Content,
		}

		posts=append(posts,&data)

	}
	return posts,nil
}