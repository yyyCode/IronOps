package logger

import (
	"IronOps/internal/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger() {
	var err error
	
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Dynamic level from config
	level := zap.InfoLevel
	if config.GlobalConfig != nil && config.GlobalConfig.Server.Mode == "debug" {
		level = zap.DebugLevel
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Use Console encoder for readability, or JSON for prod
		zapcore.AddSync(os.Stdout),
		level,
	)

	Log = zap.New(core, zap.AddCaller())
	
	if err != nil {
		panic(err)
	}
	
	zap.ReplaceGlobals(Log)
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}
