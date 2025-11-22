package categorycontroller

import (
	categoryservice "vht-go/modules/category/service"

	"github.com/gin-gonic/gin"
)

type HTTPCategoryController struct {
	svc *categoryservice.CategoryService
}

func NewHTTPCategoryController(svc *categoryservice.CategoryService) *HTTPCategoryController {
	return &HTTPCategoryController{svc: svc}
}

func (ctrl *HTTPCategoryController) SetupRoutes(v1 *gin.RouterGroup) {
	catGroup := v1.Group("/categories")
	catGroup.POST("", ctrl.CreateCategoryAPI())
	catGroup.GET("/:id", ctrl.GetCategoryByIdAPI())
	catGroup.GET("", ctrl.GetListCategoriesAPI())
	catGroup.PATCH("/:id", ctrl.UpdateCategoryAPI())
	catGroup.DELETE("/:id", ctrl.DeleteCategory())
}