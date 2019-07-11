package main

import (
	"context"
	"go-ApmCommon/models"
	"go-ApmCommon/shared/logger"
	"go-ApmCommon/shared/repository"
	"go-ApmExam3/router"

	"github.com/urfave/negroni"
)

var config models.TomlConfig

func init() {
	models.Load("config/%s/config.toml")
	config = *models.GetConfig()
	logger.Init()
	repository.Init()
	repository.NewOpenDB()
}

func main() {
	log := logger.New(context.Background())
	//n := negroni.New(negroni.HandlerFunc(middleware.Logging(config.Logconfig.Logpath)))
	n := negroni.New()
	n.UseHandler(router.NewRouter())
	log.Info("go-ApmExam3 - Server On!")
	n.Run(config.Servers["ApmExam3"].PORT)
}
