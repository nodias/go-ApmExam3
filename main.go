package main

import (
	"go-ApmCommon/logger"
	"go-ApmCommon/middleware"
	"go-ApmCommon/model"
	"go-ApmExam3/database"
	"go-ApmExam3/router"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	model.Load("config/%s/config.toml")
	config.GetConfig()
	logger.Init()
	database.Init()
	database.NewOpenDB()
}

func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware(config.Logconfig.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam3"].PORT)
}
