package route

import (
	"blogspot/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoute(db *gorm.DB) *gin.Engine {
	r:=gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db",db)
	})
	r.GET("/post",controller.GetPostAll)
	r.POST("/post",controller.CreatePostData)
	r.PATCH("/post/:id",controller.UpdateDataId)
	r.GET("/post/:id",controller.GetById)
	r.DELETE("/post/:id",controller.DeletebyID)

	return r
}