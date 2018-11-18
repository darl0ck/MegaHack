package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

// const crtPath = "server.crt"
// const keyPath = "server.key"
const crtPath = "/etc/letsencrypt/live/hackaton.paymon.org/fullchain.pem"
const keyPath = "/etc/letsencrypt/live/hackaton.paymon.org/privkey.pem"

var domainList = []string{"hackaton.paymon.org", "megahack.paymon.org"}

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

	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps)); err != nil {
			log.Fatalf("F:ERR handler.startViewOnProd err %s", err.Error())
		}
	}()

	err := http.ListenAndServeTLS(":443", crtPath, keyPath, router)
	if err != nil {
		log.Println(err.Error())
	}
}

func addRoute(router *mux.Router, path string, name string) {
	for _, domain := range domainList {
		router.Host(domain).Subrouter().HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
			fmt.Printf("%s\n", name)
			page(writer, request, name)
		})
	}
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

func redirectToHttps(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "https://"+request.Host+request.URL.Path, http.StatusTemporaryRedirect)
}
