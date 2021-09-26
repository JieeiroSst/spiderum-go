package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/internal/casbin/http"
	"gitlab.com/Spide_IT/spide_it/internal/casbin/model"
	"strconv"
)

type Http struct {
	http http.Http
}

func NewHttp(e *gin.Engine,http http.Http) {
	h:=Http{http:http}
	resource:=e.Group("/admin")
	resource.GET("/",h.CasbinRuleAll)
	resource.GET("/:id",h.CasbinRuleById)
	resource.POST("/create",h.CreateCasbinRule)
	resource.PUT("/ptype/:id",h.UpdateCasbinRulePtype)
	resource.PUT("/method/:id",h.UpdateCasbinMethod)
	resource.PUT("/endpoint/:id",h.UpdateCasbinRuleEndpoint)
	resource.PUT("/name/:id",h.UpdateCasbinRuleName)
	resource.DELETE("/:id",h.DeleteCasbinRule)
}

// CasbinRuleAll ... Casbin Rule All
// @Summary Casbin Rule All
// @Description get Casbin Rule All
// @Tags Casbins
// @Success 200 {string} string
// @Failure 500 {object} object
// @Router /admin/create [post]
func (http *Http) CasbinRuleAll(e *gin.Context){
	casbin,err:=http.http.CasbinRuleAll()
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"no data","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":casbin,"status":200})
}

// Casbin Rule By Id ... Get Casbin Rule By Id
// @Summary Get Casbin Rule By Id
// @Description get Casbin Rule By Id
// @Tags Casbins
// @Success 200 {string} string
// @Failure 500 {object} string
// @Router /admin/:id [get]
func (http *Http) CasbinRuleById(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	casbin,err:=http.http.CasbinRuleById(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"no data","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"data":casbin,"status":200})
}

// CreateCasbinRule ... Create Casbin Rule
// @Summary Create Casbin Rule
// @Description Create Casbin Rule
// @Tags Casbins
// @Success 200 {string} string
// @Failure 500 {object} string
// @Router / [post]
func (http *Http) CreateCasbinRule(e *gin.Context){
	ptype,v0,v1,v2:=e.Query("type"),e.Query("username"),e.Query("path"),e.Query("method")
	casbin:=model.CasbinRule{
		Ptype: ptype,
		V0:    v0,
		V1:    v1,
		V2:    v2,
	}
	fmt.Println(casbin.Ptype)
	err:=http.http.CreateCasbinRule(casbin)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"create failure","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":"create success","status":200})
}

// DeleteCasbinRule ... Delete Casbin Rule
// @Summary Delete Casbin Rule
// @Description Delete Casbin Rule
// @Tags Casbins
// @Success 200 {string} string
// @Failure 500 {object} string
// @Router /admin/:id [delete]
func (http *Http) DeleteCasbinRule(e *gin.Context){
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	err =http.http.DeleteCasbinRule(idInt)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"delete failure","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":"delete success","status":200})
}

// UpdateCasbinRulePtype ... Update Casbin Rule Ptype
// @Summary Update Casbin Rule Ptype
// @Description Update Casbin Rule Ptype
// @Tags Casbins
// @Success 200 {string} string
// @Failure 500 {object} string
// @Router /admin/ptype/:id [put]
func(http *Http) UpdateCasbinRulePtype(e *gin.Context){
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	ptype:=e.Query("ptype")
	err = http.http.UpdateCasbinRulePtype(idInt, ptype)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"update failure","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":"update success","status":200})
}

// UpdateCasbinRuleName ... UpdateCasbinRuleName
// @Summary UpdateCasbinRuleName
// @Description UpdateCasbinRuleName
// @Tags Casbins
// @Success 200 {string} string
// @Failure 500 {object} string
// @Router /admin/name/:id [put]
func(http *Http) UpdateCasbinRuleName(e *gin.Context) {
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	name:=e.Query("name")
	err =http.http.UpdateCasbinRuleName(idInt,name)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"update failure","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":"update success","status":200})

}

// UpdateCasbinRuleEndpoint ... UpdateCasbinRuleEndpoint
// @Summary UpdateCasbinRuleEndpoint
// @Description UpdateCasbinRuleEndpoint
// @Tags Casbins
// @Success 200 {string} string
// @Failure 500 {object} string
// @Router /admin/endponit/:id [put]
func(http *Http) UpdateCasbinRuleEndpoint(e *gin.Context){
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	endpoint:=e.Query("path")
	err =http.http.UpdateCasbinRuleEndpoint(idInt,endpoint)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"update failure","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":"update success","status":200})

}

// UpdateCasbinMethod ... UpdateCasbinMethod
// @Summary UpdateCasbinMethod
// @Description UpdateCasbinMethod
// @Tags Casbins
// @Success 200 {string} string
// @Failure 500 {object} string
// @Router /admin/methos/:id [put]
func(http *Http) UpdateCasbinMethod(e *gin.Context){
	id:=e.Param("id")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"cannot convert from string to int ","status":500})
		return
	}
	method:=e.Query("method")
	err =http.http.UpdateCasbinMethod(idInt,method)
	if err!=nil{
		e.JSON(500,map[string]interface{}{"message":"update failure","status":500})
		return
	}
	e.JSON(200,map[string]interface{}{"message":"update success","status":200})
}