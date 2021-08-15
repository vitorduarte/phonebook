package phonebook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status          int
	ResponseMessage []byte
}

func (r *StatusRecorder) WriteHeader(statusCode int) {
	r.Status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *StatusRecorder) Write(b []byte) (int, error) {
	r.ResponseMessage = b
	return r.ResponseWriter.Write(b)
}

func newStatusRecorder(w http.ResponseWriter) *StatusRecorder {
	return &StatusRecorder{
		ResponseWriter: w,
		Status:         http.StatusOK,
	}
}

func LogEndpointHitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wr := newStatusRecorder(w)
		next.ServeHTTP(wr, r)

		end := time.Now()
		elapsed := end.Sub(start)

		msg := fmt.Sprintf("%v %v %v - %v", r.Method, r.URL.Path, wr.Status, elapsed)
		if wr.Status != http.StatusOK {
			var responseMessage map[string]interface{}
			json.Unmarshal(wr.ResponseMessage, &responseMessage)
			msg = fmt.Sprintf("%v | %v", msg, responseMessage["message"])
		}

		log.Println(msg)
	})
}
