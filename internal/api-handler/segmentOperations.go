package apihandler

import (
	"net/http"

	"github.com/XineAurora/user-segmentation/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSegment(c *gin.Context) {
	segmentSlug := c.Param(segmentNameParam)

	segment, err := models.CreateSegment(h.DB, segmentSlug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, segment)
}

func (h *Handler) DeleteSegment(c *gin.Context) {
	segmentSlug := c.Param(segmentNameParam)

	err := models.DeleteSegment(h.DB, segmentSlug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":        http.StatusBadRequest,
			"description": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
