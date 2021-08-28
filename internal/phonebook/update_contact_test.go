package phonebook

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vitorduarte/phonebook/internal/contact"
	"github.com/vitorduarte/phonebook/internal/storage"
)

func TestUpdateContactHandler(t *testing.T) {
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
			name: "update_returns_200_and_returns_contact_updated_when_it_exist",
			args: args{
				storage: &storage.MockStorage{
					Contact: contact.Contact{
						Id:    "1",
						Name:  "Bob Silva",
						Phone: "999999999",
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/1",
						strings.NewReader(`{"name":"Bob","phone":"999999998"}`),
					)
				},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "update_returns_200_when_name_is_empty",
			args: args{
				storage: &storage.MockStorage{
					Contact: contact.Contact{
						Id:    "1",
						Name:  "Bob Silva",
						Phone: "999999999",
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/1", strings.NewReader(`{"name":"","phone":"999999999"}`),
					)
				},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "update_returns_400_when_name_and_phone_is_empty",
			args: args{
				storage: &storage.MockStorage{
					Contact: contact.Contact{
						Id:    "1",
						Name:  "Bob Silva",
						Phone: "999999999",
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/1",
						strings.NewReader(`{"name":"","phone":""}`),
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "update_returns_400_when_body_json_is_empty",
			args: args{
				storage: &storage.MockStorage{
					Contact: contact.Contact{
						Id:    "1",
						Name:  "Bob Silva",
						Phone: "999999999",
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/1",
						strings.NewReader(`{}`),
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "update_returns_400_when_body_is_empty",
			args: args{
				storage: &storage.MockStorage{
					Contact: contact.Contact{
						Id:    "1",
						Name:  "Bob Silva",
						Phone: "999999999",
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/1",
						nil,
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "update_returns_400_when_body_is_not_a_json",
			args: args{
				storage: &storage.MockStorage{
					Contact: contact.Contact{
						Id:    "1",
						Name:  "Bob Silva",
						Phone: "999999999",
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/1",
						strings.NewReader(`invalidJson`),
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "update_returns_404_when_contact_not_exists",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/1",
						strings.NewReader(`{"name":"Bob","phone":"999999998"}`),
					)
				},
			},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "update_returns_400_when_id_is_not_sent",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/",
						strings.NewReader(`{"name":"Bob","phone":"999999998"}`),
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "update_returns_500_when_repository_fails",
			args: args{
				storage: &storage.MockStorage{
					Error: errors.New("invalid connection"),
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPut,
						"/contact/1",
						strings.NewReader(`{"name":"Bob","phone":"999999998"}`),
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
				t.Errorf("UpdateContactHandler() = %v, want %v", result.StatusCode, tt.wantStatusCode)
			}
		})
	}
}
