package database

import (
	"context"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
)

type DB interface {
	CreateSubject(ctx context.Context, subject auth.Subject) error
	GetSubject(ctx context.Context, subjectID string) (auth.Subject, error)
}
