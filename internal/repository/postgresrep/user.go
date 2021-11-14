package postgresrep

import (
	"context"
	"fmt"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (rep *PostgresUserRepository) Create(ctx context.Context, in domain.CreateUserInput) (domain.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password, created_at) VALUES ($1, $2, $3) RETURNING id", usersTable)

	now := time.Now().Format(time.RFC3339)
	row := rep.db.QueryRow(query, in.Login, in.Password, now)

	var id int
	if err := row.Scan(&id); err != nil {
		return domain.User{}, err
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	var user domain.User
	err := rep.db.Get(&user, query, id)
	return user, err
}

func (rep *PostgresUserRepository) GetByCredentials(ctx context.Context, in domain.LoginUserInput) (domain.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE login = $1 AND password = $2", usersTable)
	var user domain.User
	err := rep.db.Get(&user, query, in.Login, in.Password)
	return user, err
}

func (rep *PostgresUserRepository) Get(ctx context.Context, id string) (domain.User, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return domain.User{}, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	var user domain.User
	err = rep.db.Get(&user, query, intID)
	return user, err
}
