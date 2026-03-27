package main

import (
	"github.com/gin-gonic/gin"
	"log"
	
	delivery "rs/internal/delivery/http"
)

func main() {
	router := gin.Default()

	router.POST("/generate", delivery.GenerateVPN)
	
	log.Println("Vpn control plane started work...")
	
	if err := router.Run(":8080"); err != nil { 
		log.Fatalf("Server error: %v", err)
	}
}
