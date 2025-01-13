package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/IndraNurfa/fastcampus/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.UserActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postID := c.Param("postID")
	postIDInt, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("post ID not valid").Error(),
		})
		return
	}
	userID := c.GetInt64("userID")
	err = h.postSvc.UpsertUserActivity(ctx, postIDInt, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
