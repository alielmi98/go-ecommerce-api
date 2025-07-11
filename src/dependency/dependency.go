package dependency

import (
	"github.com/alielmi98/go-ecommerce-api/config"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	contractRepository "github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	infraRepository "github.com/alielmi98/go-ecommerce-api/infra/db/repository"
	"github.com/alielmi98/go-ecommerce-api/infra/payment"
)

func GetUserRepository(cfg *config.Config) contractRepository.UserRepository {
	return infraRepository.NewUserRepository()
}
func GetCategoryRepository(cfg *config.Config) contractRepository.CategoryRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Products"}}
	return infraRepository.NewBaseRepository[model.Category](preloads)
}

func GetFileRepository(cfg *config.Config) contractRepository.FileRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.File](preloads)
}

func GetProductRepository(cfg *config.Config) contractRepository.ProductRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Images"}, {Entity: "Reviews"}, {Entity: "Category"}}
	return infraRepository.NewProductRepository(preloads)
}

func GetProductImageRepository(cfg *config.Config) contractRepository.ProductImageRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Image"}, {Entity: "Product"}}
	return infraRepository.NewBaseRepository[model.ProductImage](preloads)
}

func GetProductReviewRepository(cfg *config.Config) (contractRepository.ProductReviewRepository, contractRepository.ProductRepository) {
	var productReviewPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Product"}}
	var productRepoPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Images"}, {Entity: "Reviews"}, {Entity: "Category"}}
	return infraRepository.NewBaseRepository[model.ProductReview](productReviewPreloads), infraRepository.NewProductRepository(productRepoPreloads)
}

func GetCartRepository(cfg *config.Config) contractRepository.CartRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "CartItems"}}
	return infraRepository.NewCartRepository(preloads)
}

func GetCartItemRepository(cfg *config.Config) (contractRepository.CartItemRepository, contractRepository.CartRepository, contractRepository.ProductRepository) {
	// Preloads for CartItemRepository
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Product"}}
	cartItemRepo := infraRepository.NewCartItemRepository(preloads)
	// Preloads for CartRepository
	var cartPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "CartItems"}}
	cartRepo := infraRepository.NewCartRepository(cartPreloads)
	// Preloads for ProductRepository
	var productPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Images"}, {Entity: "Reviews"}, {Entity: "Category"}}
	productRepo := infraRepository.NewProductRepository(productPreloads)
	return cartItemRepo, cartRepo, productRepo
}

func GetOrderRepository(cfg *config.Config) contractRepository.OrderRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "OrderItems"}}
	return infraRepository.NewBaseRepository[model.Order](preloads)
}

func GetOrderItemRepository(cfg *config.Config) contractRepository.OrderItemRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Product"}}
	return infraRepository.NewBaseRepository[model.OrderItem](preloads)
}

func GetCheckOutRepository(cfg *config.Config) (contractRepository.CartRepository, contractRepository.OrderRepository, contractRepository.OrderItemRepository,
	contractRepository.ProductRepository, contractRepository.CartItemRepository) {
	// Preloads for CartRepository
	var cartPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "CartItems"}}
	cartRepo := infraRepository.NewCartRepository(cartPreloads)
	// Preloads for CartItemRepository
	var cartItemPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Product"}}
	cartItemRepo := infraRepository.NewCartItemRepository(cartItemPreloads)
	// Preloads for OrderRepository
	var orderPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "OrderItems"}}
	orderRepo := infraRepository.NewBaseRepository[model.Order](orderPreloads)
	// Preloads for OrderItemRepository
	var orderItemPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Product"}}
	orderItemRepo := infraRepository.NewBaseRepository[model.OrderItem](orderItemPreloads)
	// Preloads for ProductRepository
	var productPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Images"}, {Entity: "Reviews"}, {Entity: "Category"}}
	productRepo := infraRepository.NewProductRepository(productPreloads)

	return cartRepo, orderRepo, orderItemRepo, productRepo, cartItemRepo

}

func GetPaymentRepository(cfg *config.Config) (contractRepository.PaymentRepository, contractRepository.OrderRepository, contractRepository.ProductRepository) {
	var preloads []db.PreloadEntity = []db.PreloadEntity{}
	var orderPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "OrderItems"}}
	var productPreloads []db.PreloadEntity = []db.PreloadEntity{}
	return infraRepository.NewPaymentRepository(preloads), infraRepository.NewBaseRepository[model.Order](orderPreloads), infraRepository.NewProductRepository(productPreloads)
}

func GetPaymentGateway(cfg *config.Config) *payment.Zarinpal {
	return payment.NewZarinpalGateway(cfg)
}
