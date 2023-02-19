package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ongaaron96/url-shortener/backend/util"
)

const DefaultStartCount = 100000000000

func Start() {
	urlConverter := NewUrlConverter(util.NewCounter(DefaultStartCount))
	router := mux.NewRouter()

	// Home page
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "welcome to url shortener",
		}
		json.NewEncoder(rw).Encode(response)
	})

	// Convert long to short
	router.HandleFunc("/url-shortener/{url}", func(rw http.ResponseWriter, r *http.Request) {
		url := mux.Vars(r)["url"]
		shortUrl, err := urlConverter.longToShort(url)

		errorMsg := ""
		response := make(map[string]string)
		if err != nil {
			log.Printf("error converting long to short url, err: %v", err)
			errorMsg = err.Error() // TODO: better error messages
		}

		response["errorMsg"] = errorMsg
		response["shortUrl"] = shortUrl
		json.NewEncoder(rw).Encode(response)
	})

	// Convert short to long
	router.HandleFunc("/{url}", func(rw http.ResponseWriter, r *http.Request) {
		url := mux.Vars(r)["url"]
		longUrl, err := urlConverter.shortToLong(url)

		errorMsg := ""
		response := make(map[string]string)
		if err != nil {
			log.Printf("error converting short to long url, err: %v", err)
			errorMsg = err.Error() // TODO: better error messages
		}

		response["errorMsg"] = errorMsg
		response["longUrl"] = longUrl
		json.NewEncoder(rw).Encode(response)
	})

	log.Println("server running successfully!")
	fmt.Println(http.ListenAndServe(":8081", router))
}
