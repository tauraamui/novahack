package database

import "github.com/google/uuid"

type db map[string]interface{}

func (m db) Create(keyPrefix string, v interface{}) string {
	id := uuid.NewString()
	m[id+keyPrefix] = v
	return id
}

func (m db) Find(id string) interface{} {
	return m[id]
}

func (m db) Close() error {
	for k := range m {
		delete(m, k)
	}
	return nil
}
