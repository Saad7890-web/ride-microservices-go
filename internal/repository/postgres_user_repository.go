package repository

import (
	"context"
	"errors"
	"ride-microservices-go/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
    db *pgxpool.Pool
}


func NewPostgresUserRepository(db *pgxpool.Pool) UserRepository {
    return &PostgresUserRepository{
        db: db,
    }
}


func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *model.User) error {
    query := `
        INSERT INTO users (username, email, phone, password_hash, role, status)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at
    `

    row := r.db.QueryRow(ctx, query,
        user.Username,
        user.Email,
        user.Phone,
        user.PasswordHash,
        user.Role,
        user.Status,
    )

    err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return err
    }

    return nil
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
    query := `SELECT id, username, email, phone, password_hash, role, status, created_at, updated_at FROM users WHERE id=$1`

    user := &model.User{}
    err := r.db.QueryRow(ctx, query, id).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.Phone,
        &user.PasswordHash,
        &user.Role,
        &user.Status,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return user, nil
}


func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
    query := `SELECT id, username, email, phone, password_hash, role, status, created_at, updated_at FROM users WHERE email=$1`

    user := &model.User{}
    err := r.db.QueryRow(ctx, query, email).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.Phone,
        &user.PasswordHash,
        &user.Role,
        &user.Status,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return user, nil
}


func (r *PostgresUserRepository) GetUserByPhone(ctx context.Context, phone string) (*model.User, error) {
    query := `SELECT id, username, email, phone, password_hash, role, status, created_at, updated_at FROM users WHERE phone=$1`

    user := &model.User{}
    err := r.db.QueryRow(ctx, query, phone).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.Phone,
        &user.PasswordHash,
        &user.Role,
        &user.Status,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return user, nil
}


func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *model.User) error {
    query := `
        UPDATE users
        SET username=$1, email=$2, phone=$3, password_hash=$4, role=$5, status=$6, updated_at=NOW()
        WHERE id=$7
    `
    cmdTag, err := r.db.Exec(ctx, query,
        user.Username,
        user.Email,
        user.Phone,
        user.PasswordHash,
        user.Role,
        user.Status,
        user.ID,
    )
    if err != nil {
        return err
    }

    if cmdTag.RowsAffected() == 0 {
        return errors.New("user not found")
    }

    return nil
}


func (r *PostgresUserRepository) DeleteUser(ctx context.Context, id int64) error {
    query := `DELETE FROM users WHERE id=$1`
    cmdTag, err := r.db.Exec(ctx, query, id)
    if err != nil {
        return err
    }

    if cmdTag.RowsAffected() == 0 {
        return errors.New("user not found")
    }

    return nil
}
