package footballrepository

import "database/sql"

func NewFootballRepository(db *sql.DB) FootballRepository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

type FootballRepository interface {
	GetTablePernambucanoA1() error
}
