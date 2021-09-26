package post

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/post/model"
)

type PostsRepository interface {
	CreatePosts(posts model.Posts) (string,error)
	UpdatePosts(id int,posts model.Posts) (string,error)
	DeletePosts(id int)(string,error)

	PostsById(id int)(model.Posts,error)
	CreateProfile(profile model.Profiles) (string,error)
	UpdateProfile(id int,profile model.Profiles) (string,error)
	ProfileById(id int)(model.Profiles,error)
	CreatePostMetas(metas model.PostMetas) (string,error)
	UpdatePostMetas(id int,metas model.PostMetas) (string,error)
	DeletePostMetas(id int)(string,error)
	PostMetasById(id int)(model.PostMetas,error)
	CreateComment(comment model.PostComments) (string,error)
	CommentAllPost(idPost int)([]model.PostComments,error)
	CreateCategories(categories model.Categories) (string,error)
	UpdateCategories(id int,categories model.Categories) (string,error)
	DeleteCategories(id int) (string,error)
	CategoriesAll()([]model.Categories,error)
	CategoriesById(id int)(model.Categories,error)

	PostByAuthorID(authorId  int) ([]model.Posts,error)
	CategoryByPostId(parentId int) (model.Categories,error)

	PostsAll()([]model.Posts,error)
	ProfileAll() ([]model.Profiles,error)
	PostMetasAll()([]model.PostMetas,error)

	//PostsAllPagination(pagination pagination.Pagination)(*pagination.Pagination, error)
	//ProfileAllPagination(pagination pagination.Pagination) (*pagination.Pagination, error)
	//PostMetasAllPagination(pagination pagination.Pagination)(*pagination.Pagination, error)

	PublishPost(id int) (string,error)
	RemoveComment(id int) (string,error)

	ListPublishPost() ([]model.Posts,error)
	ListNotPublishPost() ([]model.Posts,error)

	RequestIpComputer(ip ip.Ip) error
}