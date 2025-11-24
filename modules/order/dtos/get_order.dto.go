package orderdtos

type GetOrderDTO struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

