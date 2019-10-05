package routes

import (
	"github.com/labstack/echo"
	"github.com/takatori/youtube-manager/api/web/api"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/related/:id", api.FetchRelatedVideos())
	}

}
