package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/hello/{to}", helloname).Methods("POST")
	http.Handle("/", r)
	fmt.Printf("Starting Levi's project...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		http.ListenAndServe(":8080", nil)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Index page of the API made by Levi lugato")
}

func helloname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %v\n", vars["to"])
}
