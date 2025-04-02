package dto

type ResponseProductImage struct {
	ProductImageId int
	ProductId      int
	ImageUrl       string
	IsMain         bool
}

type CreateProductImage struct {
	ProductId int
	ImageUrl  string
	IsMain    bool
}

type UpdateProductImage struct {
	ProductId int
	ImageUrl  string
	IsMain    bool
}
