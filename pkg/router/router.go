package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Init project", Completed: false},
	{ID: "2", Item: "Setup environment", Completed: false},
	{ID: "3", Item: "Add versionning", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var todo Todo

	if err := c.BindJSON((&todo)); err != nil {
		return
	}

	todos = append(todos, todo)
	c.IndentedJSON(http.StatusCreated, todo)

}

func getTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func toggleCheck(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)

}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleCheck)
	router.POST("/todos", addTodo)
	router.Run("localhost:3000")
	return router
}
