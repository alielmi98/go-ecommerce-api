package routers

import (
	"github.com/alielmi98/go-ecommerce-api/api/handlers"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/gin-gonic/gin"
)

// AccountRouter ...
func Account(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewAccountHandler(cfg)

	router.POST("/register-by-username", h.Create)
	router.POST("/login-by-username", h.LoginByUsername)

}
