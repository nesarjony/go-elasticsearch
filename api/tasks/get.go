// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package tasks

import (
	"fmt"
	"net/http"
	"net/url"
)

// Get - the task management API allows to retrieve information about the tasks currently executing on one or more nodes in the cluster. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/tasks.html for more info.
//
// taskID: return the task with specified id (node_id:task_number).
//
// options: optional parameters. Supports the following functional options: WithTaskID, WithWaitForCompletion, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (t *Tasks) Get(taskID string, options ...*Option) (*GetResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: t.transport.URL.Scheme,
			Host:   t.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Get"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := t.transport.Do(req)
	return &GetResponse{resp}, err
}

// GetResponse is the response for Get
type GetResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
