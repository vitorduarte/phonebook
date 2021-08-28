package phonebook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/contact"
	"github.com/vitorduarte/phonebook/internal/storage"
)

func UpdateContactHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")
		if id == "" {
			msg := fmt.Sprintf("update requires an id")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		if r.Body == nil {
			msg := fmt.Sprintf("update requires a request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		var contact contact.Contact
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			msg := fmt.Sprintf("could not update contact: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		contact.Id = id
		response, err := s.Update(contact)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to update contact: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}
