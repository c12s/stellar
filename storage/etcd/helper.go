package etcd

import (
	"strings"
)

const (
	trace = "trace/"
)

func traceKey(traceId string) string {
	s := []string{trace, traceId}
	return strings.Join(s, "/")
}
