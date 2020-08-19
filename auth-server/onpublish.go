package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const STREAMING_KEY = os.Getenv("STREAMING_KEY")

func onpublish(c *gin.Context) {
	key := c.Query("key")
	if key == STREAMING_KEY {
		c.String(http.StatusOK, "Good to go")
	} else {
		c.String(http.StatusUnauthorized, "Invalid straming key")
	}
}
