package phonebook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/contact"
	"github.com/vitorduarte/phonebook/internal/storage"
	"github.com/vitorduarte/phonebook/internal/utils"
)

func GetContactByIdHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := utils.GetIdFromPath(r, "/contact/")
		response, appErr := GetContactById(s, id)
		if appErr.Error != nil {
			msg := fmt.Sprintf("could not find contacts by id: %v", appErr.Error)
			w.WriteHeader(appErr.StatusCode)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}

func GetContactById(s storage.Storage, id string) (response contact.Contact, appErr utils.AppError) {
	response, err := s.Get(id)
	if err != nil {
		appErr = utils.AppError{Error: err, StatusCode: http.StatusInternalServerError}
		return
	}

	if (response == contact.Contact{}) {
		appErr = utils.AppError{
			Error:      fmt.Errorf("could not find contact %s", id),
			StatusCode: http.StatusNotFound,
		}
		return
	}
	return
}
