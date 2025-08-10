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
	cartItemRepo := dependency.GetCartItemRepository(cfg)
	cartRepo := dependency.GetCartRepository(cfg)
	productRepo := dependency.GetProductRepository(cfg)

	return &CartItemHandler{
		// Initialize the CartItemUsecase with the necessary dependencies
		usecase: usecase.NewCartItemUsecase(cfg, cartItemRepo, cartRepo, productRepo),
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

// UpdateCartItem godoc
// @Summary Update a CartItem
// @Description Update a CartItem
// @Tags CartItem
// @Accept json
// @produces json
// @Param id path int true "CartItem ID"
// @Param Request body dto.UpdateCartItem true "Update a CartItem"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CartItemResponse} "CartItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/cart-items/{id} [put]
// @Security AuthBearer
func (h *CartItemHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCartItem, dto.ToCartItemResponse, h.usecase.Update)
}

// DeleteCartItem godoc
// @Summary Delete a CartItem
// @Description Delete a CartItem
// @Tags CartItem
// @Accept json
// @produces json
// @Param id path int true "CartItem ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CartItemResponse} "CartItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/cart-items/{id} [delete]
// @Security AuthBearer
func (h *CartItemHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCartItemById godoc
// @Summary Get a CartItem by ID
// @Description Get a CartItem by ID
// @Tags CartItem
// @Accept json
// @produces json
// @Param id path int true "CartItem ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CartItemResponse} "CartItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/cart-items/{id} [get]
// @Security AuthBearer
func (h *CartItemHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCartItemResponse, h.usecase.GetById)
}
