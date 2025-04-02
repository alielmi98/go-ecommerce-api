package dto

// Create
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryId  int     `json:"category_id" binding:"required"`
	Status      string  `json:"status" binding:"required"`
	Slug        string  `json:"slug"`
}

type CreateProductImageRequest struct {
	ProductId int    `json:"product_id" binding:"required"`
	ImageUrl  string `json:"image_url" binding:"required"`
	IsMain    bool   `json:"is_main"`
}

type CreateProductReviewRequest struct {
	ProductId int    `json:"product_id" binding:"required"`
	Rating    int    `json:"rating" binding:"required"`
	Comment   string `json:"comment"`
	UserId    int    `json:"user_id" binding:"required"`
}

// Update
type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryId  int     `json:"category_id"`
	Status      string  `json:"status"`
	Slug        string  `json:"slug"`
}

type UpdateProductImageRequest struct {
	ProductId int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	IsMain    bool   `json:"is_main"`
}

type UpdateProductReviewRequest struct {
	ProductId int    `json:"product_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	UserId    int    `json:"user_id"`
}

// Response
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

type ProductImageResponse struct {
	ProductId int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	IsMain    bool   `json:"is_main"`
}

type ProductReviewResponse struct {
	ProductId int    `json:"product_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	UserId    int    `json:"user_id"`
}
