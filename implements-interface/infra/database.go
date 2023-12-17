package infra

type DB struct {
}

type JTDB interface {
	DBInterface
}

type INVDB interface {
	DBInterface
}

type DBInterface interface {
	Query(ID string) string
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Query(ID string) string {
	return "ID"
}
