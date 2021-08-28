package phonebook

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vitorduarte/phonebook/internal/storage"
)

func TestCreateContactHandler(t *testing.T) {
	type args struct {
		storage storage.Storage
		req     func() *http.Request
	}

	tests := []struct {
		name           string
		args           args
		wantStatusCode int
	}{
		{
			name: "create_returns_200_when_all_good",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/contact", strings.NewReader(`{"name":"Bob Silva","phone":"999999999"}`),
					)
				},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "create_returns_200_when_name_is_empty",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/contact", strings.NewReader(`{"name":"","phone":"999999999"}`),
					)
				},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "create_returns_400_when_name_and_phone_is_empty",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/contact",
						strings.NewReader(`{"name":"","phone":""}`),
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "create_returns_400_when_body_json_is_empty",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/contact",
						strings.NewReader(`{}`),
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "create_returns_400_when_body_is_empty",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/contact",
						nil,
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "create_returns_400_when_body_is_not_a_json",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/contact",
						strings.NewReader(`invalidJson`),
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "create_returns_500_when_repository_fails",
			args: args{
				storage: &storage.MockStorage{
					Error: errors.New("invalid connection"),
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
						"/contact",
						strings.NewReader(`{"name":"Bob Silva","phone":"999999999"}`),
					)
				},
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := CreateContactHandler(tt.args.storage)
			w := httptest.NewRecorder()
			handler(w, tt.args.req())
			result := w.Result()
			if result.StatusCode != tt.wantStatusCode {
				t.Errorf("createContactHandler() = %v, want %v", result.StatusCode, tt.wantStatusCode)
			}
		})
	}
}
