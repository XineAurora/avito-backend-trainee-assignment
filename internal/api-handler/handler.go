package apihandler

import (
	"github.com/XineAurora/user-segmentation/internal/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
			segments.POST("/:segment_slug", handler.CreateSegment)
			segments.DELETE("/:segment_slug", handler.DeleteSegment)

			userOperations := segments.Group("/users")
			{
				userOperations.GET("/:user_id", handler.GetUserActiveSegments)
				userOperations.POST("/:user_id", handler.UpdateUserActiveSegments)
			}
		}
	}

	return handler, nil
}
