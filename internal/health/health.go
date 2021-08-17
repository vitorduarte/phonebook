package health

import "net/http"

func Healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"service is up and running"}`))
		return
	}
}
