package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TODO struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []TODO{
	{ID: "1", Item: "Clean room", Completed: true},
	{ID: "2", Item: "Read Boook", Completed: false},
	{ID: "3", Item: "Anything", Completed: true},
}

func gettodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addtodos(context *gin.Context) {
	var newTodo TODO
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func gettodobyid(id string) (*TODO, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("to do not found")
}

func gettodo(context *gin.Context) {
	id := context.Param("id")
	TODO, err := gettodobyid(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To do not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, TODO)
}

func toogletodostatus(context *gin.Context) {
	id := context.Param("id")
	TODO, err := gettodobyid(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To do not found"})
		return
	}
	TODO.Completed = !TODO.Completed

	context.IndentedJSON(http.StatusOK, TODO)
}

func main() {
	router := gin.Default()
	router.GET("/todos", gettodos)
	router.POST("/todos", addtodos)
	router.GET("/todos/:id", gettodo)
	router.PATCH("/todos/:id", toogletodostatus)
	router.Run("localhost:9090")
}
