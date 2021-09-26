package http

import (
	"gitlab.com/Spide_IT/spide_it/internal/ip"
	"gitlab.com/Spide_IT/spide_it/internal/ip/model"
	"gitlab.com/Spide_IT/spide_it/pkg/snowflake"
)

type Http struct {
	usecase ip.IpUsecase
	snowflake snowflake.Snowflake
}

func Newhttp(usecase ip.IpUsecase,snowflake snowflake.Snowflake)*Http{
	return &Http{usecase:usecase,snowflake:snowflake}
}

func(h *Http) GetAllIp() ([]model.Ip,error){
	return h.usecase.FindIpAll()
}