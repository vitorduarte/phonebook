package health

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vitorduarte/phonebook/internal/storage"
)

func TestHealthCheckHandler(t *testing.T) {
	type args struct {
		storage storage.Storage
		req     func() *http.Request
	}

	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "create_returns_200_and_ok_when_database_is_up",
			args: args{
				storage: &storage.MockStorage{
					Health: true,
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/healthcheck",
						strings.NewReader(``),
					)
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `{"database":"OK"}`,
		},
		{
			name: "create_returns_500_and_error_message_when_database_is_down",
			args: args{
				storage: &storage.MockStorage{
					Health: false,
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/healthcheck",
						strings.NewReader(``),
					)
				},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       `{"database":"ERROR"}`,
		},
		{
			name: "create_returns_500_when_storage_fails",
			args: args{
				storage: &storage.MockStorage{
					Error: errors.New("invalid connection"),
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/healthcheck",
						strings.NewReader(``),
					)
				},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       `{"database":"ERROR"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := HealthCheckHandler(tt.args.storage)
			w := httptest.NewRecorder()
			handler(w, tt.args.req())
			result := w.Result()
			if result.StatusCode != tt.wantStatusCode {
				t.Errorf("HealthCheckHandler() = %v, want %v", result.StatusCode, tt.wantStatusCode)
			}

			if tt.wantBody != "" {
				body, _ := ioutil.ReadAll(result.Body)
				bodyString := string(body)
				if bodyString != tt.wantBody {
					t.Errorf("HealthCheckHandler() body = %v, want %v", bodyString, tt.wantBody)
				}
			}
		})
	}
}
