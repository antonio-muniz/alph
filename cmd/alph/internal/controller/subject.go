package controller

import (
	"context"

	"github.com/antonio-muniz/alph/cmd/alph/internal/database"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/request"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/sarulabs/di"
)

func CreateSubject(ctx context.Context, components di.Container, request request.CreateSubject) error {
	hashedPassword, err := password.Hash(request.Password)
	if err != nil {
		return err
	}
	subject := auth.Subject{
		ID:             request.SubjectID,
		HashedPassword: hashedPassword,
	}
	database := components.Get("database").(database.DB)
	err = database.CreateSubject(ctx, subject)
	if err != nil {
		return err
	}
	return nil
}
