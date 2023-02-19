package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ongaaron96/url-shortener/backend/handler"
)

func main() {
	fmt.Println("starting url-shortener backend")
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	log.Println("server starting...")
	handler.Start()
}
