package storage

type Contact struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type iStorage interface {
	Create(c Contact) error
	GetAll() ([]Contact, error)
	Get(id string) error
	Update(id string, c Contact) error
	Delete(id string) error
}
