package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
)

//CRUD for items table
func PersonsIndex(c *gin.Context) {
	s := models.Person{Name: "Sean", Age: "12"}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hola person: " + s.Name,
	})
}
