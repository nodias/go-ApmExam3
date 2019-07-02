package main

import (
	"go-ApmCommon/middleware"
	"go-ApmCommon/model"
	"go-ApmExam3/database"
	"go-ApmExam3/router"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	config.Load()
	database.NewOpenDB()
}

func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware(config.Logpaths.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam3"].PORT)
}
