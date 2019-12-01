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
	// Pass one: do Query by kv pairs sent from gateway
	// all pairs sent, are used as a lookup so all pairs
	// must be present in query item
	resp, err := db.Kv.Get(ctx, lookupKey(), clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, err
	}

	index := []string{}
	for _, item := range resp.Kvs {
		elem := &sPb.Tags{}
		err = proto.Unmarshal(item.Value, elem)
		if err != nil {
			return nil, err
		}

		for key, value := range req.Query {
			if val, ok := elem.Tags[key]; !ok || val != value {
				continue
			}
		}
		index = append(index, extractTraceKey(string(item.Key)))
	}

	// Pass two: based on index of keys that satisfied query
	// extrat spans and return tham
	traces := []*sPb.GetResp{}
	for _, key := range index {
		t, err := db.Get(ctx, &sPb.GetReq{TraceId: key})
		if err != nil {
			return nil, err
		}
		traces = append(traces, t)
	}

	return &sPb.ListResp{
		Traces: traces,
	}, nil
}

func (db *DB) Get(ctx context.Context, req *sPb.GetReq) (*sPb.GetResp, error) {
	trace := []*sPb.Span{}
	resp, err := db.Kv.Get(ctx, traceKey(req.TraceId), clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, err
	}

	for _, item := range resp.Kvs {
		elem := &sPb.Span{}
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
