package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	TotalRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "sentinel_requests_total",
			Help: "Total incoming requests",
		},
	)

	AllowedRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "sentinel_requests_allowed_total",
			Help: "Total allowed requests",
		},
	)

	BlockedRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "sentinel_requests_blocked_total",
			Help: "Total blocked requests",
		},
	)

	RedisErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "sentinel_redis_errors_total",
			Help: "Redis errors during rate limiting",
		},
	)

	DBFailures = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "sentinel_db_failures_total",
			Help: "Total database insert failures",
		},
	)
)

func Register() {
	prometheus.MustRegister(
		TotalRequests,
		AllowedRequests,
		BlockedRequests,
		RedisErrors,
		DBFailures,
	)
}
