package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitorduarte/phonebook/internal/phonebook"
	"github.com/vitorduarte/phonebook/internal/storage"
)

func main() {
	writeTimeout := flag.Int("w", 1, "maximum duration before timing out writes of the response in seconds")
	readTimeout := flag.Int("r", 1, "maximum duration before timing out reads of the response in seconds")
	port := flag.Int("p", 8080, "port to expose the application")
	flag.Parse()

	s := storage.NewMemoryStorage()
	router := http.NewServeMux()

	router.Handle("/healthcheck", phonebook.LogEndpointHitMiddleware(healthcheck()))
	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/contact", phonebook.LogEndpointHitMiddleware(phonebook.Contact(s)))

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", *port),
		WriteTimeout: time.Duration(*writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(*readTimeout) * time.Second,
		Handler:      router,
	}

	fmt.Printf("http server listening on port %v\n", *port)
	log.Fatal(srv.ListenAndServe())
}

func healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"service is up and running"}`))
		return
	}
}
