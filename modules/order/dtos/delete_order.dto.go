package orderdtos

type DeleteOrderDTO struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

