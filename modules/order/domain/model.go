package orderdomain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id         int         `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId     uuid.UUID   `json:"user_id" gorm:"column:user_id;type:varchar(36)"`
	FoodId     int         `json:"food_id" gorm:"column:food_id;"`
	Quantity   int         `json:"quantity" gorm:"column:quantity;default:1"`
	TotalPrice float64     `json:"total_price" gorm:"column:total_price;"`
	Status     int         `json:"status" gorm:"column:status;default:1"`
	CreatedAt  time.Time   `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"column:updated_at;"`
	User       *OrderUser  `json:"user,omitempty" gorm:"-"`
	Food       *OrderFood  `json:"food,omitempty" gorm:"-"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderUser struct {
	Id        uuid.UUID `json:"id" gorm:"column:id;"`
	Email     string    `json:"email" gorm:"column:email;"`
	FirstName string    `json:"first_name" gorm:"column:first_name;"`
	LastName  string    `json:"last_name" gorm:"column:last_name;"`
	Phone     string    `json:"phone" gorm:"column:phone;"`
}

func (OrderUser) TableName() string {
	return "users"
}

type OrderFood struct {
	Id    int     `json:"id" gorm:"column:id;"`
	Name  string  `json:"name" gorm:"column:name;"`
	Price float64 `json:"price" gorm:"column:price;"`
}

func (OrderFood) TableName() string {
	return "foods"
}

