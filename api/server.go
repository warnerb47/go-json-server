package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnerb47/go-json-server/storage"
)

func Start(url string, filePath string) {
	storage.Setpath(filePath)
	configure().Run(url)
}

func configure() *gin.Engine {

	data := storage.LoadJson()
	router := gin.Default()
	var enpoints []string

	for k := range data {
		enpoints = append(enpoints, k)
		var key = k
		router.GET("/"+key, func(c *gin.Context) {
			handleGetEntities(c, key)
		})
		router.GET("/"+key+"/:id", func(c *gin.Context) {
			handleGetEntity(c, key)
		})
		router.POST("/"+key, func(c *gin.Context) {
			handleAddEntity(c, key)
		})
		router.PATCH("/"+key+"/:id", func(c *gin.Context) {
			handleUpdateEntity(c, key)
		})
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"enpoints": enpoints})
	})
	return router
}
