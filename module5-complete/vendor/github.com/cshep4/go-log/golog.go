// Package log is a small package designed for propagating a zap logger
// through contexts. It tries to be as un-intrusive as possible whilst providing
// production-ready defaults.
package log

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// defaultLogger is a production-ready logger. From will use defaultLogger
// if there is no zap logger already present in the context.
var defaultLogger = zap.New(zapcore.NewCore(
	zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}),
	zapcore.AddSync(os.Stdout),
	zap.NewAtomicLevelAt(zapcore.InfoLevel),
))

// key can only ever be one value and will not allocate when doing lookups.
// See https://github.com/golang/go/issues/17826.
type key struct{}

// From will return the zap logger associated with the context if present,
// otherwise it will return Default.
func From(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(key{}).(*zap.Logger); ok {
		return l
	}
	return defaultLogger
}

// With will add l to the context.
func With(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, key{}, l)
}

// WithFields will return a new context with the fields added to the associated
// logger.
func WithFields(ctx context.Context, fields ...zap.Field) context.Context {
	if len(fields) == 0 {
		return ctx
	}
	return With(ctx, From(ctx).With(fields...))
}

// WithOptions will return a new context with opts applied to the associated
// logger.
func WithOptions(ctx context.Context, opts ...zap.Option) context.Context {
	if len(opts) == 0 {
		return ctx
	}
	return With(ctx, From(ctx).WithOptions(opts...))
}

// Error will log at the error level using the logger associated with the
// context.
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	From(ctx).Error(msg, fields...)
}

// Info will log at the info level using the logger associated with the
// context.
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	From(ctx).Info(msg, fields...)
}
