package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Spain Coronavirus Restrictions API")
}

func getRestrictions(w http.ResponseWriter, r *http.Request) {
	lang := mux.Vars(r)["lang"]
	com := mux.Vars(r)["com"]

	http.ServeFile(w, r, "./static/"+lang+"/"+com+".json")
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/{com}/{lang}", getRestrictions).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
