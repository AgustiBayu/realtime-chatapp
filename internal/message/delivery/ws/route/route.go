package route

import (
	"realtime-chatapp/internal/domain"
	"realtime-chatapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(r *gin.Engine, handler domain.MessageHandler) {
	messageGroup := r.Group("/messages")
	{
		messageGroup.GET("/history", middleware.JWTMiddleware(), handler.GetChatHistory)
		messageGroup.GET("/ws", handler.Connect)
	}
}
