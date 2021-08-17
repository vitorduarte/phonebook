package interceptor

import "net/http"

type ResponseRecorder struct {
	http.ResponseWriter
	Status          int
	ResponseMessage []byte
}

func (r *ResponseRecorder) WriteHeader(statusCode int) {
	r.Status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *ResponseRecorder) Write(b []byte) (int, error) {
	r.ResponseMessage = b
	return r.ResponseWriter.Write(b)
}

func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		ResponseWriter: w,
		Status:         http.StatusOK,
	}
}
