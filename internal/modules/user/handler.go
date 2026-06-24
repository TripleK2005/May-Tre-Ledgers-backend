package user

import (
	"may-tre-ledger-be/internal/core/response"
	"may-tre-ledger-be/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	currentUserID, hasUserID := c.Get(middleware.ContextUserID)
	currentRole, hasRole := c.Get(middleware.ContextRole)
	if !hasUserID || !hasRole {
		response.Error(c, http.StatusUnauthorized, "missing auth context")
		return
	}

	if currentRole != "ADMIN" && currentUserID != id {
		response.Error(c, http.StatusForbidden, "permission denied")
		return
	}

	user, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		response.Error(c, http.StatusNotFound, "user not found")
		return
	}

	resp := UserResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Role:      user.RoleName,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}

	response.Success(c, http.StatusOK, "get user success", resp)
}
