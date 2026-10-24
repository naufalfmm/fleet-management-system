package slog

import (
	"context"
	"log/slog"
	"os"

	"github.com/naufalfmm/fleet-management-system/internal/pkg/logger"
)

type slogLogger struct {
	log *slog.Logger
}

func New() (logger.Logger, error) {
	return &slogLogger{
		log: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}, nil
}

func (s slogLogger) Info(ctx context.Context, msg string, data ...any) {
	s.log.InfoContext(ctx, msg, data...)
}

func (s slogLogger) Debug(ctx context.Context, msg string, data ...any) {
	s.log.DebugContext(ctx, msg, data...)
}

func (s slogLogger) Warn(ctx context.Context, msg string, data ...any) {
	s.log.WarnContext(ctx, msg, data...)
}

func (s slogLogger) Error(ctx context.Context, msg string, data ...any) {
	s.log.ErrorContext(ctx, msg, data...)
}
