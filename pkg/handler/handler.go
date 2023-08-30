package handler

import (
	"CRUD/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		news := api.Group("/news")
		{
			news.GET("/", h.getAllNews)
			news.GET("/:id", h.getNewsById)
			news.POST("/", h.postNewNews)
			news.PUT("/:id", h.changeNewsById)
			news.DELETE("/:id", h.removeNews)
		}
		followers := api.Group("/subscribers")
		{
			followers.GET("/:id", h.getAllFollowersByNewsID)
			followers.POST("/", h.subscribeToNews)
			followers.DELETE("/:id", h.UnsubscribeFromNews)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
