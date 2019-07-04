package logger

import (
	"context"
	"go-ApmCommon/model"

	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmlogrus"
)

var log *logrus.Logger
var config model.TomlConfig

func Init() {
	config = *model.GetConfig()
	log = &logrus.Logger{
		Out:   os.Stderr,
		Hooks: make(logrus.LevelHooks),
		Level: logrus.Level(config.Logconfig.Loglevel),
		Formatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "log.level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "agro", //non-ECS
			},
		},
	}
	apm.DefaultTracer.SetLogger(log)
	log.AddHook(&apmlogrus.Hook{})
}

var instance *logrus.Logger
var once sync.Once

func GetLogger() *logrus.Logger {
	once.Do(func() {
		instance = log
	})
	return instance
}

func NewLogger(ctx context.Context) *logrus.Entry {
	return GetLogger().WithFields(apmlogrus.TraceContext(ctx))
}
