package handlers

import (
	"github.com/alielmi98/go-ecommerce-api/api/dto"
	_ "github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/dependency"
	"github.com/alielmi98/go-ecommerce-api/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase *usecase.CatecoryUsecase
}

func NewCategoryHandler(cfg *config.Config) *CategoryHandler {
	return &CategoryHandler{
		usecase: usecase.NewCategoryUsecase(cfg, dependency.GetCategoryRepository(cfg)),
	}
}

// CreateCategory godoc
// @Summary Create a Category
// @Description Create a Category
// @Tags Category
// @Accept json
// @produces json
// @Param Request body dto.CreateCategoryRequest true "Create a Category"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CategoryResponse} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/category/ [post]
// @Security AuthBearer
func (h *CategoryHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCategory, dto.ToCategoryResponse, h.usecase.Create)
}

// UpdateCategory godoc
// @Summary Update a Category
// @Description Update a Category
// @Tags Category
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCategoryRequest true "Update a Category"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CategoryResponse} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/category/{id} [put]
// @Security AuthBearer
func (h *CategoryHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCategory, dto.ToCategoryResponse, h.usecase.Update)
}

// DeleteCategory godoc
// @Summary Delete a Category
// @Description Delete a Category
// @Tags Category
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CategoryResponse} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/category/{id} [delete]
// @Security AuthBearer
func (h *CategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}
