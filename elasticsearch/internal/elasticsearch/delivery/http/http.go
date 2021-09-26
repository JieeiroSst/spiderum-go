package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/elasticsearch/http"
)

type Http struct {
	http *http.Http
}

func NewHttp(e *gin.Engine,http *http.Http){
	h:=Http{http:http}
	g:=e.Group("/api")
	g.GET("/search",h.SearchPost)
}

func (h *Http) SearchPost(e *gin.Context){
	name:=e.Query("name")
	result,err:=h.http.Query(e,name)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"data":"no data"})
		return
	}
	e.JSON(200,map[string]interface{}{"result":result})
}