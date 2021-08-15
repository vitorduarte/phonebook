package storage

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var firstContact, secondContact = Contact{
	Id:    "d38cbda4-c419-410a-8fea-2ccd2523f2b2",
	Name:  "Bob Silva",
	Phone: "999999999",
}, Contact{
	Id:    "2c928a05-21d4-445b-8644-fcc2cc40371a",
	Name:  "Alice Silva",
	Phone: "999999998",
}

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
			if got := NewMemoryStorage(); !cmp.Equal(got, tt.want) {
				t.Errorf("NewMemoryStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryStorage_Create(t *testing.T) {
	type args struct {
		c Contact
	}
	tests := []struct {
		name      string
		m         *InMemoryStorage
		args      args
		wantErr   bool
		wantCount int
	}{
		{
			name: "empty_contact_should_return_error",
			m:    NewMemoryStorage(),
			args: args{
				c: Contact{},
			},
			wantErr:   true,
			wantCount: 0,
		},
		{
			name: "empty_storage",
			m:    NewMemoryStorage(),
			args: args{c: Contact{
				Name:  "Bob",
				Phone: "999999999",
			},
			},
			wantErr:   false,
			wantCount: 1,
		},
		{
			name: "storage_with_one_contact_should_have_two_after_create",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					firstContact.Id: firstContact,
				},
			},
			args: args{
				c: Contact{
					Name:  "Alice",
					Phone: "999999999",
				},
			},
			wantErr:   false,
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultContact, err := tt.m.Create(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && resultContact.Id == "" {
				t.Errorf("InMemoryStorage.Create() created with empty id")
			}

			itemCount := len(tt.m.PhoneBook)
			if itemCount != tt.wantCount {
				t.Errorf("InMemoryStorage.Create() itemCount %d, wantCount %d", itemCount, tt.wantCount)
			}
		})
	}
}

func TestInMemoryStorage_GetAll(t *testing.T) {
	tests := []struct {
		name       string
		m          *InMemoryStorage
		wantResult []Contact
	}{
		{
			name:       "empty_storage_should_return_empty_slice",
			m:          NewMemoryStorage(),
			wantResult: []Contact{},
		},
		{
			name: "storage_with_one_contact",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					firstContact.Id: firstContact,
				},
			},
			wantResult: []Contact{
				firstContact,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.m.GetAll()
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
		m          *InMemoryStorage
		args       args
		wantErr    bool
		wantResult Contact
	}{
		{
			name:    "empty_id_should_return_error",
			m:       NewMemoryStorage(),
			args:    args{id: ""},
			wantErr: true,
		},
		{
			name:    "inexistent_id_should_return_error",
			m:       NewMemoryStorage(),
			args:    args{id: "d38cbda4-c419-410a-8fea-2ccd2523f2b2"},
			wantErr: true,
		},
		{
			name: "existent_id_should_return_valid_contact",
			args: args{id: firstContact.Id},
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					firstContact.Id: firstContact,
				},
			},
			wantErr:    false,
			wantResult: firstContact,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.m.Get(tt.args.id)
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
					firstContact.Id: firstContact,
				},
			},
			args: args{
				c: Contact{
					Id:    firstContact.Id,
					Name:  "Alice",
					Phone: "999999999",
				},
			},
			wantResponse: Contact{
				Id:    firstContact.Id,
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
			if !cmp.Equal(gotResponse, tt.wantResponse) {
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
		name      string
		m         *InMemoryStorage
		args      args
		wantErr   bool
		wantCount int
	}{
		{
			name: "existent_contact",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					firstContact.Id: firstContact,
				},
			},
			args: args{
				id: firstContact.Id,
			},
			wantErr:   false,
			wantCount: 0,
		},
		{
			name: "inexistent_contact",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					firstContact.Id: firstContact,
				},
			},
			args: args{
				id: "2ccd2523f2b2",
			},
			wantErr:   true,
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStorage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			itemCount := len(tt.m.PhoneBook)
			if itemCount != tt.wantCount {
				t.Errorf("InMemoryStorage.Delete() itemCount %d, wantCount %d", itemCount, tt.wantCount)
			}
		})
	}
}

func TestInMemoryStorage_FindByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name         string
		m            *InMemoryStorage
		args         args
		wantResponse []Contact
		wantErr      bool
	}{
		{
			name: "should_return_unique_contact",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					firstContact.Id: firstContact,
				},
			},
			args: args{
				name: firstContact.Name[:2],
			},
			wantResponse: []Contact{firstContact},
			wantErr:      false,
		},
		{
			name: "should_return_multiple_contacts",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					firstContact.Id:  firstContact,
					secondContact.Id: secondContact,
				},
			},
			args: args{
				name: "Silva",
			},
			wantResponse: []Contact{firstContact, secondContact},
			wantErr:      false,
		},
		{
			name: "should_return_none",
			m: &InMemoryStorage{
				PhoneBook: map[string]Contact{
					firstContact.Id: firstContact,
				},
			},
			args: args{
				name: "NotExistent",
			},
			wantResponse: []Contact{},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.m.FindByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStorage.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(fmt.Sprint(gotResponse) == fmt.Sprint(tt.wantResponse)) {
				t.Errorf("InMemoryStorage.FindByName() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
