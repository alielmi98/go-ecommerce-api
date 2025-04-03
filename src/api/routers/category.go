package routers

import (
	"github.com/alielmi98/go-ecommerce-api/api/handlers"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/gin-gonic/gin"
)

func Category(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCategoryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)

}
