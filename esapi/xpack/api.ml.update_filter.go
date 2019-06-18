// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMLUpdateFilterFunc(t Transport) MLUpdateFilter {
	return func(body io.Reader, filter_id string, o ...func(*MLUpdateFilterRequest)) (*Response, error) {
		var r = MLUpdateFilterRequest{Body: body, FilterID: filter_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MLUpdateFilter func(body io.Reader, filter_id string, o ...func(*MLUpdateFilterRequest)) (*Response, error)

// MLUpdateFilterRequest configures the Ml  Update Filter API request.
//
type MLUpdateFilterRequest struct {
	Body io.Reader

	FilterID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MLUpdateFilterRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("filters") + 1 + len(r.FilterID) + 1 + len("_update"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("filters")
	path.WriteString("/")
	path.WriteString(r.FilterID)
	path.WriteString("/")
	path.WriteString("_update")

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

	req, _ := newRequest(method, path.String(), r.Body)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
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
func (f MLUpdateFilter) WithContext(v context.Context) func(*MLUpdateFilterRequest) {
	return func(r *MLUpdateFilterRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MLUpdateFilter) WithPretty() func(*MLUpdateFilterRequest) {
	return func(r *MLUpdateFilterRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MLUpdateFilter) WithHuman() func(*MLUpdateFilterRequest) {
	return func(r *MLUpdateFilterRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MLUpdateFilter) WithErrorTrace() func(*MLUpdateFilterRequest) {
	return func(r *MLUpdateFilterRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MLUpdateFilter) WithFilterPath(v ...string) func(*MLUpdateFilterRequest) {
	return func(r *MLUpdateFilterRequest) {
		r.FilterPath = v
	}
}
