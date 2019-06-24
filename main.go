package main

import (
	"github.com/nodias/go-ApmCommon/middleware"
	"github.com/nodias/go-ApmExam3/database"
	"github.com/nodias/go-ApmExam3/router"
	"github.com/urfave/negroni"
)

func init() {
	database.NewOpenDB()
}

func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware("C:/workspace/logs/go-ApmExam3.log")))
	n.UseHandler(router.NewRouter())
	n.Run(":7003")
}
