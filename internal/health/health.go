package health

import (
	"encoding/json"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
)

type HealthCheckResponse struct {
	Database string `json:"database"`
}

func Healthcheck(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		healthCheckResponse := &HealthCheckResponse{}

		isDatabaseRunning := s.HealthCheck()
		if isDatabaseRunning {
			w.WriteHeader(http.StatusOK)
			healthCheckResponse.Database = "OK"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			healthCheckResponse.Database = "ERROR"
		}

		responseBytes, _ := json.Marshal(healthCheckResponse)
		w.Write(responseBytes)
	}
}
