// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newWatcherAckWatchFunc(t Transport) WatcherAckWatch {
	return func(watch_id string, o ...func(*WatcherAckWatchRequest)) (*Response, error) {
		var r = WatcherAckWatchRequest{WatchID: watch_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-ack-watch.html.
//
type WatcherAckWatch func(watch_id string, o ...func(*WatcherAckWatchRequest)) (*Response, error)

// WatcherAckWatchRequest configures the Watcher  Ack Watch API request.
//
type WatcherAckWatchRequest struct {
	ActionID []string
	WatchID  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r WatcherAckWatchRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_watcher") + 1 + len("watch") + 1 + len(r.WatchID) + 1 + len("_ack") + 1 + len(strings.Join(r.ActionID, ",")))
	path.WriteString("/")
	path.WriteString("_watcher")
	path.WriteString("/")
	path.WriteString("watch")
	path.WriteString("/")
	path.WriteString(r.WatchID)
	path.WriteString("/")
	path.WriteString("_ack")
	if len(r.ActionID) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.ActionID, ","))
	}

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f WatcherAckWatch) WithContext(v context.Context) func(*WatcherAckWatchRequest) {
	return func(r *WatcherAckWatchRequest) {
		r.ctx = v
	}
}

// WithActionID - a list of the action ids to be acked.
//
func (f WatcherAckWatch) WithActionID(v ...string) func(*WatcherAckWatchRequest) {
	return func(r *WatcherAckWatchRequest) {
		r.ActionID = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f WatcherAckWatch) WithPretty() func(*WatcherAckWatchRequest) {
	return func(r *WatcherAckWatchRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f WatcherAckWatch) WithHuman() func(*WatcherAckWatchRequest) {
	return func(r *WatcherAckWatchRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f WatcherAckWatch) WithErrorTrace() func(*WatcherAckWatchRequest) {
	return func(r *WatcherAckWatchRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f WatcherAckWatch) WithFilterPath(v ...string) func(*WatcherAckWatchRequest) {
	return func(r *WatcherAckWatchRequest) {
		r.FilterPath = v
	}
}
