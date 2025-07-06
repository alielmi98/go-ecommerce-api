package dto

type ResponseProduct struct {
	Id            int
	Name          string
	Description   string
	Price         float64
	Stock         int
	Slug          string
	Status        string
	CategoryId    int
	Category      ResponseCategory
	AverageRating float64
	CountViews    int
	Images        []ResponseProductImage
	Reviews       []ResponseProductReview
	CreatedAt     string
	UpdatedAt     string
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

// Product Category
type CreateCategory struct {
	Name        string
	Description string
}

type UpdateCategory struct {
	Name        string
	Description string
}

type ResponseCategory struct {
	Name        string
	Description string
}

// Product Image
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

// Product Review
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
