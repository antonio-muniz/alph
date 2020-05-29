package memory

import (
	"context"

	"github.com/antonio-muniz/alph/pkg/models/auth"
	"github.com/antonio-muniz/alph/pkg/password"
	"github.com/pkg/errors"
)

const (
	correctSubjectID = "someone@example.org"
	correctPassword  = "123456"
)

var ErrSubjectNotFound = errors.New("subject not found")

type Database struct{}

func (db Database) GetSubject(ctx context.Context, subjectID string) (auth.Subject, error) {
	if subjectID != correctSubjectID {
		return auth.Subject{}, ErrSubjectNotFound
	}
	hashedCorrectPassword, err := password.Hash(correctPassword)
	if err != nil {
		return auth.Subject{}, errors.Wrap(err, "hashing correct password")
	}
	subject := auth.Subject{
		ID:             subjectID,
		HashedPassword: hashedCorrectPassword,
	}
	return subject, nil
}
