package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnerb47/go-json-server/pkg/fileLoader"
)

func formatEntities(key string, value map[string]any) ([]any, error) {
	entities, ok := value[key].([]any)
	if !ok {
		return nil, errors.New("Invalid JSON structure")
	}
	return entities, nil
}

func getEntityById(id string, entities []any) (any, int, error) {
	for i, entity := range entities {
		entityMap, ok := entity.(map[string]any)
		if !ok {
			fmt.Println("Can not cast entity to map")
			continue
		}
		idFound, idExists := entityMap["id"]
		if idExists && idFound == id {
			return entityMap, i, nil
		}
	}
	return nil, -1, errors.New("Entity not found")
}

func updateEntity(c *gin.Context, data []any) {
	var newValue any
	id := c.Param("id")

	if err := c.BindJSON((&newValue)); err != nil {
		return
	}

	_, i, err := getEntityById(id, data)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}
	data[i] = newValue
	fmt.Println(data)
	c.IndentedJSON(http.StatusOK, newValue)

}

func addEntity(c *gin.Context, key string, data []any) {
	var entity any

	if err := c.BindJSON((&entity)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Can not bind json"})
	}

	data = append(data, entity)
	err := fileLoader.WriteJson(key, data)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Can not write json"})
	}
	c.IndentedJSON(http.StatusCreated, entity)

}

func getEntity(c *gin.Context, data []any) {
	id := c.Param("id")
	entity, _, err := getEntityById(id, data)
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
	var enpoints []string

	for key := range data {
		enpoints = append(enpoints, key)
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
		router.POST("/"+key, func(c *gin.Context) {
			addEntity(c, key, entities)
		})
		router.PATCH("/"+key+"/:id", func(c *gin.Context) {
			updateEntity(c, entities)
		})
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"enpoints": enpoints})
	})
	return router
}
