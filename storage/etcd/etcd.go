package etcd

import (
	"context"
	sPb "github.com/c12s/scheme/stellar"
	"github.com/c12s/stellar/model"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type DB struct {
	Kv     clientv3.KV
	Client *clientv3.Client
}

func New(conf *model.Config, timeout time.Duration) (*DB, error) {
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: timeout,
		Endpoints:   conf.Endpoints,
	})

	if err != nil {
		return nil, err
	}

	return &DB{
		Kv:     clientv3.NewKV(cli),
		Client: cli,
	}, nil
}

func (db *DB) List(context.Context, *sPb.ListReq) (*sPb.ListResp, error) {
	return nil, nil
}

func (db *DB) Get(context.Context, *sPb.GetReq) (*sPb.GetResp, error) {
	return nil, nil
}
