package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) getMessages(ctx *gin.Context) {
	var r getMessagesRequest
	if err := ctx.BindJSON(&r); err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messages, err := h.messageService.GetChatMessages(ctx, r.IDChat)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, messages)

}

type getMessagesRequest struct {
	IDChat int `json:"id_chat"`
}
