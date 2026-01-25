package route

import (
	"realtime-chatapp/internal/auth/delivery/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(app *gin.Engine, handler *http.AuthHandlerImpl) {
	authGroup := app.Group("/v1/api/auth") 
	{
		authGroup.POST("/register", handler.Register)
        authGroup.POST("/login", handler.Login)
	}
}
