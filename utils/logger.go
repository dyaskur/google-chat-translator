package utils

import (
	"cloud.google.com/go/logging"
	"context"
	"log/slog"
)

// CloudLoggingHandler implements slog.Handler
type CloudLoggingHandler struct {
	Handler slog.Handler
}

// Handle maps slog levels to Google Cloud Logging severity levels
func (h *CloudLoggingHandler) Handle(ctx context.Context, r slog.Record) error {
	// Map slog levels to Google Cloud severity
	severity := logging.Default
	switch {
	case r.Level == slog.LevelDebug:
		severity = logging.Debug
	case r.Level == slog.LevelInfo:
		severity = logging.Info
	case r.Level == slog.LevelWarn:
		severity = logging.Warning
	case r.Level == slog.LevelError:
		severity = logging.Error
	}

	// Add severity to the attributes
	attrs := make([]slog.Attr, 0, r.NumAttrs()+1)
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})
	attrs = append(attrs, slog.String("severity", severity.String()))

	// Create a new record with severity
	newRecord := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	newRecord.AddAttrs(attrs...)

	return h.Handler.Handle(ctx, newRecord)
}

func (h *CloudLoggingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CloudLoggingHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *CloudLoggingHandler) WithGroup(name string) slog.Handler {
	return &CloudLoggingHandler{Handler: h.Handler.WithGroup(name)}
}

func (h *CloudLoggingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}
