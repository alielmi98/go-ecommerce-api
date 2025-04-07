package dto

import "github.com/alielmi98/go-ecommerce-api/usecase/dto"

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
