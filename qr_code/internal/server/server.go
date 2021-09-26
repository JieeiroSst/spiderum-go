package server

import(
	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/qr_code/internal/delivery/http"
)

type QrcodeServer struct {
	server *gin.Engine
}

func NewQrcodeServer(server *gin.Engine) *QrcodeServer{
	return &QrcodeServer{server:server}
}

func (qrcode *QrcodeServer) Run(){
	http.NewHttp(qrcode.server)
}
