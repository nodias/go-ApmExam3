package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmlogrus"
)

func NewLogger(ctx context.Context) *logrus.Entry {
	return Log.WithFields(apmlogrus.TraceContext(ctx))
}

var Log = &logrus.Logger{
	Out:   os.Stderr,
	Hooks: make(logrus.LevelHooks),
	Level: logrus.InfoLevel,
	Formatter: &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "agro", //non-ECS
		},
	},
}

func init() {
	apm.DefaultTracer.SetLogger(Log)
	Log.AddHook(&apmlogrus.Hook{})
}
