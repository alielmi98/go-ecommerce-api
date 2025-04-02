package dto

type ResponseProduct struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryID  int
	Category    string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
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
