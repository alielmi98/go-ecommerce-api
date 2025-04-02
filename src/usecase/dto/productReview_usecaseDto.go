package dto

type ResponseProductReview struct {
	Id        int
	ProductId int
	UserId    int
	Rating    int
	Comment   string
}

type CreateProductReview struct {
	ProductId int
	Rating    int
	Comment   string
	UserId    int
}

type UpdateProductReview struct {
	ProductId int
	Rating    int
	Comment   string
	UserId    int
}
