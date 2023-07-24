package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEntities(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Configure(data map[string]interface{}) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	for key := range data {
		value := data[key]
		router.GET("/"+key, func(c *gin.Context) {
			getEntities(c, value)
		})
	}

	return router
}
