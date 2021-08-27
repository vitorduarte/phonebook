package phonebook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
)

func GetContactById(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		contact, err := s.Get(id)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to get contact: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		json.NewEncoder(w).Encode(contact)
		return
	}
}
