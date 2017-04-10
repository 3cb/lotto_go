package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/megamillions", MegaHandler)
	r.HandleFunc("/api/powerball", PowerballHandler)

	n := negroni.New(negroni.NewStatic(http.Dir("./static/")))
	n.UseHandler(r)
	log.Fatal(http.ListenAndServe(":3000", n))
}
