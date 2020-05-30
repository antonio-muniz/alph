package storage

import (
	"context"
	"errors"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
)

var ErrSubjectNotFound = errors.New("subject not found")

type Database interface {
	CreateSubject(ctx context.Context, subject auth.Subject) error
	GetSubject(ctx context.Context, subjectID string) (auth.Subject, error)
}
