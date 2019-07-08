package main

import (
	"go-ApmCommon/models"
	"go-ApmCommon/shared/logger"
	"go-ApmCommon/shared/middleware"
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
	n := negroni.New(negroni.HandlerFunc(middleware.Logging(config.Logconfig.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam3"].PORT)
}
