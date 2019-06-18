// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
)

func newMLGetJobsFunc(t Transport) MLGetJobs {
	return func(o ...func(*MLGetJobsRequest)) (*Response, error) {
		var r = MLGetJobsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job.html.
//
type MLGetJobs func(o ...func(*MLGetJobsRequest)) (*Response, error)

// MLGetJobsRequest configures the Ml  Get Jobs API request.
//
type MLGetJobsRequest struct {
	JobID string

	AllowNoJobs *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MLGetJobsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	if r.JobID != "" {
		path.WriteString("/")
		path.WriteString(r.JobID)
	}

	params = make(map[string]string)

	if r.AllowNoJobs != nil {
		params["allow_no_jobs"] = strconv.FormatBool(*r.AllowNoJobs)
	}

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
func (f MLGetJobs) WithContext(v context.Context) func(*MLGetJobsRequest) {
	return func(r *MLGetJobsRequest) {
		r.ctx = v
	}
}

// WithJobID - the ID of the jobs to fetch.
//
func (f MLGetJobs) WithJobID(v string) func(*MLGetJobsRequest) {
	return func(r *MLGetJobsRequest) {
		r.JobID = v
	}
}

// WithAllowNoJobs - whether to ignore if a wildcard expression matches no jobs. (this includes `_all` string or when no jobs have been specified).
//
func (f MLGetJobs) WithAllowNoJobs(v bool) func(*MLGetJobsRequest) {
	return func(r *MLGetJobsRequest) {
		r.AllowNoJobs = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MLGetJobs) WithPretty() func(*MLGetJobsRequest) {
	return func(r *MLGetJobsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MLGetJobs) WithHuman() func(*MLGetJobsRequest) {
	return func(r *MLGetJobsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MLGetJobs) WithErrorTrace() func(*MLGetJobsRequest) {
	return func(r *MLGetJobsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MLGetJobs) WithFilterPath(v ...string) func(*MLGetJobsRequest) {
	return func(r *MLGetJobsRequest) {
		r.FilterPath = v
	}
}
