package handlers

import (
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/api/dto"
	"github.com/alielmi98/go-ecommerce-api/api/helper"
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
	paymentRepo, orderRepo, productRepo := dependency.GetPaymentRepository(cfg)
	paymentGateway := dependency.GetPaymentGateway(cfg)
	return &PaymentHandler{
		usecase: usecase.NewPaymentUsecase(cfg, paymentRepo, orderRepo, productRepo, paymentGateway),
	}
}

// CreatePaymentUrl godoc
// @Summary Create a Payment URL
// @Description Create a Payment URL for a specific order
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body dto.CreatePaymentUrl true "CreatePaymentRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ResponsePaymentUrl} "Order response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/payments/create-url [post]
// @Security AuthBearer
func (h *PaymentHandler) CreatePaymentUrl(c *gin.Context) {
	Create(c, dto.ToCreatePaymentUrl, dto.ToResponsePaymentUrl, h.usecase.CreatePaymentUrl)

}

func (h *PaymentHandler) VerifyPayment(c *gin.Context) {
	// Extract the authority from the query parameters
	authority := c.Query("Authority")
	if authority == "" {
		c.JSON(400, gin.H{"error": "Authority is required"})
		return
	}
	// Call the usecase to verify the payment
	response, err := h.usecase.VerifyPayment(c, authority)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// Create Payment godoc
// @Summary Create a Payment
// @Description Create a new payment record
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body dto.CreatePayment true "CreatePaymentRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PaymentVerificationResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/payments/create [post]
func (h *PaymentHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreatePayment, dto.ToResponsePayment, h.usecase.Create)
}

// Update Payment godoc
// @Summary Update a Payment
// @Description Update an existing payment record
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Param request body dto.UpdatePayment true "UpdatePaymentRequest"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ResponsePayment}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/payments/{id} [put]
func (h *PaymentHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdatePayment, dto.ToResponsePayment, h.usecase.Update)
}

// Delete Payment godoc
// @Summary Delete a Payment
// @Description Delete an existing payment record
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Success 204 {object} helper.BaseHttpResponse "No Content"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/payments/{id} [delete]
func (h *PaymentHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// Get Payment godoc
// @Summary Get a Payment by ID
// @Description Get an existing payment record by ID
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ResponsePayment} "Payment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/payments/{id} [get]
func (h *PaymentHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToResponsePayment, h.usecase.GetById)
}

// Get By Filter godoc
// @Summary Get Payments by Filter
// @Description Get a list of payment records based on filter criteria
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body filter.PaginationInputWithFilter true "Filter criteria"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.ResponsePayment]} "Payments response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/payments [get]
func (h *PaymentHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToResponsePayment, h.usecase.GetByFilter)
}
