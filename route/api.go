package route

import (
	"crud/controllers"
	"crud/controllers/news"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", controllers.Home)

	routeNews := r.Group("/news")
	routeNews.GET("/", news.FetchNews)
	routeNews.GET("/:id", news.FindNews)
	routeNews.POST("/create", news.CreateNews)
	routeNews.POST("/:id/update", news.UpdateNews)
	routeNews.GET("/:id/delete", news.DeleteNews)

	return r
}
