package dependency

import (
	"github.com/alielmi98/go-ecommerce-api/config"
	model "github.com/alielmi98/go-ecommerce-api/domin/models"
	contractRepository "github.com/alielmi98/go-ecommerce-api/domin/repository"
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
