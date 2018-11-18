package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

	addRoute(router, "/", "index")
	addRoute(router, "/lk/chat/", "lk_chat")
	addRoute(router, "/lk/help/", "lk_help")
	addRoute(router, "/lk/main/", "lk_main")
	addRoute(router, "/main/", "main")
	addRoute(router, "/signup/", "sign_up")
	addRoute(router, "/videochat/", "videochat")

	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", router)
	if err != nil {
		log.Println(err.Error())
	}
}

func addRoute(router *mux.Router, path string, name string) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("%s\n", name)
		page(writer, request, name)
	})
}

func page(writer http.ResponseWriter, request *http.Request, name string) {
	tv, err := template.ParseGlob("template/*")
	if err != nil {
		log.Printf("ERR %s", err.Error())
		http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	println(tv.ExecuteTemplate(writer, name, nil))
}
