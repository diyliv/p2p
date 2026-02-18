package nat

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/pion/stun"
)

type STUNClient struct {
	serverAddr string
}

func NewSTUNClient(serverAddr string) *STUNClient {
	return &STUNClient{serverAddr: serverAddr}
}

func (c *STUNClient) GetPublicAddress() (net.IP, int, error) {
	conn, err := net.Dial("udp", c.serverAddr)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to dial STUN server: %w", err)
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return nil, 0, fmt.Errorf("failed to set deadline: %w", err)
	}

	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	if _, err := conn.Write(message.Raw); err != nil {
		return nil, 0, fmt.Errorf("failed to send STUN request: %w", err)
	}
	ip, port, err := c.receiveResponse(conn)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read response: %w", err)
	}
	return ip, port, nil
}

func (c *STUNClient) receiveResponse(conn net.Conn) (net.IP, int, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, 0, err
	}
	response := stun.Message{Raw: buf[:n]}
	if err := response.Decode(); err != nil {
		return nil, 0, err
	}
	var xorAddr stun.XORMappedAddress
	if err := xorAddr.GetFrom(&response); err != nil {
		return nil, 0, err
	}

	ip := xorAddr.IP
	port := xorAddr.Port
	log.Printf("STUN public address: %s:%d", ip.String(), port)
	return ip, port, nil
}

func (c *STUNClient) discoverPublicAddress() (net.IP, int, error) {
	servers := c.defaultSTUNServers()
	var lastErr error
	for _, server := range servers {
		client := NewSTUNClient(server)
		ip, port, err := client.GetPublicAddress()
		if err == nil {
			return ip, port, nil
		}
		lastErr = err
		log.Printf("STUN server %s failed: %v", server, err)
	}
	return nil, 0, fmt.Errorf("all STUN servers failed: %w", lastErr)
}

func (c *STUNClient) defaultSTUNServers() []string {
	return []string{
		"stun.l.google.com:19302",
		"stun1.l.google.com:19302",
		"stun2.l.google.com:19302",
		"stun3.l.google.com:19302",
		"stun4.l.google.com:19302",
	}
}
