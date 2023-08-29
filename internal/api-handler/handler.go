package apihandler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	segmentNameParam = "segment_slug"
	userIdParam      = "user_id"
)

type Handler struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func New(router *gin.Engine, db *gorm.DB) (*Handler, error) {
	handler := &Handler{Router: router, DB: db}

	v1 := router.Group("/api/v1")
	{
		segments := v1.Group("/segments")
		{
			segments.POST("/:"+segmentNameParam, handler.CreateSegment)
			segments.DELETE("/:"+segmentNameParam, handler.DeleteSegment)

			userOperations := segments.Group("/users")
			{
				userOperations.GET("/:"+userIdParam, handler.GetUserActiveSegments)
				userOperations.PUT("/:"+userIdParam, handler.UpdateUserActiveSegments)
			}
		}
	}

	return handler, nil
}

func (h *Handler) Start() error {
	if err := h.Router.Run(); err != nil {
		return err
	}
	return nil
}
