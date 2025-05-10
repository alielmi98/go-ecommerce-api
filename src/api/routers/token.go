package routers

import (
	"github.com/alielmi98/go-ecommerce-api/api/handlers"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/gin-gonic/gin"
)

// TokenRouter ...
func Token(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewTokenHandler(cfg)

	router.POST("/refresh-token", h.RefreshToken)

}
