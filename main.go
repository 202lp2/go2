package main

import (
	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	s := models.Item{Title: "Sean", Notes: "nnn"}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola "+s.Title,
		})
	})
	r.Run("localhost:8085") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}