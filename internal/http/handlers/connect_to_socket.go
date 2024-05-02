package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log/slog"
	"vangram_api/internal/service/message"
	"vangram_api/internal/websocket"
)

func (h *Handler) connectToSocket(ctx *gin.Context) {
	socket := ws.New()
	ws, err := socket.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		slog.Error(err.Error())
	}
	go receiver(ws, ctx, h)
}

func receiver(ws *websocket.Conn, ctx *gin.Context, h *Handler) {
	defer ws.Close()
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			slog.Error(err.Error())
		}
		var socketMessage socketMessage
		err = json.Unmarshal(p, &socketMessage)
		if err != nil {
			slog.Error(err.Error())
			ws.WriteJSON(map[string]interface{}{
				"error": "json parse error",
			})
			return
		}
		fmt.Sprintf("Received message: %q", socketMessage)
		switch socketMessage.Type {
		case "send_message":
			id, err := h.messageService.AddNewMessage(ctx, socketMessage.Message)
			if err != nil {
				slog.Error(err.Error())
				ws.WriteJSON(map[string]interface{}{
					"error": "Send message error",
				})
				return
			}
			slog.Info(string(id))
			ws.WriteJSON(map[string]interface{}{
				"id": id,
			})
		}

	}
}

type socketMessage struct {
	Type    string                `json:"type"`
	Message message.CreateMessage `json:"message"`
}
