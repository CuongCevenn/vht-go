package usermodule

import (
	"vht-go/middleware"
	usercontroller "vht-go/modules/user/infras/controller"
	userrepository "vht-go/modules/user/infras/repository"
	userservice "vht-go/modules/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserModule(v1 *gin.RouterGroup, db *gorm.DB, jwtComponent userservice.IJWTComponent, middlewareProvider middleware.IMiddlewareProvider) {
	repo := userrepository.NewGORMUserRepository(db)
	registerUserHandler := userservice.NewRegisterUserCommandHandler(repo)
	loginUserHandler := userservice.NewLoginUserCommandHandler(repo, jwtComponent)

	controller := usercontroller.NewHTTPUserController(
		registerUserHandler,
		loginUserHandler,
	)

	controller.SetupRouter(v1, middlewareProvider)
}