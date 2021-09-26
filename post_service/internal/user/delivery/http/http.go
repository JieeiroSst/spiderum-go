package http

import (
	"fmt"
	"github.com/durango/gin-passport-facebook"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/Spide_IT/spide_it/internal/user/http"
	"gitlab.com/Spide_IT/spide_it/pkg/bigcache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type Http struct {
	http *http.UserHTTP
	cache bigcache.Cache
}

func NewUserHTTP(e *gin.Engine,cache bigcache.Cache,http *http.UserHTTP){
	h:=Http{
		http:  http,
		cache: cache,
	}
	opts := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/facebook/callback",
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		Scopes:       []string{"email", "public_profile"},
		Endpoint:     facebook.Endpoint,
	}

	auth := e.Group("/auth/facebook")
	GinPassportFacebook.Routes(opts, auth)

	auth.GET("/callback", GinPassportFacebook.Middleware(),h.LoginFacebook)

	g:=e.Group("/user")
	g.POST("/login",h.Login)
	g.POST("/singup",h.SignUp)
}

func (h *Http) LoginFacebook(e *gin.Context){
	user, err := GinPassportFacebook.GetProfile(e)
	if user == nil || err != nil {
		e.AbortWithStatus(500)
		return
	}

	e.String(200, "Got it!")
}

// Login ... Login
// @Summary Login
// @Description Login
// @Tags Login
// @Success 200 {string} string
// @Failure 500 {object} object
// @Router /user/login [post]
func (h *Http) Login(e *gin.Context) {
	username,password:=e.Query("username"),e.Query("password")
	msg,id:=h.http.Login(username,password)
	if len(msg) == 0 {
		e.JSON(503,map[string]string{"token":"null"})
		return
	}
	u, err := uuid.NewRandom()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), username)
	err = h.cache.Set(sessionId,[]byte(username))
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
	e.SetCookie("current_subject", sessionId, 30*60, "/api", "", false, true)
	e.JSON(200,map[string]interface{}{"token":msg,"id":id})
}

// SignUp ... SignUp
// @Summary SignUp
// @Description SignUp
// @Tags Users
// @Success 200 {string} string
// @Failure 500 {object} object
// @Router /user/singup [post]
func (h *Http) SignUp(e *gin.Context){
	username,password:=e.Query("username"),e.Query("password")
	msg:=h.http.Signup(username,password)
	if len(msg) == 0 {
		e.JSON(500,map[string]interface{}{"data":"no data","status":500})
		return
	}
	ip:=e.ClientIP()
	method:=e.Request.Method
	err := h.http.RequestIpComputer(ip,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no get ip","status":500})
		return
	}

	e.JSON(200,map[string]interface{}{"data":msg,"status":200})
	return
}