package bootstrap

import (
	"io"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/hyfco/gopkg/conf"
	"gopkg.in/natefinch/lumberjack.v2"
)

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

	return log.With(log.NewStdLogger(multiWriter),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		//"service.id", id,
		"service.name", psm,
		//"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}
