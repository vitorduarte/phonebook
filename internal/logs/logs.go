package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	interceptor "github.com/vitorduarte/phonebook/internal/interceptor"
)

func LogEndpointHitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wr := interceptor.NewResponseRecorder(w)
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
