package mq

import "github.com/nats-io/nats.go"

var Kmq *nats.Conn

func ConnectNats(in string) error {
	nc, err := nats.Connect(in)
	if err != nil {
		return err
	} else {
		Kmq = nc
	}

	return nil
}
