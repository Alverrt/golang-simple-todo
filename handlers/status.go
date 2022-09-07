package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pong": "pong!",
	})
}
