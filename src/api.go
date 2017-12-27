package main

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"

)

type Api struct{
	r *mux.Router
	settings *Settings
}

func newApi(router *mux.Router, settings *Settings) *Api{
	return &Api{
		r:router,
		settings:settings,
	}
}

func (a *Api) init() {

	a.r.HandleFunc("/listCommands", a.getCommands).Methods("GET")
	//a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *Api) getCommands(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products := a.settings.Cmds
	//if err != nil {
	//	respondWithError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

	respondWithJSON(w, http.StatusOK, products)
}

// app.go

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}