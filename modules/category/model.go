package categorymodule

import "time"

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Icon        string `json:"icon"`
}

type CreateCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *int    `json:"status"`
}