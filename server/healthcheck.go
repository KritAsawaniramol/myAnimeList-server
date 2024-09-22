package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g *ginServer) healthCheckService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"app": "myanimelist", "status": "OK"})
}
