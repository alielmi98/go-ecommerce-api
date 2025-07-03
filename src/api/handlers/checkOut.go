package handlers

import (
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/api/dto"
	"github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/dependency"
	_ "github.com/alielmi98/go-ecommerce-api/domain/filter"
	"github.com/alielmi98/go-ecommerce-api/usecase"
	"github.com/gin-gonic/gin"
)

type CheckOutHandler struct {
	usecase *usecase.CheckOutUsecase
}

func NewCheckOutHandler(cfg *config.Config) *CheckOutHandler {
	cartRepo, orderRepo, orderItemRepo, productRepo, cartItemRepo := dependency.GetCheckOutRepository(cfg)
	return &CheckOutHandler{
		usecase: usecase.NewCheckOutUsecase(cfg, cartRepo, orderRepo, orderItemRepo, productRepo, cartItemRepo),
	}
}

// CreateOrderFromCart godoc
// @Summary Create an order from the cart
// @Description Create a new order from the items in the user's cart
// @Tags checkout
// @Accept json
// @Produce json
// @Param Request body dto.CheckOutRequest true "Check out request"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.OrderResponse} "Order response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/checkout/ [post]
// @Security AuthBearer
func (h *CheckOutHandler) CheckOut(c *gin.Context) {
	// Bind the request body to CheckOutRequest
	var checkOutRequest dto.CheckOutRequest
	if err := c.ShouldBindJSON(&checkOutRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	// map http request body to usecase input
	usecaseInput := dto.ToCheckOutRequest(checkOutRequest)
	order, err := h.usecase.CheckOut(c, usecaseInput)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	response := dto.ToOrderResponse(order)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}
