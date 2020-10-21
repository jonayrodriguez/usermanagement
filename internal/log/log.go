package log

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Constants to define the type of format
const (
	jsonFormat      = "json"
	plainTextFormat = "plain"
	SystemLogType   = "system"
	AccessLogType   = "access"
)

// Builder interface for creating a logger
type Builder interface {
	SetLoggerType(loggerType string) Builder
	SetFullFilePath(fullFilePath string) Builder
	SetFormat(format string) Builder
	SetLevel(level string) Builder
	SetMaxSize(maxSize int) Builder
	SetMaxBackup(maxBackup int) Builder
	SetMaxAge(maxAge int) Builder
	NeedCaller(addCaller bool) Builder
	Build() LoggerBuilder
}

// Logger struct containing the logger and sugared logger
type Logger struct {
	Logger        zap.Logger
	SugaredLogger zap.SugaredLogger
}

// accessLogger instance to apply the singleton pattern
var accessLogger *Logger

// systemLogger instance to apply the singleton pattern
var systemLogger *Logger

// LoggerBuilder for logger
type LoggerBuilder struct {
	loggerType   string `default:"system"`
	fullFilePath string `default:"../logs/usermanagement.log"`
	format       string `default:"plain"`
	level        string `default:"debug"`
	addCaller    bool   `default:"false"`
	maxSize      int    `default:"5"`
	maxBackup    int    `default:"10"`
	maxAge       int    `default:"15"`
}

// SetLoggerType is used to set the logger type. The default is system typet.
func (lb *LoggerBuilder) SetLoggerType(loggerType string) Builder {
	lb.loggerType = loggerType
	return lb

}

// SetFullFilePath is used to set the path of the log file
func (lb *LoggerBuilder) SetFullFilePath(fullFilePath string) Builder {
	lb.fullFilePath = fullFilePath
	return lb

}

// SetFormat is used to set the format of the encode. The default is plain text.
func (lb *LoggerBuilder) SetFormat(format string) Builder {
	lb.format = format
	return lb

}

// SetLevel is used to set the logger level. It defaults to debug level.
func (lb *LoggerBuilder) SetLevel(level string) Builder {
	lb.level = level
	return lb

}

// SetMaxSize is used to set the maximum size in megabytes of the log file before it gets
// rotated. It defaults to 5 megabytes.
func (lb *LoggerBuilder) SetMaxSize(maxSize int) Builder {
	lb.maxSize = maxSize
	return lb

}

// SetMaxBackup is used to set the maximum number of old log files to retain.  The default
// is to retain 10 old log files (though MaxAge may still cause them to get
// deleted.)
func (lb *LoggerBuilder) SetMaxBackup(maxBackup int) Builder {
	lb.maxBackup = maxBackup
	return lb

}

// SetMaxAge is used to set the maximum number of days to retain old log files based on the
// timestamp encoded in their filename. The default is 15 days.
// based on age.
func (lb *LoggerBuilder) SetMaxAge(maxAge int) Builder {
	lb.maxAge = maxAge
	return lb

}

// NeedCaller to add the caller to the log
func (lb *LoggerBuilder) NeedCaller(addCaller bool) Builder {
	lb.addCaller = addCaller
	return lb

}

// Build the logger to able to get it
func (lb *LoggerBuilder) Build() LoggerBuilder {
	// TODO- Add validation
	return LoggerBuilder{
		loggerType:   lb.loggerType,
		fullFilePath: lb.fullFilePath,
		format:       lb.format,
		level:        lb.level,
		addCaller:    lb.addCaller,
		maxSize:      lb.maxSize,
		maxBackup:    lb.maxBackup,
		maxAge:       lb.maxAge,
	}
}

// GetLogger (Singleton) according to the builder process.
func (lb LoggerBuilder) GetLogger() (*Logger, error) {
	switch strings.ToLower(lb.loggerType) {
	case SystemLogType:
		if systemLogger == nil {
			zapLogger := new(lb)
			systemLogger = &Logger{
				Logger:        *zapLogger,
				SugaredLogger: *zapLogger.Sugar(),
			}
		}
		return systemLogger, nil
	case AccessLogType:
		if accessLogger == nil {
			zapLogger := new(lb)
			accessLogger = &Logger{
				Logger:        *zapLogger,
				SugaredLogger: *zapLogger.Sugar(),
			}
		}
		return accessLogger, nil
	default:
		return nil, fmt.Errorf("Logger´s type %s not recognized", lb.loggerType)
	}

}

// Creates a new logger based on the loggerBuilder
func new(lb LoggerBuilder) *zap.Logger {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   lb.fullFilePath,
		MaxSize:    lb.maxSize, // megabytes
		MaxBackups: lb.maxBackup,
		MaxAge:     lb.maxAge, // days
	})

	core := zapcore.NewCore(
		getEncoder(lb.format),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		unmarshalText(lb.level),
	)
	// Add zap.Development() to have panic logs
	var logger *zap.Logger
	if lb.addCaller {
		logger = zap.New(core, zap.AddCaller())
	} else {
		logger = zap.New(core)

	}
	//TODO- Singleton pattern wouldn't be required if ReplaceGlobals is used,
	// but keep in mind that more than 1 loggers could be used to log in different files
	//zap.ReplaceGlobals(logger)

	return logger
}

// Get the encoder in a specific format
func getEncoder(logFormat string) zapcore.Encoder {

	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	if strings.ToLower(logFormat) == jsonFormat {
		return zapcore.NewJSONEncoder(encoderConfig)

	}
	encoderConfig.MessageKey = "msg"
	encoderConfig.TimeKey = "time"
	return zapcore.NewConsoleEncoder(encoderConfig)

}

// Conver string level to a zapcore level
func unmarshalText(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
