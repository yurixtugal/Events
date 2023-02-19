package user

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yurixtugal/Events/model"
)

var (
	psqlIsert        = "INSERT INTO users (id, email, password, is_admin, details, created_at) VALUES ($1,$2,$3,$4,$5)"
	psqlSelectGetAll = "SELECT id, email, password, details, created_at, updated_at FROM users"
)

type User struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) User {
	return User{db: db}
}

func (u User) Create(m *model.User) error {
	_, err := u.db.Exec(context.Background(), psqlIsert,
		m.Email,
		m.Password,
		m.IsAdmin,
		m.Details,
		m.CreateAt)

	if err != nil {
		return err
	}

	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	query := psqlSelectGetAll + " WHERE email = $1"
	row := u.db.QueryRow(context.Background(), query, email)
	return u.scanRow(row)
}

func (u User) GetAll() (model.Users, error) {
	rows, err := u.db.Query(context.Background(), psqlSelectGetAll)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ms := model.Users{}

	for rows.Next() {
		m, err := u.scanRow(rows)

		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (u User) scanRow(s pgx.Row) (model.User, error) {

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
	return m, nil
}
