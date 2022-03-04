package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla/Mux Example"))
}

func idEvaluator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	intID, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("Error")
		os.Exit(2)
	}

	if intID <= 0 {
		http.Error(w, "ERROR", 500)
	} else {
		w.WriteHeader((http.StatusOK))
		fmt.Fprintf(w, "ID: "+key)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/v1/items/{id}", idEvaluator)
	log.Fatal(http.ListenAndServe("localhost:8080", myRouter))
}

func main() {
	handleRequests()
}
