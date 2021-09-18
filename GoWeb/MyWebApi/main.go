package main

import (
	Mid "MyWebApi/middleware"
	// "fmt"

	// routes "github.com/cavdy-play/go_db/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func main() {
	// Init Router
	router := gin.Default()
	router.Use(cors.Default())

	// Route Handlers / Endpoints
	Mid.Api(router)

	Mid.Connect()

	log.Fatal(router.Run(":8080"))
}
