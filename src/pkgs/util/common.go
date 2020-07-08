package util

import (
	"38s-v2/src/pkgs/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"38s-v2/src/38s/config"
)

func BuildResponse(context *gin.Context, code int, body interface{}) {
	context.JSON(code, body)
}

func BuildSuccessResponse(context *gin.Context, body interface{}) {
	BuildResponse(context, http.StatusOK, dtos.BaseResponse{
		Data: body,
		Meta: dtos.Meta{Code: http.StatusOK},
	})
}

func CreateRedirectUrl(conf *config.Config, url string) string {
	return conf.ViewDir.DomainPath + url
}
