package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/artem-shestakov/LookLogsApp/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	// Create router
	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	handlers.Register(router)

	srv := http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}
	go srv.ListenAndServe()

	// waiting for interrupt signal
	<-stop

}
