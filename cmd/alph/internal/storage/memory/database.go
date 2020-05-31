package memory

import (
	"context"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
)

type database struct {
	users map[string]auth.User
}

func NewDatabase() storage.Database {
	return &database{
		users: make(map[string]auth.User),
	}
}

func (d *database) CreateUser(ctx context.Context, user auth.User) error {
	d.users[user.Username] = user
	return nil
}

func (d *database) GetUser(ctx context.Context, username string) (auth.User, error) {
	user, found := d.users[username]
	if !found {
		return auth.User{}, storage.ErrUserNotFound
	}
	return user, nil
}
