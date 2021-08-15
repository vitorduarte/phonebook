package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitorduarte/phonebook/internal/phonebook"
	"github.com/vitorduarte/phonebook/internal/storage"
)

func main() {
	s := storage.NewMemoryStorage()
	router := http.NewServeMux()

	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"service is up and running"}`))
		return
	})
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/contact", phonebook.Contact(s))

	srv := http.Server{
		Addr:         ":8080",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		Handler:      router,
	}

	fmt.Println("http server listening on port 8080")
	log.Fatal(srv.ListenAndServe())
}
