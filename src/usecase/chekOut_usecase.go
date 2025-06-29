package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/common"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type CheckOutUsecase struct {
	cartRepo      repository.CartRepository
	orderRepo     repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
	productRepo   repository.ProductRepository
	cfg           *config.Config
}

func NewCheckOutUsecase(cfg *config.Config, cartRepo repository.CartRepository, orderRepo repository.OrderRepository, orderItemRepo repository.OrderItemRepository,
	productRepo repository.ProductRepository) *CheckOutUsecase {
	return &CheckOutUsecase{
		cartRepo:      cartRepo,
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		productRepo:   productRepo,
		cfg:           cfg,
	}
}

// Create a new order from the cart
func (u *CheckOutUsecase) CheckOut(ctx context.Context) (dto.ResponseOrder, error) {
	// Extract the User ID from the request
	userId := int(ctx.Value(constants.UserIdKey).(float64))

	// Get the cart by user ID
	cart, err := u.GetCartByUserId(ctx, userId)
	if err != nil {
		return dto.ResponseOrder{}, err
	}

	// Create a new order from the cart
	order, err := u.CreateOrderFromCart(userId, cart)
	if err != nil {
		return dto.ResponseOrder{}, err
	}
	// Save the order
	savedOrder, err := u.orderRepo.Create(nil, order)
	if err != nil {
		return dto.ResponseOrder{}, err
	}
	// Create order items from the cart items
	if err := u.createOrderItems(savedOrder.Id, cart.CartItems); err != nil {
		return dto.ResponseOrder{}, err
	}

	return common.TypeConverter[dto.ResponseOrder](savedOrder)
}

func (u *CheckOutUsecase) createOrderItems(orderId int, cartItems []model.CartItem) error {
	for _, item := range cartItems {
		// Check product availability
		if !u.productRepo.CheckProductAvailability(item.ProductId, item.Quantity) {
			return &service_errors.ServiceError{EndUserMessage: service_errors.ErrItemsUnavailable}
		}
		orderItem := model.OrderItem{
			OrderId:   orderId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
		if _, err := u.orderItemRepo.Create(nil, orderItem); err != nil {
			return err
		}
	}
	return nil
}

func (u *CheckOutUsecase) GetCartByUserId(ctx context.Context, userId int) (*model.Cart, error) {
	cart, err := u.cartRepo.GetCartByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	if cart == nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound} // Cart not found
	}

	if len(cart.CartItems) == 0 {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.CartIsEmpty} // Cart is empty
	}
	return cart, nil
}

// CreateOrderFromCart creates an order from the user's cart
func (u *CheckOutUsecase) CreateOrderFromCart(userId int, cart *model.Cart) (model.Order, error) {
	// Calculate total items and total price
	totalItems := 0
	totalPrice := 0.0
	for _, item := range cart.CartItems {
		totalItems += item.Quantity
		totalPrice += float64(item.Quantity) * item.UnitPrice
	}

	// Create a new order from the cart
	order := model.Order{
		UserId:     userId,
		Status:     "Pending", // Set initial status
		TotalItems: totalItems,
		TotalPrice: totalPrice,
	}
	return order, nil

}
