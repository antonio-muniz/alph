package controller

import (
	"context"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/antonio-muniz/alph/pkg/system"
)

func CreateUser(ctx context.Context, sys system.System, request request.CreateUser) error {
	hashedPassword, err := password.Hash(request.Password)
	if err != nil {
		return err
	}
	user := auth.User{
		Username:       request.Username,
		HashedPassword: hashedPassword,
	}
	database := sys.Get("database").(storage.Database)
	err = database.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}