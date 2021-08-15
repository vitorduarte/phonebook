package storage

import (
	"reflect"
	"testing"
)

func TestNewMemoryStorage(t *testing.T) {
	tests := []struct {
		name string
		want *InMemoryStorage
	}{
		{
			name: "new_memory_storage",
			want: &InMemoryStorage{
				PhoneBook: map[string]Contact{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryStorage_Create(t *testing.T) {
	tests := []struct {
		name      string
		storage   *InMemoryStorage
		arg       Contact
		wantErr   bool
		wantCount int
	}{
		{
			name:      "empty_contact_should_return_error",
			storage:   NewMemoryStorage(),
			arg:       Contact{},
			wantErr:   true,
			wantCount: 0,
		},
		{
			name:    "empty_storage",
			storage: NewMemoryStorage(),
			arg: Contact{
				Name:  "Bob",
				Phone: "999999999",
			},
			wantErr:   false,
			wantCount: 1,
		},
		{
			name: "with_one_contact_on_storage",
			storage: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					"1": {
						Id:    "1",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			arg: Contact{
				Id:    "2",
				Name:  "Alice",
				Phone: "999999999",
			},
			wantErr:   false,
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultContact, err := tt.storage.Create(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			itemCount := len(tt.storage.PhoneBook)

			if !tt.wantErr && resultContact.Id == "" {
				t.Errorf("InMemoryStorage.Create() created with empty id")
			}

			if itemCount != tt.wantCount {
				t.Errorf("InMemoryStorage.Create() itemCount %d, wantCount %d", itemCount, tt.wantCount)
			}
		})
	}
}

func TestInMemoryStorage_GetAll(t *testing.T) {
	tests := []struct {
		name       string
		storage    *InMemoryStorage
		wantResult []Contact
	}{
		{
			name:       "empty_storage_should_return_empty_slice",
			storage:    NewMemoryStorage(),
			wantResult: []Contact{},
		},
		{
			name: "storage_with_one_contact_on_storage",
			storage: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					"1": {
						Id:    "1",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			wantResult: []Contact{
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

			if len(tt.wantResult) != len(result) {
				t.Errorf("InMemoryStorage.GetAll() result lenght = %v, wantResult lenght = %v", len(result), len(tt.wantResult))
			}

			for i := range result {
				if result[i] != tt.wantResult[i] {
					t.Errorf("InMemoryStorage.GetAll() = %v, wantResult = %v", result, tt.wantResult)
				}
			}
		})
	}
}

func TestInMemoryStorage_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name       string
		storage    *InMemoryStorage
		args       args
		wantErr    bool
		wantResult Contact
	}{
		{
			name:    "empty_id_should_return_error",
			storage: NewMemoryStorage(),
			args:    args{id: ""},
			wantErr: true,
		},
		{
			name:    "inexistent_id_should_return_error",
			storage: NewMemoryStorage(),
			args:    args{id: "d38cbda4-c419-410a-8fea-2ccd2523f2b2"},
			wantErr: true,
		},
		{
			name: "existent_id_should_return_valid_contact",
			args: args{id: "d38cbda4-c419-410a-8fea-2ccd2523f2b2"},
			storage: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					"d38cbda4-c419-410a-8fea-2ccd2523f2b2": {
						Id:    "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			wantErr: false,
			wantResult: Contact{
				Id:    "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
				Name:  "Bob",
				Phone: "999999999",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.storage.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStorage.Get() error = %v, wantErr %v", err, tt.wantErr)
			}

			if result != tt.wantResult {
				t.Errorf("InMemoryStorage.Get() = %v, wantResult = %v", result, tt.wantResult)
			}
		})
	}
}

func TestInMemoryStorage_Update(t *testing.T) {
	type args struct {
		c Contact
	}
	tests := []struct {
		name         string
		m            *InMemoryStorage
		args         args
		wantResponse Contact
		wantErr      bool
	}{
		{
			name: "existent_contact",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					"d38cbda4-c419-410a-8fea-2ccd2523f2b2": {
						Id:    "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			args: args{
				c: Contact{
					Id:    "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
					Name:  "Alice",
					Phone: "999999999",
				},
			},
			wantResponse: Contact{
				Id:    "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
				Name:  "Alice",
				Phone: "999999999",
			},
			wantErr: false,
		},
		{
			name: "inexistent_contact",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{},
			},
			args: args{
				c: Contact{
					Id:    "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
					Name:  "Alice",
					Phone: "999999999",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.m.Update(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStorage.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("InMemoryStorage.Update() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestInMemoryStorage_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		m       *InMemoryStorage
		args    args
		wantErr bool
	}{
		{
			name: "existent_contact",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					"d38cbda4-c419-410a-8fea-2ccd2523f2b2": {
						Id:    "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
						Name:  "Bob",
						Phone: "999999999",
					},
				},
			},
			args: args{
				id: "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
			},
			wantErr: false,
		},
		{
			name: "inexistent_contact",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{},
			},
			args: args{
				id: "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStorage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
