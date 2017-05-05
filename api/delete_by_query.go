// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// DeleteByQuery - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-delete-by-query.html for more info.
//
// index: a comma-separated list of index names to search; use "_all" or empty string to perform the operation on all indices.
//
// body: the search definition using the Query DSL.
//
// options: optional parameters. Supports the following functional options: WithType, WithSource, WithSourceExclude, WithSourceInclude, WithAllowNoIndices, WithAnalyzeWildcard, WithAnalyzer, WithConflicts, WithDefaultOperator, WithDf, WithExpandWildcards, WithFrom, WithIgnoreUnavailable, WithLenient, WithPreference, WithQ, WithRefresh, WithRequestCache, WithRequestsPerSecond, WithRouting, WithScroll, WithScrollSize, WithSearchTimeout, WithSearchType, WithSize, WithSlices, WithSort, WithStats, WithTerminateAfter, WithTimeout, WithVersion, WithWaitForActiveShards, WithWaitForCompletion, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) DeleteByQuery(index []string, body map[string]interface{}, options ...*Option) (*DeleteByQueryResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["DeleteByQuery"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &DeleteByQueryResponse{resp}, err
}

// DeleteByQueryResponse is the response for DeleteByQuery
type DeleteByQueryResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
