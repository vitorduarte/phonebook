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

func TestFindContactsByNameHandler(t *testing.T) {
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
			name: "find_contacts_by_name_returns_200_and_the_contact_when_it_exists",
			args: args{
				storage: &storage.MockStorage{
					Contacts: []contact.Contact{
						{
							Id:    "1",
							Name:  "Bob",
							Phone: "999999999",
						},
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodGet,
						"/contact?name=Bob",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `[{"id":"1","name":"Bob","phone":"999999999"}]`,
		},
		{
			name: "find_contacts_by_name_returns_200_and_empty_list_when_it_does_not_exists",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodGet,
						"/contact?name=Bob",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `[]`,
		},
		{
			name: "find_contacts_by_name_returns_200_and_the_contacts_when_contains_name",
			args: args{
				storage: &storage.MockStorage{
					Contacts: []contact.Contact{
						{
							Id:    "1",
							Name:  "Bob Silva",
							Phone: "999999999",
						},
						{
							Id:    "2",
							Name:  "Alice Silva",
							Phone: "999999999",
						},
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodGet,
						"/contact?name=Silva",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `[{"id":"1","name":"Bob Silva","phone":"999999999"},{"id":"2","name":"Alice Silva","phone":"999999999"}]`,
		},
		{
			name: "find_contacts_by_name_returns_500_when_repository_fails",
			args: args{
				storage: &storage.MockStorage{
					Error: errors.New("invalid connection"),
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodGet,
						"/contact?name=Bob",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := FindContactsByNameHandler(tt.args.storage)
			w := httptest.NewRecorder()
			handler(w, tt.args.req())
			result := w.Result()

			if result.StatusCode != tt.wantStatusCode {
				t.Errorf("FindContactsByNameHandler() status = %v, want %v", result.StatusCode, tt.wantStatusCode)
			}

			if tt.wantBody != "" {
				body, _ := ioutil.ReadAll(result.Body)
				bodyString := string(body)
				if bodyString != tt.wantBody {
					t.Errorf("FindContactsByNameHandler() body = %v, want %v", bodyString, tt.wantBody)
				}
			}
		})
	}
}
