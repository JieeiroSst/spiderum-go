package delivery

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/upload/internal/s3/usecase"
	"net/http"
)

type Http struct {
	usecase *usecase.Usecase
	session *session.Session
}

func NewHttp(e *gin.Engine,usecase *usecase.Usecase,session *session.Session){
	h:=Http{usecase:usecase,session:session}

	resource := e.Group("/api")

	resource.POST("/file-dir",h.UploadFileDir)
	resource.POST("/read-file",h.ReadFileStream)
	resource.POST("/read-stream",h.UploadFileStream)
}

func (h *Http) UploadFileDir(e *gin.Context){
	file, err := e.FormFile("file")
	if err != nil {
		e.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	err = h.usecase.AddFileToS3(h.session,file.Filename)
	if err!=nil{
		e.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	e.JSON(http.StatusOK,map[string]string{
		"status":"200",
		"message":"upload file success",
	})
}

func (h *Http) UploadFileStream(e *gin.Context){
	data:=e.PostForm("data")
	err := h.usecase.AddFileToS3Stream(h.session,data)
	if err!=nil{
		e.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	e.JSON(http.StatusOK,map[string]string{
		"status":"200",
		"message":"upload file success",
	})
}

func (h *Http) ReadFileStream(e *gin.Context){
	key:=e.Query("data")
	data,err:=h.usecase.ReadFile(key)
	if err!=nil{
		e.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	e.JSON(http.StatusOK,map[string]string{
		"status":"200",
		"message":"read file success",
		"data":data,
	})
}