package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "home.html")
}
var addr = flag.String("addr", ":8080", "http service address")

func main() {
	conf := readSettings()
	fmt.Printf("Hello, world.\n" + conf.Cmds[0].Cmd)
	flag.Parse()
	hub := newHub()
	a := newCmdWatcher(hub)
	a.addComd(conf.Cmds[0])
	go hub.run()
	mainRouter := mux.NewRouter()
	apiRouter := mainRouter.PathPrefix("/api").Subrouter()
	api:=newApi(apiRouter, &conf)
	api.init()
	mainRouter.HandleFunc("/", serveHome)
	fs := http.FileServer(http.Dir("static"))
	mainRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	mainRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, mainRouter)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
