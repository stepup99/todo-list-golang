package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos map[int]Todo

func init() {
	todos = map[int]Todo{
		1: {ID: 1, Title: "Learn GO", Status: "incomplete"},
		2: {ID: 2, Title: "Read Gin Documentation", Status: "incomplete"},
	}
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	fmt.Println("starting server ................")
	router.Run()

}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}
