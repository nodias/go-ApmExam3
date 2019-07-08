package middleware

import (
	"go-ApmCommon/shared/logger"
	"io"
	"net/http"
	"os"
)

func Logging(path string) func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	return func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		// Logic to write quest information, i.e. headers, user agent etc to a log file.
		var log = logger.Get()
		fpLog, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		multiWriter := io.MultiWriter(fpLog, os.Stdout)
		log.SetOutput(multiWriter)
		next(res, req)
	}
}
