package usecase

import (
	"context"
	"fmt"

	"github.com/alielmi98/go-ecommerce-api/config"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/infra/payment"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type PaymentUsecase struct {
	base           *BaseUsecase[model.Payment, dto.CreatePayment, dto.UpdatePayment, dto.ResponsePayment]
	OrderRepo      repository.OrderRepository
	cfg            *config.Config
	paymentGateway payment.Zarinpal
}

func NewPaymentUsecase(cfg *config.Config, repository repository.PaymentRepository, orderRepo repository.OrderRepository, paymentGateway *payment.Zarinpal) *PaymentUsecase {
	return &PaymentUsecase{
		base:           NewBaseUsecase[model.Payment, dto.CreatePayment, dto.UpdatePayment, dto.ResponsePayment](cfg, repository),
		OrderRepo:      orderRepo,
		cfg:            cfg,
		paymentGateway: *paymentGateway,
	}
}

// Create a new payment
func (u *PaymentUsecase) CreatePaymentUrl(ctx context.Context, req dto.CreatePayment) (dto.ResponsePayment, error) {
	Order, err := u.OrderRepo.GetById(ctx, req.OrderId)
	if err != nil {
		return dto.ResponsePayment{}, err
	}
	paymentUrl, err := u.paymentGateway.PaymentRequest(Order.TotalPrice, ("Payment for Order " + fmt.Sprint(Order.Id)))
	if err != nil {
		return dto.ResponsePayment{}, err
	}

	return dto.ResponsePayment{PaymentUrl: paymentUrl}, nil
}
