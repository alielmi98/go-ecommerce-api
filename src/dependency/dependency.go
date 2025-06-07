package dependency

import (
	"github.com/alielmi98/go-ecommerce-api/config"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	contractRepository "github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	infraRepository "github.com/alielmi98/go-ecommerce-api/infra/db/repository"
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
	return infraRepository.NewBaseRepository[model.Product](preloads)
}

func GetProductImageRepository(cfg *config.Config) contractRepository.ProductImageRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Image"}, {Entity: "Product"}}
	return infraRepository.NewBaseRepository[model.ProductImage](preloads)
}

func GetProductReviewRepository(cfg *config.Config) contractRepository.ProductReviewRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Product"}}
	return infraRepository.NewBaseRepository[model.ProductReview](preloads)
}

func GetCartRepository(cfg *config.Config) contractRepository.CartRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "CartItems"}}
	return infraRepository.NewCartRepository(preloads)
}

func GetCartItemRepository(cfg *config.Config) (contractRepository.CartItemRepository, contractRepository.CartRepository) {
	// Preloads for CartItemRepository
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Product"}}
	cartItemRepo := infraRepository.NewCartItemRepository(preloads)
	// Preloads for CartRepository
	var cartPreloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "CartItems"}}
	cartRepo := infraRepository.NewCartRepository(cartPreloads)
	return cartItemRepo, cartRepo
}

func GetOrderRepository(cfg *config.Config) contractRepository.OrderRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "OrderItems"}}
	return infraRepository.NewBaseRepository[model.Order](preloads)
}

func GetOrderItemRepository(cfg *config.Config) contractRepository.OrderItemRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Product"}}
	return infraRepository.NewBaseRepository[model.OrderItem](preloads)
}
