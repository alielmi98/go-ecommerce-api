package handlers

import (
	"github.com/alielmi98/go-ecommerce-api/api/dto"
	_ "github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/dependency"
	"github.com/alielmi98/go-ecommerce-api/usecase"

	"github.com/gin-gonic/gin"
)

type CartItemHandler struct {
	usecase *usecase.CartItemUsecase
}

func NewCartItemHandler(cfg *config.Config) *CartItemHandler {
	return &CartItemHandler{
		usecase: usecase.NewCartItemUsecase(cfg, dependency.GetCartItemRepository(cfg)),
	}
}

// CreateCartItem godoc
// @Summary Create a CartItem
// @Description Create a Category
// @Tags CartItem
// @Accept json
// @produces json
// @Param Request body dto.CreateCartItem true "Create a CartItem"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CartItemResponse} "CartItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/cart-items/ [post]
// @Security AuthBearer
func (h *CartItemHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCartItem, dto.ToCartItemResponse, h.usecase.Create)
}
