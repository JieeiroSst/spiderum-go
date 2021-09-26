package http

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/post"
	"gitlab.com/Spide_IT/spide_it/internal/post/model"
	"gitlab.com/Spide_IT/spide_it/pkg/snowflake"
	"time"
)

type PostHttp struct {
	usecase post.PostUsecase
	snowflake snowflake.Snowflake
}

func NewPostHttp(usecase post.PostUsecase,snowflake snowflake.Snowflake) *PostHttp {
	return &PostHttp{
		usecase:usecase,
		snowflake:snowflake,
	}
}

func (post *PostHttp) CreatePosts(authorId int, title string, metaTitle string, slug string, summary string,content string)(string,error) {
	posts:=model.Posts{
		Id:          post.snowflake.GearedID(),
		AuthorId:    authorId,
		Title:       title,
		MetaTitle:   metaTitle,
		Slug:        slug,
		Summary:     summary,
		Content:     content,
		CreatedAt:   time.Now(),
	}
	msg,err:=post.usecase.CreatePosts(posts)
	return msg,err
}

func (post *PostHttp) UpdatePosts(id int, title string, metaTitle string, slug string, summary string,content string) (string,error){
	posts:=model.Posts{
		Title:       title,
		MetaTitle:   metaTitle,
		Slug:        slug,
		Summary:     summary,
		Content:     content,
		UpdatedAt:   time.Now(),
	}
	msg,err:=post.usecase.UpdatePosts(id,posts)
	return msg,err
}

func (post *PostHttp) DeletePosts(id int)(string,error){
	msg,err:=post.usecase.DeletePosts(id)
	return msg,err
}

func (post *PostHttp) PostsById(id int)(model.Posts,error){
	posts,err:=post.usecase.PostsById(id)
	return posts,err
}

func (post *PostHttp) CreateProfile(userId int, firstName string, middleName string, lastName string, mobile string, email string, profile string) (string,error) {
	profiles:=model.Profiles{
		Id:           post.snowflake.GearedID(),
		UserId:       userId,
		FirstName:    firstName,
		MiddleName:   middleName,
		LastName:     lastName,
		Mobile:       mobile,
		Email:        email,
		Profile:      profile,
		RegisteredAt: time.Now(),
	}
	msg,err:=post.usecase.CreateProfile(profiles)
	return msg,err
}

func (post *PostHttp) UpdateProfile(id int, firstName string, middleName string, lastName string, mobile string, email string, profile string)(string,error) {
	profiles:=model.Profiles{
		FirstName:    firstName,
		MiddleName:   middleName,
		LastName:     lastName,
		Mobile:       mobile,
		Email:        email,
		Profile:      profile,
	}
	msg,err:=post.usecase.UpdateProfile(id,profiles)
	return msg,err
}

func (post *PostHttp) ProfileAll() ([]model.Profiles,error){
	profiles,err:=post.usecase.ProfileAll()
	return profiles,err
}

func (post *PostHttp) ProfileById(id int)(model.Profiles,error){
	profiles,err:=post.usecase.ProfileById(id)
	return profiles,err
}

func (post *PostHttp) CreatePostMetas(postId int, textKey string, content string) (string,error){
	metas:=model.PostMetas{
		Id:      post.snowflake.GearedID(),
		PostId:  postId,
		TextKey: textKey,
		Content: content,
	}
	msg,err:=post.usecase.CreatePostMetas(metas)
	return msg,err
}

func (post *PostHttp) UpdatePostMetas(id int,postId int, textKey string, content string) (string,error){
	metas:=model.PostMetas{
		PostId:  postId,
		TextKey: textKey,
		Content: content,
	}
	msg,err:=post.usecase.UpdatePostMetas(id,metas)
	return msg,err
}

func (post *PostHttp) DeletePostMetas(id int)(string,error) {
	msg,err:=post.usecase.DeletePostMetas(id)
	return msg,err
}

func (post *PostHttp) PostMetasAll() ([]model.PostMetas,error){
	metas,err:=post.usecase.PostMetasAll()
	return metas,err
}

func (post *PostHttp) PostMetasById(id int)(model.PostMetas,error){
	meta,err:=post.usecase.PostMetasById(id)
	return meta,err
}

func (post *PostHttp) CreateComment(postId int, parentId int, title string,content string) (string,error) {
	comment:=model.PostComments{
		Id:          post.snowflake.GearedID(),
		PostId:      postId,
		ParentId:    parentId,
		Title:       title,
		Published:   1,
		Content:     content,
		CreatedAt: time.Now(),
		PublishedAt: time.Now(),
	}
	msg,err:=post.usecase.CreateComment(comment)
	return msg,err
}

func (post *PostHttp) CommentAllPost(idPost int)([]model.PostComments,error){
	comments,err:=post.usecase.CommentAllPost(idPost)
	return comments,err
}

func (post *PostHttp) CreateCategories(parentId int, title string, metaTitle string, slug string, content string) (string,error){
	category:=model.Categories{
		Id:        post.snowflake.GearedID(),
		ParentId:  parentId,
		Title:     title,
		MetaTitle: metaTitle,
		Slug:      slug,
		Content:   content,
	}
	msg,err:=post.usecase.CreateCategories(category)
	return msg,err
}

func (post *PostHttp) UpdateCategories(id int,parentId int, title string, metaTitle string, slug string, content string) (string,error){
	category:=model.Categories{
		ParentId:  parentId,
		Title:     title,
		MetaTitle: metaTitle,
		Slug:      slug,
		Content:   content,
	}
	msg,err:=post.usecase.UpdateCategories(id,category)
	return msg,err
}

func (post *PostHttp) DeleteCategories(id int) (string,error){
	msg,err:=post.usecase.DeleteCategories(id)
	return msg,err	
}

func (post *PostHttp) CategoriesAll()([]model.Categories,error){
	categories,err:=post.usecase.CategoriesAll()
	return categories,err
}

func (post *PostHttp) CategoriesById(id int)(model.Categories,error){
	category,err:=post.usecase.CategoriesById(id)
	return category,err
}

func (post *PostHttp)ListProfilePost(id int)(model.ProfilePost,error){
	profilePost,err:=post.usecase.ListProfilePost(id)
	return profilePost,err
}

func (post *PostHttp) ListPostCategory() ([]model.PostCategory,error){
	postCategory,err:=post.usecase.ListPostCategory()
	return postCategory,err
}

func (post *PostHttp) PublishPost(id int) (string,error){
	msg,err:=post.usecase.PublishPost(id)
	return msg,err
}

func (post *PostHttp) DeleteComment(id int) (string,error){
	msg,err:=post.usecase.RemoveComment(id)
	return msg,err
}

func (post *PostHttp) ListPublishPost() ([]model.Posts,error){
	list,err:=post.usecase.ListPublishPost()
	return list,err
}

func (post *PostHttp) ListNotPublishPost() ([]model.Posts,error){
	list,err:=post.usecase.ListNotPublishPost()
	return list,err
}

func (post *PostHttp) PostsAll() ([]model.Posts,error){
	posts,err:=post.usecase.PostsAll()
	return posts,err
}

func (u *PostHttp) RequestIpComputer(ipAddress string,method string) error{
	ipMacAddress:=ip.Ip{
		Id:     u.snowflake.GearedID(),
		Ip:     ipAddress,
		Method: method,
		RequestAt:time.Now(),
	}
	return u.usecase.RequestIpComputer(ipMacAddress)
}