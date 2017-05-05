// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// ExistsAlias - APIs in Elasticsearch accept an index name when working against a specific index, and several indices when applicable. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-aliases.html for more info.
//
// name: a comma-separated list of alias names to return.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithName, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithLocal, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) ExistsAlias(name []string, options ...*Option) (*ExistsAliasResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "HEAD",
	}
	methodOptions := supportedOptions["ExistsAlias"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &ExistsAliasResponse{resp}, err
}

// ExistsAliasResponse is the response for ExistsAlias
type ExistsAliasResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
