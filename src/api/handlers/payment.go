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

type PaymentHandler struct {
	usecase *usecase.PaymentUsecase
}

func NewPaymentHandler(cfg *config.Config) *PaymentHandler {
	paymentRepo, orderRepo := dependency.GetPaymentRepository(cfg)
	paymentGateway := dependency.GetPaymentGateway(cfg)
	return &PaymentHandler{
		usecase: usecase.NewPaymentUsecase(cfg, paymentRepo, orderRepo, paymentGateway),
	}
}

// CreatePaymentUrl godoc
// @Summary Create a Payment URL
// @Description Create a Payment URL for a specific order
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body dto.CreatePayment true "CreatePaymentRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PaymentResponse} "Order response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/payments/ [post]
// @Security AuthBearer
func (h *PaymentHandler) CreatePaymentUrl(c *gin.Context) {
	Create(c, dto.ToCreatePayment, dto.ToPaymentResponse, h.usecase.CreatePaymentUrl)

}
