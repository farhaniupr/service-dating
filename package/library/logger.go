package library

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func ModuleLogger() LoggerZap {

	return LoggerZap{Logger: nil}
}

func Writelog(context interface{}, env Env, logLevel string, msg string) {
	var contextEcho echo.Context
	var isEcho bool
	var req *http.Request
	var res *echo.Response
	var fields []zapcore.Field

	if check, ok := context.(echo.Context); ok {
		contextEcho = check
		isEcho = true
	} else {
		isEcho = false
	}

	atomicLevel := zap.NewAtomicLevel()
	switch logLevel {
	case "debug":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zapcore.WarnLevel)
	case "err":
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	}

	endoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	writer := &lumberjack.Logger{
		Filename:   "./api.log",
		MaxSize:    5, //5mb
		MaxAge:     30,
		MaxBackups: 5, //5compress
		LocalTime:  true,
		Compress:   true,
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(endoderConfig),
		zapcore.AddSync(writer),
		atomicLevel,
	)

	logger := zap.New(zapCore, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	if isEcho {
		req = contextEcho.Request()
		res = contextEcho.Response()

		fields = []zapcore.Field{
			zap.Int("status", res.Status),
			zap.String("reques_id", contextEcho.Response().Header().Get(echo.HeaderXRequestID)),
			zap.String("method", req.Method),
			zap.String("uri", req.RequestURI),
			zap.String("host", req.Host),
			zap.String("remote_ip", contextEcho.RealIP()),
		}
		logger.Log(atomicLevel.Level(), msg, fields...)
	} else {
		logger.Log(atomicLevel.Level(), msg)
	}

	defer logger.Sync()
}
