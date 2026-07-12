package main

import (
	"log"

	"faro/backend/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	internal.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
