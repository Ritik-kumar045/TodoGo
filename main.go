package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Name string
	Done bool
}

var tasks []Task

func loadTasks() {
	data, _ := os.ReadFile("tasks.json")
	json.Unmarshal(data, &tasks)
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile("tasks.json", data, 0644)
}

func main() {
	loadTasks()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"tasks": tasks})
	})

	r.POST("/add", func(c *gin.Context) {
		task := c.PostForm("task")
		tasks = append(tasks, Task{Name: task, Done: false})
		saveTasks()
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/done/:id", func(c *gin.Context) {
		id := c.Param("id")
		var i int
		fmt.Sscanf(id, "%d", &i)
		if i >= 0 && i < len(tasks) {
			tasks[i].Done = true
			saveTasks()
		}
		c.Redirect(http.StatusFound, "/")
	})

	r.Run(":8081")
}
