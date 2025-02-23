package memberships

import (
	"context"
	"database/sql"

	"github.com/IndraNurfa/fastcampus/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error) {
	query := `SELECT id, email, username, password, created_at, created_by, updated_at, updated_by FROM users WHERE email = ? OR username = ? OR id = ?`
	row := r.db.QueryRowContext(ctx, query, email, username, userID)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Username, &response.Password, &response.CreatedAt, &response.CreatedBy, &response.UpdatedAt, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateUser(ctx context.Context, model memberships.UserModel) error {
	query := `UPDATE users SET email = ?, username = ?, updated_at = ?, updated_by = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Username, model.UpdatedAt, model.UpdatedBy, model.ID)
	if err != nil {
		return err
	}
	return nil
}
