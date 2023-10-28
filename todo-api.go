package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var todos = []todo{
	{ID: "1", Task: "Eat", Done: false},
	{ID: "2", Task: "Sleep", Done: true},
	{ID: "3", Task: "Code", Done: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8080")
}
