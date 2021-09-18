package middleware

import (

	// "os"

	"net/http"

	// routes "github.com/cavdy-play/go_db/routes"
	"github.com/gin-gonic/gin"
)

// book := Mid.Book{Name: "Java", Id: "HSA118785"}

// db.Select("Name", "Id").Create(&book)

func Api(router *gin.Engine) {
	router.GET("/", Welcome)
	router.POST("/", PostWelcome)
	router.POST("/postbook", PostBook)
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func PostWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "leo fuck devon",
	})
}

func PostBook(c *gin.Context) {
	db.AutoMigrate(&User{})
	var user User
	c.BindJSON(&user)
	c.JSON(200, gin.H{"status": "董劭威"}) // Your custom response here

	// book := Book{Name: "Java", Id: "HSA118785"}

	new_user := User{Name: user.Name, Id: user.Id}
	db.Select("Name", "Id").Create(&new_user)

}
