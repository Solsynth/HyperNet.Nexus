package kv

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var Kv *clientv3.Client

func ConnectEtcd(endpoints []string) error {
	conn, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
	})
	if err == nil {
		Kv = conn
	}
	return err
}
