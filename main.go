package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Error Response Model to explain error and http code for error
type errorModel struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

//not compatible with reflect.StructTag.Get Bad syntax???
type item struct {
	ID          int       `json: "id"`
	Date        time.Time `json: "date"`
	Description string    `json: "description"`
}

func idEvaluator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	intID, err := strconv.Atoi(key)
	// if id from URL cannot be parsed as an int, Unprocessable Entity Error
	if err != nil {
		log.Printf("error parsing id (%s): %s\r\n", key, err)

		m := errorModel{
			Message: "failed to parse id",
			Code:    "422",
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			log.Printf("error encoding response: %s\r\n", err)
			return
		}

		return
	}

	if intID <= 0 {
		http.Error(w, "ERROR", 500)
	} else if intID == 1 {
		// 404 Error, StatusNotFound
		m := errorModel{
			Message: "1 is not a valid ID",
			Code:    "404",
		}
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			log.Printf("error encoding response: %s\r\n", err)
			return
		}
		return
	} else {
		// return item
		w.WriteHeader((http.StatusOK))
		i := item{
			ID:   intID,
			Date: time.Now(),
		}
		// if I is ODD, return a different item JSON
		if intID%2 == 0 {
			i := item{
				ID:          intID,
				Date:        time.Now(),
				Description: "ID is Odd!",
			}
			json.NewEncoder(w).Encode(i)
			{
				fmt.Printf("", i)
			}
			return
		}
		json.NewEncoder(w).Encode(i)
		{
			fmt.Printf("", i)
		}
		return
	}
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/v1/items/{id}", idEvaluator)
	log.Fatal(http.ListenAndServe("localhost:8080", myRouter))
}

func main() {
	handleRequests()
}
