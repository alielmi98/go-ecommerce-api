package routers

import (
	handlers "github.com/alielmi98/go-ecommerce-api/api/handlers"
	"github.com/alielmi98/go-ecommerce-api/api/middlewares"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/gin-gonic/gin"
)

func Category(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCategoryHandler(cfg)

	r.POST("/", h.Create, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.PUT("/:id", h.Update, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.DELETE("/:id", h.Delete, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}

func Product(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewProductHandler(cfg)

	r.POST("/", h.Create, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.PUT("/:id", h.Update, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.DELETE("/:id", h.Delete, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}

func File(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)

	r.POST("/", h.Create, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.PUT("/:id", h.Update, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.DELETE("/:id", h.Delete, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.GET("/:id", h.GetById, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.POST("/get-by-filter", h.GetByFilter, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
}
func ProductImage(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewProductImageHandler(cfg)

	r.POST("/", h.Create, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.PUT("/:id", h.Update, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.DELETE("/:id", h.Delete, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}

func ProductReview(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewProductReviewHandler(cfg)

	r.POST("/", h.Create, middlewares.Authentication(cfg))
	r.PUT("/:id", h.Update, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.DELETE("/:id", h.Delete, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}
