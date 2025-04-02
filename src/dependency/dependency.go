package dependency

import (
	"github.com/alielmi98/go-ecommerce-api/config"
	model "github.com/alielmi98/go-ecommerce-api/domin/models"
	contractRepository "github.com/alielmi98/go-ecommerce-api/domin/repository"
	infraRepository "github.com/alielmi98/go-ecommerce-api/infra/db/repository"
)

func GetCategoryRepository(cfg *config.Config) contractRepository.CategoryRepository {
	return infraRepository.NewBaseRepository[model.Category]()
}
