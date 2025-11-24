package userrpcserver

import (
	userdomain "vht-go/modules/user/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetUsersRPCRequest struct {
	Ids []string `json:"ids"`
}

func (s *UserRPCServer) GetUsersRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetUsersRPCRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Convert string IDs to UUIDs
		userIds := make([]uuid.UUID, 0, len(req.Ids))
		for _, idStr := range req.Ids {
			userId, err := uuid.Parse(idStr)
			if err != nil {
				continue // Skip invalid IDs
			}
			userIds = append(userIds, userId)
		}

		var users []userdomain.User

		if err := s.db.Where("id IN (?)", userIds).Find(&users).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Mask sensitive data for all users
		for i := range users {
			users[i].Mask()
		}

		c.JSON(200, gin.H{"data": users})
	}

}

