package handler

import (
	"errors"
	"github.com/aalmat/e-commerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	Authorization = "Authorization"
	UserId        = "userId"
	UserRole      = "userRole"
)

func (h *Handler) UserIdentify(ctx *gin.Context) {
	header := ctx.GetHeader(Authorization)
	if header == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "header is empty")
	}

	headers := strings.Split(header, " ")
	if len(headers) != 2 {
		newErrorResponse(ctx, http.StatusBadRequest, "wrong header format")
	}

	token := headers[1]
	id, role, err := h.service.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	ctx.Set(UserId, id)
	ctx.Set(UserRole, role)
}

func (h *Handler) GetUserId(ctx *gin.Context) (uint, error) {
	id, ok := ctx.Get(UserId)
	if !ok {
		newErrorResponse(ctx, http.StatusBadRequest, "user not found")
		return 0, errors.New("user not found")
	}

	intId, ok := id.(uint)

	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "User id invalid type")
		return 0, errors.New("user id invalid type")
	}
	return intId, nil
}

func (h *Handler) GetUserRole(ctx *gin.Context) (models.Role, error) {
	role, ok := ctx.Get(UserRole)
	if !ok {
		newErrorResponse(ctx, http.StatusBadRequest, "role not found")
		return 0, errors.New("role not found")
	}
	conRole, ok := role.(models.Role)
	if !ok {
		newErrorResponse(ctx, http.StatusBadRequest, "role not found")
		return 0, errors.New("role not found")
	}

	return conRole, nil
}
