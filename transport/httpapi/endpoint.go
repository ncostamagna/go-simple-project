package httpapi

import (
	"net/http"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ncostamagna/go-simple-project/domain"
	"github.com/ncostamagna/go-simple-project/internal/title"
)

type (
	Endpoints struct {
		Get    gin.HandlerFunc
		GetAll gin.HandlerFunc
		Store  gin.HandlerFunc
	}
)

func MakePostsEndpoints(s title.Service) Endpoints {
	return Endpoints{
		Get:    makeGet(s),
		GetAll: makeGetAll(s),
		Store:  makeStore(s),
	}
}

func makeGet(s title.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id") 

		result, err := s.Get(id)
		if err != nil {
			if errors.Is(err, title.ErrNoTitlesInDatabase) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			if errors.Is(err, title.ErrTitleNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func makeStore(s title.Service) gin.HandlerFunc {
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