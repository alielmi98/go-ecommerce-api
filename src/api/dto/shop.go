package dto

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryId  int     `json:"category_id" binding:"required"`
	Status      string  `json:"status" binding:"required"`
	Slug        string  `json:"slug"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
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
