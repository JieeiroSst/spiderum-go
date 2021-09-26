package repository

import (
	ip "gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/internal/post/model"
	"gorm.io/gorm"
	"time"
)

type PostsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) *PostsRepository{
	_ = db.AutoMigrate(&model.Posts{},&model.Categories{},&model.PostComments{},&model.Profiles{})
	_ = db.AutoMigrate(&model.PostMetas{})
	return &PostsRepository{db:db}
}

func (post *PostsRepository) CreatePosts(posts model.Posts) (string,error){
	err:=post.db.Create(&posts).Error
	if err!=nil{
		return "create failed",err
	}
	return "create success", nil
}

func (post *PostsRepository) UpdatePosts(id int,posts model.Posts) (string,error){
	err:=post.db.Model(model.Posts{}).Where("id = ? ", id).Updates(posts).Error
	if err!=nil{
		return "update failed",err
	}
	return "update success", nil
}

func (post *PostsRepository) DeletePosts(id int)(string,error){
	err:=post.db.Delete(model.Posts{}, "id = ?", id).Error
	if err!=nil{
		return "delete failed",err
	}
	return "delete success", nil
}

func (post *PostsRepository) PostsAll()([]model.Posts,error){
	var posts []model.Posts
	post.db.Find(&posts)
	return posts, nil
}

func (post *PostsRepository) PostsById(id int)(model.Posts,error){
	var posts model.Posts
	post.db.Where("id = ?", id).Find(&posts)
	return posts, nil
}

func (post *PostsRepository) CreateCategories(categories model.Categories) (string,error){
	err:=post.db.Create(&categories).Error
	if err!=nil{
		return "create failed",err
	}
	return "create success", nil
}

func (post *PostsRepository) UpdateCategories(id int,categories model.Categories) (string,error){
	err:=post.db.Model(model.Categories{}).Where("id = ? ", id).Updates(categories).Error
	if err!=nil{
		return "update failed",err
	}
	return "update success", nil
}

func (post *PostsRepository) DeleteCategories(id int) (string,error){
	err:=post.db.Delete(model.Categories{}, "id = ?", id).Error
	if err!=nil{
		return "delete failed",err
	}
	return "delete success", nil
}

func (post *PostsRepository) CategoriesAll()([]model.Categories,error){
	var categories []model.Categories
	post.db.Find(&categories)
	return categories, nil
}

func (post *PostsRepository) CategoriesById(id int)(model.Categories,error){
	var category model.Categories
	post.db.Where("id = ?", id).Find(&category)
	return category, nil
}

func (post *PostsRepository) CreateComment(comment model.PostComments) (string,error){
	err:=post.db.Create(&comment).Error
	if err!=nil{
		return "create failed",err
	}
	return "create success", nil
}

func (post *PostsRepository) CommentAllPost(idPost int)([]model.PostComments,error){
	var comments []model.PostComments
	post.db.Where("postId=?",idPost).Find(&comments)
	return comments, nil
}

func (post *PostsRepository) CreatePostMetas(metas model.PostMetas) (string,error){
	err:=post.db.Create(&metas).Error
	if err!=nil{
		return "create failed",err
	}
	return "create success", nil
}

func (post *PostsRepository) UpdatePostMetas(id int,metas model.PostMetas) (string,error){
	err:=post.db.Model(model.PostMetas{}).Where("id = ? ", id).Updates(metas).Error
	if err!=nil{
		return "update failed",err
	}
	return "update success", nil
}

func (post *PostsRepository) DeletePostMetas(id int)(string,error){
	err:=post.db.Delete(model.PostMetas{}, "id = ?", id).Error
	if err!=nil{
		return "delete failed",err
	}
	return "delete success", nil
}

func (post *PostsRepository) PostMetasAll()([]model.PostMetas,error){
	var postMetas []model.PostMetas
	result := post.db.Find(&postMetas)
	if result != nil{
		return nil, result.Error
	}
	return postMetas, nil
}

func (post *PostsRepository) PostMetasById(id int)(model.PostMetas,error){
	var postMeta model.PostMetas
	post.db.Where("id = ?", id).Find(&postMeta)
	return postMeta, nil
}

func (post *PostsRepository) CreateProfile(profile model.Profiles) (string,error){
	err:=post.db.Create(&profile).Error
	if err!=nil{
		return "create failed",err
	}
	return "create success", nil
}

func (post *PostsRepository)UpdateProfile(id int,profile model.Profiles) (string,error){
	err:=post.db.Model(model.Profiles{}).Where("id = ? ", id).Updates(profile).Error
	if err!=nil{
		return "update failed",err
	}
	return "update success", nil
}

func (post *PostsRepository) ProfileAll() ([]model.Profiles,error){
	var profiles []model.Profiles
	post.db.Find(&profiles)
	return profiles, nil
}

func (post *PostsRepository) ProfileById(id int)(model.Profiles,error){
	var profile model.Profiles
	post.db.Where("id = ?", id).Find(&profile)
	return profile, nil
}

func (post *PostsRepository) PostByAuthorID(authorId  int) ([]model.Posts,error){
	var posts []model.Posts
	post.db.Where("authorId = ?",authorId).Find(&posts)
	return posts, nil
}

func (post *PostsRepository) CategoryByPostId(parentId int) (model.Categories,error){
	var category model.Categories
	post.db.Where("parentId = ?",parentId).Find(&category)
	return category, nil
}

func (post *PostsRepository) PublishPost(id int) (string,error) {
	err:=post.db.Model(model.Posts{}).Where("id = ? ", id).UpdateColumns(model.Posts{Published:1,PublishedAt:time.Now()})
	if err!=nil{
		return "update failed",err.Error
	}
	return "update success", nil
}

func (post *PostsRepository) RemoveComment(id int) (string,error){
	err:=post.db.Model(model.Posts{}).Where("id = ? ", id).Update("published", 0).Error
	if err!=nil{
		return "remove failed",err
	}
	return "remove success", nil
}

func (post *PostsRepository) ListPublishPost() ([]model.Posts,error){
	var posts []model.Posts
	post.db.Where("published = ?" ,1).Find(&posts)
	return posts,nil
}

func (post *PostsRepository) ListNotPublishPost() ([]model.Posts,error){
	var posts []model.Posts
	post.db.Where("published = ?" ,0).Find(&posts)
	return posts,nil
}

func (post *PostsRepository) RequestIpComputer(ip ip.Ip) error{
	err:=post.db.Create(&ip).Error
	return err
}
