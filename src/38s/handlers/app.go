package handlers

import (
	"38s-v2/src/38s/config"
	"38s-v2/src/38s/repositories"
	"38s-v2/src/38s/services"
	"38s-v2/src/pkgs/constants"
	"38s-v2/src/pkgs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)


type Handler struct {
	storeService   repositories.Store
	sessionService services.SessionService
	config         *config.Config
}

func NewHandler(
	storeService repositories.Store, sessionService services.SessionService,
	config *config.Config) *Handler {
	return &Handler{
		storeService:   storeService,
		sessionService: sessionService,
		config:         config,

	}
}

func (h *Handler) Health(ctx *gin.Context) {
	util.BuildSuccessResponse(ctx, "Alive")
}

func (h *Handler) HomePage(ctx *gin.Context) {
	dataView := gin.H{}
	ctx.HTML(http.StatusOK, constants.HtmlTemplateNameHomePage, dataView)
}