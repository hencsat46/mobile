package logger

import (
	"log"
	"log/slog"
	"os"

	"hackathon/pkg/config"
)

type Logger struct {
	l *slog.Logger
}

func New(cfg *config.Config) *Logger {
	if cfg.Environment == "dev" {
		log := new(Logger)
		log.l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}))
		return log
	} else if cfg.Environment == "prod" {
		log := new(Logger)
		log.l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelInfo,
		}))

		return log
	} else {
		log.Fatal("choose the environment")
	}

	return nil
}

func (l *Logger) SetAsDefault() {
	slog.SetDefault(l.l)
}
