package post

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/post/model"
)

type PostUsecase interface {
	CreatePosts(posts model.Posts) (string,error)
	UpdatePosts(id int,posts model.Posts) (string,error)
	DeletePosts(id int)(string,error)
	PostsAll()([]model.Posts,error)
	PostsById(id int)(model.Posts,error)
	CreateProfile(profile model.Profiles) (string,error)
	UpdateProfile(id int,profile model.Profiles) (string,error)
	ProfileAll() ([]model.Profiles,error)
	ProfileById(id int)(model.Profiles,error)
	CreatePostMetas(metas model.PostMetas) (string,error)
	UpdatePostMetas(id int,metas model.PostMetas) (string,error)
	DeletePostMetas(id int)(string,error)
	PostMetasAll()([]model.PostMetas,error)
	PostMetasById(id int)(model.PostMetas,error)
	CreateComment(comment model.PostComments) (string,error)
	CommentAllPost(idPost int)([]model.PostComments,error)
	CreateCategories(categories model.Categories) (string,error)
	UpdateCategories(id int,categories model.Categories) (string,error)
	DeleteCategories(id int) (string,error)
	CategoriesAll()([]model.Categories,error)
	CategoriesById(id int)(model.Categories,error)
	ListProfilePost(id int)(model.ProfilePost,error)
	ListPostCategory() ([]model.PostCategory,error)
	PublishPost(id int) (string,error)
	RemoveComment(id int) (string,error)
	ListPublishPost() ([]model.Posts,error)
	ListNotPublishPost() ([]model.Posts,error)

	RequestIpComputer(ip ip.Ip) error
}