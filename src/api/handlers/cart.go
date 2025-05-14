package handlers

import (
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/api/dto"
	"github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/dependency"
	"github.com/alielmi98/go-ecommerce-api/usecase"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase *usecase.CartUsecase
}

func NewCartHandler(cfg *config.Config) *CartHandler {
	return &CartHandler{
		usecase: usecase.NewCartUsecase(cfg, dependency.GetCartRepository(cfg)),
	}
}

// CreateCart godoc
// @Summary Create a Cart
// @Description Create a Category
// @Tags Cart
// @Accept json
// @produces json
// @Param Request body dto.CreateCart true "Create a Cart"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CartResponse} "Cart response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/carts/ [post]
// @Security AuthBearer
func (h *CartHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCart, dto.ToCartResponse, h.usecase.Create)
}

// UpdateCart godoc
// @Summary Update a Cart
// @Description Update a Cart
// @Tags Cart
// @Accept json
// @produces json
// @Param id path int true "Cart ID"
// @Param Request body dto.UpdateCart true "Update a Cart"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CartResponse} "Cart response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/carts/{id} [put]
// @Security AuthBearer
func (h *CartHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCart, dto.ToCartResponse, h.usecase.Update)
}

// DeleteCart godoc
// @Summary Delete a Cart
// @Description Delete a Cart
// @Tags Cart
// @Accept json
// @produces json
// @Param id path int true "Cart ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CartResponse} "Cart response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/carts/{id} [delete]
// @Security AuthBearer
func (h *CartHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCartById godoc
// @Summary Get a Cart by ID
// @Description Get a Cart by ID
// @Tags Cart
// @Accept json
// @produces json
// @Param id path int true "Cart ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CartResponse} "Cart response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/carts/{id} [get]
// @Security AuthBearer
func (h *CartHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCartResponse, h.usecase.GetById)
}

// GetCartByUserId godoc
// @Summary Get a User Cart
// @Description Get a User Cart By User ID
// @Tags Cart
// @Accept json
// @produces json
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CartResponse} "Cart response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/carts/ [get]
// @Security AuthBearer
func (h *CartHandler) GetByUserId(c *gin.Context) {
	userId := int(c.Value(constants.UserIdKey).(float64))
	if userId == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}
	cart, err := h.usecase.GetByUserId(c, userId)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := dto.ToCartResponse(cart)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
