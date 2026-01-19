package peer

import (
	"fmt"
	"log"
	"net"

	"github.com/diyliv/p2p/internal/models"
	"github.com/diyliv/p2p/internal/models/interfaces"
)

type Peer struct {
	// instead of using hard-coded port value
	// reading from config file would be better
	listenPort    int
	listener      net.Listener
	privateKey    interface{}
	publicKey     interface{}
	cacheConn     interfaces.Cache
	commands      models.Command
	mode          string // broadcast or private
	privateTarget string // address for private mode
}

func NewPeer(port int) (*Peer, error) {
	return nil, nil
}

func (p *Peer) StartListening() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", p.listenPort))
	if err != nil {
		return err
	}
	p.listener = listener
	log.Printf("Listening on port %d", p.listenPort)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Connection accept error: %v", err)
				continue
			}
			go p.handleIncomingConnection(conn)
		}
	}()

	return nil
}

func (p *Peer) handleIncomingConnection(conn net.Conn) {

}
