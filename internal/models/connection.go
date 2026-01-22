package models

import "net"

type Connection struct {
	Conn            net.Conn
	RemotePublicKey interface{}
	SendChan        chan []byte
	CloseChan       chan struct{}
}
