package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitorduarte/phonebook/internal/health"
	"github.com/vitorduarte/phonebook/internal/logs"
	"github.com/vitorduarte/phonebook/internal/phonebook"
	"github.com/vitorduarte/phonebook/internal/prometheus"
	"github.com/vitorduarte/phonebook/internal/storage"
)

func GetRoutes(s storage.Storage) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/healthcheck", prometheus.Middleware(logs.LogEndpointHitMiddleware(health.HealthCheckHandler(s))))
	router.Handle("/contact", prometheus.Middleware(logs.LogEndpointHitMiddleware(phonebook.ContactHandler(s))))
	router.Handle("/contact/", prometheus.Middleware(logs.LogEndpointHitMiddleware(phonebook.ContactHandler(s))))

	return router
}
