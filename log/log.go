package log

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New setups Zap to the correct log level and correct output format.
func New(logFormat, logLevel, path string) (*zap.Logger, error) {
	var zapConfig zap.Config

	switch logFormat {
	case "json":
		zapConfig = zap.NewProductionConfig()
		zapConfig.DisableStacktrace = true
	default:
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.DisableStacktrace = true
		zapConfig.DisableCaller = true
		zapConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {}
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig.OutputPaths = []string{path}
	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.LevelKey = "log_level"
	zapConfig.EncoderConfig.MessageKey = "message"
	zapConfig.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02T15:04:05.000Z07:00"))
	}

	// Set the logger
	switch logLevel {
	case "debug":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "fatal":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	go watchLevel(&zapConfig, logger)

	zap.RedirectStdLog(logger)

	return logger, nil
}

func watchLevel(config *zap.Config, logger *zap.Logger) {
	var (
		elevated     bool
		defaultLevel = config.Level
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGUSR1)

	for s := range c {
		if s == syscall.SIGINT {
			return
		}

		elevated = !elevated
		if elevated {
			config.Level.SetLevel(zap.DebugLevel)
			logger.Info("Log level elevated to debug")
		} else {
			logger.Info("Log level restored to original configuration", zap.String("level", defaultLevel.String()))
			config.Level.SetLevel(defaultLevel.Level())
		}
	}
}
