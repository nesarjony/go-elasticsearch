// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"
	"net/url"
)

// PutSettings - allows to update cluster wide specific settings. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-update-settings.html for more info.
//
// options: optional parameters. Supports the following functional options: WithFlatSettings, WithMasterTimeout, WithTimeout, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cluster) PutSettings(options ...*Option) (*PutSettingsResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "PUT",
	}
	methodOptions := supportedOptions["PutSettings"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &PutSettingsResponse{resp}, err
}

// PutSettingsResponse is the response for PutSettings
type PutSettingsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
