package dto

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
