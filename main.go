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
	router.HandleFunc("/lk/chat/", lkChat)
	router.HandleFunc("/lk/help/", lkHelp)
	router.HandleFunc("/lk/main/", lkMain)
	router.HandleFunc("/main/", mainT)
	router.HandleFunc("/signup/", signup)
	router.HandleFunc("/videochat/", videochat)
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
	println(tv.ExecuteTemplate(writer, "index", nil))
}

func lkChat(writer http.ResponseWriter, request *http.Request) {
	tv, err := template.ParseGlob("template/*")
	if err != nil {
		log.Printf("ERR %s", err.Error())
		http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	println(tv.ExecuteTemplate(writer, "lk_chat", nil))
}

func lkHelp(writer http.ResponseWriter, request *http.Request) {
	tv, err := template.ParseGlob("template/*")
	if err != nil {
		log.Printf("ERR %s", err.Error())
		http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	println(tv.ExecuteTemplate(writer, "lk_help", nil))
}

func lkMain(writer http.ResponseWriter, request *http.Request) {
	tv, err := template.ParseGlob("template/*")
	if err != nil {
		log.Printf("ERR %s", err.Error())
		http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	println(tv.ExecuteTemplate(writer, "lk_main", nil))
}

func mainT(writer http.ResponseWriter, request *http.Request) {
	tv, err := template.ParseGlob("template/*")
	if err != nil {
		log.Printf("ERR %s", err.Error())
		http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	println(tv.ExecuteTemplate(writer, "main", nil))
}

func signup(writer http.ResponseWriter, request *http.Request) {
	tv, err := template.ParseGlob("template/*")
	if err != nil {
		log.Printf("ERR %s", err.Error())
		http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	println(tv.ExecuteTemplate(writer, "sign_up", nil))
}

func videochat(writer http.ResponseWriter, request *http.Request) {
	tv, err := template.ParseGlob("template/*")
	if err != nil {
		log.Printf("ERR %s", err.Error())
		http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	println(tv.ExecuteTemplate(writer, "videochat", nil))
}
