package bootstrap

import (
	kl "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hyfco/gopkg/conf"
	"github.com/hyfco/gopkg/ipx"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type LogrusHook struct {
	Id  string
	Ip  string
	psm string
}

func NewLogrusHook(psm string) logrus.Hook {
	id, _ := os.Hostname()

	return &LogrusHook{
		Id:  id,
		Ip:  "todo",
		psm: psm,
	}
}

func (l LogrusHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
	}
}

func (l LogrusHook) Fire(entry *logrus.Entry) error {
	entry.WithField("id", l.Id)
	entry.WithField("ip", l.Ip)
	entry.WithField("psm", l.psm)

	ip := ipx.GetClientIPFromContext(entry.Context)
	entry.WithField("client_ip", ip)

	span := trace.SpanContextFromContext(entry.Context)
	if span.HasTraceID() {
		entry.WithField("trace_id", span.TraceID())
	}

	if span.HasSpanID() {
		entry.WithField("span_id", span.SpanID())
	}

	return nil
}

func parseLevel(level string) logrus.Level {
	var logLevel logrus.Level
	switch level {
	case "debug":
		logLevel = logrus.DebugLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	default:
		logLevel = logrus.InfoLevel
	}
	return logLevel
}

func NewLogger(psm string, config *conf.Bootstrap) log.Logger {

	lc := config.GetLogger()
	var ws []io.Writer
	for _, output := range lc.OutputPaths {
		if output == "stdout" {
			ws = append(ws, os.Stdout)
		} else {
			iow := &lumberjack.Logger{
				Filename:   os.ExpandEnv(output),
				MaxSize:    int(lc.GetMaxSize()),
				MaxAge:     int(lc.GetMaxAge()),
				MaxBackups: int(lc.GetMaxBackup()),
				LocalTime:  true,
				Compress:   true,
			}
			ws = append(ws, iow)
		}
	}

	multiWriter := io.MultiWriter(ws...)

	logrus.SetLevel(parseLevel(lc.GetLevel()))
	logrus.SetOutput(multiWriter)
	if lc.GetFormat() == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.AddHook(NewLogrusHook(psm))
	return kl.NewLogger(logrus.StandardLogger())
}
