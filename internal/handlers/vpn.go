package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cvrseq/rk-core/internal/models"
)

type VpnConfigService interface {
	GenerateConfig(ctx context.Context, input VpnConfigRequest) (VpnConfigResponse, error)
	GetConfig(ctx context.Context) ([]models.VpnConfigModel, error) // need think it from gorm to local json
}

type VpnConfigHandler struct {
	serv VpnConfigService
}

func NewConfigHandler(serv VpnConfigService) *VpnConfigHandler {
	return &VpnConfigHandler{
		serv: serv,
	}
}

type VpnConfigRequest struct {
	UserID  uint   `json:"user_id"`
	OrderID uint   `json:"order_id"`
	Region  string `json:"region"`
}

type VpnConfigResponse struct {
	UserID  uint   `json:"user_id"`
	OrderID uint   `json:"order_id"`
	Region  string `json:"region"`
	Cfg     string `json:"config"`
}

func (h *VpnConfigHandler) Generate(w http.ResponseWriter, r *http.Request) {
	var cfg VpnConfigRequest

	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req := &VpnConfigRequest{
		UserID:  cfg.UserID,
		OrderID: cfg.OrderID,
		Region:  cfg.Region,
	}

	config, err := h.serv.GenerateConfig(r.Context(), *req)
	if err != nil {
		http.Error(w, "failed to generate config", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(config)
}
