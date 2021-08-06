package storage

import "github.com/vitorduarte/phonebook/internal/contact"

type iStorage interface {
	Create(c contact.Contact) error
	GetAll() ([]contact.Contact, error)
	Get(id string) error
	Update(id string, c contact.Contact) error
	Delete(id string) error
}
