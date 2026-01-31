package route

import (
	"realtime-chatapp/internal/message/delivery/ws"
	"realtime-chatapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(r *gin.Engine, handler *ws.MessageHandlerImpl) {
	messageGroup := r.Group("/messages")
	{
		messageGroup.GET("/history", middleware.JWTMiddleware(), handler.GetChatHistory)
		messageGroup.GET("/ws", handler.Connect)
	}
}
