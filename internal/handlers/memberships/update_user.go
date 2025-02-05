package memberships

import (
	"log"
	"net/http"

	"github.com/IndraNurfa/fastcampus/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	log.Println(userID)
	err := h.membershipSvc.UpdateUser(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
