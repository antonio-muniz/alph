package storage

import (
	"context"
	"errors"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
)

var ErrUserNotFound = errors.New("user not found")

type Database interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, username string) (model.User, error)
}
