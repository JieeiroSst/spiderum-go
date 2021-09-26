package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	email "gitlab.com/Spide_IT/spide_it/email/internal"
	"gitlab.com/Spide_IT/spide_it/email/internal/model"
)

type Http struct {
	usecase email.EmailsUseCase
}

func NewHttp(e *gin.Engine,usecase email.EmailsUseCase) {
	h:=&Http{usecase:usecase}
	g:=e.Group("/api")
	g.POST("/send-mail",h.SendEmail)
}

func (h *Http) SendEmail(e *gin.Context){
	email:=model.Email{
		Email:       e.PostForm("email"),
		To:          e.PostForm("to"),
		From:        e.PostForm("from"),
		Body:        e.PostForm("body"),
		Subject:     e.PostForm("subject"),
		ContentType: e.PostForm("content_type"),
	}
	dataByte,err:=json.Marshal(&email)
	if err!=nil{
		e.JSON(400,"Bad Request")
		return
	}
	err = h.usecase.SendEmail(e,dataByte)
	if err!=nil{
		e.JSON(400,"Bad Request")
		return
	}
	e.JSON(200,"Email sent successfully ")
}