package storage

type MockFileStorage struct{}

func NewMockFileStorage() *MockFileStorage {
	s := &MockFileStorage{}
	return s
}

func (s *MockFileStorage) Load() map[string]any {
	var result map[string]any
	return result
}

func (s *MockFileStorage) Write(key string, data any) error {
	return nil
}
