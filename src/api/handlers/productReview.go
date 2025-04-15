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

type ProductReviewHandler struct {
	usecase *usecase.ProductReviewUsecase
}

func NewProductReviewHandler(cfg *config.Config) *ProductReviewHandler {
	return &ProductReviewHandler{
		usecase: usecase.NewProductReviewUsecase(cfg, dependency.GetProductReviewRepository(cfg)),
	}
}

// CreateProductReview godoc
// @Summary Create a ProductReview
// @Description Create a ProductReview
// @Tags ProductReview
// @Accept json
// @produces json
// @Param Request body dto.CreateProductReviewRequest true "Create a ProductReview"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ProductReviewResponse} "ProductReview response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-reviews/ [post]
// @Security AuthBearer
func (h *ProductReviewHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateProductReview, dto.ToProductReviewResponse, h.usecase.Create)
}

// UpdateProductReview godoc
// @Summary Update a ProductReview
// @Description Update a ProductReview
// @Tags ProductReview
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateProductReviewRequest true "Update a ProductReview"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductReviewResponse} "ProductReview response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-reviews/{id} [put]
// @Security AuthBearer
func (h *ProductReviewHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateProductReview, dto.ToProductReviewResponse, h.usecase.Update)
}

// DeleteProductReview godoc
// @Summary Delete a ProductReview
// @Description Delete a ProductReview
// @Tags ProductReview
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductReviewResponse} "ProductReview response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-reviews/{id} [delete]
// @Security AuthBearer
func (h *ProductReviewHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetProductReview godoc
// @Summary Get a ProductReview
// @Description Get a ProductReview
// @Tags ProductReview
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductReviewResponse} "ProductReview response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-reviews/{id} [get]
// @Security AuthBearer
func (h *ProductReviewHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToProductReviewResponse, h.usecase.GetById)
}

// GetProductReviews godoc
// @Summary Get ProductReviews
// @Description Get ProductReviews
// @Tags ProductReview
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.ProductReviewResponse]} "ProductReview response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/product-reviews/get-by-filter [post]
// @Security AuthBearer
func (h *ProductReviewHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToProductReviewResponse, h.usecase.GetByFilter)
}
