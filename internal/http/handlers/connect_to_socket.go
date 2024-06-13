package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"vangram_api/internal/http/middleware"
	"vangram_api/internal/service/message"
	ws "vangram_api/internal/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*Client]bool)
var broadcast = make(chan *socketMessage)

type Client struct {
	Conn   *websocket.Conn
	UserID int
}

func (h *Handler) connectToSocket(ctx *gin.Context) {
	socket := ws.New()
	ws, err := socket.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		slog.Error(err.Error())
	}
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	client := &Client{Conn: ws, UserID: userID}
	clients[client] = true
	go broadcastMessage(h, ctx)
	fmt.Println("clients:", len(clients), clients, ws.RemoteAddr())
	receiver(client)
	delete(clients, client)
}

func receiver(client *Client) {
	defer client.Conn.Close()
	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		var socketMessage socketMessage
		err = json.Unmarshal(p, &socketMessage)
		if err != nil {
			slog.Error(err.Error())
			err := client.Conn.WriteJSON(map[string]interface{}{
				"error": "json parse error",
			})
			if err != nil {
				slog.Error(err.Error())
				return
			}
			return
		}
		broadcast <- &socketMessage
		fmt.Println("host", client.Conn.RemoteAddr())
	}
}

func broadcastMessage(h *Handler, ctx *gin.Context) {
	for {
		m := <-broadcast
		for client := range clients {
			userID, err := middleware.GetUserID(ctx)
			if err != nil {
				slog.Error(err.Error())
				return
			}
			if client.UserID == m.Message.IDUser {
				_, err := h.messageService.AddNewMessage(ctx, m.Message, userID)
				if err != nil {
					slog.Error(err.Error())
					err := client.Conn.WriteJSON(map[string]interface{}{
						"error": "Send message error",
					})
					if err != nil {
						return
					}
					return
				}
				err = client.Conn.WriteJSON(m)
				if err != nil {
					slog.Error("Websocket error: %s", err)
					err := client.Conn.Close()
					if err != nil {
						return
					}
					delete(clients, client)
				}
			}
		}

	}
}

type socketMessage struct {
	Type    string                `json:"type"`
	Message message.CreateMessage `json:"message"`
}
