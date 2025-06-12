package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/common"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type CheckOutUsecase struct {
	cartRepo      repository.CartRepository
	orderRepo     repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
	cfg           *config.Config
}

func NewCheckOutUsecase(cfg *config.Config, cartRepo repository.CartRepository, orderRepo repository.OrderRepository, orderItemRepo repository.OrderItemRepository) *CheckOutUsecase {
	return &CheckOutUsecase{
		cartRepo:      cartRepo,
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		cfg:           cfg,
	}
}

// Create a new order from the cart
func (u *CheckOutUsecase) CreateOrderFromCart(ctx context.Context) (dto.ResponseOrder, error) {
	// Extract the User ID from the request
	userId := int(ctx.Value(constants.UserIdKey).(float64))
	// Get the cart by user ID
	cart, err := u.cartRepo.GetCartByUserId(nil, userId)
	if err != nil {
		return dto.ResponseOrder{}, err
	}
	if cart == nil {
		return dto.ResponseOrder{}, nil // No cart found for the user
	}

	// Create a new order from the cart
	order := model.Order{
		UserId: userId,
		Status: "Pending", // Set initial status
	}

	// Save the order
	savedOrder, err := u.orderRepo.Create(nil, order)
	if err != nil {
		return dto.ResponseOrder{}, err
	}

	// Create order items from the cart items
	for _, item := range cart.CartItems {
		orderItem := model.OrderItem{
			OrderId:   savedOrder.Id,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
		if _, err := u.orderItemRepo.Create(nil, orderItem); err != nil {
			return dto.ResponseOrder{}, err
		}
	}

	return common.TypeConverter[dto.ResponseOrder](savedOrder)
}
