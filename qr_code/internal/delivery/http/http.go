package http

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"log"
)

type Http struct {}

func NewHttp(e *gin.Engine) {
	h:=Http{}
	resource := e.Group("/api")
	resource.GET("/qrcode",h.CreateQrCode)
}

func (h *Http) CreateQrCode(e *gin.Context){
	dataString := e.PostForm("dataString")
	img, err := qrcode.Encode(dataString, qrcode.Medium, 256)
	img2 := base64.StdEncoding.EncodeToString(img)
	if err != nil {
		log.Println(err)
	}

	e.JSON(200,gin.H{
		"images": img2,
	})
}
