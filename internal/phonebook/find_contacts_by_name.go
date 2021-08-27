package phonebook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
)

func FindContactsByName(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		contacts, err := s.FindByName(name)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to get contact: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		response := make([]interface{}, len(contacts))
		for i, c := range contacts {
			response[i] = c
		}

		json.NewEncoder(w).Encode(response)
	}
}
