package http

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/internal/ip/http"
	"gitlab.com/Spide_IT/spide_it/middleware"
	"gorm.io/gorm"
)

type Http struct {
	http http.Http
	auth middleware.Authorization
}

func NewIpHttp(e *gin.Engine,http http.Http,auth middleware.Authorization,db *gorm.DB){
	h:=Http{
		http: http,
		auth: auth,
	}
	adapter,_:=gormadapter.NewAdapterByDB(db)

	resource := e.Group("/api")
	resource.Use(auth.Authenticate())
	{
		resourceIp:=resource.Group("/admin")
		resourceIp.GET("/ip",auth.Authorize("/api/admin/*","GET",adapter),h.GetIpAll)
	}
}

func (h *Http) GetIpAll(e *gin.Context){
	ips,err:=h.http.GetAllIp()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"status":500,"data":"no data"})
		return
	}
	e.JSON(200,map[string]interface{}{"status":200,"data":ips})
}