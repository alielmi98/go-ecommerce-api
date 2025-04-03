package dependency

import (
	"github.com/alielmi98/go-ecommerce-api/config"
	model "github.com/alielmi98/go-ecommerce-api/domin/models"
	contractRepository "github.com/alielmi98/go-ecommerce-api/domin/repository"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	infraRepository "github.com/alielmi98/go-ecommerce-api/infra/db/repository"
)

func GetCategoryRepository(cfg *config.Config) contractRepository.CategoryRepository {
	var preloads []db.PreloadEntity = []db.PreloadEntity{{Entity: "Products"}}
	return infraRepository.NewBaseRepository[model.Category](preloads)
}
