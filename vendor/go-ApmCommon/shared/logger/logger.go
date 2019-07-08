package logger

import (
	"context"
	"go-ApmCommon/models"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmlogrus"
)

var log *logrus.Logger
var config models.TomlConfig

func Init() {
	config = *models.GetConfig()
	log = &logrus.Logger{
		Out:   os.Stderr,
		Hooks: make(logrus.LevelHooks),
		Level: logrus.Level(config.Logconfig.Loglevel),
		Formatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "log.level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "function.name", //non-ECS
			},
		},
	}
	apm.DefaultTracer.SetLogger(log)
	log.AddHook(&apmlogrus.Hook{})
}

//sigltone
var instance *logrus.Logger
var once sync.Once

func Get() *logrus.Logger {
	once.Do(func() {
		instance = log
	})
	return instance
}

func New(ctx context.Context) *logrus.Entry {
	return Get().WithFields(apmlogrus.TraceContext(ctx))
}
