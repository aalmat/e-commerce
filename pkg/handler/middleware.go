package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const userCtx = "userId"
const AuthHeader = "Authorization"

func (h *Handler) UserIdentify(ctx *gin.Context) {
	header := ctx.GetHeader(AuthHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "header is empty")
		return
	}

	headers := strings.Split(header, " ")
	if len(headers) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "wrong header format")
		return
	}

	id, err := h.service.ParseToken(headers[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(userCtx, id)

}

func GetUserId(ctx *gin.Context) (uint, error) {
	id, ok := ctx.Get(userCtx)

	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "List not found")
		return 0, errors.New("List not found")
	}

	intId, ok := id.(uint)

	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "User id invalid type")
		return 0, errors.New("User id invalid type")
	}

	return intId, nil
}
