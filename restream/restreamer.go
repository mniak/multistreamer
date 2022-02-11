package restream

import (
	"errors"
	"net"

	"github.com/gwuhaolin/livego/av"
	"github.com/gwuhaolin/livego/protocol/rtmp"
	"github.com/gwuhaolin/livego/protocol/rtmp/core"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	handler av.Handler
	targets []string
}

func Restreamer(h av.Handler, targets []string) *Server {
	return &Server{
		handler: h,
		targets: targets,
	}
}

func (s *Server) Serve(listener net.Listener) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("relayserver serve panic: ", r)
		}
	}()

	for {
		var netconn net.Conn
		netconn, err = listener.Accept()
		if err != nil {
			return
		}
		conn := core.NewConn(netconn, 4*1024)
		log.Debug("new client, connect remote: ", conn.RemoteAddr().String(),
			"local:", conn.LocalAddr().String())
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn *core.Conn) error {
	if err := conn.HandshakeServer(); err != nil {
		conn.Close()
		log.Error("handleConn HandshakeServer err: ", err)
		return err
	}
	connServer := core.NewConnServer(conn)

	if err := connServer.ReadMsg(); err != nil {
		conn.Close()
		log.Error("handleConn read msg err: ", err)
		return err
	}

	connServer.PublishInfo.Name = "stream"
	log.Debugf("handleConn: IsPublisher=%v", connServer.IsPublisher())
	if !connServer.IsPublisher() {
		err := errors.New("the server does not accept play streams")
		connServer.Close(err)
		log.Warn("handleConn closed connection because it is play stream")
		return err
	}

	reader := rtmp.NewVirReader(connServer)
	s.handler.HandleReader(reader)

	for _, target := range s.targets {
		connClient := core.NewConnClient()
		if err := connClient.Start(target, av.PUBLISH); err != nil {
			conn.Close()
			log.Error("handleConn client conn start err: ", err)
			return err
		}
		writer := rtmp.NewVirWriter(connClient)
		wrappedWriter := NewWriterWithKey(writer, reader.Info().Key)
		s.handler.HandleWriter(wrappedWriter)
	}

	return nil
}
