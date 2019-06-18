// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newMLPreviewDatafeedFunc(t Transport) MLPreviewDatafeed {
	return func(datafeed_id string, o ...func(*MLPreviewDatafeedRequest)) (*Response, error) {
		var r = MLPreviewDatafeedRequest{DatafeedID: datafeed_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-preview-datafeed.html.
//
type MLPreviewDatafeed func(datafeed_id string, o ...func(*MLPreviewDatafeedRequest)) (*Response, error)

// MLPreviewDatafeedRequest configures the Ml  Preview Datafeed API request.
//
type MLPreviewDatafeedRequest struct {
	DatafeedID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MLPreviewDatafeedRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID) + 1 + len("_preview"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	path.WriteString("/")
	path.WriteString(r.DatafeedID)
	path.WriteString("/")
	path.WriteString("_preview")

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
func (f MLPreviewDatafeed) WithContext(v context.Context) func(*MLPreviewDatafeedRequest) {
	return func(r *MLPreviewDatafeedRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MLPreviewDatafeed) WithPretty() func(*MLPreviewDatafeedRequest) {
	return func(r *MLPreviewDatafeedRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MLPreviewDatafeed) WithHuman() func(*MLPreviewDatafeedRequest) {
	return func(r *MLPreviewDatafeedRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MLPreviewDatafeed) WithErrorTrace() func(*MLPreviewDatafeedRequest) {
	return func(r *MLPreviewDatafeedRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MLPreviewDatafeed) WithFilterPath(v ...string) func(*MLPreviewDatafeedRequest) {
	return func(r *MLPreviewDatafeedRequest) {
		r.FilterPath = v
	}
}
