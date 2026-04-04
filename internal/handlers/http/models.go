package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenerateVPNRequest struct {
	UserID  string `json:"user_id"`
	OrderID string `json:"order_id"`
	Region  string `json:"region"`
	Cfg     string `json:"config"`
}

