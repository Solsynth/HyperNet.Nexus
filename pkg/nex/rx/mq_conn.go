package rx

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"github.com/nats-io/nats.go"
)

type MqConn struct {
	n *nex.Conn

	Nt *nats.Conn
}

func NewMqConn(conn *nex.Conn) (*MqConn, error) {
	c := &MqConn{
		n: conn,
	}

	mqAddr := conn.AllocResource(nex.AllocatableResourceMq)
	if mqAddr == nil {
		return nil, fmt.Errorf("unable to allocate resource: message queue")
	} else if addr, ok := mqAddr.(string); !ok {
		return nil, fmt.Errorf("alloced mq resource address is not a string")
	} else if nc, err := nats.Connect(addr); err != nil {
		return nil, fmt.Errorf("unable to connect to nats server: %v", err)
	} else {
		c.Nt = nc
	}

	return c, nil
}
