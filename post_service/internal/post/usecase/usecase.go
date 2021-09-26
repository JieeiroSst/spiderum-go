package usecase

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/post"
	"gitlab.com/Spide_IT/spide_it/internal/post/model"
)

type PostsUsecase struct {
	repository post.PostsRepository
}


func NewPostsUsecase(repository post.PostsRepository) *PostsUsecase {
	return &PostsUsecase{repository:repository}
}

func (post *PostsUsecase) PublishPost(id int) (string, error) {
	msg,err:=post.repository.PublishPost(id)
	return msg,err
}

func (post *PostsUsecase) CreatePosts(posts model.Posts) (string,error){
	msg,err:=post.repository.CreatePosts(posts)
	return msg,err
}

func (post *PostsUsecase) UpdatePosts(id int,posts model.Posts) (string,error)  {
	msg,err:= post.repository.UpdatePosts(id,posts)
	return msg,err
}

func (post *PostsUsecase) DeletePosts(id int)(string,error){
	msg,err:=post.repository.DeletePosts(id)
	return msg,err
}

func (post *PostsUsecase) PostsAll()([]model.Posts,error){
	posts,err:=post.repository.PostsAll()
	return posts,err
}

func (post * PostsUsecase) PostsById(id int)(model.Posts,error){
	posts,err:=post.repository.PostsById(id)
	return posts,err
}

func (post *PostsUsecase) CreateProfile(profile model.Profiles) (string,error){
	msg,err:=post.repository.CreateProfile(profile)
	return msg,err
}

func (post *PostsUsecase) UpdateProfile(id int,profile model.Profiles) (string,error){
	msg,err:=post.repository.UpdateProfile(id,profile)
	return msg,err
}

func (post *PostsUsecase) ProfileAll() ([]model.Profiles,error){
	profile,err:=post.repository.ProfileAll()
	return profile,err
}

func (post *PostsUsecase) ProfileById(id int)(model.Profiles,error){
	profile,err:=post.repository.ProfileById(id)
	return profile,err
}

func (post *PostsUsecase) CreatePostMetas(metas model.PostMetas) (string,error){
	msg,err:=post.repository.CreatePostMetas(metas)
	return msg,err
}

func (post *PostsUsecase) UpdatePostMetas(id int,metas model.PostMetas) (string,error){
	msg,err:=post.repository.UpdatePostMetas(id,metas)
	return msg,err
}

func (post *PostsUsecase) DeletePostMetas(id int)(string,error){
	msg,err:=post.repository.DeletePostMetas(id)
	return msg,err
}

func (post *PostsUsecase)PostMetasAll()([]model.PostMetas,error){
	metas,err:=post.repository.PostMetasAll()
	return metas,err
}

func (post *PostsUsecase) PostMetasById(id int)(model.PostMetas,error){
	meta,err:=post.repository.PostMetasById(id)
	return meta,err
}

func (post *PostsUsecase) CreateComment(comment model.PostComments) (string,error){
	msg,err:=post.repository.CreateComment(comment)
	return msg,err
}

func (post *PostsUsecase) CommentAllPost(idPost int)([]model.PostComments,error){
	comment,err:=post.repository.CommentAllPost(idPost)
	return comment,err
}

func (post *PostsUsecase) CreateCategories(categories model.Categories) (string,error){
	msg,err:=post.repository.CreateCategories(categories)
	return msg,err
}

func (post *PostsUsecase) UpdateCategories(id int,categories model.Categories) (string,error){
	msg,err:=post.repository.UpdateCategories(id,categories)
	return msg,err
}

func (post *PostsUsecase) DeleteCategories(id int) (string,error) {
	msg,err:=post.repository.DeleteCategories(id)
	return msg,err
}

func (post *PostsUsecase) CategoriesAll()([]model.Categories,error){
	categories,err:=post.repository.CategoriesAll()
	return categories,err
}

func (post *PostsUsecase) CategoriesById(id int)(model.Categories,error){
	category,err:=post.repository.CategoriesById(id)
	return category,err
}

func (post *PostsUsecase) ListProfilePost(id int)(model.ProfilePost,error) {
	profile,err:=post.repository.ProfileById(id)
	if err!=nil{
		return model.ProfilePost{}, nil
	}
	posts,err:=post.repository.PostByAuthorID(id)
	if err!=nil{
		return model.ProfilePost{}, nil
	}
	return model.ProfilePost{
		Profile:profile,
		Posts: posts,
	}, nil
}

func (post *PostsUsecase) ListPostCategory() ([]model.PostCategory,error){
	var postCategories []model.PostCategory
	posts,err:=post.repository.PostsAll()
	if err!=nil{
		return nil, nil
	}
	for _,p:= range posts {
		cateogory,err:=post.repository.CategoryByPostId(p.Id)
		if err!=nil {
			return nil, nil
		}
		postCategory:=model.PostCategory{
			Post:     p,
			Category: cateogory,
		}
		postCategories=append(postCategories,postCategory)
	}
	return postCategories,nil
}

func (post *PostsUsecase) publishPost(id int) (string,error){
	msg,err:=post.repository.PublishPost(id)

	return msg,err
}

func (post *PostsUsecase) RemoveComment(id int) (string,error){
	msg,err:=post.repository.RemoveComment(id)
	return msg,err
}

func (post *PostsUsecase) ListPublishPost() ([]model.Posts,error){
	list,err:=post.repository.ListPublishPost()
	return list,err
}

func (post *PostsUsecase) ListNotPublishPost() ([]model.Posts,error){
	list,err:=post.repository.ListNotPublishPost()
	return list,err
}

func (post *PostsUsecase) RequestIpComputer(ip ip.Ip) error{
	return post.repository.RequestIpComputer(ip)
}