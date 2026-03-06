package httpapi

import (
	"github.com/gin-gonic/gin"
)

func New(endpoints Endpoints) *gin.Engine {
	app := gin.Default()

	app.POST("/titles", endpoints.Store)
	app.GET("/titles", endpoints.GetAll)
	app.GET("/titles/:id", endpoints.Get)

	return app
}
