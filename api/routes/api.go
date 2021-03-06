package routes

import (
	"github.com/labstack/echo"
	"github.com/takatori/youtube-manager/api/middlewares"
	"github.com/takatori/youtube-manager/api/web/api"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo(), middlewares.FirebaseAuth())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}
	fg := g.Group("/favorite", middlewares.FirebaseGuard())
	{
		fg.GET("", api.FetchFavoriteVideos())
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
	}
}
