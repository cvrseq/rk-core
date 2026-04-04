package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	log.Println("Vpn control plane started work...")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
