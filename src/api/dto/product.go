package dto

import "github.com/alielmi98/go-ecommerce-api/usecase/dto"

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryId  int     `json:"category_id" binding:"required"`
	Status      string  `json:"status" binding:"required"`
	Slug        string  `json:"slug"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryId  int     `json:"category_id"`
	Status      string  `json:"status"`
	Slug        string  `json:"slug"`
}

type ProductResponse struct {
	Name         string                  `json:"name"`
	Description  string                  `json:"description"`
	Price        float64                 `json:"price"`
	Stock        int                     `json:"stock"`
	CategoryId   int                     `json:"category_id"`
	Status       string                  `json:"status"`
	Slug         string                  `json:"slug"`
	AvrageRating float64                 `json:"avrage_rating"`
	CountViews   int                     `json:"count_views"`
	Images       []ProductImageResponse  `json:"images"`
	Reviews      []ProductReviewResponse `json:"reviews"`
	Category     CategoryResponse        `json:"category"`
}

func ToProductResponse(from ProductResponse, images []ProductImageResponse, reviews []ProductReviewResponse, category CategoryResponse) ProductResponse {
	return ProductResponse{
		Name:         from.Name,
		Description:  from.Description,
		Price:        from.Price,
		Stock:        from.Stock,
		CategoryId:   from.CategoryId,
		Status:       from.Status,
		Slug:         from.Slug,
		AvrageRating: from.AvrageRating,
		CountViews:   from.CountViews,
		Images:       images,
		Reviews:      reviews,
		Category:     category,
	}
}

func ToProductCreateRequest(from CreateProductRequest) dto.CreateProduct {
	return dto.CreateProduct{
		Name:        from.Name,
		Description: from.Description,
		Price:       from.Price,
		Stock:       from.Stock,
		CategoryId:  from.CategoryId,
		Status:      from.Status,
		Slug:        from.Slug,
	}
}

func ToProductUpdateRequest(from UpdateProductRequest) dto.UpdateProduct {
	return dto.UpdateProduct{
		Name:        from.Name,
		Description: from.Description,
		Price:       from.Price,
		Stock:       from.Stock,
		CategoryId:  from.CategoryId,
		Status:      from.Status,
		Slug:        from.Slug,
	}
}
