package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncostamagna/go-simple-project/domain"
)

type (
	Endpoints struct {
		Get    gin.HandlerFunc
		GetAll gin.HandlerFunc
		Store  gin.HandlerFunc
	}
)

func MakePostsEndpoints() Endpoints {
	return Endpoints{
		Get:    makeGet(),
		GetAll: makeGetAll(),
		Store:  makeStore(),
	}
}

func makeGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := domain.Title{
			Name: "exaple1",
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func makeGetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := []domain.Title{{
			Name: "exaple1",
		},
			{
				Name: "exaple2",
			},
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func makeStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := domain.Title{
			Name:        "exaple1",
			Description: "desc1",
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
