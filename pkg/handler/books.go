package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createBook(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{"id:": id})
}

func (h *Handler) getAllBooks(c *gin.Context) {

}

func (h *Handler) getBookById(c *gin.Context) {

}

func (h *Handler) setBookById(c *gin.Context) {

}

func (h *Handler) deleteBookById(c *gin.Context) {

}
