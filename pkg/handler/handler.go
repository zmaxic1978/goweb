package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zmaxic1978/goweb/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/", h.userIdentity)
	{
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getAllBooks)
			books.GET("/:id", h.getBookById)
			books.PUT("/:id", h.setBookById)
			books.DELETE("/:id", h.deleteBookById)
			books.PUT("/:bookId/authors/authorId", h.setBookAuthor)
		}

		authors := api.Group("/authors")
		{
			authors.POST("/", h.createAuthor)
			authors.GET("/", h.getAllAuthors)
			authors.GET("/:id", h.getAuthorById)
			authors.PUT("/:id", h.setAuthorById)
			authors.DELETE("/:id", h.deleteAuthorById)
		}
	}

	return router
}
