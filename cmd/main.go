package main

import (
	"log"
	"net/http"

	"github.com/cvrseq/rk-core/internal/db"
	"github.com/cvrseq/rk-core/internal/handlers"
	"github.com/cvrseq/rk-core/internal/repository"
	"github.com/cvrseq/rk-core/internal/service"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Could not connect database")
	}

	router := http.NewServeMux()

	vpnConfigRepo := repository.NewVpnConfigRepository(database)
	vpnConfigService := service.NewVpnConfigDataService(vpnConfigRepo)
	vpnConfigHandler := handlers.NewConfigHandler(vpnConfigService)

	router.HandleFunc("POST /generate", vpnConfigHandler.Generate)
}
