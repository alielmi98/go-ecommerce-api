package handlers

import (
	"github.com/alielmi98/go-ecommerce-api/api/dto"
	_ "github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/dependency"
	_ "github.com/alielmi98/go-ecommerce-api/domin/filter"
	"github.com/alielmi98/go-ecommerce-api/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase *usecase.CategoryUsecase
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
// @Router /v1/shop/categories/ [post]
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
// @Router /v1/shop/categories/{id} [put]
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
// @Router /v1/shop/categories/{id} [delete]
// @Security AuthBearer
func (h *CategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCategory godoc
// @Summary Get a Category
// @Description Get a Category
// @Tags Category
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CategoryResponse} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/categories/{id} [get]
// @Security AuthBearer
func (h *CategoryHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCategoryResponse, h.usecase.GetById)
}

// GetCategories godoc
// @Summary Get Categories
// @Description Get Categories
// @Tags Category
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CategoryResponse]} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/categories/get-by-filter [post]
// @Security AuthBearer
func (h *CategoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCategoryResponse, h.usecase.GetByFilter)
}
