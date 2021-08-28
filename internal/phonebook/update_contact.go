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

func UpdateContactHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := utils.GetIdFromPath(r, "/contact/")
		if id == "" {
			msg := fmt.Sprintf("update requires an id")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		if r.Body == http.NoBody {
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

		response, appErr := UpdateContact(s, contact)
		if appErr.Error != nil {
			msg := fmt.Sprintf("could not update contact: %v", appErr.Error)
			w.WriteHeader(appErr.StatusCode)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}

func UpdateContact(s storage.Storage, c contact.Contact) (response contact.Contact, appErr utils.AppError) {
	receivedContact, err := s.Get(c.Id)
	if err != nil {
		appErr = utils.AppError{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
		return
	}

	// Check if the contact is empty
	if (receivedContact == contact.Contact{}) {
		appErr = utils.AppError{
			Error:      fmt.Errorf("could not find contact %s", c.Id),
			StatusCode: http.StatusNotFound,
		}
		return
	}

	// Name and phone number are required
	if c.Name == "" && c.Phone == "" {
		err := errors.New("name and phone cannot be empty")
		appErr = utils.AppError{Error: err, StatusCode: http.StatusBadRequest}
		return
	}

	response, err = s.Update(c)
	if err != nil {
		appErr = utils.AppError{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return
}
