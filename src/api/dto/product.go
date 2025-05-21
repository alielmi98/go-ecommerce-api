package dto

import (
	"sync"

	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

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
	CreatedAt    string                  `json:"created_at"`
	UpdatedAt    string                  `json:"updated_at"`
}

func ToProductResponse(from dto.ResponseProduct) ProductResponse {
	images := []ProductImageResponse{}
	reviews := []ProductReviewResponse{}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for _, item := range from.Images {
			images = append(images, ToProductImageResponse(item))
		}
		wg.Done()
	}()

	go func() {
		for _, item := range from.Reviews {
			reviews = append(reviews, ToProductReviewResponse(item))
		}
		wg.Done()
	}()

	wg.Wait()

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
		Category:     ToCategoryResponse(from.Category),
	}
}

func ToCreateProduct(from CreateProductRequest) dto.CreateProduct {
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

func ToUpdateProduct(from UpdateProductRequest) dto.UpdateProduct {
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

// Product Category
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToCategoryResponse(from dto.ResponseCategory) CategoryResponse {
	return CategoryResponse{
		Name:        from.Name,
		Description: from.Description,
	}
}

func ToCreateCategory(from dto.CreateCategory) dto.CreateCategory {
	return dto.CreateCategory{
		Name:        from.Name,
		Description: from.Description,
	}
}

func ToUpdateCategory(from dto.UpdateCategory) dto.UpdateCategory {
	return dto.UpdateCategory{
		Name:        from.Name,
		Description: from.Description,
	}
}

// Product Image
type CreateProductImageRequest struct {
	ProductId int  `json:"product_id" binding:"required"`
	ImageId   int  `json:"image_id" binding:"required"`
	IsMain    bool `json:"is_main"`
}

type UpdateProductImageRequest struct {
	ProductId int  `json:"product_id" binding:"required"`
	ImageId   int  `json:"image_id"`
	IsMain    bool `json:"is_main" binding:"required"`
}

type ProductImageResponse struct {
	Id        int  `json:"id"` // This field is not used in the request, but it's included for consistency with the response
	ProductId int  `json:"product_id"`
	ImageId   int  `json:"image_id"`
	IsMain    bool `json:"is_main"`
}

func ToProductImageResponse(from dto.ResponseProductImage) ProductImageResponse {
	return ProductImageResponse{
		Id:        from.Id,
		ProductId: from.ProductId,
		ImageId:   from.ImageId,
		IsMain:    from.IsMain,
	}
}
func ToCreateProductImage(from CreateProductImageRequest) dto.CreateProductImage {
	return dto.CreateProductImage{
		ProductId: from.ProductId,
		ImageId:   from.ImageId,
		IsMain:    from.IsMain,
	}
}
func ToUpdateProductImage(from UpdateProductImageRequest) dto.UpdateProductImage {
	return dto.UpdateProductImage{
		ProductId: from.ProductId,
		ImageId:   from.ImageId,
		IsMain:    from.IsMain,
	}
}

// Product Review
type CreateProductReviewRequest struct {
	ProductId int    `json:"product_id" binding:"required"`
	Rating    int    `json:"rating" binding:"required"`
	Comment   string `json:"comment"`
	UserId    int    `json:"user_id" binding:"required"`
}

type UpdateProductReviewRequest struct {
	ProductId int    `json:"product_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	UserId    int    `json:"user_id"`
}

type ProductReviewResponse struct {
	ProductId int    `json:"product_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	UserId    int    `json:"user_id"`
}

func ToProductReviewResponse(from dto.ResponseProductReview) ProductReviewResponse {
	return ProductReviewResponse{
		ProductId: from.ProductId,
		Rating:    from.Rating,
		Comment:   from.Comment,
		UserId:    from.UserId,
	}
}
func ToCreateProductReview(from CreateProductReviewRequest) dto.CreateProductReview {
	return dto.CreateProductReview{
		ProductId: from.ProductId,
		Rating:    from.Rating,
		Comment:   from.Comment,
		UserId:    from.UserId,
	}
}
func ToUpdateProductReview(from UpdateProductReviewRequest) dto.UpdateProductReview {
	return dto.UpdateProductReview{
		ProductId: from.ProductId,
		Rating:    from.Rating,
		Comment:   from.Comment,
		UserId:    from.UserId,
	}
}
