package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/dist/").Handler(http.FileServer(http.Dir("./static/")))

	r.HandleFunc("/api/megamillions", MegaHandler)
	r.HandleFunc("/api/powerball", PowerballHandler)

	log.Fatal(http.ListenAndServe(":5000", r))
}
