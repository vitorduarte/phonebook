package phonebook

import (
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
)

func GetContactsHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")
		isGetAll := id == "" && name == ""
		isGetById := id != ""

		if isGetAll {
			GetAllContacts(s)(w, r)
			return
		}
		if isGetById {
			GetContactById(s)(w, r)
			return
		}
		FindContactsByNameHandler(s)(w, r)

	}
}
