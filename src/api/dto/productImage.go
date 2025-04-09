package dto

import "github.com/alielmi98/go-ecommerce-api/usecase/dto"

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
