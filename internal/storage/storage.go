package storage

import "github.com/vitorduarte/phonebook/internal/contact"

type Storage interface {
	Create(c contact.Contact) (contact.Contact, error)
	GetAll() ([]contact.Contact, error)
	Get(id string) (contact.Contact, error)
	Update(c contact.Contact) (contact.Contact, error)
	Delete(id string) error
	FindByName(name string) ([]contact.Contact, error)
	HealthCheck() bool
}
