package dto

import "github.com/alielmi98/go-ecommerce-api/usecase/dto"

type CreateProductImageRequest struct {
	ProductId int    `json:"product_id" binding:"required"`
	ImageUrl  string `json:"image_url" binding:"required"`
	IsMain    bool   `json:"is_main"`
}

type UpdateProductImageRequest struct {
	ProductId int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	IsMain    bool   `json:"is_main"`
}

type ProductImageResponse struct {
	Id        int    `json:"id"` // This field is not used in the request, but it's included for consistency with the response
	ProductId int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	IsMain    bool   `json:"is_main"`
}

func ToProductImageResponse(from dto.ResponseProductImage) ProductImageResponse {
	return ProductImageResponse{
		Id:        from.ProductImageId,
		ProductId: from.ProductId,
		ImageUrl:  from.ImageUrl,
		IsMain:    from.IsMain,
	}
}
func ToCreateProductImage(from CreateProductImageRequest) dto.CreateProductImage {
	return dto.CreateProductImage{
		ProductId: from.ProductId,
		ImageUrl:  from.ImageUrl,
		IsMain:    from.IsMain,
	}
}
func ToUpdateProductImage(from dto.UpdateProductImage) dto.UpdateProductImage {
	return dto.UpdateProductImage{
		ProductId: from.ProductId,
		ImageUrl:  from.ImageUrl,
		IsMain:    from.IsMain,
	}
}
