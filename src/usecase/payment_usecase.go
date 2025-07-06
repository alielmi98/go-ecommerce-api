package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/infra/payment"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type PaymentUsecase struct {
	base           *BaseUsecase[model.Payment, dto.CreatePayment, dto.UpdatePayment, dto.ResponsePayment]
	paymentRepo    repository.PaymentRepository
	OrderRepo      repository.OrderRepository
	ProductRepo    repository.ProductRepository
	cfg            *config.Config
	paymentGateway payment.Zarinpal
}

func NewPaymentUsecase(cfg *config.Config, repository repository.PaymentRepository, orderRepo repository.OrderRepository, productRepo repository.ProductRepository, paymentGateway *payment.Zarinpal) *PaymentUsecase {
	return &PaymentUsecase{
		base:           NewBaseUsecase[model.Payment, dto.CreatePayment, dto.UpdatePayment, dto.ResponsePayment](cfg, repository),
		paymentRepo:    repository,
		OrderRepo:      orderRepo,
		ProductRepo:    productRepo,
		cfg:            cfg,
		paymentGateway: *paymentGateway,
	}
}

// Create a new payment
func (u *PaymentUsecase) CreatePaymentUrl(ctx context.Context, req dto.CreatePaymentUrl) (dto.ResponsePaymentUrl, error) {
	order, err := u.OrderRepo.GetById(ctx, req.OrderId)
	if err != nil {
		return dto.ResponsePaymentUrl{}, err
	}
	// Validate the order before proceeding with payment
	// This function checks if the order is valid for payment
	if err := validateOrderForPayment(order, u.ProductRepo); err != nil {
		return dto.ResponsePaymentUrl{}, err
	}
	paymentUrl, authorityId, err := u.paymentGateway.PaymentRequest(order.TotalPrice, ("Payment for Order " + fmt.Sprint(order.Id)))
	if err != nil {
		return dto.ResponsePaymentUrl{}, err
	}

	// Create a new payment record
	payment := dto.CreatePayment{
		Amount:      order.TotalPrice,
		Status:      "pending",
		AuthorityId: authorityId,
		RefId:       0, // This will be updated after payment verification
		UserId:      order.UserId,
		OrderId:     order.Id,
	}
	_, err = u.base.Create(ctx, payment)
	if err != nil {
		return dto.ResponsePaymentUrl{}, err
	}

	return dto.ResponsePaymentUrl{PaymentUrl: paymentUrl + authorityId}, nil
}

// Verify the payment
func (u *PaymentUsecase) VerifyPayment(ctx context.Context, authority string) (dto.PaymentVerificationResponse, error) {
	// Get Payment by authority
	payment, err := u.paymentRepo.GetByAuthority(ctx, authority)
	if err != nil {
		return dto.PaymentVerificationResponse{}, err
	}

	// Call the payment gateway's verify method
	refId, err := u.paymentGateway.PaymentVerify(authority, payment.Amount)
	if err != nil {
		// Update payment status to failed
		_, updateErr := u.base.Update(ctx, payment.Id, dto.UpdatePayment{
			Status: "failed",
			RefId:  refId,
		})
		if updateErr != nil {
			return dto.PaymentVerificationResponse{}, updateErr
		}
		return dto.PaymentVerificationResponse{
			PaymentId: payment.Id,
			Status:    "failed: " + err.Error(),
		}, nil
	}

	// Update payment status to verified
	_, err = u.base.Update(ctx, payment.Id, dto.UpdatePayment{
		Status: "verified",
		RefId:  refId,
	})
	if err != nil {
		return dto.PaymentVerificationResponse{}, err
	}

	// Update order status to paid
	_, err = u.OrderRepo.Update(ctx, payment.OrderId, model.Order{
		Status: "paid",
	})
	if err != nil {
		log.Printf("Order status update failed for payment %d: %v", payment.Id, err)
		return dto.PaymentVerificationResponse{
			PaymentId: payment.Id,
			Status:    "payment verified but order update failed",
			RefId:     refId,
		}, nil
	}

	// Deduct the product stock asynchronously
	go func() {
		if err := u.DeductProductStock(payment.OrderId); err != nil {
			log.Printf("caller:%s  Level:%s Msg:%v", constants.Postgres, constants.UseCase, err)
		}
	}()

	// Return the response with payment ID, status, and reference ID
	return dto.PaymentVerificationResponse{
		PaymentId: payment.Id,
		Status:    "verified",
		RefId:     refId,
	}, nil
}

// Deduct the product stock after payment verification
func (u *PaymentUsecase) DeductProductStock(orderId int) error {
	order, err := u.OrderRepo.GetById(context.Background(), orderId)
	if err != nil {
		return err
	}

	for _, item := range order.OrderItems {
		err := u.ProductRepo.DeductProductStock(item.ProductId, item.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateOrderForPayment(order model.Order, productRepo repository.ProductRepository) error {
	if order.Status == "paid" {
		return fmt.Errorf("order is already paid")
	}
	if order.Status == "canceled" || order.Status == "expired" {
		return fmt.Errorf("order is not active")
	}
	if order.TotalPrice <= 0 {
		return fmt.Errorf("invalid order total")
	}
	for _, item := range order.OrderItems {
		if !productRepo.CheckProductAvailability(item.ProductId, item.Quantity) {
			return fmt.Errorf("insufficient stock for product %d", item.ProductId)
		}
	}
	return nil
}

// Create a new payment
func (u *PaymentUsecase) Create(ctx context.Context, req dto.CreatePayment) (dto.ResponsePayment, error) {
	// Call the base Create method
	return u.base.Create(ctx, req)
}

func (u *PaymentUsecase) Update(ctx context.Context, id int, req dto.UpdatePayment) (dto.ResponsePayment, error) {
	// Call the base Update method
	return u.base.Update(ctx, id, req)
}

func (u *PaymentUsecase) Delete(ctx context.Context, id int) error {
	// Call the base Delete method
	return u.base.Delete(ctx, id)
}
func (u *PaymentUsecase) GetById(ctx context.Context, id int) (dto.ResponsePayment, error) {
	// Call the base GetById method
	return u.base.GetById(ctx, id)
}

func (u *PaymentUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponsePayment], error) {
	return u.base.GetByFilter(ctx, req)
}
