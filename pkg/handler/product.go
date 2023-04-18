package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(ctx *gin.Context) {
	id, err := GetUserId(ctx)
	if err != nil {
		return
	}
	fmt.Print(id)

}
