package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.POST("/on_publish", func(c *gin.Context) {
			c.Request.ParseForm()
			key := c.PostForm("name")
			if key == viper.GetString("STREAMING_KEY") {
				c.String(http.StatusOK, "Good to go")
			} else {
				c.String(http.StatusUnauthorized, "Invalid straming key")
			}
		})
		r.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
