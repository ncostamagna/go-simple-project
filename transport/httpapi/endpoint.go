package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncostamagna/go-simple-project/domain"
	"github.com/ncostamagna/go-simple-project/service"
)

type (
	Endpoints struct {
		Get    gin.HandlerFunc
		GetAll gin.HandlerFunc
		Store  gin.HandlerFunc
	}
)

func MakePostsEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Get:    makeGet(s),
		GetAll: makeGetAll(s),
		Store:  makeStore(s),
	}
}

func makeGet(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id") 

		title, err := s.Get(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "title not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": title})

	}
}

func makeGetAll(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		titles := s.GetAll()

		c.JSON(http.StatusOK, gin.H{"data": titles})
	}
}



func makeStore(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Title
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		res, err := s.Store(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	    c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
