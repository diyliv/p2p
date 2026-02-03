package nat

import (
	"net"

	"github.com/huin/goupnp/dcps/internetgateway2"
)

type UPnPPortForwarder struct {
	client     *internetgateway2.WANIPConnection1
	internalIP net.IP
	externalIP net.IP
	port       int
}

func NewUPnPPortForwarder(port int) (*UPnPPortForwarder, error) {
	return nil, nil
}

func (u *UPnPPortForwarder) AddMapping()
func (u *UPnPPortForwarder) Removemapping()
func (u *UPnPPortForwarder) GetExternalIP()
func (u *UPnPPortForwarder) GetInternalIP()
