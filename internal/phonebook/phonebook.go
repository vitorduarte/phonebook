package phonebook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vitorduarte/phonebook/internal/storage"
)

var badRequestResponse = []byte(`{"message":"bad request"}`)
var methodNotAllowedResponse = []byte(`{"message":"method not allowed"}`)
var notFoundResponse = []byte(`{"message":"contact not foud"}`)

func Contact(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetContacts(s)(w, r)
		case http.MethodPost:
			CreateContact(s)(w, r)
		case http.MethodPut:
			UpdateContact(s)(w, r)
		case http.MethodDelete:
			DeleteContact(s)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(methodNotAllowedResponse)
		}
	}
}

func CreateContact(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Body == nil {
			log.Println("create requires a request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(badRequestResponse)
			return
		}

		var contact storage.Contact
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			msg := fmt.Sprintf("could not create contact: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(badRequestResponse)
			return
		}

		response, err := s.Create(contact)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to create contact: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func GetContacts(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")
		if id == "" {
			contacts, err := s.GetAll()
			if err != nil {
				msg := fmt.Sprintf("an error occurred while trying to get all contacts: %v", err)
				log.Println(msg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
				return
			}

			response := make([]interface{}, len(contacts))
			for i, c := range contacts {
				response[i] = c
			}

			json.NewEncoder(w).Encode(response)
			return
		}

		contact, err := s.Get(id)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to get contact: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		json.NewEncoder(w).Encode(contact)
	}
}

func UpdateContact(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")
		if id == "" {
			log.Println("update requires an id")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(badRequestResponse)
			return
		}

		if r.Body == nil {
			log.Println("update requires a request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(badRequestResponse)
			return
		}

		var contact storage.Contact
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			msg := fmt.Sprintf("could not update contact: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(badRequestResponse)
			return
		}

		contact.Id = id
		response, err := s.Update(contact)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to update contact: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func DeleteContact(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")
		if id == "" {
			log.Println("delete requires an id")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(badRequestResponse)
			return
		}

		err := s.Delete(id)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to delete contact: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
			return
		}
	}
}
