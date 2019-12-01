package etcd

import (
	"context"
	sPb "github.com/c12s/scheme/stellar"
	"github.com/c12s/stellar/model"
	"github.com/coreos/etcd/clientv3"
	"github.com/golang/protobuf/proto"
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

func (db *DB) List(ctx context.Context, req *sPb.ListReq) (*sPb.ListResp, error) {
	return nil, nil
}

func (db *DB) Get(ctx context.Context, req *sPb.GetReq) (*sPb.GetResp, error) {
	trace := []*sPb.TracePart{}
	resp, err := db.Kv.Get(ctx, traceKey(req.TraceId), clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, err
	}

	for _, item := range resp.Kvs {
		elem := &sPb.TracePart{}
		err = proto.Unmarshal(item.Value, elem)
		if err != nil {
			return nil, err
		}
		trace = append(trace, elem)
	}

	return &sPb.GetResp{
		Trace: trace,
	}, nil
}
