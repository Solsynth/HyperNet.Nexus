package kv

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var Kv *clientv3.Client

func ConnectEtcd(endpoints []string) error {
	conn, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		return err
	}
	var status []bool
	for _, endpoint := range endpoints {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		_, err := conn.Status(ctx, endpoint)
		if err != nil {
			log.Warn().Str("endpoint", endpoint).Err(err).Msg("An KV endpoint is not available...")
		}
		status = append(status, err == nil)
		cancel()
	}
	if len(lo.Filter(status, func(s bool, _ int) bool { return s })) == 0 {
		return fmt.Errorf("unable to connect to all KV endpoints")
	}
	Kv = conn
	return err
}
