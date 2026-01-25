package main

import (
	"realtime-chatapp/internal/auth/delivery/http"
	"realtime-chatapp/internal/auth/delivery/http/route"
	"realtime-chatapp/internal/auth/repository"
	"realtime-chatapp/internal/auth/usecase"
	"realtime-chatapp/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db := config.NewDB()
	defer db.Close()
	r := gin.Default()

	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := http.NewAuthHandler(authUsecase)

	route.SetupAuthRoute(r, authHandler)
	r.Run(":8080")
}
