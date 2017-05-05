// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// Forcemerge - the force merge API allows to force merging of one or more indices through an API. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-forcemerge.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithFlush, WithIgnoreUnavailable, WithMaxNumSegments, WithOnlyExpungeDeletes, WithOperationThreading, WithWaitForMerge, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Forcemerge(options ...*Option) (*ForcemergeResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["Forcemerge"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &ForcemergeResponse{resp}, err
}

// ForcemergeResponse is the response for Forcemerge
type ForcemergeResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
