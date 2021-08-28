package phonebook

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vitorduarte/phonebook/internal/storage"
	"github.com/vitorduarte/phonebook/internal/utils"
)

func DeleteContactHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := strings.TrimPrefix(r.URL.Path, "/contact/")
		if id == "" {
			msg := fmt.Sprintf("delete requires an id")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		appErr := DeleteContact(s, id)
		if appErr.Error != nil {
			msg := fmt.Sprintf("an error occurred while trying to delete contact: %v", appErr.Error)
			w.WriteHeader(appErr.StatusCode)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}
	}
}

func DeleteContact(s storage.Storage, id string) (appErr utils.AppError) {
	_, err := s.Get(id)
	if err != nil {
		return utils.AppError{
			Error:      err,
			StatusCode: http.StatusNotFound,
		}
	}

	err = s.Delete(id)
	if err != nil {
		return utils.AppError{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return
}
