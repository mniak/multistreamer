package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/mniak/multistreamer/tools/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mkconfigCmd = &cobra.Command{
	Use: "mkconfig",
}

var mkconfigNginxCmd = &cobra.Command{
	Use: "nginx",
	Run: func(cmd *cobra.Command, args []string) {

		t := template.Must(template.New("nginx-config").Parse(`
events {}
rtmp {
	server {
		listen {{.Port}};
		chunk_size 4096;
		application live {
			live on;
			record off;
			push {{.YoutubeURL}}/{{.YoutubeKey}};
			push rtmp://127.0.0.1:{{.STunnelPort}}/rtmp/{{.FacebookKey}};
			on_publish http://127.0.0.1:8080/on_publish;
		}
	}
}
`))
		var err error
		data := make(map[string]interface{})
		if data["Port"], err = cmd.Flags().GetInt("port"); err != nil {
			fmt.Println("Invalid port.", err.Error())
			os.Exit(1)
		}
		if data["STunnelPort"], err = cmd.Flags().GetInt("stunnel_port"); err != nil {
			fmt.Println("Invalid STunnel port.", err.Error())
			os.Exit(1)
		}
		if yturl, err := utils.GetRTMPURLFlag(cmd, "youtube_url"); err == nil {
			data["YoutubeURL"] = yturl.String()
		} else {
			fmt.Println("Invalid YouTube URL.", err.Error())
			os.Exit(1)
		}
		if data["YoutubeKey"], err = cmd.Flags().GetString("youtube_key"); err != nil {
			fmt.Println("Invalid YouTube Key.", err.Error())
			os.Exit(1)
		}
		if data["FacebookKey"], err = cmd.Flags().GetString("facebook_key"); err != nil {
			fmt.Println("Invalid Facebook Key.", err.Error())
			os.Exit(1)
		}
		var filename string
		if filename, err = cmd.Flags().GetString("output"); err != nil {
			fmt.Println("Invalid output filename.", err.Error())
			os.Exit(1)
		}
		var file *os.File
		if file, err = os.Create(filename); err != nil {
			fmt.Println("Could not open file.", err.Error())
			os.Exit(1)
		}
		defer file.Close()
		if err = t.Execute(file, data); err != nil {
			fmt.Println("Could not process template.", err.Error())
			os.Exit(1)
		}
	},
}

var mkconfigSTunnelCmd = &cobra.Command{
	Use: "stunnel",
	Run: func(cmd *cobra.Command, args []string) {

		t := template.Must(template.New("nginx-config").Parse(`
pid = /var/run/stunnel4/stunnel.pid
output = /var/log/stunnel4/stunnel.log

setuid = stunnel4
setgid = stunnel4

# https://www.stunnel.org/faq.html
socket = r:TCP_NODELAY=1
socket = l:TCP_NODELAY=1

debug = 4

[fb-live]
client = yes
accept = {{.STunnelPort}}
connect = {{.FacebookHost}}:{{.FacebookPort}}
verifyChain = no
`))
		var err error
		data := make(map[string]interface{})
		if data["STunnelPort"], err = cmd.Flags().GetInt("stunnel_port"); err != nil {
			fmt.Println("Invalid STunnel port.", err.Error())
			os.Exit(1)
		}
		if url, err := utils.GetRTMPURLFlag(cmd, "facebook_url"); err == nil {
			data["FacebookHost"] = url.Hostname()
			port := url.Port()
			if port == "" {
				if url.Scheme == "rtmps" {
					port = "443"
				} else {
					fmt.Println("The Facebook URL must use scheme rtmps://.")
					os.Exit(1)
				}
			}
			data["FacebookPort"] = port
		} else {
			fmt.Println("Invalid Facebook URL.", err.Error())
			os.Exit(1)
		}
		var filename string
		if filename, err = cmd.Flags().GetString("output"); err != nil {
			fmt.Println("Invalid output filename.", err.Error())
			os.Exit(1)
		}
		var file *os.File
		if file, err = os.Create(filename); err != nil {
			fmt.Println("Could not open file.", err.Error())
			os.Exit(1)
		}
		defer file.Close()
		if err = t.Execute(file, data); err != nil {
			fmt.Println("Could not process template.", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	viper.BindEnv("STREAMING_PORT")
	viper.BindEnv("STUNNEL_PORT")
	viper.BindEnv("YOUTUBE_URL")
	viper.BindEnv("YOUTUBE_KEY")
	viper.BindEnv("FACEBOOK_URL")
	viper.BindEnv("FACEBOOK_KEY")

	rootCmd.AddCommand(mkconfigCmd)

	// nginx
	mkconfigCmd.AddCommand(mkconfigNginxCmd)

	mkconfigNginxCmd.Flags().StringP("output", "o", "", "File where to save the results")
	mkconfigNginxCmd.MarkFlagRequired("output")
	mkconfigNginxCmd.MarkFlagFilename("output")

	viper.SetDefault("STREAMING_PORT", 1935)
	mkconfigNginxCmd.Flags().Int("port", viper.GetInt("STREAMING_PORT"), "Streaming Port")
	viper.SetDefault("STUNNEL_PORT", 1936)
	mkconfigNginxCmd.Flags().Int("stunnel_port", viper.GetInt("STUNNEL_PORT"), "STunnel4 Port")

	viper.SetDefault("YOUTUBE_URL", "rtmp://a.rtmp.youtube.com/live2")
	mkconfigNginxCmd.Flags().String("youtube_url", viper.GetString("YOUTUBE_URL"), "YouTube URL")
	if mkconfigNginxCmd.Flag("youtube_url").DefValue == "" {
		mkconfigNginxCmd.MarkFlagRequired("youtube_url")
	}

	mkconfigNginxCmd.Flags().String("youtube_key", viper.GetString("YOUTUBE_KEY"), "YouTube Key")
	if mkconfigNginxCmd.Flag("youtube_key").DefValue == "" {
		mkconfigNginxCmd.MarkFlagRequired("youtube_key")
	}

	viper.BindEnv("FACEBOOK_KEY")
	mkconfigNginxCmd.Flags().String("facebook_key", viper.GetString("FACEBOOK_KEY"), "Facebook Key")
	if mkconfigNginxCmd.Flag("facebook_key").DefValue == "" {
		mkconfigNginxCmd.MarkFlagRequired("facebook_key")
	}

	// stunnel
	mkconfigCmd.AddCommand(mkconfigSTunnelCmd)

	mkconfigSTunnelCmd.Flags().AddFlag(mkconfigNginxCmd.Flag("stunnel_port"))
	mkconfigSTunnelCmd.Flags().AddFlag(mkconfigNginxCmd.Flag("output"))

	viper.SetDefault("FACEBOOK_URL", "rtmps://live-api-s.facebook.com:443/rtmp/")
	mkconfigSTunnelCmd.Flags().String("facebook_url", viper.GetString("FACEBOOK_URL"), "Facebook URL")
	if mkconfigSTunnelCmd.Flag("facebook_url").DefValue == "" {
		mkconfigSTunnelCmd.MarkFlagRequired("facebook_url")
	}
}
