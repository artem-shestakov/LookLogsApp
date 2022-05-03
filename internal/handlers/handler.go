package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func Register(router *mux.Router) {
	router.HandleFunc("/", loadTemp).Methods(http.MethodGet)
	router.HandleFunc("/to-log", toLog).Methods(http.MethodPost)
}

func loadTemp(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")

	// Return page to user
	tmpl.Execute(w, nil)
}

func toLog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	logrus.WithFields(logrus.Fields{
		"time": time.Now().Format("2006-01-02 15:04:05"),
		"from": r.RemoteAddr,
	}).Info(r.PostForm.Get("logText"))
	http.Redirect(w, r, "/", 302)
}
