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

var list []GenerateVPNRequest

func GenerateVPN(c *gin.Context) {
	var req GenerateVPNRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "error",
			"config":  "no data",
		})
		return
	}

	list = append(list, req)

	c.JSON(http.StatusOK, gin.H{
		"message": "vpn successfully configured",
		"status":  "ok",
	})
}

func GetAllUsersVpn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"all data": list,
	})
}

func GetUserByID(c *gin.Context) {
	var req GenerateVPNRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user id":  req.UserID,
		"order id": req.OrderID,
		"region":   req.Region,
		"config":   req.Cfg,
	})
}
