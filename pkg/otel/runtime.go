package otel

import (
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
)

func StartRuntimePublisher() error {
	// // This reader is used as a stand-in for a reader that will actually export
	// // data. See https://pkg.go.dev/go.opentelemetry.io/otel/exporters for
	// // exporters that can be used as or with readers.
	// reader := metric.NewManualReader(
	// 	// Add the runtime producer to get histograms fromc the Go runtime.
	// 	metric.WithProducer(runtime.NewProducer()),
	// )
	// provider := metric.NewMeterProvider(metric.WithReader(reader))
	// lc.Append(fx.StopHook(provider.Shutdown))
	// otel.SetMeterProvider(provider)

	// Start go runtime metric collection.
	return runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second))
}
