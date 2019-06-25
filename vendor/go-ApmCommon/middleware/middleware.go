package middleware

import (
	"io"
	"net/http"
	"os"

	"go-ApmCommon/logger"
)

func NewLoggingMiddleware(path string) func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	return func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		// Logic to write request information, i.e. headers, user agent etc to a log file.
		var log = logger.Log
		fpLog, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer fpLog.Close()
		multiWriter := io.MultiWriter(fpLog, os.Stdout)
		log.SetOutput(multiWriter)
		next(res, req)
	}
}
