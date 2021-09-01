package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestE2ECreateContact(t *testing.T) {
	type args struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	tests := []struct {
		name           string
		args           args
		wantStatusCode int
	}{
		{
			name: "create_returns_200_when_all_good",
			args: args{
				Name:  "Bob",
				Phone: "1234567890",
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "create_returns_200_when_name_is_empty",
			args: args{
				Name:  "",
				Phone: "1234567890",
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "create_returns_400_when_name_and_phone_is_empty",
			args: args{
				Name:  "",
				Phone: "",
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "create_returns_400_when_body_json_is_empty",
			wantStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.args)
			b := bytes.NewBuffer(body)

			result, _ := http.Post("http://localhost:8080/contact", "application/json", b)

			if result.StatusCode != tt.wantStatusCode {
				t.Errorf("[POST] /contact status = %v, want %v", result.StatusCode, tt.wantStatusCode)
			}
		})
	}
}
