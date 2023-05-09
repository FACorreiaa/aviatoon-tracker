package user

import (
	"context"
	"github.com/FACorreiaa/aviatoon-tracker/internal/structs"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (q *Repository) GetUsers() ([]structs.User, error) {

	// Define users variable.
	var users []structs.User

	// Send query to database.
	rows, err := q.db.Query(context.Background(), `SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser func for getting one user by given ID.
func (q *Repository) GetUser(id uuid.UUID) (structs.User, error) {
	// Define user variable.
	var user structs.User

	// Send query to database.
	if err := q.db.QueryRow(context.Background(), `SELECT * FROM users WHERE id = $1`, id).Scan(&user.ID, &user.UserName, &user.Email); err != nil {
		return structs.User{}, err
	}

	return user, nil
}

// CreateUser func for creating user by given User object.
func (q *Repository) CreateUser(u *structs.User) error {
	// Send query to database
	if _, err := q.db.Exec(context.Background(),
		`INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6)`,
		u.ID,
		u.CreatedAt,
		u.UpdatedAt,
		u.Email,
		u.UserStatus,
		u.UserAttrs,
	); err != nil {
		return err
	}

	return nil
}

// UpdateUser func for updating user by given User object.
func (q *Repository) UpdateUser(u *structs.User) error {
	// Send query to database.
	if _, err := q.db.Exec(context.Background(),
		`UPDATE users SET updated_at = $2, email = $3, user_attrs = $4 WHERE id = $1`,
		u.ID,
		u.UpdatedAt,
		u.Email,
		u.UserAttrs,
	); err != nil {
		return err
	}

	return nil
}

// DeleteUser func for delete user by given ID.
func (q *Repository) DeleteUser(id uuid.UUID) error {
	// Send query to database.

	if _, err := q.db.Exec(context.Background(), `DELETE FROM users WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}
