package userrpcserver

import (
	userdomain "vht-go/modules/user/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GetUserRPCRequest struct {
	Id string `json:"id" binding:"required"`
}

func (s *UserRPCServer) GetUserRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetUserRPCRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		userId, err := uuid.Parse(req.Id)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid user id format"})
			return
		}

		var user userdomain.User

		if err := s.db.First(&user, "id = ?", userId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{"error": userdomain.ErrUserNotFound})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Mask sensitive data
		user.Mask()

		c.JSON(200, gin.H{"data": &user})
	}
}

