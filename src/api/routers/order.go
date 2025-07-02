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

func Cart(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCartHandler(cfg)
	// Admin
	r.POST("/", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Delete)
	r.GET("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.GetById)
	// Customer
	r.GET("/", middlewares.Authentication(cfg), h.GetByUserId)

}

func OrderItem(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewOrderItemHandler(cfg)

	// Admin
	r.POST("/", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Delete)
	r.GET("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.GetById)
	r.POST("/get-by-filter", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.GetByFilter)
}

func Order(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewOrderHandler(cfg)

	// Admin
	r.POST("/", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Delete)
	r.GET("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.GetById)
}

func CheckOutHandler(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCheckOutHandler(cfg)

	r.POST("/", middlewares.Authentication(cfg), h.CheckOut)

}

func Payment(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPaymentHandler(cfg)

	r.POST("/", middlewares.Authentication(cfg), h.CreatePaymentUrl)
}
