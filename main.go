package main

import (
	"fmt"
	"go-ApmCommon/middleware"
	"go-ApmCommon/model"
	"go-ApmExam3/database"
	"go-ApmExam3/router"
	"os"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	database.NewOpenDB()
	config.New("config.toml")
	//EXPORT APM EXVIRONMENT
	apmurl := fmt.Sprintf("%s%s", config.Servers["APM_TESTSERVER"].IP, config.Servers["APM_TESTSERVER"].PORT)
	os.Setenv("ELASTIC_APM_SERVER_URL", apmurl)
	os.Setenv("ELASTIC_APM_SERVICE_NAME", config.Title)
}

func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware(config.Logpaths["local"].Path)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["local3"].PORT)
}
