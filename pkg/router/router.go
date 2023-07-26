package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func formatEntities(key string, value map[string]any) ([]any, error) {
	entities, ok := value[key].([]any)
	if !ok {
		return nil, errors.New("Invalid JSON structure")
	}
	return entities, nil
}

func getEntityById(id string, entities []any) (any, error) {
	for _, entity := range entities {
		entityMap, ok := entity.(map[string]any)
		if !ok {
			fmt.Println("Can not cast entity to map")
			continue
		}
		idFound, idExists := entityMap["id"]
		if idExists && idFound == id {
			return entityMap, nil
		}
	}
	return nil, errors.New("Entity not found")
}

func getEntity(c *gin.Context, data []any) {
	id := c.Param("id")
	entity, err := getEntityById(id, data)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, entity)
}

func getEntities(c *gin.Context, data []any) {
	c.JSON(http.StatusOK, data)
}

func Configure(data map[string]any) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	for key := range data {
		entities, ok := formatEntities(key, data)
		if ok != nil {
			fmt.Printf("Invalid JSON structure: %v\n", key)
		}
		router.GET("/"+key, func(c *gin.Context) {
			getEntities(c, entities)
		})
		router.GET("/"+key+"/:id", func(c *gin.Context) {
			getEntity(c, entities)
		})
	}
	return router
}
