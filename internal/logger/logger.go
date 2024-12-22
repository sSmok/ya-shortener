package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func init() {
	if globalLogger == nil {
		globalLogger = zap.New(core())
	}
}

// Logger - возвращает глобальный объект логера
func Logger() *zap.Logger {
	return globalLogger
}

// Debug - логирует сообщения на уровне Debug
func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

// Info - логирует сообщения на уровне Info
func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

// Warn - логирует сообщения на уровне Warn
func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}

// Error - логирует сообщения на уровне Error
func Error(msg string, fields ...zap.Field) {
	globalLogger.Error(msg, fields...)
}

// Fatal - логирует сообщения на уровне Fatal
func Fatal(msg string, fields ...zap.Field) {
	globalLogger.Fatal(msg, fields...)
}

func core() zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(cfg)
	return zapcore.NewCore(consoleEncoder, stdout, zap.InfoLevel)
}
