package etcd

import (
	"strings"
)

const (
	trace = "trace"
	tags  = "tags"
	ltags = "tags/"
	sep   = "/"
)

func traceKey(traceId string) string {
	s := []string{trace, traceId}
	return strings.Join(s, "/")
}

func formTagKey(parts []string) string {
	s := []string{trace, tags}
	return strings.Join(append(s, parts...), "/")
}

func formKey(parts []string) string {
	s := []string{trace}
	return strings.Join(append(s, parts...), "/")
}

func lookupKey() string {
	s := []string{trace, tags}
	return strings.Join(s, "/")
}

func transformKey(lookupKey string) string {
	return strings.Replace(lookupKey, ltags, "", -1)
}

func extractTraceKey(key string) string {
	keyTransformed := transformKey(key)
	return strings.Split(keyTransformed, sep)[1]
}
