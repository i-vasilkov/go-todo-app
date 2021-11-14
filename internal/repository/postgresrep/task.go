package postgresrep

import (
	"context"
	"fmt"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type PostgresTaskRepository struct {
	db *sqlx.DB
}

func NewPostgresTaskRepository(db *sqlx.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (rep *PostgresTaskRepository) Get(ctx context.Context, id, userId string) (domain.Task, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return domain.Task{}, err
	}

	intUserID, err := strconv.Atoi(userId)
	if err != nil {
		return domain.Task{}, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND user_id = $2", tasksTable)

	var task domain.Task
	err = rep.db.Get(&task, query, intID, intUserID)

	return task, err
}

func (rep *PostgresTaskRepository) GetAll(ctx context.Context, userId string) ([]domain.Task, error) {
	intUserID, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", tasksTable)

	tasks := make([]domain.Task, 0)
	err = rep.db.Select(&tasks, query, intUserID)

	return tasks, err
}

func (rep *PostgresTaskRepository) Create(ctx context.Context, userId string, in domain.CreateTaskInput) (domain.Task, error) {
	intUserID, err := strconv.Atoi(userId)
	if err != nil {
		return domain.Task{}, err
	}

	query := fmt.Sprintf("INSERT INTO %s (name, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", tasksTable)

	now := time.Now().Format(time.RFC3339)
	row := rep.db.QueryRow(query, in.Name, intUserID, now, now)

	var id int
	if err := row.Scan(&id); err != nil {
		return domain.Task{}, err
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND user_id = $2", tasksTable)

	var task domain.Task
	err = rep.db.Get(&task, query, id, intUserID)

	return task, err
}

func (rep *PostgresTaskRepository) Update(ctx context.Context, id, userId string, in domain.UpdateTaskInput) (domain.Task, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return domain.Task{}, err
	}

	intUserID, err := strconv.Atoi(userId)
	if err != nil {
		return domain.Task{}, err
	}

	query := fmt.Sprintf("UPDATE %s SET name = $1, updated_at = $2 WHERE id = $3 AND user_id = $4", tasksTable)
	now := time.Now().Format(time.RFC3339)
	_, err = rep.db.Exec(query, in.Name, now, intID, intUserID)
	if err != nil {
		return domain.Task{}, err
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND user_id = $2", tasksTable)

	var task domain.Task
	err = rep.db.Get(&task, query, id, intUserID)

	return task, err
}

func (rep *PostgresTaskRepository) Delete(ctx context.Context, id, userId string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	intUserID, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND user_id = $2", tasksTable)

	_, err = rep.db.Exec(query, intID, intUserID)
	return err
}
