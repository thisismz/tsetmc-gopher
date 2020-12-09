package core

import (
	"fmt"
	"tsetmc-gopher/global"
	"tsetmc-gopher/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.BRC_CONFIG.Zap.Director); !ok { // Determine whether there is a Director folder
		fmt.Printf("create %v directory\n", global.BRC_CONFIG.Zap.Director)
		_ = os.Mkdir(global.BRC_CONFIG.Zap.Director, os.ModePerm)
	}

	switch global.BRC_CONFIG.Zap.Level { // Initialize the level of the configuration file
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.BRC_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.BRC_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.BRC_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // Lowercase encoder (default)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.BRC_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // Lowercase encoder with color
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.BRC_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // Uppercase encoder
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.BRC_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // Uppercase encoder with color
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder Obtain zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.BRC_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := utils.GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// Custom log output time format
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.BRC_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
