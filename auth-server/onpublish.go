package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var STREAMING_KEY string

func onpublish(c *gin.Context) {
	var b bytes.Buffer
	_, _ = b.ReadFrom(c.Request.Body)
	fmt.Println("Payload:", b.String())
	fmt.Println(c.Request.Header)
	key := c.DefaultPostForm("name", "no-name-found")
	if key == STREAMING_KEY {
		fmt.Printf("good key %s == %s\n", key, STREAMING_KEY)
		c.String(http.StatusOK, "Good to go")
	} else {
		fmt.Printf("bad key %s != %s\n", key, STREAMING_KEY)
		c.String(http.StatusUnauthorized, "Invalid straming key")
	}
}
