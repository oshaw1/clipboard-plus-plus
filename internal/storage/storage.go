package storage

type Storer interface {
	Get(slot int) (string, error)
	Set(slot int, content string) error
}
