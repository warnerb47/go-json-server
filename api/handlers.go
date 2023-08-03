package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnerb47/go-json-server/storage"
)

func handleUpdateEntity(c *gin.Context, key string) {
	var newValue any
	id := c.Param("id")
	if err := c.BindJSON((&newValue)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "can not bind json"})
		return
	}

	jsonData := storage.LoadJson()
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
	if err := storage.WriteJson(key, data); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Can not write json"})
		return
	}
	c.IndentedJSON(http.StatusOK, newValue)

}

func handleAddEntity(c *gin.Context, key string) {
	var entity any
	if err := c.BindJSON((&entity)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Can not bind json"})
		return
	}

	jsonData := storage.LoadJson()
	data, ok := formatEntities(key, jsonData)
	if ok != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid JSON structure"})
		return
	}

	data = append(data, entity)
	err := storage.WriteJson(key, data)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Can not write json"})
		return
	}
	c.IndentedJSON(http.StatusCreated, entity)

}

func handleGetEntity(c *gin.Context, key string) {
	id := c.Param("id")
	jsonData := storage.LoadJson()
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

func handleGetEntities(c *gin.Context, key string) {
	jsonData := storage.LoadJson()
	data, ok := formatEntities(key, jsonData)
	if ok != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid JSON structure"})
		return
	}
	c.JSON(http.StatusOK, data)
}

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
