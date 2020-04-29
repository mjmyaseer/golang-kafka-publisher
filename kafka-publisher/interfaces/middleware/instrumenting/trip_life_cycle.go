package instrumenting

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"context"
	"fmt"
 "Company import path goes here"

"time"
)

var (
	fieldKeys    = []string{"method", "error"}
	requestCount = kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "publisher",
		Subsystem: "name",
		Name:      "count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "namespace name",
		Subsystem: "name",
		Name:      "count",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
)

//TripLifeCycleMiddleware
type TripLifeCycle struct {
	publisher_service.Publisher
}

func (m TripLifeCycle) Publish(ctx context.Context, event string, data interface{}, key int) (output []publisher_service.Output, err error) {
	start := time.Now()
	defer func(start time.Time) {
		lvs := []string{"method", "trip_lifecycle_publisher", "error", fmt.Sprint(err != nil)}
		requestCount.With(lvs...).Add(1)
		requestLatency.With(lvs...).Observe(float64(time.Now().Sub(start).Nanoseconds()) / 1000)
	}(start)

	output, err = m.Publisher.Publish(ctx, event, data, key)
	return
}
