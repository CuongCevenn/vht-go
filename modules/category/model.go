package categorymodule

import (
	"time"

	"github.com/google/uuid"
)

type Paging struct {
	Page  int `json:"page" form:"page"`
	Limit  int `json:"limit" form:"limit"`
	Total int64 `json:"total"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
}

type Category struct {
	ID          uuid.UUID `json:"id" gorm:"column:id;"`
	Name        string    `json:"name" gorm:"column:name;"`
	Description string    `json:"description" gorm:"column:description;"`
	Status      int       `json:"status" gorm:"column:status;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;"`
	// Icon        string `json:"icon"`
}

func (Category) TableName() string {
	return "categories"
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

func (UpdateCategoryDTO) TableName() string {
	return Category{}.TableName()
}