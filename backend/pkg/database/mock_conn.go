package database

import "github.com/google/uuid"

type db map[string]interface{}

func (m db) Create(keyPrefix string, v interface{}) string {
	id := uuid.NewString()
	m[keyPrefix+id] = v
	return id
}

func (m db) Find(keyPrefix, id string) interface{} {
	return m[keyPrefix+id]
}

func (m db) Close() error {
	for k := range m {
		delete(m, k)
	}
	return nil
}
