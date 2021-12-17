package database

import "github.com/google/uuid"

type Entry interface {
	ID() string
}

type db map[string]Entry

func (m db) Create(keyPrefix string, v Entry) string {
	id := uuid.NewString()
	m[keyPrefix+id] = v
	return id
}

func (m db) Find(id string) Entry {
	return m[id]
}

func (m db) Replace(v Entry) {
	if _, ok := m[v.ID()]; ok {
		m[v.ID()] = v
	}
}

func (m db) Close() error {
	for k := range m {
		delete(m, k)
	}
	return nil
}
