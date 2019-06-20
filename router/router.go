package router

import (
	"log"
	"net/http"

	"github.com/nodias/go-ApmCommon/model"
	"github.com/nodias/go-ApmExam3/database"

	"github.com/gorilla/mux"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgorilla"
)

func NewRouter() *mux.Router {
	return router()
}

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/userInfo/{id}", getUserInfoHandler)
	router.HandleFunc("/hello/{name}", helloHandler)
	router.HandleFunc("/users", getUsersHandler)
	router.HandleFunc("/user/{id}", getUserHandler)
	router.Use(apmgorilla.Middleware())
	return router
}

func getUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	user := &model.User{}

	ctx := req.Context()
	span, ctx := apm.StartSpan(ctx, "getUserInfoHandler", "custom")
	defer span.End()

	id := mux.Vars(req)["id"]

	user, err := database.GetUserInfo(id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("router.getUserInfoHandler data : %s", user)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	span, ctx := apm.StartSpan(ctx, "helloHandler", "custom")
	defer span.End()

	data := mux.Vars(req)["name"]
	log.Println(data)
	w.Write([]byte(data))
	return
}

func getUsersHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	span, ctx := apm.StartSpan(ctx, "getUsersHandler", "custom")
	defer span.End()

	data := []byte("end of GetUsers")
	log.Println(data)
	w.Write(data)
	return
}

func getUserHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	span, ctx := apm.StartSpan(ctx, "getUserHandler", "custom")
	defer span.End()

	data := []byte("end of GetUser")
	log.Println(data)
	w.Write(data)
	return
}
