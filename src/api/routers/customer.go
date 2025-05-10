package routers

import (
	handlers "github.com/alielmi98/go-ecommerce-api/api/handlers"
	"github.com/alielmi98/go-ecommerce-api/api/middlewares"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/events"
	"github.com/gin-gonic/gin"
)

func CartItem(r *gin.RouterGroup, cfg *config.Config, dispatcher *events.EventDispatcher) {
	h := handlers.NewCartItemHandler(cfg)

	dispatcher.Register("ProductPriceChanged", h.HandleProductPriceChanged)
	r.POST("/", middlewares.Authentication(cfg), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), h.Delete)
	r.GET("/:id", h.GetById)
}
