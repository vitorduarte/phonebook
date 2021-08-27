package storage

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/vitorduarte/phonebook/internal/contact"
)

type InMemoryStorage struct {
	PhoneBook map[string]contact.Contact
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		PhoneBook: make(map[string]contact.Contact),
	}
}

func (m *InMemoryStorage) Create(c contact.Contact) (contactResponse contact.Contact, err error) {
	if c.Name == "" && c.Phone == "" {
		err = fmt.Errorf("Name and phone cannot be empty")
		return
	}

	c.Id = uuid.New().String()
	for {
		if _, ok := m.PhoneBook[c.Id]; !ok {
			break
		}
		c.Id = uuid.New().String()
	}

	m.PhoneBook[c.Id] = c
	contactResponse = c
	return
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

func (m *InMemoryStorage) Update(c contact.Contact) (response contact.Contact, err error) {
	if _, ok := m.PhoneBook[c.Id]; !ok {
		err = fmt.Errorf("contact with id: %s does not exist on database", c.Id)
		return
	}

	m.PhoneBook[c.Id] = c
	response = c
	return
}

func (m *InMemoryStorage) Delete(id string) error {
	if _, ok := m.PhoneBook[id]; !ok {
		return fmt.Errorf("contact with id: %s does not exist on database", id)
	}

	delete(m.PhoneBook, id)
	return nil
}

func (m *InMemoryStorage) FindByName(name string) (response []contact.Contact, err error) {
	for _, c := range m.PhoneBook {
		if strings.Contains(c.Name, name) {
			response = append(response, c)
		}
	}

	return
}

func (m *InMemoryStorage) HealthCheck() (bool, error) {
	return true, nil
}
