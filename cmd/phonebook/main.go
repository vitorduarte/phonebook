package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vitorduarte/phonebook/internal/storage"
)

func main() {
	writeTimeout := flag.Int("w", 1, "maximum duration before timing out writes of the response in seconds")
	readTimeout := flag.Int("r", 1, "maximum duration before timing out reads of the response in seconds")
	port := flag.Int("p", 8080, "port to expose the application")
	flag.Parse()

	ms, err := storage.NewMongoStorage("mongodb://mongodb:27017")
	if err != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", *port),
		WriteTimeout: time.Duration(*writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(*readTimeout) * time.Second,
		Handler:      GetRoutes(ms),
	}

	fmt.Printf("http server listening on port %v\n", *port)
	log.Fatal(srv.ListenAndServe())
}
