package handler

import (
	"errors"
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AuthHeader = "Authorization"
	userCtx    = "userId"
	userRole   = "userRole"
)

func (h *Handler) UserIdentify(c *gin.Context) {
	header := c.GetHeader(AuthHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "header is empty")
		return
	}

	headers := strings.Split(header, " ")
	if len(headers) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "wrong header format")
		return
	}

	userId, role, err := h.service.ParseToken(headers[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
	c.Set(userRole, role)
}

func GetUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return 0, errors.New("user not found")
	}

	intId, ok := id.(uint)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User id invalid type")
		return 0, errors.New("user id invalid type")
	}
	return intId, nil
}

func GetUserRole(c *gin.Context) (models.Role, error) {
	role, ok := c.Get(userRole)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return 0, errors.New("user not found")
	}

	r, ok := role.(models.Role)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User role invalid type")
		return 0, errors.New("user role invalid type")
	}

	return r, nil
}
