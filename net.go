package gardendocker

import (
	"fmt"
	"net"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry/gunk/localip"
	"github.com/docker/docker/pkg/iptables"
)

//go:generate counterfeiter . Chain
type Chain interface {
	Forward(action iptables.Action, ip net.IP, port int, proto, dest_addr string, dest_port int) error
}

type NetHandler struct {
	ContainerIP string
	Chain       Chain
}

func (c *NetHandler) NetIn(hostPort, containerPort uint32) (uint32, uint32, error) {
	externalIP, _ := localip.LocalIP()

	if err := c.Chain.Forward(iptables.Add, net.ParseIP(externalIP), int(hostPort), "tcp", c.ContainerIP, int(containerPort)); err != nil {
		return 0, 0, fmt.Errorf("netin %d to %d: %s", hostPort, containerPort, err)
	}

	return 0, 0, nil
}

func (c *NetHandler) NetOut(netOutRule garden.NetOutRule) error {
	return nil
}
