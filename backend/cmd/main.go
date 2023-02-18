package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting url-shortener backend")
	log.Println("server starting...")

	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "welcome to dockerized app",
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/{msg}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		msg := vars["msg"]

		if msg == "" {
			msg = "hello"
		}

		response := map[string]string{
			"message": msg,
		}
		json.NewEncoder(rw).Encode(response)
	})

	log.Println("server running successfully!")
	fmt.Println(http.ListenAndServe(":8081", router))
}
