package main

import (
	// Import Auth dengan Alias

	authHandler "realtime-chatapp/internal/auth/delivery/http"
	authRoute "realtime-chatapp/internal/auth/delivery/http/route"
	authRepo "realtime-chatapp/internal/auth/repository"
	authUsecase "realtime-chatapp/internal/auth/usecase"

	// Import Message dengan Alias
	msgHandler "realtime-chatapp/internal/message/delivery/ws"
	msgRoute "realtime-chatapp/internal/message/delivery/ws/route"
	msgRepo "realtime-chatapp/internal/message/repository"
	msgUsecase "realtime-chatapp/internal/message/usecase"

	"realtime-chatapp/internal/config"
	"realtime-chatapp/internal/message/delivery/ws/socket"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db := config.NewDB()
	defer db.Close()
	r := gin.Default()

	// 1. Inisialisasi Auth (Gunakan alias)
	aRepo := authRepo.NewAuthRepository(db)
	aUsecase := authUsecase.NewAuthUsecase(aRepo)
	aHandler := authHandler.NewAuthHandler(aUsecase)

	// 2. Inisialisasi Message & WebSocket
	hub := socket.NewHub()
	go hub.Run()

	mRepo := msgRepo.NewMessageRepository(db)
	mUsecase := msgUsecase.NewMessageUsecase(mRepo)
	mHandler := msgHandler.NewMessageHandler(mUsecase, hub)

	authRoute.SetupAuthRoute(r, aHandler)
	msgRoute.SetupMessageRoutes(r, mHandler)

	r.Run(":8080")
}
