package log

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"go.uber.org/fx"

	"github.com/matt-e/seed/pkg/stage"
)

var Module = fx.Module("log",
	fx.Provide(New),
)

func New(s stage.Stage) *slog.Logger {
	if s == stage.Dev {
		return slog.New(tint.NewHandler(os.Stderr, nil))
	}
	return slog.New(slog.NewJSONHandler(os.Stderr, nil))
}
