package phonebook

import (
	"fmt"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
)

func DeleteContactHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := r.URL.Query().Get("id")
		if id == "" {
			msg := fmt.Sprintf("delete requires an id")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		err := s.Delete(id)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to delete contact: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}
	}
}
