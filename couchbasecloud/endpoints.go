package couchbasecloud

import (
	"fmt"
	"net/url"
)

type endpoint string

const (
	listCloudsEndpoint   endpoint = "/clouds"
	listClustersEndpoint endpoint = "/clusters"
)

// endpoint formats an endpoint based on a template string. It will URL path escape the provided arguments.
func (e endpoint) format(args ...string) string {
	escaped := make([]interface{}, len(args))
	for index, arg := range args {
		escaped[index] = url.PathEscape(arg)
	}

	return fmt.Sprintf(string(e), escaped...)
}
