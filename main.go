package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

type Todo struct {
	CreatedBy string
	Content   string
	isDone    bool
	Status    int
}

func main() {

	todo := make([]Todo, 0, 5)
	t := Todo{
		CreatedBy: "eiji",
		Content:   "早起き",
		Status:    0,
		isDone:    false,
	}
	todo = append(todo, t)
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",map[string]interface{}{
			"todo": todo,
		})
	})

	router.Run()
}
