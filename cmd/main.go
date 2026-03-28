package main

import (
	"log"

	"github.com/gin-gonic/gin"

	delivery "rs/internal/delivery/http"
)

func main() {
	router := gin.Default()

	router.POST("/vpn", delivery.GenerateVPN)
	router.GET("/vpn/all", delivery.GetAllUsersVpn)

	log.Println("Vpn control plane started work...")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
