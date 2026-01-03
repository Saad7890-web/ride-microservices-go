package repository

import (
	"context"
	"ride-microservices-go/internal/model"
)


type UserRepository interface {
    CreateUser(ctx context.Context, user *model.User) error
    GetUserByID(ctx context.Context, id int64) (*model.User, error)
    GetUserByEmail(ctx context.Context, email string) (*model.User, error)
    GetUserByPhone(ctx context.Context, phone string) (*model.User, error)
    UpdateUser(ctx context.Context, user *model.User) error
    DeleteUser(ctx context.Context, id int64) error
}
