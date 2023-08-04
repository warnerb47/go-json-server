package storage

type TestFileStorage struct{}

func NewTestFileStorage() *TestFileStorage {
	s := &TestFileStorage{}
	return s
}

func (s *TestFileStorage) Load() map[string]any {
	var result map[string]any
	return result
}

func (s *TestFileStorage) Write(key string, data any) error {
	jsonData := s.Load()
	jsonData[key] = data
	return nil
}
