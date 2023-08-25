package apihandler

import "github.com/gin-gonic/gin"

type Handler struct {
	Router *gin.Engine
}

func New() *Handler {
	router := gin.Default()
	handler := &Handler{Router: router}

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

	return handler
}
