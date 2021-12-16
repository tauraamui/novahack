package database

type DBConn interface {
	Create(string, Entry) string
	Find(string) Entry
	Replace(Entry)
	Close() error
}

func NewConn() DBConn {
	return &db{}
}
