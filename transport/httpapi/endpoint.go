package httpapi

import (
	"net/http"
	"errors"

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

func makeStore(s title.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Title

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		dbTarget := c.Query("db")

		res, err := s.Store(req, dbTarget)

		if err != nil {

			if errors.Is(err,  title.ErrNameAndDescriptionRequired) || errors.Is(err, title.ErrInvalidDatabaseTarget) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	    c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func makeGet(s title.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id") 
		dbTarget := c.Query("db")

		result, err := s.Get(id, dbTarget)

		if err != nil {

			if errors.Is(err, title.ErrInvalidDatabaseTarget) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if errors.Is(err, title.ErrTitleNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func makeGetAll(s title.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		dbTarget := c.Query("db")

		result, err := s.GetAll(dbTarget)

		if err != nil {

			if errors.Is(err, title.ErrInvalidDatabaseTarget) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

