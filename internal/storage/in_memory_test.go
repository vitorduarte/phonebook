package storage

import (
	"testing"

	"github.com/vitorduarte/phonebook/internal/contact"
)

func TestInMemoryStorage_Create(t *testing.T) {
	tests := []struct {
		name      string
		storage   *InMemoryStorage
		arg       contact.Contact
		wantErr   bool
		wantCount int
	}{
		{
			name:      "empty_contact_should_return_error",
			storage:   NewMemoryStorage(),
			arg:       contact.Contact{},
			wantErr:   true,
			wantCount: 0,
		},
		{
			name:    "empty_storage",
			storage: NewMemoryStorage(),
			arg: contact.Contact{
				Id:    "1",
				Name:  "Bob",
				Phone: "999999999",
			},
			wantErr:   false,
			wantCount: 1,
		},
		{
			name: "with_one_contact_on_storage",
			storage: &InMemoryStorage{
				PhoneBook: map[string]contact.Contact{
					"1": {
						Id:    "1",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			arg: contact.Contact{
				Id:    "2",
				Name:  "Alice",
				Phone: "999999999",
			},
			wantErr:   false,
			wantCount: 2,
		},
		{
			name: "with_id_already_on_storage",
			storage: &InMemoryStorage{
				PhoneBook: map[string]contact.Contact{
					"1": {
						Id:    "1",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			arg: contact.Contact{
				Id:    "1",
				Name:  "Alice",
				Phone: "999999999",
			},
			wantErr:   true,
			wantCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.storage.Create(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			itemCount := len(tt.storage.PhoneBook)

			if itemCount != tt.wantCount {
				t.Errorf("itemCount %d, wantCount %d", itemCount, tt.wantCount)
			}
		})
	}
}

func TestInMemoryStorage_GetAll(t *testing.T) {
	tests := []struct {
		name         string
		storage      *InMemoryStorage
		wantedResult []contact.Contact
	}{
		{
			name:         "empty_storage_should_return_empty_slice",
			storage:      NewMemoryStorage(),
			wantedResult: []contact.Contact{},
		},
		{
			name: "storage_with_one_contact_on_storage",
			storage: &InMemoryStorage{
				PhoneBook: map[string]contact.Contact{
					"1": {
						Id:    "1",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			wantedResult: []contact.Contact{
				{
					Id:    "1",
					Name:  "Bob",
					Phone: "999999999",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.storage.GetAll()
			if err != nil {
				t.Errorf("error = %v", err)
			}

			if len(tt.wantedResult) != len(result) {
				t.Errorf("result lenght = %v, wantedResult lenght = %v", len(result), len(tt.wantedResult))
			}

			for i := range result {
				if result[i] != tt.wantedResult[i] {
					t.Errorf("result = %v, wantedResult = %v", result, tt.wantedResult)
				}
			}
		})
	}
}

func TestInMemoryStorage_Get(t *testing.T) {
	tests := []struct {
		name         string
		storage      *InMemoryStorage
		arg          string
		wantErr      bool
		wantedResult contact.Contact
	}{
		{
			name:    "empty_id_should_return_error",
			storage: NewMemoryStorage(),
			arg:     "",
			wantErr: true,
		},
		{
			name:    "inexistent_id_should_return_error",
			storage: NewMemoryStorage(),
			arg:     "1",
			wantErr: true,
		},
		{
			name: "existent_id_should_return_valid_contact",
			arg:  "1",
			storage: &InMemoryStorage{
				PhoneBook: map[string]contact.Contact{
					"1": {
						Id:    "1",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			wantErr: false,
			wantedResult: contact.Contact{
				Id:    "1",
				Name:  "Bob",
				Phone: "999999999",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.storage.Get(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if result != tt.wantedResult {
				t.Errorf("result = %v, wantedResult = %v", result, tt.wantedResult)
			}
		})
	}
}
