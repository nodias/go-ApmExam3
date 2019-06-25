package main

import (
	"go-ApmCommon/middleware"
	"go-ApmCommon/model"
	"go-ApmExam3/database"
	"go-ApmExam3/router"
	"os"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

// var phase *string

func init() {
	database.NewOpenDB()
	config.New("config.toml")

	// phase = flag.String("phase", "local", "select phase, (ex. local, dv) If it is 'local', modify config.toml to fit your own.")
	// flag.Parse()

	//EXPORT APM EXVIRONMENT
	os.Setenv("ELASTIC_APM_SERVER_URL", config.ApmServerUrl())
	os.Setenv("ELASTIC_APM_SERVICE_NAME", config.Service)
}

func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware(config.Logpaths.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam3"].PORT)
}
