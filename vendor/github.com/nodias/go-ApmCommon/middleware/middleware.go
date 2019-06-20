package middleware

import (
	"io"
	"log"
	"net/http"
	"os"
)

func LoggingMiddleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// Logic to write request information, i.e. headers, user agent etc to a log file.
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)
	next(res, req)
}
