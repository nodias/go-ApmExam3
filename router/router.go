package router

import (
	"encoding/json"
	"go-ApmCommon/logger"
	"go-ApmCommon/response"
	"go-ApmExam3/service"
	"net/http"
	"strings"
	"unicode"

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
	router.Use(apmgorilla.Middleware())
	return router
}

//getUserInfoHandler is a function, gets the information of one User
func getUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	log := logger.NewMyLogger(req)
	id := mux.Vars(req)["id"]
	log.WithField("id", id).Info("handling hello request")
	if strings.IndexFunc(id, func(r rune) bool { return r >= unicode.MaxASCII }) >= 0 {
		panic("non-ASCII id!")
	}

	user, rerr := service.GetUserInfo(req.Context(), id)
	if rerr != nil {
		log.WithError(rerr.Err).Error("failed to GetUserInfo")
		apm.CaptureError(req.Context(), rerr.Err).Send()
		w.WriteHeader(rerr.Code)
	}
	err := json.NewEncoder(w).Encode(response.Response{
		Id:    response.ID(id),
		User:  user,
		Error: rerr,
	})
	if err != nil {
		log.WithError(err).Error("failed to GetUserInfo")
		http.Error(w, "failed encode to json", 500)
		return
	}
}
