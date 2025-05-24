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

type OrderItemHandler struct {
	usecase *usecase.OrderItemUsecase
}

func NewOrderItemHandler(cfg *config.Config) *OrderItemHandler {
	return &OrderItemHandler{
		usecase: usecase.NewOrderItemUsecase(cfg, dependency.GetOrderItemRepository(cfg)),
	}
}

// CreateOrderItem godoc
// @Summary Create a OrderItem
// @Description Create a Category
// @Tags OrderItem
// @Accept json
// @produces json
// @Param Request body dto.CreateOrderItem true "Create a OrderItem"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.OrderItemResponse} "OrderItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/order-items/ [post]
// @Security AuthBearer
func (h *OrderItemHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateOrderItem, dto.ToOrderItemResponse, h.usecase.Create)
}

// UpdateOrderItem godoc
// @Summary Update a OrderItem
// @Description Update a OrderItem
// @Tags OrderItem
// @Accept json
// @produces json
// @Param id path int true "OrderItem ID"
// @Param Request body dto.UpdateOrderItem true "Update a OrderItem"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.OrderItemResponse} "OrderItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/order-items/{id} [put]
// @Security AuthBearer
func (h *OrderItemHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateOrderItem, dto.ToOrderItemResponse, h.usecase.Update)
}

// DeleteOrderItem godoc
// @Summary Delete a OrderItem
// @Description Delete a OrderItem
// @Tags OrderItem
// @Accept json
// @produces json
// @Param id path int true "OrderItem ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.OrderItemResponse} "OrderItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/order-items/{id} [delete]
// @Security AuthBearer
func (h *OrderItemHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetOrderItem godoc
// @Summary Get a OrderItemBy ID
// @Description Get a OrderItem
// @Tags OrderItem
// @Accept json
// @produces json
// @Param id path int true "OrderItem ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.OrderItemResponse} "OrderItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/order-items/{id} [get]
// @Security AuthBearer
func (h *OrderItemHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToOrderItemResponse, h.usecase.GetById)
}

// GetOrderItems By Filter  godoc
// @Summary Get a OrderItem By Filter
// @Description Get a OrderItem By Filter
// @Tags OrderItem
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Get a OrderItem By Filter"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.OrderItemResponse} "OrderItem response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/shop/order-items/get-by-filter [post]
// @Security AuthBearer
func (h *OrderItemHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToOrderItemResponse, h.usecase.GetByFilter)
}
