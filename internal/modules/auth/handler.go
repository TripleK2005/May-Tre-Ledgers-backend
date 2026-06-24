package auth

import (
	"errors"
	"net/http"

	"may-tre-ledger-be/internal/core/response"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Register(c.Request.Context(), req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "register success", gin.H{})
}

func (h *handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	authResponse, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		writeAuthError(c, err)
		return
	}

	response.Success(c, http.StatusOK, "login success", authResponse)
}

func (h *handler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	authResponse, err := h.service.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		writeAuthError(c, err)
		return
	}

	response.Success(c, http.StatusOK, "refresh token success", authResponse)
}

func (h *handler) Logout(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Logout(c.Request.Context(), req.RefreshToken); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "logout success", gin.H{})
}

func writeAuthError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrInvalidCredentials):
		response.Error(c, http.StatusUnauthorized, "invalid username or password")
	case errors.Is(err, ErrInactiveUser):
		response.Error(c, http.StatusForbidden, "user is inactive")
	case errors.Is(err, ErrInvalidToken):
		response.Error(c, http.StatusUnauthorized, "invalid token")
	case errors.Is(err, ErrExpiredToken):
		response.Error(c, http.StatusUnauthorized, "expired token")
	default:
		response.Error(c, http.StatusInternalServerError, err.Error())
	}
}
