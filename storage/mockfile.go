package storage

import "encoding/json"

type MockFileStorage struct{}

var jsonString = `{
	"hospitals": {
		{"adress": "Dakar", "id": "1", "name": "Hôpital Fann"},
		{"adress": "Dakar", "id": "2", "name": "Dalal Jam"},
	},
	"schools": {
		{"adress": "Dakar", "id": "1", "name": "Lycée Jean Mermoz"},
	},
	"users": {
		{"age": 23, "id": "1", "name": "Elliot", "type": "Reader"},
		{"age": 17, "id": "2", "name": "Fraser", "type": "Author"},
	},
}`

func NewMockFileStorage() *MockFileStorage {
	s := &MockFileStorage{}
	return s
}

func (s *MockFileStorage) Load() map[string]any {
	var data map[string]any
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil
	}
	return data
}

func (s *MockFileStorage) Write(key string, data any) error {
	return nil
}
