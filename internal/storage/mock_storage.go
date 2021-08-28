package storage

import "github.com/vitorduarte/phonebook/internal/contact"

type MockStorage struct {
	Error    error
	Contact  contact.Contact
	Contacts []contact.Contact
	Health   bool
}

func (m *MockStorage) Create(c contact.Contact) (contact.Contact, error) {
	return m.Contact, m.Error
}

func (m *MockStorage) GetAll() ([]contact.Contact, error) {
	return m.Contacts, m.Error
}

func (m *MockStorage) Get(id string) (contact.Contact, error) {
	return m.Contact, m.Error
}

func (m *MockStorage) Update(c contact.Contact) (contact.Contact, error) {
	return m.Contact, m.Error
}

func (m *MockStorage) Delete(id string) error {
	return m.Error
}

func (m *MockStorage) FindByName(name string) ([]contact.Contact, error) {
	return m.Contacts, m.Error
}

func (m *MockStorage) HealthCheck() bool {
	return m.Health
}
