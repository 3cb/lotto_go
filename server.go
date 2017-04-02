package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	// r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/api/megamillions", MegaHandler)
	r.HandleFunc("/api/powerball", PowerballHandler)

	n := negroni.New(negroni.NewStatic(http.Dir("./")))
	n.UseHandler(r)
	log.Fatal(http.ListenAndServe(":3000", n))
}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(200)
// 	io.WriteString(w, "Welcome to the Homepage!!")
// }

func MegaHandler(w http.ResponseWriter, req *http.Request) {
	dateQuery := []string{}
	if err := json.NewDecoder(req.Body).Decode(&dateQuery); err != nil {
		log.Println(err)
	}

}

func PowerballHandler(w http.ResponseWriter, req *http.Request) {

}

type DateQuery struct {
	Start string
	End   string
}
