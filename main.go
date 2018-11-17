package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	router.HandleFunc("/", index)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", router)
	if err != nil {
		log.Println(err.Error())
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	tv, err := template.ParseGlob("template/*")
	if err != nil {
		log.Printf("ERR %s", err.Error())
		http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	println(tv.ExecuteTemplate(writer, "videochat", nil))
}
