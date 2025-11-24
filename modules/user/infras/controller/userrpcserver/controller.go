package userrpcserver

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRPCServer struct {
	db *gorm.DB
}

func NewUserRPCServer(db *gorm.DB) *UserRPCServer {
	return &UserRPCServer{db: db}
}

func (s *UserRPCServer) SetupRouter(v1 *gin.RouterGroup) {
	userRPCGroup := v1.Group("/rpc/users")
	userRPCGroup.POST("/get-user", s.GetUserRPCAPI())
	userRPCGroup.POST("/get-users", s.GetUsersRPCAPI())
}

