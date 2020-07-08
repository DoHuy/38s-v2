package routers

import (
	"38s-v2/src/38s/config"
	"38s-v2/src/38s/handlers"
	"38s-v2/src/38s/repositories"
	"38s-v2/src/38s/services"
	"38s-v2/src/pkgs/constants"
	"38s-v2/src/pkgs/util"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router     *gin.Engine
	config     *config.Config
	httpServer *http.Server
}

func createRender(conf *config.Config) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	path := conf.ViewDir.Template
	urlDomain := func() string {
		return conf.ViewDir.DomainPath
	}
	//
	r.AddFromFilesFuncs(constants.HtmlTemplateNameHomePage, template.FuncMap{
		"urlDomain": urlDomain,
	}, path+"base.html")
	return r
}

func NewServer(
	router *gin.Engine,
	storeService repositories.Store,
	sessionService services.SessionService,
	conf *config.Config) *Server {
	handler := handlers.NewHandler(storeService, sessionService, conf)
	router = SetupRouter(router, handler, conf)

	httpServer := &http.Server{
		Addr:    conf.Server.ListenAddress,
		Handler: router,
	}

	server := &Server{
		router:     router,
		config:     conf,
		httpServer: httpServer,
	}
	return server
}

func SetupRouter(router *gin.Engine, handler *handlers.Handler, conf *config.Config) *gin.Engine {
	router.Use(util.Zap(zap.L(), time.RFC3339, true))
	router.HTMLRender = createRender(conf)
	router.Static("static", conf.ViewDir.Static)
	app := router.Group("/38s")
	{
		app.GET("/health", handler.Health)
		app.GET("/", handler.HomePage)

	}
	return router
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return nil
}
