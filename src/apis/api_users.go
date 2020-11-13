package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for items table
func UsersIndex(c *gin.Context) {
	var lis []models.User

	db, _ := c.Get("db")

	conn := db.(gorm.DB)
	// Migrate the schema
	//conn.AutoMigrate(&models.User{})

	conn.Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func UsersCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	//var d User
	//d := models.User{Email: c.PostForm("email")} //, PasswordHash: c.PostForm("password_hash")
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	d.PasswordHash = string("123")
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}

func UsersGet(c *gin.Context) {

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.User
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)
}

func UsersUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.User
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	d.Email = c.PostForm("email")
	d.PasswordHash = c.PostForm("password_hash")
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}

func UsersDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.User

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
