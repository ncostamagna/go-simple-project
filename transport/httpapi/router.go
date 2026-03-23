package httpapi

import (
	"github.com/gin-gonic/gin"
)

func New(e Endpoints) *gin.Engine {
	app := gin.Default()

	app.POST("/titles", e.Store)
	app.GET("/titles", e.GetAll)
	app.GET("/titles/:id", e.Get)

	return app
}
