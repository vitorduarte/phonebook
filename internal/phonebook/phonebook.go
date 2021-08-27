package phonebook

import (
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
)

var badRequestResponse = []byte(`{"message":"bad request"}`)
var methodNotAllowedResponse = []byte(`{"message":"method not allowed"}`)
var notFoundResponse = []byte(`{"message":"contact not foud"}`)

func ContactHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetContactsHandler(s)(w, r)
		case http.MethodPost:
			CreateContactHandler(s)(w, r)
		case http.MethodPut:
			UpdateContactHandler(s)(w, r)
		case http.MethodDelete:
			DeleteContactHandler(s)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(methodNotAllowedResponse)
		}
	}
}
