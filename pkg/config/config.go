package config

import (
	"context"
	"log/slog"

	"go.uber.org/fx"

	"github.com/matt-e/seed/pkg/stage"
)

var Module = fx.Module(
	"config",
	fx.Provide(New),
)

type Config struct {
	OtelCollectorEndpoint string
	OtelCollectorInsecure bool
}

func New(ctx context.Context, log *slog.Logger, s stage.Stage) *Config {
	log.InfoContext(ctx, "Loading configuration for stage", slog.String("stage", s.String()))
	return &Config{
		OtelCollectorEndpoint: "127.0.0.1:4317",
		OtelCollectorInsecure: true,
	}
}
