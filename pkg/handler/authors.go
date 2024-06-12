package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	todo "github.com/zmaxic1978/goweb"
	"net/http"
	"strconv"
	"time"
)

const (
	errIncorrectAuthorFormat = "неверное передана информация по автору"
	errInvalidDateFormat     = "неверно передана дата в информации по автору, ожидается формат yyyy-mm-dd"
	errInvalidAuthorId       = "передан неверный Id автора"
)

type responseGetAllAuthors struct {
	Data []todo.Author `json:"data"`
}

func (h *Handler) createAuthor(c *gin.Context) {

	// пользователь авторизован
	/*if _, err := getUserId(c); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}*/

	var input todo.Author
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errIncorrectAuthorFormat, err).Error())
		return
	}

	if _, err := time.Parse("2006-01-02", input.Birthday); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidDateFormat, err).Error())
		return
	}

	id, err := h.services.Api.CreateAuthor(input)
	if err != nil {
		if errors.As(err, new(todo.BadFormatError)) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) getAllAuthors(c *gin.Context) {
	list, err := h.services.Api.GetAllAuthors()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responseGetAllAuthors{Data: list})
}

func (h *Handler) getAuthorById(c *gin.Context) {

	authorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidAuthorId, err).Error())
		return
	}

	author, err := h.services.Api.GetAuthorById(authorId)
	if err != nil && errors.As(err, new(todo.NoDataFound)) {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, author)
}

func (h *Handler) setAuthorById(c *gin.Context) {
	authorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidAuthorId, err).Error())
		return
	}

	var input todo.Author
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errIncorrectAuthorFormat, err).Error())
		return
	}

	if _, err := time.Parse("2006-01-02", input.Birthday); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidDateFormat, err).Error())
		return
	}

	input.Id = authorId
	cnt, err := h.services.Api.SetAuthorById(input)
	if err != nil {
		if errors.As(err, new(todo.BadFormatError)) || errors.As(err, new(todo.NoDataFound)) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"записей обновлено": cnt})

}

func (h *Handler) deleteAuthorById(c *gin.Context) {

	authorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("%s: %w", errInvalidAuthorId, err).Error())
		return
	}

	cnt, err := h.services.Api.DeleteAuthorById(authorId)
	if err != nil {
		if errors.As(err, new(todo.BadFormatError)) || errors.As(err, new(todo.NoDataFound)) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"записей удалено": cnt})

}
