package main

import (
	"net"
	"time"

	"github.com/gwuhaolin/livego/configure"
	"github.com/gwuhaolin/livego/protocol/rtmp"
	"github.com/mniak/multistreamer/v2/restream"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

var VERSION = "master"

func init() {
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("livego panic: ", r)
			time.Sleep(1 * time.Second)
		}
	}()

	stream := rtmp.NewRtmpStream()
	startRelay(stream)
}

func startRelay(stream *rtmp.RtmpStream) {
	rtmpAddr := configure.Config.GetString("rtmp_addr")

	rtmpListen, err := net.Listen("tcp", rtmpAddr)
	if err != nil {
		log.Fatal(err)
	}

	relayServer := restream.Restreamer(stream, []string{
		"rtmp://localhost:8888/abc/123",
		"rtmp://localhost:9999/abc/123",
	})
	log.Info("HLS server disable....")

	defer func() {
		if r := recover(); r != nil {
			log.Error("RTMP server panic: ", r)
		}
	}()
	log.Info("RTMP Listen On ", rtmpAddr)
	relayServer.Serve(rtmpListen)
}
