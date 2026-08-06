package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	redirectCountMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "router_redirect_handler_redirect_total",
			Help: "Number of redirects served by redirect handlers",
		},
		[]string{
			PromLblRedirectType,
		},
	)

	backendRequestCountMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "router_backend_handler_request_total",
			Help: "Number of requests served by backend handlers",
		},
		[]string{
			PromLblBackendId,
			PromLblRequestMethod,
		},
	)

	backendResponseDurationSecondsMetric = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "router_backend_handler_response_duration_seconds",
			Help: "Histogram of response durations by backend",
		},
		[]string{
			PromLblBackendId,
			PromLblRequestMethod,
			PromLblResponseCode,
		},
	)
)

func RegisterMetrics(r prometheus.Registerer) {
	r.MustRegister(
		backendRequestCountMetric,
		backendResponseDurationSecondsMetric,
		redirectCountMetric,
	)
}
