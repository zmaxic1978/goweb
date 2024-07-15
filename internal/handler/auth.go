package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	todo2 "github.com/zmaxic1978/goweb/todo"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo2.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {
	var input todo2.Login
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.CreateToken(input)
	if err != nil {
		if errors.As(err, new(todo2.AuthorizationError)) {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}

		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
