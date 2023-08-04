package storage

type Storage interface {
	Load() map[string]any
	Write(key string, data any) error
}
