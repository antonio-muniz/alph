package storage

import (
	"context"
	"errors"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
)

var ErrUserNotFound = errors.New("user not found")

type Database interface {
	CreateUser(ctx context.Context, user auth.User) error
	GetUser(ctx context.Context, username string) (auth.User, error)
}
