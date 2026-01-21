package observability

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// InitLogger initializes the global logger.
// In a real app, you would pass config here (e.g. env, log level).
func InitLogger() {
	once.Do(func() {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		)

		logger = zap.New(core, zap.AddCaller())
	})
}

// GetLogger returns the global logger instance.
// It initializes it if it hasn't been initialized yet.
func GetLogger() *zap.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

// Sugar returns the sugared logger.
func Sugar() *zap.SugaredLogger {
	return GetLogger().Sugar()
}
