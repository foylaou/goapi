package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TestHandler
// @Summary 測試參數 API
// @Tags Test
// @Param any path string true "任意參數"
// @Success 200 {object} map[string]string
// @Router /test/{any} [get]
func TestHandler(c *gin.Context) {
	message := c.Param("any")
	c.JSON(http.StatusOK, gin.H{"message": message + " " + strconv.Itoa(123)})
}

// PingHandler
// @Summary 健康檢查
// @Tags Health
// @Param any path string true "任意參數"
// @Success 200 {object} map[string]string
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
