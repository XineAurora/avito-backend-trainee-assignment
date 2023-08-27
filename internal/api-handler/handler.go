package apihandler

import (
	"github.com/XineAurora/user-segmentation/internal/db"
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

func New() (*Handler, error) {
	router := gin.Default()
	db, err := db.New()
	if err != nil {
		//TODO: add description to error
		return nil, err
	}
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
				userOperations.POST("/:"+userIdParam, handler.UpdateUserActiveSegments)
			}
		}
	}

	return handler, nil
}
