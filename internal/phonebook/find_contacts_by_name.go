package phonebook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/contact"
	"github.com/vitorduarte/phonebook/internal/storage"
	"github.com/vitorduarte/phonebook/internal/utils"
)

func FindContactsByNameHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		response, appErr := FuindContactsByName(s, name)
		if appErr.Error != nil {
			msg := fmt.Sprintf("could not find contacts by name: %v", appErr.Error)
			w.WriteHeader(appErr.StatusCode)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}

func FuindContactsByName(s storage.Storage, name string) (response []contact.Contact, appErr utils.AppError) {
	contacts, err := s.FindByName(name)
	if err != nil {
		appErr = utils.AppError{Error: err, StatusCode: http.StatusInternalServerError}
		return
	}

	response = make([]contact.Contact, len(contacts))
	for i, c := range contacts {
		response[i] = c
	}

	return
}
