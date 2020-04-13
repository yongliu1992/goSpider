package logrus

import (
	"io"
	"os"
	"yongliu1992/goSpider/helper/log/base"
	"yongliu1992/goSpider/helper/log/field"
)
import log "github.com/sirupsen/logrus"

// loggerLogrus 代表基于logrus的日志记录器的类型。
type loggerLogrus struct {
	// level 代表日志级别。
	level base.LogLevel
	// format 代表日志格式。
	format base.LogFormat
	// optWithLocation 代表OptWithLocation选项。
	// 该选项表示记录日志时是否带有调用方的代码位置。
	optWithLocation base.OptWithLocation
	// inner 代表内部使用的日志记录器。
	inner *log.Entry
}

// NewLogger 会新建并返回一个日志记录器。
func NewLogger() base.MyLogger {
	return NewLoggerBy(base.LevelInfo, base.FormatText, os.Stdout, nil)
}

// NewLoggerBy 会根据指定的参数新建并返回一个日志记录器。
func NewLoggerBy(
	level base.LogLevel,
	format base.LogFormat,
	writer io.Writer,
	options []base.Option) base.MyLogger {
	var logrusLevel log.Level
	switch level {
	default:
		level = base.LevelInfo
		logrusLevel = log.InfoLevel
	case base.LevelDebug:
		logrusLevel = log.DebugLevel
	case base.LevelWarn:
		logrusLevel = log.WarnLevel
	case base.LevelError:
		logrusLevel = log.ErrorLevel
	case base.LevelFatal:
		logrusLevel = log.FatalLevel
	case base.LevelPanic:
		logrusLevel = log.PanicLevel
	}
	var optWithLocation base.OptWithLocation
	if options != nil {
		for _, opt := range options {
			if opt.Name() == "with location" {
				optWithLocation, _ = opt.(base.OptWithLocation)
			}
		}
	}
	return &loggerLogrus{
		level:           level,
		format:          format,
		optWithLocation: optWithLocation,
		inner:           initInnerLogger(logrusLevel, format, writer),
	}
}

// initInnerLogger 会初始化内部使用的日志记录器。
func initInnerLogger(
	level log.Level,
	format base.LogFormat,
	writer io.Writer) *log.Entry {
	innerLogger := log.New()

	switch format {
	case base.FormatJson:
		innerLogger.Formatter = &log.JSONFormatter{
			TimestampFormat: base.TimeStampFormat,
		}
	default:
		innerLogger.Formatter = &log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: base.TimeStampFormat,
			DisableSorting:  false,
		}
	}
	innerLogger.Level = level
	innerLogger.Out = writer
	return log.NewEntry(innerLogger)
}

func (logger *loggerLogrus) Name() string {
	return "logrus"
}

func (logger *loggerLogrus) Level() base.LogLevel {
	return logger.level
}

func (logger *loggerLogrus) Format() base.LogFormat {
	return logger.format
}

func (logger *loggerLogrus) Options() []base.Option {
	return []base.Option{logger.optWithLocation}
}

func (logger *loggerLogrus) Debug(v ...interface{}) {
	logger.getInner().Debug(v...)
}

func (logger *loggerLogrus) Debugf(format string, v ...interface{}) {
	logger.getInner().Debugf(format, v...)
}

func (logger *loggerLogrus) Debugln(v ...interface{}) {
	logger.getInner().Debug(v...)
}

func (logger *loggerLogrus) Error(v ...interface{}) {
	logger.getInner().Error(v...)
}

func (logger *loggerLogrus) Errorf(format string, v ...interface{}) {
	logger.getInner().Errorf(format, v...)
}

func (logger *loggerLogrus) Errorln(v ...interface{}) {
	logger.getInner().Errorln(v...)
}

func (logger *loggerLogrus) Fatal(v ...interface{}) {
	logger.getInner().Fatal(v...)
}

func (logger *loggerLogrus) Fatalf(format string, v ...interface{}) {
	logger.getInner().Fatalf(format, v...)
}

func (logger *loggerLogrus) Fatalln(v ...interface{}) {
	logger.getInner().Fatalln(v...)
}

func (logger *loggerLogrus) Info(v ...interface{}) {
	logger.getInner().Info(v...)
}

func (logger *loggerLogrus) Infof(format string, v ...interface{}) {
	logger.getInner().Infof(format, v...)
}

func (logger *loggerLogrus) Infoln(v ...interface{}) {
	logger.getInner().Infoln(v...)
}

func (logger *loggerLogrus) Panic(v ...interface{}) {
	logger.getInner().Panic(v...)
}

func (logger *loggerLogrus) Panicf(format string, v ...interface{}) {
	logger.getInner().Panicf(format, v...)
}

func (logger *loggerLogrus) Panicln(v ...interface{}) {
	logger.getInner().Panicln(v...)
}

func (logger *loggerLogrus) Warn(v ...interface{}) {
	logger.getInner().Warning(v...)
}

func (logger *loggerLogrus) Warnf(format string, v ...interface{}) {
	logger.getInner().Warningf(format, v...)
}

func (logger *loggerLogrus) Warnln(v ...interface{}) {
	logger.getInner().Warningln(v...)
}

func (logger *loggerLogrus) WithFields(fields ...field.Field) base.MyLogger {
	fieldsLen := len(fields)
	if fieldsLen == 0 {
		return logger
	}
	logrusFields := make(map[string]interface{}, fieldsLen)
	for _, curField := range fields {
		logrusFields[curField.Name()] = curField.Value()
	}
	return &loggerLogrus{
		level:           logger.level,
		format:          logger.format,
		optWithLocation: logger.optWithLocation,
		inner:           logger.inner.WithFields(logrusFields),
	}
}

func (logger *loggerLogrus) getInner() *log.Entry {
	inner := logger.inner
	if logger.optWithLocation.Value {
		inner = WithLocation(inner)
	}
	return inner
}

// WithLocation 用于附加记录日志的代码的位置。
func WithLocation(entry *log.Entry) *log.Entry {
	funcPath, fileName, line := base.GetInvokerLocation(4)
	return entry.WithField(
		"location", map[string]interface{}{"func_path": funcPath, "file_name": fileName, "line": line},
	)
}
