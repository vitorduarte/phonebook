package phonebook

import (
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
	"github.com/vitorduarte/phonebook/internal/utils"
)

func GetContactsHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := utils.GetIdFromPath(r, "/contact/")
		name := r.URL.Query().Get("name")
		isGetAll := id == "" && name == ""
		isGetById := id != ""

		if isGetAll {
			GetAllContactsHandler(s)(w, r)
			return
		}
		if isGetById {
			GetContactByIdHandler(s)(w, r)
			return
		}
		FindContactsByNameHandler(s)(w, r)

	}
}
