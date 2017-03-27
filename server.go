package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	n := negroni.New(negroni.NewStatic(http.Dir("./")))
	n.UseHandler(r)
	log.Fatal(http.ListenAndServe(":3000", n))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "Welcome to the Homepage!!")
}
