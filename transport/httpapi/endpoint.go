package httpapi

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ncostamagna/go-simple-project/domain"

	"github.com/google/uuid"
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
		id := c.Param("id") 

		dbMu.Lock()
		defer dbMu.Unlock()

		for _, title := range database {
			if title.ID == id {
				c.JSON(http.StatusOK, gin.H{"data": title})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "title not found"})
	}
}

func makeGetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		dbMu.Lock()
		defer dbMu.Unlock()

		c.JSON(http.StatusOK, gin.H{"data": database})
	}
}

var (
	database []domain.Title
	dbMu sync.Mutex
)

func makeStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Title
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		if req.Name == "" || req.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Description are required"})
				return
		}
		req.ID = uuid.NewString()

		dbMu.Lock()
		database = append(database, req)
		dbMu.Unlock()

	    c.JSON(http.StatusOK, gin.H{"data": req})
	}
}
