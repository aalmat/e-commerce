package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const userCtx = "userId"

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
