package main

import (
	"net/http"

	"go.elastic.co/apm/module/apmgorilla"

	"github.com/gorilla/mux"
	common "github.com/nodias/go-ApmCommon"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser)
	router.Use(apmgorilla.Middleware())
	n := negroni.New(negroni.HandlerFunc(common.LoggingMiddleware))
	n.UseHandler(router)
	n.Run(":7003")
}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	data := []byte("end of GetUsers")
	w.Write(data)
	return
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	data := []byte("end of GetUser")
	w.Write(data)
	return
}
