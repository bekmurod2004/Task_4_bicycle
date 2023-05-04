package handler

import (
	"app/api/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetPromo(c *gin.Context) {
	var action models.StigmaApi

	err := c.ShouldBindJSON(&action)
	if err != nil {
		h.handlerResponse(c, "new error", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.storages.Code().PromoView(context.Background(), &action)
	if err != nil {
		h.handlerResponse(c, "_", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "answer", http.StatusCreated, resp)

}
