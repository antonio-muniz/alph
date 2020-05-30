package memory

import (
	"context"

	"github.com/antonio-muniz/alph/cmd/alph/internal/database"
	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
	"github.com/pkg/errors"
)

var ErrSubjectNotFound = errors.New("subject not found")

type db struct {
	subjects map[string]auth.Subject
}

func NewDB() database.DB {
	return &db{
		subjects: make(map[string]auth.Subject),
	}
}

func (db *db) CreateSubject(ctx context.Context, subject auth.Subject) error {
	db.subjects[subject.ID] = subject
	return nil
}

func (db *db) GetSubject(ctx context.Context, subjectID string) (auth.Subject, error) {
	subject, found := db.subjects[subjectID]
	if !found {
		return auth.Subject{}, ErrSubjectNotFound
	}
	return subject, nil
}
