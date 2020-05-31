package memory

import (
	"context"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
)

type database struct {
	users map[string]model.User
}

func NewDatabase() storage.Database {
	return &database{
		users: make(map[string]model.User),
	}
}

func (d *database) CreateUser(ctx context.Context, user model.User) error {
	d.users[user.Username] = user
	return nil
}

func (d *database) GetUser(ctx context.Context, username string) (model.User, error) {
	user, found := d.users[username]
	if !found {
		return model.User{}, storage.ErrUserNotFound
	}
	return user, nil
}
