package user

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yurixtugal/Events/infrastructure/postgres"
	"github.com/yurixtugal/Events/model"
)

const table = "users"

var fields = []string{
	"id",
	"email",
	"password",
	"details",
	"created_at",
	"updated_at",
}

var (
	psqlInsert       = postgres.BuildSQLInsert(table, fields)
	psqlSelectGetAll = postgres.BuildSQLSelect(table, fields)
)

type User struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) User {
	return User{db: db}
}

func (u User) Create(m *model.User) error {
	_, err := u.db.Exec(context.Background(), psqlInsert,
		m.Email,
		m.Password,
		m.IsAdmin,
		m.Details,
		m.CreateAt,
		postgres.Int64ToNull(m.UpdateAt))

	if err != nil {
		return err
	}

	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	query := psqlSelectGetAll + " WHERE email = $1"
	row := u.db.QueryRow(context.Background(), query, email)
	return u.scanRow(row, true)
}

func (u User) GetAll() (model.Users, error) {
	rows, err := u.db.Query(context.Background(), psqlSelectGetAll)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ms := model.Users{}

	for rows.Next() {
		m, err := u.scanRow(rows, false)

		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (u User) scanRow(s pgx.Row, withPassword bool) (model.User, error) {

	m := model.User{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Email,
		&m.Password,
		&m.IsAdmin,
		&m.Details,
		&m.CreateAt,
		&updatedAtNull)

	if err != nil {
		return m, err
	}
	m.UpdateAt = updatedAtNull.Int64

	if !withPassword {
		m.Password = ""
	}

	return m, nil
}
