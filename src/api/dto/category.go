package dto

import "github.com/alielmi98/go-ecommerce-api/usecase/dto"

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToCategoryResponse(from dto.ResponseCategory) CategoryResponse {
	return CategoryResponse{
		Name:        from.Name,
		Description: from.Description,
	}
}

func ToCreateCategory(from dto.CreateCategory) dto.CreateCategory {
	return dto.CreateCategory{
		Name:        from.Name,
		Description: from.Description,
	}
}

func ToUpdateCategory(from dto.UpdateCategory) dto.UpdateCategory {
	return dto.UpdateCategory{
		Name:        from.Name,
		Description: from.Description,
	}
}
