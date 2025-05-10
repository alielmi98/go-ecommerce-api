package handlers

import (
	"github.com/alielmi98/go-ecommerce-api/api/dto"
	_ "github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/dependency"
	_ "github.com/alielmi98/go-ecommerce-api/domain/filter"
	"github.com/alielmi98/go-ecommerce-api/events"
	"github.com/alielmi98/go-ecommerce-api/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase *usecase.ProductUsecase
}

func NewProductHandler(cfg *config.Config, dipatcher *events.EventDispatcher) *ProductHandler {
	return &ProductHandler{
		usecase: usecase.NewProductUsecase(cfg, dependency.GetProductRepository(cfg), dipatcher),
	}
}

// CreateProduct godoc
// @Summary Create a Product
// @Description Create a Product
// @Tags Product
// @Accept json
// @produces json
// @Param Request body dto.CreateProductRequest true "Create a Product"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ProductResponse} "Product response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/products/ [post]
// @Security AuthBearer
func (h *ProductHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateProduct, dto.ToProductResponse, h.usecase.Create)
}

// UpdateProduct godoc
// @Summary Update a Product
// @Description Update a Product
// @Tags Product
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateProductRequest true "Update a Product"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductResponse} "Product response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/products/{id} [put]
// @Security AuthBearer
func (h *ProductHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateProduct, dto.ToProductResponse, h.usecase.Update)
}

// DeleteProduct godoc
// @Summary Delete a Product
// @Description Delete a Product
// @Tags Product
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductResponse} "Product response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/products/{id} [delete]
// @Security AuthBearer
func (h *ProductHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetProduct godoc
// @Summary Get a Product
// @Description Get a Product
// @Tags Product
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ProductResponse} "Product response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/products/{id} [get]
// @Security AuthBearer
func (h *ProductHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToProductResponse, h.usecase.GetById)
}

// Getproducts godoc
// @Summary Get products
// @Description Get products
// @Tags Product
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.ProductResponse]} "Product response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/shop/products/get-by-filter [post]
// @Security AuthBearer
func (h *ProductHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToProductResponse, h.usecase.GetByFilter)
}
