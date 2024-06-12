package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	todo "github.com/zmaxic1978/goweb"
	"net/http"
	"strconv"
)

func (h *Handler) setBookAuthor(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidBookId, err).Error())
		return
	}

	authorId, err := strconv.Atoi(c.Param("authorid"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidAuthorId, err).Error())
		return
	}

	var bookauthor todo.BookAuthor
	if err := c.BindJSON(&bookauthor); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errIncorrectBookFormat, err).Error())
		return
	}

	bookauthor.Book.Id = bookId
	bookauthor.Author.Id = authorId
	updated, err := h.services.Api.SetBookAuthorById(bookauthor)
	if err != nil {
		if errors.As(err, new(todo.BadFormatError)) || errors.As(err, new(todo.NoDataFound)) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"информация обновлена: ": bool(updated > 0)})
}
