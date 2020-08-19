package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var STREAMING_KEY string

func onpublish(c *gin.Context) {
	c.Request.ParseForm()
	key := c.PostForm("name")
	if key == STREAMING_KEY {
		c.String(http.StatusOK, "Good to go")
	} else {
		c.String(http.StatusUnauthorized, "Invalid straming key")
	}
}
