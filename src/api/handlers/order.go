package handlers

import (
	"github.com/alielmi98/go-ecommerce-api/api/dto"
	_ "github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/dependency"
	_ "github.com/alielmi98/go-ecommerce-api/domain/filter"

	"github.com/alielmi98/go-ecommerce-api/usecase"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(cfg *config.Config) *OrderHandler {
	return &OrderHandler{
		usecase: usecase.NewOrderUsecase(cfg, dependency.GetOrderRepository(cfg)),
	}
}

// CreateOrder godoc
// @Summary Create a Order
// @Description Create a Category
// @Tags Order
// @Accept json
// @produces json
// @Param Request body dto.CreateOrder true "Create a Order"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.OrderResponse} "Order response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/orders/ [post]
// @Security AuthBearer
func (h *OrderHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateOrder, dto.ToOrderResponse, h.usecase.Create)
}

// UpdateOrder godoc
// @Summary Update a Order
// @Description Update a Order
// @Tags Order
// @Accept json
// @produces json
// @Param id path int true "Order ID"
// @Param Request body dto.UpdateOrder true "Update a Order"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.OrderResponse} "Order response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/orders/{id} [put]
// @Security AuthBearer
func (h *OrderHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateOrder, dto.ToOrderResponse, h.usecase.Update)
}

// DeleteOrder godoc
// @Summary Delete a Order
// @Description Delete a Order
// @Tags Order
// @Accept json
// @produces json
// @Param id path int true "Order ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.OrderResponse} "Order response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/orders/{id} [delete]
// @Security AuthBearer
func (h *OrderHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetOrder godoc
// @Summary Get a OrderBy ID
// @Description Get a Order
// @Tags Order
// @Accept json
// @produces json
// @Param id path int true "Order ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.OrderResponse} "Order response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/orders/{id} [get]
// @Security AuthBearer
func (h *OrderHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToOrderResponse, h.usecase.GetById)
}

// GetOrders By Filter  godoc
// @Summary Get a Order By Filter
// @Description Get a Order By Filter
// @Tags Order
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Get a Order By Filter"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.OrderResponse} "Order response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/orders/get-by-filter [post]
// @Security AuthBearer
func (h *OrderHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToOrderResponse, h.usecase.GetByFilter)
}
