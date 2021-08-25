package storage

type Contact struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Storage interface {
	Create(c Contact) (Contact, error)
	GetAll() ([]Contact, error)
	Get(id string) (Contact, error)
	Update(c Contact) (Contact, error)
	Delete(id string) error
	FindByName(name string) ([]Contact, error)
	HealthCheck() bool
}
