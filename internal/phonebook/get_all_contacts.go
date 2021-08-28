package phonebook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
)

func GetAllContacts(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contacts, err := s.GetAll()
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to get all contacts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		response := make([]interface{}, len(contacts))
		for i, c := range contacts {
			response[i] = c
		}

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
		return
	}
}
