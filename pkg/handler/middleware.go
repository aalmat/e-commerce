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

func ValidateHeader(header string) (string, error) {
	if header == "" {
		return "", errors.New("header is empty")
	}

	headers := strings.Split(header, " ")
	if len(headers) != 2 {
		return "", errors.New("wrong header format")
	}

	return headers[1], nil
}

func (h *Handler) UserIdentify(ctx *gin.Context) {
	header := ctx.GetHeader(Authorization)
	token, err := ValidateHeader(header)

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	id, role, err := h.service.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
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
