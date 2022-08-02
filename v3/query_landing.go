package v3

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IdpPingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func IdpLanding(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": LANDING,
	})
}

func IdpVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "v0.0.1",
	})
}
