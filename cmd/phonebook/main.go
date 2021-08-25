package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitorduarte/phonebook/internal/health"
	"github.com/vitorduarte/phonebook/internal/logs"
	"github.com/vitorduarte/phonebook/internal/phonebook"
	"github.com/vitorduarte/phonebook/internal/prometheus"
	"github.com/vitorduarte/phonebook/internal/storage"
)

func main() {
	writeTimeout := flag.Int("w", 1, "maximum duration before timing out writes of the response in seconds")
	readTimeout := flag.Int("r", 1, "maximum duration before timing out reads of the response in seconds")
	port := flag.Int("p", 8080, "port to expose the application")
	flag.Parse()

	// s := storage.NewInMemoryStorage()
	ms, err := storage.NewMongoDBStorage("mongodb://mongodb:27017")
	if err != nil {
		log.Fatal(err)
	}
	router := http.NewServeMux()
	prometheus.Init()

	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/healthcheck", prometheus.Middleware(logs.LogEndpointHitMiddleware(health.Healthcheck(ms))))
	router.Handle("/contact", prometheus.Middleware(logs.LogEndpointHitMiddleware(phonebook.Contact(ms))))

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", *port),
		WriteTimeout: time.Duration(*writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(*readTimeout) * time.Second,
		Handler:      router,
	}

	fmt.Printf("http server listening on port %v\n", *port)
	log.Fatal(srv.ListenAndServe())
}
