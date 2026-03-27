package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenerateVPNRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`

	Region string `json:"region" binding:"required,len=2"`
}

func GenerateVPN(c *gin.Context) {
	var req GenerateVPNRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "vpn successfully configured",
		"status":  "ok",
		"data":    req,
	})
}
