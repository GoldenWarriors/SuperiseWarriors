package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!");
}

func main() {
	router := mux.NewRouter().StrictSlash(true);
	router.HandleFunc("/", Index);

	log.Fatal(http.ListenAndServe(":8080", router));
}
