package http

import (
	"net/http"
	"realtime-chatapp/internal/domain"

	"github.com/gin-gonic/gin"
)

type AuthHandlerImpl struct {
	Usecase domain.AuthUsecase
}

func NewAuthHandler(usescase domain.AuthUsecase) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		Usecase: usescase,
	}
}

func (a *AuthHandlerImpl) Register(c *gin.Context) {
	var request domain.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Data not falid",
		})
		return
	}
	err := a.Usecase.Regiter(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, domain.WebResponse{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Successful register",
	})
}

func (a *AuthHandlerImpl) Login(c *gin.Context) {
	var request domain.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email dan password diperlukan"})
		return
	}
	response, err := a.Usecase.Login(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Email dan password diperlukan",
		})
		return
	}

	response, err = a.Usecase.Login(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successful login",
		Data:    response,
	})
}
