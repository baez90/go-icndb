// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// NewPrometheusMetrics creates a new http.Handler for the prometheus metrics operation
func NewPrometheusMetrics() *PrometheusMetrics {
	return &PrometheusMetrics{
		prometheusHandler: promhttp.Handler(),
	}
}

/*PrometheusMetrics swagger:route GET /metrics prometheusMetrics

Get metrics in a Prometheus compatible format

*/
type PrometheusMetrics struct {
	prometheusHandler http.Handler
}

func (o *PrometheusMetrics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	o.prometheusHandler.ServeHTTP(rw, r)
}
