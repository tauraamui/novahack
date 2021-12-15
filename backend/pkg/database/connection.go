package database

type DBConn interface {
	Create(string, interface{}) string
	Find(string) interface{}
	Close() error
}

func NewConn() DBConn {
	return &db{}
}
