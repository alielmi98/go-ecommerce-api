package routers

import (
	handlers "github.com/alielmi98/go-ecommerce-api/api/handlers"
	"github.com/alielmi98/go-ecommerce-api/api/middlewares"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/events"
	"github.com/gin-gonic/gin"
)

func Category(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCategoryHandler(cfg)

	r.POST("/", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}

func Product(r *gin.RouterGroup, cfg *config.Config, dispatcher *events.EventDispatcher) {
	h := handlers.NewProductHandler(cfg, dispatcher)

	r.POST("/", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}

func File(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)

	r.POST("/", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Delete)
	r.GET("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.GetById)
	r.POST("/get-by-filter", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.GetByFilter)
}
func ProductImage(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewProductImageHandler(cfg)

	r.POST("/", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}

func ProductReview(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewProductReviewHandler(cfg)

	r.POST("/", middlewares.Authentication(cfg), h.Create)
	r.PUT("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Update)
	r.DELETE("/:id", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}), h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}
