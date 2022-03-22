package storage

type Storage interface {
	Get(id string, title string) string
	GetAll(id string) bool
	Set(id string, title string, password string) map[string]string
	Update(id string, title string, password string) bool
	Remove(id string, title string) bool
}
