package logger

import (
	"github.com/DalvinCodes/bookish-couscous/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strconv"
)

var Logger *zap.Logger
var SLogger *zap.SugaredLogger

func ConfigureLogger(config *config.Config) {
	writeSyncer := logWriter(config)
	encoder := encoder()

	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)

	Logger = zap.New(core, zap.AddCaller())
	SLogger = Logger.Sugar()
	defer Logger.Sync()
}

func encoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func logWriter(config *config.Config) zapcore.WriteSyncer {

	filename := config.Logger.Filename
	maxSize, _ := strconv.Atoi(config.Logger.MaxSize)
	localTime, _ := strconv.ParseBool(config.Logger.LocalTime)
	compress, _ := strconv.ParseBool(config.Logger.Compress)

	lumberJack := &lumberjack.Logger{
		Filename:  filename,
		MaxSize:   maxSize,
		LocalTime: localTime,
		Compress:  compress,
	}

	return zapcore.AddSync(lumberJack)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...zap.Field) {
	Logger.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)

}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)

}

// DebugS uses fmt.Sprint to construct and log a message.
func DebugS(args ...interface{}) {
	SLogger.Debug(args)
}

// InfoS uses fmt.Sprint to construct and log a message.
func InfoS(args ...interface{}) {
	SLogger.Info(args)
}

// WarnS uses fmt.Sprint to construct and log a message.
func WarnS(args ...interface{}) {
	SLogger.Warn(args)
}

// ErrorS uses fmt.Sprint to construct and log a message.
func ErrorS(args ...interface{}) {
	SLogger.Error(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	SLogger.Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	SLogger.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	SLogger.Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	SLogger.Warnf(template, args)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	SLogger.DPanicf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	SLogger.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	SLogger.Fatalf(template, args)
}
