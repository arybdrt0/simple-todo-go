package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Todo struct {
    Task string
    Done bool
}

var todos []Todo

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
    })

    router.POST("/add", func(c *gin.Context) {
        task := c.PostForm("task")
        todos = append(todos, Todo{Task: task, Done: false})
        c.Redirect(http.StatusFound, "/")
    })

    router.POST("/complete", func(c *gin.Context) {
    	taskNum, err := strconv.Atoi(c.PostForm("taskNum"))
    	if err != nil || taskNum < 0 || taskNum >= len(todos) {
        	c.Redirect(http.StatusFound, "/")
        	return
    	}

    	todos = append(todos[:taskNum], todos[taskNum+1:]...)
    	c.Redirect(http.StatusFound, "/")
	})	


    router.Run(":8080")
}
