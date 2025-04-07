package otel

import (
	"time"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
)

var stdoutExportInterval = 3 * time.Second

var StdoutModule = fx.Module("otel(stdout)",
	fx.Provide(
		NewStdoutTraceProvider,
		NewStdoutMeterProvider,
		NewStdoutLogProvider,
	),
)

func NewStdoutTraceProvider(lc fx.Lifecycle) (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),
	)

	lc.Append(fx.StopHook(traceProvider.Shutdown))

	return traceProvider, nil
}

func NewStdoutMeterProvider(lc fx.Lifecycle) (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(stdoutExportInterval))),
	)
	lc.Append(fx.StopHook(meterProvider.Shutdown))
	return meterProvider, nil
}

func NewStdoutLogProvider(lc fx.Lifecycle) (*log.LoggerProvider, error) {
	// Set up logger provider.
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)

	lc.Append(fx.StopHook(loggerProvider.Shutdown))
	return loggerProvider, nil
}
