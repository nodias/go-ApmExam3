package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nodias/go-ApmCommon/model"
	"github.com/nodias/go-ApmCommon/response"
	"github.com/nodias/go-ApmExam3/service"

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
	router.HandleFunc("/panic/{rune}", panicHandler)
	router.Use(apmgorilla.Middleware())
	return router
}

//getUserInfoHandler is a function, gets the information of one User
func getUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	user := &model.User{}

	ctx := req.Context()
	span, ctx := apm.StartSpan(ctx, "getUserInfoHandler", "custom")
	defer span.End()

	id := mux.Vars(req)["id"]

	user, err := service.GetUserInfo(req.Context(), id)
	if err != nil {
		apm.CaptureError(req.Context(), err).Send()
	}
	err = json.NewEncoder(w).Encode(response.Response{
		Id:   response.ID(id),
		User: user,
		Err:  response.ResponseErr{err},
	})
	if err != nil {
		log.Println(err)
	}
	log.Printf("router.getUserInfoHandler data : %s", user)
}

func panicHandler(w http.ResponseWriter, req *http.Request) {
	p_rune := mux.Vars(req)["rune"]
	log.Printf("panicHandler : %v", p_rune)
	service.PanicGenerator(req.Context(), p_rune)
}
