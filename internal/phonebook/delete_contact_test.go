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

func TestDeleteContactHandler(t *testing.T) {
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
			name: "delete_returns_200_when_it_exist",
			args: args{
				storage: &storage.MockStorage{
					Contact: contact.Contact{
						Id:    "1",
						Name:  "Vitor",
						Phone: "999999999",
					},
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodDelete,
						"/contact/1",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "delete_returns_404_when_does_not_exist",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodDelete,
						"/contact/1",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "delete_returns_400_when_id_is_not_sent",
			args: args{
				storage: &storage.MockStorage{},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodDelete,
						"/contact/",
						strings.NewReader(""),
					)
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "delete_returns_500_when_repository_fails",
			args: args{
				storage: &storage.MockStorage{
					Error: errors.New("invalid connection"),
				},
				req: func() *http.Request {
					return httptest.NewRequest(
						http.MethodPost,
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
			handler := DeleteContactHandler(tt.args.storage)
			w := httptest.NewRecorder()
			handler(w, tt.args.req())
			result := w.Result()
			if result.StatusCode != tt.wantStatusCode {
				t.Errorf("deleteContactHandler() = %v, want %v", result.StatusCode, tt.wantStatusCode)
			}
		})
	}
}
