package phonebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/contact"
	"github.com/vitorduarte/phonebook/internal/storage"
	"github.com/vitorduarte/phonebook/internal/utils"
)

func CreateContactHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Body == http.NoBody {
			msg := fmt.Sprintf("create requires a request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		var contact contact.Contact
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			msg := fmt.Sprintf("could not create contact: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		response, appErr := CreateContact(s, contact)

		if appErr.Error != nil {
			msg := fmt.Sprintf("could not create contact: %v", appErr.Error)
			w.WriteHeader(appErr.StatusCode)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func CreateContact(s storage.Storage, c contact.Contact) (contactResponse contact.Contact, appErr utils.AppError) {
	// Name and phone number are required
	if c.Name == "" && c.Phone == "" {
		err := errors.New("name and phone cannot be empty")
		appErr = utils.AppError{Error: err, StatusCode: http.StatusBadRequest}
		return
	}

	contactResponse, err := s.Create(c)
	if err != nil {
		appErr = utils.AppError{Error: err, StatusCode: http.StatusInternalServerError}
		return
	}

	return
}
