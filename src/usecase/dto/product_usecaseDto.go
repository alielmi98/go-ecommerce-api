package dto

type ResponseProduct struct {
	Id           int
	Name         string
	Description  string
	Price        float64
	Stock        int
	Slug         string
	Status       string
	CategoryId   int
	Category     ResponseCategory
	AvrageRating float64
	CountViews   int
	Images       []ResponseProductImage
	Reviews      []ResponseProductReview
	CreatedAt    string
	UpdatedAt    string
}

type CreateProduct struct {
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryId  int
	Slug        string
	Status      string
}

type UpdateProduct struct {
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryId  int
	Slug        string
	Status      string
}
