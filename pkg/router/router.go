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

func updateEntity(c *gin.Context, key string) {
	var newValue any
	id := c.Param("id")
	if err := c.BindJSON((&newValue)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "can not bind json"})
		return
	}

	jsonData := fileLoader.LoadJson()
	data, ok := formatEntities(key, jsonData)
	if ok != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid JSON structure"})
		return
	}

	_, i, err := getEntityById(id, data)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}
	data[i] = newValue
	if err := fileLoader.WriteJson(key, data); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Can not write json"})
		return
	}
	c.IndentedJSON(http.StatusOK, newValue)

}

func addEntity(c *gin.Context, key string) {
	var entity any
	if err := c.BindJSON((&entity)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Can not bind json"})
		return
	}

	jsonData := fileLoader.LoadJson()
	data, ok := formatEntities(key, jsonData)
	if ok != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid JSON structure"})
		return
	}

	data = append(data, entity)
	err := fileLoader.WriteJson(key, data)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Can not write json"})
		return
	}
	c.IndentedJSON(http.StatusCreated, entity)

}

func getEntity(c *gin.Context, key string) {
	id := c.Param("id")
	jsonData := fileLoader.LoadJson()
	data, ok := formatEntities(key, jsonData)
	if ok != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid JSON structure"})
		return
	}
	entity, _, err := getEntityById(id, data)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, entity)
}

func getEntities(c *gin.Context, key string) {
	jsonData := fileLoader.LoadJson()
	data, ok := formatEntities(key, jsonData)
	if ok != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid JSON structure"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func Configure() *gin.Engine {
	data := fileLoader.LoadJson()
	router := gin.Default()
	var enpoints []string

	for key := range data {
		enpoints = append(enpoints, key)
		router.GET("/"+key, func(c *gin.Context) {
			getEntities(c, key)
		})
		router.GET("/"+key+"/:id", func(c *gin.Context) {
			getEntity(c, key)
		})
		router.POST("/"+key, func(c *gin.Context) {
			addEntity(c, key)
		})
		router.PATCH("/"+key+"/:id", func(c *gin.Context) {
			updateEntity(c, key)
		})
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"enpoints": enpoints})
	})
	return router
}
