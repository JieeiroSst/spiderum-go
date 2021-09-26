package http

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/internal/post/http"
	"gitlab.com/Spide_IT/spide_it/middleware"
	"gorm.io/gorm"
	"strconv"
)

type Http struct {
	http http.PostHttp
	auth middleware.Authorization
}

func NewPostHttp(e *gin.Engine,http http.PostHttp,auth middleware.Authorization,db *gorm.DB){
	h:=Http{
		http:http,
		auth:auth,
	}
	adapter,_:=gormadapter.NewAdapterByDB(db)

	resource := e.Group("/api")
	resource.Use(auth.Authenticate())
	{
		resourceUser:=resource.Group("/user")
		resourceUser.POST("/profile",auth.Authorize("/api/user/*","POST",adapter),h.CreateProfile)
		resourceUser.PUT("/profile/:id",auth.Authorize("/api/user/*","PUT",adapter),h.UpdateProfile)
		resourceUser.POST("/post",auth.Authorize("/api/user/*","POST",adapter),h.CreatePosts)
		resourceUser.PUT("/post/:id",auth.Authorize("/api/user/*","PUT",adapter),h.UpdatePosts)
		resourceUser.DELETE("/post/:id",auth.Authorize("/api/user/*","DELETE",adapter),h.DeletePosts)
		resourceUser.GET("/post",auth.Authorize("/api/user/*","GET",adapter),h.PostsAll)
		resourceUser.GET("/post/:id",auth.Authorize("/api/user/*","GET",adapter),h.PostsById)
		resourceUser.POST("/comment",auth.Authorize("/api/user/*","POST",adapter),h.CreateComment)
		resourceUser.GET("/category",auth.Authorize("/api/user/*","GET",adapter),h.CategoriesAll)
		resourceUser.GET("/category/:id",auth.Authorize("/api/user/*","GET",adapter),h.CategoriesById)
		resourceUser.GET("/profile_post/:id",auth.Authorize("/api/user/*","GET",adapter),h.ListProfilePost)
		resourceUser.GET("/post_category",auth.Authorize("/api/user/*","GET",adapter),h.ListPostCategory)
		resourceUser.GET("/post/publish/:id",auth.Authorize("/api/user/*","GET",adapter),h.PublishPost)
		resourceUser.DELETE("/comment/:id",auth.Authorize("/api/user/*","DELETE",adapter),h.RemoveComment)
		resourceUser.GET("/list_post",auth.Authorize("/api/user/*","GET",adapter),h.ListPublishPost)
		resourceUser.GET("/list_not_post",auth.Authorize("/api/user/*","GET",adapter),h.ListNotPublishPost)

		resourceAdmin:=resource.Group("/admin")
		resourceAdmin.GET("/profile",auth.Authorize("/api/admin/*","GET",adapter),h.ProfileAll)
		resourceAdmin.GET("/profile/:id",auth.Authorize("/api/admin/*","GET",adapter),h.ProfileById)
		resourceAdmin.POST("/post_metas",auth.Authorize("/api/admin/*","GET",adapter),h.CreatePostMetas)
		resourceAdmin.PUT("/post_metas/:id",auth.Authorize("/api/admin/*","PUT",adapter),h.UpdatePostMetas)
		resourceAdmin.DELETE("/post_metas/:id",auth.Authorize("/api/admin/*","DELETE",adapter),h.DeletePostMetas)
		resourceAdmin.GET("/post_metas",auth.Authorize("/api/admin/*","GET",adapter),h.PostMetasAll)
		resourceAdmin.GET("/post_metas/:id",auth.Authorize("/api/admin/*","GET",adapter),h.PostMetasById)
		resourceAdmin.GET("/comment",auth.Authorize("/api/admin/*","GET",adapter),h.CommentAllPost)
		resourceAdmin.POST("/category",auth.Authorize("/api/admin/*","POST",adapter),h.CreateCategories)
		resourceAdmin.PUT("/category/:id",auth.Authorize("/api/admin/*","PUT",adapter),h.UpdateCategories)
		resourceAdmin.GET("/category",auth.Authorize("/api/admin/*","GET",adapter),h.CategoriesAll)
		resourceAdmin.GET("/category/:id",auth.Authorize("/api/admin/*","GET",adapter),h.CategoriesById)
		resourceAdmin.GET("/post",auth.Authorize("/api/admin/*","GET",adapter),h.PostsAll)

		resourceClient:=resource.Group("/client")
		resourceClient.GET("/post",auth.Authorize("/api/client/*","GET",adapter),h.PostsAll)
		resourceClient.GET("/post/:id",auth.Authorize("/api/client/*","GET",adapter),h.PostsById)
	}

}

// CreatePosts ... CreatePosts
// @Summary CreatePosts
// @Description CreatePosts
// @Tags CreatePosts
// @Success 200 {string} string
// @Failure 500 {object} object
// @Router /api/posts [post]
func (h *Http) CreatePosts(e *gin.Context) {
	authorId,title,metaTitle:=e.Query("authorId"),e.Query("title"),e.Query("metaTitle")
	slug,summary,content:=e.Query("slug"),e.Query("summary"),e.Query("content")
	authorIdInt,err:=strconv.Atoi(authorId)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.CreatePosts(authorIdInt,title,metaTitle,slug,summary,content)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// UpdatePosts ... UpdatePosts
// @Summary UpdatePosts
// @Description UpdatePosts
// @Tags Posts
// @Success 200 {string} string
// @Failure 500 {object} object
// @Router /api/posts/:id [put]
func (h *Http) UpdatePosts(e *gin.Context){
	title,metaTitle:=e.Query("title"),e.Query("metaTitle")
	slug,summary,content:=e.Query("slug"),e.Query("summary"),e.Query("content")

	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.UpdatePosts(idInt,title,metaTitle,slug,summary,content)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// DeletePosts ... DeletePosts
// @Summary Delete Posts
// @Description Delete Posts
// @Tags posts
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router /api/delete [delete]
func (h *Http) DeletePosts(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.DeletePosts(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// PostsById ... PostsById
// @Summary Get user by id
// @Description get all post
// @Tags post
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router /api/posts/:id [get]
func (h *Http) PostsById(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.PostsById(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"post":msg,"status":200})
}

// CreateProfile ... Create Profile
// @Summary Create Profile
// @Description Create Profile
// @Tags Users
// @Success 200 {string} string
// @Failure 500 {object} object
// @Router /api/profile [post]
func (h *Http) CreateProfile(e *gin.Context) {
	userId,firstName,middleName:=e.Query("userId"),e.Query("firstName"),e.Query("middleName")
	lastName,mobile,email,profile:=e.Query("lastName"),e.Query("mobile"),e.Query("email"),e.Query("profile")

	idInt,err:=strconv.Atoi(userId)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}

	msg,err:=h.http.CreateProfile(idInt,firstName,middleName,lastName,mobile,email,profile)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// UpdateProfile ... Update Profile
// @Summary Update Profile
// @Description Update Profile
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router /api/profile/:id [put]
func (h *Http) UpdateProfile(e *gin.Context) {
	firstName,middleName:=e.Query("firstName"),e.Query("middleName")
	lastName,mobile,email,profile:=e.Query("lastName"),e.Query("mobile"),e.Query("email"),e.Query("profile")
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.UpdateProfile(idInt,firstName,middleName,lastName,mobile,email,profile)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// ProfileAll ... Profile All
// @Summary Profile All
// @Description Profile All
// @Tags profile
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) ProfileAll(e *gin.Context){
	profile,err:=h.http.ProfileAll()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"profile":profile,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) ProfileById(e *gin.Context) {
	id:=e.Param("id")

	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	profile,err:=h.http.ProfileById(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"profile":profile,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) CreatePostMetas(e *gin.Context){
	textKey,content,postId:=e.Query("textKey"),e.Query("content"),e.Query("postId")
	postIdInt,err:=strconv.Atoi(postId)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.CreatePostMetas(postIdInt,textKey,content)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) UpdatePostMetas(e *gin.Context) {
	textKey,content,postId:=e.Query("textKey"),e.Query("content"),e.Query("postId")
	postIdInt,err:=strconv.Atoi(postId)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.UpdatePostMetas(idInt,postIdInt,textKey,content)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) DeletePostMetas(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.DeletePostMetas(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) PostMetasAll(e *gin.Context) {
	post,err:=h.http.PostMetasAll()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":post,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) PostMetasById(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	post,err:=h.http.PostMetasById(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":post,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) CreateComment(e *gin.Context) {
	postId,parentId,title,content:=e.Query("postId"),e.Query("parentId"),e.Query("title"),e.Query("content")
	postIdInt,err:=strconv.Atoi(postId)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	parentIdInt,err:=strconv.Atoi(parentId)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.CreateComment(postIdInt,parentIdInt,title,content)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) CommentAllPost(e *gin.Context) {
	idPost:=e.Param("idPost")
	idPostInt,err:=strconv.Atoi(idPost)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	comment,err:=h.http.CommentAllPost(idPostInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":comment,"status":200})

}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) CreateCategories(e *gin.Context){
	parentId,title,metaTitle,slug,content:=e.Query("parentId"),e.Query("title"),e.Query("metaTitle"),e.Query("slug"),e.Query("content")
	parentIdInt,err:=strconv.Atoi(parentId)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}

	msg,err:=h.http.CreateCategories(parentIdInt,title,metaTitle,slug,content)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) UpdateCategories(e *gin.Context){
	parentId,title,metaTitle,slug,content:=e.Query("parentId"),e.Query("title"),e.Query("metaTitle"),e.Query("slug"),e.Query("content")
	parentIdInt,err:=strconv.Atoi(parentId)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.UpdateCategories(idInt,parentIdInt,title,metaTitle,slug,content)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) DeleteCategories(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.DeleteCategories(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) CategoriesAll(e *gin.Context) {
	categories,err:=h.http.CategoriesAll()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":categories,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) CategoriesById(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	category,err:=h.http.CategoriesById(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":category,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) ListProfilePost(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	profilePost,err:=h.http.ListProfilePost(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":profilePost,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) ListPostCategory(e *gin.Context) {
	postcategory,err:=h.http.ListPostCategory()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":postcategory,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) PublishPost(e *gin.Context){
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.PublishPost(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) RemoveComment(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	msg,err:=h.http.DeleteComment(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":msg,"status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":msg,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) ListPublishPost(e *gin.Context) {
	list,err:=h.http.ListPublishPost()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":list,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) ListNotPublishPost(e *gin.Context) {
	list,err:=h.http.ListNotPublishPost()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":list,"status":200})
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {string} string
// @Failure 404 {object} object
// @Router / [get]
func (h *Http) PostsAll(e *gin.Context) {
	post,err:=h.http.PostsAll()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err = h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":post,"status":200})
}