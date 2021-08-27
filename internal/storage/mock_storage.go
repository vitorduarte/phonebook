package storage

import "github.com/vitorduarte/phonebook/internal/contact"

type MockStorage struct {
	err         error
	contact     contact.Contact
	contacts    []contact.Contact
	healthCheck bool
}

func (m *MockStorage) Create(c contact.Contact) (contact.Contact, error) {
	return m.contact, m.err
}

func (m *MockStorage) GetAll() ([]contact.Contact, error) {
	return m.contacts, m.err
}

func (m *MockStorage) Get(id string) (contact.Contact, error) {
	return m.contact, m.err
}

func (m *MockStorage) Update(c contact.Contact) (contact.Contact, error) {
	return m.contact, m.err
}

func (m *MockStorage) Delete(id string) error {
	return m.err
}

func (m *MockStorage) FindByName(name string) ([]contact.Contact, error) {
	return m.contacts, m.err
}

func (m *MockStorage) HealthCheck() bool {
	return m.healthCheck
}
