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

func (s slogLogger) Info(ctx context.Context, msg string, data logger.LoggerData) {
	if data == nil {
		s.log.InfoContext(ctx, msg, nil)
		return
	}
	s.log.InfoContext(ctx, msg, data.Args()...)
}

func (s slogLogger) Debug(ctx context.Context, msg string, data logger.LoggerData) {
	if data == nil {
		s.log.DebugContext(ctx, msg, nil)
		return
	}
	s.log.DebugContext(ctx, msg, data.Args()...)
}

func (s slogLogger) Warn(ctx context.Context, msg string, data logger.LoggerData) {
	if data == nil {
		s.log.WarnContext(ctx, msg, nil)
		return
	}
	s.log.WarnContext(ctx, msg, data.Args()...)
}

func (s slogLogger) Error(ctx context.Context, msg string, data logger.LoggerData) {
	if data == nil {
		s.log.ErrorContext(ctx, msg, nil)
		return
	}
	s.log.ErrorContext(ctx, msg, data.Args()...)
}
