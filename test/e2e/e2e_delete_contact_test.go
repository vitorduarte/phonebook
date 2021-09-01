package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestE2EDeleteContact(t *testing.T) {
	type args struct {
		Id string
	}

	type Contact struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	tests := []struct {
		name           string
		args           args
		initContact    Contact
		wantStatusCode int
	}{
		{
			name: "delete_returns_200_when_it_exist",
			initContact: Contact{
				Name:  "Bob",
				Phone: "1234567890",
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "delete_returns_404_when_does_not_exist",
			args:           args{Id: "61ffffffff7ab995fba84c39"},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "delete_returns_400_when_id_is_not_sent",
			wantStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var id string
			id = tt.args.Id

			if tt.initContact != (Contact{}) {
				body, _ := json.Marshal(tt.initContact)
				b := bytes.NewBuffer(body)

				result, _ := http.Post("http://localhost:8080/contact", "application/json", b)
				if result.StatusCode == http.StatusOK {
					var contact Contact
					body, _ := ioutil.ReadAll(result.Body)
					json.Unmarshal(body, &contact)
					id = contact.Id
				}
			}

			client := &http.Client{}
			deleteUrl := fmt.Sprintf("http://localhost:8080/contact/%s", id)
			request, _ := http.NewRequest(http.MethodDelete, deleteUrl, nil)
			result, _ := client.Do(request)

			if result.StatusCode != tt.wantStatusCode {
				t.Errorf("[DELETE] /contact/%v status = %v, want %v", id, result.StatusCode, tt.wantStatusCode)
			}
		})
	}
}
