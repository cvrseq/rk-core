package main

import (
	"log"

	"github.com/gin-gonic/gin"

	handlers "rs/internal/handlers/http"
)

func main() {
	router := gin.Default()

	router.POST("/vpn", handlers.GenerateVPN)
	router.GET("/vpn/all", handlers.GetAllUsersVpn)
	router.GET("/vpn/check", handlers.GetUserByID)

	log.Println("Vpn control plane started work...")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
