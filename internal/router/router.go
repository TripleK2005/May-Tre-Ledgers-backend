package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"may-tre-ledger-be/internal/core/config"
	"may-tre-ledger-be/internal/core/response"

	"may-tre-ledger-be/internal/middleware"
	"may-tre-ledger-be/internal/modules/auth"
	"may-tre-ledger-be/internal/modules/role"
	"may-tre-ledger-be/internal/modules/token"
	"may-tre-ledger-be/internal/modules/user"
)

type Router struct {
	Config *config.Config
	DB     *pgxpool.Pool
}

func Setup(cfg *config.Config, db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(c.Request.Context()); err != nil {
			response.Error(c, http.StatusServiceUnavailable, "database unavailable")
			return
		}

		response.Success(c, http.StatusOK, "ok", gin.H{})
	})

	userRepo := user.NewRepository(db)
	roleRepo := role.NewRepository(db)
	tokenRepo := token.NewRepository(db)

	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	authService := auth.NewService(userRepo, roleRepo, tokenRepo, cfg)
	authHandler := auth.NewHandler(authService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/refresh", authHandler.RefreshToken)
		authGroup.POST("/logout", authHandler.Logout)
	}

	userGroup := r.Group("/users")
	userGroup.Use(middleware.Auth(cfg.JWTSecret))
	{
		userGroup.Use(middleware.RequireRoles("ADMIN", "STAFF"))
		userGroup.GET("/:id", userHandler.GetByID)
	}

	return r
}
