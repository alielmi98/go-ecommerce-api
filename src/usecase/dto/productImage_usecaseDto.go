package dto

type ResponseProductImage struct {
	Id        int
	ImageId   int
	ProductId int
	ImageUrl  string
	IsMain    bool
}

type CreateProductImage struct {
	ProductId int
	ImageId   int
	IsMain    bool
}

type UpdateProductImage struct {
	ProductId int
	ImageId   int
	IsMain    bool
}
