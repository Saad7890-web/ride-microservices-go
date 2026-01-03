package model

import "time"

type User struct {
    ID           int64     `db:"id"`
    Username     string    `db:"username"`
    Email        string    `db:"email"`
    Phone        string    `db:"phone"`
    PasswordHash string    `db:"password_hash"`
    Role         string    `db:"role"`
    Status       string    `db:"status"`
    CreatedAt    time.Time `db:"created_at"`
    UpdatedAt    time.Time `db:"updated_at"`
}
