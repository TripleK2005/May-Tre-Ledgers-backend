package middleware

import (
	"net/http"
	"strings"

	"may-tre-ledger-be/internal/core/response"
	"may-tre-ledger-be/internal/utils"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserID = "user_id"
	ContextRole   = "role"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Error(c, http.StatusUnauthorized, "invalid authorization header")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1], secret)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		c.Set(ContextUserID, claims.UserID)
		c.Set(ContextRole, claims.Role)
		c.Next()
	}
}

func RequireRoles(roles ...string) gin.HandlerFunc {
	allowedRoles := make(map[string]struct{}, len(roles))
	for _, role := range roles {
		allowedRoles[strings.ToUpper(role)] = struct{}{}
	}

	return func(c *gin.Context) {
		role, ok := c.Get(ContextRole)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "missing auth context")
			c.Abort()
			return
		}

		roleName, ok := role.(string)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "invalid auth context")
			c.Abort()
			return
		}

		if _, ok := allowedRoles[strings.ToUpper(roleName)]; !ok {
			response.Error(c, http.StatusForbidden, "permission denied")
			c.Abort()
			return
		}

		c.Next()
	}
}
