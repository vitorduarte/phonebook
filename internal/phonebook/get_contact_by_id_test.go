package phonebook

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vitorduarte/phonebook/internal/contact"
	"github.com/vitorduarte/phonebook/internal/storage"
)

func TestGetContactByIdHandler(t *testing.T) {
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
			name: "get_contact_by_id_returns_200_and_the_contact_when_it_exist",
			args: args{
				storage: &storage.MockStorage{
					Contact: contact.Contact{
						Id:    "1",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodGet,
						"/contact/1",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `{"id":"1","name":"Bob","phone":"999999999"}`,
		},
		{
			name: "get_contact_by_id_returns_404_when_does_not_exist",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodGet,
						"/contact/1",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "get_contact_by_id_returns_500_when_repository_fails",
			args: args{
				storage: &storage.MockStorage{
					Error: errors.New("invalid connection"),
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodGet,
						"/contact/1",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := ContactHandler(tt.args.storage)
			w := httptest.NewRecorder()
			handler(w, tt.args.req())
			result := w.Result()

			if result.StatusCode != tt.wantStatusCode {
				t.Errorf("GetContactByIdHandler() status = %v, want %v", result.StatusCode, tt.wantStatusCode)
			}

			if tt.wantBody != "" {
				body, _ := ioutil.ReadAll(result.Body)
				bodyString := string(body)
				if bodyString != tt.wantBody {
					t.Errorf("GetContactByIdHandler() body = %v, want %v", bodyString, tt.wantBody)
				}
			}
		})
	}
}
