package storage

import (
	"fmt"

	"github.com/vitorduarte/phonebook/internal/contact"
)

type InMemoryStorage struct {
	PhoneBook map[string]contact.Contact
}

func (m *InMemoryStorage) Create(c contact.Contact) (err error) {
	if c.Id == "" {
		return fmt.Errorf("Could not add contact to database, invalid id")
	}

	if _, ok := m.PhoneBook[c.Id]; ok {
		return fmt.Errorf("Could not add contact to database, id: %s already exists on database", c.Id)
	}

	m.PhoneBook[c.Id] = c
	return nil
}

func (m *InMemoryStorage) GetAll() (response []contact.Contact, err error) {
	for _, c := range m.PhoneBook {
		response = append(response, c)
	}

	return response, nil
}

func (m *InMemoryStorage) Get(id string) (response contact.Contact, err error) {
	response, ok := m.PhoneBook[id]
	if !ok {
		return response, fmt.Errorf("contact with id: %s does not exist on database", id)
	}

	return
}

func NewMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		PhoneBook: make(map[string]contact.Contact),
	}
}
