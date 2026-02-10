package httpapi

import (
	"github.com/gin-gonic/gin"
)

func New(endpoints Endpoints) *gin.Engine {
	app := gin.Default()
	app.GET("/titles/:id", endpoints.Get)
	app.GET("/titles", endpoints.GetAll)
	app.POST("/titles", endpoints.Store)

	return app
}
