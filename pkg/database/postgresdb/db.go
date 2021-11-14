package postgresdb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(con *Connection) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			con.Host, con.Port, con.Username, con.DBName, con.Password, con.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
