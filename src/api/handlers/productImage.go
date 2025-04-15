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

type ProductImageHandler struct {
	usecase *usecase.ProductImageUsecase
}

func NewProductImageHandler(cfg *config.Config) *ProductImageHandler {
	return &ProductImageHandler{
		usecase: usecase.NewProductImageUsecase(cfg, dependency.GetProductImageRepository(cfg)),
	}
}

// CreateProductImage godoc
// @Summary Create a ProductImage
// @Description Create a ProductImage
// @Tags ProductImage
// @Accept json
// @produces json
// @Param Request body dto.CreateProductImageRequest true "Create a ProductImage"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ProductImageResponse} "ProductImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-images/ [post]
// @Security AuthBearer
func (h *ProductImageHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateProductImage, dto.ToProductImageResponse, h.usecase.Create)
}

// UpdateProductImage godoc
// @Summary Update a ProductImage
// @Description Update a ProductImage
// @Tags ProductImage
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateProductImageRequest true "Update a ProductImage"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductImageResponse} "ProductImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-images/{id} [put]
// @Security AuthBearer
func (h *ProductImageHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateProductImage, dto.ToProductImageResponse, h.usecase.Update)
}

// DeleteProductImage godoc
// @Summary Delete a ProductImage
// @Description Delete a ProductImage
// @Tags ProductImage
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductImageResponse} "ProductImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-images/{id} [delete]
// @Security AuthBearer
func (h *ProductImageHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetProductImage godoc
// @Summary Get a ProductImage
// @Description Get a ProductImage
// @Tags ProductImage
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductImageResponse} "ProductImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-images/{id} [get]
// @Security AuthBearer
func (h *ProductImageHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToProductImageResponse, h.usecase.GetById)
}

// GetProductImages godoc
// @Summary Get ProductImages
// @Description Get ProductImages
// @Tags ProductImage
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.ProductImageResponse]} "ProductImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-images/get-by-filter [post]
// @Security AuthBearer
func (h *ProductImageHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToProductImageResponse, h.usecase.GetByFilter)
}
