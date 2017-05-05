// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"
	"net/url"
)

// Reroute - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-reroute.html for more info.
//
// options: optional parameters. Supports the following functional options: WithDryRun, WithExplain, WithMasterTimeout, WithMetric, WithRetryFailed, WithTimeout, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cluster) Reroute(options ...*Option) (*RerouteResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["Reroute"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &RerouteResponse{resp}, err
}

// RerouteResponse is the response for Reroute
type RerouteResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
