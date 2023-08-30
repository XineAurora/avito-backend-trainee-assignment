package apihandler

import (
	"net/http"
	"strconv"

	"github.com/XineAurora/user-segmentation/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserActiveSegments(c *gin.Context) {
	userIdString := c.Param(userIdParam)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": "user_id is invalid",
		})
		return
	}

	segments, err := models.GetUserActiveSegments(h.DB, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, segments)
}

func (h *Handler) UpdateUserActiveSegments(c *gin.Context) {
	userIdString := c.Param(userIdParam)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": "user_id is invalid",
		})
		return
	}

	var body struct {
		Include []string
		Exclude []string
	}
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": "request body is invalid",
		})
		return
	}

	err = models.UpdateUserActiveSegments(h.DB, userId, body.Include, body.Exclude)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) VerifyUserId(c *gin.Context) {
	//verify userid or discard request
	c.Next()
}
